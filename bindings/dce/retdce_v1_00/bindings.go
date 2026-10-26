// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package retdce_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAmb string

type TCOrgaoIBGE string

type TCRT string

type TChDFe string

type TCNPJ string

type TCNPJOpc string

type TCNPJVar string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

type TCompet string

type TCPF string

type TCPFVar string

type TDCeInfDCeIde struct {
	CUF          TCodUFIBGE   `xml:"http://www.portalfiscal.inf.br/dce cUF" json:"cUF"`
	CDC          string       `xml:"http://www.portalfiscal.inf.br/dce cDC" json:"cDC"`
	Mod          TModDCe      `xml:"http://www.portalfiscal.inf.br/dce mod" json:"mod"`
	Serie        TSerie       `xml:"http://www.portalfiscal.inf.br/dce serie" json:"serie"`
	NDC          TNF          `xml:"http://www.portalfiscal.inf.br/dce nDC" json:"nDC"`
	DhEmi        TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/dce dhEmi" json:"dhEmi"`
	TpEmis       TtpEmis      `xml:"http://www.portalfiscal.inf.br/dce tpEmis" json:"tpEmis"`
	TpEmit       string       `xml:"http://www.portalfiscal.inf.br/dce tpEmit" json:"tpEmit"`
	NSiteAutoriz string       `xml:"http://www.portalfiscal.inf.br/dce nSiteAutoriz" json:"nSiteAutoriz"`
	CDV          string       `xml:"http://www.portalfiscal.inf.br/dce cDV" json:"cDV"`
	TpAmb        TAmb         `xml:"http://www.portalfiscal.inf.br/dce tpAmb" json:"tpAmb"`
	VerProc      string       `xml:"http://www.portalfiscal.inf.br/dce verProc" json:"verProc"`
}

type TDCeInfDCeEmit struct {
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/dce CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/dce CPF,omitempty" json:"CPF,omitempty"`
	IDOutros  *string   `xml:"http://www.portalfiscal.inf.br/dce idOutros,omitempty" json:"idOutros,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/dce xNome" json:"xNome"`
	EnderEmit TEndeEmi  `xml:"http://www.portalfiscal.inf.br/dce enderEmit" json:"enderEmit"`
}

type TDCeInfDCeFisco struct {
	CNPJ   TCNPJOpc `xml:"http://www.portalfiscal.inf.br/dce CNPJ" json:"CNPJ"`
	XOrgao string   `xml:"http://www.portalfiscal.inf.br/dce xOrgao" json:"xOrgao"`
	UF     TUF      `xml:"http://www.portalfiscal.inf.br/dce UF" json:"UF"`
}

type TDCeInfDCeMarketplace struct {
	CNPJ  TCNPJOpc `xml:"http://www.portalfiscal.inf.br/dce CNPJ" json:"CNPJ"`
	XNome string   `xml:"http://www.portalfiscal.inf.br/dce xNome" json:"xNome"`
	Site  string   `xml:"http://www.portalfiscal.inf.br/dce Site" json:"Site"`
}

type TDCeInfDCeTransportadora struct {
	CNPJ  TCNPJOpc `xml:"http://www.portalfiscal.inf.br/dce CNPJ" json:"CNPJ"`
	XNome string   `xml:"http://www.portalfiscal.inf.br/dce xNome" json:"xNome"`
}

type TDCeInfDCeECT struct {
	CNPJ  TCNPJOpc `xml:"http://www.portalfiscal.inf.br/dce CNPJ" json:"CNPJ"`
	XNome string   `xml:"http://www.portalfiscal.inf.br/dce xNome" json:"xNome"`
}

type TDCeInfDCeDest struct {
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/dce CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/dce CPF,omitempty" json:"CPF,omitempty"`
	IDOutros  *string   `xml:"http://www.portalfiscal.inf.br/dce idOutros,omitempty" json:"idOutros,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/dce xNome" json:"xNome"`
	EnderDest TEndeDest `xml:"http://www.portalfiscal.inf.br/dce enderDest" json:"enderDest"`
}

type TDCeInfDCeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/dce CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/dce CPF,omitempty" json:"CPF,omitempty"`
}

type TDCeInfDCeDetProd struct {
	XProd  string      `xml:"http://www.portalfiscal.inf.br/dce xProd" json:"xProd"`
	NCM    string      `xml:"http://www.portalfiscal.inf.br/dce NCM" json:"NCM"`
	QCom   TDec1104Opc `xml:"http://www.portalfiscal.inf.br/dce qCom" json:"qCom"`
	VUnCom string      `xml:"http://www.portalfiscal.inf.br/dce vUnCom" json:"vUnCom"`
	VProd  string      `xml:"http://www.portalfiscal.inf.br/dce vProd" json:"vProd"`
}

type TDCeInfDCeDet struct {
	NItem     string            `xml:"nItem,attr" json:"nItem"`
	Prod      TDCeInfDCeDetProd `xml:"http://www.portalfiscal.inf.br/dce prod" json:"prod"`
	InfAdProd *string           `xml:"http://www.portalfiscal.inf.br/dce infAdProd,omitempty" json:"infAdProd,omitempty"`
}

