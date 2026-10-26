// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package evalteracaopagtoservmdfe_v3_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAmb string

type TCOrgaoIBGE string

type TChCTe string

type TChMDFe string

type TChNFe string

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

type TDec0303 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec0906Opc string

type TDec1104 string

type TDec1104Opc string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec1302 string

type TDec1302Opc string

type TEmit string

type TEventoInfEventoDetEvento struct {
	VersaoEvento TVerEvento   `xml:"versaoEvento,attr" json:"versaoEvento"`
	Any          []RawElement `xml:",any" json:"any,omitempty"`
}

type TEventoInfEventoInfSolicNFF struct {
	XSolic string `xml:"http://www.portalfiscal.inf.br/mdfe xSolic" json:"xSolic"`
}

type TEventoInfEventoInfPAAPAASignature struct {
	SignatureValue Base64Binary     `xml:"http://www.portalfiscal.inf.br/mdfe SignatureValue" json:"SignatureValue"`
	RSAKeyValue    TRSAKeyValueType `xml:"http://www.portalfiscal.inf.br/mdfe RSAKeyValue" json:"RSAKeyValue"`
}

type TEventoInfEventoInfPAA struct {
	CNPJPAA      TCNPJ                              `xml:"http://www.portalfiscal.inf.br/mdfe CNPJPAA" json:"CNPJPAA"`
	PAASignature TEventoInfEventoInfPAAPAASignature `xml:"http://www.portalfiscal.inf.br/mdfe PAASignature" json:"PAASignature"`
}

type TEventoInfEvento struct {
	ID          string                       `xml:"Id,attr" json:"Id"`
	COrgao      TCOrgaoIBGE                  `xml:"http://www.portalfiscal.inf.br/mdfe cOrgao" json:"cOrgao"`
	TpAmb       TAmb                         `xml:"http://www.portalfiscal.inf.br/mdfe tpAmb" json:"tpAmb"`
	CNPJ        *TCNPJ                       `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF         *TCPF                        `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	ChMDFe      TChMDFe                      `xml:"http://www.portalfiscal.inf.br/mdfe chMDFe" json:"chMDFe"`
	DhEvento    TDateTimeUTC                 `xml:"http://www.portalfiscal.inf.br/mdfe dhEvento" json:"dhEvento"`
	TpEvento    String                       `xml:"http://www.portalfiscal.inf.br/mdfe tpEvento" json:"tpEvento"`
	NSeqEvento  string                       `xml:"http://www.portalfiscal.inf.br/mdfe nSeqEvento" json:"nSeqEvento"`
	DetEvento   TEventoInfEventoDetEvento    `xml:"http://www.portalfiscal.inf.br/mdfe detEvento" json:"detEvento"`
	InfSolicNFF *TEventoInfEventoInfSolicNFF `xml:"http://www.portalfiscal.inf.br/mdfe infSolicNFF,omitempty" json:"infSolicNFF,omitempty"`
	InfPAA      *TEventoInfEventoInfPAA      `xml:"http://www.portalfiscal.inf.br/mdfe infPAA,omitempty" json:"infPAA,omitempty"`
}

