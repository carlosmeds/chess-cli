package cli

import (
	"testing"

	"github.com/example/chess-cli/internal/chess"
)

func TestParseMoveAndPromotion(t *testing.T) {
	command, err := Parse("e7 e8 q")
	if err != nil {
		t.Fatal(err)
	}
	if command.Kind != MoveCommand || command.Move.Promotion != chess.Queen || command.Move.From.String() != "e7" {
		t.Fatalf("comando inesperado: %+v", command)
	}
}

func TestParseRejectsBadInput(t *testing.T) {
	for _, input := range []string{"", "e9 e4", "e2", "e7 e8 king"} {
		if _, err := Parse(input); err == nil {
			t.Errorf("Parse(%q) deveria falhar", input)
		}
	}
}
