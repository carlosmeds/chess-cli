package retrieval

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/example/chess-cli/rag-experiment/internal/chunking"
	"github.com/example/chess-cli/rag-experiment/internal/config"
	"github.com/example/chess-cli/rag-experiment/internal/embedding"
	"github.com/example/chess-cli/rag-experiment/internal/indexing"
	"github.com/example/chess-cli/rag-experiment/internal/tokens"
)

type Strategy string

const (
	Full    Strategy = "full"
	Lexical Strategy = "lexical"
	Vector  Strategy = "vector"
	Hybrid  Strategy = "hybrid"
)

type Result struct {
	Chunk chunking.Chunk
	Score float64
	Rank  int
}
type Response struct {
	Results       []Result
	ContextTokens int
	Duration      time.Duration
}

func Search(idx indexing.Index, c config.Config, strategy Strategy, query string) (Response, error) {
	start := time.Now()
	if strategy != Full && strategy != Lexical && strategy != Vector && strategy != Hybrid {
		return Response{}, fmt.Errorf("unknown retrieval strategy %q", strategy)
	}
	if strategy == Full {
		results := make([]Result, len(idx.Chunks))
		total := 0
		for i, ch := range idx.Chunks {
			results[i] = Result{Chunk: ch, Rank: i + 1}
			total += tokens.Estimate(ch.Heading + "\n" + ch.Content)
		}
		return Response{Results: results, ContextTokens: total, Duration: time.Since(start)}, nil
	}
	lex := bm25(idx.Chunks, query)
	var ranked []Result
	if strategy == Lexical {
		ranked = lex
	} else {
		provider, err := embedding.New(idx.Provider, c.EmbeddingDims)
		if err != nil {
			return Response{}, err
		}
		qv, err := provider.Embed(query)
		if err != nil {
			return Response{}, err
		}
		vec := make([]Result, 0, len(idx.Chunks))
		for _, ch := range idx.Chunks {
			score := embedding.Cosine(qv, ch.Vector)
			if score > 0 {
				vec = append(vec, Result{Chunk: ch, Score: score})
			}
		}
		sort.SliceStable(vec, func(i, j int) bool { return vec[i].Score > vec[j].Score })
		if strategy == Vector {
			ranked = vec
		} else {
			ranked = rrf(lex[:min(len(lex), c.CandidateLimit)], vec[:min(len(vec), c.CandidateLimit)], c.RRFK)
		}
	}
	selected := make([]Result, 0, c.ResultLimit)
	used := 0
	for _, r := range ranked {
		n := tokens.Estimate(r.Chunk.Heading + "\n" + r.Chunk.Content)
		if used+n > c.TokenBudget {
			continue
		}
		r.Rank = len(selected) + 1
		selected = append(selected, r)
		used += n
		if len(selected) >= c.ResultLimit {
			break
		}
	}
	return Response{Results: selected, ContextTokens: used, Duration: time.Since(start)}, nil
}

func bm25(chunks []chunking.Chunk, query string) []Result {
	q := unique(embedding.Terms(query))
	n := float64(len(chunks))
	dfs := map[string]int{}
	docs := make([][]string, len(chunks))
	avg := 0.0
	for i, ch := range chunks {
		docs[i] = embedding.Terms(ch.Heading + " " + strings.Join(ch.Symbols, " ") + " " + ch.Content)
		avg += float64(len(docs[i]))
		seen := map[string]bool{}
		for _, t := range docs[i] {
			if !seen[t] {
				dfs[t]++
				seen[t] = true
			}
		}
	}
	if n > 0 {
		avg /= n
	}
	var out []Result
	for i, terms := range docs {
		tf := map[string]int{}
		for _, t := range terms {
			tf[t]++
		}
		score := 0.0
		for _, term := range q {
			f := float64(tf[term])
			if f == 0 {
				continue
			}
			idf := math.Log(1 + (n-float64(dfs[term])+.5)/(float64(dfs[term])+.5))
			score += idf * (f * 2.2) / (f + 1.2*(.25+.75*float64(len(terms))/avg))
		}
		if score > 0 {
			out = append(out, Result{Chunk: chunks[i], Score: score})
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Score == out[j].Score {
			return out[i].Chunk.ID < out[j].Chunk.ID
		}
		return out[i].Score > out[j].Score
	})
	for i := range out {
		out[i].Rank = i + 1
	}
	return out
}

func rrf(a, b []Result, k int) []Result {
	scores := map[string]float64{}
	chunks := map[string]chunking.Chunk{}
	for _, list := range [][]Result{a, b} {
		for i, r := range list {
			scores[r.Chunk.ID] += 1 / float64(k+i+1)
			chunks[r.Chunk.ID] = r.Chunk
		}
	}
	out := make([]Result, 0, len(scores))
	for id, score := range scores {
		out = append(out, Result{Chunk: chunks[id], Score: score})
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Score == out[j].Score {
			return out[i].Chunk.ID < out[j].Chunk.ID
		}
		return out[i].Score > out[j].Score
	})
	return out
}
func unique(in []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, s := range in {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}
