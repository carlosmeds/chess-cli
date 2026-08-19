# Estado da partida

`Game`, em `internal/chess/game.go`, concentra:

- `board`: peças e casas em um `Board` de valores;
- `turn`: jogador que deve mover;
- `status`: em andamento, xeque-mate, afogamento ou repetição tripla;
- `enPassantTarget`: casa temporária criada por avanço duplo de peão;
- `positions`: contagem das identidades de posição já ocorridas.

Os direitos de roque são derivados da posição, cor e campo `Moved` do rei e das torres, não armazenados em um contador separado. A possibilidade de en passant só participa da identidade quando há captura legal para o jogador da vez.

`NewGame` e `NewGameWithBoard` registram a posição inicial. `Play` rejeita partidas encerradas e jogadas inválidas antes de qualquer alteração; depois aplica a jogada, alterna o turno, verifica mate ou afogamento e registra a posição. `Restart` substitui todo esse estado por uma nova partida.

Veja [histórico de posições](position/history.md) e [término](termination/README.md).
