package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunnerReportsThreefoldRepetition(t *testing.T) {
	input := strings.NewReader(strings.Join([]string{
		"g1 f3", "g8 f6", "f3 g1", "f6 g8",
		"g1 f3", "g8 f6", "f3 g1", "f6 g8",
		"quit",
	}, "\n"))
	var output bytes.Buffer
	if err := NewRunner(input, &output).Run(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "Empate por repetição tripla de posição.") {
		t.Fatalf("mensagem de empate ausente na saída:\n%s", output.String())
	}
}
