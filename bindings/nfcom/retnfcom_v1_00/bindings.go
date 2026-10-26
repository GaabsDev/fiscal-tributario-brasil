// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package retnfcom_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TALCZFMCBS struct {
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vTribRegCBS" json:"vTribRegCBS"`
}

type TALCZFMCBSNFe struct {
	TpALCZFMCBS     string        `xml:"http://www.portalfiscal.inf.br/nfcom tpALCZFMCBS" json:"tpALCZFMCBS"`
	NProcSuframa    *string       `xml:"http://www.portalfiscal.inf.br/nfcom nProcSuframa,omitempty" json:"nProcSuframa,omitempty"`
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vTribRegCBS" json:"vTribRegCBS"`
}

type TAjusteCompet struct {
	CompetApur TCompetApur `xml:"http://www.portalfiscal.inf.br/nfcom competApur" json:"competApur"`
	VIBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBS" json:"vIBS"`
	VCBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBS" json:"vCBS"`
}

type TAjusteNFComAnt string

type TAjusteNFComAntV2 string

type TAmb string

type TAssinante string

type TCIBSGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfcom gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfcom gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfcom gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSUF" json:"vIBSUF"`
}

type TCIBSGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfcom gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfcom gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfcom gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMun" json:"vIBSMun"`
}

type TCIBSGCBS struct {
	PCBS       TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pCBS" json:"pCBS"`
	GDif       *TDif         `xml:"http://www.portalfiscal.inf.br/nfcom gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfcom gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed         `xml:"http://www.portalfiscal.inf.br/nfcom gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBS   `xml:"http://www.portalfiscal.inf.br/nfcom gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vCBS" json:"vCBS"`
}

type TCIBS struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	GIBSUF         TCIBSGIBSUF     `xml:"http://www.portalfiscal.inf.br/nfcom gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSGIBSMun    `xml:"http://www.portalfiscal.inf.br/nfcom gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfcom vIBS" json:"vIBS"`
	GCBS           TCIBSGCBS       `xml:"http://www.portalfiscal.inf.br/nfcom gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nfcom gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfcom gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCIBSNFeGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfcom gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfcom gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfcom gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSUF" json:"vIBSUF"`
}

type TCIBSNFeGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfcom gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfcom gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfcom gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMun" json:"vIBSMun"`
}

type TCIBSNFeGCBS struct {
	PCBS       TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfcom pCBS" json:"pCBS"`
	GDif       *TDif          `xml:"http://www.portalfiscal.inf.br/nfcom gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib      `xml:"http://www.portalfiscal.inf.br/nfcom gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed          `xml:"http://www.portalfiscal.inf.br/nfcom gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBSNFe `xml:"http://www.portalfiscal.inf.br/nfcom gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfcom vCBS" json:"vCBS"`
}

type TCIBSNFe struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	GIBSUF         TCIBSNFeGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfcom gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSNFeGIBSMun `xml:"http://www.portalfiscal.inf.br/nfcom gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfcom vIBS" json:"vIBS"`
	GCBS           TCIBSNFeGCBS    `xml:"http://www.portalfiscal.inf.br/nfcom gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nfcom gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfcom gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCOrgaoIBGE string

type TCRT string

type TCST string

type TCFOP string

type TChDFe string

type TChDFeRTC string

type TCNPJ string

type TCNPJBaseRTC string

type TCNPJOpc string

type TCNPJRTC string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

type TCompFat string

type TCompet string

type TCompetApur string

type TCompraGov struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nfcom tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfcom pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nfcom tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nfcom refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCompraGovReduzido struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nfcom tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfcom pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nfcom tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nfcom refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCPF string

type TCPFVar string

type TCredPres struct {
	PCredPres        TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pCredPres" json:"pCredPres"`
	VCredPres        *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfcom vCredPres,omitempty" json:"vCredPres,omitempty"`
	VCredPresCondSus *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfcom vCredPresCondSus,omitempty" json:"vCredPresCondSus,omitempty"`
}

type TCredPresIBSZFM struct {
	CompetApur       TCompetApur       `xml:"http://www.portalfiscal.inf.br/nfcom competApur" json:"competApur"`
	TpCredPresIBSZFM TTpCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nfcom tpCredPresIBSZFM" json:"tpCredPresIBSZFM"`
	VCredPresIBSZFM  TDec1302RTC       `xml:"http://www.portalfiscal.inf.br/nfcom vCredPresIBSZFM" json:"vCredPresIBSZFM"`
}

type TCredPresOper struct {
	VBCCredPres  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vBCCredPres" json:"vBCCredPres"`
	CCredPres    TcCredPres  `xml:"http://www.portalfiscal.inf.br/nfcom cCredPres" json:"cCredPres"`
	GIBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCredPres,omitempty" json:"gIBSCredPres,omitempty"`
	GCBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nfcom gCBSCredPres,omitempty" json:"gCBSCredPres,omitempty"`
}

type TData string

type TDateTimeUTC string

type TDec1104RTC string

type TDec1302RTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec030204RTC string

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

type TDec1104OpRTC string

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

type TDevTrib struct {
	PDevTrib *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pDevTrib,omitempty" json:"pDevTrib,omitempty"`
	VDevTrib TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfcom vDevTrib" json:"vDevTrib"`
}

type TDif struct {
	PDif TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pDif" json:"pDif"`
	VDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vDif" json:"vDif"`
}

type TDoc string

type TEmail string

type TEmpresaSoft struct {
	CNPJ     *TCNPJ `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF      *TCPF  `xml:"http://www.portalfiscal.inf.br/nfcom CPF,omitempty" json:"CPF,omitempty"`
	XContato string `xml:"http://www.portalfiscal.inf.br/nfcom xContato" json:"xContato"`
	Email    TEmail `xml:"http://www.portalfiscal.inf.br/nfcom email" json:"email"`
	Fone     string `xml:"http://www.portalfiscal.inf.br/nfcom fone" json:"fone"`
}

type TEndeDest struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfcom xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfcom nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfcom xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfcom xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfcom cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfcom xMun" json:"xMun"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/nfcom CEP" json:"CEP"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/nfcom UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/nfcom cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/nfcom xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfcom fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nfcom email,omitempty" json:"email,omitempty"`
}

type TEndeEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfcom xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfcom nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfcom xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfcom xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfcom cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfcom xMun" json:"xMun"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/nfcom CEP" json:"CEP"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/nfcom UF" json:"UF"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfcom fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nfcom email,omitempty" json:"email,omitempty"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfcom xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfcom nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfcom xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfcom xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfcom cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfcom xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/nfcom CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/nfcom UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/nfcom cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/nfcom xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfcom fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nfcom email,omitempty" json:"email,omitempty"`
}

type TEnteGov string

