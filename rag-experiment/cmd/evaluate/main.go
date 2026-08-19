package main

import (
	"flag"
	"fmt"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/config"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/evaluation"
	"github.com/carlosmeds/context-engineering-chess-lab/rag-experiment/internal/indexing"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path := flag.String("config", "rag-experiment/config.example.yaml", "configuration file")
	flag.Parse()
	c, err := config.Load(*path)
	if err != nil {
		log.Fatal(err)
	}
	idx, err := indexing.Load(c.IndexPath)
	if err != nil {
		log.Fatal("run make rag-index first: ", err)
	}
	r, err := evaluation.Run(idx, c, "rag-experiment/evals/questions.json", "rag-experiment/evals/expected-sources.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = os.MkdirAll(filepath.Dir(c.ReportPath), 0755); err != nil {
		log.Fatal(err)
	}
	if err = evaluation.WriteMarkdown(c.ReportPath, r, c); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("report=%s criteria_met=%t\n", c.ReportPath, r.CriteriaMet)
	for _, s := range r.Summaries {
		fmt.Printf("%s recall=%.3f precision=%.3f mrr=%.3f concepts=%.3f grounded=%.3f refusal=%.3f additional_search=%.3f tokens=%d reduction=%.1f%% latency=%s\n", s.Strategy, s.Recall, s.Precision, s.MRR, s.ConceptAccuracy, s.GroundedRate, s.CorrectRefusalRate, s.AdditionalSearchRate, s.MeanTokens, s.TokenReduction*100, s.MeanLatency)
	}
}
