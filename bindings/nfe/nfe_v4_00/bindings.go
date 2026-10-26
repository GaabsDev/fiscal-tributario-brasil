// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package nfe_v4_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TALCZFMCBS struct {
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vTribRegCBS" json:"vTribRegCBS"`
}

type TALCZFMCBSNFe struct {
	TpALCZFMCBS     string        `xml:"http://www.portalfiscal.inf.br/nfe tpALCZFMCBS" json:"tpALCZFMCBS"`
	NProcSuframa    *string       `xml:"http://www.portalfiscal.inf.br/nfe nProcSuframa,omitempty" json:"nProcSuframa,omitempty"`
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vTribRegCBS" json:"vTribRegCBS"`
}

type TAjusteCompet struct {
	CompetApur TCompetApur `xml:"http://www.portalfiscal.inf.br/nfe competApur" json:"competApur"`
	VIBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBS" json:"vIBS"`
	VCBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBS" json:"vCBS"`
}

type TAmb string

type TCIBSGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfe gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSUF" json:"vIBSUF"`
}

type TCIBSGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfe gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSMun" json:"vIBSMun"`
}

type TCIBSGCBS struct {
	PCBS       TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pCBS" json:"pCBS"`
	GDif       *TDif         `xml:"http://www.portalfiscal.inf.br/nfe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed         `xml:"http://www.portalfiscal.inf.br/nfe gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBS   `xml:"http://www.portalfiscal.inf.br/nfe gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vCBS" json:"vCBS"`
}

type TCIBS struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfe vBC" json:"vBC"`
	GIBSUF         TCIBSGIBSUF     `xml:"http://www.portalfiscal.inf.br/nfe gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSGIBSMun    `xml:"http://www.portalfiscal.inf.br/nfe gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfe vIBS" json:"vIBS"`
	GCBS           TCIBSGCBS       `xml:"http://www.portalfiscal.inf.br/nfe gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nfe gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfe gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCIBSNFeGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfe gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSUF" json:"vIBSUF"`
}

type TCIBSNFeGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfe gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSMun" json:"vIBSMun"`
}

type TCIBSNFeGCBS struct {
	PCBS       TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfe pCBS" json:"pCBS"`
	GDif       *TDif          `xml:"http://www.portalfiscal.inf.br/nfe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib      `xml:"http://www.portalfiscal.inf.br/nfe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed          `xml:"http://www.portalfiscal.inf.br/nfe gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBSNFe `xml:"http://www.portalfiscal.inf.br/nfe gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfe vCBS" json:"vCBS"`
}

type TCIBSNFe struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfe vBC" json:"vBC"`
	GIBSUF         TCIBSNFeGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfe gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSNFeGIBSMun `xml:"http://www.portalfiscal.inf.br/nfe gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfe vIBS" json:"vIBS"`
	GCBS           TCIBSNFeGCBS    `xml:"http://www.portalfiscal.inf.br/nfe gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nfe gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfe gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCListServ string

type TCOrgaoIBGE string

type TCST string

type TChDFeRTC string

type TChNFe string

type TCNPJ string

type TCNPJBaseRTC string

type TCNPJOpc string

type TCNPJRTC string

type TCNPJVar string

type TCodMunIBGE string

type TCodUFIBGE string

type TCompetApur string

type TCompraGov struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nfe tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfe pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nfe tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nfe refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCompraGovReduzido struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nfe tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfe pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nfe tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nfe refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TConsReciNFe struct {
	Versao TVerNFe `xml:"versao,attr" json:"versao"`
	TpAmb  TAmb    `xml:"http://www.portalfiscal.inf.br/nfe tpAmb" json:"tpAmb"`
	NRec   TRec    `xml:"http://www.portalfiscal.inf.br/nfe nRec" json:"nRec"`
}

type TCPF string

type TCPFVar string

type TCredPres struct {
	PCredPres        TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pCredPres" json:"pCredPres"`
	VCredPres        *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfe vCredPres,omitempty" json:"vCredPres,omitempty"`
	VCredPresCondSus *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfe vCredPresCondSus,omitempty" json:"vCredPresCondSus,omitempty"`
}

type TCredPresIBSZFM struct {
	CompetApur       TCompetApur       `xml:"http://www.portalfiscal.inf.br/nfe competApur" json:"competApur"`
	TpCredPresIBSZFM TTpCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nfe tpCredPresIBSZFM" json:"tpCredPresIBSZFM"`
	VCredPresIBSZFM  TDec1302RTC       `xml:"http://www.portalfiscal.inf.br/nfe vCredPresIBSZFM" json:"vCredPresIBSZFM"`
}

type TCredPresOper struct {
	VBCCredPres  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vBCCredPres" json:"vBCCredPres"`
	CCredPres    TcCredPres  `xml:"http://www.portalfiscal.inf.br/nfe cCredPres" json:"cCredPres"`
	GIBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nfe gIBSCredPres,omitempty" json:"gIBSCredPres,omitempty"`
	GCBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nfe gCBSCredPres,omitempty" json:"gCBSCredPres,omitempty"`
}

type TData string

type TDateTimeUTC string

type TDec1104RTC string

type TDec1302RTC string

type TDec0104v string

type TDec0204v string

type TDec0302Max100 string

type TDec030204RTC string

type TDec0302a04 string

type TDec0302a04Max100 string

type TDec0302a04Opc string

type TDec0304Max100 string

type TDec03v00a04Max100Opc string

type TDec0803v string

type TDec1104 string

type TDec1104OpRTC string

type TDec1104Opc string

type TDec1104v string

type TDec1110v string

type TDec1203 string

type TDec1204 string

type TDec1204Opc string

type TDec1204temperatura string

type TDec1204v string

type TDec1302 string

type TDec1302Opc string

type TDevTrib struct {
	PDevTrib *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pDevTrib,omitempty" json:"pDevTrib,omitempty"`
	VDevTrib TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfe vDevTrib" json:"vDevTrib"`
}

type TDif struct {
	PDif TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pDif" json:"pDif"`
	VDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vDif" json:"vDif"`
}

type TEnderEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfe xMun" json:"xMun"`
	UF      TUFEmi      `xml:"http://www.portalfiscal.inf.br/nfe UF" json:"UF"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/nfe CEP" json:"CEP"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/nfe cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/nfe xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfe fone,omitempty" json:"fone,omitempty"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfe xMun" json:"xMun"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/nfe UF" json:"UF"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/nfe CEP,omitempty" json:"CEP,omitempty"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/nfe cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/nfe xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfe fone,omitempty" json:"fone,omitempty"`
}

type TEnteGov string

type TEnviNFe struct {
	Versao  TVerNFe `xml:"versao,attr" json:"versao"`
	IDLote  TIDLote `xml:"http://www.portalfiscal.inf.br/nfe idLote" json:"idLote"`
	IndSinc string  `xml:"http://www.portalfiscal.inf.br/nfe indSinc" json:"indSinc"`
	NFe     []TNFe  `xml:"http://www.portalfiscal.inf.br/nfe NFe,omitempty" json:"NFe,omitempty"`
}

type TEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSEstCred" json:"vCBSEstCred"`
}

type TFinNFe string

type TGuid string

type TIBSCBSMonoTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSUF" json:"vIBSUF"`
}

type TIBSCBSMonoTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSMun" json:"vIBSMun"`
}

type TIBSCBSMonoTotGIBS struct {
	GIBSUF           TIBSCBSMonoTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfe gIBSUF" json:"gIBSUF"`
	GIBSMun          TIBSCBSMonoTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nfe gIBSMun" json:"gIBSMun"`
	VIBS             TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfe vIBS" json:"vIBS"`
	VCredPres        TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfe vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfe vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGCBS struct {
	VDif             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDif" json:"vDif"`
	VDevTrib         TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDevTrib" json:"vDevTrib"`
	VCBS             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBS" json:"vCBS"`
	VCredPres        TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGMono struct {
	VIBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSMono" json:"vIBSMono"`
	VCBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSMono" json:"vCBSMono"`
	VIBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSMonoReten" json:"vIBSMonoReten"`
	VCBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSMonoReten" json:"vCBSMonoReten"`
	VIBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSMonoRet" json:"vIBSMonoRet"`
	VCBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSMonoRet" json:"vCBSMonoRet"`
}

type TIBSCBSMonoTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSMonoTot struct {
	VBCIBSCBS    TDec1302RTC                 `xml:"http://www.portalfiscal.inf.br/nfe vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         *TIBSCBSMonoTotGIBS         `xml:"http://www.portalfiscal.inf.br/nfe gIBS,omitempty" json:"gIBS,omitempty"`
	GCBS         *TIBSCBSMonoTotGCBS         `xml:"http://www.portalfiscal.inf.br/nfe gCBS,omitempty" json:"gCBS,omitempty"`
	GMono        *TIBSCBSMonoTotGMono        `xml:"http://www.portalfiscal.inf.br/nfe gMono,omitempty" json:"gMono,omitempty"`
	GEstornoCred *TIBSCBSMonoTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIBSCBSTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSUF" json:"vIBSUF"`
}

type TIBSCBSTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSMun" json:"vIBSMun"`
}

type TIBSCBSTotGIBS struct {
	GIBSUF  TIBSCBSTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfe gIBSUF" json:"gIBSUF"`
	GIBSMun TIBSCBSTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nfe gIBSMun" json:"gIBSMun"`
	VIBS    TDec1302RTC           `xml:"http://www.portalfiscal.inf.br/nfe vIBS" json:"vIBS"`
}

type TIBSCBSTotGCBS struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vDevTrib" json:"vDevTrib"`
	VCBS     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBS" json:"vCBS"`
}

type TIBSCBSTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSTot struct {
	VBCIBSCBS    TDec1302RTC             `xml:"http://www.portalfiscal.inf.br/nfe vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         TIBSCBSTotGIBS          `xml:"http://www.portalfiscal.inf.br/nfe gIBS" json:"gIBS"`
	GCBS         TIBSCBSTotGCBS          `xml:"http://www.portalfiscal.inf.br/nfe gCBS" json:"gCBS"`
	GEstornoCred *TIBSCBSTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIS struct {
	CSTIS        TCST           `xml:"http://www.portalfiscal.inf.br/nfe CSTIS" json:"CSTIS"`
	CClassTribIS TcClassTrib    `xml:"http://www.portalfiscal.inf.br/nfe cClassTribIS" json:"cClassTribIS"`
	VBCIS        *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vBCIS,omitempty" json:"vBCIS,omitempty"`
	PIS          *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pIS,omitempty" json:"pIS,omitempty"`
	AdRemIS      *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe adRemIS,omitempty" json:"adRemIS,omitempty"`
	UTrib        *string        `xml:"http://www.portalfiscal.inf.br/nfe uTrib,omitempty" json:"uTrib,omitempty"`
	QTrib        *TDec1104OpRTC `xml:"http://www.portalfiscal.inf.br/nfe qTrib,omitempty" json:"qTrib,omitempty"`
	VIS          *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIS,omitempty" json:"vIS,omitempty"`
}

type TISTot struct {
	VIS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIS" json:"vIS"`
}

type TIDLote string

type TIe string

type TIeDest string

type TIeDestNaoIsento string

type TIeST string

type TIndDoacao string

