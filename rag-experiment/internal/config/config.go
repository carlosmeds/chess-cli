package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	CorpusPaths       []string
	DistractorPaths   []string
	IndexPath         string
	ReportPath        string
	EmbeddingProvider string
	EmbeddingDims     int
	CandidateLimit    int
	ResultLimit       int
	TokenBudget       int
	RRFK              int
	MinRecall         float64
	MinTokenReduction float64
}

func Default() Config {
	return Config{
		CorpusPaths:     []string{"docs/domain", "docs/ARCHITECTURE.md", "docs/DECISIONS.md", "docs/TESTING.md"},
		DistractorPaths: []string{"rag-experiment/distractors"}, IndexPath: "rag-experiment/data/index.json",
		ReportPath: "rag-experiment/data/report.md", EmbeddingProvider: "local-hash", EmbeddingDims: 256,
		CandidateLimit: 20, ResultLimit: 6, TokenBudget: 1400, RRFK: 60, MinRecall: .90, MinTokenReduction: .60,
	}
}

// Load parses the deliberately small YAML subset used by config.example.yaml.
func Load(path string) (Config, error) {
	c := Default()
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return c, nil
		}
		return c, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(strings.SplitN(s.Text(), "#", 2)[0])
		if line == "" || strings.HasPrefix(line, "-") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return c, fmt.Errorf("invalid config line %q", line)
		}
		key, value := strings.TrimSpace(parts[0]), strings.Trim(strings.TrimSpace(parts[1]), "\"'")
		switch key {
		case "embedding_provider":
			c.EmbeddingProvider = value
		case "embedding_dimensions":
			c.EmbeddingDims, err = strconv.Atoi(value)
		case "candidate_limit":
			c.CandidateLimit, err = strconv.Atoi(value)
		case "result_limit":
			c.ResultLimit, err = strconv.Atoi(value)
		case "token_budget":
			c.TokenBudget, err = strconv.Atoi(value)
		case "rrf_k":
			c.RRFK, err = strconv.Atoi(value)
		case "minimum_recall":
			c.MinRecall, err = strconv.ParseFloat(value, 64)
		case "minimum_token_reduction":
			c.MinTokenReduction, err = strconv.ParseFloat(value, 64)
		case "index_path":
			c.IndexPath = value
		case "report_path":
			c.ReportPath = value
		default:
			return c, fmt.Errorf("unknown config key %q", key)
		}
		if err != nil {
			return c, fmt.Errorf("%s: %w", key, err)
		}
	}
	return c, s.Err()
}
