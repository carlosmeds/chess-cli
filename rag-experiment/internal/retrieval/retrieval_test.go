package retrieval

import (
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/chunking"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/config"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/embedding"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/indexing"
	"testing"
)

func TestLexicalFindsExactSymbol(t *testing.T) {
	chunks := []chunking.Chunk{{ID: "a", Source: "identity.md", Heading: "Identity", Content: "positionKey contains castling rights"}, {ID: "b", Source: "other.md", Heading: "Other", Content: "board game"}}
	p, _ := embedding.New("local-hash", 64)
	for i := range chunks {
		chunks[i].Vector, _ = p.Embed(chunks[i].Content)
	}
	c := config.Default()
	c.ResultLimit = 1
	c.TokenBudget = 100
	r, err := Search(indexing.Index{Provider: "local-hash", Chunks: chunks}, c, Lexical, "positionKey")
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Results) != 1 || r.Results[0].Chunk.ID != "a" {
		t.Fatalf("unexpected: %#v", r.Results)
	}
}
func TestHybridFindsConfiguredParaphrase(t *testing.T) {
	chunks := []chunking.Chunk{{ID: "a", Source: "en.md", Heading: "En passant", Content: "A captura en passant vale na resposta imediata."}, {ID: "b", Source: "other.md", Heading: "Other", Content: "tabuleiro comum"}}
	p, _ := embedding.New("local-hash", 64)
	for i := range chunks {
		chunks[i].Vector, _ = p.Embed(chunks[i].Heading + "\n" + chunks[i].Content)
	}
	c := config.Default()
	c.EmbeddingDims = 64
	c.ResultLimit = 1
	c.TokenBudget = 100
	r, err := Search(indexing.Index{Provider: "local-hash", Chunks: chunks}, c, Hybrid, "captura de passagem")
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Results) != 1 || r.Results[0].Chunk.ID != "a" {
		t.Fatalf("unexpected: %#v", r.Results)
	}
}

func TestSearchRejectsUnknownStrategy(t *testing.T) {
	_, err := Search(indexing.Index{}, config.Default(), Strategy("typo"), "query")
	if err == nil {
		t.Fatal("expected an error")
	}
}
