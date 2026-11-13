# Fiscal Tributário Brasil

Biblioteca Go para representar documentos fiscais eletrônicos brasileiros, ler e gerar XML e validar a estrutura dos documentos contra os schemas oficiais capturados.

## Instalação

Requer Go 1.22 ou superior. Para validar XML, instale também o executável `xmllint` do pacote libxml2.

```sh
go get github.com/GaabsDev/fiscal-tributario-brasil@v0.2.0
```

O nome do projeto é **Fiscal Tributário Brasil**. O caminho usado nos imports é `github.com/GaabsDev/fiscal-tributario-brasil`.

## Documentos e serviços suportados

Esta versão inclui **24 perfis e 470 schemas de entrada**, com tipos Go e validação XSD por serviço.

| Documento | Modelo / referência | Schemas de entrada | Pacotes de tipos e validação |
| --- | --- | ---: | --- |
| NF-e | 55 | 3 | [bindings/nfe](bindings/nfe) |
| NFC-e | 65 | 3 | [bindings/nfce](bindings/nfce) |
| CT-e | 57 | 40 | [bindings/cte](bindings/cte) |
| CT-e | 57 / Simples Nacional (referência técnica) | 42 | [bindings/cte_simples_nacional](bindings/cte_simples_nacional) |
| CT-e OS | 67 | 40 | [bindings/cteos](bindings/cteos) |
| CT-e OS | 67 / Simples Nacional (referência técnica) | 42 | [bindings/cteos_simples_nacional](bindings/cteos_simples_nacional) |
| GTV-e | 64 | 40 | [bindings/gtve](bindings/gtve) |
| GTV-e | 64 / Simples Nacional (referência técnica) | 42 | [bindings/gtve_simples_nacional](bindings/gtve_simples_nacional) |
| MDF-e | 58 | 31 | [bindings/mdfe](bindings/mdfe) |
| BP-e | 63 | 19 | [bindings/bpe](bindings/bpe) |
| BP-e | 63 / Simples Nacional (referência técnica) | 21 | [bindings/bpe_simples_nacional](bindings/bpe_simples_nacional) |
| NFCom | 62 | 13 | [bindings/nfcom](bindings/nfcom) |
| NFCom | 62 / Simples Nacional (referência técnica) | 15 | [bindings/nfcom_simples_nacional](bindings/nfcom_simples_nacional) |
| NF3e | 66 | 15 | [bindings/nf3e](bindings/nf3e) |
| NFS-e Nacional | NFS | 10 | [bindings/nfse](bindings/nfse) |
| DC-e | 99 | 11 | [bindings/dce](bindings/dce) |
| NFAg | 75 | 13 | [bindings/nfag](bindings/nfag) |
| NFGas | 76 | 13 | [bindings/nfgas](bindings/nfgas) |
| NF-e ABI | 77 | 13 | [bindings/nfeabi](bindings/nfeabi) |
| NFS-e Via | NFSVia | 4 | [bindings/nfsevia](bindings/nfsevia) |
| NF-e Ouro | Ouro | 2 | [bindings/nfeouro](bindings/nfeouro) |
| DAO | DAO | 2 | [bindings/dao](bindings/dao) |
| DeRE | Prévia técnica | 31 | [bindings/dere_previa](bindings/dere_previa) |
| DeRE | Produção | 5 | [bindings/dere_producao](bindings/dere_producao) |


Cada linha corresponde a um perfil efetivamente incluído no SDK. A contagem considera os schemas de entrada gerados para aquele perfil, incluindo serviços e eventos; não equivale ao número de serviços de autorização disponíveis em produção.

Os pacotes `bindings` contêm tipos derivados dos XSDs e as APIs `Parse`, `Marshal` e `ValidateXML`. Os pacotes de primeiro nível, como `nfe`, `cte` e `mdfe`, contêm auxiliares com cobertura própria; não substituem os tipos e schemas dos serviços.

Além dos documentos, o pacote `cclasstrib` consulta a tabela tributária capturada, com indicadores por modelo e datas de vigência. O pacote `protocolos` oferece estruturas auxiliares.

