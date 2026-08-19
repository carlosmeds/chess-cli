# Contexto do Chess CLI

Jogo de xadrez local para dois humanos no terminal, escrito em Go 1.23 somente com a biblioteca padrão. Suporta movimentos legais, xeque, xeque-mate, afogamento, roque, promoção, en passant e repetição tripla. IA, rede, relógio, persistência e GUI estão fora do escopo atual.

## Mapa do projeto

- `cmd/chess`: composição do executável.
- `internal/chess`: domínio, estado e regras; nunca depende da CLI.
- `internal/cli`: parsing, apresentação e sessão interativa.
- `docs/domain`: regras divididas por contexto.
- `docs`: arquitetura, decisões, testes e roadmap.
- `.claude/agents` e `.codex/agents`: agentes especializados.
- `.claude/skills`: procedimentos reutilizáveis; `.codex/skills` aponta para a mesma fonte no Codex.

## Roteamento de contexto

| Tipo de tarefa | Contexto inicial |
| --- | --- |
| Regra de xadrez | `docs/domain/README.md` |
| Movimento especial | `docs/domain/moves/` |
| Término da partida | `docs/domain/termination/` |
| Identidade de posição | `docs/domain/position/` |
| Alteração na CLI | `internal/cli` e `README.md` |
| Testes | `docs/TESTING.md` |
| Arquitetura | `docs/ARCHITECTURE.md` |

Busque símbolos e referências antes de abrir arquivos inteiros. Leia apenas o ramo de documentação aplicável.

## Comandos essenciais

Use `make run`, `make test`, `make test-race`, `make fmt`, `make vet` e `make check`. Para iteração, execute primeiro um teste focado, como `go test ./internal/chess -run TestCastling`.

## Invariantes arquiteturais

- `internal/chess` não depende de `internal/cli` nem faz I/O.
- Não há estado global mutável.
- Uma jogada é validada contra auto-xeque antes de alterar a partida.
- Somente jogadas válidas alteram estado e histórico.
- APIs e comportamentos existentes são preservados fora do escopo pedido.
- Não adicionar dependências externas sem necessidade e decisão registrada.

## Agentes e skills

Agentes: `code-explorer`, `chess-domain`, `test-engineer`, `code-reviewer` e `context-maintainer`. Use agentes apenas quando houver trabalho independente ou especialização útil; veja a política em `AGENTS.md`.

Skills disponíveis: `implement-domain-rule`, `fix-bug`, `add-cli-command`, `review-change` e `update-context`.

```text
# Claude Code
/implement-domain-rule empate por repetição tripla

# Codex CLI ou IDE
$implement-domain-rule empate por repetição tripla
```

## Checklist mínimo

- Alterar somente arquivos relacionados e evitar refatoração incidental.
- Adicionar ou ajustar testes e documentação afetados.
- Executar testes focados, `gofmt` e `make check`.
- Revisar o diff e informar arquivos, validações e limitações.