type TEnviNFCom struct {
	Versao TVerNFCom `xml:"versao,attr" json:"versao"`
	IDLote TIDLote   `xml:"http://www.portalfiscal.inf.br/nfcom idLote" json:"idLote"`
	NFCom  []TNFCom  `xml:"http://www.portalfiscal.inf.br/nfcom NFCom,omitempty" json:"NFCom,omitempty"`
}

type TEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSEstCred" json:"vCBSEstCred"`
}

type TFatNFCom string

type TFinNFCom string

type TIBSCBSMonoTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSUF" json:"vIBSUF"`
}

type TIBSCBSMonoTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMun" json:"vIBSMun"`
}

type TIBSCBSMonoTotGIBS struct {
	GIBSUF           TIBSCBSMonoTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfcom gIBSUF" json:"gIBSUF"`
	GIBSMun          TIBSCBSMonoTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nfcom gIBSMun" json:"gIBSMun"`
	VIBS             TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfcom vIBS" json:"vIBS"`
	VCredPres        TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfcom vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfcom vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGCBS struct {
	VDif             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDif" json:"vDif"`
	VDevTrib         TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDevTrib" json:"vDevTrib"`
	VCBS             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBS" json:"vCBS"`
	VCredPres        TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGMono struct {
	VIBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMono" json:"vIBSMono"`
	VCBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSMono" json:"vCBSMono"`
	VIBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMonoReten" json:"vIBSMonoReten"`
	VCBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSMonoReten" json:"vCBSMonoReten"`
	VIBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMonoRet" json:"vIBSMonoRet"`
	VCBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSMonoRet" json:"vCBSMonoRet"`
}

type TIBSCBSMonoTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSMonoTot struct {
	VBCIBSCBS    TDec1302RTC                 `xml:"http://www.portalfiscal.inf.br/nfcom vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         *TIBSCBSMonoTotGIBS         `xml:"http://www.portalfiscal.inf.br/nfcom gIBS,omitempty" json:"gIBS,omitempty"`
	GCBS         *TIBSCBSMonoTotGCBS         `xml:"http://www.portalfiscal.inf.br/nfcom gCBS,omitempty" json:"gCBS,omitempty"`
	GMono        *TIBSCBSMonoTotGMono        `xml:"http://www.portalfiscal.inf.br/nfcom gMono,omitempty" json:"gMono,omitempty"`
	GEstornoCred *TIBSCBSMonoTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIBSCBSTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSUF" json:"vIBSUF"`
}

type TIBSCBSTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMun" json:"vIBSMun"`
}

type TIBSCBSTotGIBS struct {
	GIBSUF  TIBSCBSTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfcom gIBSUF" json:"gIBSUF"`
	GIBSMun TIBSCBSTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nfcom gIBSMun" json:"gIBSMun"`
	VIBS    TDec1302RTC           `xml:"http://www.portalfiscal.inf.br/nfcom vIBS" json:"vIBS"`
}

type TIBSCBSTotGCBS struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vDevTrib" json:"vDevTrib"`
	VCBS     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBS" json:"vCBS"`
}

type TIBSCBSTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSTot struct {
	VBCIBSCBS    TDec1302RTC             `xml:"http://www.portalfiscal.inf.br/nfcom vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         TIBSCBSTotGIBS          `xml:"http://www.portalfiscal.inf.br/nfcom gIBS" json:"gIBS"`
	GCBS         TIBSCBSTotGCBS          `xml:"http://www.portalfiscal.inf.br/nfcom gCBS" json:"gCBS"`
	GEstornoCred *TIBSCBSTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIPv4 string

type TIS struct {
	CSTIS        TCST           `xml:"http://www.portalfiscal.inf.br/nfcom CSTIS" json:"CSTIS"`
	CClassTribIS TcClassTrib    `xml:"http://www.portalfiscal.inf.br/nfcom cClassTribIS" json:"cClassTribIS"`
	VBCIS        *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vBCIS,omitempty" json:"vBCIS,omitempty"`
	PIS          *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pIS,omitempty" json:"pIS,omitempty"`
	AdRemIS      *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom adRemIS,omitempty" json:"adRemIS,omitempty"`
	UTrib        *string        `xml:"http://www.portalfiscal.inf.br/nfcom uTrib,omitempty" json:"uTrib,omitempty"`
	QTrib        *TDec1104OpRTC `xml:"http://www.portalfiscal.inf.br/nfcom qTrib,omitempty" json:"qTrib,omitempty"`
	VIS          *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIS,omitempty" json:"vIS,omitempty"`
}

type TISTot struct {
	VIS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIS" json:"vIS"`
}

type TIDLote string

type TIe string

type TIeDest string

