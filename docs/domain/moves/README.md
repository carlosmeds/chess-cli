# Movimentos

- [Movimentos comuns](standard-moves.md): geometria, ocupação, caminho e auto-xeque.
- [Roque](castling.md): movimento conjunto de rei e torre.
- [Promoção](promotion.md): substituição do peão na última fileira.
- [En passant](en-passant.md): captura temporária após avanço duplo.

`Game.validate` faz as verificações gerais; `validatePieceMove` e auxiliares em `internal/chess/rules.go` tratam as regras por peça. A aplicação ocorre somente depois da simulação contra auto-xeque.
