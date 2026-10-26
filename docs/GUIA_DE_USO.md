# Fiscal Tributário Brasil — guia de integração

## Instalar

```sh
go mod init exemplo-integracao
go get github.com/GaabsDev/fiscal-tributario-brasil@v0.1.0
```

## Consultar uma classificação tributária

Crie main.go:

```go
package main

import (
    "fmt"
    "log"
    "github.com/GaabsDev/fiscal-tributario-brasil/cclasstrib"
)

func main() {
    classificacao, existe := cclasstrib.Get("000001")
    if !existe {
        log.Fatal("classificação ausente do catálogo desta versão")
    }
    fmt.Println(classificacao.Code, classificacao.Description)
    fmt.Println("Indicador para NF-e:", classificacao.Applicability["55"])
    fmt.Println("Vigência:", classificacao.ValidFrom, classificacao.ValidTo)
}
```

Execute go run . no diretório do exemplo. A consulta demonstra a API; a classificação adequada depende da operação. Um registro presente no catálogo pode ter datas que limitem sua vigência.

## Montar uma NF-e

Consulte [o exemplo completo](../examples/01_basic_nfe/main.go). A partir da raiz do repositório, execute go run ./examples/01_basic_nfe.

O exemplo usa dados fictícios e demonstra construção e validação local. Para sua integração, preencha emitente, destinatário, identificação, itens, tributos, totais, transporte e pagamento conforme a operação e a revisão fiscal adotadas.

## Tratar erros

```go
if err := doc.Validate(); err != nil {
    return fmt.Errorf("documento inválido: %w", err)
}
xmlBytes, err := doc.ToXMLIndent("", "  ")
if err != nil {
    return fmt.Errorf("não foi possível gerar XML: %w", err)
}
// Encaminhe xmlBytes para a validação XSD e para as próximas etapas da integração.
```

Esse trecho pressupõe doc preenchido, importação de fmt e uma função que retorne error. O exemplo de erros estruturados está em [examples/03_validation_errors](../examples/03_validation_errors/main.go).

## Validar o XML e integrar o autorizador

Utilize o schema raiz e todas as dependências da revisão selecionada. Não misture schemas de pacotes diferentes. Uma validação local bem-sucedida não comprova assinatura, transmissão ou autorização.

Consulte [VALIDACOES.md](VALIDACOES.md) e [FONTES_OFICIAIS.md](FONTES_OFICIAIS.md) para distinguir a cobertura e os artefatos utilizados.