## Primeiro uso

Escolha o serviço na [matriz de schemas e serviços](docs/SCHEMAS_E_SERVICOS.md). O exemplo abaixo lê uma NF-e já preenchida e valida seu XML antes de decodificá-la:

```go
package main

import (
    "context"
    "log"
    "os"

    nfe "github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfe/nfe_v4_00"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatal("uso: go run . documento.xml")
    }
    dados, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    var documento nfe.NFe
    if err := nfe.Parse(context.Background(), dados, &documento); err != nil {
        log.Fatal(err)
    }
    log.Print("XML validado e documento decodificado")
}
```

Para outro documento ou evento, use o pacote e o tipo raiz informados no [manifesto dos perfis](schema-profiles-manifest.json). Os exemplos executáveis da pasta `examples` demonstram montagem, campos opcionais, tratamento de erros e serialização com os auxiliares de NF-e.

## Versões oficiais e rastreabilidade

A relação de pacotes capturados, datas de consulta e URLs oficiais está em [Versões e fontes oficiais](docs/VERSOES_OFICIAIS.md). Os manifestos registram hashes, serviços e referências normativas da geração.

A versão do SDK, o leiaute XML, a revisão de uma nota técnica e a revisão do pacote XSD são informações distintas. Publicação de um documento técnico não significa entrada em produção.

## Limites da validação

A validação XSD verifica a estrutura do XML. Ela não define o enquadramento tributário da operação, não substitui regras de negócio específicas e não realiza assinatura digital ou transmissão ao autorizador.

Para a NFAg, consulte [Compatibilidade NFAg](docs/COMPATIBILIDADE_NFAG.md) para saber se o pacote capturado exige uma alternativa explícita de validação. O modo padrão usa o XSD oficial.

Para a NF-e/NFC-e, a [cobertura da NT 2026.005](docs/NT_2026_005.md) distingue os campos e regras implementados diretamente da nota, o pacote XSD oficial ainda sem esses campos e o evento futuro 412140.

As revisões oficiais de 30/10/2026 também estão refletidas no modelo: NT 2025.002 v1.53 (referenciamento de devoluções), NT 2026.006 v1.01 (vinculação de pagamentos), NT 2026.007 v1.11 (emitente exclusivo de IBS/CBS) e NT 2026.008 v1.01 (valores líquidos e ICMS previsto). Alterações marcadas como futuras pela própria nota permanecem documentadas, sem serem apresentadas como exigência já ativa.

A DeRE de produção e a prévia técnica são perfis separados. Uma prévia técnica não deve ser tratada como disponível para recepção em produção.

## Documentação

- [Guia de integração](docs/GUIA_DE_USO.md).
- [Matriz completa de schemas, serviços e imports](docs/SCHEMAS_E_SERVICOS.md).
- [Versões e fontes oficiais](docs/VERSOES_OFICIAIS.md).
- [Cobertura normativa verificada](docs/COBERTURA_NORMATIVA.md).
- [NF-e/NFC-e NT 2026.005](docs/NT_2026_005.md).
- [Conceitos e responsabilidades](docs/CONCEITOS.md).
- [IBS/CBS e classificação tributária](docs/REFORMA_TRIBUTARIA.md).
- [Validações implementadas](docs/VALIDACOES.md).
- [Perguntas frequentes](docs/FAQ.md).
- [Política de versionamento e branches](docs/VERSIONAMENTO.md).
- [Notas de versão](RELEASE_NOTES.md) e [changelog](CHANGELOG.md).
- [Documentação Go](https://pkg.go.dev/github.com/GaabsDev/fiscal-tributario-brasil@v0.2.0).

## Verificar o SDK

Com `xmllint` instalado:

```sh
go test ./...
go vet ./...
go build ./examples/...
```

## Distribuição

O módulo Go é publicado pela tag do repositório e disponibilizado no ecossistema Go. O arquivo completo do SDK, com código, schemas, licença e documentação, é anexado à release do GitHub. Não há imagem Docker pública deste SDK.

## Licença

MIT. Consulte o texto completo em [LICENSE](LICENSE).
