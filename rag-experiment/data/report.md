# Relatório do experimento RAG

Contagens de tokens são estimativas (`caracteres UTF-8 / 4`). A qualidade abaixo usa recuperação de fontes, cobertura de conceitos e um respondedor extrativo determinístico; não é julgamento por LLM.

Índice: 199448 bytes; indexação: 2.934002ms.

| Pergunta | Estratégia | Fontes recuperadas | Recall@K | Tokens | Resultado |
|---|---|---|---:|---:|---|
| q01 — Como o cavalo se move e uma jogada pode deixar o próprio rei atacado? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q01 — Como o cavalo se move e uma jogada pode deixar o próprio rei atacado? | lexical | docs/ARCHITECTURE.md<br>docs/domain/glossary.md<br>docs/domain/moves/promotion.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 0.00 | 772 | falha |
| q01 — Como o cavalo se move e uma jogada pode deixar o próprio rei atacado? | hybrid | docs/ARCHITECTURE.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/promotion.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md | 0.00 | 920 | falha |
| q02 — Quais são as condições para fazer roque? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q02 — Quais são as condições para fazer roque? | lexical | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/identity.md | 1.00 | 981 | ok |
| q02 — Quais são as condições para fazer roque? | hybrid | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/standard-moves.md | 1.00 | 869 | ok |
| q03 — Por quanto tempo vale a captura de passagem do peão? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | falha |
| q03 — Por quanto tempo vale a captura de passagem do peão? | lexical | docs/domain/game-state.md<br>docs/domain/moves/README.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/position/identity.md<br>rag-experiment/distractors/board-games.md | 1.00 | 945 | falha |
| q03 — Por quanto tempo vale a captura de passagem do peão? | hybrid | docs/domain/moves/README.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/position/identity.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 754 | falha |
| q04 — Em quais peças um peão pode se transformar ao chegar ao fim? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q04 — Em quais peças um peão pode se transformar ao chegar ao fim? | lexical | docs/ARCHITECTURE.md<br>docs/domain/moves/README.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 634 | ok |
| q04 — Em quais peças um peão pode se transformar ao chegar ao fim? | hybrid | docs/DECISIONS.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/identity.md<br>rag-experiment/distractors/board-games.md | 1.00 | 846 | ok |
| q05 — Qual a diferença entre xeque-mate e afogamento? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q05 — Qual a diferença entre xeque-mate e afogamento? | lexical | docs/DECISIONS.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md | 1.00 | 834 | ok |
| q05 — Qual a diferença entre xeque-mate e afogamento? | hybrid | docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/game-state.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md | 1.00 | 1045 | ok |
| q06 — Quando a mesma configuração do tabuleiro encerra a partida empatada? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q06 — Quando a mesma configuração do tabuleiro encerra a partida empatada? | lexical | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/domain/README.md<br>docs/domain/invariants.md<br>docs/domain/termination/threefold-repetition.md | 0.50 | 1062 | falha |
| q06 — Quando a mesma configuração do tabuleiro encerra a partida empatada? | hybrid | docs/DECISIONS.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/position/identity.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md | 1.00 | 1021 | ok |
| q07 — Quais campos formam positionKey? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q07 — Quais campos formam positionKey? | lexical | docs/DECISIONS.md<br>docs/domain/game-state.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md | 1.00 | 813 | ok |
| q07 — Quais campos formam positionKey? | hybrid | docs/DECISIONS.md<br>docs/domain/game-state.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>rag-experiment/distractors/board-games.md | 1.00 | 973 | ok |
| q08 — Mover uma torre e devolvê-la ao lugar preserva a identidade da posição? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q08 — Mover uma torre e devolvê-la ao lugar preserva a identidade da posição? | lexical | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/domain/game-state.md<br>docs/domain/moves/castling.md<br>docs/domain/position/identity.md<br>docs/domain/termination/threefold-repetition.md | 1.00 | 1151 | ok |
| q08 — Mover uma torre e devolvê-la ao lugar preserva a identidade da posição? | hybrid | docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/game-state.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 0.00 | 1121 | falha |
| q09 — Qual camada pode depender de internal/chess e onde fica o I/O? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q09 — Qual camada pode depender de internal/chess e onde fica o I/O? | lexical | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/domain/README.md<br>docs/domain/invariants.md | 1.00 | 563 | ok |
| q09 — Qual camada pode depender de internal/chess e onde fica o I/O? | hybrid | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/invariants.md<br>rag-experiment/distractors/board-games.md | 1.00 | 897 | ok |
| q10 — Por que o núcleo do jogo não escreve diretamente no terminal? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q10 — Por que o núcleo do jogo não escreve diretamente no terminal? | lexical | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/domain/invariants.md<br>docs/domain/position/identity.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 695 | ok |
| q10 — Por que o núcleo do jogo não escreve diretamente no terminal? | hybrid | docs/DECISIONS.md<br>docs/domain/invariants.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md | 0.50 | 961 | falha |
| q11 — O que acontece com estado e histórico após uma jogada inválida? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q11 — O que acontece com estado e histórico após uma jogada inválida? | lexical | docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md | 1.00 | 1289 | ok |
| q11 — O que acontece com estado e histórico após uma jogada inválida? | hybrid | docs/ARCHITECTURE.md<br>docs/TESTING.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/position/history.md<br>docs/domain/termination/threefold-repetition.md | 0.50 | 1328 | falha |
| q12 — Como testar regras do domínio e quando usar o runner? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q12 — Como testar regras do domínio e quando usar o runner? | lexical | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 954 | ok |
| q12 — Como testar regras do domínio e quando usar o runner? | hybrid | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md | 1.00 | 1072 | ok |
| q13 — Quais helpers são recomendados para montar posições de teste sem CLI? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q13 — Quais helpers são recomendados para montar posições de teste sem CLI? | lexical | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/invariants.md<br>docs/domain/termination/threefold-repetition.md | 1.00 | 1070 | ok |
| q13 — Quais helpers são recomendados para montar posições de teste sem CLI? | hybrid | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/invariants.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 989 | ok |
| q14 — Onde validateCastle é implementado e quais testes cobrem a passagem por ataque? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q14 — Onde validateCastle é implementado e quais testes cobrem a passagem por ataque? | lexical | docs/ARCHITECTURE.md<br>docs/TESTING.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md | 1.00 | 1223 | ok |
| q14 — Onde validateCastle é implementado e quais testes cobrem a passagem por ataque? | hybrid | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/glossary.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md | 1.00 | 1087 | ok |
| q15 — Por que o tabuleiro usa [8][8]Piece em vez de ponteiros? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 1.00 | 4874 | ok |
| q15 — Por que o tabuleiro usa [8][8]Piece em vez de ponteiros? | lexical | docs/DECISIONS.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/position/identity.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md | 1.00 | 909 | ok |
| q15 — Por que o tabuleiro usa [8][8]Piece em vez de ponteiros? | hybrid | docs/DECISIONS.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/position/identity.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md | 1.00 | 909 | ok |
| q16 — Qual algoritmo controla o incremento do relógio de cada jogador? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 0.00 | 4874 | recusa correta |
| q16 — Qual algoritmo controla o incremento do relógio de cada jogador? | lexical | docs/TESTING.md<br>docs/domain/invariants.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md | 0.00 | 1235 | recusa correta |
| q16 — Qual algoritmo controla o incremento do relógio de cada jogador? | hybrid | docs/domain/invariants.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/threefold-repetition.md | 0.00 | 1167 | recusa correta |
| q17 — Como configurar partidas pela rede com autenticação de usuários? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 0.00 | 4874 | recusa correta |
| q17 — Como configurar partidas pela rede com autenticação de usuários? | lexical | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/domain/game-state.md<br>docs/domain/position/identity.md<br>rag-experiment/distractors/operational-notes.md | 0.00 | 805 | recusa correta |
| q17 — Como configurar partidas pela rede com autenticação de usuários? | hybrid | docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/position/identity.md | 0.00 | 1312 | recusa correta |
| q18 — Qual engine escolhe as jogadas da inteligência artificial? | full | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/TESTING.md<br>docs/domain/README.md<br>docs/domain/game-state.md<br>docs/domain/glossary.md<br>docs/domain/invariants.md<br>docs/domain/moves/README.md<br>docs/domain/moves/castling.md<br>docs/domain/moves/en-passant.md<br>docs/domain/moves/promotion.md<br>docs/domain/moves/standard-moves.md<br>docs/domain/position/history.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/checkmate.md<br>docs/domain/termination/stalemate.md<br>docs/domain/termination/threefold-repetition.md<br>rag-experiment/distractors/board-games.md<br>rag-experiment/distractors/operational-notes.md | 0.00 | 4874 | recusa correta |
| q18 — Qual engine escolhe as jogadas da inteligência artificial? | lexical | docs/ARCHITECTURE.md<br>docs/domain/game-state.md<br>docs/domain/invariants.md<br>docs/domain/position/identity.md<br>docs/domain/termination/README.md<br>docs/domain/termination/threefold-repetition.md | 0.00 | 1323 | recusa correta |
| q18 — Qual engine escolhe as jogadas da inteligência artificial? | hybrid | docs/ARCHITECTURE.md<br>docs/DECISIONS.md<br>docs/domain/game-state.md<br>docs/domain/invariants.md<br>docs/domain/termination/README.md<br>docs/domain/termination/threefold-repetition.md | 0.00 | 1262 | recusa correta |

