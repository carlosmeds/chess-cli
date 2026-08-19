package chunking

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileSplitsSectionsAndKeepsADRSubheadings(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "DECISIONS.md")
	if err := os.WriteFile(path, []byte("# D\n\nintro\n## ADR-001 — X\n### Contexto\na\n### Decisão\nb\n## ADR-002 — Y\nc\n"), 0644); err != nil {
		t.Fatal(err)
	}
	chunks, err := File(path, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(chunks) != 3 {
		t.Fatalf("got %d chunks", len(chunks))
	}
	if chunks[1].Heading != "ADR-001 — X" || chunks[1].ContentType != "adr" {
		t.Fatalf("unexpected ADR chunk: %#v", chunks[1])
	}
}
