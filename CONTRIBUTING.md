# Guia de Contribuição — Fiscal Tributário Brasil

Agradecemos pelo interesse em contribuir com o Fiscal Tributário Brasil.

## Fluxo de contribuição

1. Abra uma issue detalhando o erro, a atualização normativa ou a melhoria.
2. Crie uma branch temporária com prefixo `fix/`, `feat/` ou `docs/`.
3. Envie um pull request para `main`. Correções destinadas a uma linha de manutenção devem usar a branch `release/vX.Y` correspondente.
4. Aguarde a execução integral do CI e a revisão do mantenedor antes do merge.

Contribuições externas nunca são incorporadas automaticamente. Somente pessoas com permissão de escrita podem concluir o merge no repositório.

## Validação obrigatória

- Execute `gofmt` nos arquivos Go modificados.
- Execute `go test -race ./...`, `go vet ./...` e `go build ./examples/...`.
- Não utilize `float32` ou `float64` para valores monetários fiscais.
- Atualize a documentação e as evidências normativas quando a mudança alterar a cobertura do SDK.

## Versões

Não crie, mova, reutilize ou exclua tags de versão. Tags e releases são produzidas pelo fluxo de publicação mantido pelo projeto. Consulte [a política de versionamento](docs/VERSIONAMENTO.md).
