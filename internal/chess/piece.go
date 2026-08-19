package chess

type Color uint8

const (
	White Color = iota
	Black
)

func (c Color) Opponent() Color {
	if c == White {
		return Black
	}
	return White
}

func (c Color) String() string {
	if c == White {
		return "brancas"
	}
	return "pretas"
}

type PieceType uint8

const (
	NoPiece PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Type  PieceType
	Color Color
	Moved bool
}

func (p Piece) Empty() bool { return p.Type == NoPiece }

func (p Piece) Symbol() string {
	white := map[PieceType]string{Pawn: "♙", Knight: "♘", Bishop: "♗", Rook: "♖", Queen: "♕", King: "♔"}
	black := map[PieceType]string{Pawn: "♟", Knight: "♞", Bishop: "♝", Rook: "♜", Queen: "♛", King: "♚"}
	if p.Color == White {
		return white[p.Type]
	}
	return black[p.Type]
}
