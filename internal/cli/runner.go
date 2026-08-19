package cli

import (
	"bufio"
	"fmt"
	"io"

	"github.com/carlosmeds/context-engineering-chess-lab/internal/chess"
)

type Runner struct {
	in   io.Reader
	out  io.Writer
	game *chess.Game
}

func NewRunner(in io.Reader, out io.Writer) *Runner {
	return &Runner{in: in, out: out, game: chess.NewGame()}
}

func (r *Runner) Run() error {
	fmt.Fprintln(r.out, "Xadrez CLI — digite 'help' para ajuda.")
	RenderBoard(r.out, r.game.Board())
	scanner := bufio.NewScanner(r.in)
	for {
		fmt.Fprintf(r.out, "\n%s> ", r.game.Turn())
		if !scanner.Scan() {
			fmt.Fprintln(r.out)
			return scanner.Err()
		}
		command, err := Parse(scanner.Text())
		if err != nil {
			fmt.Fprintf(r.out, "Erro: %v\n", err)
			continue
		}
		switch command.Kind {
		case HelpCommand:
			RenderHelp(r.out)
		case BoardCommand:
			RenderBoard(r.out, r.game.Board())
		case RestartCommand:
			r.game.Restart()
			fmt.Fprintln(r.out, "Partida reiniciada.")
			RenderBoard(r.out, r.game.Board())
		case QuitCommand:
			fmt.Fprintln(r.out, "Até a próxima!")
			return nil
		case MoveCommand:
			if err := r.game.Play(command.Move); err != nil {
				fmt.Fprintf(r.out, "Jogada inválida: %v\n", err)
				continue
			}
			RenderBoard(r.out, r.game.Board())
			switch r.game.Status() {
			case chess.Checkmate:
				fmt.Fprintf(r.out, "Xeque-mate! %s vencem. Use 'restart' ou 'quit'.\n", r.game.Turn().Opponent())
			case chess.Stalemate:
				fmt.Fprintln(r.out, "Afogamento: empate. Use 'restart' ou 'quit'.")
			case chess.ThreefoldRepetition:
				fmt.Fprintln(r.out, "Empate por repetição tripla de posição. Use 'restart' ou 'quit'.")
			default:
				if r.game.InCheck(r.game.Turn()) {
					fmt.Fprintf(r.out, "Xeque às %s!\n", r.game.Turn())
				}
			}
		}
	}
}