## Médias por estratégia

| Estratégia | Recall | Precision | MRR | Conceitos | Fundamentadas | Recusa correta | Busca adicional | Tokens | Redução | Latência |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| full | 1.000 | 0.075 | 0.183 | 1.000 | 0.944 | 1.000 | 0.000 | 4874 | 0.0% | 28.367µs |
| lexical | 0.900 | 0.217 | 0.682 | 0.972 | 0.889 | 1.000 | 0.133 | 958 | 80.3% | 1.726388ms |
| hybrid | 0.800 | 0.200 | 0.572 | 0.917 | 0.833 | 1.000 | 0.267 | 1029 | 78.9% | 1.853813ms |

## Perguntas em que o RAG falhou

- q01: Como o cavalo se move e uma jogada pode deixar o próprio rei atacado?
- q08: Mover uma torre e devolvê-la ao lugar preserva a identidade da posição?
- q10: Por que o núcleo do jogo não escreve diretamente no terminal?
- q11: O que acontece com estado e histórico após uma jogada inválida?

## Perguntas em que o lexical venceu o vetorial

- q02: Quais são as condições para fazer roque?
- q04: Em quais peças um peão pode se transformar ao chegar ao fim?
- q08: Mover uma torre e devolvê-la ao lugar preserva a identidade da posição?
- q10: Por que o núcleo do jogo não escreve diretamente no terminal?
- q11: O que acontece com estado e histórico após uma jogada inválida?

