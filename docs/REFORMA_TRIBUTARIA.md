# Fiscal Tributário Brasil — IBS/CBS e classificação tributária

As estruturas tributárias devem ser preenchidas conforme a operação e a revisão técnica aplicável. Os exemplos não definem alíquotas ou enquadramento fiscal para uso em produção.

O catálogo cClassTrib é capturado da fonte oficial e inclui indicadores por modelo e datas de vigência. Quantidade de registros não equivale à quantidade de códigos vigentes na data da emissão.

A regra I05k-10, rejeição 1165, trata da classificação para subapuração do IBS na ZFM em prod.tpCredPresIBSZFM na NFC-e modelo 65. Ela não proíbe genericamente IBS/CBS na NFC-e.

A NT 2026.005 v1.00 acrescenta tpRegimeIBSCBS, cAtivSN e gTribSN. O SDK implementa o modelo e as totalizações a partir da nota oficial e preserva o XSD 010g v1.03 sem alterações. vRBSNitem e tpRBSN pertencem ao evento futuro 412140 gerado pela Administração Tributária.

A geração também cobre as revisões NT 2025.002 v1.53, NT 2026.006 v1.01, NT 2026.007 v1.11 e NT 2026.008 v1.01, incluindo referenciamento de DF-e por item, vinculação de pagamentos, IE opcional para o emitente exclusivo de IBS/CBS no modelo 55 e os campos de valor líquido. As datas de implantação são registradas separadamente da disponibilidade estrutural no SDK.

Consulte as fontes e os hashes no manifesto da geração. Publicação de uma NT, disponibilidade em homologação e vigência em produção devem ser verificadas separadamente.
