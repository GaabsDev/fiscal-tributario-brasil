// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package evtoperfinanc_v0_0_2

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

type DeREContentEvtOperFinancIdeEvento struct {
	FinEvt   string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 finEvt" json:"finEvt"`
	TpAmb    string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 tpAmb" json:"tpAmb"`
	AplicEmi string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 aplicEmi" json:"aplicEmi"`
	VerAplic string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 verAplic" json:"verAplic"`
}

type DeREContentEvtOperFinancIdeContrib struct {
	NrInsc string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 nrInsc" json:"nrInsc"`
}

type DeREContentEvtOperFinancIdePeriodo struct {
	PerApur string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 perApur" json:"perApur"`
	SubPer  string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 subPer" json:"subPer"`
}

type DeREContentEvtOperFinancInfoOperGOperIdePartesInfoAdqAdqExterior struct {
	NIF        *string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 NIF,omitempty" json:"NIF,omitempty"`
	CNaoNIF    *string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 cNaoNIF,omitempty" json:"cNaoNIF,omitempty"`
	NrDocIdent *string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 nrDocIdent,omitempty" json:"nrDocIdent,omitempty"`
	CPais      string  `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 cPais" json:"cPais"`
}

type DeREContentEvtOperFinancInfoOperGOperIdePartesInfoAdq struct {
	CPF         *string                                                           `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 CPF,omitempty" json:"CPF,omitempty"`
	CNPJ        *string                                                           `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 CNPJ,omitempty" json:"CNPJ,omitempty"`
	CMun        *string                                                           `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 cMun,omitempty" json:"cMun,omitempty"`
	AdqExterior *DeREContentEvtOperFinancInfoOperGOperIdePartesInfoAdqAdqExterior `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 adqExterior,omitempty" json:"adqExterior,omitempty"`
}

type DeREContentEvtOperFinancInfoOperGOperIdePartes struct {
	InfoAdq DeREContentEvtOperFinancInfoOperGOperIdePartesInfoAdq `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 infoAdq" json:"infoAdq"`
}

type DeREContentEvtOperFinancInfoOperGOperDadosOperDetOperQualifOper struct {
	IndMoedaEstr string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 indMoedaEstr" json:"indMoedaEstr"`
}

type DeREContentEvtOperFinancInfoOperGOperDadosOperDetOperGBC struct {
	VBCApur string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 vBCApur" json:"vBCApur"`
}

type DeREContentEvtOperFinancInfoOperGOperDadosOperDetOperGTributos struct {
	VBCTrib     string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 vBCTrib" json:"vBCTrib"`
	PIBSMunTrib string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 pIBSMunTrib" json:"pIBSMunTrib"`
	VIBSMunTrib string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 vIBSMunTrib" json:"vIBSMunTrib"`
	PIBSUFTrib  string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 pIBSUFTrib" json:"pIBSUFTrib"`
	VIBSUFTrib  string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 vIBSUFTrib" json:"vIBSUFTrib"`
	PCBSTrib    string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 pCBSTrib" json:"pCBSTrib"`
	VCBSTrib    string `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 vCBSTrib" json:"vCBSTrib"`
}

type DeREContentEvtOperFinancInfoOperGOperDadosOperDetOper struct {
	Seq        string                                                          `xml:"seq,attr" json:"seq"`
	SeqRef     *string                                                         `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 seqRef,omitempty" json:"seqRef,omitempty"`
	DhOper     string                                                          `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 dhOper" json:"dhOper"`
	QualifOper DeREContentEvtOperFinancInfoOperGOperDadosOperDetOperQualifOper `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 qualifOper" json:"qualifOper"`
	GBC        DeREContentEvtOperFinancInfoOperGOperDadosOperDetOperGBC        `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 gBC" json:"gBC"`
	GTributos  DeREContentEvtOperFinancInfoOperGOperDadosOperDetOperGTributos  `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 gTributos" json:"gTributos"`
}

type DeREContentEvtOperFinancInfoOperGOperDadosOper struct {
	DetOper []DeREContentEvtOperFinancInfoOperGOperDadosOperDetOper `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 detOper,omitempty" json:"detOper,omitempty"`
}

type DeREContentEvtOperFinancInfoOperGOper struct {
	ChAgrup    string                                         `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 chAgrup" json:"chAgrup"`
	CodBC      string                                         `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 codBC" json:"codBC"`
	ChAgrupRef *string                                        `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 chAgrupRef,omitempty" json:"chAgrupRef,omitempty"`
	IdePartes  DeREContentEvtOperFinancInfoOperGOperIdePartes `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 idePartes" json:"idePartes"`
	DadosOper  DeREContentEvtOperFinancInfoOperGOperDadosOper `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 dadosOper" json:"dadosOper"`
}

type DeREContentEvtOperFinancInfoOper struct {
	GOper DeREContentEvtOperFinancInfoOperGOper `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 gOper" json:"gOper"`
}

type DeREContentEvtOperFinanc struct {
	ID         string                             `xml:"id,attr" json:"id"`
	IdeEvento  DeREContentEvtOperFinancIdeEvento  `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 ideEvento" json:"ideEvento"`
	IdeContrib DeREContentEvtOperFinancIdeContrib `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 ideContrib" json:"ideContrib"`
	IdePeriodo DeREContentEvtOperFinancIdePeriodo `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 idePeriodo" json:"idePeriodo"`
	InfoOper   DeREContentEvtOperFinancInfoOper   `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 infoOper" json:"infoOper"`
}

type DeREContent struct {
	EvtOperFinanc DeREContentEvtOperFinanc `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 evtOperFinanc" json:"evtOperFinanc"`
	Signature     SignatureType            `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

// DeRE represents the global element {http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2}DeRE.
type DeRE struct {
	XMLName xml.Name `xml:"http://www.dere.gov.br/schemas/evtOperFinanc/v0_0_2 DeRE" json:"-"`
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
