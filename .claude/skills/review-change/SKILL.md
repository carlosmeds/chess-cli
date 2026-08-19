---
name: review-change
description: Revise uma mudança concluída no Chess CLI em busca de defeitos, regressões, invariantes quebrados, testes ausentes e documentação divergente.
---

# Revisar mudança

Escopo indicado: `$ARGUMENTS`.

Use o agente `code-reviewer` em contexto isolado quando disponível. A revisão é somente leitura: examine o diff, abra apenas o contexto necessário e priorize problemas reais por severidade. Verifique correção, regressões, invariantes, movimentos especiais, condições de término, concorrência quando aplicável, testes ausentes e documentação divergente.

Entregue primeiro os achados com arquivo e linha, depois dúvidas e um resumo curto. Se não houver achados, diga explicitamente. Não modifique arquivos durante a revisão.
