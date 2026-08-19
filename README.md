# Chess CLI

Jogo de xadrez local para dois jogadores, executado inteiramente no terminal e escrito somente com a biblioteca padrão do Go.

## Requisitos e execução

- Go 1.23 ou superior
- Terminal com suporte a UTF-8 para os símbolos das peças

```sh
go run ./cmd/chess
# ou
make run
```

Digite movimentos como `e2 e4`. Na promoção, acrescente `q`, `r`, `b` ou `n`, por exemplo `e7 e8 q`. O roque usa o movimento normal do rei (`e1 g1` ou `e1 c1`).

Exemplo curto:

```text
brancas> e2 e4
pretas> e7 e5
brancas> g1 f3
```

Comandos disponíveis: `help`, `board`, `restart` e `quit` (também aceita `exit`). O jogo valida movimentos, capturas, xeque e auto-xeque; detecta xeque-mate, afogamento e empate por repetição tripla; e suporta roque, promoção e en passant.

## Qualidade

```sh
make test
make test-race
make vet
make check
```

O domínio fica em `internal/chess`, a interação de terminal em `internal/cli` e o executável em `cmd/chess`. Consulte [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) e [docs/DOMAIN.md](docs/DOMAIN.md) para detalhes.
