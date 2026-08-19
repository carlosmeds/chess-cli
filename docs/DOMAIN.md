# Domínio e regras

`Position` usa arquivo `a`–`h` e fileira `1`–`8`, armazenados internamente como índices zero-based. `Board` contém 64 casas. `Piece` tem tipo, cor e indicador de movimento. `Move` contém origem, destino e promoção opcional. `Game` mantém tabuleiro, turno, estado e alvo temporário de en passant.

Peões avançam na direção da cor, podem avançar duas casas desde a origem e capturam na diagonal. Cavalos movem em L. Bispos usam diagonais, torres linhas, damas ambos e reis uma casa. Peças deslizantes exigem caminho livre; destino aliado é proibido. Toda jogada é simulada e rejeitada se o rei do jogador permanecer atacado.

Xeque ocorre quando a casa do rei é atacada. Sem jogadas legais, há xeque-mate se o rei está em xeque e afogamento caso contrário.

- Roque: rei e torre não moveram, caminho livre, rei fora de xeque e nenhuma casa atravessada atacada. Entrada: `e1 g1`, `e1 c1`, `e8 g8` ou `e8 c8`.
- Promoção: ao chegar à última fileira, o peão deve virar dama, torre, bispo ou cavalo: `e7 e8 q`.
- En passant: disponível somente na resposta imediata ao avanço duplo de um peão adversário adjacente.

Invariantes: exatamente um lado joga por vez; rei não é capturado; jogada válida não deixa o próprio rei em xeque; alvo de en passant expira após uma jogada.

Exemplos válidos na posição inicial: `e2 e4`, `g1 f3`. Inválidos: `e2 e5` (distância), `c1 h6` (caminho bloqueado), mover uma peça preta no turno branco ou ocupar casa de peça aliada.
