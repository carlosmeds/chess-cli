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
