---
name: implement-domain-rule
description: Implemente uma regra de xadrez no domínio quando a tarefa exigir estado, legalidade, movimentos especiais ou término; não use para mudanças exclusivas da CLI.
---

# Implementar regra de domínio

Regra solicitada: `$ARGUMENTS`.

1. Leia `CLAUDE.md`, localize símbolos com `rg` e siga somente o ramo aplicável de `docs/domain/README.md`.
2. Identifique estado, invariantes, interações com regras especiais e condições de término antes de editar.
3. Planeje a menor mudança que preserve a API e a independência de `internal/chess`.
4. Quando disponível e a regra for complexa, delegue a implementação ao agente `chess-domain` em contexto isolado. Não delegue tarefas pequenas e lineares.
5. Implemente a regra e testes de domínio; use CLI apenas para integração ou mensagem observável.
6. Execute testes focados, `gofmt -w .` e `make check`.
7. Depois da validação, atualize somente a documentação afetada e o roadmap se a funcionalidade estiver realmente concluída.
8. Para mudança não trivial, solicite revisão independente com `code-reviewer` ou a skill `review-change` e consolide os achados.

Não crie dependência do domínio para a CLI, não faça refatoração incidental e não execute commit ou push sem pedido.

Exemplos: `/implement-domain-rule empate por repetição tripla` no Claude Code; `$implement-domain-rule empate por repetição tripla` no Codex.
