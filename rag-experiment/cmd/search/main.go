package main

import (
	"flag"
	"fmt"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/config"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/indexing"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/retrieval"
	"log"
)

func main() {
	path := flag.String("config", "rag-experiment/config.example.yaml", "configuration file")
	query := flag.String("query", "", "query")
	strategy := flag.String("strategy", "hybrid", "full, lexical, vector, or hybrid")
	flag.Parse()
	if *query == "" {
		log.Fatal("-query is required")
	}
	c, err := config.Load(*path)
	if err != nil {
		log.Fatal(err)
	}
	idx, err := indexing.Load(c.IndexPath)
	if err != nil {
		log.Fatal("run make rag-index first: ", err)
	}
	r, err := retrieval.Search(idx, c, retrieval.Strategy(*strategy), *query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("strategy=%s chunks=%d estimated_tokens=%d latency=%s\n", *strategy, len(r.Results), r.ContextTokens, r.Duration)
	for _, x := range r.Results {
		fmt.Printf("\n%d. %s — %s (%.5f)\n%s\n", x.Rank, x.Chunk.Source, x.Chunk.Heading, x.Score, x.Chunk.Content)
	}
}