type TImpICMS00 struct {
	CST   string          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	VBC   TDec1302        `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	PICMS TDec0302        `xml:"http://www.portalfiscal.inf.br/nfcom pICMS" json:"pICMS"`
	VICMS TDec1302        `xml:"http://www.portalfiscal.inf.br/nfcom vICMS" json:"vICMS"`
	PFCP  *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfcom pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP  *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfcom vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpICMS20 struct {
	CST        string          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	PRedBC     TDec0302Opc     `xml:"http://www.portalfiscal.inf.br/nfcom pRedBC" json:"pRedBC"`
	VBC        TDec1302        `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	PICMS      TDec0302        `xml:"http://www.portalfiscal.inf.br/nfcom pICMS" json:"pICMS"`
	VICMS      TDec1302        `xml:"http://www.portalfiscal.inf.br/nfcom vICMS" json:"vICMS"`
	VICMSDeson *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfcom vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string         `xml:"http://www.portalfiscal.inf.br/nfcom cBenef,omitempty" json:"cBenef,omitempty"`
	PFCP       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfcom pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfcom vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpICMS40 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/nfcom cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS51 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/nfcom cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS90 struct {
	CST        string          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	VBC        *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfcom vBC,omitempty" json:"vBC,omitempty"`
	PICMS      *TDec0302       `xml:"http://www.portalfiscal.inf.br/nfcom pICMS,omitempty" json:"pICMS,omitempty"`
	VICMS      *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfcom vICMS,omitempty" json:"vICMS,omitempty"`
	VICMSDeson *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfcom vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string         `xml:"http://www.portalfiscal.inf.br/nfcom cBenef,omitempty" json:"cBenef,omitempty"`
	PFCP       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfcom pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfcom vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpICMSSN struct {
	CST   string `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	IndSN string `xml:"http://www.portalfiscal.inf.br/nfcom indSN" json:"indSN"`
}

type TImpGICMSPrevistoPagtoAntecip struct {
	VICMSPrevisto TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vICMSPrevisto" json:"vICMSPrevisto"`
}

type TImpICMSUFDest struct {
	CUFDest      TCodUFIBGE `xml:"cUFDest,attr" json:"cUFDest"`
	VBCUFDest    TDec1302   `xml:"http://www.portalfiscal.inf.br/nfcom vBCUFDest" json:"vBCUFDest"`
	PFCPUFDest   TDec0302   `xml:"http://www.portalfiscal.inf.br/nfcom pFCPUFDest" json:"pFCPUFDest"`
	PICMSUFDest  TDec0302   `xml:"http://www.portalfiscal.inf.br/nfcom pICMSUFDest" json:"pICMSUFDest"`
	VFCPUFDest   TDec1302   `xml:"http://www.portalfiscal.inf.br/nfcom vFCPUFDest" json:"vFCPUFDest"`
	VICMSUFDest  TDec1302   `xml:"http://www.portalfiscal.inf.br/nfcom vICMSUFDest" json:"vICMSUFDest"`
	VICMSUFEmi   TDec1302   `xml:"http://www.portalfiscal.inf.br/nfcom vICMSUFEmi" json:"vICMSUFEmi"`
	CBenefUFDest *string    `xml:"http://www.portalfiscal.inf.br/nfcom cBenefUFDest,omitempty" json:"cBenefUFDest,omitempty"`
}

type TImpPIS struct {
	CST  string      `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	VBC  TDec1302    `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	PPIS TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfcom pPIS" json:"pPIS"`
	VPIS TDec1302    `xml:"http://www.portalfiscal.inf.br/nfcom vPIS" json:"vPIS"`
}

type TImpCOFINS struct {
	CST     string      `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	VBC     TDec1302    `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	PCOFINS TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfcom pCOFINS" json:"pCOFINS"`
	VCOFINS TDec1302    `xml:"http://www.portalfiscal.inf.br/nfcom vCOFINS" json:"vCOFINS"`
}

type TImpFUST struct {
	VBC   TDec1302    `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	PFUST TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfcom pFUST" json:"pFUST"`
	VFUST TDec1302    `xml:"http://www.portalfiscal.inf.br/nfcom vFUST" json:"vFUST"`
}

type TImpFUNTTEL struct {
	VBC      TDec1302    `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	PFUNTTEL TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfcom pFUNTTEL" json:"pFUNTTEL"`
	VFUNTTEL TDec1302    `xml:"http://www.portalfiscal.inf.br/nfcom vFUNTTEL" json:"vFUNTTEL"`
}

type TImpRetTrib struct {
	VRetPIS    TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfcom vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfcom vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfcom vRetCSLL" json:"vRetCSLL"`
	VBCIRRF    TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfcom vBCIRRF" json:"vBCIRRF"`
	VIRRF      TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfcom vIRRF" json:"vIRRF"`
}

type TImp struct {
	ICMS00                    *TImpICMS00                    `xml:"http://www.portalfiscal.inf.br/nfcom ICMS00,omitempty" json:"ICMS00,omitempty"`
	ICMS20                    *TImpICMS20                    `xml:"http://www.portalfiscal.inf.br/nfcom ICMS20,omitempty" json:"ICMS20,omitempty"`
	ICMS40                    *TImpICMS40                    `xml:"http://www.portalfiscal.inf.br/nfcom ICMS40,omitempty" json:"ICMS40,omitempty"`
	ICMS51                    *TImpICMS51                    `xml:"http://www.portalfiscal.inf.br/nfcom ICMS51,omitempty" json:"ICMS51,omitempty"`
	ICMS90                    *TImpICMS90                    `xml:"http://www.portalfiscal.inf.br/nfcom ICMS90,omitempty" json:"ICMS90,omitempty"`
	ICMSSN                    *TImpICMSSN                    `xml:"http://www.portalfiscal.inf.br/nfcom ICMSSN,omitempty" json:"ICMSSN,omitempty"`
	GICMSPrevistoPagtoAntecip *TImpGICMSPrevistoPagtoAntecip `xml:"http://www.portalfiscal.inf.br/nfcom gICMSPrevistoPagtoAntecip,omitempty" json:"gICMSPrevistoPagtoAntecip,omitempty"`
	ICMSUFDest                []TImpICMSUFDest               `xml:"http://www.portalfiscal.inf.br/nfcom ICMSUFDest,omitempty" json:"ICMSUFDest,omitempty"`
	IndSemCST                 *string                        `xml:"http://www.portalfiscal.inf.br/nfcom indSemCST,omitempty" json:"indSemCST,omitempty"`
	PIS                       *TImpPIS                       `xml:"http://www.portalfiscal.inf.br/nfcom PIS,omitempty" json:"PIS,omitempty"`
	COFINS                    *TImpCOFINS                    `xml:"http://www.portalfiscal.inf.br/nfcom COFINS,omitempty" json:"COFINS,omitempty"`
	FUST                      *TImpFUST                      `xml:"http://www.portalfiscal.inf.br/nfcom FUST,omitempty" json:"FUST,omitempty"`
	FUNTTEL                   *TImpFUNTTEL                   `xml:"http://www.portalfiscal.inf.br/nfcom FUNTTEL,omitempty" json:"FUNTTEL,omitempty"`
	RetTrib                   *TImpRetTrib                   `xml:"http://www.portalfiscal.inf.br/nfcom retTrib,omitempty" json:"retTrib,omitempty"`
	IBSCBS                    *TTribNFCom                    `xml:"http://www.portalfiscal.inf.br/nfcom IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type TIndDoacao string

type TIndIEDest string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModNFCom string

type TModNFPapel string

type TMonofasiaGMonoPadrao struct {
	QBCMono  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfcom qBCMono" json:"qBCMono"`
	AdRemIBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom adRemIBS" json:"adRemIBS"`
	AdRemCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom adRemCBS" json:"adRemCBS"`
	VIBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMono" json:"vIBSMono"`
	VCBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vCBSMono" json:"vCBSMono"`
}

type TMonofasiaGMonoReten struct {
	QBCMonoReten  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfcom qBCMonoReten" json:"qBCMonoReten"`
	AdRemIBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom adRemIBSReten" json:"adRemIBSReten"`
	VIBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMonoReten" json:"vIBSMonoReten"`
	AdRemCBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom adRemCBSReten" json:"adRemCBSReten"`
	VCBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vCBSMonoReten" json:"vCBSMonoReten"`
}

type TMonofasiaGMonoRet struct {
	QBCMonoRet  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfcom qBCMonoRet" json:"qBCMonoRet"`
	AdRemIBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom adRemIBSRet" json:"adRemIBSRet"`
	VIBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMonoRet" json:"vIBSMonoRet"`
	AdRemCBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom adRemCBSRet" json:"adRemCBSRet"`
	VCBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vCBSMonoRet" json:"vCBSMonoRet"`
}

type TMonofasiaGMonoDif struct {
	PDifIBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pDifIBS" json:"pDifIBS"`
	VIBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMonoDif" json:"vIBSMonoDif"`
	PDifCBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pDifCBS" json:"pDifCBS"`
	VCBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vCBSMonoDif" json:"vCBSMonoDif"`
}

type TMonofasia struct {
	GMonoPadrao     *TMonofasiaGMonoPadrao `xml:"http://www.portalfiscal.inf.br/nfcom gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten      *TMonofasiaGMonoReten  `xml:"http://www.portalfiscal.inf.br/nfcom gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet        *TMonofasiaGMonoRet    `xml:"http://www.portalfiscal.inf.br/nfcom gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GMonoDif        *TMonofasiaGMonoDif    `xml:"http://www.portalfiscal.inf.br/nfcom gMonoDif,omitempty" json:"gMonoDif,omitempty"`
	VTotIBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/nfcom vTotIBSMonoItem" json:"vTotIBSMonoItem"`
	VTotCBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/nfcom vTotCBSMonoItem" json:"vTotCBSMonoItem"`
}

type TMotAjuste string

type TMotSub string

type TMotivo string

type TNF string

type TNFComInfNFComIde struct {
	CUF                TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/nfcom cUF" json:"cUF"`
	TpAmb              TAmb                `xml:"http://www.portalfiscal.inf.br/nfcom tpAmb" json:"tpAmb"`
	Mod                TModNFCom           `xml:"http://www.portalfiscal.inf.br/nfcom mod" json:"mod"`
	Serie              TSerie              `xml:"http://www.portalfiscal.inf.br/nfcom serie" json:"serie"`
	NNF                TNF                 `xml:"http://www.portalfiscal.inf.br/nfcom nNF" json:"nNF"`
	CNF                string              `xml:"http://www.portalfiscal.inf.br/nfcom cNF" json:"cNF"`
	CDV                string              `xml:"http://www.portalfiscal.inf.br/nfcom cDV" json:"cDV"`
	DhEmi              TDateTimeUTC        `xml:"http://www.portalfiscal.inf.br/nfcom dhEmi" json:"dhEmi"`
	TpEmis             TtpEmis             `xml:"http://www.portalfiscal.inf.br/nfcom tpEmis" json:"tpEmis"`
	NSiteAutoriz       string              `xml:"http://www.portalfiscal.inf.br/nfcom nSiteAutoriz" json:"nSiteAutoriz"`
	CMunFG             TCodMunIBGE         `xml:"http://www.portalfiscal.inf.br/nfcom cMunFG" json:"cMunFG"`
	FinNFCom           TFinNFCom           `xml:"http://www.portalfiscal.inf.br/nfcom finNFCom" json:"finNFCom"`
	TpFat              TFatNFCom           `xml:"http://www.portalfiscal.inf.br/nfcom tpFat" json:"tpFat"`
	VerProc            string              `xml:"http://www.portalfiscal.inf.br/nfcom verProc" json:"verProc"`
	IndPrePago         *string             `xml:"http://www.portalfiscal.inf.br/nfcom indPrePago,omitempty" json:"indPrePago,omitempty"`
	IndCessaoMeiosRede *string             `xml:"http://www.portalfiscal.inf.br/nfcom indCessaoMeiosRede,omitempty" json:"indCessaoMeiosRede,omitempty"`
	IndNotaEntrada     *string             `xml:"http://www.portalfiscal.inf.br/nfcom indNotaEntrada,omitempty" json:"indNotaEntrada,omitempty"`
	DhCont             *TDateTimeUTC       `xml:"http://www.portalfiscal.inf.br/nfcom dhCont,omitempty" json:"dhCont,omitempty"`
	XJust              *string             `xml:"http://www.portalfiscal.inf.br/nfcom xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov         *TCompraGovReduzido `xml:"http://www.portalfiscal.inf.br/nfcom gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	TpPagAnt           *TPagAnt            `xml:"http://www.portalfiscal.inf.br/nfcom tpPagAnt,omitempty" json:"tpPagAnt,omitempty"`
}

type TNFComInfNFComEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ" json:"CNPJ"`
	IE        TIe      `xml:"http://www.portalfiscal.inf.br/nfcom IE" json:"IE"`
	IEUFDest  *string  `xml:"http://www.portalfiscal.inf.br/nfcom IEUFDest,omitempty" json:"IEUFDest,omitempty"`
	CRT       TCRT     `xml:"http://www.portalfiscal.inf.br/nfcom CRT" json:"CRT"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/nfcom xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/nfcom xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/nfcom enderEmit" json:"enderEmit"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/nfcom ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TNFComInfNFComDest struct {
	XNome     string     `xml:"http://www.portalfiscal.inf.br/nfcom xNome" json:"xNome"`
	CNPJ      *TCNPJOpc  `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF      `xml:"http://www.portalfiscal.inf.br/nfcom CPF,omitempty" json:"CPF,omitempty"`
	IDOutros  *string    `xml:"http://www.portalfiscal.inf.br/nfcom idOutros,omitempty" json:"idOutros,omitempty"`
	IndIEDest TIndIEDest `xml:"http://www.portalfiscal.inf.br/nfcom indIEDest" json:"indIEDest"`
	IE        *TIeDest   `xml:"http://www.portalfiscal.inf.br/nfcom IE,omitempty" json:"IE,omitempty"`
	IM        *string    `xml:"http://www.portalfiscal.inf.br/nfcom IM,omitempty" json:"IM,omitempty"`
	EnderDest TEndeDest  `xml:"http://www.portalfiscal.inf.br/nfcom enderDest" json:"enderDest"`
}

type TNFComInfNFComAssinante struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type TNFComInfNFComGSubGNF struct {
	CNPJ       TCNPJ       `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ" json:"CNPJ"`
	Mod        TModNFPapel `xml:"http://www.portalfiscal.inf.br/nfcom mod" json:"mod"`
	Serie      string      `xml:"http://www.portalfiscal.inf.br/nfcom serie" json:"serie"`
	NNF        TNF         `xml:"http://www.portalfiscal.inf.br/nfcom nNF" json:"nNF"`
	CompetEmis TCompet     `xml:"http://www.portalfiscal.inf.br/nfcom CompetEmis" json:"CompetEmis"`
	Hash115    *string     `xml:"http://www.portalfiscal.inf.br/nfcom hash115,omitempty" json:"hash115,omitempty"`
}

type TNFComInfNFComGSub struct {
	ChNFCom *TChDFe                `xml:"http://www.portalfiscal.inf.br/nfcom chNFCom,omitempty" json:"chNFCom,omitempty"`
	GNF     *TNFComInfNFComGSubGNF `xml:"http://www.portalfiscal.inf.br/nfcom gNF,omitempty" json:"gNF,omitempty"`
	MotSub  TMotSub                `xml:"http://www.portalfiscal.inf.br/nfcom motSub" json:"motSub"`
}

type TNFComInfNFComGCofatGNF struct {
	CNPJ       TCNPJ       `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ" json:"CNPJ"`
	Mod        TModNFPapel `xml:"http://www.portalfiscal.inf.br/nfcom mod" json:"mod"`
	Serie      string      `xml:"http://www.portalfiscal.inf.br/nfcom serie" json:"serie"`
	NNF        TNF         `xml:"http://www.portalfiscal.inf.br/nfcom nNF" json:"nNF"`
	CompetEmis TCompet     `xml:"http://www.portalfiscal.inf.br/nfcom CompetEmis" json:"CompetEmis"`
	Hash115    *string     `xml:"http://www.portalfiscal.inf.br/nfcom hash115,omitempty" json:"hash115,omitempty"`
}

type TNFComInfNFComGCofat struct {
	ChNFComLocal *TChDFe                  `xml:"http://www.portalfiscal.inf.br/nfcom chNFComLocal,omitempty" json:"chNFComLocal,omitempty"`
	GNF          *TNFComInfNFComGCofatGNF `xml:"http://www.portalfiscal.inf.br/nfcom gNF,omitempty" json:"gNF,omitempty"`
}

type TNFComInfNFComDetProdGPagAntecipado struct {
	ChDFePagAnt TChDFe  `xml:"http://www.portalfiscal.inf.br/nfcom chDFePagAnt" json:"chDFePagAnt"`
	NItemPagAnt *string `xml:"http://www.portalfiscal.inf.br/nfcom nItemPagAnt,omitempty" json:"nItemPagAnt,omitempty"`
}

type TNFComInfNFComDetProd struct {
	CProd          string                               `xml:"http://www.portalfiscal.inf.br/nfcom cProd" json:"cProd"`
	XProd          string                               `xml:"http://www.portalfiscal.inf.br/nfcom xProd" json:"xProd"`
	CClass         string                               `xml:"http://www.portalfiscal.inf.br/nfcom cClass" json:"cClass"`
	CFOP           *TCFOP                               `xml:"http://www.portalfiscal.inf.br/nfcom CFOP,omitempty" json:"CFOP,omitempty"`
	CNPJLD         *TCNPJ                               `xml:"http://www.portalfiscal.inf.br/nfcom CNPJLD,omitempty" json:"CNPJLD,omitempty"`
	UMed           TUMedItem                            `xml:"http://www.portalfiscal.inf.br/nfcom uMed" json:"uMed"`
	QFaturada      TDec11001104                         `xml:"http://www.portalfiscal.inf.br/nfcom qFaturada" json:"qFaturada"`
	VItem          string                               `xml:"http://www.portalfiscal.inf.br/nfcom vItem" json:"vItem"`
	VDesc          *TDec1302Opc                         `xml:"http://www.portalfiscal.inf.br/nfcom vDesc,omitempty" json:"vDesc,omitempty"`
	VOutro         *TDec1302Opc                         `xml:"http://www.portalfiscal.inf.br/nfcom vOutro,omitempty" json:"vOutro,omitempty"`
	VProd          string                               `xml:"http://www.portalfiscal.inf.br/nfcom vProd" json:"vProd"`
	VItemLiq       *string                              `xml:"http://www.portalfiscal.inf.br/nfcom vItemLiq,omitempty" json:"vItemLiq,omitempty"`
	VProdLiq       *string                              `xml:"http://www.portalfiscal.inf.br/nfcom vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	DExpiracao     *TData                               `xml:"http://www.portalfiscal.inf.br/nfcom dExpiracao,omitempty" json:"dExpiracao,omitempty"`
	IndDevolucao   *string                              `xml:"http://www.portalfiscal.inf.br/nfcom indDevolucao,omitempty" json:"indDevolucao,omitempty"`
	CNPJCobrTerc   *TCNPJ                               `xml:"http://www.portalfiscal.inf.br/nfcom CNPJCobrTerc,omitempty" json:"CNPJCobrTerc,omitempty"`
	GPagAntecipado *TNFComInfNFComDetProdGPagAntecipado `xml:"http://www.portalfiscal.inf.br/nfcom gPagAntecipado,omitempty" json:"gPagAntecipado,omitempty"`
}

type TNFComInfNFComDetGProcRefGIBSCBS struct {
	VBC     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	VIBSUF  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSUF" json:"vIBSUF"`
	VIBSMun TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSMun" json:"vIBSMun"`
	VCBS    TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBS" json:"vCBS"`
}

type TNFComInfNFComDetGProcRefGProc struct {
	TpProc    TProcesso `xml:"http://www.portalfiscal.inf.br/nfcom tpProc" json:"tpProc"`
	NProcesso string    `xml:"http://www.portalfiscal.inf.br/nfcom nProcesso" json:"nProcesso"`
}

type TNFComInfNFComDetGProcRef struct {
	VItem        string                            `xml:"http://www.portalfiscal.inf.br/nfcom vItem" json:"vItem"`
	QFaturada    TDec11001104                      `xml:"http://www.portalfiscal.inf.br/nfcom qFaturada" json:"qFaturada"`
	VProd        string                            `xml:"http://www.portalfiscal.inf.br/nfcom vProd" json:"vProd"`
	VDesc        *TDec1302Opc                      `xml:"http://www.portalfiscal.inf.br/nfcom vDesc,omitempty" json:"vDesc,omitempty"`
	VOutro       *TDec1302Opc                      `xml:"http://www.portalfiscal.inf.br/nfcom vOutro,omitempty" json:"vOutro,omitempty"`
	IndDevolucao *string                           `xml:"http://www.portalfiscal.inf.br/nfcom indDevolucao,omitempty" json:"indDevolucao,omitempty"`
	VBC          *TDec1302                         `xml:"http://www.portalfiscal.inf.br/nfcom vBC,omitempty" json:"vBC,omitempty"`
	PICMS        *TDec0302                         `xml:"http://www.portalfiscal.inf.br/nfcom pICMS,omitempty" json:"pICMS,omitempty"`
	VICMS        *TDec1302                         `xml:"http://www.portalfiscal.inf.br/nfcom vICMS,omitempty" json:"vICMS,omitempty"`
	VPIS         *TDec1302                         `xml:"http://www.portalfiscal.inf.br/nfcom vPIS,omitempty" json:"vPIS,omitempty"`
	VCOFINS      *TDec1302                         `xml:"http://www.portalfiscal.inf.br/nfcom vCOFINS,omitempty" json:"vCOFINS,omitempty"`
	VFCP         *TDec1302                         `xml:"http://www.portalfiscal.inf.br/nfcom vFCP,omitempty" json:"vFCP,omitempty"`
	GIBSCBS      *TNFComInfNFComDetGProcRefGIBSCBS `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GProc        []TNFComInfNFComDetGProcRefGProc  `xml:"http://www.portalfiscal.inf.br/nfcom gProc,omitempty" json:"gProc,omitempty"`
}

type TNFComInfNFComDetGRessarc struct {
	TpRessarc    TRessarcimento `xml:"http://www.portalfiscal.inf.br/nfcom tpRessarc" json:"tpRessarc"`
	DRef         TData          `xml:"http://www.portalfiscal.inf.br/nfcom dRef" json:"dRef"`
	NProcesso    *string        `xml:"http://www.portalfiscal.inf.br/nfcom nProcesso,omitempty" json:"nProcesso,omitempty"`
	NProtReclama *string        `xml:"http://www.portalfiscal.inf.br/nfcom nProtReclama,omitempty" json:"nProtReclama,omitempty"`
	XObs         *string        `xml:"http://www.portalfiscal.inf.br/nfcom xObs,omitempty" json:"xObs,omitempty"`
}

type TNFComInfNFComDet struct {
	NItem                      string                     `xml:"nItem,attr" json:"nItem"`
	ChNFComAnt                 *TChDFe                    `xml:"chNFComAnt,attr,omitempty" json:"chNFComAnt,omitempty"`
	NItemAnt                   *string                    `xml:"nItemAnt,attr,omitempty" json:"nItemAnt,omitempty"`
	IndNFComAntPapelFatCentral *string                    `xml:"indNFComAntPapelFatCentral,attr,omitempty" json:"indNFComAntPapelFatCentral,omitempty"`
	Prod                       TNFComInfNFComDetProd      `xml:"http://www.portalfiscal.inf.br/nfcom prod" json:"prod"`
	Imposto                    TImp                       `xml:"http://www.portalfiscal.inf.br/nfcom imposto" json:"imposto"`
	GProcRef                   *TNFComInfNFComDetGProcRef `xml:"http://www.portalfiscal.inf.br/nfcom gProcRef,omitempty" json:"gProcRef,omitempty"`
	GRessarc                   *TNFComInfNFComDetGRessarc `xml:"http://www.portalfiscal.inf.br/nfcom gRessarc,omitempty" json:"gRessarc,omitempty"`
	InfAdProd                  *string                    `xml:"http://www.portalfiscal.inf.br/nfcom infAdProd,omitempty" json:"infAdProd,omitempty"`
}

type TNFComInfNFComPgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/nfcom pgto,omitempty" json:"pgto,omitempty"`
}

type TNFComInfNFComTotalICMSTot struct {
	VBC        TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vBC" json:"vBC"`
	VICMS      TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vICMS" json:"vICMS"`
	VICMSDeson TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vICMSDeson" json:"vICMSDeson"`
	VFCP       TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vFCP" json:"vFCP"`
}

type TNFComInfNFComTotalVRetTribTot struct {
	VRetPIS    TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vRetCSLL" json:"vRetCSLL"`
	VIRRF      TDec1302 `xml:"http://www.portalfiscal.inf.br/nfcom vIRRF" json:"vIRRF"`
}

type TNFComInfNFComTotal struct {
	VProd       TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfcom vProd" json:"vProd"`
	VProdLiq    *TDec1302                      `xml:"http://www.portalfiscal.inf.br/nfcom vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	ICMSTot     TNFComInfNFComTotalICMSTot     `xml:"http://www.portalfiscal.inf.br/nfcom ICMSTot" json:"ICMSTot"`
	VCOFINS     TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfcom vCOFINS" json:"vCOFINS"`
	VPIS        TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfcom vPIS" json:"vPIS"`
	VFUNTTEL    TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfcom vFUNTTEL" json:"vFUNTTEL"`
	VFUST       TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfcom vFUST" json:"vFUST"`
	VRetTribTot TNFComInfNFComTotalVRetTribTot `xml:"http://www.portalfiscal.inf.br/nfcom vRetTribTot" json:"vRetTribTot"`
	VDesc       TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfcom vDesc" json:"vDesc"`
	VOutro      TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfcom vOutro" json:"vOutro"`
	VNF         TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfcom vNF" json:"vNF"`
	IBSCBSTot   *TIBSCBSTot                    `xml:"http://www.portalfiscal.inf.br/nfcom IBSCBSTot,omitempty" json:"IBSCBSTot,omitempty"`
	VTotDFe     *TDec1302                      `xml:"http://www.portalfiscal.inf.br/nfcom vTotDFe,omitempty" json:"vTotDFe,omitempty"`
}

type TNFComInfNFComGFidelidade struct {
	QtdSaldoPts  string `xml:"http://www.portalfiscal.inf.br/nfcom qtdSaldoPts" json:"qtdSaldoPts"`
	DRefSaldoPts TData  `xml:"http://www.portalfiscal.inf.br/nfcom dRefSaldoPts" json:"dRefSaldoPts"`
	QtdPtsResg   string `xml:"http://www.portalfiscal.inf.br/nfcom qtdPtsResg" json:"qtdPtsResg"`
	DRefResgPts  TData  `xml:"http://www.portalfiscal.inf.br/nfcom dRefResgPts" json:"dRefResgPts"`
}

type TNFComInfNFComGFatGPIX struct {
	URLQRCodePIX string `xml:"http://www.portalfiscal.inf.br/nfcom urlQRCodePIX" json:"urlQRCodePIX"`
}

type TNFComInfNFComGFat struct {
	CompetFat    TCompet                 `xml:"http://www.portalfiscal.inf.br/nfcom CompetFat" json:"CompetFat"`
	DVencFat     TData                   `xml:"http://www.portalfiscal.inf.br/nfcom dVencFat" json:"dVencFat"`
	DPerUsoIni   *TData                  `xml:"http://www.portalfiscal.inf.br/nfcom dPerUsoIni,omitempty" json:"dPerUsoIni,omitempty"`
	DPerUsoFim   *TData                  `xml:"http://www.portalfiscal.inf.br/nfcom dPerUsoFim,omitempty" json:"dPerUsoFim,omitempty"`
	CodBarras    string                  `xml:"http://www.portalfiscal.inf.br/nfcom codBarras" json:"codBarras"`
	CodDebAuto   *string                 `xml:"http://www.portalfiscal.inf.br/nfcom codDebAuto,omitempty" json:"codDebAuto,omitempty"`
	CodBanco     *string                 `xml:"http://www.portalfiscal.inf.br/nfcom codBanco,omitempty" json:"codBanco,omitempty"`
	CodAgencia   *string                 `xml:"http://www.portalfiscal.inf.br/nfcom codAgencia,omitempty" json:"codAgencia,omitempty"`
	EnderCorresp *TEndeEmi               `xml:"http://www.portalfiscal.inf.br/nfcom enderCorresp,omitempty" json:"enderCorresp,omitempty"`
	GPIX         *TNFComInfNFComGFatGPIX `xml:"http://www.portalfiscal.inf.br/nfcom gPIX,omitempty" json:"gPIX,omitempty"`
}

type TNFComInfNFComGFatCentral struct {
	CNPJ TCNPJ      `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ" json:"CNPJ"`
	CUF  TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/nfcom cUF" json:"cUF"`
}

type TNFComInfNFComAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/nfcom CPF,omitempty" json:"CPF,omitempty"`
}

