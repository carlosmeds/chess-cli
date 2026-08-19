# Afogamento

Há afogamento quando o jogador da vez não está em xeque, mas não possui jogada legal. O resultado é empate e usa a mesma enumeração de `hasLegalMove` do xeque-mate.

O domínio define `Stalemate`; a CLI apresenta `Afogamento: empate.`. `TestStalemate` monta uma posição mínima com `NewEmptyBoard` e confirma também a ausência de xeque.