type TInfRespTec struct {
	CNPJ     TCNPJOpc `xml:"http://www.portalfiscal.inf.br/nfe CNPJ" json:"CNPJ"`
	XContato string   `xml:"http://www.portalfiscal.inf.br/nfe xContato" json:"xContato"`
	Email    string   `xml:"http://www.portalfiscal.inf.br/nfe email" json:"email"`
	Fone     string   `xml:"http://www.portalfiscal.inf.br/nfe fone" json:"fone"`
	IDCSRT   *string  `xml:"http://www.portalfiscal.inf.br/nfe idCSRT,omitempty" json:"idCSRT,omitempty"`
	HashCSRT *string  `xml:"http://www.portalfiscal.inf.br/nfe hashCSRT,omitempty" json:"hashCSRT,omitempty"`
}

type TIPIIPITrib struct {
	CST   string       `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	VBC   *TDec1302    `xml:"http://www.portalfiscal.inf.br/nfe vBC,omitempty" json:"vBC,omitempty"`
	PIPI  *TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfe pIPI,omitempty" json:"pIPI,omitempty"`
	QUnid *TDec1204v   `xml:"http://www.portalfiscal.inf.br/nfe qUnid,omitempty" json:"qUnid,omitempty"`
	VUnid *TDec1104    `xml:"http://www.portalfiscal.inf.br/nfe vUnid,omitempty" json:"vUnid,omitempty"`
	VIPI  TDec1302     `xml:"http://www.portalfiscal.inf.br/nfe vIPI" json:"vIPI"`
}

type TIPIIPINT struct {
	CST string `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
}

type TIPI struct {
	CNPJProd *TCNPJ       `xml:"http://www.portalfiscal.inf.br/nfe CNPJProd,omitempty" json:"CNPJProd,omitempty"`
	CSelo    *string      `xml:"http://www.portalfiscal.inf.br/nfe cSelo,omitempty" json:"cSelo,omitempty"`
	QSelo    *string      `xml:"http://www.portalfiscal.inf.br/nfe qSelo,omitempty" json:"qSelo,omitempty"`
	CEnq     string       `xml:"http://www.portalfiscal.inf.br/nfe cEnq" json:"cEnq"`
	IPITrib  *TIPIIPITrib `xml:"http://www.portalfiscal.inf.br/nfe IPITrib,omitempty" json:"IPITrib,omitempty"`
	IPINT    *TIPIIPINT   `xml:"http://www.portalfiscal.inf.br/nfe IPINT,omitempty" json:"IPINT,omitempty"`
}

type TJust string

type TLocal struct {
	CNPJ    *TCNPJOpc   `xml:"http://www.portalfiscal.inf.br/nfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TCPF       `xml:"http://www.portalfiscal.inf.br/nfe CPF,omitempty" json:"CPF,omitempty"`
	XNome   *string     `xml:"http://www.portalfiscal.inf.br/nfe xNome,omitempty" json:"xNome,omitempty"`
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfe xMun" json:"xMun"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/nfe UF" json:"UF"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/nfe CEP,omitempty" json:"CEP,omitempty"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/nfe cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/nfe xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfe fone,omitempty" json:"fone,omitempty"`
	Email   *string     `xml:"http://www.portalfiscal.inf.br/nfe email,omitempty" json:"email,omitempty"`
	IE      *TIe        `xml:"http://www.portalfiscal.inf.br/nfe IE,omitempty" json:"IE,omitempty"`
}

type TMed string

type TMod string

type TMonofasiaGIBSMonoAdRemGMonoPadrao struct {
	QBCMono  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfe qBCMono" json:"qBCMono"`
	AdRemIBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe adRemIBS" json:"adRemIBS"`
	VIBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSMono" json:"vIBSMono"`
}

type TMonofasiaGIBSMonoAdRemGMonoReten struct {
	QBCMonoReten  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfe qBCMonoReten" json:"qBCMonoReten"`
	AdRemIBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe adRemIBSReten" json:"adRemIBSReten"`
	VIBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSMonoReten" json:"vIBSMonoReten"`
}

type TMonofasiaGIBSMonoAdRemGMonoRet struct {
	VIBSMonoRet TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSMonoRet" json:"vIBSMonoRet"`
}

type TMonofasiaGIBSMonoAdRemGpBioDiferenca struct {
	QBCBioComb    TDec1104RTC `xml:"http://www.portalfiscal.inf.br/nfe qBCBioComb" json:"qBCBioComb"`
	VIBSDiferenca TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSDiferenca" json:"vIBSDiferenca"`
}

type TMonofasiaGIBSMonoAdRem struct {
	GMonoPadrao    *TMonofasiaGIBSMonoAdRemGMonoPadrao    `xml:"http://www.portalfiscal.inf.br/nfe gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten     *TMonofasiaGIBSMonoAdRemGMonoReten     `xml:"http://www.portalfiscal.inf.br/nfe gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet       *TMonofasiaGIBSMonoAdRemGMonoRet       `xml:"http://www.portalfiscal.inf.br/nfe gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GpBioDiferenca *TMonofasiaGIBSMonoAdRemGpBioDiferenca `xml:"http://www.portalfiscal.inf.br/nfe gpBioDiferenca,omitempty" json:"gpBioDiferenca,omitempty"`
}

type TMonofasiaGIBSMonoAdValoremGMonoPadrao struct {
	VBCMono      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vBCMono" json:"vBCMono"`
	PAliqMonoUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqMonoUF" json:"pAliqMonoUF"`
	VIBSMonoUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSMonoUF" json:"vIBSMonoUF"`
	PAliqMonoMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqMonoMun" json:"pAliqMonoMun"`
	VIBSMonoMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSMonoMun" json:"vIBSMonoMun"`
	VIBSMono     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSMono" json:"vIBSMono"`
}

type TMonofasiaGIBSMonoAdValoremGMonoReten struct {
	VBCMonoReten   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vBCMonoReten" json:"vBCMonoReten"`
	PAliqMonoReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqMonoReten" json:"pAliqMonoReten"`
	VIBSMonoReten  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vIBSMonoReten" json:"vIBSMonoReten"`
}

type TMonofasiaGIBSMonoAdValoremGMonoRet struct {
	VIBSMonoRet TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSMonoRet" json:"vIBSMonoRet"`
}

type TMonofasiaGIBSMonoAdValoremGpBioDiferenca struct {
	QBCBioComb    TDec1104RTC `xml:"http://www.portalfiscal.inf.br/nfe qBCBioComb" json:"qBCBioComb"`
	VIBSDiferenca TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBSDiferenca" json:"vIBSDiferenca"`
}

type TMonofasiaGIBSMonoAdValorem struct {
	GMonoPadrao    *TMonofasiaGIBSMonoAdValoremGMonoPadrao    `xml:"http://www.portalfiscal.inf.br/nfe gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten     *TMonofasiaGIBSMonoAdValoremGMonoReten     `xml:"http://www.portalfiscal.inf.br/nfe gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet       *TMonofasiaGIBSMonoAdValoremGMonoRet       `xml:"http://www.portalfiscal.inf.br/nfe gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GpBioDiferenca *TMonofasiaGIBSMonoAdValoremGpBioDiferenca `xml:"http://www.portalfiscal.inf.br/nfe gpBioDiferenca,omitempty" json:"gpBioDiferenca,omitempty"`
}

type TMonofasiaGCBSMonoAdRemGMonoPadrao struct {
	QBCMono  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfe qBCMono" json:"qBCMono"`
	AdRemCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe adRemCBS" json:"adRemCBS"`
	VCBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vCBSMono" json:"vCBSMono"`
}

type TMonofasiaGCBSMonoAdRemGMonoReten struct {
	QBCMonoReten  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfe qBCMonoReten" json:"qBCMonoReten"`
	AdRemCBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe adRemCBSReten" json:"adRemCBSReten"`
	VCBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vCBSMonoReten" json:"vCBSMonoReten"`
}

type TMonofasiaGCBSMonoAdRemGMonoRet struct {
	VCBSMonoRet TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSMonoRet" json:"vCBSMonoRet"`
}

type TMonofasiaGCBSMonoAdRemGpBioDiferenca struct {
	QBCBioComb    TDec1104RTC `xml:"http://www.portalfiscal.inf.br/nfe qBCBioComb" json:"qBCBioComb"`
	VCBSDiferenca TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSDiferenca" json:"vCBSDiferenca"`
}

type TMonofasiaGCBSMonoAdRem struct {
	GMonoPadrao    *TMonofasiaGCBSMonoAdRemGMonoPadrao    `xml:"http://www.portalfiscal.inf.br/nfe gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten     *TMonofasiaGCBSMonoAdRemGMonoReten     `xml:"http://www.portalfiscal.inf.br/nfe gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet       *TMonofasiaGCBSMonoAdRemGMonoRet       `xml:"http://www.portalfiscal.inf.br/nfe gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GpBioDiferenca *TMonofasiaGCBSMonoAdRemGpBioDiferenca `xml:"http://www.portalfiscal.inf.br/nfe gpBioDiferenca,omitempty" json:"gpBioDiferenca,omitempty"`
}

type TMonofasiaGCBSMonoAdValoremGMonoPadrao struct {
	VBCMono      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vBCMono" json:"vBCMono"`
	PAliqMonoCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqMonoCBS" json:"pAliqMonoCBS"`
	VCBSMono     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vCBSMono" json:"vCBSMono"`
}

type TMonofasiaGCBSMonoAdValoremGMonoReten struct {
	VBCMonoReten   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vBCMonoReten" json:"vBCMonoReten"`
	PAliqMonoReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqMonoReten" json:"pAliqMonoReten"`
	VCBSMonoReten  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vCBSMonoReten" json:"vCBSMonoReten"`
}

type TMonofasiaGCBSMonoAdValoremGMonoRet struct {
	VCBSMonoRet TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSMonoRet" json:"vCBSMonoRet"`
}

type TMonofasiaGCBSMonoAdValoremGpBioDiferenca struct {
	QBCBioComb    TDec1104RTC `xml:"http://www.portalfiscal.inf.br/nfe qBCBioComb" json:"qBCBioComb"`
	VCBSDiferenca TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBSDiferenca" json:"vCBSDiferenca"`
}

type TMonofasiaGCBSMonoAdValorem struct {
	GMonoPadrao    *TMonofasiaGCBSMonoAdValoremGMonoPadrao    `xml:"http://www.portalfiscal.inf.br/nfe gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten     *TMonofasiaGCBSMonoAdValoremGMonoReten     `xml:"http://www.portalfiscal.inf.br/nfe gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet       *TMonofasiaGCBSMonoAdValoremGMonoRet       `xml:"http://www.portalfiscal.inf.br/nfe gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GpBioDiferenca *TMonofasiaGCBSMonoAdValoremGpBioDiferenca `xml:"http://www.portalfiscal.inf.br/nfe gpBioDiferenca,omitempty" json:"gpBioDiferenca,omitempty"`
}

