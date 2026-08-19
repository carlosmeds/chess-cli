# Movimentos comuns

Peões avançam na direção da cor, uma casa ou duas desde a origem se o caminho estiver livre, e capturam uma peça adversária na diagonal. Cavalos movem em L. Bispos percorrem diagonais, torres linhas e colunas, damas combinam ambos, e reis movem uma casa.

Pré-condições comuns:

- origem e destino são casas válidas e diferentes;
- a origem contém uma peça do jogador da vez;
- o destino não contém peça aliada nem rei adversário;
- peças deslizantes possuem caminho livre;
- a posição simulada não deixa o rei do jogador em xeque.

Uma jogada inválida retorna erro sem alterar estado ou histórico. `internal/chess/game_test.go` cobre movimentos iniciais, caminhos bloqueados, peça aliada, xeque e auto-xeque. Promoção, roque e en passant têm documentos próprios.
