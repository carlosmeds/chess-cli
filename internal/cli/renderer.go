package cli

import (
	"fmt"
	"io"

	"github.com/carlosmeds/context-engineering-chess-lab/internal/chess"
)

func RenderBoard(w io.Writer, board chess.Board) {
	fmt.Fprintln(w, "")
	for rank := 7; rank >= 0; rank-- {
		fmt.Fprintf(w, "%d  ", rank+1)
		for file := 0; file < 8; file++ {
			piece := board.At(chess.NewPosition(file, rank))
			if piece.Empty() {
				fmt.Fprint(w, "· ")
			} else {
				fmt.Fprintf(w, "%s ", piece.Symbol())
			}
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w, "\n   a b c d e f g h")
}

func RenderHelp(w io.Writer) {
	fmt.Fprintln(w, `Comandos:
  e2 e4       move uma peça
  e7 e8 q     promove um peão (q, r, b ou n)
  board       mostra o tabuleiro
  restart     reinicia a partida
  help        mostra esta ajuda
  quit        encerra o jogo`)
}
