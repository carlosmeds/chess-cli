# Testes

Os testes unitários ficam junto aos pacotes. `internal/chess/game_test.go` cobre regras e estados; `internal/cli/parser_test.go` cobre tradução de entrada. Use tabelas para variações com a mesma preparação.

```sh
go test ./internal/chess -run TestCastling
go test ./internal/cli -run TestParse
go test ./...
go test -race ./...
```

Cenários críticos: movimento e bloqueio de cada peça, captura aliada proibida, auto-xeque, ataques ao rei, mate/afogamento e pré-condições/efeitos das três regras especiais.

Para repetição tripla, cubra a segunda e a terceira ocorrências, um ciclo real de movimentos, jogador da vez, direitos de roque, en passant efetivamente capturável, jogadas inválidas, reinício e precedência dos estados de mate e afogamento. O teste do runner verifica também a mensagem apresentada no terminal.

Para criar posições sem CLI, use `NewEmptyBoard`, `Board.Set` e `NewGameWithBoard`. Inclua os dois reis; a ausência de um rei representa estado inválido e é tratada como xeque. Prefira coordenadas convertidas por `ParsePosition` a índices mágicos.
