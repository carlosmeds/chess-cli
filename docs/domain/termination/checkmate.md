# Xeque-mate

Há xeque-mate quando o jogador da vez está em xeque e não possui jogada legal. `hasLegalMove` enumera candidatos e usa a mesma validação aplicada por `Play`, incluindo promoção e auto-xeque.

O domínio define `Checkmate`; a CLI anuncia como vencedor o oponente do jogador sem movimentos. `TestCheckmateFoolsMate` cobre o mate do louco. Alterações na enumeração de movimentos devem preservar roque, promoção e en passant.
