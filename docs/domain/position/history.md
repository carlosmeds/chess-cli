# Histórico de posições

`Game.positions` é um mapa de `positionKey` para número de ocorrências. Os construtores criam o mapa e registram a posição inicial. Após uma jogada válida, `Play` registra a posição resultante; ao atingir três ocorrências em uma partida ainda em andamento, define empate por repetição.

Validação e comandos da CLI não chamam `recordPosition`, portanto erros, `help` e `board` não alteram a contagem. `Restart` substitui `Game` por uma nova instância, descartando o mapa anterior.

O histórico existe apenas durante a sessão e não é uma lista de movimentos, FEN ou PGN. Persistência e exportação não estão implementadas.