type TMonofasia struct {
	GIBSMonoAdRem     *TMonofasiaGIBSMonoAdRem     `xml:"http://www.portalfiscal.inf.br/nfe gIBSMonoAdRem,omitempty" json:"gIBSMonoAdRem,omitempty"`
	GIBSMonoAdValorem *TMonofasiaGIBSMonoAdValorem `xml:"http://www.portalfiscal.inf.br/nfe gIBSMonoAdValorem,omitempty" json:"gIBSMonoAdValorem,omitempty"`
	GCBSMonoAdRem     *TMonofasiaGCBSMonoAdRem     `xml:"http://www.portalfiscal.inf.br/nfe gCBSMonoAdRem,omitempty" json:"gCBSMonoAdRem,omitempty"`
	GCBSMonoAdValorem *TMonofasiaGCBSMonoAdValorem `xml:"http://www.portalfiscal.inf.br/nfe gCBSMonoAdValorem,omitempty" json:"gCBSMonoAdValorem,omitempty"`
	VTotIBSMonoItem   TDec1302RTC                  `xml:"http://www.portalfiscal.inf.br/nfe vTotIBSMonoItem" json:"vTotIBSMonoItem"`
	VTotCBSMonoItem   TDec1302RTC                  `xml:"http://www.portalfiscal.inf.br/nfe vTotCBSMonoItem" json:"vTotCBSMonoItem"`
}

type TMotivo string

type TNF string

type TNFeInfNFeIdeNFrefRefNF struct {
	CUF   TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/nfe cUF" json:"cUF"`
	AAMM  string     `xml:"http://www.portalfiscal.inf.br/nfe AAMM" json:"AAMM"`
	CNPJ  TCNPJ      `xml:"http://www.portalfiscal.inf.br/nfe CNPJ" json:"CNPJ"`
	Mod   string     `xml:"http://www.portalfiscal.inf.br/nfe mod" json:"mod"`
	Serie TSerie     `xml:"http://www.portalfiscal.inf.br/nfe serie" json:"serie"`
	NNF   TNF        `xml:"http://www.portalfiscal.inf.br/nfe nNF" json:"nNF"`
}

type TNFeInfNFeIdeNFrefRefNFP struct {
	CUF   TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/nfe cUF" json:"cUF"`
	AAMM  string     `xml:"http://www.portalfiscal.inf.br/nfe AAMM" json:"AAMM"`
	CNPJ  *TCNPJ     `xml:"http://www.portalfiscal.inf.br/nfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *TCPF      `xml:"http://www.portalfiscal.inf.br/nfe CPF,omitempty" json:"CPF,omitempty"`
	IE    TIeDest    `xml:"http://www.portalfiscal.inf.br/nfe IE" json:"IE"`
	Mod   string     `xml:"http://www.portalfiscal.inf.br/nfe mod" json:"mod"`
	Serie TSerie     `xml:"http://www.portalfiscal.inf.br/nfe serie" json:"serie"`
	NNF   TNF        `xml:"http://www.portalfiscal.inf.br/nfe nNF" json:"nNF"`
}

type TNFeInfNFeIdeNFrefRefECF struct {
	Mod  string `xml:"http://www.portalfiscal.inf.br/nfe mod" json:"mod"`
	NECF string `xml:"http://www.portalfiscal.inf.br/nfe nECF" json:"nECF"`
	NCOO string `xml:"http://www.portalfiscal.inf.br/nfe nCOO" json:"nCOO"`
}

type TNFeInfNFeIdeNFref struct {
	RefNFe    *TChNFe                   `xml:"http://www.portalfiscal.inf.br/nfe refNFe,omitempty" json:"refNFe,omitempty"`
	RefNFeSig *TChNFe                   `xml:"http://www.portalfiscal.inf.br/nfe refNFeSig,omitempty" json:"refNFeSig,omitempty"`
	RefNF     *TNFeInfNFeIdeNFrefRefNF  `xml:"http://www.portalfiscal.inf.br/nfe refNF,omitempty" json:"refNF,omitempty"`
	RefNFP    *TNFeInfNFeIdeNFrefRefNFP `xml:"http://www.portalfiscal.inf.br/nfe refNFP,omitempty" json:"refNFP,omitempty"`
	RefCTe    *TChNFe                   `xml:"http://www.portalfiscal.inf.br/nfe refCTe,omitempty" json:"refCTe,omitempty"`
	RefECF    *TNFeInfNFeIdeNFrefRefECF `xml:"http://www.portalfiscal.inf.br/nfe refECF,omitempty" json:"refECF,omitempty"`
}

type TNFeInfNFeIdeGPagAntecipado struct {
	RefNFe []TChNFe `xml:"http://www.portalfiscal.inf.br/nfe refNFe,omitempty" json:"refNFe,omitempty"`
}

type TNFeInfNFeIde struct {
	CUF            TCodUFIBGE                   `xml:"http://www.portalfiscal.inf.br/nfe cUF" json:"cUF"`
	CNF            string                       `xml:"http://www.portalfiscal.inf.br/nfe cNF" json:"cNF"`
	NatOp          string                       `xml:"http://www.portalfiscal.inf.br/nfe natOp" json:"natOp"`
	Mod            TMod                         `xml:"http://www.portalfiscal.inf.br/nfe mod" json:"mod"`
	Serie          TSerie                       `xml:"http://www.portalfiscal.inf.br/nfe serie" json:"serie"`
	NNF            TNF                          `xml:"http://www.portalfiscal.inf.br/nfe nNF" json:"nNF"`
	DhEmi          TDateTimeUTC                 `xml:"http://www.portalfiscal.inf.br/nfe dhEmi" json:"dhEmi"`
	DhSaiEnt       *TDateTimeUTC                `xml:"http://www.portalfiscal.inf.br/nfe dhSaiEnt,omitempty" json:"dhSaiEnt,omitempty"`
	DPrevEntrega   *TData                       `xml:"http://www.portalfiscal.inf.br/nfe dPrevEntrega,omitempty" json:"dPrevEntrega,omitempty"`
	TpNF           string                       `xml:"http://www.portalfiscal.inf.br/nfe tpNF" json:"tpNF"`
	IDDest         string                       `xml:"http://www.portalfiscal.inf.br/nfe idDest" json:"idDest"`
	CMunFG         TCodMunIBGE                  `xml:"http://www.portalfiscal.inf.br/nfe cMunFG" json:"cMunFG"`
	CMunFGIBS      *TCodMunIBGE                 `xml:"http://www.portalfiscal.inf.br/nfe cMunFGIBS,omitempty" json:"cMunFGIBS,omitempty"`
	TpImp          string                       `xml:"http://www.portalfiscal.inf.br/nfe tpImp" json:"tpImp"`
	TpEmis         string                       `xml:"http://www.portalfiscal.inf.br/nfe tpEmis" json:"tpEmis"`
	CDV            string                       `xml:"http://www.portalfiscal.inf.br/nfe cDV" json:"cDV"`
	TpAmb          TAmb                         `xml:"http://www.portalfiscal.inf.br/nfe tpAmb" json:"tpAmb"`
	FinNFe         TFinNFe                      `xml:"http://www.portalfiscal.inf.br/nfe finNFe" json:"finNFe"`
	TpNFDebito     *TTpNFDebito                 `xml:"http://www.portalfiscal.inf.br/nfe tpNFDebito,omitempty" json:"tpNFDebito,omitempty"`
	TpNFCredito    *TTpNFCredito                `xml:"http://www.portalfiscal.inf.br/nfe tpNFCredito,omitempty" json:"tpNFCredito,omitempty"`
	IndFinal       string                       `xml:"http://www.portalfiscal.inf.br/nfe indFinal" json:"indFinal"`
	IndPres        string                       `xml:"http://www.portalfiscal.inf.br/nfe indPres" json:"indPres"`
	IndIntermed    *string                      `xml:"http://www.portalfiscal.inf.br/nfe indIntermed,omitempty" json:"indIntermed,omitempty"`
	CIndOp         *string                      `xml:"http://www.portalfiscal.inf.br/nfe cIndOp,omitempty" json:"cIndOp,omitempty"`
	ProcEmi        TProcEmi                     `xml:"http://www.portalfiscal.inf.br/nfe procEmi" json:"procEmi"`
	VerProc        string                       `xml:"http://www.portalfiscal.inf.br/nfe verProc" json:"verProc"`
	DhCont         *TDateTimeUTC                `xml:"http://www.portalfiscal.inf.br/nfe dhCont,omitempty" json:"dhCont,omitempty"`
	XJust          *string                      `xml:"http://www.portalfiscal.inf.br/nfe xJust,omitempty" json:"xJust,omitempty"`
	NFref          []TNFeInfNFeIdeNFref         `xml:"http://www.portalfiscal.inf.br/nfe NFref,omitempty" json:"NFref,omitempty"`
	GCompraGov     *TCompraGov                  `xml:"http://www.portalfiscal.inf.br/nfe gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	GPagAntecipado *TNFeInfNFeIdeGPagAntecipado `xml:"http://www.portalfiscal.inf.br/nfe gPagAntecipado,omitempty" json:"gPagAntecipado,omitempty"`
}

type TNFeInfNFeEmit struct {
	CNPJ      *TCNPJ    `xml:"http://www.portalfiscal.inf.br/nfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/nfe CPF,omitempty" json:"CPF,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/nfe xNome" json:"xNome"`
	XFant     *string   `xml:"http://www.portalfiscal.inf.br/nfe xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEnderEmi `xml:"http://www.portalfiscal.inf.br/nfe enderEmit" json:"enderEmit"`
	IE        *TIe      `xml:"http://www.portalfiscal.inf.br/nfe IE,omitempty" json:"IE,omitempty"`
	IEST      *TIeST    `xml:"http://www.portalfiscal.inf.br/nfe IEST,omitempty" json:"IEST,omitempty"`
	IM        *string   `xml:"http://www.portalfiscal.inf.br/nfe IM,omitempty" json:"IM,omitempty"`
	CNAE      *string   `xml:"http://www.portalfiscal.inf.br/nfe CNAE,omitempty" json:"CNAE,omitempty"`
	CRT       string    `xml:"http://www.portalfiscal.inf.br/nfe CRT" json:"CRT"`
	ISUFEmit  *string   `xml:"http://www.portalfiscal.inf.br/nfe ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TNFeInfNFeAvulsa struct {
	CNPJ    TCNPJ     `xml:"http://www.portalfiscal.inf.br/nfe CNPJ" json:"CNPJ"`
	XOrgao  string    `xml:"http://www.portalfiscal.inf.br/nfe xOrgao" json:"xOrgao"`
	Matr    string    `xml:"http://www.portalfiscal.inf.br/nfe matr" json:"matr"`
	XAgente string    `xml:"http://www.portalfiscal.inf.br/nfe xAgente" json:"xAgente"`
	Fone    *string   `xml:"http://www.portalfiscal.inf.br/nfe fone,omitempty" json:"fone,omitempty"`
	UF      TUFEmi    `xml:"http://www.portalfiscal.inf.br/nfe UF" json:"UF"`
	NDAR    *string   `xml:"http://www.portalfiscal.inf.br/nfe nDAR,omitempty" json:"nDAR,omitempty"`
	DEmi    *TData    `xml:"http://www.portalfiscal.inf.br/nfe dEmi,omitempty" json:"dEmi,omitempty"`
	VDAR    *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vDAR,omitempty" json:"vDAR,omitempty"`
	RepEmi  string    `xml:"http://www.portalfiscal.inf.br/nfe repEmi" json:"repEmi"`
	DPag    *TData    `xml:"http://www.portalfiscal.inf.br/nfe dPag,omitempty" json:"dPag,omitempty"`
}

