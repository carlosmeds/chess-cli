# Domínio do xadrez

Este índice é a entrada para tarefas em `internal/chess`. Leia somente o tópico relacionado à mudança.

## Conceitos centrais

- [Glossário](glossary.md): nomes usados no código e na documentação.
- [Invariantes](invariants.md): propriedades que toda mudança deve preservar.
- [Estado da partida](game-state.md): tabuleiro, turno, regras temporárias, histórico e término.
- [Movimentos](moves/README.md): movimentos comuns e especiais.
- [Término](termination/README.md): xeque-mate, afogamento e repetição tripla.
- [Identidade e histórico de posição](position/identity.md): comparação e contagem de posições.

O domínio valida e altera a partida sem conhecer terminal ou comandos. A arquitetura geral está em [ARCHITECTURE.md](../ARCHITECTURE.md), os testes em [TESTING.md](../TESTING.md) e as decisões em [DECISIONS.md](../DECISIONS.md).
