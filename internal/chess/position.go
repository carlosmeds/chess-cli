package chess

import (
	"fmt"
	"strings"
)

// Position identifies a square using zero-based file (a=0) and rank (1=0).
type Position struct {
	File int
	Rank int
}

func NewPosition(file, rank int) Position { return Position{File: file, Rank: rank} }

func ParsePosition(value string) (Position, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	if len(value) != 2 || value[0] < 'a' || value[0] > 'h' || value[1] < '1' || value[1] > '8' {
		return Position{}, fmt.Errorf("posição inválida %q: use uma coordenada entre a1 e h8", value)
	}
	return Position{File: int(value[0] - 'a'), Rank: int(value[1] - '1')}, nil
}

func (p Position) Valid() bool { return p.File >= 0 && p.File < 8 && p.Rank >= 0 && p.Rank < 8 }

func (p Position) String() string {
	if !p.Valid() {
		return "??"
	}
	return string([]byte{byte('a' + p.File), byte('1' + p.Rank)})
}
