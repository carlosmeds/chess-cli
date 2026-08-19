package embedding

import (
	"fmt"
	"hash/fnv"
	"math"
	"strings"
	"unicode"
)

type Provider interface {
	Name() string
	Embed(string) ([]float64, error)
}

func New(name string, dimensions int) (Provider, error) {
	switch name {
	case "local-hash":
		return LocalHash{Dimensions: dimensions}, nil
	case "none", "":
		return Disabled{}, nil
	default:
		return nil, fmt.Errorf("unknown embedding provider %q", name)
	}
}

type Disabled struct{}

func (Disabled) Name() string                    { return "none" }
func (Disabled) Embed(string) ([]float64, error) { return nil, nil }

// LocalHash is a deterministic feature-hashing baseline, not a learned model.
// Canonical concepts make paraphrases comparable without external services.
type LocalHash struct{ Dimensions int }

func (l LocalHash) Name() string { return "local-hash" }
func (l LocalHash) Embed(text string) ([]float64, error) {
	if l.Dimensions < 32 {
		return nil, fmt.Errorf("embedding dimensions must be at least 32")
	}
	v := make([]float64, l.Dimensions)
	words := normalize(text)
	for i, word := range words {
		add(v, word, 1)
		if i+1 < len(words) {
			add(v, word+"_"+words[i+1], .6)
		}
	}
	var norm float64
	for _, x := range v {
		norm += x * x
	}
	norm = math.Sqrt(norm)
	if norm > 0 {
		for i := range v {
			v[i] /= norm
		}
	}
	return v, nil
}

var concepts = map[string]string{
	"rei afogado": "afogamento", "sem jogada legal": "sem_movimento", "não possui jogada legal": "sem_movimento",
	"nao possui jogada legal": "sem_movimento", "empate por repetição": "repeticao_tripla", "empate por repeticao": "repeticao_tripla",
	"três vezes": "repeticao_tripla", "tres vezes": "repeticao_tripla", "mesma configuração": "identidade_posicao",
	"mesma configuracao": "identidade_posicao", "posição igual": "identidade_posicao", "posicao igual": "identidade_posicao",
	"captura de passagem": "en_passant", "passagem do rei": "roque", "troca do peão": "promocao", "troca do peao": "promocao",
	"deixar o próprio rei atacado": "auto_xeque", "deixar o proprio rei atacado": "auto_xeque",
	"não deixa o próprio rei em xeque": "auto_xeque", "nao deixa o proprio rei em xeque": "auto_xeque",
	"direitos de roque": "direito_roque", "mover uma torre e devolvê-la": "direito_roque",
	"camada de terminal": "cli", "linha de comando": "cli", "núcleo do jogo": "dominio", "nucleo do jogo": "dominio",
	"matriz de casas": "tabuleiro", "posição estrutural": "positionkey", "posicao estrutural": "positionkey",
}

func normalize(text string) []string {
	s := strings.ToLower(text)
	for phrase, canonical := range concepts {
		s = strings.ReplaceAll(s, phrase, " "+canonical+" ")
	}
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteByte(' ')
		}
	}
	fields := strings.Fields(b.String())
	out := fields[:0]
	for _, f := range fields {
		if !stop[f] {
			out = append(out, stem(f))
		}
	}
	return out
}

var stop = map[string]bool{"a": true, "o": true, "as": true, "os": true, "de": true, "da": true, "do": true, "das": true, "dos": true, "e": true, "em": true, "um": true, "uma": true, "para": true, "por": true, "que": true, "como": true, "qual": true, "quais": true, "quando": true, "no": true, "na": true, "nos": true, "nas": true, "se": true}

func stem(s string) string {
	for _, suffix := range []string{"mente", "ções", "cao", "ção", "ados", "adas", "ido", "ida", "s"} {
		if len([]rune(s)) > len([]rune(suffix))+3 && strings.HasSuffix(s, suffix) {
			return strings.TrimSuffix(s, suffix)
		}
	}
	return s
}
func add(v []float64, feature string, weight float64) {
	h := fnv.New64a()
	_, _ = h.Write([]byte(feature))
	v[int(h.Sum64()%uint64(len(v)))] += weight
}

func Cosine(a, b []float64) float64 {
	if len(a) == 0 || len(a) != len(b) {
		return 0
	}
	var sum float64
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func Terms(text string) []string { return normalize(text) }
