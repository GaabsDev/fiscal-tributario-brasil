# Guia de Contribuição — Fiscal Tributário Brasil

Agradecemos pelo interesse em contribuir com o **Fiscal Tributário Brasil**!

## Diretrizes
1. Abra uma Issue detalhando o bug ou melhoria.
2. Certifique-se de que os testes passam com 100% de cobertura: `go test -v ./...`.
3. Formate todo o código com `gofmt -w .`.
4. Não utilize `float32` ou `float64` para valores monetários fiscais.
