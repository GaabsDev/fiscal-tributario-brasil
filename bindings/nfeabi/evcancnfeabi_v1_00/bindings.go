// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package evcancnfeabi_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAdquirApropCredInd struct {
	NAdquir   TSeq1a99 `xml:"nAdquir,attr" json:"nAdquir"`
	CNPJ      *TCNPJ   `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF    `xml:"http://www.portalfiscal.inf.br/nfeabi CPF,omitempty" json:"CPF,omitempty"`
	XNome     TNome    `xml:"http://www.portalfiscal.inf.br/nfeabi xNome" json:"xNome"`
	PParticip TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pParticip" json:"pParticip"`
}

type TAmb string

type TCOrgaoIBGE string

type TCSTIBSCBS string

type TChDFe string

type TChNFeABI string

type TCNPJ string

type TCNPJOpc string

type TCNPJVar string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

type TCPF string

type TCPFVar string

type TData string

type TDate string

type TDateTimeUTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec030204 string

type TDec0302a04 string

type TDec0302a04Max100 string

type TDec0302a04Opc string

type TDec0303 string

type TDec0303Opc string

type TDec0304 string

type TDec04020408 string

type TDec0503 string

type TDec0702 string

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

type TEmail string

type TEndereco struct {
	XLgr    string  `xml:"http://www.portalfiscal.inf.br/nfeabi xLgr" json:"xLgr"`
	Nro     string  `xml:"http://www.portalfiscal.inf.br/nfeabi nro" json:"nro"`
	XCpl    *string `xml:"http://www.portalfiscal.inf.br/nfeabi xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro *string `xml:"http://www.portalfiscal.inf.br/nfeabi xBairro,omitempty" json:"xBairro,omitempty"`
	CMun    string  `xml:"http://www.portalfiscal.inf.br/nfeabi cMun" json:"cMun"`
	XMun    string  `xml:"http://www.portalfiscal.inf.br/nfeabi xMun" json:"xMun"`
	UF      string  `xml:"http://www.portalfiscal.inf.br/nfeabi UF" json:"UF"`
	CEP     *string `xml:"http://www.portalfiscal.inf.br/nfeabi CEP,omitempty" json:"CEP,omitempty"`
	CPais   *string `xml:"http://www.portalfiscal.inf.br/nfeabi cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string `xml:"http://www.portalfiscal.inf.br/nfeabi xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string `xml:"http://www.portalfiscal.inf.br/nfeabi fone,omitempty" json:"fone,omitempty"`
}

type TEventoInfEventoDetEvento struct {
	VersaoEvento TVerEvento   `xml:"versaoEvento,attr" json:"versaoEvento"`
	Any          []RawElement `xml:",any" json:"any,omitempty"`
}

type TEventoInfEvento struct {
	ID         string                    `xml:"Id,attr" json:"Id"`
	COrgao     string                    `xml:"http://www.portalfiscal.inf.br/nfeabi cOrgao" json:"cOrgao"`
	TpAmb      string                    `xml:"http://www.portalfiscal.inf.br/nfeabi tpAmb" json:"tpAmb"`
	CNPJ       *TCNPJ                    `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF        *TCPF                     `xml:"http://www.portalfiscal.inf.br/nfeabi CPF,omitempty" json:"CPF,omitempty"`
	ChNFeABI   string                    `xml:"http://www.portalfiscal.inf.br/nfeabi chNFeABI" json:"chNFeABI"`
	DhEvento   string                    `xml:"http://www.portalfiscal.inf.br/nfeabi dhEvento" json:"dhEvento"`
	TpEvento   string                    `xml:"http://www.portalfiscal.inf.br/nfeabi tpEvento" json:"tpEvento"`
	NSeqEvento string                    `xml:"http://www.portalfiscal.inf.br/nfeabi nSeqEvento" json:"nSeqEvento"`
	VerEvento  TVerEvento                `xml:"http://www.portalfiscal.inf.br/nfeabi verEvento" json:"verEvento"`
	DetEvento  TEventoInfEventoDetEvento `xml:"http://www.portalfiscal.inf.br/nfeabi detEvento" json:"detEvento"`
}

type TEvento struct {
	Versao    TVerEvento       `xml:"versao,attr" json:"versao"`
	InfEvento TEventoInfEvento `xml:"http://www.portalfiscal.inf.br/nfeabi infEvento" json:"infEvento"`
	Signature SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TGCBS struct {
	PCBS TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pCBS" json:"pCBS"`
	GRed TGRed    `xml:"http://www.portalfiscal.inf.br/nfeabi gRed" json:"gRed"`
	VCBS TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vCBS" json:"vCBS"`
}

type TGCompraGov struct {
	TpEnteGov string     `xml:"http://www.portalfiscal.inf.br/nfeabi tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204 `xml:"http://www.portalfiscal.inf.br/nfeabi pRedutor" json:"pRedutor"`
	TpOperGov string     `xml:"http://www.portalfiscal.inf.br/nfeabi tpOperGov" json:"tpOperGov"`
}

type TGEstornoCred struct {
	VIBSEstCred TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vCBSEstCred" json:"vCBSEstCred"`
}

