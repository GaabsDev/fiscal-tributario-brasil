# Fiscal Tributário Brasil — validações

## Validação local

Validate retorna as verificações estruturais e de negócio implementadas no pacote. Em NF-e, os erros estruturados incluem caminho, código, severidade e, quando disponível, referência da regra.

- I05k-10 / 1165: classificação prod.tpCredPresIBSZFM informada na NFC-e modelo 65.
- LA01-30 / 1106: ausência do grupo de combustível quando a classificação tributária estiver abrangida pela regra.

A existência dessas verificações não demonstra a implementação de todas as regras do autorizador.

## Validação XSD

Execute a validação do XML produzido contra o pacote oficial correspondente. Os testes que exigem schemas e xmllint falham se esses recursos estiverem ausentes.

Verificar que um arquivo XSD tem sintaxe XML válida não equivale a validar um documento produzido pelo SDK.

## Evidências da versão

Consulte build-report.json e fiscal-sources-manifest.json. Uma norma descoberta só deve ser tratada como implementada quando houver a ligação entre fonte, implementação e validação.
