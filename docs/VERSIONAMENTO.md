# Fiscal Tributário Brasil — política de versionamento

## Princípios

- O módulo segue versionamento semântico e usa tags Go no formato `vMAJOR.MINOR.PATCH`.
- Cada tag identifica exatamente um commit e é imutável depois de publicada.
- Uma versão já distribuída pelo proxy Go nunca é reutilizada, movida ou sobrescrita.
- Cada tag publicada deve corresponder ao mesmo código da release do GitHub, do Package OCI e da documentação no pkg.go.dev.

## Branches

- `main` contém a próxima versão estável e é a origem normal das releases.
- `release/vX.Y` existe somente quando a linha `vX.Y.z` precisa receber manutenção enquanto `main` já avança para outra linha.
- `feat/*`, `fix/*` e `docs/*` são branches temporárias e devem entrar por pull request.

Não se cria uma tag para representar uma branch. A tag é criada no commit exato aprovado para a versão.

## Publicação

1. Gerar o candidato de release de forma reproduzível.
2. Validar testes, vet, exemplos, documentação, licença, manifestos e hashes.
3. Confirmar que o commit está em `main` ou na branch de manutenção `release/vX.Y` correspondente.
4. Criar uma tag anotada `vX.Y.Z` no commit aprovado.
5. Publicar a GitHub Release, o Package OCI do SDK e solicitar a indexação no proxy Go e no pkg.go.dev.
6. Verificar publicamente todos os artefatos e registrar seus identificadores.

Falhas posteriores exigem uma nova versão. Uma correção para `v0.2.0`, por exemplo, deve ser publicada como `v0.2.1`.
