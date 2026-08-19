# Repetição tripla

A partida termina empatada na terceira ocorrência da mesma posição. A posição inicial e o resultado de cada jogada válida são registrados; tentativas inválidas e comandos de apresentação não entram no histórico. `Restart` limpa a contagem e registra uma nova posição inicial.

Duas posições são iguais somente com as mesmas peças nas mesmas casas, o mesmo jogador da vez, os mesmos direitos de roque e a mesma possibilidade efetiva de captura en passant. Número da jogada, contador de 50 movimentos, peças capturadas e histórico completo não participam.

`internal/chess/repetition.go` produz a chave e `Game.Play` define `ThreefoldRepetition` na terceira ocorrência, sem sobrescrever mate ou afogamento. `repetition_test.go` cobre identidade, contagem, reinício, jogadas inválidas e o ciclo real dos cavalos; `runner_test.go` cobre a mensagem da CLI.
