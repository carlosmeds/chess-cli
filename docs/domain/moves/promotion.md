# Promoção

Ao alcançar a última fileira, um peão deve ser promovido a dama, torre, bispo ou cavalo. A entrada inclui `q`, `r`, `b` ou `n`, por exemplo `e7 e8 q`.

Promoção é inválida para outra peça, fora da última fileira, sem escolha válida ou para rei/peão. `validatePawn` verifica a escolha e `applyUnchecked` substitui o tipo depois de mover. O parser converte a letra em `Move.Promotion`.

`TestPromotion` cobre a alteração no domínio e `TestParseMoveAndPromotion` cobre a entrada da CLI.
