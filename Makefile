.PHONY: run test test-race fmt vet check rag-index rag-search rag-eval rag-clean

run:
	go run ./cmd/chess

test:
	go test ./...

test-race:
	go test -race ./...

fmt:
	gofmt -w cmd internal

vet:
	go vet ./...

check:
	test -z "$$(gofmt -l cmd internal)"
	go test ./...
	go test -race ./...
	go vet ./...

rag-index:
	GOCACHE=/tmp/context-engineering-chess-lab-rag-go-cache go run ./rag-experiment/cmd/index

rag-search:
	GOCACHE=/tmp/context-engineering-chess-lab-rag-go-cache go run ./rag-experiment/cmd/search -query "$(QUERY)"

rag-eval:
	GOCACHE=/tmp/context-engineering-chess-lab-rag-go-cache go run ./rag-experiment/cmd/evaluate

rag-clean:
	find rag-experiment/data -type f ! -name .gitkeep -delete