type TNFComInfNFComInfAdic struct {
	InfAdFisco *string  `xml:"http://www.portalfiscal.inf.br/nfcom infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     []string `xml:"http://www.portalfiscal.inf.br/nfcom infCpl,omitempty" json:"infCpl,omitempty"`
}

type TNFComInfNFCom struct {
	Versao      TVerNFCom                  `xml:"versao,attr" json:"versao"`
	ID          string                     `xml:"Id,attr" json:"Id"`
	Ide         TNFComInfNFComIde          `xml:"http://www.portalfiscal.inf.br/nfcom ide" json:"ide"`
	Emit        TNFComInfNFComEmit         `xml:"http://www.portalfiscal.inf.br/nfcom emit" json:"emit"`
	Dest        TNFComInfNFComDest         `xml:"http://www.portalfiscal.inf.br/nfcom dest" json:"dest"`
	Assinante   TNFComInfNFComAssinante    `xml:"http://www.portalfiscal.inf.br/nfcom assinante" json:"assinante"`
	GSub        *TNFComInfNFComGSub        `xml:"http://www.portalfiscal.inf.br/nfcom gSub,omitempty" json:"gSub,omitempty"`
	GCofat      *TNFComInfNFComGCofat      `xml:"http://www.portalfiscal.inf.br/nfcom gCofat,omitempty" json:"gCofat,omitempty"`
	Det         []TNFComInfNFComDet        `xml:"http://www.portalfiscal.inf.br/nfcom det,omitempty" json:"det,omitempty"`
	PgtoVinc    *TNFComInfNFComPgtoVinc    `xml:"http://www.portalfiscal.inf.br/nfcom pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	Total       TNFComInfNFComTotal        `xml:"http://www.portalfiscal.inf.br/nfcom total" json:"total"`
	GFidelidade *TNFComInfNFComGFidelidade `xml:"http://www.portalfiscal.inf.br/nfcom gFidelidade,omitempty" json:"gFidelidade,omitempty"`
	GFat        *TNFComInfNFComGFat        `xml:"http://www.portalfiscal.inf.br/nfcom gFat,omitempty" json:"gFat,omitempty"`
	GFatCentral *TNFComInfNFComGFatCentral `xml:"http://www.portalfiscal.inf.br/nfcom gFatCentral,omitempty" json:"gFatCentral,omitempty"`
	AutXML      []TNFComInfNFComAutXML     `xml:"http://www.portalfiscal.inf.br/nfcom autXML,omitempty" json:"autXML,omitempty"`
	InfAdic     *TNFComInfNFComInfAdic     `xml:"http://www.portalfiscal.inf.br/nfcom infAdic,omitempty" json:"infAdic,omitempty"`
	GRespTec    *TRespTec                  `xml:"http://www.portalfiscal.inf.br/nfcom gRespTec,omitempty" json:"gRespTec,omitempty"`
}

