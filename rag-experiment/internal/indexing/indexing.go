package indexing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/chunking"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/config"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/embedding"
)

type Index struct {
	Version         int              `json:"version"`
	Provider        string           `json:"provider"`
	Dimensions      int              `json:"dimensions"`
	CreatedAt       time.Time        `json:"created_at"`
	BuildDurationNS int64            `json:"build_duration_ns"`
	Chunks          []chunking.Chunk `json:"chunks"`
}

type Stats struct {
	Chunks, Reused, Embedded, Bytes int
	Duration                        time.Duration
}

func Build(c config.Config) (Index, Stats, error) {
	start := time.Now()
	chunks, err := chunking.Corpus(c.CorpusPaths, c.DistractorPaths)
	if err != nil {
		return Index{}, Stats{}, err
	}
	provider, err := embedding.New(c.EmbeddingProvider, c.EmbeddingDims)
	if err != nil {
		return Index{}, Stats{}, err
	}
	old, _ := Load(c.IndexPath)
	cached := map[string][]float64{}
	if old.Provider == provider.Name() && old.Dimensions == c.EmbeddingDims {
		for _, ch := range old.Chunks {
			cached[ch.ContentHash] = ch.Vector
		}
	}
	stats := Stats{Chunks: len(chunks)}
	for i := range chunks {
		if vector, ok := cached[chunks[i].ContentHash]; ok {
			chunks[i].Vector = vector
			stats.Reused++
			continue
		}
		chunks[i].Vector, err = provider.Embed(chunks[i].Heading + "\n" + chunks[i].Content)
		if err != nil {
			return Index{}, stats, err
		}
		if len(chunks[i].Vector) > 0 {
			stats.Embedded++
		}
	}
	idx := Index{Version: 1, Provider: provider.Name(), Dimensions: c.EmbeddingDims, CreatedAt: time.Now().UTC(), Chunks: chunks}
	idx.BuildDurationNS = time.Since(start).Nanoseconds()
	data, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return Index{}, stats, err
	}
	if err = os.MkdirAll(filepath.Dir(c.IndexPath), 0755); err != nil {
		return Index{}, stats, err
	}
	tmp := c.IndexPath + ".tmp"
	if err = os.WriteFile(tmp, data, 0644); err != nil {
		return Index{}, stats, err
	}
	if err = os.Rename(tmp, c.IndexPath); err != nil {
		return Index{}, stats, err
	}
	stats.Bytes = len(data)
	stats.Duration = time.Since(start)
	return idx, stats, nil
}

func Load(path string) (Index, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Index{}, err
	}
	var idx Index
	if err = json.Unmarshal(data, &idx); err != nil {
		return Index{}, fmt.Errorf("decode index: %w", err)
	}
	return idx, nil
}
