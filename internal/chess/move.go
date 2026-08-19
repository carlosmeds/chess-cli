package chess

type Move struct {
	From      Position
	To        Position
	Promotion PieceType
}

func (m Move) String() string { return m.From.String() + " " + m.To.String() }