type TNFeInfNFeDest struct {
	CNPJ          *TCNPJ            `xml:"http://www.portalfiscal.inf.br/nfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF           *TCPF             `xml:"http://www.portalfiscal.inf.br/nfe CPF,omitempty" json:"CPF,omitempty"`
	IDEstrangeiro *string           `xml:"http://www.portalfiscal.inf.br/nfe idEstrangeiro,omitempty" json:"idEstrangeiro,omitempty"`
	XNome         *string           `xml:"http://www.portalfiscal.inf.br/nfe xNome,omitempty" json:"xNome,omitempty"`
	EnderDest     *TEndereco        `xml:"http://www.portalfiscal.inf.br/nfe enderDest,omitempty" json:"enderDest,omitempty"`
	IndIEDest     string            `xml:"http://www.portalfiscal.inf.br/nfe indIEDest" json:"indIEDest"`
	IE            *TIeDestNaoIsento `xml:"http://www.portalfiscal.inf.br/nfe IE,omitempty" json:"IE,omitempty"`
	ISUF          *string           `xml:"http://www.portalfiscal.inf.br/nfe ISUF,omitempty" json:"ISUF,omitempty"`
	IM            *string           `xml:"http://www.portalfiscal.inf.br/nfe IM,omitempty" json:"IM,omitempty"`
	Email         *string           `xml:"http://www.portalfiscal.inf.br/nfe email,omitempty" json:"email,omitempty"`
}

type TNFeInfNFeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/nfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/nfe CPF,omitempty" json:"CPF,omitempty"`
}

type TNFeInfNFeDetProdGCred struct {
	CCredPresumido string      `xml:"http://www.portalfiscal.inf.br/nfe cCredPresumido" json:"cCredPresumido"`
	PCredPresumido TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfe pCredPresumido" json:"pCredPresumido"`
	VCredPresumido TDec1302    `xml:"http://www.portalfiscal.inf.br/nfe vCredPresumido" json:"vCredPresumido"`
}

type TNFeInfNFeDetProdDIAdi struct {
	NAdicao     *string      `xml:"http://www.portalfiscal.inf.br/nfe nAdicao,omitempty" json:"nAdicao,omitempty"`
	NSeqAdic    string       `xml:"http://www.portalfiscal.inf.br/nfe nSeqAdic" json:"nSeqAdic"`
	CFabricante string       `xml:"http://www.portalfiscal.inf.br/nfe cFabricante" json:"cFabricante"`
	VDescDI     *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vDescDI,omitempty" json:"vDescDI,omitempty"`
	NDraw       *string      `xml:"http://www.portalfiscal.inf.br/nfe nDraw,omitempty" json:"nDraw,omitempty"`
}

type TNFeInfNFeDetProdDI struct {
	NDI          string                   `xml:"http://www.portalfiscal.inf.br/nfe nDI" json:"nDI"`
	DDI          TData                    `xml:"http://www.portalfiscal.inf.br/nfe dDI" json:"dDI"`
	XLocDesemb   string                   `xml:"http://www.portalfiscal.inf.br/nfe xLocDesemb" json:"xLocDesemb"`
	UFDesemb     TUFEmi                   `xml:"http://www.portalfiscal.inf.br/nfe UFDesemb" json:"UFDesemb"`
	DDesemb      TData                    `xml:"http://www.portalfiscal.inf.br/nfe dDesemb" json:"dDesemb"`
	TpViaTransp  string                   `xml:"http://www.portalfiscal.inf.br/nfe tpViaTransp" json:"tpViaTransp"`
	VAFRMM       *TDec1302                `xml:"http://www.portalfiscal.inf.br/nfe vAFRMM,omitempty" json:"vAFRMM,omitempty"`
	TpIntermedio string                   `xml:"http://www.portalfiscal.inf.br/nfe tpIntermedio" json:"tpIntermedio"`
	CNPJ         *TCNPJ                   `xml:"http://www.portalfiscal.inf.br/nfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF          *TCPF                    `xml:"http://www.portalfiscal.inf.br/nfe CPF,omitempty" json:"CPF,omitempty"`
	UFTerceiro   *TUFEmi                  `xml:"http://www.portalfiscal.inf.br/nfe UFTerceiro,omitempty" json:"UFTerceiro,omitempty"`
	CExportador  string                   `xml:"http://www.portalfiscal.inf.br/nfe cExportador" json:"cExportador"`
	Adi          []TNFeInfNFeDetProdDIAdi `xml:"http://www.portalfiscal.inf.br/nfe adi,omitempty" json:"adi,omitempty"`
}

type TNFeInfNFeDetProdDetExportExportInd struct {
	NRE     string    `xml:"http://www.portalfiscal.inf.br/nfe nRE" json:"nRE"`
	ChNFe   TChNFe    `xml:"http://www.portalfiscal.inf.br/nfe chNFe" json:"chNFe"`
	QExport TDec1104v `xml:"http://www.portalfiscal.inf.br/nfe qExport" json:"qExport"`
}

type TNFeInfNFeDetProdDetExport struct {
	NDraw     *string                              `xml:"http://www.portalfiscal.inf.br/nfe nDraw,omitempty" json:"nDraw,omitempty"`
	ExportInd *TNFeInfNFeDetProdDetExportExportInd `xml:"http://www.portalfiscal.inf.br/nfe exportInd,omitempty" json:"exportInd,omitempty"`
}

type TNFeInfNFeDetProdRastro struct {
	NLote  string    `xml:"http://www.portalfiscal.inf.br/nfe nLote" json:"nLote"`
	QLote  TDec0803v `xml:"http://www.portalfiscal.inf.br/nfe qLote" json:"qLote"`
	DFab   TData     `xml:"http://www.portalfiscal.inf.br/nfe dFab" json:"dFab"`
	DVal   TData     `xml:"http://www.portalfiscal.inf.br/nfe dVal" json:"dVal"`
	CAgreg *string   `xml:"http://www.portalfiscal.inf.br/nfe cAgreg,omitempty" json:"cAgreg,omitempty"`
}

type TNFeInfNFeDetProdInfProdNFF struct {
	CProdFisco string `xml:"http://www.portalfiscal.inf.br/nfe cProdFisco" json:"cProdFisco"`
	COperNFF   string `xml:"http://www.portalfiscal.inf.br/nfe cOperNFF" json:"cOperNFF"`
}

type TNFeInfNFeDetProdInfProdEmb struct {
	XEmb    string    `xml:"http://www.portalfiscal.inf.br/nfe xEmb" json:"xEmb"`
	QVolEmb TDec0803v `xml:"http://www.portalfiscal.inf.br/nfe qVolEmb" json:"qVolEmb"`
	UEmb    string    `xml:"http://www.portalfiscal.inf.br/nfe uEmb" json:"uEmb"`
}

type TNFeInfNFeDetProdVeicProd struct {
	TpOp         string `xml:"http://www.portalfiscal.inf.br/nfe tpOp" json:"tpOp"`
	Chassi       string `xml:"http://www.portalfiscal.inf.br/nfe chassi" json:"chassi"`
	CCor         string `xml:"http://www.portalfiscal.inf.br/nfe cCor" json:"cCor"`
	XCor         string `xml:"http://www.portalfiscal.inf.br/nfe xCor" json:"xCor"`
	Pot          string `xml:"http://www.portalfiscal.inf.br/nfe pot" json:"pot"`
	Cilin        string `xml:"http://www.portalfiscal.inf.br/nfe cilin" json:"cilin"`
	PesoL        string `xml:"http://www.portalfiscal.inf.br/nfe pesoL" json:"pesoL"`
	PesoB        string `xml:"http://www.portalfiscal.inf.br/nfe pesoB" json:"pesoB"`
	NSerie       string `xml:"http://www.portalfiscal.inf.br/nfe nSerie" json:"nSerie"`
	TpComb       string `xml:"http://www.portalfiscal.inf.br/nfe tpComb" json:"tpComb"`
	NMotor       string `xml:"http://www.portalfiscal.inf.br/nfe nMotor" json:"nMotor"`
	CMT          string `xml:"http://www.portalfiscal.inf.br/nfe CMT" json:"CMT"`
	Dist         string `xml:"http://www.portalfiscal.inf.br/nfe dist" json:"dist"`
	AnoMod       string `xml:"http://www.portalfiscal.inf.br/nfe anoMod" json:"anoMod"`
	AnoFab       string `xml:"http://www.portalfiscal.inf.br/nfe anoFab" json:"anoFab"`
	TpPint       string `xml:"http://www.portalfiscal.inf.br/nfe tpPint" json:"tpPint"`
	TpVeic       string `xml:"http://www.portalfiscal.inf.br/nfe tpVeic" json:"tpVeic"`
	EspVeic      string `xml:"http://www.portalfiscal.inf.br/nfe espVeic" json:"espVeic"`
	VIN          string `xml:"http://www.portalfiscal.inf.br/nfe VIN" json:"VIN"`
	CondVeic     string `xml:"http://www.portalfiscal.inf.br/nfe condVeic" json:"condVeic"`
	CMod         string `xml:"http://www.portalfiscal.inf.br/nfe cMod" json:"cMod"`
	CCorDENATRAN string `xml:"http://www.portalfiscal.inf.br/nfe cCorDENATRAN" json:"cCorDENATRAN"`
	Lota         string `xml:"http://www.portalfiscal.inf.br/nfe lota" json:"lota"`
	TpRest       string `xml:"http://www.portalfiscal.inf.br/nfe tpRest" json:"tpRest"`
}

type TNFeInfNFeDetProdMed struct {
	CProdANVISA    string   `xml:"http://www.portalfiscal.inf.br/nfe cProdANVISA" json:"cProdANVISA"`
	XMotivoIsencao *string  `xml:"http://www.portalfiscal.inf.br/nfe xMotivoIsencao,omitempty" json:"xMotivoIsencao,omitempty"`
	VPMC           TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vPMC" json:"vPMC"`
}

type TNFeInfNFeDetProdArma struct {
	TpArma string `xml:"http://www.portalfiscal.inf.br/nfe tpArma" json:"tpArma"`
	NSerie string `xml:"http://www.portalfiscal.inf.br/nfe nSerie" json:"nSerie"`
	NCano  string `xml:"http://www.portalfiscal.inf.br/nfe nCano" json:"nCano"`
	Descr  string `xml:"http://www.portalfiscal.inf.br/nfe descr" json:"descr"`
}

type TNFeInfNFeDetProdCombCIDE struct {
	QBCProd   TDec1204v `xml:"http://www.portalfiscal.inf.br/nfe qBCProd" json:"qBCProd"`
	VAliqProd TDec1104  `xml:"http://www.portalfiscal.inf.br/nfe vAliqProd" json:"vAliqProd"`
	VCIDE     TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vCIDE" json:"vCIDE"`
}

type TNFeInfNFeDetProdCombEncerrante struct {
	NBico   string   `xml:"http://www.portalfiscal.inf.br/nfe nBico" json:"nBico"`
	NBomba  *string  `xml:"http://www.portalfiscal.inf.br/nfe nBomba,omitempty" json:"nBomba,omitempty"`
	NTanque string   `xml:"http://www.portalfiscal.inf.br/nfe nTanque" json:"nTanque"`
	VEncIni TDec1203 `xml:"http://www.portalfiscal.inf.br/nfe vEncIni" json:"vEncIni"`
	VEncFin TDec1203 `xml:"http://www.portalfiscal.inf.br/nfe vEncFin" json:"vEncFin"`
}

