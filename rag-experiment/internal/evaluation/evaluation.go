package evaluation

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/example/chess-cli/rag-experiment/internal/config"
	"github.com/example/chess-cli/rag-experiment/internal/embedding"
	"github.com/example/chess-cli/rag-experiment/internal/indexing"
	"github.com/example/chess-cli/rag-experiment/internal/retrieval"
)

type Question struct {
	ID               string   `json:"id"`
	Category         string   `json:"category"`
	Question         string   `json:"question"`
	RequiredConcepts []string `json:"required_concepts"`
	ExpectedStrategy string   `json:"expected_strategy"`
	Answerable       bool     `json:"answerable"`
}
type Row struct {
	Question, ID                            string
	Strategy                                retrieval.Strategy
	Sources                                 []string
	Recall, Precision, MRR, ConceptAccuracy float64
	Tokens                                  int
	Latency                                 time.Duration
	Refused, Grounded                       bool
	Answerable                              bool
	AdditionalSearch                        bool
	Irrelevant                              int
	Result                                  string
}
type Summary struct {
	Strategy                                                                                                        retrieval.Strategy
	Recall, Precision, MRR, ConceptAccuracy, GroundedRate, CorrectRefusalRate, AdditionalSearchRate, TokenReduction float64
	MeanTokens                                                                                                      int
	MeanLatency                                                                                                     time.Duration
}
type Report struct {
	Rows                                                 []Row
	Summaries                                            []Summary
	HybridFailures, LexicalBeatsVector, IrrelevantChunks []string
	CriteriaMet                                          bool
	FullTokens                                           int
	IndexBytes                                           int64
	IndexDuration                                        time.Duration
}

func Run(idx indexing.Index, c config.Config, questionsPath, sourcesPath string) (Report, error) {
	var qs []Question
	if err := readJSON(questionsPath, &qs); err != nil {
		return Report{}, err
	}
	expected := map[string][]string{}
	if err := readJSON(sourcesPath, &expected); err != nil {
		return Report{}, err
	}
	report := Report{IndexDuration: time.Duration(idx.BuildDurationNS)}
	if info, statErr := os.Stat(c.IndexPath); statErr == nil {
		report.IndexBytes = info.Size()
	}
	vectorRecall := map[string]float64{}
	lexicalRecall := map[string]float64{}
	for _, q := range qs {
		for _, strategy := range []retrieval.Strategy{retrieval.Full, retrieval.Lexical, retrieval.Vector, retrieval.Hybrid} {
			resp, err := retrieval.Search(idx, c, strategy, q.Question)
			if err != nil {
				return report, err
			}
			row := score(q, strategy, resp, expected[q.ID])
			if strategy == retrieval.Vector {
				vectorRecall[q.ID] = row.Recall
				continue
			}
			if strategy == retrieval.Lexical {
				lexicalRecall[q.ID] = row.Recall
			}
			report.Rows = append(report.Rows, row)
			if strategy == retrieval.Full && report.FullTokens == 0 {
				report.FullTokens = row.Tokens
			}
			if strategy == retrieval.Hybrid && q.Answerable && row.Recall < 1 {
				report.HybridFailures = append(report.HybridFailures, q.ID+": "+q.Question)
			}
			for _, r := range resp.Results {
				if r.Chunk.Distractor {
					report.IrrelevantChunks = append(report.IrrelevantChunks, q.ID+"/"+string(strategy)+": "+r.Chunk.Source+" — "+r.Chunk.Heading)
				}
			}
		}
	}
	for _, q := range qs {
		if lexicalRecall[q.ID] > vectorRecall[q.ID] {
			report.LexicalBeatsVector = append(report.LexicalBeatsVector, q.ID+": "+q.Question)
		}
	}
	for _, s := range []retrieval.Strategy{retrieval.Full, retrieval.Lexical, retrieval.Hybrid} {
		report.Summaries = append(report.Summaries, summarize(report.Rows, s, report.FullTokens))
	}
	hybrid := report.Summaries[2]
	answerableRecall := meanFiltered(report.Rows, retrieval.Hybrid, func(r Row) bool { return len(expected[r.ID]) > 0 }, func(r Row) float64 { return boolFloat(r.Recall > 0) })
	report.CriteriaMet = answerableRecall >= c.MinRecall && hybrid.CorrectRefusalRate == 1 && hybrid.TokenReduction >= c.MinTokenReduction && hybrid.ConceptAccuracy >= report.Summaries[0].ConceptAccuracy*.95
	report.IrrelevantChunks = unique(report.IrrelevantChunks)
	return report, nil
}

