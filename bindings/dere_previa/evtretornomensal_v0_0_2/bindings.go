// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package evtretornomensal_v0_0_2

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type CanonicalizationMethodType struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type CryptoBinary string

type DSAKeyValueType struct {
	P           *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# P,omitempty" json:"P,omitempty"`
	Q           *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# Q,omitempty" json:"Q,omitempty"`
	G           *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# G,omitempty" json:"G,omitempty"`
	Y           CryptoBinary  `xml:"http://www.w3.org/2000/09/xmldsig# Y" json:"Y"`
	J           *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# J,omitempty" json:"J,omitempty"`
	Seed        *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# Seed,omitempty" json:"Seed,omitempty"`
	PgenCounter *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# PgenCounter,omitempty" json:"PgenCounter,omitempty"`
}

type DigestMethodType struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type DigestValueType string

type HMACOutputLengthType string

type KeyInfoType struct {
	ID *ID `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type KeyValueType struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type ManifestType struct {
	ID        *ID             `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	Reference []ReferenceType `xml:"http://www.w3.org/2000/09/xmldsig# Reference,omitempty" json:"Reference,omitempty"`
}

type ObjectType struct {
	ID       *ID     `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	MimeType *String `xml:"MimeType,attr,omitempty" json:"MimeType,omitempty"`
	Encoding *AnyURI `xml:"Encoding,attr,omitempty" json:"Encoding,omitempty"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type PGPDataType struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type RSAKeyValueType struct {
	Modulus  CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# Modulus" json:"Modulus"`
	Exponent CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# Exponent" json:"Exponent"`
}

type ReferenceType struct {
	ID           *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	URI          *AnyURI          `xml:"URI,attr,omitempty" json:"URI,omitempty"`
	TypeField    *AnyURI          `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	Transforms   *TransformsType  `xml:"http://www.w3.org/2000/09/xmldsig# Transforms,omitempty" json:"Transforms,omitempty"`
	DigestMethod DigestMethodType `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod" json:"DigestMethod"`
	DigestValue  DigestValueType  `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue" json:"DigestValue"`
}

type RetrievalMethodType struct {
	URI        *AnyURI         `xml:"URI,attr,omitempty" json:"URI,omitempty"`
	TypeField  *AnyURI         `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	Transforms *TransformsType `xml:"http://www.w3.org/2000/09/xmldsig# Transforms,omitempty" json:"Transforms,omitempty"`
}

type SPKIDataType struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type SignatureMethodType struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type SignaturePropertiesType struct {
	ID                *ID                     `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	SignatureProperty []SignaturePropertyType `xml:"http://www.w3.org/2000/09/xmldsig# SignatureProperty,omitempty" json:"SignatureProperty,omitempty"`
}

type SignaturePropertyType struct {
	Target AnyURI `xml:"Target,attr" json:"Target"`
	ID     *ID    `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type SignatureType struct {
	ID             *ID                `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	SignedInfo     SignedInfoType     `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo" json:"SignedInfo"`
	SignatureValue SignatureValueType `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue" json:"SignatureValue"`
	KeyInfo        *KeyInfoType       `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo,omitempty" json:"KeyInfo,omitempty"`
	Object         []ObjectType       `xml:"http://www.w3.org/2000/09/xmldsig# Object,omitempty" json:"Object,omitempty"`
}

type SignatureValueType struct {
	Value Base64Binary `xml:",chardata" json:"value"`
	ID    *ID          `xml:"Id,attr,omitempty" json:"Id,omitempty"`
}

type SignedInfoType struct {
	ID                     *ID                        `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	CanonicalizationMethod CanonicalizationMethodType `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod" json:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethodType        `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod" json:"SignatureMethod"`
	Reference              []ReferenceType            `xml:"http://www.w3.org/2000/09/xmldsig# Reference,omitempty" json:"Reference,omitempty"`
}

type TransformType struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type TransformsType struct {
	Transform []TransformType `xml:"http://www.w3.org/2000/09/xmldsig# Transform,omitempty" json:"Transform,omitempty"`
}

type X509DataType struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type X509IssuerSerialType struct {
	X509IssuerName   String  `xml:"http://www.w3.org/2000/09/xmldsig# X509IssuerName" json:"X509IssuerName"`
	X509SerialNumber Integer `xml:"http://www.w3.org/2000/09/xmldsig# X509SerialNumber" json:"X509SerialNumber"`
}

type ID string

type AnyURI string

type Base64Binary string

type Integer string

type String string

type DeREContentEvtRetornoMensalIdeContrib struct {
	NrInsc string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 nrInsc" json:"nrInsc"`
}

type DeREContentEvtRetornoMensalIdeStatusOcorrencias struct {
	Codigo      string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codigo" json:"codigo"`
	Descricao   string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 descricao" json:"descricao"`
	Tipo        string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 tipo" json:"tipo"`
	Localizacao *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 localizacao,omitempty" json:"localizacao,omitempty"`
}

type DeREContentEvtRetornoMensalIdeStatus struct {
	CdRetorno   string                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 cdRetorno" json:"cdRetorno"`
	DescRetorno string                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 descRetorno" json:"descRetorno"`
	Ocorrencias []DeREContentEvtRetornoMensalIdeStatusOcorrencias `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 ocorrencias,omitempty" json:"ocorrencias,omitempty"`
}

type DeREContentEvtRetornoMensalInfoRecEv struct {
	NrRecibo      *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 nrRecibo,omitempty" json:"nrRecibo,omitempty"`
	SeqEvento     *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 seqEvento,omitempty" json:"seqEvento,omitempty"`
	ProtocoloLote *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 protocoloLote,omitempty" json:"protocoloLote,omitempty"`
	DhRecepcao    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 dhRecepcao" json:"dhRecepcao"`
	DhProcess     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 dhProcess" json:"dhProcess"`
	TpEv          string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 tpEv" json:"tpEv"`
	Hash          string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 hash" json:"hash"`
}

type DeREContentEvtRetornoMensalInfoEventoIdePeriodo struct {
	PerApur string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 perApur" json:"perApur"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoAdic struct {
	NrReciboBalancete  string   `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 nrReciboBalancete" json:"nrReciboBalancete"`
	NrReciboAplicFin   *string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 nrReciboAplicFin,omitempty" json:"nrReciboAplicFin,omitempty"`
	NrReciboRelDedu    []string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 nrReciboRelDedu,omitempty" json:"nrReciboRelDedu,omitempty"`
	NrReciboOperTitPub *string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 nrReciboOperTitPub,omitempty" json:"nrReciboOperTitPub,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCGCoeficientes struct {
	PRateioDespCapt *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pRateioDespCapt,omitempty" json:"pRateioDespCapt,omitempty"`
	PRevDedAtoCoop  *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pRevDedAtoCoop,omitempty" json:"pRevDedAtoCoop,omitempty"`
	PExportIaV      *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pExportIaV,omitempty" json:"pExportIaV,omitempty"`
	PExportVI       *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pExportVI,omitempty" json:"pExportVI,omitempty"`
	PExportIX       *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pExportIX,omitempty" json:"pExportIX,omitempty"`
	PExportXI       *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pExportXI,omitempty" json:"pExportXI,omitempty"`
	PExportXIII     *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pExportXIII,omitempty" json:"pExportXIII,omitempty"`
	PExportXIV      *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pExportXIV,omitempty" json:"pExportXIV,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCGBCIBS struct {
	VBCIBS     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCIBS" json:"vBCIBS"`
	VBCNIBS    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCNIBS,omitempty" json:"vBCNIBS,omitempty"`
	VDedBCN    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vDedBCN,omitempty" json:"vDedBCN,omitempty"`
	VBCApurIBS string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCApurIBS" json:"vBCApurIBS"`
	PIBSMun    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBSMun" json:"pIBSMun"`
	VIBSMun    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSMun" json:"vIBSMun"`
	PIBSUF     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBSUF" json:"pIBSUF"`
	VIBSUF     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSUF" json:"vIBSUF"`
	PIBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBS" json:"pIBS"`
	VIBSTot    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSTot" json:"vIBSTot"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCGBCCBS struct {
	VBCCBS     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCCBS" json:"vBCCBS"`
	VBCNCBS    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCNCBS,omitempty" json:"vBCNCBS,omitempty"`
	VDedBCN    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vDedBCN,omitempty" json:"vDedBCN,omitempty"`
	VBCApurCBS string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCApurCBS" json:"vBCApurCBS"`
	PCBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pCBS" json:"pCBS"`
	VCBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vCBS" json:"vCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCNGBCNIBSDetBCN struct {
	CodBCN      string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCN" json:"codBCN"`
	PerOrigem   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 perOrigem" json:"perOrigem"`
	VOrigemIni  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vOrigemIni" json:"vOrigemIni"`
	VSaldoAnt   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCNGBCNIBS struct {
	CodBCNRaiz  string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCNRaiz" json:"codBCNRaiz"`
	VSaldoAnt   string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
	QtdOrigens  string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 qtdOrigens" json:"qtdOrigens"`
	DetBCN      []DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCNGBCNIBSDetBCN `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBCN,omitempty" json:"detBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCNGBCNCBSDetBCN struct {
	CodBCN      string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCN" json:"codBCN"`
	PerOrigem   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 perOrigem" json:"perOrigem"`
	VOrigemIni  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vOrigemIni" json:"vOrigemIni"`
	VSaldoAnt   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCNGBCNCBS struct {
	CodBCNRaiz  string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCNRaiz" json:"codBCNRaiz"`
	VSaldoAnt   string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
	QtdOrigens  string                                                                            `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 qtdOrigens" json:"qtdOrigens"`
	DetBCN      []DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCNGBCNCBSDetBCN `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBCN,omitempty" json:"detBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCN struct {
	GBCNIBS DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCNGBCNIBS `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCNIBS" json:"gBCNIBS"`
	GBCNCBS DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCNGBCNCBS `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCNCBS" json:"gBCNCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBC struct {
	CodBC          string                                                                    `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBC" json:"codBC"`
	XDetBC         string                                                                    `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 xDetBC" json:"xDetBC"`
	GCoeficientes  *DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCGCoeficientes `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gCoeficientes,omitempty" json:"gCoeficientes,omitempty"`
	MemoriaCalculo string                                                                    `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 memoriaCalculo" json:"memoriaCalculo"`
	GBCIBS         DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCGBCIBS         `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCIBS" json:"gBCIBS"`
	GBCCBS         DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCGBCCBS         `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCCBS" json:"gBCCBS"`
	InfoBCN        *DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBCInfoBCN       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoBCN,omitempty" json:"infoBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroTotalTributos struct {
	VIBSMun string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSMun" json:"vIBSMun"`
	VIBSUF  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSUF" json:"vIBSUF"`
	VIBSTot string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSTot" json:"vIBSTot"`
	VCBS    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vCBS" json:"vCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiro struct {
	DetBC         []DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroDetBC       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBC,omitempty" json:"detBC,omitempty"`
	TotalTributos DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiroTotalTributos `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 totalTributos" json:"totalTributos"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCGBCIBS struct {
	VBCIBS     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCIBS" json:"vBCIBS"`
	VBCNIBS    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCNIBS,omitempty" json:"vBCNIBS,omitempty"`
	VDedBCN    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vDedBCN,omitempty" json:"vDedBCN,omitempty"`
	VBCApurIBS string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCApurIBS" json:"vBCApurIBS"`
	PIBSMun    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBSMun" json:"pIBSMun"`
	VIBSMun    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSMun" json:"vIBSMun"`
	PIBSUF     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBSUF" json:"pIBSUF"`
	VIBSUF     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSUF" json:"vIBSUF"`
	PIBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBS" json:"pIBS"`
	VIBSTot    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSTot" json:"vIBSTot"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCGBCCBS struct {
	VBCCBS     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCCBS" json:"vBCCBS"`
	VBCNCBS    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCNCBS,omitempty" json:"vBCNCBS,omitempty"`
	VDedBCN    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vDedBCN,omitempty" json:"vDedBCN,omitempty"`
	VBCApurCBS string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCApurCBS" json:"vBCApurCBS"`
	PCBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pCBS" json:"pCBS"`
	VCBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vCBS" json:"vCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCNGBCNIBSDetBCN struct {
	CodBCN      string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCN" json:"codBCN"`
	PerOrigem   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 perOrigem" json:"perOrigem"`
	VOrigemIni  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vOrigemIni" json:"vOrigemIni"`
	VSaldoAnt   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCNGBCNIBS struct {
	CodBCNRaiz  string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCNRaiz" json:"codBCNRaiz"`
	VSaldoAnt   string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
	QtdOrigens  string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 qtdOrigens" json:"qtdOrigens"`
	DetBCN      []DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCNGBCNIBSDetBCN `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBCN,omitempty" json:"detBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCNGBCNCBSDetBCN struct {
	CodBCN      string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCN" json:"codBCN"`
	PerOrigem   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 perOrigem" json:"perOrigem"`
	VOrigemIni  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vOrigemIni" json:"vOrigemIni"`
	VSaldoAnt   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCNGBCNCBS struct {
	CodBCNRaiz  string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCNRaiz" json:"codBCNRaiz"`
	VSaldoAnt   string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
	QtdOrigens  string                                                                       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 qtdOrigens" json:"qtdOrigens"`
	DetBCN      []DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCNGBCNCBSDetBCN `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBCN,omitempty" json:"detBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCN struct {
	GBCNIBS DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCNGBCNIBS `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCNIBS" json:"gBCNIBS"`
	GBCNCBS DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCNGBCNCBS `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCNCBS" json:"gBCNCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBC struct {
	CodBC          string                                                         `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBC" json:"codBC"`
	XDetBC         string                                                         `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 xDetBC" json:"xDetBC"`
	MemoriaCalculo string                                                         `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 memoriaCalculo" json:"memoriaCalculo"`
	GBCIBS         DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCGBCIBS   `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCIBS" json:"gBCIBS"`
	GBCCBS         DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCGBCCBS   `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCCBS" json:"gBCCBS"`
	InfoBCN        *DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBCInfoBCN `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoBCN,omitempty" json:"infoBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeTotalTributos struct {
	VIBSMun string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSMun" json:"vIBSMun"`
	VIBSUF  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSUF" json:"vIBSUF"`
	VIBSTot string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSTot" json:"vIBSTot"`
	VCBS    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vCBS" json:"vCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotSaude struct {
	DetBC         []DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeDetBC       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBC,omitempty" json:"detBC,omitempty"`
	TotalTributos DeREContentEvtRetornoMensalInfoEventoInfoTotSaudeTotalTributos `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 totalTributos" json:"totalTributos"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCGCoeficientes struct {
	PExportProg *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pExportProg,omitempty" json:"pExportProg,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCGBCIBS struct {
	VBCIBS     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCIBS" json:"vBCIBS"`
	VBCNIBS    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCNIBS,omitempty" json:"vBCNIBS,omitempty"`
	VDedBCN    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vDedBCN,omitempty" json:"vDedBCN,omitempty"`
	VBCApurIBS string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCApurIBS" json:"vBCApurIBS"`
	PIBSMun    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBSMun" json:"pIBSMun"`
	VIBSMun    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSMun" json:"vIBSMun"`
	PIBSUF     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBSUF" json:"pIBSUF"`
	VIBSUF     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSUF" json:"vIBSUF"`
	PIBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIBS" json:"pIBS"`
	VIBSTot    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSTot" json:"vIBSTot"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCGBCCBS struct {
	VBCCBS     string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCCBS" json:"vBCCBS"`
	VBCNCBS    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCNCBS,omitempty" json:"vBCNCBS,omitempty"`
	VDedBCN    *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vDedBCN,omitempty" json:"vDedBCN,omitempty"`
	VBCApurCBS string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCApurCBS" json:"vBCApurCBS"`
	PCBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pCBS" json:"pCBS"`
	VCBS       string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vCBS" json:"vCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCGBCIS struct {
	VBCIS     string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCIS" json:"vBCIS"`
	VBCApurIS string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vBCApurIS" json:"vBCApurIS"`
	PIS       string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 pIS" json:"pIS"`
	VIS       string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIS" json:"vIS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCNGBCNIBSDetBCN struct {
	CodBCN      string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCN" json:"codBCN"`
	PerOrigem   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 perOrigem" json:"perOrigem"`
	VOrigemIni  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vOrigemIni" json:"vOrigemIni"`
	VSaldoAnt   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCNGBCNIBS struct {
	CodBCNRaiz  string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCNRaiz" json:"codBCNRaiz"`
	VSaldoAnt   string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
	QtdOrigens  string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 qtdOrigens" json:"qtdOrigens"`
	DetBCN      []DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCNGBCNIBSDetBCN `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBCN,omitempty" json:"detBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCNGBCNCBSDetBCN struct {
	CodBCN      string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCN" json:"codBCN"`
	PerOrigem   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 perOrigem" json:"perOrigem"`
	VOrigemIni  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vOrigemIni" json:"vOrigemIni"`
	VSaldoAnt   string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCNGBCNCBS struct {
	CodBCNRaiz  string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBCNRaiz" json:"codBCNRaiz"`
	VSaldoAnt   string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoAnt" json:"vSaldoAnt"`
	VUtilPer    string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vUtilPer" json:"vUtilPer"`
	VSaldoFinal string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vSaldoFinal" json:"vSaldoFinal"`
	QtdOrigens  string                                                                      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 qtdOrigens" json:"qtdOrigens"`
	DetBCN      []DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCNGBCNCBSDetBCN `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBCN,omitempty" json:"detBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCN struct {
	GBCNIBS DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCNGBCNIBS `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCNIBS" json:"gBCNIBS"`
	GBCNCBS DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCNGBCNCBS `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCNCBS" json:"gBCNCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBC struct {
	CodBC          string                                                              `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 codBC" json:"codBC"`
	XDetBC         string                                                              `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 xDetBC" json:"xDetBC"`
	GCoeficientes  *DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCGCoeficientes `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gCoeficientes,omitempty" json:"gCoeficientes,omitempty"`
	MemoriaCalculo string                                                              `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 memoriaCalculo" json:"memoriaCalculo"`
	GBCIBS         DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCGBCIBS         `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCIBS" json:"gBCIBS"`
	GBCCBS         DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCGBCCBS         `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCCBS" json:"gBCCBS"`
	GBCIS          DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCGBCIS          `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 gBCIS" json:"gBCIS"`
	InfoBCN        *DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBCInfoBCN       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoBCN,omitempty" json:"infoBCN,omitempty"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProgTotalTributos struct {
	VIS     string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIS" json:"vIS"`
	VIBSMun string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSMun" json:"vIBSMun"`
	VIBSUF  string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSUF" json:"vIBSUF"`
	VIBSTot string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSTot" json:"vIBSTot"`
	VCBS    string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vCBS" json:"vCBS"`
}

type DeREContentEvtRetornoMensalInfoEventoInfoTotProg struct {
	DetBC         []DeREContentEvtRetornoMensalInfoEventoInfoTotProgDetBC       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 detBC,omitempty" json:"detBC,omitempty"`
	TotalTributos DeREContentEvtRetornoMensalInfoEventoInfoTotProgTotalTributos `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 totalTributos" json:"totalTributos"`
}

type DeREContentEvtRetornoMensalInfoEventoTotalTributosGeral struct {
	VIS     *string `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIS,omitempty" json:"vIS,omitempty"`
	VIBSMun string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSMun" json:"vIBSMun"`
	VIBSUF  string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSUF" json:"vIBSUF"`
	VIBSTot string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vIBSTot" json:"vIBSTot"`
	VCBS    string  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 vCBS" json:"vCBS"`
}

type DeREContentEvtRetornoMensalInfoEvento struct {
	IdePeriodo         DeREContentEvtRetornoMensalInfoEventoIdePeriodo         `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 idePeriodo" json:"idePeriodo"`
	InfoAdic           DeREContentEvtRetornoMensalInfoEventoInfoAdic           `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoAdic" json:"infoAdic"`
	InfoTotFinanceiro  *DeREContentEvtRetornoMensalInfoEventoInfoTotFinanceiro `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoTotFinanceiro,omitempty" json:"infoTotFinanceiro,omitempty"`
	InfoTotSaude       *DeREContentEvtRetornoMensalInfoEventoInfoTotSaude      `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoTotSaude,omitempty" json:"infoTotSaude,omitempty"`
	InfoTotProg        *DeREContentEvtRetornoMensalInfoEventoInfoTotProg       `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoTotProg,omitempty" json:"infoTotProg,omitempty"`
	TotalTributosGeral DeREContentEvtRetornoMensalInfoEventoTotalTributosGeral `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 totalTributosGeral" json:"totalTributosGeral"`
}

type DeREContentEvtRetornoMensal struct {
	ID         string                                 `xml:"id,attr" json:"id"`
	IdeContrib DeREContentEvtRetornoMensalIdeContrib  `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 ideContrib" json:"ideContrib"`
	IdeStatus  DeREContentEvtRetornoMensalIdeStatus   `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 ideStatus" json:"ideStatus"`
	InfoRecEv  DeREContentEvtRetornoMensalInfoRecEv   `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoRecEv" json:"infoRecEv"`
	InfoEvento *DeREContentEvtRetornoMensalInfoEvento `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 infoEvento,omitempty" json:"infoEvento,omitempty"`
}

type DeREContent struct {
	EvtRetornoMensal DeREContentEvtRetornoMensal `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 evtRetornoMensal" json:"evtRetornoMensal"`
	Signature        SignatureType               `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

// DeRE represents the global element {http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2}DeRE.
type DeRE struct {
	XMLName xml.Name `xml:"http://www.dere.gov.br/schemas/evtRetornoMensal/v0_0_2 DeRE" json:"-"`
	DeREContent
}

// CanonicalizationMethod represents the global element {http://www.w3.org/2000/09/xmldsig#}CanonicalizationMethod.
type CanonicalizationMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod" json:"-"`
	CanonicalizationMethodType
}

// DSAKeyValue represents the global element {http://www.w3.org/2000/09/xmldsig#}DSAKeyValue.
type DSAKeyValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# DSAKeyValue" json:"-"`
	DSAKeyValueType
}

// DigestMethod represents the global element {http://www.w3.org/2000/09/xmldsig#}DigestMethod.
type DigestMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod" json:"-"`
	DigestMethodType
}

// DigestValue represents the global element {http://www.w3.org/2000/09/xmldsig#}DigestValue.
type DigestValue struct {
	XMLName xml.Name        `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue" json:"-"`
	Value   DigestValueType `xml:",chardata" json:"value"`
}

// KeyInfo represents the global element {http://www.w3.org/2000/09/xmldsig#}KeyInfo.
type KeyInfo struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo" json:"-"`
	KeyInfoType
}

// KeyName represents the global element {http://www.w3.org/2000/09/xmldsig#}KeyName.
type KeyName struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# KeyName" json:"-"`
	Value   String   `xml:",chardata" json:"value"`
}

// KeyValue represents the global element {http://www.w3.org/2000/09/xmldsig#}KeyValue.
type KeyValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# KeyValue" json:"-"`
	KeyValueType
}

// Manifest represents the global element {http://www.w3.org/2000/09/xmldsig#}Manifest.
type Manifest struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Manifest" json:"-"`
	ManifestType
}

// MgmtData represents the global element {http://www.w3.org/2000/09/xmldsig#}MgmtData.
type MgmtData struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# MgmtData" json:"-"`
	Value   String   `xml:",chardata" json:"value"`
}

// Object represents the global element {http://www.w3.org/2000/09/xmldsig#}Object.
type Object struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Object" json:"-"`
	ObjectType
}

// PGPData represents the global element {http://www.w3.org/2000/09/xmldsig#}PGPData.
type PGPData struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# PGPData" json:"-"`
	PGPDataType
}

// RSAKeyValue represents the global element {http://www.w3.org/2000/09/xmldsig#}RSAKeyValue.
type RSAKeyValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# RSAKeyValue" json:"-"`
	RSAKeyValueType
}

// Reference represents the global element {http://www.w3.org/2000/09/xmldsig#}Reference.
type Reference struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Reference" json:"-"`
	ReferenceType
}

// RetrievalMethod represents the global element {http://www.w3.org/2000/09/xmldsig#}RetrievalMethod.
type RetrievalMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# RetrievalMethod" json:"-"`
	RetrievalMethodType
}

// SPKIData represents the global element {http://www.w3.org/2000/09/xmldsig#}SPKIData.
type SPKIData struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SPKIData" json:"-"`
	SPKIDataType
}

// Signature represents the global element {http://www.w3.org/2000/09/xmldsig#}Signature.
type Signature struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"-"`
	SignatureType
}

// SignatureMethod represents the global element {http://www.w3.org/2000/09/xmldsig#}SignatureMethod.
type SignatureMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod" json:"-"`
	SignatureMethodType
}

// SignatureProperties represents the global element {http://www.w3.org/2000/09/xmldsig#}SignatureProperties.
type SignatureProperties struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureProperties" json:"-"`
	SignaturePropertiesType
}

// SignatureProperty represents the global element {http://www.w3.org/2000/09/xmldsig#}SignatureProperty.
type SignatureProperty struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureProperty" json:"-"`
	SignaturePropertyType
}

// SignatureValue represents the global element {http://www.w3.org/2000/09/xmldsig#}SignatureValue.
type SignatureValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue" json:"-"`
	SignatureValueType
}

// SignedInfo represents the global element {http://www.w3.org/2000/09/xmldsig#}SignedInfo.
type SignedInfo struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo" json:"-"`
	SignedInfoType
}

// Transform represents the global element {http://www.w3.org/2000/09/xmldsig#}Transform.
type Transform struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Transform" json:"-"`
	TransformType
}

// Transforms represents the global element {http://www.w3.org/2000/09/xmldsig#}Transforms.
type Transforms struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Transforms" json:"-"`
	TransformsType
}

// X509Data represents the global element {http://www.w3.org/2000/09/xmldsig#}X509Data.
type X509Data struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# X509Data" json:"-"`
	X509DataType
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