type TEvento struct {
	Versao    TVerEvento       `xml:"versao,attr" json:"versao"`
	InfEvento TEventoInfEvento `xml:"http://www.portalfiscal.inf.br/mdfe infEvento" json:"infEvento"`
	Signature SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TIPv4 string

type TIe string

type TIeDest string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModMD string

type TMotivo string

type TNF string

type TNSU string

type TPlaca string

type TProcEvento struct {
	Versao        TVerEvento    `xml:"versao,attr" json:"versao"`
	IpTransmissor *TIPv4        `xml:"ipTransmissor,attr,omitempty" json:"ipTransmissor,omitempty"`
	NPortaCon     *string       `xml:"nPortaCon,attr,omitempty" json:"nPortaCon,omitempty"`
	DhConexao     *TDateTimeUTC `xml:"dhConexao,attr,omitempty" json:"dhConexao,omitempty"`
	EventoMDFe    TEvento       `xml:"http://www.portalfiscal.inf.br/mdfe eventoMDFe" json:"eventoMDFe"`
	RetEventoMDFe TRetEvento    `xml:"http://www.portalfiscal.inf.br/mdfe retEventoMDFe" json:"retEventoMDFe"`
}

type TProt string

type TRSAKeyValueType struct {
	Modulus  Base64Binary `xml:"http://www.portalfiscal.inf.br/mdfe Modulus" json:"Modulus"`
	Exponent Base64Binary `xml:"http://www.portalfiscal.inf.br/mdfe Exponent" json:"Exponent"`
}

type TRec string

type TRetEventoInfEvento struct {
	ID          *string       `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb       TAmb          `xml:"http://www.portalfiscal.inf.br/mdfe tpAmb" json:"tpAmb"`
	VerAplic    TVerAplic     `xml:"http://www.portalfiscal.inf.br/mdfe verAplic" json:"verAplic"`
	COrgao      TCOrgaoIBGE   `xml:"http://www.portalfiscal.inf.br/mdfe cOrgao" json:"cOrgao"`
	CStat       TStat         `xml:"http://www.portalfiscal.inf.br/mdfe cStat" json:"cStat"`
	XMotivo     TMotivo       `xml:"http://www.portalfiscal.inf.br/mdfe xMotivo" json:"xMotivo"`
	ChMDFe      *TChMDFe      `xml:"http://www.portalfiscal.inf.br/mdfe chMDFe,omitempty" json:"chMDFe,omitempty"`
	TpEvento    *string       `xml:"http://www.portalfiscal.inf.br/mdfe tpEvento,omitempty" json:"tpEvento,omitempty"`
	XEvento     *string       `xml:"http://www.portalfiscal.inf.br/mdfe xEvento,omitempty" json:"xEvento,omitempty"`
	NSeqEvento  *string       `xml:"http://www.portalfiscal.inf.br/mdfe nSeqEvento,omitempty" json:"nSeqEvento,omitempty"`
	DhRegEvento *TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/mdfe dhRegEvento,omitempty" json:"dhRegEvento,omitempty"`
	NProt       *TProt        `xml:"http://www.portalfiscal.inf.br/mdfe nProt,omitempty" json:"nProt,omitempty"`
}

type TRetEvento struct {
	Versao    TVerEvento          `xml:"versao,attr" json:"versao"`
	InfEvento TRetEventoInfEvento `xml:"http://www.portalfiscal.inf.br/mdfe infEvento" json:"infEvento"`
	Signature *SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TTransp string

type TUF string

type TVerAplic string

type TVerEvento string

type Tano string

type TtipoUnidCarga string

type TtipoUnidTransp string

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

type EvAlteracaoPagtoServMDFeContentInfPagComp struct {
	TpComp string   `xml:"http://www.portalfiscal.inf.br/mdfe tpComp" json:"tpComp"`
	VComp  TDec1302 `xml:"http://www.portalfiscal.inf.br/mdfe vComp" json:"vComp"`
	XComp  *string  `xml:"http://www.portalfiscal.inf.br/mdfe xComp,omitempty" json:"xComp,omitempty"`
}

type EvAlteracaoPagtoServMDFeContentInfPagInfPrazo struct {
	NParcela string      `xml:"http://www.portalfiscal.inf.br/mdfe nParcela" json:"nParcela"`
	DVenc    TData       `xml:"http://www.portalfiscal.inf.br/mdfe dVenc" json:"dVenc"`
	VParcela TDec1302Opc `xml:"http://www.portalfiscal.inf.br/mdfe vParcela" json:"vParcela"`
}

type EvAlteracaoPagtoServMDFeContentInfPagInfBanc struct {
	CodBanco   *string   `xml:"http://www.portalfiscal.inf.br/mdfe codBanco,omitempty" json:"codBanco,omitempty"`
	CodAgencia *string   `xml:"http://www.portalfiscal.inf.br/mdfe codAgencia,omitempty" json:"codAgencia,omitempty"`
	CNPJIPEF   *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/mdfe CNPJIPEF,omitempty" json:"CNPJIPEF,omitempty"`
	PIX        *string   `xml:"http://www.portalfiscal.inf.br/mdfe PIX,omitempty" json:"PIX,omitempty"`
}

type EvAlteracaoPagtoServMDFeContentInfPag struct {
	XNome             *string                                         `xml:"http://www.portalfiscal.inf.br/mdfe xNome,omitempty" json:"xNome,omitempty"`
	CPF               *TCPF                                           `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	CNPJ              *TCNPJOpc                                       `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	IDEstrangeiro     *string                                         `xml:"http://www.portalfiscal.inf.br/mdfe idEstrangeiro,omitempty" json:"idEstrangeiro,omitempty"`
	Comp              []EvAlteracaoPagtoServMDFeContentInfPagComp     `xml:"http://www.portalfiscal.inf.br/mdfe Comp,omitempty" json:"Comp,omitempty"`
	VContrato         TDec1302                                        `xml:"http://www.portalfiscal.inf.br/mdfe vContrato" json:"vContrato"`
	IndPag            string                                          `xml:"http://www.portalfiscal.inf.br/mdfe indPag" json:"indPag"`
	VAdiant           *TDec1302                                       `xml:"http://www.portalfiscal.inf.br/mdfe vAdiant,omitempty" json:"vAdiant,omitempty"`
	IndAntecipaAdiant *string                                         `xml:"http://www.portalfiscal.inf.br/mdfe indAntecipaAdiant,omitempty" json:"indAntecipaAdiant,omitempty"`
	InfPrazo          []EvAlteracaoPagtoServMDFeContentInfPagInfPrazo `xml:"http://www.portalfiscal.inf.br/mdfe infPrazo,omitempty" json:"infPrazo,omitempty"`
	TpAntecip         *string                                         `xml:"http://www.portalfiscal.inf.br/mdfe tpAntecip,omitempty" json:"tpAntecip,omitempty"`
	InfBanc           EvAlteracaoPagtoServMDFeContentInfPagInfBanc    `xml:"http://www.portalfiscal.inf.br/mdfe infBanc" json:"infBanc"`
}

type EvAlteracaoPagtoServMDFeContent struct {
	DescEvento string                                  `xml:"http://www.portalfiscal.inf.br/mdfe descEvento" json:"descEvento"`
	NProt      TProt                                   `xml:"http://www.portalfiscal.inf.br/mdfe nProt" json:"nProt"`
	InfPag     []EvAlteracaoPagtoServMDFeContentInfPag `xml:"http://www.portalfiscal.inf.br/mdfe infPag,omitempty" json:"infPag,omitempty"`
}

// EvAlteracaoPagtoServMDFe represents the global element {http://www.portalfiscal.inf.br/mdfe}evAlteracaoPagtoServMDFe.
type EvAlteracaoPagtoServMDFe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/mdfe evAlteracaoPagtoServMDFe" json:"-"`
	EvAlteracaoPagtoServMDFeContent
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
