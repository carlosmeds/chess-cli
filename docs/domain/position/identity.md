# Identidade da posição

`positionKey` é um valor Go comparável e determinístico formado por:

- `[64]uint8` com tipo e cor de cada peça, na ordem das casas;
- jogador da vez;
- máscara de quatro direitos de roque;
- índice da casa en passant ou sentinela quando não há captura efetiva.

Os direitos são derivados de rei e torres nas casas iniciais com `Moved == false`. Caminho bloqueado ou casa atacada não elimina o direito: apenas impede o roque naquela jogada. O campo `Moved` de outras peças não participa.

O alvo en passant participa somente se `effectiveEnPassantTarget` encontra um peão do jogador da vez cuja captura passa pela validação completa, incluindo auto-xeque. Assim, um avanço duplo sem capturador adjacente não cria uma identidade diferente.

A chave não inclui número da jogada, regra dos 50 movimentos, peças fora do tabuleiro nem histórico. Ela é usada diretamente como chave de mapa, sem serialização ou hash próprio.