type TGIBSCBS struct {
	VOperacIndiv    TDec1302         `xml:"http://www.portalfiscal.inf.br/nfeabi vOperacIndiv" json:"vOperacIndiv"`
	VTornaIndiv     *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vTornaIndiv,omitempty" json:"vTornaIndiv,omitempty"`
	VRedAjusteIndiv *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vRedAjusteIndiv,omitempty" json:"vRedAjusteIndiv,omitempty"`
	VRedSocialIndiv *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vRedSocialIndiv,omitempty" json:"vRedSocialIndiv,omitempty"`
	VBC             TDec1302         `xml:"http://www.portalfiscal.inf.br/nfeabi vBC" json:"vBC"`
	GIBSUF          TGIBSUF          `xml:"http://www.portalfiscal.inf.br/nfeabi gIBSUF" json:"gIBSUF"`
	GIBSMun         TGIBSMun         `xml:"http://www.portalfiscal.inf.br/nfeabi gIBSMun" json:"gIBSMun"`
	GCBS            TGCBS            `xml:"http://www.portalfiscal.inf.br/nfeabi gCBS" json:"gCBS"`
	GTribCompraGov  *TGTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfeabi gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TGIBSCBSEventoPagParcela struct {
	VOperacIndiv    *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vOperacIndiv,omitempty" json:"vOperacIndiv,omitempty"`
	VRedAjusteIndiv *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vRedAjusteIndiv,omitempty" json:"vRedAjusteIndiv,omitempty"`
	VRedSocialIndiv *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vRedSocialIndiv,omitempty" json:"vRedSocialIndiv,omitempty"`
	VBC             TDec1302         `xml:"http://www.portalfiscal.inf.br/nfeabi vBC" json:"vBC"`
	GIBSUF          TGIBSUF          `xml:"http://www.portalfiscal.inf.br/nfeabi gIBSUF" json:"gIBSUF"`
	GIBSMun         TGIBSMun         `xml:"http://www.portalfiscal.inf.br/nfeabi gIBSMun" json:"gIBSMun"`
	GCBS            TGCBS            `xml:"http://www.portalfiscal.inf.br/nfeabi gCBS" json:"gCBS"`
	GTribCompraGov  *TGTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfeabi gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TGIBSMun struct {
	PIBSMun TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pIBSMun" json:"pIBSMun"`
	GRed    TGRed    `xml:"http://www.portalfiscal.inf.br/nfeabi gRed" json:"gRed"`
	VIBSMun TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vIBSMun" json:"vIBSMun"`
}

type TGIBSUF struct {
	PIBSUF TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pIBSUF" json:"pIBSUF"`
	GRed   TGRed    `xml:"http://www.portalfiscal.inf.br/nfeabi gRed" json:"gRed"`
	VIBSUF TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vIBSUF" json:"vIBSUF"`
}

type TGIntermedCorret struct {
	NCorretagem string   `xml:"nCorretagem,attr" json:"nCorretagem"`
	CNPJ        *TCNPJ   `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF         *TCPF    `xml:"http://www.portalfiscal.inf.br/nfeabi CPF,omitempty" json:"CPF,omitempty"`
	XNome       string   `xml:"http://www.portalfiscal.inf.br/nfeabi xNome" json:"xNome"`
	VCorretagem TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vCorretagem" json:"vCorretagem"`
}

type TGRed struct {
	PRedAliq  TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAliqEfet" json:"pAliqEfet"`
}

type TGTribCompraGov struct {
	PAliqIBSUF  TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vTribCBS" json:"vTribCBS"`
}

type TIBSCBS struct {
	CST          string         `xml:"http://www.portalfiscal.inf.br/nfeabi CST" json:"CST"`
	CClassTrib   string         `xml:"http://www.portalfiscal.inf.br/nfeabi cClassTrib" json:"cClassTrib"`
	GIBSCBS      *TGIBSCBS      `xml:"http://www.portalfiscal.inf.br/nfeabi gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfeabi gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIPv4 string

type TIe string

type TIeDest string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModNFeABI string

type TMotivo string

type TNF string

type TNSU string

type TNome string

type TProcEvento struct {
	Versao          TVerEvento `xml:"versao,attr" json:"versao"`
	IpTransmissor   *string    `xml:"ipTransmissor,attr,omitempty" json:"ipTransmissor,omitempty"`
	NPortaCon       *string    `xml:"nPortaCon,attr,omitempty" json:"nPortaCon,omitempty"`
	DhConexao       *string    `xml:"dhConexao,attr,omitempty" json:"dhConexao,omitempty"`
	EventoNFeABI    TEvento    `xml:"http://www.portalfiscal.inf.br/nfeabi eventoNFeABI" json:"eventoNFeABI"`
	RetEventoNFeABI TRetEvento `xml:"http://www.portalfiscal.inf.br/nfeabi retEventoNFeABI" json:"retEventoNFeABI"`
}

type TProt string

type TProtNFeABIInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/nfeabi tpAmb" json:"tpAmb"`
	VerAplic string           `xml:"http://www.portalfiscal.inf.br/nfeabi verAplic" json:"verAplic"`
	ChNFeABI TChNFeABI        `xml:"http://www.portalfiscal.inf.br/nfeabi chNFeABI" json:"chNFeABI"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/nfeabi dhRecbto" json:"dhRecbto"`
	NProt    *string          `xml:"http://www.portalfiscal.inf.br/nfeabi nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/nfeabi digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/nfeabi cStat" json:"cStat"`
	XMotivo  string           `xml:"http://www.portalfiscal.inf.br/nfeabi xMotivo" json:"xMotivo"`
}

type TProtNFeABI struct {
	Versao    TVerNFeABI         `xml:"versao,attr" json:"versao"`
	InfProt   TProtNFeABIInfProt `xml:"http://www.portalfiscal.inf.br/nfeabi infProt" json:"infProt"`
	Signature *SignatureType     `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRec string

type TRetEventoInfEvento struct {
	ID          *string `xml:"http://www.portalfiscal.inf.br/nfeabi Id,omitempty" json:"Id,omitempty"`
	TpAmb       string  `xml:"http://www.portalfiscal.inf.br/nfeabi tpAmb" json:"tpAmb"`
	VerAplic    string  `xml:"http://www.portalfiscal.inf.br/nfeabi verAplic" json:"verAplic"`
	COrgao      string  `xml:"http://www.portalfiscal.inf.br/nfeabi cOrgao" json:"cOrgao"`
	CStat       string  `xml:"http://www.portalfiscal.inf.br/nfeabi cStat" json:"cStat"`
	XMotivo     string  `xml:"http://www.portalfiscal.inf.br/nfeabi xMotivo" json:"xMotivo"`
	ChNFeABI    *string `xml:"http://www.portalfiscal.inf.br/nfeabi chNFeABI,omitempty" json:"chNFeABI,omitempty"`
	TpEvento    *string `xml:"http://www.portalfiscal.inf.br/nfeabi tpEvento,omitempty" json:"tpEvento,omitempty"`
	XEvento     *string `xml:"http://www.portalfiscal.inf.br/nfeabi xEvento,omitempty" json:"xEvento,omitempty"`
	NSeqEvento  *string `xml:"http://www.portalfiscal.inf.br/nfeabi nSeqEvento,omitempty" json:"nSeqEvento,omitempty"`
	DhRegEvento *string `xml:"http://www.portalfiscal.inf.br/nfeabi dhRegEvento,omitempty" json:"dhRegEvento,omitempty"`
	NProt       *string `xml:"http://www.portalfiscal.inf.br/nfeabi nProt,omitempty" json:"nProt,omitempty"`
}

type TRetEvento struct {
	Versao    TVerEvento          `xml:"versao,attr" json:"versao"`
	InfEvento TRetEventoInfEvento `xml:"http://www.portalfiscal.inf.br/nfeabi infEvento" json:"infEvento"`
	Signature *SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRetNFeABI struct {
	Versao     TVerNFeABI   `xml:"versao,attr" json:"versao"`
	TpAmb      TAmb         `xml:"http://www.portalfiscal.inf.br/nfeabi tpAmb" json:"tpAmb"`
	VerAplic   string       `xml:"http://www.portalfiscal.inf.br/nfeabi verAplic" json:"verAplic"`
	CStat      string       `xml:"http://www.portalfiscal.inf.br/nfeabi cStat" json:"cStat"`
	XMotivo    string       `xml:"http://www.portalfiscal.inf.br/nfeabi xMotivo" json:"xMotivo"`
	CUF        TCodUFIBGE   `xml:"http://www.portalfiscal.inf.br/nfeabi cUF" json:"cUF"`
	DhRecbto   TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfeabi dhRecbto" json:"dhRecbto"`
	ProtNFeABI *TProtNFeABI `xml:"http://www.portalfiscal.inf.br/nfeabi protNFeABI,omitempty" json:"protNFeABI,omitempty"`
}

type TSeq1a99 string

type TSerie string

type TServ string

type TStat string

type TString string

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerEvento string

type TVerNFeABI string

type Tano string

type TcClassTribIBSCBS string

type TtpEventoPagParcela string

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

type EvCancNFeABIContent struct {
	Versao     string `xml:"versao,attr" json:"versao"`
	DescEvento string `xml:"http://www.portalfiscal.inf.br/nfeabi descEvento" json:"descEvento"`
	NProt      TProt  `xml:"http://www.portalfiscal.inf.br/nfeabi nProt" json:"nProt"`
	XJust      TJust  `xml:"http://www.portalfiscal.inf.br/nfeabi xJust" json:"xJust"`
}

// EvCancNFeABI represents the global element {http://www.portalfiscal.inf.br/nfeabi}evCancNFeABI.
type EvCancNFeABI struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfeabi evCancNFeABI" json:"-"`
	EvCancNFeABIContent
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
