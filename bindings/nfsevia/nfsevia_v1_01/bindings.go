// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package nfsevia_v1_01

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TCCBS struct {
	PAliq TSPercCBSDesc2V2   `xml:"http://www.sped.fazenda.gov.br/nfsevia pAliq" json:"pAliq"`
	VCBS  TSValorCBSDesc15V2 `xml:"http://www.sped.fazenda.gov.br/nfsevia vCBS" json:"vCBS"`
}

type TCDescCondIncond struct {
	VDescIncond *TSDescIncondDesc15V2 `xml:"http://www.sped.fazenda.gov.br/nfsevia vDescIncond,omitempty" json:"vDescIncond,omitempty"`
	VDescCond   *TSDescCondDesc15V2   `xml:"http://www.sped.fazenda.gov.br/nfsevia vDescCond,omitempty" json:"vDescCond,omitempty"`
}

type TCExigSusp struct {
	TpSusp    TSOpExigSuspensa      `xml:"http://www.sped.fazenda.gov.br/nfsevia tpSusp" json:"tpSusp"`
	NProcesso TSNumProcExigSuspensa `xml:"http://www.sped.fazenda.gov.br/nfsevia nProcesso" json:"nProcesso"`
}

type TCExplVia struct {
	CategVeic    TSCategoriaVeiculo `xml:"http://www.sped.fazenda.gov.br/nfsevia categVeic" json:"categVeic"`
	Sentido      TSSentido          `xml:"http://www.sped.fazenda.gov.br/nfsevia sentido" json:"sentido"`
	Placa        TSPlaca            `xml:"http://www.sped.fazenda.gov.br/nfsevia placa" json:"placa"`
	CodAcessoPed TSCodAcessoPedagio `xml:"http://www.sped.fazenda.gov.br/nfsevia codAcessoPed" json:"codAcessoPed"`
	ModPassagem  TSModPassagem      `xml:"http://www.sped.fazenda.gov.br/nfsevia modPassagem" json:"modPassagem"`
	CodContrato  TSCodContrato      `xml:"http://www.sped.fazenda.gov.br/nfsevia codContrato" json:"codContrato"`
	XInfComp     *TSDesc2000        `xml:"http://www.sped.fazenda.gov.br/nfsevia xInfComp,omitempty" json:"xInfComp,omitempty"`
}

type TCFornec struct {
	CNPJ  TSCNPJFornec         `xml:"http://www.sped.fazenda.gov.br/nfsevia CNPJ" json:"CNPJ"`
	XNome *TSNomeFornecDesc150 `xml:"http://www.sped.fazenda.gov.br/nfsevia xNome,omitempty" json:"xNome,omitempty"`
	CMun  TSMunFornec          `xml:"http://www.sped.fazenda.gov.br/nfsevia cMun" json:"cMun"`
}

type TCGIBS struct {
	VIBSTot TSIBSTotDesc15V2 `xml:"http://www.sped.fazenda.gov.br/nfsevia vIBSTot" json:"vIBSTot"`
	UF      TCIBSUF          `xml:"http://www.sped.fazenda.gov.br/nfsevia uf" json:"uf"`
	Mun     TCIBSMun         `xml:"http://www.sped.fazenda.gov.br/nfsevia mun" json:"mun"`
}

type TCIBSCBS struct {
	CST        TSCodCST              `xml:"http://www.sped.fazenda.gov.br/nfsevia CST" json:"CST"`
	CClassTrib TSCodClassTrib        `xml:"http://www.sped.fazenda.gov.br/nfsevia cClassTrib" json:"cClassTrib"`
	CIndOp     TSIndOperacao         `xml:"http://www.sped.fazenda.gov.br/nfsevia cIndOp" json:"cIndOp"`
	VBC        TSBCIBSCBSDesc15V2    `xml:"http://www.sped.fazenda.gov.br/nfsevia vBC" json:"vBC"`
	VLiq       TSLiqIBSCBSDesc15V2   `xml:"http://www.sped.fazenda.gov.br/nfsevia vLiq" json:"vLiq"`
	VTotNF     TSTotNFIBSCBSDesc15V2 `xml:"http://www.sped.fazenda.gov.br/nfsevia vTotNF" json:"vTotNF"`
	GIBS       TCGIBS                `xml:"http://www.sped.fazenda.gov.br/nfsevia gIBS" json:"gIBS"`
	GCBS       TCCBS                 `xml:"http://www.sped.fazenda.gov.br/nfsevia gCBS" json:"gCBS"`
}

