package chess

import "testing"

func playMoves(t *testing.T, g *Game, moves ...Move) {
	t.Helper()
	for _, m := range moves {
		if err := g.Play(m); err != nil {
			t.Fatalf("%s: %v", m, err)
		}
	}
}

func knightCycle() []Move {
	return []Move{
		move("g1", "f3"), move("g8", "f6"),
		move("f3", "g1"), move("f6", "g8"),
	}
}

func TestThreefoldRepetitionWithKnightCycle(t *testing.T) {
	g := NewGame()
	playMoves(t, g, knightCycle()...)
	if g.Status() != InProgress {
		t.Fatalf("duas ocorrências encerraram a partida: status=%v", g.Status())
	}
	playMoves(t, g, knightCycle()...)
	if g.Status() != ThreefoldRepetition {
		t.Fatalf("esperado empate por repetição tripla, status=%v", g.Status())
	}
}

func TestPositionIdentityIncludesRelevantState(t *testing.T) {
	base := NewEmptyBoard()
	base.Set(pos("e1"), Piece{Type: King, Color: White})
	base.Set(pos("h1"), Piece{Type: Rook, Color: White})
	base.Set(pos("e8"), Piece{Type: King, Color: Black})

	withoutCastling := base
	king := withoutCastling.At(pos("e1"))
	king.Moved = true
	withoutCastling.Set(pos("e1"), king)

	tests := []struct {
		name string
		a, b *Game
	}{
		{"jogador da vez", NewGameWithBoard(base, White), NewGameWithBoard(base, Black)},
		{"direitos de roque", NewGameWithBoard(base, White), NewGameWithBoard(withoutCastling, White)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.a.positionKey() == tc.b.positionKey() {
				t.Fatal("posições deveriam ser distintas")
			}
		})
	}
}

func TestPositionIdentityIncludesOnlyEffectiveEnPassant(t *testing.T) {
	board := NewEmptyBoard()
	board.Set(pos("e1"), Piece{Type: King, Color: White})
	board.Set(pos("e8"), Piece{Type: King, Color: Black})
	board.Set(pos("e5"), Piece{Type: Pawn, Color: White})
	board.Set(pos("d5"), Piece{Type: Pawn, Color: Black})

	withoutTarget := NewGameWithBoard(board, White)
	withTarget := NewGameWithBoard(board, White)
	target := pos("d6")
	withTarget.enPassantTarget = &target
	if withoutTarget.positionKey() == withTarget.positionKey() {
		t.Fatal("possibilidade efetiva de en passant não diferenciou as posições")
	}

	board.Set(pos("e5"), Piece{})
	withoutCapturer := NewGameWithBoard(board, White)
	withoutCapturer.enPassantTarget = &target
	withoutSameBoardTarget := NewGameWithBoard(board, White)
	if withoutCapturer.positionKey() != withoutSameBoardTarget.positionKey() {
		t.Fatal("alvo sem captura efetiva alterou a identidade")
	}
}

func TestInvalidMoveDoesNotCountPosition(t *testing.T) {
	g := NewGame()
	initial := g.positionKey()
	if err := g.Play(move("e2", "e5")); err == nil {
		t.Fatal("jogada inválida foi aceita")
	}
	if got := g.positions[initial]; got != 1 {
		t.Fatalf("jogada inválida alterou a contagem: %d", got)
	}
}

func TestRestartResetsRepetitionHistory(t *testing.T) {
	g := NewGame()
	playMoves(t, g, knightCycle()...)
	initial := NewGame().positionKey()
	if g.positions[initial] != 2 {
		t.Fatal("preparação não registrou a segunda ocorrência")
	}

	g.Restart()
	if g.Status() != InProgress || len(g.positions) != 1 || g.positions[initial] != 1 {
		t.Fatalf("reinício não restaurou o histórico inicial: status=%v posições=%v", g.Status(), g.positions)
	}
}

func TestRepetitionDoesNotOverrideTerminalStatus(t *testing.T) {
	tests := []struct {
		name string
		game *Game
		last Move
		want Status
	}{
		{
			name: "xeque-mate",
			game: func() *Game {
				g := NewGame()
				playMoves(t, g, move("f2", "f3"), move("e7", "e5"), move("g2", "g4"))
				return g
			}(),
			last: move("d8", "h4"),
			want: Checkmate,
		},
		{
			name: "afogamento",
			game: func() *Game {
				b := NewEmptyBoard()
				b.Set(pos("a8"), Piece{Type: King, Color: Black})
				b.Set(pos("c6"), Piece{Type: King, Color: White})
				b.Set(pos("b6"), Piece{Type: Queen, Color: White})
				return NewGameWithBoard(b, White)
			}(),
			last: move("b6", "c7"),
			want: Stalemate,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := *tc.game
			result.applyUnchecked(tc.last)
			result.turn = result.turn.Opponent()
			tc.game.positions[result.positionKey()] = 2
			playMoves(t, tc.game, tc.last)
			if tc.game.Status() != tc.want {
				t.Fatalf("status terminal sobrescrito: esperado=%v obtido=%v", tc.want, tc.game.Status())
			}
		})
	}
}