type TDCeInfDCeTotal struct {
	VDC TDec1302 `xml:"http://www.portalfiscal.inf.br/dce vDC" json:"vDC"`
}

type TDCeInfDCeTransp struct {
	ModTrans   string   `xml:"http://www.portalfiscal.inf.br/dce modTrans" json:"modTrans"`
	CNPJTransp TCNPJOpc `xml:"http://www.portalfiscal.inf.br/dce CNPJTransp" json:"CNPJTransp"`
}

type TDCeInfDCeInfAdicObsEmit struct {
	XCampo *string `xml:"xCampo,attr,omitempty" json:"xCampo,omitempty"`
	XTexto string  `xml:"http://www.portalfiscal.inf.br/dce xTexto" json:"xTexto"`
}

type TDCeInfDCeInfAdicObsFisco struct {
	XCampo *string `xml:"xCampo,attr,omitempty" json:"xCampo,omitempty"`
	XTexto string  `xml:"http://www.portalfiscal.inf.br/dce xTexto" json:"xTexto"`
}

type TDCeInfDCeInfAdicObsMarketplace struct {
	XCampo *string `xml:"xCampo,attr,omitempty" json:"xCampo,omitempty"`
	XTexto string  `xml:"http://www.portalfiscal.inf.br/dce xTexto" json:"xTexto"`
}

type TDCeInfDCeInfAdicObsECT struct {
	XCampo *string `xml:"xCampo,attr,omitempty" json:"xCampo,omitempty"`
	XTexto string  `xml:"http://www.portalfiscal.inf.br/dce xTexto" json:"xTexto"`
}

type TDCeInfDCeInfAdic struct {
	InfAdFisco       *string                           `xml:"http://www.portalfiscal.inf.br/dce infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl           *string                           `xml:"http://www.portalfiscal.inf.br/dce infCpl,omitempty" json:"infCpl,omitempty"`
	InfAdMarketplace *string                           `xml:"http://www.portalfiscal.inf.br/dce infAdMarketplace,omitempty" json:"infAdMarketplace,omitempty"`
	InfAdECT         *string                           `xml:"http://www.portalfiscal.inf.br/dce infAdECT,omitempty" json:"infAdECT,omitempty"`
	ObsEmit          []TDCeInfDCeInfAdicObsEmit        `xml:"http://www.portalfiscal.inf.br/dce obsEmit,omitempty" json:"obsEmit,omitempty"`
	ObsFisco         []TDCeInfDCeInfAdicObsFisco       `xml:"http://www.portalfiscal.inf.br/dce obsFisco,omitempty" json:"obsFisco,omitempty"`
	ObsMarketplace   []TDCeInfDCeInfAdicObsMarketplace `xml:"http://www.portalfiscal.inf.br/dce obsMarketplace,omitempty" json:"obsMarketplace,omitempty"`
	ObsECT           []TDCeInfDCeInfAdicObsECT         `xml:"http://www.portalfiscal.inf.br/dce obsECT,omitempty" json:"obsECT,omitempty"`
}

type TDCeInfDCeInfDec struct {
	XObs1 string `xml:"http://www.portalfiscal.inf.br/dce xObs1" json:"xObs1"`
	XObs2 string `xml:"http://www.portalfiscal.inf.br/dce xObs2" json:"xObs2"`
}

type TDCeInfDCe struct {
	Versao         TVerDCe                   `xml:"versao,attr" json:"versao"`
	ID             string                    `xml:"Id,attr" json:"Id"`
	Ide            TDCeInfDCeIde             `xml:"http://www.portalfiscal.inf.br/dce ide" json:"ide"`
	Emit           TDCeInfDCeEmit            `xml:"http://www.portalfiscal.inf.br/dce emit" json:"emit"`
	Fisco          *TDCeInfDCeFisco          `xml:"http://www.portalfiscal.inf.br/dce Fisco,omitempty" json:"Fisco,omitempty"`
	Marketplace    *TDCeInfDCeMarketplace    `xml:"http://www.portalfiscal.inf.br/dce Marketplace,omitempty" json:"Marketplace,omitempty"`
	Transportadora *TDCeInfDCeTransportadora `xml:"http://www.portalfiscal.inf.br/dce Transportadora,omitempty" json:"Transportadora,omitempty"`
	ECT            *TDCeInfDCeECT            `xml:"http://www.portalfiscal.inf.br/dce ECT,omitempty" json:"ECT,omitempty"`
	Dest           TDCeInfDCeDest            `xml:"http://www.portalfiscal.inf.br/dce dest" json:"dest"`
	AutXML         []TDCeInfDCeAutXML        `xml:"http://www.portalfiscal.inf.br/dce autXML,omitempty" json:"autXML,omitempty"`
	Det            []TDCeInfDCeDet           `xml:"http://www.portalfiscal.inf.br/dce det,omitempty" json:"det,omitempty"`
	Total          TDCeInfDCeTotal           `xml:"http://www.portalfiscal.inf.br/dce total" json:"total"`
	Transp         TDCeInfDCeTransp          `xml:"http://www.portalfiscal.inf.br/dce transp" json:"transp"`
	InfAdic        *TDCeInfDCeInfAdic        `xml:"http://www.portalfiscal.inf.br/dce infAdic,omitempty" json:"infAdic,omitempty"`
	InfDec         TDCeInfDCeInfDec          `xml:"http://www.portalfiscal.inf.br/dce infDec" json:"infDec"`
}

type TDCeInfSolicDCe struct {
	XSolic string `xml:"http://www.portalfiscal.inf.br/dce xSolic" json:"xSolic"`
}

type TDCeInfDCeSupl struct {
	QRCodDCe string `xml:"http://www.portalfiscal.inf.br/dce qrCodDCe" json:"qrCodDCe"`
	URLChave string `xml:"http://www.portalfiscal.inf.br/dce urlChave" json:"urlChave"`
}

type TDCe struct {
	InfDCe      TDCeInfDCe      `xml:"http://www.portalfiscal.inf.br/dce infDCe" json:"infDCe"`
	InfSolicDCe TDCeInfSolicDCe `xml:"http://www.portalfiscal.inf.br/dce infSolicDCe" json:"infSolicDCe"`
	InfDCeSupl  TDCeInfDCeSupl  `xml:"http://www.portalfiscal.inf.br/dce infDCeSupl" json:"infDCeSupl"`
	Signature   SignatureType   `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

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

type TDoc string

type TEmail string

type TEndeDest struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/dce xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/dce nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/dce xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/dce xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/dce cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/dce xMun" json:"xMun"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/dce UF" json:"UF"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/dce CEP" json:"CEP"`
	CPais   string      `xml:"http://www.portalfiscal.inf.br/dce cPais" json:"cPais"`
	XPais   string      `xml:"http://www.portalfiscal.inf.br/dce xPais" json:"xPais"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/dce fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/dce email,omitempty" json:"email,omitempty"`
}

type TEndeEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/dce xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/dce nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/dce xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/dce xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/dce cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/dce xMun" json:"xMun"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/dce UF" json:"UF"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/dce CEP" json:"CEP"`
	CPais   string      `xml:"http://www.portalfiscal.inf.br/dce cPais" json:"cPais"`
	XPais   string      `xml:"http://www.portalfiscal.inf.br/dce xPais" json:"xPais"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/dce fone,omitempty" json:"fone,omitempty"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/dce xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/dce nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/dce xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/dce xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/dce cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/dce xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/dce CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/dce UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/dce cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/dce xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/dce fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/dce email,omitempty" json:"email,omitempty"`
}

type TFinDCe string

type TIPv4 string

type TIe string

type TIeDest string

type TImp struct {
}

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModDCe string

type TMotivo string

type TNF string

type TNSU string

type TProcEmi string

type TProt string

type TProtDCeInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/dce tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/dce verAplic" json:"verAplic"`
	ChDCe    TChDFe           `xml:"http://www.portalfiscal.inf.br/dce chDCe" json:"chDCe"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/dce dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/dce nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/dce digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/dce cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/dce xMotivo" json:"xMotivo"`
}

type TProtDCeInfFisco struct {
	CMsg string  `xml:"http://www.portalfiscal.inf.br/dce cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/dce xMsg" json:"xMsg"`
}

type TProtDCe struct {
	Versao    TVerDCe           `xml:"versao,attr" json:"versao"`
	InfProt   TProtDCeInfProt   `xml:"http://www.portalfiscal.inf.br/dce infProt" json:"infProt"`
	InfFisco  *TProtDCeInfFisco `xml:"http://www.portalfiscal.inf.br/dce infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRec string

type TRetDCe struct {
	Versao   TVerDCe      `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb         `xml:"http://www.portalfiscal.inf.br/dce tpAmb" json:"tpAmb"`
	VerAplic TVerAplic    `xml:"http://www.portalfiscal.inf.br/dce verAplic" json:"verAplic"`
	CStat    TStat        `xml:"http://www.portalfiscal.inf.br/dce cStat" json:"cStat"`
	XMotivo  TMotivo      `xml:"http://www.portalfiscal.inf.br/dce xMotivo" json:"xMotivo"`
	CUF      TCodUFIBGE   `xml:"http://www.portalfiscal.inf.br/dce cUF" json:"cUF"`
	DhRecbto TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/dce dhRecbto" json:"dhRecbto"`
	ProtDCe  *TProtDCe    `xml:"http://www.portalfiscal.inf.br/dce protDCe,omitempty" json:"protDCe,omitempty"`
}

type TSerie string

type TServ string

type TStat string

type TString string

type TTime string

type TUF string

type TUFSemEX string

type TUtilizacao string

type TVerAplic string

type TVerDCe string

type Tano string

type TtpEmis string

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

// RetDCe represents the global element {http://www.portalfiscal.inf.br/dce}retDCe.
type RetDCe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/dce retDCe" json:"-"`
	TRetDCe
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
