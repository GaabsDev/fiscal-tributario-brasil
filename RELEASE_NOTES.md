## Fiscal Tributário Brasil v0.1.0

Versão do SDK independente Fiscal Tributário Brasil em Go para documentos fiscais eletrônicos e Reforma Tributária.

A cobertura completa de documentos, eventos e serviços está na matriz do README e em docs/SCHEMAS_E_SERVICOS.md.

### Instalação

```bash
go get github.com/GaabsDev/fiscal-tributario-brasil@v0.1.0
```

### Resumo de Compatibilidade e Alterações

- **Tipo de Mudança (SemVer):** `initial`
- Primeira versão deste módulo.
- Primeira publicação do módulo, versão v0.1.0.

### Regras Fiscais Automatizadas / Verificadas

- Regra I05k-10 (Rejeição oficial 1165): validação em NFC-e modelo 65

### Limitações Desta Versão

- Esta biblioteca cobre modelos de dados, serialização XML/JSON e validações de regras normativas.
- Assinatura digital e comunicação com o autorizador exigem integração própria.

### Verificação do pacote

Hash do SDK: `056d8496a71741ad55307ae7a83ddeadddfeba0e409da3cc6bfdb98e7bfde521`. Os arquivos de distribuição e os manifestos estão anexados à release.

## Modelos e serviços

19 perfis e 308 schemas de entrada gerados. Referências normativas, caminhos de importação e hashes estão em schema-profiles-manifest.json.

- NF-e: modelo 55.
- NFC-e: modelo 65.
- CT-e: modelo 57.
- CT-e OS: modelo 67.
- GTV-e: modelo 64.
- MDF-e: modelo 58.
- BP-e: modelo 63.
- NFCom: modelo 62.
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

647 testes de topo aprovados; 647 funções de teste geradas. Cobertura global: 54.62%. gofmt e go vet 100% em conformidade sem advertências.

## Compatibilidade NFAg

O SDK preserva a validação oficial estrita. A correção local da escolha vazia no XSD da NT 2026.004 exige ParseCompatible, MarshalCompatible ou ValidateXMLCompatible; ela não garante recepção pelo autorizador. Consulte docs/COMPATIBILIDADE_NFAG.md e schema-compatibility-manifest.json.
