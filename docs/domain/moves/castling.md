# Roque

O roque é informado como o movimento normal do rei: `e1 g1`, `e1 c1`, `e8 g8` ou `e8 c8`.

É permitido quando rei e torre correspondentes estão nas casas iniciais e nunca moveram, o caminho está vazio, o rei não está em xeque e nenhuma casa atravessada está atacada. Ao aplicar, `applyUnchecked` move também a torre e marca ambas as peças como movidas.

É inválido se qualquer pré-condição falhar. `validateCastle` em `internal/chess/rules.go` concentra a regra; `TestCastling` e `TestCastlingThroughCheckIsRejected` cobrem efeito e passagem por ataque. Os direitos remanescentes também participam da [identidade da posição](../position/identity.md).