type TNFComInfNFComSupl struct {
	QRCodNFCom string `xml:"http://www.portalfiscal.inf.br/nfcom qrCodNFCom" json:"qrCodNFCom"`
}

type TNFCom struct {
	InfNFCom     TNFComInfNFCom     `xml:"http://www.portalfiscal.inf.br/nfcom infNFCom" json:"infNFCom"`
	InfNFComSupl TNFComInfNFComSupl `xml:"http://www.portalfiscal.inf.br/nfcom infNFComSupl" json:"infNFComSupl"`
	Signature    SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TNSU string

type TOperCompraGov string

type TPagAnt string

type TPagRef struct {
	RefDFe []TChDFeRTC `xml:"http://www.portalfiscal.inf.br/nfcom refDFe,omitempty" json:"refDFe,omitempty"`
}

type TPagamentoRTC struct {
	NPag        string       `xml:"nPag,attr" json:"nPag"`
	IDTransacao string       `xml:"idTransacao,attr" json:"idTransacao"`
	TpMeioPgto  string       `xml:"http://www.portalfiscal.inf.br/nfcom tpMeioPgto" json:"tpMeioPgto"`
	CNPJReceb   TCNPJRTC     `xml:"http://www.portalfiscal.inf.br/nfcom CNPJReceb" json:"CNPJReceb"`
	CNPJBasePSP TCNPJBaseRTC `xml:"http://www.portalfiscal.inf.br/nfcom CNPJBasePSP" json:"CNPJBasePSP"`
}

type TProcEmi string

type TProcesso string

type TProt string

type TProtNFComInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/nfcom tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/nfcom verAplic" json:"verAplic"`
	ChNFCom  TChDFe           `xml:"http://www.portalfiscal.inf.br/nfcom chNFCom" json:"chNFCom"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/nfcom dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/nfcom nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/nfcom digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/nfcom cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/nfcom xMotivo" json:"xMotivo"`
}

