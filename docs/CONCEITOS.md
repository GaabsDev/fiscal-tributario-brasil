# Fiscal Tributário Brasil — conceitos

O SDK fornece modelos de dados e métodos de validação e serialização. Seu sistema preenche os dados da operação, decide a tributação e integra os serviços fiscais necessários.

## Fluxo de integração

1. Escolha o pacote correspondente ao documento.
2. Preencha os campos usando os tipos do pacote e dados reais da operação.
3. Execute Validate e trate os erros retornados.
4. Serialize o documento com ToXML ou ToXMLIndent, verificando o erro.
5. Valide o XML contra o pacote XSD aplicável ao ambiente e à revisão utilizados.
6. Integre assinatura, transmissão e tratamento de respostas no seu sistema.

Validate executa verificações implementadas em Go. Ele não substitui a validação XSD nem garante autorização pelo Fisco.

Use Decimal para valores monetários e respeite as regras de escala e arredondamento da operação. Não transforme exemplos fictícios em regras tributárias universais.

As versões do XML, do manual, da nota técnica e do pacote XSD são informações diferentes. Consulte o manifesto de fontes da versão.
