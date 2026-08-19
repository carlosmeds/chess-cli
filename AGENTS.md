# Instruções para agentes Codex

## Contexto e alterações

1. Leia `CLAUDE.md` como mapa, depois somente a documentação indicada para a tarefa.
2. Use `rg` para localizar símbolos e referências antes de abrir arquivos inteiros. Evite varreduras amplas sem necessidade.
3. Preserve APIs, comportamento existente e a separação `cmd` → `internal/cli` → `internal/chess`; o domínio não depende da CLI.
4. Não refatore fora do escopo nem introduza dependências externas sem necessidade registrada.
5. Execute primeiro testes focados, depois `gofmt -w .` e `make check`.
6. Atualize somente a documentação afetada. Regras detalhadas pertencem a `docs/domain`, não a `CLAUDE.md`.

## Delegação

- Não use agentes em tarefas pequenas e lineares.
- `code_explorer`: exploração somente leitura quando a localização do código não for conhecida.
- `chess_domain`: implementação de regras complexas e seus testes de domínio.
- `test_engineer`: casos extremos e cobertura; trabalhe em paralelo apenas quando não editar os mesmos arquivos do implementador.
- `code_reviewer`: revisão independente após a implementação, preferencialmente somente leitura.
- `context_maintainer`: documentação e consistência depois que o comportamento estiver validado.
- Agentes de leitura podem atuar em paralelo. Agentes que editam o mesmo pacote ou documento devem atuar sequencialmente.
- Não crie um agente por peça ou por regra individual. Solicite subagentes apenas para tarefas independentes, espere todos os solicitados e consolide os resultados no agente principal.

As definições ficam em `.codex/agents`; as equivalentes do Claude Code ficam em `.claude/agents`.

Para uma regra complexa, o fluxo de referência é: agente principal lê o mapa; `code_explorer` localiza o contexto; `chess_domain` implementa; `test_engineer` verifica casos extremos sem disputar arquivos; `code_reviewer` revisa; `context_maintainer` atualiza documentos após validação; agente principal executa as verificações e consolida. Adapte ou omita etapas em tarefas triviais.

## Skills

As skills canônicas ficam em `.claude/skills`; `.codex/skills` é um link para a mesma fonte, sem cópia do procedimento. Invoque com `/nome-da-skill argumentos` no Claude Code ou `$nome-da-skill argumentos` no Codex. Use `implement-domain-rule`, `fix-bug`, `add-cli-command`, `review-change` e `update-context` conforme o tipo da tarefa.

## Entrega

Informe arquivos alterados, testes executados e limitações restantes. Antes de concluir, confira o diff, links da documentação e afirmações sobre funcionalidades implementadas.
