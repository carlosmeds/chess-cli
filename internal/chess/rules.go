package chess

import (
	"fmt"
	"math"
)

func (g *Game) validatePieceMove(m Move, piece Piece) error {
	if piece.Type != Pawn && m.Promotion != NoPiece {
		return fmt.Errorf("somente peões podem ser promovidos")
	}
	df, dr := m.To.File-m.From.File, m.To.Rank-m.From.Rank
	absF, absR := abs(df), abs(dr)
	switch piece.Type {
	case Pawn:
		return g.validatePawn(m, piece, df, dr)
	case Knight:
		if !((absF == 1 && absR == 2) || (absF == 2 && absR == 1)) {
			return illegal(piece)
		}
	case Bishop:
		if absF != absR || !pathClear(g.board, m.From, m.To) {
			return illegal(piece)
		}
	case Rook:
		if (df != 0 && dr != 0) || !pathClear(g.board, m.From, m.To) {
			return illegal(piece)
		}
	case Queen:
		if !((absF == absR) || df == 0 || dr == 0) || !pathClear(g.board, m.From, m.To) {
			return illegal(piece)
		}
	case King:
		if absF == 2 && dr == 0 {
			return g.validateCastle(m, piece)
		}
		if absF > 1 || absR > 1 {
			return illegal(piece)
		}
	default:
		return fmt.Errorf("peça inválida")
	}
	return nil
}

func (g *Game) validatePawn(m Move, piece Piece, df, dr int) error {
	direction, start, last := 1, 1, 7
	if piece.Color == Black {
		direction, start, last = -1, 6, 0
	}
	target := g.board.At(m.To)
	valid := false
	if df == 0 && dr == direction && target.Empty() {
		valid = true
	}
	if df == 0 && dr == 2*direction && m.From.Rank == start && target.Empty() {
		between := Position{File: m.From.File, Rank: m.From.Rank + direction}
		valid = g.board.At(between).Empty()
	}
	if abs(df) == 1 && dr == direction {
		if !target.Empty() && target.Color != piece.Color {
			valid = true
		}
		if target.Empty() && g.enPassantTarget != nil && *g.enPassantTarget == m.To {
			valid = true
		}
	}
	if !valid {
		return illegal(piece)
	}
	if m.To.Rank == last {
		if m.Promotion != Queen && m.Promotion != Rook && m.Promotion != Bishop && m.Promotion != Knight {
			return fmt.Errorf("promoção exige q, r, b ou n")
		}
	} else if m.Promotion != NoPiece {
		return fmt.Errorf("promoção só é permitida na última fileira")
	}
	return nil
}

func (g *Game) validateCastle(m Move, king Piece) error {
	if king.Moved || isInCheck(g.board, king.Color) {
		return fmt.Errorf("roque não permitido: rei já moveu ou está em xeque")
	}
	homeRank := 0
	if king.Color == Black {
		homeRank = 7
	}
	if m.From != (Position{File: 4, Rank: homeRank}) {
		return fmt.Errorf("roque não permitido nesta posição")
	}
	rookFile := 7
	step := 1
	if m.To.File == 2 {
		rookFile, step = 0, -1
	} else if m.To.File != 6 {
		return illegal(king)
	}
	rook := g.board.At(Position{File: rookFile, Rank: homeRank})
	if rook.Type != Rook || rook.Color != king.Color || rook.Moved {
		return fmt.Errorf("roque não permitido: torre indisponível")
	}
	for file := m.From.File + step; file != rookFile; file += step {
		if !g.board.At(Position{File: file, Rank: homeRank}).Empty() {
			return fmt.Errorf("roque não permitido: caminho bloqueado")
		}
	}
	for _, file := range []int{m.From.File + step, m.From.File + 2*step} {
		if squareAttacked(g.board, Position{File: file, Rank: homeRank}, king.Color.Opponent()) {
			return fmt.Errorf("roque não permitido através de casa atacada")
		}
	}
	return nil
}

func (g *Game) applyUnchecked(m Move) {
	piece := g.board.At(m.From)
	// En passant removes the pawn behind the empty target square.
	if piece.Type == Pawn && m.From.File != m.To.File && g.board.At(m.To).Empty() {
		g.board.Set(Position{File: m.To.File, Rank: m.From.Rank}, Piece{})
	}
	// Castling moves the rook together with the king.
	if piece.Type == King && abs(m.To.File-m.From.File) == 2 {
		if m.To.File == 6 {
			g.board.move(Position{File: 7, Rank: m.From.Rank}, Position{File: 5, Rank: m.From.Rank})
		} else {
			g.board.move(Position{File: 0, Rank: m.From.Rank}, Position{File: 3, Rank: m.From.Rank})
		}
	}
	g.board.move(m.From, m.To)
	if piece.Type == Pawn && (m.To.Rank == 0 || m.To.Rank == 7) {
		promoted := g.board.At(m.To)
		promoted.Type = m.Promotion
		g.board.Set(m.To, promoted)
	}
	g.enPassantTarget = nil
	if piece.Type == Pawn && abs(m.To.Rank-m.From.Rank) == 2 {
		target := Position{File: m.From.File, Rank: (m.From.Rank + m.To.Rank) / 2}
		g.enPassantTarget = &target
	}
}

func isInCheck(board Board, color Color) bool {
	king, found := board.findKing(color)
	return !found || squareAttacked(board, king, color.Opponent())
}

func squareAttacked(board Board, target Position, by Color) bool {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			from := Position{File: file, Rank: rank}
			p := board.At(from)
			if p.Empty() || p.Color != by {
				continue
			}
			df, dr := target.File-file, target.Rank-rank
			switch p.Type {
			case Pawn:
				dir := 1
				if by == Black {
					dir = -1
				}
				if abs(df) == 1 && dr == dir {
					return true
				}
			case Knight:
				if (abs(df) == 1 && abs(dr) == 2) || (abs(df) == 2 && abs(dr) == 1) {
					return true
				}
			case Bishop:
				if abs(df) == abs(dr) && pathClear(board, from, target) {
					return true
				}
			case Rook:
				if (df == 0 || dr == 0) && pathClear(board, from, target) {
					return true
				}
			case Queen:
				if (abs(df) == abs(dr) || df == 0 || dr == 0) && pathClear(board, from, target) {
					return true
				}
			case King:
				if abs(df) <= 1 && abs(dr) <= 1 {
					return true
				}
			}
		}
	}
	return false
}

func pathClear(board Board, from, to Position) bool {
	df, dr := sign(to.File-from.File), sign(to.Rank-from.Rank)
	for p := (Position{File: from.File + df, Rank: from.Rank + dr}); p != to; p = (Position{File: p.File + df, Rank: p.Rank + dr}) {
		if !board.At(p).Empty() {
			return false
		}
	}
	return true
}

func illegal(piece Piece) error { return fmt.Errorf("movimento inválido para %s", piece.Symbol()) }
func abs(value int) int         { return int(math.Abs(float64(value))) }
func sign(value int) int {
	if value < 0 {
		return -1
	}
	if value > 0 {
		return 1
	}
	return 0
}
