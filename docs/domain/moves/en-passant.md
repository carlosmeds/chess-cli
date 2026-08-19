# En passant

Um avanço duplo de peão cria `enPassantTarget` na casa atravessada. Na resposta imediata, um peão adversário adjacente pode mover diagonalmente para essa casa vazia e remover o peão que avançou.

Qualquer jogada seguinte limpa o alvo; por isso uma tentativa tardia é inválida. `validatePawn` reconhece o destino e `applyUnchecked` remove o peão capturado e atualiza o alvo.

Para repetição de posição, o alvo só diferencia posições quando o jogador da vez possui uma captura en passant legal, inclusive sem expor o próprio rei. Veja [identidade](../position/identity.md). `TestEnPassantAndImmediateWindow` cobre captura e expiração; testes de repetição cobrem a possibilidade efetiva.