type TProtNFComInfFisco struct {
	CMsg string  `xml:"http://www.portalfiscal.inf.br/nfcom cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/nfcom xMsg" json:"xMsg"`
}

type TProtNFCom struct {
	Versao    TVerNFCom           `xml:"versao,attr" json:"versao"`
	InfProt   TProtNFComInfProt   `xml:"http://www.portalfiscal.inf.br/nfcom infProt" json:"infProt"`
	InfFisco  *TProtNFComInfFisco `xml:"http://www.portalfiscal.inf.br/nfcom infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRBSN string

type TRec string

type TRed struct {
	PRedAliq  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqEfet" json:"pAliqEfet"`
}

type TRespTec struct {
	CNPJ     TCNPJ   `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ" json:"CNPJ"`
	XContato string  `xml:"http://www.portalfiscal.inf.br/nfcom xContato" json:"xContato"`
	Email    TEmail  `xml:"http://www.portalfiscal.inf.br/nfcom email" json:"email"`
	Fone     string  `xml:"http://www.portalfiscal.inf.br/nfcom fone" json:"fone"`
	IDCSRT   *string `xml:"http://www.portalfiscal.inf.br/nfcom idCSRT,omitempty" json:"idCSRT,omitempty"`
	HashCSRT *string `xml:"http://www.portalfiscal.inf.br/nfcom hashCSRT,omitempty" json:"hashCSRT,omitempty"`
}