type TNFeInfNFeDetProdCombOrigComb struct {
	IndImport string                `xml:"http://www.portalfiscal.inf.br/nfe indImport" json:"indImport"`
	CUFOrig   TCodUFIBGE            `xml:"http://www.portalfiscal.inf.br/nfe cUFOrig" json:"cUFOrig"`
	POrig     TDec03v00a04Max100Opc `xml:"http://www.portalfiscal.inf.br/nfe pOrig" json:"pOrig"`
}

type TNFeInfNFeDetProdComb struct {
	CProdANP   string                           `xml:"http://www.portalfiscal.inf.br/nfe cProdANP" json:"cProdANP"`
	DescANP    string                           `xml:"http://www.portalfiscal.inf.br/nfe descANP" json:"descANP"`
	PGLP       *TDec0302a04Max100               `xml:"http://www.portalfiscal.inf.br/nfe pGLP,omitempty" json:"pGLP,omitempty"`
	PGNn       *TDec0302a04Max100               `xml:"http://www.portalfiscal.inf.br/nfe pGNn,omitempty" json:"pGNn,omitempty"`
	PGNi       *TDec0302a04Max100               `xml:"http://www.portalfiscal.inf.br/nfe pGNi,omitempty" json:"pGNi,omitempty"`
	VPart      *TDec1302                        `xml:"http://www.portalfiscal.inf.br/nfe vPart,omitempty" json:"vPart,omitempty"`
	CODIF      *string                          `xml:"http://www.portalfiscal.inf.br/nfe CODIF,omitempty" json:"CODIF,omitempty"`
	QTemp      *TDec1204temperatura             `xml:"http://www.portalfiscal.inf.br/nfe qTemp,omitempty" json:"qTemp,omitempty"`
	UFCons     TUF                              `xml:"http://www.portalfiscal.inf.br/nfe UFCons" json:"UFCons"`
	CIDE       *TNFeInfNFeDetProdCombCIDE       `xml:"http://www.portalfiscal.inf.br/nfe CIDE,omitempty" json:"CIDE,omitempty"`
	Encerrante *TNFeInfNFeDetProdCombEncerrante `xml:"http://www.portalfiscal.inf.br/nfe encerrante,omitempty" json:"encerrante,omitempty"`
	PBio       *TDec03v00a04Max100Opc           `xml:"http://www.portalfiscal.inf.br/nfe pBio,omitempty" json:"pBio,omitempty"`
	OrigComb   []TNFeInfNFeDetProdCombOrigComb  `xml:"http://www.portalfiscal.inf.br/nfe origComb,omitempty" json:"origComb,omitempty"`
}

type TNFeInfNFeDetProd struct {
	CProd            string                       `xml:"http://www.portalfiscal.inf.br/nfe cProd" json:"cProd"`
	CEAN             string                       `xml:"http://www.portalfiscal.inf.br/nfe cEAN" json:"cEAN"`
	CBarra           *string                      `xml:"http://www.portalfiscal.inf.br/nfe cBarra,omitempty" json:"cBarra,omitempty"`
	XProd            string                       `xml:"http://www.portalfiscal.inf.br/nfe xProd" json:"xProd"`
	NCM              string                       `xml:"http://www.portalfiscal.inf.br/nfe NCM" json:"NCM"`
	NVE              []string                     `xml:"http://www.portalfiscal.inf.br/nfe NVE,omitempty" json:"NVE,omitempty"`
	CEST             *string                      `xml:"http://www.portalfiscal.inf.br/nfe CEST,omitempty" json:"CEST,omitempty"`
	IndEscala        *string                      `xml:"http://www.portalfiscal.inf.br/nfe indEscala,omitempty" json:"indEscala,omitempty"`
	CNPJFab          *TCNPJ                       `xml:"http://www.portalfiscal.inf.br/nfe CNPJFab,omitempty" json:"CNPJFab,omitempty"`
	CBenef           *string                      `xml:"http://www.portalfiscal.inf.br/nfe cBenef,omitempty" json:"cBenef,omitempty"`
	GCred            []TNFeInfNFeDetProdGCred     `xml:"http://www.portalfiscal.inf.br/nfe gCred,omitempty" json:"gCred,omitempty"`
	TpCredPresIBSZFM *TTpCredPresIBSZFM           `xml:"http://www.portalfiscal.inf.br/nfe tpCredPresIBSZFM,omitempty" json:"tpCredPresIBSZFM,omitempty"`
	EXTIPI           *string                      `xml:"http://www.portalfiscal.inf.br/nfe EXTIPI,omitempty" json:"EXTIPI,omitempty"`
	CFOP             string                       `xml:"http://www.portalfiscal.inf.br/nfe CFOP" json:"CFOP"`
	UCom             string                       `xml:"http://www.portalfiscal.inf.br/nfe uCom" json:"uCom"`
	QCom             TDec1104v                    `xml:"http://www.portalfiscal.inf.br/nfe qCom" json:"qCom"`
	VUnCom           TDec1110v                    `xml:"http://www.portalfiscal.inf.br/nfe vUnCom" json:"vUnCom"`
	VProd            TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfe vProd" json:"vProd"`
	VUnComLiq        *TDec1110v                   `xml:"http://www.portalfiscal.inf.br/nfe vUnComLiq,omitempty" json:"vUnComLiq,omitempty"`
	VProdLiq         *TDec1302                    `xml:"http://www.portalfiscal.inf.br/nfe vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	CEANTrib         string                       `xml:"http://www.portalfiscal.inf.br/nfe cEANTrib" json:"cEANTrib"`
	CBarraTrib       *string                      `xml:"http://www.portalfiscal.inf.br/nfe cBarraTrib,omitempty" json:"cBarraTrib,omitempty"`
	UTrib            string                       `xml:"http://www.portalfiscal.inf.br/nfe uTrib" json:"uTrib"`
	QTrib            TDec1104v                    `xml:"http://www.portalfiscal.inf.br/nfe qTrib" json:"qTrib"`
	VUnTrib          TDec1110v                    `xml:"http://www.portalfiscal.inf.br/nfe vUnTrib" json:"vUnTrib"`
	VFrete           *TDec1302Opc                 `xml:"http://www.portalfiscal.inf.br/nfe vFrete,omitempty" json:"vFrete,omitempty"`
	VSeg             *TDec1302Opc                 `xml:"http://www.portalfiscal.inf.br/nfe vSeg,omitempty" json:"vSeg,omitempty"`
	VDesc            *TDec1302Opc                 `xml:"http://www.portalfiscal.inf.br/nfe vDesc,omitempty" json:"vDesc,omitempty"`
	VOutro           *TDec1302Opc                 `xml:"http://www.portalfiscal.inf.br/nfe vOutro,omitempty" json:"vOutro,omitempty"`
	IndTot           string                       `xml:"http://www.portalfiscal.inf.br/nfe indTot" json:"indTot"`
	IndBemMovelUsado *string                      `xml:"http://www.portalfiscal.inf.br/nfe indBemMovelUsado,omitempty" json:"indBemMovelUsado,omitempty"`
	DI               []TNFeInfNFeDetProdDI        `xml:"http://www.portalfiscal.inf.br/nfe DI,omitempty" json:"DI,omitempty"`
	DetExport        []TNFeInfNFeDetProdDetExport `xml:"http://www.portalfiscal.inf.br/nfe detExport,omitempty" json:"detExport,omitempty"`
	XPed             *string                      `xml:"http://www.portalfiscal.inf.br/nfe xPed,omitempty" json:"xPed,omitempty"`
	NItemPed         *string                      `xml:"http://www.portalfiscal.inf.br/nfe nItemPed,omitempty" json:"nItemPed,omitempty"`
	NFCI             *TGuid                       `xml:"http://www.portalfiscal.inf.br/nfe nFCI,omitempty" json:"nFCI,omitempty"`
	Rastro           []TNFeInfNFeDetProdRastro    `xml:"http://www.portalfiscal.inf.br/nfe rastro,omitempty" json:"rastro,omitempty"`
	InfProdNFF       *TNFeInfNFeDetProdInfProdNFF `xml:"http://www.portalfiscal.inf.br/nfe infProdNFF,omitempty" json:"infProdNFF,omitempty"`
	InfProdEmb       *TNFeInfNFeDetProdInfProdEmb `xml:"http://www.portalfiscal.inf.br/nfe infProdEmb,omitempty" json:"infProdEmb,omitempty"`
	VeicProd         *TNFeInfNFeDetProdVeicProd   `xml:"http://www.portalfiscal.inf.br/nfe veicProd,omitempty" json:"veicProd,omitempty"`
	Med              *TNFeInfNFeDetProdMed        `xml:"http://www.portalfiscal.inf.br/nfe med,omitempty" json:"med,omitempty"`
	Arma             []TNFeInfNFeDetProdArma      `xml:"http://www.portalfiscal.inf.br/nfe arma,omitempty" json:"arma,omitempty"`
	Comb             *TNFeInfNFeDetProdComb       `xml:"http://www.portalfiscal.inf.br/nfe comb,omitempty" json:"comb,omitempty"`
	NRECOPI          *string                      `xml:"http://www.portalfiscal.inf.br/nfe nRECOPI,omitempty" json:"nRECOPI,omitempty"`
}

type TNFeInfNFeDetImposto struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type TNFeInfNFeDetImpostoDevolIPI struct {
	VIPIDevol TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vIPIDevol" json:"vIPIDevol"`
}

type TNFeInfNFeDetImpostoDevol struct {
	PDevol TDec0302Max100               `xml:"http://www.portalfiscal.inf.br/nfe pDevol" json:"pDevol"`
	IPI    TNFeInfNFeDetImpostoDevolIPI `xml:"http://www.portalfiscal.inf.br/nfe IPI" json:"IPI"`
}

type TNFeInfNFeDetObsItemObsCont struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/nfe xTexto" json:"xTexto"`
}

type TNFeInfNFeDetObsItemObsFisco struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/nfe xTexto" json:"xTexto"`
}

type TNFeInfNFeDetObsItem struct {
	ObsCont  *TNFeInfNFeDetObsItemObsCont  `xml:"http://www.portalfiscal.inf.br/nfe obsCont,omitempty" json:"obsCont,omitempty"`
	ObsFisco *TNFeInfNFeDetObsItemObsFisco `xml:"http://www.portalfiscal.inf.br/nfe obsFisco,omitempty" json:"obsFisco,omitempty"`
}

type TNFeInfNFeDetDFeReferenciado struct {
	ChaveAcesso TChNFe  `xml:"http://www.portalfiscal.inf.br/nfe chaveAcesso" json:"chaveAcesso"`
	NItem       *string `xml:"http://www.portalfiscal.inf.br/nfe nItem,omitempty" json:"nItem,omitempty"`
}

