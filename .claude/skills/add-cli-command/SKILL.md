---
name: add-cli-command
description: Adicione ou altere um comando textual do Chess CLI mantendo parsing, sessão e apresentação fora do domínio.
---

# Adicionar comando da CLI

Comando solicitado: `$ARGUMENTS`.

- Localize `CommandKind`, `Parse`, `Runner` e os testes relacionados antes de editar.
- Mantenha parsing, validação textual e mensagens em `internal/cli`; exponha no domínio somente operações de estado reutilizáveis.
- Preserve comandos existentes e rejeite formatos inválidos com mensagens claras.
- Adicione testes do parser e, quando houver fluxo de sessão, do runner com buffers.
- Atualize ajuda e `README.md` quando a interface pública mudar.
- Execute testes focados de `internal/cli`, `gofmt -w .` e `make check`.

Não adicione dependência de `internal/chess` para `internal/cli`.
