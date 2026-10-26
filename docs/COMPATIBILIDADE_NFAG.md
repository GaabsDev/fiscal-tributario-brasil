# Fiscal Tributário Brasil — compatibilidade NFAg

O pacote oficial da NT 2026.004 corrigido em 12/10/2026 contém uma escolha obrigatória vazia antes de det. O SDK conserva os arquivos originais e seus hashes. Uma cópia derivada remove exatamente essa escolha vazia; nenhuma outra regra ou campo é removido.

## Escolha explícita da validação

- ValidateXML, Parse e Marshal usam sempre o XSD oficial original.
- ValidateXMLCompatible, ParseCompatible e MarshalCompatible existem apenas nos serviços afetados e selecionam a correção local explicitamente.
- A validação compatível não comprova conformidade com o XSD oficial inalterado nem garante recepção pelo autorizador.
- Uma nova revisão oficial com outro hash nunca recebe essa transformação automaticamente.

Serviços afetados: bindings/nfag/nfag_v1_00 e bindings/nfag/procnfag_v1_00. As funções recebem os mesmos argumentos das funções estritas correspondentes.

## Rastreabilidade

O arquivo schema-compatibility-manifest.json registra operação, escopo, hashes e exigência de escolha explícita. Os bytes oficiais ficam em schemafiles/nfag/data; a cópia derivada fica em schemafiles/nfag/compatibility. O pacote de origem continua identificado pela sua revisão oficial; a cópia derivada é identificada por uma regra de compatibilidade do Fiscal Tributário Brasil.

Os testes confirmam simultaneamente a rejeição pelo original, a conversão XML/Go/XML pelo derivado, a rejeição de det inválido e a imutabilidade do original.

- Regra: `nfag-nt2026.004-empty-choice/v1`.
- Arquivo: `PL_NFAg_NT2026.004 RTC_1.00/nfagTiposBasico_v1.00.xsd`.
- SHA-256 oficial: `e72acd32937077c33831aaaf87dae91f96fef65206accadb4eee9f948826e7a4`.
- SHA-256 derivado: `dd30fd9d456c28d92b2ca73c5d12d1588a13e2ffcfa3374da7424257d4fb8b9b`.