type TRessarcimento string

type TRetEnviNFComInfRec struct {
	NRec     TRec         `xml:"http://www.portalfiscal.inf.br/nfcom nRec" json:"nRec"`
	DhRecbto TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfcom dhRecbto" json:"dhRecbto"`
	TMed     string       `xml:"http://www.portalfiscal.inf.br/nfcom tMed" json:"tMed"`
}

type TRetEnviNFCom struct {
	Versao   TVerNFCom            `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb                 `xml:"http://www.portalfiscal.inf.br/nfcom tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE           `xml:"http://www.portalfiscal.inf.br/nfcom cUF" json:"cUF"`
	VerAplic TVerAplic            `xml:"http://www.portalfiscal.inf.br/nfcom verAplic" json:"verAplic"`
	CStat    TStat                `xml:"http://www.portalfiscal.inf.br/nfcom cStat" json:"cStat"`
	XMotivo  TMotivo              `xml:"http://www.portalfiscal.inf.br/nfcom xMotivo" json:"xMotivo"`
	InfRec   *TRetEnviNFComInfRec `xml:"http://www.portalfiscal.inf.br/nfcom infRec,omitempty" json:"infRec,omitempty"`
}

type TRetNFCom struct {
	Versao    TVerNFCom   `xml:"versao,attr" json:"versao"`
	TpAmb     TAmb        `xml:"http://www.portalfiscal.inf.br/nfcom tpAmb" json:"tpAmb"`
	CUF       TCodUFIBGE  `xml:"http://www.portalfiscal.inf.br/nfcom cUF" json:"cUF"`
	VerAplic  TVerAplic   `xml:"http://www.portalfiscal.inf.br/nfcom verAplic" json:"verAplic"`
	CStat     TStat       `xml:"http://www.portalfiscal.inf.br/nfcom cStat" json:"cStat"`
	XMotivo   TMotivo     `xml:"http://www.portalfiscal.inf.br/nfcom xMotivo" json:"xMotivo"`
	ProtNFCom *TProtNFCom `xml:"http://www.portalfiscal.inf.br/nfcom protNFCom,omitempty" json:"protNFCom,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TStringRTC string

type TTarifa string

type TTime string

type TTotalSN struct {
	VRBSNTot          TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfcom vRBSNTot" json:"vRBSNTot"`
	VIBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSSNTot,omitempty" json:"vIBSSNTot,omitempty"`
	VIBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSSNtotPendSusp,omitempty" json:"vIBSSNtotPendSusp,omitempty"`
	VCBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSSNTot,omitempty" json:"vCBSSNTot,omitempty"`
	VCBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSSNtotPendSusp,omitempty" json:"vCBSSNtotPendSusp,omitempty"`
}