type TCIBSMun struct {
	PAliqEfetMun TSAliqEfetMunDesc2V2  `xml:"http://www.sped.fazenda.gov.br/nfsevia pAliqEfetMun" json:"pAliqEfetMun"`
	VIBSMun      TSValorIBSMunDesc15V2 `xml:"http://www.sped.fazenda.gov.br/nfsevia vIBSMun" json:"vIBSMun"`
}

type TCIBSUF struct {
	PAliqEfetUF TSAliqEfetUFDesc2V2  `xml:"http://www.sped.fazenda.gov.br/nfsevia pAliqEfetUF" json:"pAliqEfetUF"`
	VIBSUF      TSValorIBSUFDesc15V2 `xml:"http://www.sped.fazenda.gov.br/nfsevia vIBSUF" json:"vIBSUF"`
}

type TCISSQN struct {
	VBC    TSBCDesc15V2    `xml:"http://www.sped.fazenda.gov.br/nfsevia vBC" json:"vBC"`
	PAliq  TSAliqDesc2V2   `xml:"http://www.sped.fazenda.gov.br/nfsevia pAliq" json:"pAliq"`
	VISSQN TSISSQNDesc15V2 `xml:"http://www.sped.fazenda.gov.br/nfsevia vISSQN" json:"vISSQN"`
}

type TCInfNFSeVia struct {
	ID       TSIDNFSeVia              `xml:"id,attr" json:"id"`
	TpAmb    TSTipoAmbiente           `xml:"http://www.sped.fazenda.gov.br/nfsevia tpAmb" json:"tpAmb"`
	VerAplic TSVersaoEmissor          `xml:"http://www.sped.fazenda.gov.br/nfsevia verAplic" json:"verAplic"`
	CStat    TStat                    `xml:"http://www.sped.fazenda.gov.br/nfsevia cStat" json:"cStat"`
	DhEmi    TSEmissaoDateTimeUTC     `xml:"http://www.sped.fazenda.gov.br/nfsevia dhEmi" json:"dhEmi"`
	DhCompet TSCompetenciaDateTimeUTC `xml:"http://www.sped.fazenda.gov.br/nfsevia dhCompet" json:"dhCompet"`
	Subst    *TCSubst                 `xml:"http://www.sped.fazenda.gov.br/nfsevia subst,omitempty" json:"subst,omitempty"`
	Fornec   TCFornec                 `xml:"http://www.sped.fazenda.gov.br/nfsevia fornec" json:"fornec"`
	Serv     TCServ                   `xml:"http://www.sped.fazenda.gov.br/nfsevia serv" json:"serv"`
	Valores  TCValores                `xml:"http://www.sped.fazenda.gov.br/nfsevia valores" json:"valores"`
}

type TCInfProtNFSeVia struct {
	InfProt TCProtNFSeVia `xml:"http://www.sped.fazenda.gov.br/nfsevia infProt" json:"infProt"`
}

