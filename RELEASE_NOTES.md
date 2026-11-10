## Fiscal Tributário Brasil v0.2.0

Versão do SDK independente Fiscal Tributário Brasil em Go para documentos fiscais eletrônicos e Reforma Tributária.

A cobertura completa de documentos, eventos e serviços está na matriz do README e em docs/SCHEMAS_E_SERVICOS.md.

### Instalação

```bash
go get github.com/GaabsDev/fiscal-tributario-brasil@v0.2.0
```

### Resumo de Compatibilidade e Alterações

- **Tipo de Mudança (SemVer):** `minor`
- O snapshot local da API anterior não estava disponível; nenhuma conclusão de compatibilidade símbolo a símbolo foi presumida.
- Atualização minor sobre v0.1.0; o snapshot local da API anterior não estava disponível para comparação símbolo a símbolo.

### Cobertura normativa desta versão

- NFe leiaute 4.00.
- NT 2025.002 v1.53.
- NT 2026.005 v1.00.
- NT 2026.006 v1.01.
- NT 2026.007 v1.11.
- NT 2026.008 v1.01.
- NT 2026.009 v1.00.
- NT 2026.010 v1.00.

### Regras Fiscais Automatizadas / Verificadas

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

### Limitações Desta Versão

- Esta biblioteca cobre modelos de dados, serialização XML/JSON e validações de regras normativas.
- Assinatura digital e comunicação com o autorizador exigem integração própria.

### Verificação do pacote

Hash do SDK: `38cee700bc1d6bd4535d7ea7400a704618f90f9f177509babf5d14240f16addf`. Os arquivos de distribuição e os manifestos estão anexados à release.

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
