// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package retornolotedere_v1_0_1

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TDadosProcessamento struct {
	DhProcessamento               DateTime `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 dhProcessamento" json:"dhProcessamento"`
	VersaoAplicativoProcessamento *string  `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 versaoAplicativoProcessamento,omitempty" json:"versaoAplicativoProcessamento,omitempty"`
}

type TDadosRecepcao struct {
	DhRecepcao               DateTime `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 dhRecepcao" json:"dhRecepcao"`
	VersaoAplicativoRecepcao string   `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 versaoAplicativoRecepcao" json:"versaoAplicativoRecepcao"`
	Protocolo                string   `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 protocolo" json:"protocolo"`
}

type TEventoDere struct {
	ID  ID           `xml:"id,attr" json:"id"`
	Any []RawElement `xml:",any" json:"any,omitempty"`
}

type TOcorrenciasOcorrencia struct {
	Codigo      string       `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 codigo" json:"codigo"`
	Descricao   string       `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 descricao" json:"descricao"`
	Tipo        UnsignedByte `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 tipo" json:"tipo"`
	Localizacao *string      `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 localizacao,omitempty" json:"localizacao,omitempty"`
}

type TOcorrencias struct {
	Ocorrencia []TOcorrenciasOcorrencia `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 ocorrencia,omitempty" json:"ocorrencia,omitempty"`
}

type TStatus struct {
	CdResposta   Int           `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 cdResposta" json:"cdResposta"`
	DescResposta string        `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 descResposta" json:"descResposta"`
	Ocorrencias  *TOcorrencias `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 ocorrencias,omitempty" json:"ocorrencias,omitempty"`
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

type DateTime string

type Int string

type Integer string

type String string

type UnsignedByte string

type DeREContentRetornoLoteEventosIdeContrib struct {
	NrInsc string `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 nrInsc" json:"nrInsc"`
}

type DeREContentRetornoLoteEventosIdeTransmissor struct {
	NiTransmissor string `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 niTransmissor" json:"niTransmissor"`
}

type DeREContentRetornoLoteEventosRetornoEventos struct {
	Evento []TEventoDere `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 evento,omitempty" json:"evento,omitempty"`
}

type DeREContentRetornoLoteEventos struct {
	ID                     ID                                           `xml:"id,attr" json:"id"`
	IdeContrib             DeREContentRetornoLoteEventosIdeContrib      `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 ideContrib" json:"ideContrib"`
	IdeTransmissor         DeREContentRetornoLoteEventosIdeTransmissor  `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 ideTransmissor" json:"ideTransmissor"`
	Status                 TStatus                                      `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 status" json:"status"`
	DadosRecepcaoLote      *TDadosRecepcao                              `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 dadosRecepcaoLote,omitempty" json:"dadosRecepcaoLote,omitempty"`
	DadosProcessamentoLote *TDadosProcessamento                         `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 dadosProcessamentoLote,omitempty" json:"dadosProcessamentoLote,omitempty"`
	RetornoEventos         *DeREContentRetornoLoteEventosRetornoEventos `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 retornoEventos,omitempty" json:"retornoEventos,omitempty"`
}

type DeREContent struct {
	RetornoLoteEventos DeREContentRetornoLoteEventos `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 retornoLoteEventos" json:"retornoLoteEventos"`
	Signature          SignatureType                 `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

// DeRE represents the global element {http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1}DeRE.
type DeRE struct {
	XMLName xml.Name `xml:"http://www.dere.gov.br/schemas/retornoLoteDere/v1_0_1 DeRE" json:"-"`
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
