---
name: fix-bug
description: Corrija um defeito reproduzível no domínio ou na CLI com teste de regressão e mudança mínima; use quando há comportamento atual incorreto descrito.
---

# Corrigir defeito

Defeito relatado: `$ARGUMENTS`.

1. Leia o contexto roteado por `CLAUDE.md`, localize o fluxo e reproduza o defeito antes de alterar produção.
2. Registre a reprodução em um teste que falhe pelo motivo esperado.
3. Determine a causa raiz e aplique a menor correção segura, sem refatorar código não relacionado.
4. Valide o teste de regressão e comportamentos vizinhos, incluindo movimentos especiais ou términos quando afetados.
5. Execute `gofmt -w .` e `make check`; atualize somente documentação cujo comportamento mudou.

Use `code-explorer` apenas se o ponto da falha não for conhecido e `code-reviewer` para correções de alto risco. Não enfraqueça um teste para fazê-lo passar.

Exemplo: `/fix-bug o roque é permitido quando o rei atravessa uma casa atacada` ou `$fix-bug ...`.
