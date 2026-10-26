// Package fiscaltributariobrasil apresenta o SDK Fiscal Tributário Brasil.
//
// O módulo fornece tipos Go derivados dos schemas oficiais capturados e APIs
// para leitura, serialização e validação estrutural de XML por serviço fiscal.
//
// # Documentos
//
// A matriz inclui NF-e, NFC-e, CT-e, CT-e OS, GTV-e, MDF-e, BP-e, NFCom,
// NF3e, NFS-e Nacional, DC-e, NFAg, NFGas, NF-e ABI, NFS-e Via,
// NF-e Ouro, DAO e DeRE. A DeRE de produção e a prévia técnica são separadas.
//
// # Integração
//
// Escolha o pacote correspondente em bindings e use Parse para validar e
// decodificar, Marshal para serializar e validar ou ValidateXML para validar
// um XML. Os tipos raiz e imports exatos constam em schema-profiles-manifest.json.
// O README e docs/GUIA_DE_USO.md contêm exemplos de integração.
//
// # Limites
//
// A validação estrutural não realiza assinatura digital nem transmissão ao
// autorizador. A NFAg exige escolha explícita das APIs de compatibilidade para
// o defeito documentado no XSD capturado. O modo padrão preserva o original.
//
// O SDK é distribuído sob a licença MIT completa incluída no módulo.
package fiscaltributariobrasil
