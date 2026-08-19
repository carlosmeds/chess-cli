package chess

type Board struct{ squares [8][8]Piece }

func NewEmptyBoard() Board { return Board{} }

func NewBoard() Board {
	b := NewEmptyBoard()
	back := [...]PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for file, kind := range back {
		b.squares[0][file] = Piece{Type: kind, Color: White}
		b.squares[1][file] = Piece{Type: Pawn, Color: White}
		b.squares[6][file] = Piece{Type: Pawn, Color: Black}
		b.squares[7][file] = Piece{Type: kind, Color: Black}
	}
	return b
}

func (b Board) At(p Position) Piece {
	if !p.Valid() {
		return Piece{}
	}
	return b.squares[p.Rank][p.File]
}

// Set is useful for composing explicit positions in tests and tools.
func (b *Board) Set(p Position, piece Piece) {
	if p.Valid() {
		b.squares[p.Rank][p.File] = piece
	}
}

func (b *Board) move(from, to Position) {
	p := b.At(from)
	p.Moved = true
	b.Set(to, p)
	b.Set(from, Piece{})
}

func (b Board) findKing(color Color) (Position, bool) {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			p := Position{File: file, Rank: rank}
			piece := b.At(p)
			if piece.Type == King && piece.Color == color {
				return p, true
			}
		}
	}
	return Position{}, false
}
