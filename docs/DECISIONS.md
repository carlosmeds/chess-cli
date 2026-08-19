# Decisões arquiteturais

## ADR-001 — Separação entre domínio e CLI

### Contexto
Regras devem ser testáveis sem terminal.
### Decisão
`internal/chess` não faz I/O; `internal/cli` traduz e apresenta.
### Consequências
Testes montam partidas diretamente e outros frontends podem reutilizar o domínio.

## ADR-002 — Tabuleiro como matriz de valores

### Contexto
O tabuleiro tem tamanho fixo e peças pequenas.
### Decisão
Usar `[8][8]Piece`, sem ponteiros por casa.
### Consequências
Cópias para simulação são simples, baratas e livres de aliasing.

## ADR-003 — Movimento como valor explícito

### Contexto
Origem, destino e promoção precisam atravessar as camadas.
### Decisão
Usar `Move{From, To, Promotion}`; roque e en passant são inferidos do estado.
### Consequências
A API permanece pequena, sem hierarquia de tipos de movimento.

## ADR-004 — Validação por geometria e simulação

### Contexto
Auto-xeque e regras especiais tornam validação local insuficiente.
### Decisão
Validar a regra da peça e simular em uma cópia antes de aceitar.
### Consequências
Uma única estratégia cobre peças cravadas e exposição do rei.

## ADR-005 — Sem interfaces internas

### Contexto
Há uma implementação de domínio e uma de terminal.
### Decisão
Usar tipos concretos e `io.Reader`/`io.Writer` da biblioteca padrão.
### Consequências
Menos abstrações; testes ainda desacoplam I/O com buffers.

## ADR-006 — Somente biblioteca padrão

### Contexto
O projeto não necessita framework ou parser externo.
### Decisão
Não adicionar dependências externas sem necessidade e ADR novos.
### Consequências
Build simples, pequeno e reproduzível.

## ADR-007 — Identidade estrutural para repetição de posição

### Contexto
Uma disposição de peças isolada não distingue posições com movimentos legais diferentes por turno, roque ou en passant.
### Decisão
Representar a identidade por um valor comparável contendo os códigos das 64 casas, o jogador da vez, uma máscara dos quatro direitos de roque e a casa de en passant somente quando existe captura legal. Contar esses valores em um mapa mantido por `Game`.
### Consequências
A comparação é determinística e não depende de serialização ou hash sujeito a colisões. O campo `Moved` só influencia a máscara de roque; movimentos anteriores sem efeito nas jogadas disponíveis não diferenciam posições.

## ADR-008 — Histórico de repetição mantido por Game

### Contexto
A repetição depende de todas as posições aceitas da partida, mas não deve ser afetada por comandos ou tentativas inválidas.
### Decisão
`Game` mantém um mapa de identidade para contagem, registra posições nos construtores e após `Play` aceitar uma jogada, e recria o mapa em `Restart`.
### Consequências
O histórico fica no domínio e a CLI não conhece sua representação. Não há persistência nem lista de notações de movimentos.

## ADR-009 — Documentação de domínio por contexto

### Contexto
Um único arquivo de regras força tarefas pequenas a carregar assuntos não relacionados.
### Decisão
Usar `docs/domain/README.md` como índice e dividir estado, invariantes, movimentos, términos e posições em documentos focados. Manter `docs/DOMAIN.md` como ponte compatível.
### Consequências
Agentes podem carregar somente o ramo relevante; índices e links precisam permanecer consistentes.

## ADR-010 — Agentes especializados e skills compartilhadas

### Contexto
Exploração, implementação, testes, revisão e documentação têm necessidades diferentes de ferramentas, custo e contexto.
### Decisão
Definir cinco agentes equivalentes para Claude Code e Codex, com leitura restrita para exploração e revisão. Manter skills canônicas em `.claude/skills` no formato Agent Skills e expor o mesmo diretório ao Codex pelo link `.codex/skills`.
### Consequências
Procedimentos não são duplicados. Tarefas triviais continuam no agente principal; trabalhos que editam os mesmos arquivos são sequenciais, e o agente principal consolida toda delegação.