type TTpCredPresIBSZFM string

type TTransfCred struct {
	VIBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBS" json:"vIBS"`
	VCBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBS" json:"vCBS"`
}

type TTribBPe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfcom cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfcom indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCTe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfcom cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfcom indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCompraGov struct {
	PAliqIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vTribCBS" json:"vTribCBS"`
}

type TTribItemSN struct {
	NItem        string         `xml:"nItem,attr" json:"nItem"`
	VRBSNItem    TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfcom vRBSNItem" json:"vRBSNItem"`
	TpRBSN       TRBSN          `xml:"http://www.portalfiscal.inf.br/nfcom tpRBSN" json:"tpRBSN"`
	PIBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pIBSSN,omitempty" json:"pIBSSN,omitempty"`
	VIBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSSN,omitempty" json:"vIBSSN,omitempty"`
	PCBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pCBSSN,omitempty" json:"pCBSSN,omitempty"`
	VCBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vCBSSN,omitempty" json:"vCBSSN,omitempty"`
	VIBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vIBSPendSusp,omitempty" json:"vIBSPendSusp,omitempty"`
	VCBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vCBSPendSusp,omitempty" json:"vCBSPendSusp,omitempty"`
}

type TTribNF3e struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfcom cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfcom indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFAg struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfcom cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfcom indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFCe struct {
	CST         TCST        `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	CClassTrib  TcClassTrib `xml:"http://www.portalfiscal.inf.br/nfcom cClassTrib" json:"cClassTrib"`
	IndDoacao   *TIndDoacao `xml:"http://www.portalfiscal.inf.br/nfcom indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS     *TCIBSNFe   `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono *TMonofasia `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
}

type TTribNFCom struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfcom cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfcom indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFGas struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfcom cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfcom indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFe struct {
	CST             TCST             `xml:"http://www.portalfiscal.inf.br/nfcom CST" json:"CST"`
	CClassTrib      TcClassTrib      `xml:"http://www.portalfiscal.inf.br/nfcom cClassTrib" json:"cClassTrib"`
	IndDoacao       *TIndDoacao      `xml:"http://www.portalfiscal.inf.br/nfcom indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS         *TCIBSNFe        `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono     *TMonofasia      `xml:"http://www.portalfiscal.inf.br/nfcom gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
	GTransfCred     *TTransfCred     `xml:"http://www.portalfiscal.inf.br/nfcom gTransfCred,omitempty" json:"gTransfCred,omitempty"`
	GAjusteCompet   *TAjusteCompet   `xml:"http://www.portalfiscal.inf.br/nfcom gAjusteCompet,omitempty" json:"gAjusteCompet,omitempty"`
	GEstornoCred    *TEstornoCred    `xml:"http://www.portalfiscal.inf.br/nfcom gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
	GCredPresOper   *TCredPresOper   `xml:"http://www.portalfiscal.inf.br/nfcom gCredPresOper,omitempty" json:"gCredPresOper,omitempty"`
	GCredPresIBSZFM *TCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nfcom gCredPresIBSZFM,omitempty" json:"gCredPresIBSZFM,omitempty"`
}

type TTribRegular struct {
	CSTReg             TCST          `xml:"http://www.portalfiscal.inf.br/nfcom CSTReg" json:"CSTReg"`
	CClassTribReg      TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfcom cClassTribReg" json:"cClassTribReg"`
	PAliqEfetRegIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqEfetRegIBSUF" json:"pAliqEfetRegIBSUF"`
	VTribRegIBSUF      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vTribRegIBSUF" json:"vTribRegIBSUF"`
	PAliqEfetRegIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqEfetRegIBSMun" json:"pAliqEfetRegIBSMun"`
	VTribRegIBSMun     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vTribRegIBSMun" json:"vTribRegIBSMun"`
	PAliqEfetRegCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS        TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfcom vTribRegCBS" json:"vTribRegCBS"`
}

type TUMedItem string

type TUF string

type TUFSemEX string

type TUtilizacao string

type TVerAplic string

type TVerNFCom string

type Tano string

type TcClassTrib string

type TcCredPres string

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

type GYearMonth string

type Integer string

type String string

// RetNFCom represents the global element {http://www.portalfiscal.inf.br/nfcom}retNFCom.
type RetNFCom struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfcom retNFCom" json:"-"`
	TRetNFCom
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
