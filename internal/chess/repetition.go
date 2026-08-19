package chess

// positionKey contains only the state that can change the legal moves from a
// position. It deliberately excludes Piece.Moved except where it determines a
// castling right.
type positionKey struct {
	squares   [64]uint8
	turn      Color
	castling  uint8
	enPassant int8
}

const noEnPassant int8 = -1

func (g *Game) recordPosition() uint8 {
	key := g.positionKey()
	g.positions[key]++
	return g.positions[key]
}

func (g *Game) positionKey() positionKey {
	key := positionKey{turn: g.turn, enPassant: noEnPassant}
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			piece := g.board.At(Position{File: file, Rank: rank})
			if !piece.Empty() {
				key.squares[rank*8+file] = uint8(piece.Type) + uint8(piece.Color)*uint8(King)
			}
		}
	}
	key.castling = g.castlingRights()
	if target, ok := g.effectiveEnPassantTarget(); ok {
		key.enPassant = int8(target.Rank*8 + target.File)
	}
	return key
}

func (g *Game) castlingRights() uint8 {
	var rights uint8
	for color := White; color <= Black; color++ {
		rank := 0
		shift := uint8(0)
		if color == Black {
			rank, shift = 7, 2
		}
		king := g.board.At(Position{File: 4, Rank: rank})
		if king.Type != King || king.Color != color || king.Moved {
			continue
		}
		for _, rook := range []struct {
			file int
			bit  uint8
		}{{7, 1 << shift}, {0, 2 << shift}} {
			piece := g.board.At(Position{File: rook.file, Rank: rank})
			if piece.Type == Rook && piece.Color == color && !piece.Moved {
				rights |= rook.bit
			}
		}
	}
	return rights
}

func (g *Game) effectiveEnPassantTarget() (Position, bool) {
	if g.enPassantTarget == nil {
		return Position{}, false
	}
	target := *g.enPassantTarget
	direction := 1
	if g.turn == Black {
		direction = -1
	}
	fromRank := target.Rank - direction
	for _, file := range []int{target.File - 1, target.File + 1} {
		from := Position{File: file, Rank: fromRank}
		piece := g.board.At(from)
		if piece.Type == Pawn && piece.Color == g.turn && g.validate(Move{From: from, To: target}) == nil {
			return target, true
		}
	}
	return Position{}, false
}
