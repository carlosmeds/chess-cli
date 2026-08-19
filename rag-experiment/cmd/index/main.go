package main

import (
	"flag"
	"fmt"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/config"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/indexing"
	"log"
)

func main() {
	path := flag.String("config", "rag-experiment/config.example.yaml", "configuration file")
	flag.Parse()
	c, err := config.Load(*path)
	if err != nil {
		log.Fatal(err)
	}
	_, s, err := indexing.Build(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("chunks=%d reused=%d embedded=%d bytes=%d duration=%s\n", s.Chunks, s.Reused, s.Embedded, s.Bytes, s.Duration)
}