type TNFeInfNFeDet struct {
	NItem           string                        `xml:"nItem,attr" json:"nItem"`
	Prod            TNFeInfNFeDetProd             `xml:"http://www.portalfiscal.inf.br/nfe prod" json:"prod"`
	Imposto         TNFeInfNFeDetImposto          `xml:"http://www.portalfiscal.inf.br/nfe imposto" json:"imposto"`
	ImpostoDevol    *TNFeInfNFeDetImpostoDevol    `xml:"http://www.portalfiscal.inf.br/nfe impostoDevol,omitempty" json:"impostoDevol,omitempty"`
	InfAdProd       *string                       `xml:"http://www.portalfiscal.inf.br/nfe infAdProd,omitempty" json:"infAdProd,omitempty"`
	ObsItem         *TNFeInfNFeDetObsItem         `xml:"http://www.portalfiscal.inf.br/nfe obsItem,omitempty" json:"obsItem,omitempty"`
	VItem           *TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfe vItem,omitempty" json:"vItem,omitempty"`
	DFeReferenciado *TNFeInfNFeDetDFeReferenciado `xml:"http://www.portalfiscal.inf.br/nfe DFeReferenciado,omitempty" json:"DFeReferenciado,omitempty"`
}

type TNFeInfNFeTotalICMSTot struct {
	VBC            TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vBC" json:"vBC"`
	VICMS          TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vICMS" json:"vICMS"`
	VICMSDeson     TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vICMSDeson" json:"vICMSDeson"`
	VFCPUFDest     *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vFCPUFDest,omitempty" json:"vFCPUFDest,omitempty"`
	VICMSUFDest    *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vICMSUFDest,omitempty" json:"vICMSUFDest,omitempty"`
	VICMSUFRemet   *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vICMSUFRemet,omitempty" json:"vICMSUFRemet,omitempty"`
	VFCP           TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vFCP" json:"vFCP"`
	VBCST          TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vBCST" json:"vBCST"`
	VST            TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vST" json:"vST"`
	VFCPST         TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vFCPST" json:"vFCPST"`
	VFCPSTRet      TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vFCPSTRet" json:"vFCPSTRet"`
	QBCMono        *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe qBCMono,omitempty" json:"qBCMono,omitempty"`
	VICMSMono      *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vICMSMono,omitempty" json:"vICMSMono,omitempty"`
	QBCMonoReten   *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe qBCMonoReten,omitempty" json:"qBCMonoReten,omitempty"`
	VICMSMonoReten *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vICMSMonoReten,omitempty" json:"vICMSMonoReten,omitempty"`
	QBCMonoRet     *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe qBCMonoRet,omitempty" json:"qBCMonoRet,omitempty"`
	VICMSMonoRet   *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vICMSMonoRet,omitempty" json:"vICMSMonoRet,omitempty"`
	VProd          TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vProd" json:"vProd"`
	VFrete         TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vFrete" json:"vFrete"`
	VSeg           TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vSeg" json:"vSeg"`
	VDesc          TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vDesc" json:"vDesc"`
	VII            TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vII" json:"vII"`
	VIPI           TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vIPI" json:"vIPI"`
	VIPIDevol      TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vIPIDevol" json:"vIPIDevol"`
	VPIS           TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vPIS" json:"vPIS"`
	VCOFINS        TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vCOFINS" json:"vCOFINS"`
	VOutro         TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vOutro" json:"vOutro"`
	VNF            TDec1302  `xml:"http://www.portalfiscal.inf.br/nfe vNF" json:"vNF"`
	VTotTrib       *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vTotTrib,omitempty" json:"vTotTrib,omitempty"`
}

type TNFeInfNFeTotalISSQNtot struct {
	VServ       *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vServ,omitempty" json:"vServ,omitempty"`
	VBC         *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vBC,omitempty" json:"vBC,omitempty"`
	VISS        *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vISS,omitempty" json:"vISS,omitempty"`
	VPIS        *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vPIS,omitempty" json:"vPIS,omitempty"`
	VCOFINS     *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vCOFINS,omitempty" json:"vCOFINS,omitempty"`
	DCompet     TData        `xml:"http://www.portalfiscal.inf.br/nfe dCompet" json:"dCompet"`
	VDeducao    *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vDeducao,omitempty" json:"vDeducao,omitempty"`
	VOutro      *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vOutro,omitempty" json:"vOutro,omitempty"`
	VDescIncond *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vDescIncond,omitempty" json:"vDescIncond,omitempty"`
	VDescCond   *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vDescCond,omitempty" json:"vDescCond,omitempty"`
	VISSRet     *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vISSRet,omitempty" json:"vISSRet,omitempty"`
	CRegTrib    *string      `xml:"http://www.portalfiscal.inf.br/nfe cRegTrib,omitempty" json:"cRegTrib,omitempty"`
}

type TNFeInfNFeTotalRetTrib struct {
	VRetPIS    *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vRetPIS,omitempty" json:"vRetPIS,omitempty"`
	VRetCOFINS *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vRetCOFINS,omitempty" json:"vRetCOFINS,omitempty"`
	VRetCSLL   *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vRetCSLL,omitempty" json:"vRetCSLL,omitempty"`
	VBCIRRF    *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vBCIRRF,omitempty" json:"vBCIRRF,omitempty"`
	VIRRF      *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vIRRF,omitempty" json:"vIRRF,omitempty"`
	VBCRetPrev *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vBCRetPrev,omitempty" json:"vBCRetPrev,omitempty"`
	VRetPrev   *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vRetPrev,omitempty" json:"vRetPrev,omitempty"`
}

type TNFeInfNFeTotal struct {
	VProdLiqTot *TDec1302                `xml:"http://www.portalfiscal.inf.br/nfe vProdLiqTot,omitempty" json:"vProdLiqTot,omitempty"`
	ICMSTot     TNFeInfNFeTotalICMSTot   `xml:"http://www.portalfiscal.inf.br/nfe ICMSTot" json:"ICMSTot"`
	ISSQNtot    *TNFeInfNFeTotalISSQNtot `xml:"http://www.portalfiscal.inf.br/nfe ISSQNtot,omitempty" json:"ISSQNtot,omitempty"`
	RetTrib     *TNFeInfNFeTotalRetTrib  `xml:"http://www.portalfiscal.inf.br/nfe retTrib,omitempty" json:"retTrib,omitempty"`
	ISTot       *TISTot                  `xml:"http://www.portalfiscal.inf.br/nfe ISTot,omitempty" json:"ISTot,omitempty"`
	IBSCBSTot   *TIBSCBSMonoTot          `xml:"http://www.portalfiscal.inf.br/nfe IBSCBSTot,omitempty" json:"IBSCBSTot,omitempty"`
	VNFTot      *TDec1302                `xml:"http://www.portalfiscal.inf.br/nfe vNFTot,omitempty" json:"vNFTot,omitempty"`
}

type TNFeInfNFeTranspTransporta struct {
	CNPJ   *TCNPJ   `xml:"http://www.portalfiscal.inf.br/nfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF    *TCPF    `xml:"http://www.portalfiscal.inf.br/nfe CPF,omitempty" json:"CPF,omitempty"`
	XNome  *string  `xml:"http://www.portalfiscal.inf.br/nfe xNome,omitempty" json:"xNome,omitempty"`
	IE     *TIeDest `xml:"http://www.portalfiscal.inf.br/nfe IE,omitempty" json:"IE,omitempty"`
	XEnder *string  `xml:"http://www.portalfiscal.inf.br/nfe xEnder,omitempty" json:"xEnder,omitempty"`
	XMun   *string  `xml:"http://www.portalfiscal.inf.br/nfe xMun,omitempty" json:"xMun,omitempty"`
	UF     *TUF     `xml:"http://www.portalfiscal.inf.br/nfe UF,omitempty" json:"UF,omitempty"`
}

type TNFeInfNFeTranspRetTransp struct {
	VServ    TDec1302    `xml:"http://www.portalfiscal.inf.br/nfe vServ" json:"vServ"`
	VBCRet   TDec1302    `xml:"http://www.portalfiscal.inf.br/nfe vBCRet" json:"vBCRet"`
	PICMSRet TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfe pICMSRet" json:"pICMSRet"`
	VICMSRet TDec1302    `xml:"http://www.portalfiscal.inf.br/nfe vICMSRet" json:"vICMSRet"`
	CFOP     string      `xml:"http://www.portalfiscal.inf.br/nfe CFOP" json:"CFOP"`
	CMunFG   TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfe cMunFG" json:"cMunFG"`
}

type TNFeInfNFeTranspVolLacres struct {
	NLacre string `xml:"http://www.portalfiscal.inf.br/nfe nLacre" json:"nLacre"`
}

type TNFeInfNFeTranspVol struct {
	QVol   *string                     `xml:"http://www.portalfiscal.inf.br/nfe qVol,omitempty" json:"qVol,omitempty"`
	Esp    *string                     `xml:"http://www.portalfiscal.inf.br/nfe esp,omitempty" json:"esp,omitempty"`
	Marca  *string                     `xml:"http://www.portalfiscal.inf.br/nfe marca,omitempty" json:"marca,omitempty"`
	NVol   *string                     `xml:"http://www.portalfiscal.inf.br/nfe nVol,omitempty" json:"nVol,omitempty"`
	PesoL  *TDec1203                   `xml:"http://www.portalfiscal.inf.br/nfe pesoL,omitempty" json:"pesoL,omitempty"`
	PesoB  *TDec1203                   `xml:"http://www.portalfiscal.inf.br/nfe pesoB,omitempty" json:"pesoB,omitempty"`
	Lacres []TNFeInfNFeTranspVolLacres `xml:"http://www.portalfiscal.inf.br/nfe lacres,omitempty" json:"lacres,omitempty"`
}

type TNFeInfNFeTransp struct {
	ModFrete   string                      `xml:"http://www.portalfiscal.inf.br/nfe modFrete" json:"modFrete"`
	Transporta *TNFeInfNFeTranspTransporta `xml:"http://www.portalfiscal.inf.br/nfe transporta,omitempty" json:"transporta,omitempty"`
	RetTransp  *TNFeInfNFeTranspRetTransp  `xml:"http://www.portalfiscal.inf.br/nfe retTransp,omitempty" json:"retTransp,omitempty"`
	VeicTransp *TVeiculo                   `xml:"http://www.portalfiscal.inf.br/nfe veicTransp,omitempty" json:"veicTransp,omitempty"`
	Reboque    []TVeiculo                  `xml:"http://www.portalfiscal.inf.br/nfe reboque,omitempty" json:"reboque,omitempty"`
	Vagao      *string                     `xml:"http://www.portalfiscal.inf.br/nfe vagao,omitempty" json:"vagao,omitempty"`
	Balsa      *string                     `xml:"http://www.portalfiscal.inf.br/nfe balsa,omitempty" json:"balsa,omitempty"`
	Vol        []TNFeInfNFeTranspVol       `xml:"http://www.portalfiscal.inf.br/nfe vol,omitempty" json:"vol,omitempty"`
}

type TNFeInfNFeCobrFat struct {
	NFat  *string   `xml:"http://www.portalfiscal.inf.br/nfe nFat,omitempty" json:"nFat,omitempty"`
	VOrig *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vOrig,omitempty" json:"vOrig,omitempty"`
	VDesc *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vDesc,omitempty" json:"vDesc,omitempty"`
	VLiq  *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vLiq,omitempty" json:"vLiq,omitempty"`
}

type TNFeInfNFeCobrDup struct {
	NDup  *string     `xml:"http://www.portalfiscal.inf.br/nfe nDup,omitempty" json:"nDup,omitempty"`
	DVenc *TData      `xml:"http://www.portalfiscal.inf.br/nfe dVenc,omitempty" json:"dVenc,omitempty"`
	VDup  TDec1302Opc `xml:"http://www.portalfiscal.inf.br/nfe vDup" json:"vDup"`
}

