# Invariantes do domínio

- Uma posição válida possui exatamente um rei de cada cor. Os construtores de teste permitem montar estados intermediários, mas a ausência de rei é tratada como xeque.
- Exatamente um lado joga por vez; uma jogada válida alterna o turno.
- O rei não é capturado e o jogador não conclui uma jogada deixando o próprio rei em xeque.
- Somente jogadas validadas alteram tabuleiro, turno, alvo en passant ou histórico de posições.
- O alvo en passant expira após a resposta imediata.
- Um estado terminal impede novas jogadas; `Restart` cria uma nova partida.
- `internal/chess` não depende da CLI nem faz I/O.
- Não existe estado global mutável.

Essas propriedades são protegidas principalmente por `Game.Play`, `Game.validate`, simulação em cópia e testes em `internal/chess`.
