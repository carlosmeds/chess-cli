# Contexto do Chess CLI

Jogo de xadrez local para dois humanos no terminal, em Go 1.23 e sem dependências externas. Inclui regras legais, xeque/mate/afogamento, roque, promoção e en passant. Não inclui IA, rede, relógio, persistência ou GUI.

## Arquitetura e mapa

- `cmd/chess`: composição e entrada do executável.
- `internal/chess`: domínio e regras; nunca depende da CLI.
- `internal/cli`: parser, renderização e loop interativo.
- `docs`: arquitetura, domínio, decisões, testes e roadmap.

Leia [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) para mudanças estruturais, [docs/DOMAIN.md](docs/DOMAIN.md) para regras, [docs/DECISIONS.md](docs/DECISIONS.md) para decisões e [docs/TESTING.md](docs/TESTING.md) para testes.

## Comandos e convenções

Use `make run`, `make test`, `make test-race`, `make fmt`, `make vet` e `make check`. Código idiomático, `gofmt`, erros claros em português e testes orientados a tabela quando apropriado. Interfaces só com benefício concreto.

Não quebre: domínio independente da CLI; ausência de estado global mutável; validação de auto-xeque antes de alterar a partida; biblioteca padrão como política; APIs existentes sem necessidade explícita.

## Checklist obrigatório

- Investigar apenas símbolos e arquivos relacionados.
- Atualizar testes e documentação afetados.
- Executar testes focados, `make fmt` e `make check`.
- Relatar arquivos alterados, verificações e limitações.
