package main

import (
	"fmt"
	"os"

	"github.com/carlosmeds/context-engineering-chess-lab/internal/cli"
)

func main() {
	if err := cli.NewRunner(os.Stdin, os.Stdout).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}
}
