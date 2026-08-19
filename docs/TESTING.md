# Testes

Os testes ficam junto aos pacotes. `internal/chess/game_test.go` cobre regras e estados, `repetition_test.go` cobre identidade e histórico, e os testes de `internal/cli` cobrem parser e sessão. Use tabelas para variações com a mesma preparação.

```sh
go test ./internal/chess -run TestCastling
go test ./internal/chess -run TestThreefoldRepetition
go test ./internal/cli -run TestParse
go test ./...
go test -race ./...
go vet ./...
make check
```

Cenários críticos: movimento e bloqueio de cada peça, captura aliada proibida, auto-xeque, ataques ao rei, mate/afogamento e pré-condições/efeitos das três regras especiais.

Para repetição tripla, cubra a segunda e a terceira ocorrências, um ciclo real de movimentos, jogador da vez, direitos de roque, en passant efetivamente capturável, jogadas inválidas, reinício e precedência dos estados de mate e afogamento. O teste do runner verifica também a mensagem apresentada no terminal.

Para criar posições sem CLI, use `NewEmptyBoard`, `Board.Set` e `NewGameWithBoard`. Inclua os dois reis; a ausência de um rei representa estado inválido e é tratada como xeque. Prefira coordenadas convertidas por `ParsePosition` a índices mágicos.

Uma correção de defeito deve começar com um teste de regressão que reproduza o comportamento. Regras do domínio são testadas sem parser; use o runner somente quando mensagem, comando ou coordenação da sessão fizer parte do requisito. Antes de concluir, execute primeiro o teste focado, depois `gofmt -w .` e `make check`, que inclui testes normais, race detector e vet.
