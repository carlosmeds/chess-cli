package chunking

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Chunk struct {
	ID, Source, Heading, ContentType, Domain, Status, ContentHash string
	Symbols                                                       []string
	Content                                                       string
	Distractor                                                    bool
	Vector                                                        []float64
}

var symbolRE = regexp.MustCompile("`([A-Za-z][A-Za-z0-9_.]*(?:\\([^`]*\\))?)`")

func Corpus(paths, distractors []string) ([]Chunk, error) {
	var out []Chunk
	for _, group := range []struct {
		paths      []string
		distractor bool
	}{{paths, false}, {distractors, true}} {
		files, err := markdownFiles(group.paths)
		if err != nil {
			return nil, err
		}
		for _, path := range files {
			chunks, err := File(path, group.distractor)
			if err != nil {
				return nil, err
			}
			out = append(out, chunks...)
		}
	}
	return out, nil
}

func markdownFiles(paths []string) ([]string, error) {
	var files []string
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		if !info.IsDir() {
			if strings.HasSuffix(path, ".md") {
				files = append(files, path)
			}
			continue
		}
		err = filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && strings.HasSuffix(p, ".md") {
				files = append(files, p)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Strings(files)
	return files, nil
}

func File(path string, distractor bool) ([]Chunk, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var title, heading string
	var body []string
	var chunks []Chunk
	flush := func() {
		content := strings.TrimSpace(strings.Join(body, "\n"))
		body = nil
		if content == "" {
			return
		}
		h := heading
		if h == "" {
			h = title
		}
		sum := sha256.Sum256([]byte(path + "\n" + h + "\n" + content))
		id := fmt.Sprintf("%x", sum[:8])
		matches := symbolRE.FindAllStringSubmatch(content, -1)
		symbols := make([]string, 0, len(matches))
		for _, m := range matches {
			symbols = append(symbols, strings.TrimSuffix(m[1], "()"))
		}
		chunks = append(chunks, Chunk{ID: id, Source: filepath.ToSlash(path), Heading: h, ContentType: contentType(path, h), Domain: domain(path), Symbols: symbols, Status: "accepted", ContentHash: fmt.Sprintf("%x", sum[:]), Content: content, Distractor: distractor})
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "# ") {
			flush()
			title = strings.TrimSpace(line[2:])
			heading = title
			continue
		}
		if strings.HasPrefix(line, "## ") {
			flush()
			heading = strings.TrimSpace(line[3:])
			continue
		}
		// ADR subheadings stay with their ADR instead of becoming tiny chunks.
		body = append(body, line)
	}
	flush()
	return chunks, s.Err()
}

func domain(path string) string {
	p := filepath.ToSlash(path)
	for _, d := range []string{"moves", "position", "termination"} {
		if strings.Contains(p, "/"+d+"/") {
			return d
		}
	}
	if strings.Contains(p, "docs/domain/") {
		return "domain"
	}
	if strings.Contains(p, "distractors") {
		return "distractor"
	}
	return "project"
}

func contentType(path, heading string) string {
	if strings.HasPrefix(heading, "ADR-") {
		return "adr"
	}
	if strings.Contains(path, "/moves/") {
		return "move_rule"
	}
	if strings.Contains(path, "/termination/") {
		return "termination_rule"
	}
	if strings.Contains(path, "/position/") {
		return "position_rule"
	}
	if strings.HasSuffix(path, "TESTING.md") {
		return "testing_convention"
	}
	return "documentation"
}