## Chunks irrelevantes recuperados

- q01/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q01/full: rag-experiment/distractors/board-games.md — Promoção e término
- q01/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q01/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q01/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q01/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q01/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q01/hybrid: rag-experiment/distractors/board-games.md — Posição e repetição
- q01/lexical: rag-experiment/distractors/board-games.md — Posição e repetição
- q01/lexical: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q02/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q02/full: rag-experiment/distractors/board-games.md — Promoção e término
- q02/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q02/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q02/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q02/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q02/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q03/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q03/full: rag-experiment/distractors/board-games.md — Promoção e término
- q03/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q03/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q03/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q03/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q03/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q03/hybrid: rag-experiment/distractors/board-games.md — Promoção e término
- q03/hybrid: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q03/lexical: rag-experiment/distractors/board-games.md — Promoção e término
- q04/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q04/full: rag-experiment/distractors/board-games.md — Promoção e término
- q04/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q04/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q04/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q04/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q04/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q04/hybrid: rag-experiment/distractors/board-games.md — Promoção e término
- q04/lexical: rag-experiment/distractors/board-games.md — Posição e repetição
- q04/lexical: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q05/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q05/full: rag-experiment/distractors/board-games.md — Promoção e término
- q05/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q05/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q05/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q05/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q05/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q06/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q06/full: rag-experiment/distractors/board-games.md — Promoção e término
- q06/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q06/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q06/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q06/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q06/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q06/hybrid: rag-experiment/distractors/board-games.md — Promoção e término
- q07/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q07/full: rag-experiment/distractors/board-games.md — Promoção e término
- q07/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q07/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q07/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q07/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q07/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q07/hybrid: rag-experiment/distractors/board-games.md — Promoção e término
- q08/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q08/full: rag-experiment/distractors/board-games.md — Promoção e término
- q08/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q08/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q08/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q08/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q08/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q08/hybrid: rag-experiment/distractors/board-games.md — Posição e repetição
- q08/hybrid: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q09/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q09/full: rag-experiment/distractors/board-games.md — Promoção e término
- q09/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q09/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q09/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q09/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q09/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q09/hybrid: rag-experiment/distractors/board-games.md — Posição e repetição
- q10/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q10/full: rag-experiment/distractors/board-games.md — Promoção e término
- q10/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q10/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q10/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q10/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q10/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q10/lexical: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q11/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q11/full: rag-experiment/distractors/board-games.md — Promoção e término
- q11/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q11/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q11/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q11/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q11/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q12/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q12/full: rag-experiment/distractors/board-games.md — Promoção e término
- q12/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q12/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q12/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q12/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q12/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q12/lexical: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q12/lexical: rag-experiment/distractors/operational-notes.md — Testes de interface
- q13/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q13/full: rag-experiment/distractors/board-games.md — Promoção e término
- q13/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q13/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q13/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q13/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q13/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q13/hybrid: rag-experiment/distractors/operational-notes.md — Testes de interface
- q14/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q14/full: rag-experiment/distractors/board-games.md — Promoção e término
- q14/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q14/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q14/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q14/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q14/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q15/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q15/full: rag-experiment/distractors/board-games.md — Promoção e término
- q15/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q15/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q15/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q15/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q15/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q16/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q16/full: rag-experiment/distractors/board-games.md — Promoção e término
- q16/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q16/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q16/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q16/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q16/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q16/lexical: rag-experiment/distractors/board-games.md — Promoção e término
- q17/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q17/full: rag-experiment/distractors/board-games.md — Promoção e término
- q17/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q17/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q17/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q17/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q17/full: rag-experiment/distractors/operational-notes.md — Testes de interface
- q17/lexical: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q18/full: rag-experiment/distractors/board-games.md — Posição e repetição
- q18/full: rag-experiment/distractors/board-games.md — Promoção e término
- q18/full: rag-experiment/distractors/board-games.md — Vocabulário semelhante em outros jogos
- q18/full: rag-experiment/distractors/operational-notes.md — Estado de conexão
- q18/full: rag-experiment/distractors/operational-notes.md — Histórico de comandos
- q18/full: rag-experiment/distractors/operational-notes.md — Notas operacionais fictícias para recuperação
- q18/full: rag-experiment/distractors/operational-notes.md — Testes de interface

## Conclusão

Os limiares configurados não foram atingidos. Recomendação experimental: ajustar ou abandonar; não adotar este RAG com base neste ensaio.
