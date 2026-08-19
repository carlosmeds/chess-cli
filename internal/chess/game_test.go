package chess

import "testing"

func pos(value string) Position {
	p, err := ParsePosition(value)
	if err != nil {
		panic(err)
	}
	return p
}
func move(from, to string) Move { return Move{From: pos(from), To: pos(to)} }

func TestInitialMovesAndTurn(t *testing.T) {
	g := NewGame()
	if err := g.Play(move("e2", "e4")); err != nil {
		t.Fatal(err)
	}
	if g.Turn() != Black {
		t.Fatal("turno não alternou")
	}
	if err := g.Play(move("e4", "e5")); err == nil {
		t.Fatal("permitiu jogar peça adversária")
	}
	if err := g.Play(move("e7", "e5")); err != nil {
		t.Fatal(err)
	}
}

func TestPieceMovementAndBlockedPath(t *testing.T) {
	tests := []struct {
		name, from, to string
		legal          bool
	}{
		{"cavalo", "g1", "f3", true}, {"bispo bloqueado", "c1", "h6", false},
		{"torre bloqueada", "a1", "a4", false}, {"peão para trás", "e2", "e1", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGame()
			err := g.Play(move(tc.from, tc.to))
			if (err == nil) != tc.legal {
				t.Fatalf("legal=%v, erro=%v", tc.legal, err)
			}
		})
	}
}

func TestCannotCaptureFriendlyPiece(t *testing.T) {
	g := NewGame()
	if err := g.Play(move("g1", "e2")); err == nil {
		t.Fatal("permitiu capturar peça da mesma cor")
	}
}

func TestDetectsCheck(t *testing.T) {
	b := NewEmptyBoard()
	b.Set(pos("e1"), Piece{Type: King, Color: White})
	b.Set(pos("e8"), Piece{Type: King, Color: Black})
	b.Set(pos("a1"), Piece{Type: Rook, Color: Black})
	g := NewGameWithBoard(b, White)
	if !g.InCheck(White) {
		t.Fatal("não detectou xeque")
	}
}

func TestCannotLeaveOwnKingInCheck(t *testing.T) {
	b := NewEmptyBoard()
	b.Set(pos("e1"), Piece{Type: King, Color: White})
	b.Set(pos("e2"), Piece{Type: Rook, Color: White})
	b.Set(pos("e8"), Piece{Type: Rook, Color: Black})
	b.Set(pos("a8"), Piece{Type: King, Color: Black})
	g := NewGameWithBoard(b, White)
	if err := g.Play(move("e2", "f2")); err == nil {
		t.Fatal("permitiu expor o próprio rei")
	}
}

func TestCastling(t *testing.T) {
	b := NewEmptyBoard()
	b.Set(pos("e1"), Piece{Type: King, Color: White})
	b.Set(pos("h1"), Piece{Type: Rook, Color: White})
	b.Set(pos("e8"), Piece{Type: King, Color: Black})
	g := NewGameWithBoard(b, White)
	if err := g.Play(move("e1", "g1")); err != nil {
		t.Fatal(err)
	}
	if g.Board().At(pos("f1")).Type != Rook || g.Board().At(pos("g1")).Type != King {
		t.Fatal("roque não moveu rei e torre")
	}
}

func TestCastlingThroughCheckIsRejected(t *testing.T) {
	b := NewEmptyBoard()
	b.Set(pos("e1"), Piece{Type: King, Color: White})
	b.Set(pos("h1"), Piece{Type: Rook, Color: White})
	b.Set(pos("f8"), Piece{Type: Rook, Color: Black})
	b.Set(pos("a8"), Piece{Type: King, Color: Black})
	g := NewGameWithBoard(b, White)
	if err := g.Play(move("e1", "g1")); err == nil {
		t.Fatal("permitiu roque através de xeque")
	}
}

func TestEnPassantAndImmediateWindow(t *testing.T) {
	g := NewGame()
	for _, m := range []Move{move("e2", "e4"), move("a7", "a6"), move("e4", "e5"), move("d7", "d5"), move("e5", "d6")} {
		if err := g.Play(m); err != nil {
			t.Fatalf("%s: %v", m, err)
		}
	}
	if !g.Board().At(pos("d5")).Empty() || g.Board().At(pos("d6")).Type != Pawn {
		t.Fatal("en passant incorreto")
	}

	g = NewGame()
	for _, m := range []Move{move("e2", "e4"), move("a7", "a6"), move("e4", "e5"), move("d7", "d5"), move("h2", "h3"), move("a6", "a5")} {
		if err := g.Play(m); err != nil {
			t.Fatal(err)
		}
	}
	if err := g.Play(move("e5", "d6")); err == nil {
		t.Fatal("permitiu en passant tardio")
	}
}

func TestPromotion(t *testing.T) {
	b := NewEmptyBoard()
	b.Set(pos("a1"), Piece{Type: King, Color: White})
	b.Set(pos("h8"), Piece{Type: King, Color: Black})
	b.Set(pos("e7"), Piece{Type: Pawn, Color: White})
	g := NewGameWithBoard(b, White)
	m := move("e7", "e8")
	m.Promotion = Knight
	if err := g.Play(m); err != nil {
		t.Fatal(err)
	}
	if g.Board().At(pos("e8")).Type != Knight {
		t.Fatal("promoção não aplicada")
	}
}

func TestCheckmateFoolsMate(t *testing.T) {
	g := NewGame()
	for _, m := range []Move{move("f2", "f3"), move("e7", "e5"), move("g2", "g4"), move("d8", "h4")} {
		if err := g.Play(m); err != nil {
			t.Fatal(err)
		}
	}
	if g.Status() != Checkmate || !g.InCheck(White) {
		t.Fatalf("esperado xeque-mate, status=%v", g.Status())
	}
}

func TestStalemate(t *testing.T) {
	b := NewEmptyBoard()
	b.Set(pos("a8"), Piece{Type: King, Color: Black})
	b.Set(pos("c6"), Piece{Type: King, Color: White})
	b.Set(pos("b6"), Piece{Type: Queen, Color: White})
	g := NewGameWithBoard(b, White)
	if err := g.Play(move("b6", "c7")); err != nil {
		t.Fatal(err)
	}
	if g.Status() != Stalemate || g.InCheck(Black) {
		t.Fatalf("esperado afogamento, status=%v check=%v", g.Status(), g.InCheck(Black))
	}
}