type TNFeInfNFeCobr struct {
	Fat *TNFeInfNFeCobrFat  `xml:"http://www.portalfiscal.inf.br/nfe fat,omitempty" json:"fat,omitempty"`
	Dup []TNFeInfNFeCobrDup `xml:"http://www.portalfiscal.inf.br/nfe dup,omitempty" json:"dup,omitempty"`
}

type TNFeInfNFePagDetPagCard struct {
	TpIntegra string  `xml:"http://www.portalfiscal.inf.br/nfe tpIntegra" json:"tpIntegra"`
	CNPJ      *TCNPJ  `xml:"http://www.portalfiscal.inf.br/nfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	TBand     *string `xml:"http://www.portalfiscal.inf.br/nfe tBand,omitempty" json:"tBand,omitempty"`
	CAut      *string `xml:"http://www.portalfiscal.inf.br/nfe cAut,omitempty" json:"cAut,omitempty"`
	CNPJReceb *TCNPJ  `xml:"http://www.portalfiscal.inf.br/nfe CNPJReceb,omitempty" json:"CNPJReceb,omitempty"`
	IDTermPag *string `xml:"http://www.portalfiscal.inf.br/nfe idTermPag,omitempty" json:"idTermPag,omitempty"`
}

type TNFeInfNFePagDetPag struct {
	IndPag  *string                  `xml:"http://www.portalfiscal.inf.br/nfe indPag,omitempty" json:"indPag,omitempty"`
	TPag    string                   `xml:"http://www.portalfiscal.inf.br/nfe tPag" json:"tPag"`
	XPag    *string                  `xml:"http://www.portalfiscal.inf.br/nfe xPag,omitempty" json:"xPag,omitempty"`
	VPag    TDec1302                 `xml:"http://www.portalfiscal.inf.br/nfe vPag" json:"vPag"`
	DPag    *TData                   `xml:"http://www.portalfiscal.inf.br/nfe dPag,omitempty" json:"dPag,omitempty"`
	CNPJPag *TCNPJ                   `xml:"http://www.portalfiscal.inf.br/nfe CNPJPag,omitempty" json:"CNPJPag,omitempty"`
	UFPag   *TUFEmi                  `xml:"http://www.portalfiscal.inf.br/nfe UFPag,omitempty" json:"UFPag,omitempty"`
	Card    *TNFeInfNFePagDetPagCard `xml:"http://www.portalfiscal.inf.br/nfe card,omitempty" json:"card,omitempty"`
}

type TNFeInfNFePag struct {
	DetPag []TNFeInfNFePagDetPag `xml:"http://www.portalfiscal.inf.br/nfe detPag,omitempty" json:"detPag,omitempty"`
	VTroco *TDec1302             `xml:"http://www.portalfiscal.inf.br/nfe vTroco,omitempty" json:"vTroco,omitempty"`
}

type TNFeInfNFeInfIntermed struct {
	CNPJ         TCNPJ  `xml:"http://www.portalfiscal.inf.br/nfe CNPJ" json:"CNPJ"`
	IDCadIntTran string `xml:"http://www.portalfiscal.inf.br/nfe idCadIntTran" json:"idCadIntTran"`
}

type TNFeInfNFeInfAdicObsCont struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/nfe xTexto" json:"xTexto"`
}

type TNFeInfNFeInfAdicObsFisco struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/nfe xTexto" json:"xTexto"`
}

type TNFeInfNFeInfAdicProcRef struct {
	NProc   string  `xml:"http://www.portalfiscal.inf.br/nfe nProc" json:"nProc"`
	IndProc string  `xml:"http://www.portalfiscal.inf.br/nfe indProc" json:"indProc"`
	TpAto   *string `xml:"http://www.portalfiscal.inf.br/nfe tpAto,omitempty" json:"tpAto,omitempty"`
}