func score(q Question, strategy retrieval.Strategy, resp retrieval.Response, expected []string) Row {
	sources := []string{}
	relevant := 0
	first := 0
	context := strings.Builder{}
	for i, r := range resp.Results {
		sources = append(sources, r.Chunk.Source)
		context.WriteString(r.Chunk.Heading + "\n" + r.Chunk.Content + "\n")
		if contains(expected, r.Chunk.Source) {
			relevant++
			if first == 0 {
				first = i + 1
			}
		}
	}
	recall := 0.0
	if len(expected) > 0 {
		seen := map[string]bool{}
		for _, s := range sources {
			if contains(expected, s) {
				seen[s] = true
			}
		}
		recall = float64(len(seen)) / float64(len(expected))
	}
	precision := 0.0
	if len(resp.Results) > 0 {
		precision = float64(relevant) / float64(len(resp.Results))
	}
	mrr := 0.0
	if first > 0 {
		mrr = 1 / float64(first)
	}
	concepts := 1.0
	if len(q.RequiredConcepts) > 0 {
		hit := 0
		for _, concept := range q.RequiredConcepts {
			if conceptCovered(concept, context.String()) {
				hit++
			}
		}
		concepts = float64(hit) / float64(len(q.RequiredConcepts))
	}
	refused := !hasEvidence(q.Question, resp)
	grounded := (q.Answerable && !refused && relevant > 0) || (!q.Answerable && refused)
	result := "falha"
	if q.Answerable && recall == 1 && concepts == 1 && !refused {
		result = "ok"
	}
	if !q.Answerable && refused {
		result = "recusa correta"
	}
	if !q.Answerable && !refused {
		result = "resposta sem evidência"
	}
	return Row{Question: q.Question, ID: q.ID, Strategy: strategy, Sources: unique(sources), Recall: recall, Precision: precision, MRR: mrr, ConceptAccuracy: concepts, Tokens: resp.ContextTokens, Latency: resp.Duration, Refused: refused, Grounded: grounded, Answerable: q.Answerable, AdditionalSearch: q.Answerable && recall < 1, Irrelevant: countDistractors(resp), Result: result}
}
func hasEvidence(query string, response retrieval.Response) bool {
	q := unique(embedding.Terms(query))
	for _, result := range response.Results {
		if result.Chunk.Distractor {
			continue
		}
		terms := map[string]bool{}
		for _, t := range embedding.Terms(result.Chunk.Heading + " " + result.Chunk.Content) {
			terms[t] = true
		}
		hits := 0
		for _, t := range q {
			if len(t) > 3 && !genericEvidenceTerm[t] && terms[t] {
				hits++
			}
		}
		if hits >= 2 {
			return true
		}
	}
	return false
}

var genericEvidenceTerm = map[string]bool{"jogador": true, "partida": true, "jogada": true, "configurar": true, "controla": true}

