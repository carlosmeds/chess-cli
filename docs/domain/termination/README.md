# Término da partida

`Game.Status` representa as condições implementadas:

- [Xeque-mate](checkmate.md)
- [Afogamento](stalemate.md)
- [Repetição tripla](threefold-repetition.md)

Depois de cada jogada válida, o domínio verifica se o próximo jogador possui movimento legal. Sem movimentos, xeque produz mate e ausência de xeque produz afogamento. A posição resultante também é contada para repetição; mate e afogamento têm precedência. Qualquer status terminal faz `Play` rejeitar novas jogadas até `Restart`.

A CLI apenas traduz o status em mensagem; ela não decide o resultado.