type TNFeInfNFeInfAdic struct {
	InfAdFisco *string                     `xml:"http://www.portalfiscal.inf.br/nfe infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     *string                     `xml:"http://www.portalfiscal.inf.br/nfe infCpl,omitempty" json:"infCpl,omitempty"`
	ObsCont    []TNFeInfNFeInfAdicObsCont  `xml:"http://www.portalfiscal.inf.br/nfe obsCont,omitempty" json:"obsCont,omitempty"`
	ObsFisco   []TNFeInfNFeInfAdicObsFisco `xml:"http://www.portalfiscal.inf.br/nfe obsFisco,omitempty" json:"obsFisco,omitempty"`
	ProcRef    []TNFeInfNFeInfAdicProcRef  `xml:"http://www.portalfiscal.inf.br/nfe procRef,omitempty" json:"procRef,omitempty"`
}

type TNFeInfNFeExporta struct {
	UFSaidaPais  TUFEmi  `xml:"http://www.portalfiscal.inf.br/nfe UFSaidaPais" json:"UFSaidaPais"`
	XLocExporta  string  `xml:"http://www.portalfiscal.inf.br/nfe xLocExporta" json:"xLocExporta"`
	XLocDespacho *string `xml:"http://www.portalfiscal.inf.br/nfe xLocDespacho,omitempty" json:"xLocDespacho,omitempty"`
}

type TNFeInfNFeCompra struct {
	XNEmp *string `xml:"http://www.portalfiscal.inf.br/nfe xNEmp,omitempty" json:"xNEmp,omitempty"`
	XPed  *string `xml:"http://www.portalfiscal.inf.br/nfe xPed,omitempty" json:"xPed,omitempty"`
	XCont *string `xml:"http://www.portalfiscal.inf.br/nfe xCont,omitempty" json:"xCont,omitempty"`
}

type TNFeInfNFeCanaForDia struct {
	Dia  string    `xml:"dia,attr" json:"dia"`
	Qtde TDec1110v `xml:"http://www.portalfiscal.inf.br/nfe qtde" json:"qtde"`
}

type TNFeInfNFeCanaDeduc struct {
	XDed string   `xml:"http://www.portalfiscal.inf.br/nfe xDed" json:"xDed"`
	VDed TDec1302 `xml:"http://www.portalfiscal.inf.br/nfe vDed" json:"vDed"`
}

type TNFeInfNFeCana struct {
	Safra   string                 `xml:"http://www.portalfiscal.inf.br/nfe safra" json:"safra"`
	Ref     string                 `xml:"http://www.portalfiscal.inf.br/nfe ref" json:"ref"`
	ForDia  []TNFeInfNFeCanaForDia `xml:"http://www.portalfiscal.inf.br/nfe forDia,omitempty" json:"forDia,omitempty"`
	QTotMes TDec1110v              `xml:"http://www.portalfiscal.inf.br/nfe qTotMes" json:"qTotMes"`
	QTotAnt TDec1110v              `xml:"http://www.portalfiscal.inf.br/nfe qTotAnt" json:"qTotAnt"`
	QTotGer TDec1110v              `xml:"http://www.portalfiscal.inf.br/nfe qTotGer" json:"qTotGer"`
	Deduc   []TNFeInfNFeCanaDeduc  `xml:"http://www.portalfiscal.inf.br/nfe deduc,omitempty" json:"deduc,omitempty"`
	VFor    TDec1302               `xml:"http://www.portalfiscal.inf.br/nfe vFor" json:"vFor"`
	VTotDed TDec1302               `xml:"http://www.portalfiscal.inf.br/nfe vTotDed" json:"vTotDed"`
	VLiqFor TDec1302               `xml:"http://www.portalfiscal.inf.br/nfe vLiqFor" json:"vLiqFor"`
}

type TNFeInfNFeInfSolicNFF struct {
	XSolic string `xml:"http://www.portalfiscal.inf.br/nfe xSolic" json:"xSolic"`
}

type TNFeInfNFeAgropecuarioDefensivo struct {
	NReceituario string `xml:"http://www.portalfiscal.inf.br/nfe nReceituario" json:"nReceituario"`
	CPFRespTec   TCPF   `xml:"http://www.portalfiscal.inf.br/nfe CPFRespTec" json:"CPFRespTec"`
}

type TNFeInfNFeAgropecuarioGuiaTransito struct {
	TpGuia    string  `xml:"http://www.portalfiscal.inf.br/nfe tpGuia" json:"tpGuia"`
	UFGuia    TUFEmi  `xml:"http://www.portalfiscal.inf.br/nfe UFGuia" json:"UFGuia"`
	SerieGuia *string `xml:"http://www.portalfiscal.inf.br/nfe serieGuia,omitempty" json:"serieGuia,omitempty"`
	NGuia     string  `xml:"http://www.portalfiscal.inf.br/nfe nGuia" json:"nGuia"`
}

type TNFeInfNFeAgropecuario struct {
	Defensivo    []TNFeInfNFeAgropecuarioDefensivo   `xml:"http://www.portalfiscal.inf.br/nfe defensivo,omitempty" json:"defensivo,omitempty"`
	GuiaTransito *TNFeInfNFeAgropecuarioGuiaTransito `xml:"http://www.portalfiscal.inf.br/nfe guiaTransito,omitempty" json:"guiaTransito,omitempty"`
}

type TNFeInfNFeInfPAAPAASignature struct {
	SignatureValue Base64Binary     `xml:"http://www.portalfiscal.inf.br/nfe SignatureValue" json:"SignatureValue"`
	RSAKeyValue    TRSAKeyValueType `xml:"http://www.portalfiscal.inf.br/nfe RSAKeyValue" json:"RSAKeyValue"`
}

type TNFeInfNFeInfPAA struct {
	CNPJPAA      TCNPJ                        `xml:"http://www.portalfiscal.inf.br/nfe CNPJPAA" json:"CNPJPAA"`
	PAASignature TNFeInfNFeInfPAAPAASignature `xml:"http://www.portalfiscal.inf.br/nfe PAASignature" json:"PAASignature"`
}

type TNFeInfNFe struct {
	Versao       TVerNFe                 `xml:"versao,attr" json:"versao"`
	ID           string                  `xml:"Id,attr" json:"Id"`
	Ide          TNFeInfNFeIde           `xml:"http://www.portalfiscal.inf.br/nfe ide" json:"ide"`
	Emit         TNFeInfNFeEmit          `xml:"http://www.portalfiscal.inf.br/nfe emit" json:"emit"`
	Avulsa       *TNFeInfNFeAvulsa       `xml:"http://www.portalfiscal.inf.br/nfe avulsa,omitempty" json:"avulsa,omitempty"`
	Dest         *TNFeInfNFeDest         `xml:"http://www.portalfiscal.inf.br/nfe dest,omitempty" json:"dest,omitempty"`
	Retirada     *TLocal                 `xml:"http://www.portalfiscal.inf.br/nfe retirada,omitempty" json:"retirada,omitempty"`
	Entrega      *TLocal                 `xml:"http://www.portalfiscal.inf.br/nfe entrega,omitempty" json:"entrega,omitempty"`
	AutXML       []TNFeInfNFeAutXML      `xml:"http://www.portalfiscal.inf.br/nfe autXML,omitempty" json:"autXML,omitempty"`
	Det          []TNFeInfNFeDet         `xml:"http://www.portalfiscal.inf.br/nfe det,omitempty" json:"det,omitempty"`
	Total        TNFeInfNFeTotal         `xml:"http://www.portalfiscal.inf.br/nfe total" json:"total"`
	Transp       TNFeInfNFeTransp        `xml:"http://www.portalfiscal.inf.br/nfe transp" json:"transp"`
	Cobr         *TNFeInfNFeCobr         `xml:"http://www.portalfiscal.inf.br/nfe cobr,omitempty" json:"cobr,omitempty"`
	Pag          TNFeInfNFePag           `xml:"http://www.portalfiscal.inf.br/nfe pag" json:"pag"`
	InfIntermed  *TNFeInfNFeInfIntermed  `xml:"http://www.portalfiscal.inf.br/nfe infIntermed,omitempty" json:"infIntermed,omitempty"`
	GPgtoVinc    *TPgtoVinc              `xml:"http://www.portalfiscal.inf.br/nfe gPgtoVinc,omitempty" json:"gPgtoVinc,omitempty"`
	InfAdic      *TNFeInfNFeInfAdic      `xml:"http://www.portalfiscal.inf.br/nfe infAdic,omitempty" json:"infAdic,omitempty"`
	Exporta      *TNFeInfNFeExporta      `xml:"http://www.portalfiscal.inf.br/nfe exporta,omitempty" json:"exporta,omitempty"`
	Compra       *TNFeInfNFeCompra       `xml:"http://www.portalfiscal.inf.br/nfe compra,omitempty" json:"compra,omitempty"`
	Cana         *TNFeInfNFeCana         `xml:"http://www.portalfiscal.inf.br/nfe cana,omitempty" json:"cana,omitempty"`
	InfRespTec   *TInfRespTec            `xml:"http://www.portalfiscal.inf.br/nfe infRespTec,omitempty" json:"infRespTec,omitempty"`
	InfSolicNFF  *TNFeInfNFeInfSolicNFF  `xml:"http://www.portalfiscal.inf.br/nfe infSolicNFF,omitempty" json:"infSolicNFF,omitempty"`
	Agropecuario *TNFeInfNFeAgropecuario `xml:"http://www.portalfiscal.inf.br/nfe agropecuario,omitempty" json:"agropecuario,omitempty"`
	InfPAA       *TNFeInfNFeInfPAA       `xml:"http://www.portalfiscal.inf.br/nfe infPAA,omitempty" json:"infPAA,omitempty"`
}

type TNFeInfNFeSupl struct {
	QRCode   string `xml:"http://www.portalfiscal.inf.br/nfe qrCode" json:"qrCode"`
	URLChave string `xml:"http://www.portalfiscal.inf.br/nfe urlChave" json:"urlChave"`
}

type TNFe struct {
	InfNFe     TNFeInfNFe      `xml:"http://www.portalfiscal.inf.br/nfe infNFe" json:"infNFe"`
	InfNFeSupl *TNFeInfNFeSupl `xml:"http://www.portalfiscal.inf.br/nfe infNFeSupl,omitempty" json:"infNFeSupl,omitempty"`
	Signature  SignatureType   `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TNFeProc struct {
	Versao  TVerNFe  `xml:"versao,attr" json:"versao"`
	NFe     TNFe     `xml:"http://www.portalfiscal.inf.br/nfe NFe" json:"NFe"`
	ProtNFe TProtNFe `xml:"http://www.portalfiscal.inf.br/nfe protNFe" json:"protNFe"`
}

type TOperCompraGov string

type TPagRef struct {
	RefDFe []TChDFeRTC `xml:"http://www.portalfiscal.inf.br/nfe refDFe,omitempty" json:"refDFe,omitempty"`
}

type TPgtoVincGPgto struct {
	NPag        string       `xml:"nPag,attr" json:"nPag"`
	IDTransacao string       `xml:"idTransacao,attr" json:"idTransacao"`
	TpMeioPgto  string       `xml:"http://www.portalfiscal.inf.br/nfe tpMeioPgto" json:"tpMeioPgto"`
	CNPJReceb   TCNPJRTC     `xml:"http://www.portalfiscal.inf.br/nfe CNPJReceb" json:"CNPJReceb"`
	CNPJBasePSP TCNPJBaseRTC `xml:"http://www.portalfiscal.inf.br/nfe CNPJBasePSP" json:"CNPJBasePSP"`
}

type TPgtoVinc struct {
	GPgto []TPgtoVincGPgto `xml:"http://www.portalfiscal.inf.br/nfe gPgto,omitempty" json:"gPgto,omitempty"`
}

type TPlaca string

type TProcEmi string

type TProt string

type TProtNFeInfProt struct {
	ID *ID `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type TProtNFe struct {
	Versao    TVerNFe         `xml:"versao,attr" json:"versao"`
	InfProt   TProtNFeInfProt `xml:"http://www.portalfiscal.inf.br/nfe infProt" json:"infProt"`
	Signature *SignatureType  `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRSAKeyValueType struct {
	Modulus  Base64Binary `xml:"http://www.portalfiscal.inf.br/nfe Modulus" json:"Modulus"`
	Exponent Base64Binary `xml:"http://www.portalfiscal.inf.br/nfe Exponent" json:"Exponent"`
}

type TRec string

type TRed struct {
	PRedAliq  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqEfet" json:"pAliqEfet"`
}

type TRetConsReciNFe struct {
	Versao   TVerNFe      `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb         `xml:"http://www.portalfiscal.inf.br/nfe tpAmb" json:"tpAmb"`
	VerAplic TVerAplic    `xml:"http://www.portalfiscal.inf.br/nfe verAplic" json:"verAplic"`
	NRec     TRec         `xml:"http://www.portalfiscal.inf.br/nfe nRec" json:"nRec"`
	CStat    TStat        `xml:"http://www.portalfiscal.inf.br/nfe cStat" json:"cStat"`
	XMotivo  TMotivo      `xml:"http://www.portalfiscal.inf.br/nfe xMotivo" json:"xMotivo"`
	CUF      TCodUFIBGE   `xml:"http://www.portalfiscal.inf.br/nfe cUF" json:"cUF"`
	DhRecbto TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfe dhRecbto" json:"dhRecbto"`
	CMsg     *string      `xml:"http://www.portalfiscal.inf.br/nfe cMsg,omitempty" json:"cMsg,omitempty"`
	XMsg     *string      `xml:"http://www.portalfiscal.inf.br/nfe xMsg,omitempty" json:"xMsg,omitempty"`
	ProtNFe  []TProtNFe   `xml:"http://www.portalfiscal.inf.br/nfe protNFe,omitempty" json:"protNFe,omitempty"`
}

type TRetEnviNFeInfRec struct {
	NRec TRec `xml:"http://www.portalfiscal.inf.br/nfe nRec" json:"nRec"`
	TMed TMed `xml:"http://www.portalfiscal.inf.br/nfe tMed" json:"tMed"`
}

type TRetEnviNFe struct {
	Versao   TVerNFe            `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb               `xml:"http://www.portalfiscal.inf.br/nfe tpAmb" json:"tpAmb"`
	VerAplic TVerAplic          `xml:"http://www.portalfiscal.inf.br/nfe verAplic" json:"verAplic"`
	CStat    TStat              `xml:"http://www.portalfiscal.inf.br/nfe cStat" json:"cStat"`
	XMotivo  TMotivo            `xml:"http://www.portalfiscal.inf.br/nfe xMotivo" json:"xMotivo"`
	CUF      TCodUFIBGE         `xml:"http://www.portalfiscal.inf.br/nfe cUF" json:"cUF"`
	DhRecbto TDateTimeUTC       `xml:"http://www.portalfiscal.inf.br/nfe dhRecbto" json:"dhRecbto"`
	InfRec   *TRetEnviNFeInfRec `xml:"http://www.portalfiscal.inf.br/nfe infRec,omitempty" json:"infRec,omitempty"`
	ProtNFe  *TProtNFe          `xml:"http://www.portalfiscal.inf.br/nfe protNFe,omitempty" json:"protNFe,omitempty"`
}

type TSerie string

type TServ string

type TStat string

type TString string

type TStringRTC string

type TTime string

type TTpCredPresIBSZFM string

type TTpNFCredito string

type TTpNFDebito string

type TTransfCred struct {
	VIBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vIBS" json:"vIBS"`
	VCBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfe vCBS" json:"vCBS"`
}

type TTribBPe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCTe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCompraGov struct {
	PAliqIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vTribCBS" json:"vTribCBS"`
}

type TTribNF3e struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFAg struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFCe struct {
	CST        TCST        `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	CClassTrib TcClassTrib `xml:"http://www.portalfiscal.inf.br/nfe cClassTrib" json:"cClassTrib"`
	IndDoacao  *TIndDoacao `xml:"http://www.portalfiscal.inf.br/nfe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS    *TCIBSNFe   `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
}

type TTribNFCom struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFGas struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFe struct {
	CST             TCST             `xml:"http://www.portalfiscal.inf.br/nfe CST" json:"CST"`
	CClassTrib      TcClassTrib      `xml:"http://www.portalfiscal.inf.br/nfe cClassTrib" json:"cClassTrib"`
	IndDoacao       *TIndDoacao      `xml:"http://www.portalfiscal.inf.br/nfe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS         *TCIBSNFe        `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono     *TMonofasia      `xml:"http://www.portalfiscal.inf.br/nfe gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
	GTransfCred     *TTransfCred     `xml:"http://www.portalfiscal.inf.br/nfe gTransfCred,omitempty" json:"gTransfCred,omitempty"`
	GAjusteCompet   *TAjusteCompet   `xml:"http://www.portalfiscal.inf.br/nfe gAjusteCompet,omitempty" json:"gAjusteCompet,omitempty"`
	GEstornoCred    *TEstornoCred    `xml:"http://www.portalfiscal.inf.br/nfe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
	GCredPresOper   *TCredPresOper   `xml:"http://www.portalfiscal.inf.br/nfe gCredPresOper,omitempty" json:"gCredPresOper,omitempty"`
	GCredPresIBSZFM *TCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nfe gCredPresIBSZFM,omitempty" json:"gCredPresIBSZFM,omitempty"`
}

type TTribRegular struct {
	CSTReg             TCST          `xml:"http://www.portalfiscal.inf.br/nfe CSTReg" json:"CSTReg"`
	CClassTribReg      TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfe cClassTribReg" json:"cClassTribReg"`
	PAliqEfetRegIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqEfetRegIBSUF" json:"pAliqEfetRegIBSUF"`
	VTribRegIBSUF      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vTribRegIBSUF" json:"vTribRegIBSUF"`
	PAliqEfetRegIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqEfetRegIBSMun" json:"pAliqEfetRegIBSMun"`
	VTribRegIBSMun     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vTribRegIBSMun" json:"vTribRegIBSMun"`
	PAliqEfetRegCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfe pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS        TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfe vTribRegCBS" json:"vTribRegCBS"`
}

type TUF string

type TUFEmi string

type TVeiculo struct {
	Placa string  `xml:"http://www.portalfiscal.inf.br/nfe placa" json:"placa"`
	UF    *TUF    `xml:"http://www.portalfiscal.inf.br/nfe UF,omitempty" json:"UF,omitempty"`
	RNTC  *string `xml:"http://www.portalfiscal.inf.br/nfe RNTC,omitempty" json:"RNTC,omitempty"`
}

type TVerAplic string

type TVerNFe string

type Tano string

type TcClassTrib string

type TcCredPres string

type Torig string

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

type String string

// NFe represents the global element {http://www.portalfiscal.inf.br/nfe}NFe.
type NFe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe NFe" json:"-"`
	TNFe
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
