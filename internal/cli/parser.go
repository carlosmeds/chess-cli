package cli

import (
	"fmt"
	"strings"

	"github.com/example/chess-cli/internal/chess"
)

type CommandKind uint8

const (
	MoveCommand CommandKind = iota
	HelpCommand
	BoardCommand
	RestartCommand
	QuitCommand
)

type Command struct {
	Kind CommandKind
	Move chess.Move
}

func Parse(input string) (Command, error) {
	fields := strings.Fields(strings.ToLower(input))
	if len(fields) == 0 {
		return Command{}, fmt.Errorf("digite uma jogada ou 'help'")
	}
	if len(fields) == 1 {
		switch fields[0] {
		case "help":
			return Command{Kind: HelpCommand}, nil
		case "board":
			return Command{Kind: BoardCommand}, nil
		case "restart":
			return Command{Kind: RestartCommand}, nil
		case "quit", "exit":
			return Command{Kind: QuitCommand}, nil
		}
	}
	if len(fields) < 2 || len(fields) > 3 {
		return Command{}, fmt.Errorf("formato inválido: use 'e2 e4' ou 'e7 e8 q'")
	}
	from, err := chess.ParsePosition(fields[0])
	if err != nil {
		return Command{}, err
	}
	to, err := chess.ParsePosition(fields[1])
	if err != nil {
		return Command{}, err
	}
	move := chess.Move{From: from, To: to}
	if len(fields) == 3 {
		promotions := map[string]chess.PieceType{"q": chess.Queen, "r": chess.Rook, "b": chess.Bishop, "n": chess.Knight}
		var ok bool
		move.Promotion, ok = promotions[fields[2]]
		if !ok {
			return Command{}, fmt.Errorf("promoção inválida: use q, r, b ou n")
		}
	}
	return Command{Kind: MoveCommand, Move: move}, nil
}