func conceptCovered(concept, context string) bool {
	available := map[string]bool{}
	for _, term := range embedding.Terms(context) {
		available[term] = true
	}
	terms := embedding.Terms(concept)
	for _, term := range terms {
		if !available[term] {
			return false
		}
	}
	return len(terms) > 0
}
func countDistractors(r retrieval.Response) int {
	n := 0
	for _, x := range r.Results {
		if x.Chunk.Distractor {
			n++
		}
	}
	return n
}
func summarize(rows []Row, s retrieval.Strategy, full int) Summary {
	var n, refusalN, refusalOK int
	var recall, precision, mrr, concept, grounded, additional, tokens, latency float64
	for _, r := range rows {
		if r.Strategy != s {
			continue
		}
		n++
		if r.Answerable {
			recall += r.Recall
			precision += r.Precision
			mrr += r.MRR
		}
		concept += r.ConceptAccuracy
		grounded += boolFloat(r.Grounded)
		additional += boolFloat(r.AdditionalSearch)
		tokens += float64(r.Tokens)
		latency += float64(r.Latency)
		if !r.Answerable {
			refusalN++
			refusalOK += int(boolFloat(r.Refused))
		}
	}
	if n == 0 {
		return Summary{Strategy: s}
	}
	reduction := 0.0
	if full > 0 {
		reduction = 1 - (tokens/float64(n))/float64(full)
	}
	refusalRate := 0.0
	if refusalN > 0 {
		refusalRate = float64(refusalOK) / float64(refusalN)
	}
	answerableN := n - refusalN
	return Summary{Strategy: s, Recall: recall / float64(answerableN), Precision: precision / float64(answerableN), MRR: mrr / float64(answerableN), ConceptAccuracy: concept / float64(n), GroundedRate: grounded / float64(n), CorrectRefusalRate: refusalRate, AdditionalSearchRate: additional / float64(answerableN), MeanTokens: int(tokens / float64(n)), MeanLatency: time.Duration(latency / float64(n)), TokenReduction: reduction}
}
func WriteMarkdown(path string, r Report, c config.Config) error {
	var b strings.Builder
	b.WriteString("# Relatório do experimento RAG\n\nContagens de tokens são estimativas (`caracteres UTF-8 / 4`). A qualidade abaixo usa recuperação de fontes, cobertura de conceitos e um respondedor extrativo determinístico; não é julgamento por LLM.\n\n")
	fmt.Fprintf(&b, "Índice: %d bytes; indexação: %s.\n\n", r.IndexBytes, r.IndexDuration)
	b.WriteString("| Pergunta | Estratégia | Fontes recuperadas | Recall@K | Tokens | Resultado |\n|---|---|---|---:|---:|---|\n")
	for _, row := range r.Rows {
		fmt.Fprintf(&b, "| %s | %s | %s | %.2f | %d | %s |\n", escape(row.ID+" — "+row.Question), row.Strategy, escape(strings.Join(row.Sources, "<br>")), row.Recall, row.Tokens, row.Result)
	}
	b.WriteString("\n## Médias por estratégia\n\n| Estratégia | Recall | Precision | MRR | Conceitos | Fundamentadas | Recusa correta | Busca adicional | Tokens | Redução | Latência |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, s := range r.Summaries {
		fmt.Fprintf(&b, "| %s | %.3f | %.3f | %.3f | %.3f | %.3f | %.3f | %.3f | %d | %.1f%% | %s |\n", s.Strategy, s.Recall, s.Precision, s.MRR, s.ConceptAccuracy, s.GroundedRate, s.CorrectRefusalRate, s.AdditionalSearchRate, s.MeanTokens, s.TokenReduction*100, s.MeanLatency)
	}
	list(&b, "Perguntas em que o RAG falhou", r.HybridFailures)
	list(&b, "Perguntas em que o lexical venceu o vetorial", r.LexicalBeatsVector)
	list(&b, "Chunks irrelevantes recuperados", r.IrrelevantChunks)
	b.WriteString("\n## Conclusão\n\n")
	if r.CriteriaMet {
		b.WriteString("Os limiares configurados foram atingidos. Recomendação experimental: manter e validar com um modelo real antes de qualquer adoção.\n")
	} else {
		b.WriteString("Os limiares configurados não foram atingidos. Recomendação experimental: ajustar ou abandonar; não adotar este RAG com base neste ensaio.\n")
	}
	return os.WriteFile(path, []byte(b.String()), 0644)
}
func list(b *strings.Builder, title string, items []string) {
	b.WriteString("\n## " + title + "\n\n")
	if len(items) == 0 {
		b.WriteString("- Nenhuma.\n")
		return
	}
	for _, x := range items {
		b.WriteString("- " + x + "\n")
	}
}
func readJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
func contains(items []string, s string) bool {
	for _, x := range items {
		if x == s {
			return true
		}
	}
	return false
}
func unique(items []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, x := range items {
		if !seen[x] {
			seen[x] = true
			out = append(out, x)
		}
	}
	sort.Strings(out)
	return out
}
func boolFloat(v bool) float64 {
	if v {
		return 1
	}
	return 0
}
func meanFiltered(rows []Row, s retrieval.Strategy, filter func(Row) bool, value func(Row) float64) float64 {
	sum := 0.0
	n := 0
	for _, r := range rows {
		if r.Strategy == s && filter(r) {
			sum += value(r)
			n++
		}
	}
	if n == 0 {
		return 0
	}
	return sum / float64(n)
}
func escape(s string) string { return strings.ReplaceAll(s, "|", "\\|") }