type TCNFSeVia struct {
	Versao     TSVerNFSeVia  `xml:"versao,attr" json:"versao"`
	InfNFSeVia TCInfNFSeVia  `xml:"http://www.sped.fazenda.gov.br/nfsevia infNFSeVia" json:"infNFSeVia"`
	Signature  SignatureType `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCNFSeViaProc struct {
	NFSeVia     TCNFSeVia        `xml:"http://www.sped.fazenda.gov.br/nfsevia NFSeVia" json:"NFSeVia"`
	ProtNFSeVia TCInfProtNFSeVia `xml:"http://www.sped.fazenda.gov.br/nfsevia protNFSeVia" json:"protNFSeVia"`
}

type TCPiscofins struct {
	CST          TSTipoCST              `xml:"http://www.sped.fazenda.gov.br/nfsevia CST" json:"CST"`
	VBCPISCOFINS *TSBCPISCOFINSDesc15V2 `xml:"http://www.sped.fazenda.gov.br/nfsevia vBCPisCofins,omitempty" json:"vBCPisCofins,omitempty"`
	PAliqPIS     *TSAliqPISDesc2V2      `xml:"http://www.sped.fazenda.gov.br/nfsevia pAliqPis,omitempty" json:"pAliqPis,omitempty"`
	PAliqCOFINS  *TSAliqCOFINSDesc2V2   `xml:"http://www.sped.fazenda.gov.br/nfsevia pAliqCofins,omitempty" json:"pAliqCofins,omitempty"`
	VPIS         *TSPISDesc15V2         `xml:"http://www.sped.fazenda.gov.br/nfsevia vPis,omitempty" json:"vPis,omitempty"`
	VCOFINS      *TSCOFINSDesc15V2      `xml:"http://www.sped.fazenda.gov.br/nfsevia vCofins,omitempty" json:"vCofins,omitempty"`
}

type TCProtNFSeVia struct {
	TpAmb     TSTipoAmbiente  `xml:"http://www.sped.fazenda.gov.br/nfsevia tpAmb" json:"tpAmb"`
	VerAplic  TSVersaoEmissor `xml:"http://www.sped.fazenda.gov.br/nfsevia verAplic" json:"verAplic"`
	ChNFSeVia TSChaveNFSeVia  `xml:"http://www.sped.fazenda.gov.br/nfsevia chNFSeVia" json:"chNFSeVia"`
	DhRecbto  TSDateTimeUTC   `xml:"http://www.sped.fazenda.gov.br/nfsevia dhRecbto" json:"dhRecbto"`
	NProt     TSNumProtocolo  `xml:"http://www.sped.fazenda.gov.br/nfsevia nProt" json:"nProt"`
}

type TCServ struct {
	CLocOper   TSLocOperCodMunIBGE `xml:"http://www.sped.fazenda.gov.br/nfsevia cLocOper" json:"cLocOper"`
	XLocOper   TSLocOperDesc150    `xml:"http://www.sped.fazenda.gov.br/nfsevia xLocOper" json:"xLocOper"`
	CPracaEmi  TSPracaEmiCodigo    `xml:"http://www.sped.fazenda.gov.br/nfsevia cPracaEmi" json:"cPracaEmi"`
	XPracaEmi  TSPracaEmiDesc150   `xml:"http://www.sped.fazenda.gov.br/nfsevia xPracaEmi" json:"xPracaEmi"`
	KmPracaEmi TSKmPracaEmi        `xml:"http://www.sped.fazenda.gov.br/nfsevia kmPracaEmi" json:"kmPracaEmi"`
	CLocEmi    TSLocEmiCodMunIBGE  `xml:"http://www.sped.fazenda.gov.br/nfsevia cLocEmi" json:"cLocEmi"`
	XLocEmi    TSLocEmiDesc150     `xml:"http://www.sped.fazenda.gov.br/nfsevia xLocEmi" json:"xLocEmi"`
	NNFSeVia   TSNumNFSeVia        `xml:"http://www.sped.fazenda.gov.br/nfsevia nNFSeVia" json:"nNFSeVia"`
	Serie      TSSerie             `xml:"http://www.sped.fazenda.gov.br/nfsevia serie" json:"serie"`
	CTribNac   TSCodTribNac        `xml:"http://www.sped.fazenda.gov.br/nfsevia cTribNac" json:"cTribNac"`
	CNBS       TSNBS               `xml:"http://www.sped.fazenda.gov.br/nfsevia cNBS" json:"cNBS"`
	ExplVia    TCExplVia           `xml:"http://www.sped.fazenda.gov.br/nfsevia explVia" json:"explVia"`
}

type TCSubst struct {
	ChSubstda TSChaveNFSeVia                 `xml:"http://www.sped.fazenda.gov.br/nfsevia chSubstda" json:"chSubstda"`
	CMotivo   TSMotivoSubstituicao           `xml:"http://www.sped.fazenda.gov.br/nfsevia cMotivo" json:"cMotivo"`
	XMotivo   *TSDescricaoMotivoSubstituicao `xml:"http://www.sped.fazenda.gov.br/nfsevia xMotivo,omitempty" json:"xMotivo,omitempty"`
}

type TCTrib struct {
	ISSQN   TCISSQN    `xml:"http://www.sped.fazenda.gov.br/nfsevia issqn" json:"issqn"`
	TribFed *TCTribFed `xml:"http://www.sped.fazenda.gov.br/nfsevia tribFed,omitempty" json:"tribFed,omitempty"`
	Ibscbs  *TCIBSCBS  `xml:"http://www.sped.fazenda.gov.br/nfsevia ibscbs,omitempty" json:"ibscbs,omitempty"`
}

type TCTribFed struct {
	Piscofins *TCPiscofins `xml:"http://www.sped.fazenda.gov.br/nfsevia piscofins,omitempty" json:"piscofins,omitempty"`
}

type TCValores struct {
	VServ           TSServDesc15V2    `xml:"http://www.sped.fazenda.gov.br/nfsevia vServ" json:"vServ"`
	ModPagamento    TSModPagamento    `xml:"http://www.sped.fazenda.gov.br/nfsevia modPagamento" json:"modPagamento"`
	VDescCondIncond *TCDescCondIncond `xml:"http://www.sped.fazenda.gov.br/nfsevia vDescCondIncond,omitempty" json:"vDescCondIncond,omitempty"`
	Trib            TCTrib            `xml:"http://www.sped.fazenda.gov.br/nfsevia trib" json:"trib"`
}

type TSAliqCOFINSDesc2V2 string

type TSAliqDesc2V2 string

type TSAliqEfetMunDesc2V2 string

type TSAliqEfetUFDesc2V2 string

type TSAliqPISDesc2V2 string

type TSBCDesc15V2 string

type TSBCIBSCBSDesc15V2 string

type TSBCPISCOFINSDesc15V2 string

type TSCNPJ string

type TSCNPJFornec string

type TSCPF string

type TSCategoriaVeiculo string

type TSChaveNFSeVia string

type TSCodAcessoPedagio string

type TSCodCST string

type TSCodClassTrib string

type TSCodContrato string

type TSCodJustCanc string

type TSCodJustSubst string

type TSCodMotivoRejeicaoManif string

type TSCodMunIBGE string

type TSCodNaoNIF string

type TSCodTribNac string

type TSCOFINSDesc15V2 string

type TSCompetenciaDateTimeUTC string

type TSData string

type TSDateTimeUTC string

type TSDec15V2 string

type TSDec2V2 string

type TSDec3V2 string

type TSDec7V3 string

type TSDec7V6 string

type TSDesc150 string

type TSDesc2000 string

type TSDesc600 string

type TSDescCondDesc15V2 string

type TSDescIncondDesc15V2 string

type TSDescricaoMotivoSubstituicao string

type TSEmissaoDateTimeUTC string

type TSIBSTotDesc15V2 string

type TSISSQNDesc15V2 string

type TSIDEvento string

type TSIDNFSeVia string

type TSIndOperacao string

type TSKmPracaEmi string

type TSLiqIBSCBSDesc15V2 string

type TSLocEmiCodMunIBGE string

type TSLocEmiDesc150 string

type TSLocOperCodMunIBGE string

type TSLocOperDesc150 string

type TSModPagamento string

type TSModPassagem string

type TSMotivo string

type TSMotivoSubstituicao string

type TSMunFornec string

type TSNBS string

type TSNIF string

type TSNomeFornecDesc150 string

type TSNum3Dig string

type TSNumNFSeVia string

type TSNumProcExigSuspensa string

type TSNumProtocolo string

type TSOpExigSuspensa string

type TSPercCBSDesc2V2 string

type TSPISDesc15V2 string

type TSPlaca string

type TSPracaEmiCodigo string

type TSPracaEmiDesc150 string

type TSRetCPDesc15V2 string

type TSRetCSLLDesc15V2 string

type TSRetIRRFDesc15V2 string

type TSRodagem string

type TSSentido string

type TSSerie string

type TSServDesc15V2 string

type TSString string

type TSStringComQuebraDeLinha string

type TSTipoAmbiente string

type TSTipoCST string

type TSTipoRetPISCOFINS string

type TSTotNFIBSCBSDesc15V2 string

type TSTribISSQN string

type TSValorCBSDesc15V2 string

type TSValorIBSMunDesc15V2 string

type TSValorIBSUFDesc15V2 string

type TSVerNFSeVia string

type TSVersaoEmissor string

type TStat string

type DigestValueType string

type KeyInfoType struct {
	ID       *ID          `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	X509Data X509DataType `xml:"http://www.w3.org/2000/09/xmldsig# X509Data" json:"X509Data"`
}

type ReferenceTypeDigestMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type ReferenceType struct {
	ID           *ID                       `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	URI          string                    `xml:"URI,attr" json:"URI"`
	TypeField    *AnyURI                   `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	Transforms   TransformsType            `xml:"http://www.w3.org/2000/09/xmldsig# Transforms" json:"Transforms"`
	DigestMethod ReferenceTypeDigestMethod `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod" json:"DigestMethod"`
	DigestValue  DigestValueType           `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue" json:"DigestValue"`
}

type SignatureType struct {
	ID             *ID                `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	SignedInfo     SignedInfoType     `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo" json:"SignedInfo"`
	SignatureValue SignatureValueType `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue" json:"SignatureValue"`
	KeyInfo        KeyInfoType        `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo" json:"KeyInfo"`
}

type SignatureValueType struct {
	Value Base64Binary `xml:",chardata" json:"value"`
	ID    *ID          `xml:"Id,attr,omitempty" json:"Id,omitempty"`
}

type SignedInfoTypeCanonicalizationMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type SignedInfoTypeSignatureMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type SignedInfoType struct {
	ID                     *ID                                  `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	CanonicalizationMethod SignedInfoTypeCanonicalizationMethod `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod" json:"CanonicalizationMethod"`
	SignatureMethod        SignedInfoTypeSignatureMethod        `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod" json:"SignatureMethod"`
	Reference              ReferenceType                        `xml:"http://www.w3.org/2000/09/xmldsig# Reference" json:"Reference"`
}

