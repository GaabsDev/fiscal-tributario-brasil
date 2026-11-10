# Fiscal Tributário Brasil — validações

## Validação local

Validate retorna as verificações estruturais e de negócio implementadas no pacote. Em NF-e, os erros estruturados incluem caminho, código, severidade e, quando disponível, referência da regra.

- I05k-10 / 1165: classificação prod.tpCredPresIBSZFM informada na NFC-e modelo 65.
- LA01-30 / 1106: ausência do grupo de combustível quando a classificação tributária estiver abrangida pela regra.
- W59h-10 / 1275: total gTribSN ausente quando algum item possui gTribSN.
- W59h-20 / 1276: total gTribSN informado quando nenhum item possui gTribSN.
- W59i-10 / 1277: total vIBSSN diferente da soma exata dos itens.
- W59j-10 / 1284: total vCBSSN diferente da soma exata dos itens.
- B25-40, VC02-14, VC02-20, VC02-30 e VC03-20: consistência do referenciamento de NF-e/DF-e na devolução.
- YC03-10, YC04-20 e YC06-10: unicidade e identificação dos pagamentos vinculados.
- C17-42, C17-43, C18-50, N01-10 e UB12-11: emissão por contribuinte exclusivo de IBS/CBS.
- I11c-20, NB01-10, NB01-20, W01a-10 e W01a-20: valor líquido e ICMS previsto no pagamento antecipado.

A existência dessas verificações não demonstra a implementação de todas as regras do autorizador.

## Validação XSD

Execute a validação do XML produzido contra o pacote oficial correspondente. Os testes que exigem schemas e xmllint falham se esses recursos estiverem ausentes.

Verificar que um arquivo XSD tem sintaxe XML válida não equivale a validar um documento produzido pelo SDK.

## Evidências da versão

Consulte build-report.json e fiscal-sources-manifest.json. Uma norma descoberta só deve ser tratada como implementada quando houver a ligação entre fonte, implementação e validação.
