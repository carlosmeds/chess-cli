# Glossário

- **Board**: matriz de 64 casas contendo valores `Piece`.
- **Color**: lado branco ou preto; também determina direção dos peões e turno.
- **Game**: agregado que mantém tabuleiro, turno, estado de término, alvo en passant e histórico de posições.
- **Move**: valor com origem, destino e promoção opcional.
- **Piece**: tipo, cor e indicador `Moved`; casa com `NoPiece` é vazia.
- **Position**: casa por arquivo e fileira, armazenados como índices de zero a sete.
- **Xeque**: rei do jogador está sob ataque.
- **Jogada legal**: movimento válido para a peça que não deixa o próprio rei em xeque.
- **Direito de roque**: possibilidade ainda preservada pelo rei e por uma torre em suas casas iniciais, independentemente de o caminho estar atualmente livre.
- **En passant efetivo**: alvo temporário para o qual o jogador da vez possui uma captura en passant legal.
