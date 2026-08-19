# Experimento local de RAG

Este diretório mede se selecionar trechos reduz contexto sem perder as fontes necessárias. Ele não integra o executável de xadrez, não indexa código-fonte e não é uma proposta de produção.

## Estratégias

- `full`: todos os chunks do corpus e dos distratores, sem orçamento.
- `lexical`: BM25 local, até 20 candidatos e no máximo 6 chunks dentro do orçamento.
- `hybrid`: BM25 + vetores locais, combinados por Reciprocal Rank Fusion (RRF), com os mesmos limites.

O corpus oficial é fixado no código para evitar inclusão acidental: `docs/domain/`, `docs/ARCHITECTURE.md`, `docs/DECISIONS.md` e `docs/TESTING.md`. `docs/conventions/` seria aceito pelo escopo, mas não existe hoje. Os arquivos em `distractors/` são rotulados no índice e não afirmam comportamento oficial.

## Execução

Na raiz do repositório:

```sh
make rag-index
make rag-search QUERY="como funciona a repetição tripla?"
make rag-eval
make rag-clean
```

O índice e o relatório são gravados em `rag-experiment/data/` e ignorados pelo Git. `rag-clean` usa um alvo explícito e preserva `.gitkeep`.

## Embeddings e segurança

`embedding_provider: local-hash` é o padrão: um vetor determinístico de feature hashing com um vocabulário pequeno de paráfrases em português. Ele serve como baseline vetorial reproduzível, não é um modelo aprendido. Não usa rede, credenciais ou dependências externas. `embedding_provider: none` mantém indexação e busca lexical sem vetores; nesse modo, use `-strategy lexical`.

A interface `embedding.Provider` isola o provedor para um ensaio posterior com embeddings aprendidos. Esta primeira versão deliberadamente não implementa chamadas HTTP: não há infraestrutura autorizada no projeto, e documentos não devem ser enviados externamente sem decisão explícita. Se um provedor remoto for acrescentado, a chave deve vir apenas de variável de ambiente, nunca deste YAML, e o envio precisa ser habilitado conscientemente.

Chunks cujo `content_hash` não mudou reutilizam o vetor do índice anterior quando o provedor também é o mesmo. Alterar dimensão ou algoritmo exige mudar a identificação do provedor ou executar `make rag-clean` antes de reindexar.

## Chunking e metadados

Markdown é dividido por título e seções `##`. Subseções de um ADR permanecem no mesmo chunk, preservando contexto, decisão e consequências. Não há cortes cegos nem sobreposição nesta versão, pois as seções atuais são pequenas e autocontidas. Cada registro contém `id`, `source`, `heading`, `content_type`, `domain`, `symbols`, `status`, `content_hash`, conteúdo, marcador de distrator e vetor.

## Avaliação e limites

As 18 perguntas e seus conceitos ficam em `evals/questions.json`; fontes esperadas ficam separadas em `evals/expected-sources.json`. O relatório calcula Recall@K de fontes, Precision@K por chunk, MRR, cobertura literal de conceitos, fundamentação, recusa, estimativa de tokens, redução, latência e chunks distratores.

Não há cliente de LLM no projeto. Para manter o ensaio executável e não introduzir credenciais, a “resposta” é uma decisão extrativa determinística baseada apenas no contexto recuperado. Portanto, cobertura de conceitos e fundamentação são proxies auditáveis, não uma nota de fluência ou correção dada por um modelo. A tabela completa permite revisão humana. Uma segunda rodada com modelo deve usar exatamente o mesmo prompt, temperatura e limites nas três estratégias e registrar tokens reais da API separadamente.

Tokens são estimados como `ceil(caracteres UTF-8 / 4)` e claramente marcados como estimativa. Latências variam por máquina e devem ser lidas como comparação local, não benchmark absoluto. Os limiares configuráveis estão em `config.example.yaml`.

## Custo de manutenção

O custo imediato é baixo: corpus explícito, JSON pequeno e índice local. O custo cresce se o vocabulário manual de paráfrases precisar acompanhar novos temas; isso é a principal fragilidade do baseline vetorial. Um embedding aprendido reduziria essa manutenção, mas acrescentaria custo, privacidade, versionamento do modelo e reindexação.
