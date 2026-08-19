package chess

import "fmt"

type Status uint8

const (
	InProgress Status = iota
	Checkmate
	Stalemate
	ThreefoldRepetition
)

type Game struct {
	board           Board
	turn            Color
	status          Status
	enPassantTarget *Position
	positions       map[positionKey]uint8
}

func NewGame() *Game { return newGame(NewBoard(), White) }

// NewGameWithBoard creates a game at a deliberate position, primarily for tests.
func NewGameWithBoard(board Board, turn Color) *Game { return newGame(board, turn) }

func newGame(board Board, turn Color) *Game {
	g := &Game{board: board, turn: turn, positions: make(map[positionKey]uint8)}
	g.recordPosition()
	return g
}

func (g *Game) Board() Board             { return g.board }
func (g *Game) Turn() Color              { return g.turn }
func (g *Game) Status() Status           { return g.status }
func (g *Game) InCheck(color Color) bool { return isInCheck(g.board, color) }

func (g *Game) Restart() { *g = *NewGame() }

func (g *Game) Play(move Move) error {
	if g.status != InProgress {
		return fmt.Errorf("a partida já terminou")
	}
	if err := g.validate(move); err != nil {
		return err
	}
	g.applyUnchecked(move)
	g.turn = g.turn.Opponent()
	if !g.hasLegalMove(g.turn) {
		if g.InCheck(g.turn) {
			g.status = Checkmate
		} else {
			g.status = Stalemate
		}
	}
	if g.recordPosition() >= 3 && g.status == InProgress {
		g.status = ThreefoldRepetition
	}
	return nil
}

func (g *Game) validate(move Move) error {
	if !move.From.Valid() || !move.To.Valid() || move.From == move.To {
		return fmt.Errorf("origem e destino devem ser casas diferentes e válidas")
	}
	piece := g.board.At(move.From)
	if piece.Empty() {
		return fmt.Errorf("não há peça em %s", move.From)
	}
	if piece.Color != g.turn {
		return fmt.Errorf("é a vez das %s", g.turn)
	}
	target := g.board.At(move.To)
	if !target.Empty() && target.Color == piece.Color {
		return fmt.Errorf("%s está ocupada por uma peça da mesma cor", move.To)
	}
	if target.Type == King {
		return fmt.Errorf("o rei não pode ser capturado")
	}
	if err := g.validatePieceMove(move, piece); err != nil {
		return err
	}
	copy := *g
	copy.applyUnchecked(move)
	if isInCheck(copy.board, piece.Color) {
		return fmt.Errorf("a jogada deixaria o próprio rei em xeque")
	}
	return nil
}

func (g *Game) hasLegalMove(color Color) bool {
	old := g.turn
	g.turn = color
	defer func() { g.turn = old }()
	for fr := 0; fr < 8; fr++ {
		for ff := 0; ff < 8; ff++ {
			from := Position{File: ff, Rank: fr}
			if p := g.board.At(from); p.Empty() || p.Color != color {
				continue
			}
			for tr := 0; tr < 8; tr++ {
				for tf := 0; tf < 8; tf++ {
					to := Position{File: tf, Rank: tr}
					promotion := NoPiece
					if g.board.At(from).Type == Pawn && (tr == 0 || tr == 7) {
						promotion = Queen
					}
					if g.validate(Move{From: from, To: to, Promotion: promotion}) == nil {
						return true
					}
				}
			}
		}
	}
	return false
}
