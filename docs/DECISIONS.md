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
