# Changelog — Fiscal Tributário Brasil

Todas as alterações notáveis neste projeto são documentadas deterministicamente neste arquivo.
O formato baseia-se no [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/) e adere ao [Semantic Versioning](https://semver.org/lang/pt-BR/).

## [v0.1.0] - 2026-10-26

### Especificações

- Documentos e serviços gerados: consulte a matriz do README e docs/SCHEMAS_E_SERVICOS.md.
- Referências normativas e schemas capturados: consulte schema-profiles-manifest.json e fiscal-sources-manifest.json.
- Fiscal IR Hash: `649b403f43c9daeff0c631434c7ce475c5a6a8c8314fe92f34c79bb34a0e5986`
- Source Set Hash: `a469cd2f828a9df07d23d6fdd34486a5fa2a0a3e738763ebc0033aca57748322`

### Regras fiscais

- Regra I05k-10 (Rejeição oficial 1165): validação em NFC-e modelo 65

### Adicionado

- func CalculateChaveAcessoDV
- func DecimalPtr
- func GenerateChaveAcesso
- func GetDocumentedRules
- func GetSpecificationMetadata
- func MarshalXML
- func NewDate
- func NewDateTime
- func NewDecimal
- func NewDecimalFromFloat
- func NewNFe
- func RegexHelper
- func StringPtr
- func SupportedSpecifications
- func UnmarshalJSON
- func UnmarshalXML
- func Validate
- method (CClassTrib).Definition
- method (CClassTrib).IsCombustivel
- method (CClassTrib).String
- method (CClassTrib).Validate
- method (CClassTrib).ValidateForModel
- method (CFOP).String
- method (CFOP).Validate
- method (CNPJ).String
- method (CNPJ).Validate
- method (CPF).String
- method (CPF).Validate
- method (CSOSN).String
- method (CSOSN).Validate
- method (CST).String
- method (CST).Validate
- method (ChaveAcesso).String
- method (ChaveAcesso).Validate
- method (CodigoMunicipio).String
- method (CodigoMunicipio).Validate
- method (Date).String
- method (Date).Validate
- method (DateTime).String
- method (DateTime).Validate
- method (Decimal).Float64
- method (Decimal).IsZero
- method (Decimal).String
- method (Decimal).Validate
- method (FinalidadeNFe).String
- method (FinalidadeNFe).Validate
- method (IBSCBSTot).VCBS
- method (IBSCBSTot).VIBS
- method (ModeloDocumento).String
- method (ModeloDocumento).Validate
- method (NCM).String
- method (NCM).Validate
- method (NFe).AddItem
- method (NFe).AddPagamento
- method (NFe).MarshalJSON
- method (NFe).SetDestinatario
- method (NFe).SetEmitente
- method (NFe).SetTotal
- method (NFe).ToJSON
- method (NFe).ToJSONIndent
- method (NFe).ToXML
- method (NFe).ToXMLCompact
- method (NFe).ToXMLIndent
- method (NFe).Validate
- method (TAmb).String
- method (TAmb).Validate
- method (TCOrgaoIBGE).String
- method (TCOrgaoIBGE).Validate
- method (TCodUfIBGE).String
- method (TCodUfIBGE).Validate
- method (TEnteGov).String
- method (TEnteGov).Validate
- method (TIndDoacao).String
- method (TIndDoacao).Validate
- method (TOperCompraGov).String
- method (TOperCompraGov).Validate
- method (TProcEmi).String
- method (TProcEmi).Validate
- method (TTpCredPresIBSZFM).String
- method (TTpCredPresIBSZFM).Validate
- method (TTpNFCredito).String
- method (TTpNFCredito).Validate
- method (TTpNFDebito).String
- method (TTpNFDebito).Validate
- method (TTransformURI).String
- method (TTransformURI).Validate
- method (TUf).String
- method (TUf).Validate
- method (TUfEmi).String
- method (TUfEmi).Validate
- method (TipoAmbiente).String
- method (TipoAmbiente).Validate
- method (TipoEmissao).String
- method (TipoEmissao).Validate
- method (Torig).String
- method (Torig).Validate
- method (Tpais).String
- method (Tpais).Validate
- method (UF).String
- method (UF).Validate
- method (ValidationError).Error
- method (ValidationErrors).AddBusinessRule
- method (ValidationErrors).AddStructural
- method (ValidationErrors).Error
- method (ValidationErrors).Errors
- method (ValidationErrors).HasErrors
- method (ValidationErrors).Warnings
- struct Adi
- struct Arma
- struct AutXML
- struct Avulsa
- struct C14NMethod
- struct CIDE
- struct COFINS
- struct COFINSAliq
- struct COFINSNT
- struct COFINSST
- struct Cana
- struct Card
- struct Cobr
- struct Comb
- struct Compra
- struct DI
- struct Deduc
- struct Dest
- struct Det
- struct DetExport
- struct DetPag
- struct DigestMethod
- struct DocumentedRule
- struct Dup
- struct Emit
- struct Encerrante
- struct EnderDest
- struct EnderEmit
- struct ExportInd
- struct Exporta
- struct Fat
- struct ForDia
- struct GAjusteCompet
- struct GCBS
- struct GCBSTot
- struct GCredPresIBSZFM
- struct GCredPresOper
- struct GDevTrib
- struct GDif
- struct GEstornoCred
- struct GEstornoCredTot
- struct GIBS
- struct GIBSCBS
- struct GIBSCBSMono
- struct GIBSMun
- struct GIBSMunTot
- struct GIBSUF
- struct GIBSUFTot
- struct GMonoTot
- struct GRed
- struct GTransfCred
- struct IBSCBS
- struct IBSCBSTot
- struct ICMS
- struct ICMS00
- struct ICMS10
- struct ICMS20
- struct ICMS30
- struct ICMS40
- struct ICMS51
- struct ICMS60
- struct ICMS70
- struct ICMS90
- struct ICMSSN101
- struct ICMSSN102
- struct ICMSSN201
- struct ICMSSN202
- struct ICMSSN500
- struct ICMSSN900
- struct ICMSTot
- struct ICMSUFDest
- struct II
- struct IPI
- struct IPINT
- struct IPITrib
- struct IS
- struct ISSQN
- struct ISSQNtot
- struct ISTot
- struct Ide
- struct Imposto
- struct InfAdic
- struct InfIntermed
- struct InfNFe
- struct InfNFeSupl
- struct InfRespTec
- struct KeyInfo
- struct Lacres
- struct Local
- struct Med
- struct NFe
- struct NFref
- struct ObsCont
- struct ObsFisco
- struct PIS
- struct PISAliq
- struct PISNT
- struct PISST
- struct Pag
- struct ProcRef
- struct Prod
- struct Rastro
- struct Reboque
- struct Reference
- struct RetTransp
- struct RetTrib
- struct SigMethod
- struct Signature
- struct SignedInfo
- struct SpecMetadata
- struct Total
- struct Transform
- struct Transforms
- struct Transp
- struct Transporta
- struct ValidationError
- struct VeicProd
- struct VeicTransp
- struct Vol
- struct X509Data
- type CClassTrib
- type CFOP
- type CNPJ
- type CPF
- type CSOSN
- type CST
- type ChaveAcesso
- type CodigoMunicipio
- type Date
- type DateTime
- type Decimal
- type FinalidadeNFe
- type ModeloDocumento
- type NCM
- type RuleStatus
- type Severity
- type TAmb
- type TCOrgaoIBGE
- type TCodUfIBGE
- type TEnteGov
- type TIndDoacao
- type TOperCompraGov
- type TProcEmi
- type TTpCredPresIBSZFM
- type TTpNFCredito
- type TTpNFDebito
- type TTransformURI
- type TUf
- type TUfEmi
- type TipoAmbiente
- type TipoEmissao
- type Torig
- type Tpais
- type UF
- type ValidationErrors
- type ValidationSource

### Correções

- Consulte as alterações de API e regras acima; correções não são presumidas sem evidências.


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
