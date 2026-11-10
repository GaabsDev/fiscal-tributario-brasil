# Changelog — Fiscal Tributário Brasil

Todas as alterações notáveis neste projeto são documentadas deterministicamente neste arquivo.
O formato baseia-se no [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/) e adere ao [Semantic Versioning](https://semver.org/lang/pt-BR/).

## [v0.2.0] - 2026-11-10

### Especificações

- Documentos e serviços gerados: consulte a matriz do README e docs/SCHEMAS_E_SERVICOS.md.
- Referências normativas e schemas capturados: consulte schema-profiles-manifest.json e fiscal-sources-manifest.json.
- NFe leiaute 4.00.
- NT 2025.002 v1.53.
- NT 2026.005 v1.00.
- NT 2026.006 v1.01.
- NT 2026.007 v1.11.
- NT 2026.008 v1.01.
- NT 2026.009 v1.00.
- NT 2026.010 v1.00.
- Fiscal IR Hash: `2156e69df852c615df5e5af2c9c2848d2fc619e40c43880b71b7eccdc6413df4`
- Source Set Hash: `867ff0360723d4b418f15e6b98dbeec6b631ded020bf18c10a1ecf82ad8c0f8e`

### Regras fiscais

- Regra B25-40 (Rejeição 255)
- Regra C17-42 (Rejeição 156)
- Regra C17-43 (Rejeição 157)
- Regra C18-50 (Rejeição 158)
- Regra I11c-20 (Rejeição 1279)
- Regra N01-10 (Rejeição 161)
- Regra NB01-10 (Rejeição 1280)
- Regra NB01-20 (Rejeição 1280)
- Regra UB12-11 (Rejeição 162)
- Regra VC02-14 (Rejeição 321)
- Regra VC02-20 (Rejeição 1072)
- Regra VC02-30 (Rejeição 1130)
- Regra VC03-20 (Rejeição 1048)
- Regra W01a-10 (Rejeição 1282)
- Regra W01a-20 (Rejeição 1283)
- Regra W59h-10 (Rejeição 1275): exige o total gTribSN quando ao menos um item informa gTribSN
- Regra W59h-20 (Rejeição 1276): proíbe o total gTribSN quando nenhum item informa gTribSN
- Regra W59i-10 (Rejeição 1277): confere vIBSSN total com a soma exata dos itens
- Regra W59j-10 (Rejeição 1284): confere vCBSSN total com a soma exata dos itens
- Regra YC03-10 (Rejeição 215)
- Regra YC04-20 (Rejeição 215)
- Regra YC06-10 (Rejeição 1274)

### Correções

- Consulte as alterações de API e regras acima; correções não são presumidas sem evidências.


## Modelos e serviços

24 perfis e 470 schemas de entrada gerados. Referências normativas, caminhos de importação e hashes estão em schema-profiles-manifest.json.

- NF-e: modelo 55.
- NFC-e: modelo 65.
- CT-e: modelo 57.
- CT-e: modelo 57 (simples_nacional).
- CT-e OS: modelo 67.
- CT-e OS: modelo 67 (simples_nacional).
- GTV-e: modelo 64.
- GTV-e: modelo 64 (simples_nacional).
- MDF-e: modelo 58.
- BP-e: modelo 63.
- BP-e: modelo 63 (simples_nacional).
- NFCom: modelo 62.
- NFCom: modelo 62 (simples_nacional).
- NF3e: modelo 66.
- NFS-e Nacional: modelo NFS.
- DC-e: modelo 99.
- NFAg: modelo 75.
- NFGas: modelo 76.
- NF-e ABI: modelo 77.
- NFS-e Via: modelo NFSVia.
- NF-e Ouro: modelo Ouro.
- DAO: modelo DAO.
- DeRE: modelo DeRE (prévia técnica).
- DeRE: modelo DeRE (produção).

## Validação

972 testes de topo aprovados; 972 funções de teste geradas. Cobertura global: 55.44%. gofmt e go vet 100% em conformidade sem advertências.
