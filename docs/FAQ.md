# Fiscal Tributário Brasil — perguntas frequentes

## Qual é o nome e qual é o import?

O projeto se chama Fiscal Tributário Brasil. O caminho técnico do módulo é github.com/GaabsDev/fiscal-tributario-brasil.

## O SDK emite e autoriza notas sozinho?

A biblioteca oferece modelos, serialização e validações implementadas. Seu sistema integra as demais etapas, incluindo assinatura e comunicação com o autorizador.

## Validate garante que a SEFAZ aceitará o documento?

Não. A aceitação depende também do XSD aplicável, das regras do ambiente, do cadastro, da assinatura e da resposta do serviço fiscal.

## Como conferir as versões técnicas utilizadas?

Consulte o manifesto de fontes, seus hashes e o relatório de build. Versões de leiaute, NT, manual e XSD devem ser conferidas separadamente.

## Posso reutilizar uma versão Go já publicada?

Versões distribuídas pelo ecossistema Go são imutáveis. Uma correção deve receber uma nova versão. Excluir a tag no GitHub não altera o conteúdo já armazenado pelos proxies.

## Onde encontro exemplos?

Na pasta examples e no guia GUIA_DE_USO.md. Os dados dos exemplos são fictícios.