type TTransformURI string

type TransformType struct {
	Algorithm TTransformURI `xml:"Algorithm,attr" json:"Algorithm"`
	XPath     []String      `xml:"http://www.w3.org/2000/09/xmldsig# XPath,omitempty" json:"XPath,omitempty"`
}

type TransformsType struct {
	Transform []TransformType `xml:"http://www.w3.org/2000/09/xmldsig# Transform,omitempty" json:"Transform,omitempty"`
}

type X509DataType struct {
	X509Certificate Base64Binary `xml:"http://www.w3.org/2000/09/xmldsig# X509Certificate" json:"X509Certificate"`
}

type ID string

type AnyURI string

type Base64Binary string

type String string

// NFSeVia represents the global element {http://www.sped.fazenda.gov.br/nfsevia}NFSeVia.
type NFSeVia struct {
	XMLName xml.Name `xml:"http://www.sped.fazenda.gov.br/nfsevia NFSeVia" json:"-"`
	TCNFSeVia
}

// Signature represents the global element {http://www.w3.org/2000/09/xmldsig#}Signature.
type Signature struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"-"`
	SignatureType
}

// MarshalXML avoids re-emitting a default xmlns already supplied by encoding/xml.
func (value RawElement) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	if value.XMLName.Local != "" {
		start.Name = value.XMLName
	}
	for _, attr := range value.Attrs {
		if attr.Name.Space == "" && attr.Name.Local == "xmlns" && attr.Value == start.Name.Space {
			continue
		}
		if attr.Name.Space == "xmlns" {
			attr.Name = xml.Name{Local: "xmlns:" + attr.Name.Local}
		}
		start.Attr = append(start.Attr, attr)
	}
	body := struct {
		InnerXML string `xml:",innerxml"`
	}{value.InnerXML}
	return encoder.EncodeElement(body, start)
}
