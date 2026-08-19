.PHONY: run test test-race fmt vet check

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
