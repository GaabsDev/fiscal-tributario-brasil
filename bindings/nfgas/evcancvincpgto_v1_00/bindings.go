// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package evcancvincpgto_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAmb string

type TCOrgaoIBGE string

type TChDFe string

type TCNPJ string

type TCNPJOpc string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

type TCPF string

type TCPFVar string

type TData string

type TDateTimeUTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec0302a04 string

type TDec0302a04Max100 string

type TDec0302a04Opc string

type TDec0303 string

type TDec0303Opc string

type TDec04020408 string

type TDec0503 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec0906 string

type TDec0906Opc string

type TDec0908 string

type TDec0908Opc string

type TDec11001104 string

type TDec1104 string

type TDec1104Opc string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec1302 string

type TDec1302Opc string

type TDec13021304 string

type TDec13021306 string

type TDec13021308 string

type TDec13021310 string

type TEventoInfEventoDetEvento struct {
	VersaoEvento string       `xml:"versaoEvento,attr" json:"versaoEvento"`
	Any          []RawElement `xml:",any" json:"any,omitempty"`
}

type TEventoInfEvento struct {
	ID         string                    `xml:"Id,attr" json:"Id"`
	COrgao     TCOrgaoIBGE               `xml:"http://www.portalfiscal.inf.br/nfgas cOrgao" json:"cOrgao"`
	TpAmb      TAmb                      `xml:"http://www.portalfiscal.inf.br/nfgas tpAmb" json:"tpAmb"`
	CNPJ       TCNPJ                     `xml:"http://www.portalfiscal.inf.br/nfgas CNPJ" json:"CNPJ"`
	ChNFGas    TChDFe                    `xml:"http://www.portalfiscal.inf.br/nfgas chNFGas" json:"chNFGas"`
	DhEvento   TDateTimeUTC              `xml:"http://www.portalfiscal.inf.br/nfgas dhEvento" json:"dhEvento"`
	TpEvento   string                    `xml:"http://www.portalfiscal.inf.br/nfgas tpEvento" json:"tpEvento"`
	NSeqEvento string                    `xml:"http://www.portalfiscal.inf.br/nfgas nSeqEvento" json:"nSeqEvento"`
	DetEvento  TEventoInfEventoDetEvento `xml:"http://www.portalfiscal.inf.br/nfgas detEvento" json:"detEvento"`
}

type TEvento struct {
	Versao    TVerEvento       `xml:"versao,attr" json:"versao"`
	InfEvento TEventoInfEvento `xml:"http://www.portalfiscal.inf.br/nfgas infEvento" json:"infEvento"`
	Signature SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TIPv4 string

type TIe string

type TIeDest string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModNF string

type TMotivo string

type TNF string

type TNSU string

type TProcEvento struct {
	Versao         TVerEvento    `xml:"versao,attr" json:"versao"`
	IpTransmissor  *TIPv4        `xml:"ipTransmissor,attr,omitempty" json:"ipTransmissor,omitempty"`
	NPortaCon      *string       `xml:"nPortaCon,attr,omitempty" json:"nPortaCon,omitempty"`
	DhConexao      *TDateTimeUTC `xml:"dhConexao,attr,omitempty" json:"dhConexao,omitempty"`
	EventoNFGas    TEvento       `xml:"http://www.portalfiscal.inf.br/nfgas eventoNFGas" json:"eventoNFGas"`
	RetEventoNFGas TRetEvento    `xml:"http://www.portalfiscal.inf.br/nfgas retEventoNFGas" json:"retEventoNFGas"`
}

type TProt string

type TRec string

type TRetEventoInfEvento struct {
	ID          *string       `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb       TAmb          `xml:"http://www.portalfiscal.inf.br/nfgas tpAmb" json:"tpAmb"`
	VerAplic    TVerAplic     `xml:"http://www.portalfiscal.inf.br/nfgas verAplic" json:"verAplic"`
	COrgao      TCOrgaoIBGE   `xml:"http://www.portalfiscal.inf.br/nfgas cOrgao" json:"cOrgao"`
	CStat       TStat         `xml:"http://www.portalfiscal.inf.br/nfgas cStat" json:"cStat"`
	XMotivo     TMotivo       `xml:"http://www.portalfiscal.inf.br/nfgas xMotivo" json:"xMotivo"`
	ChNFGas     *TChDFe       `xml:"http://www.portalfiscal.inf.br/nfgas chNFGas,omitempty" json:"chNFGas,omitempty"`
	TpEvento    *string       `xml:"http://www.portalfiscal.inf.br/nfgas tpEvento,omitempty" json:"tpEvento,omitempty"`
	XEvento     *string       `xml:"http://www.portalfiscal.inf.br/nfgas xEvento,omitempty" json:"xEvento,omitempty"`
	NSeqEvento  *string       `xml:"http://www.portalfiscal.inf.br/nfgas nSeqEvento,omitempty" json:"nSeqEvento,omitempty"`
	DhRegEvento *TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfgas dhRegEvento,omitempty" json:"dhRegEvento,omitempty"`
	NProt       *TProt        `xml:"http://www.portalfiscal.inf.br/nfgas nProt,omitempty" json:"nProt,omitempty"`
}

type TRetEvento struct {
	Versao    string              `xml:"versao,attr" json:"versao"`
	InfEvento TRetEventoInfEvento `xml:"http://www.portalfiscal.inf.br/nfgas infEvento" json:"infEvento"`
	Signature *SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerEvento string

type Tano string

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

type EvCancVincPgtoContent struct {
	DescEvento    string `xml:"http://www.portalfiscal.inf.br/nfgas descEvento" json:"descEvento"`
	NProt         TProt  `xml:"http://www.portalfiscal.inf.br/nfgas nProt" json:"nProt"`
	NProtVincPgto TProt  `xml:"http://www.portalfiscal.inf.br/nfgas nProtVincPgto" json:"nProtVincPgto"`
}

// EvCancVincPgto represents the global element {http://www.portalfiscal.inf.br/nfgas}evCancVincPgto.
type EvCancVincPgto struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfgas evCancVincPgto" json:"-"`
	EvCancVincPgtoContent
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
