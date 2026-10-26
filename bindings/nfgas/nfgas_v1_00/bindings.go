// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package nfgas_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TALCZFMCBS struct {
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vTribRegCBS" json:"vTribRegCBS"`
}

type TALCZFMCBSNFe struct {
	TpALCZFMCBS     string        `xml:"http://www.portalfiscal.inf.br/nfgas tpALCZFMCBS" json:"tpALCZFMCBS"`
	NProcSuframa    *string       `xml:"http://www.portalfiscal.inf.br/nfgas nProcSuframa,omitempty" json:"nProcSuframa,omitempty"`
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vTribRegCBS" json:"vTribRegCBS"`
}

type TAcessante string

type TAjusteCompet struct {
	CompetApur TCompetApur `xml:"http://www.portalfiscal.inf.br/nfgas competApur" json:"competApur"`
	VIBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBS" json:"vIBS"`
	VCBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBS" json:"vCBS"`
}

type TAmb string

type TBand string

type TCIBSGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfgas gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfgas gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfgas gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSUF" json:"vIBSUF"`
}

type TCIBSGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfgas gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfgas gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfgas gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMun" json:"vIBSMun"`
}

type TCIBSGCBS struct {
	PCBS       TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pCBS" json:"pCBS"`
	GDif       *TDif         `xml:"http://www.portalfiscal.inf.br/nfgas gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfgas gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed         `xml:"http://www.portalfiscal.inf.br/nfgas gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBS   `xml:"http://www.portalfiscal.inf.br/nfgas gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vCBS" json:"vCBS"`
}

type TCIBS struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	GIBSUF         TCIBSGIBSUF     `xml:"http://www.portalfiscal.inf.br/nfgas gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSGIBSMun    `xml:"http://www.portalfiscal.inf.br/nfgas gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfgas vIBS" json:"vIBS"`
	GCBS           TCIBSGCBS       `xml:"http://www.portalfiscal.inf.br/nfgas gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nfgas gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfgas gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCIBSNFeGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfgas gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfgas gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfgas gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSUF" json:"vIBSUF"`
}

type TCIBSNFeGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfgas gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfgas gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfgas gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMun" json:"vIBSMun"`
}

type TCIBSNFeGCBS struct {
	PCBS       TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfgas pCBS" json:"pCBS"`
	GDif       *TDif          `xml:"http://www.portalfiscal.inf.br/nfgas gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib      `xml:"http://www.portalfiscal.inf.br/nfgas gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed          `xml:"http://www.portalfiscal.inf.br/nfgas gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBSNFe `xml:"http://www.portalfiscal.inf.br/nfgas gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfgas vCBS" json:"vCBS"`
}

type TCIBSNFe struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	GIBSUF         TCIBSNFeGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfgas gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSNFeGIBSMun `xml:"http://www.portalfiscal.inf.br/nfgas gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfgas vIBS" json:"vIBS"`
	GCBS           TCIBSNFeGCBS    `xml:"http://www.portalfiscal.inf.br/nfgas gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nfgas gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfgas gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCOrgaoIBGE string

type TCST string

type TCateg string

type TCFOP string

type TChDFe string

type TChDFeRTC string

type TClasse string

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
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nfgas tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfgas pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nfgas tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nfgas refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCompraGovReduzido struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nfgas tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfgas pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nfgas tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nfgas refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCPF string

type TCPFVar string

type TCredPres struct {
	PCredPres        TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pCredPres" json:"pCredPres"`
	VCredPres        *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfgas vCredPres,omitempty" json:"vCredPres,omitempty"`
	VCredPresCondSus *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfgas vCredPresCondSus,omitempty" json:"vCredPresCondSus,omitempty"`
}

type TCredPresIBSZFM struct {
	CompetApur       TCompetApur       `xml:"http://www.portalfiscal.inf.br/nfgas competApur" json:"competApur"`
	TpCredPresIBSZFM TTpCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nfgas tpCredPresIBSZFM" json:"tpCredPresIBSZFM"`
	VCredPresIBSZFM  TDec1302RTC       `xml:"http://www.portalfiscal.inf.br/nfgas vCredPresIBSZFM" json:"vCredPresIBSZFM"`
}

type TCredPresOper struct {
	VBCCredPres  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vBCCredPres" json:"vBCCredPres"`
	CCredPres    TcCredPres  `xml:"http://www.portalfiscal.inf.br/nfgas cCredPres" json:"cCredPres"`
	GIBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCredPres,omitempty" json:"gIBSCredPres,omitempty"`
	GCBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nfgas gCBSCredPres,omitempty" json:"gCBSCredPres,omitempty"`
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

type TDec13021310 string

type TDevTrib struct {
	PDevTrib *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pDevTrib,omitempty" json:"pDevTrib,omitempty"`
	VDevTrib TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfgas vDevTrib" json:"vDevTrib"`
}

type TDif struct {
	PDif TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pDif" json:"pDif"`
	VDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vDif" json:"vDif"`
}

type TDoc string

type TEmail string

type TEmpresaSoft struct {
	CNPJ     *TCNPJ `xml:"http://www.portalfiscal.inf.br/nfgas CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF      *TCPF  `xml:"http://www.portalfiscal.inf.br/nfgas CPF,omitempty" json:"CPF,omitempty"`
	XContato string `xml:"http://www.portalfiscal.inf.br/nfgas xContato" json:"xContato"`
	Email    TEmail `xml:"http://www.portalfiscal.inf.br/nfgas email" json:"email"`
	Fone     string `xml:"http://www.portalfiscal.inf.br/nfgas fone" json:"fone"`
}

type TEndeEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfgas xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfgas nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfgas xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfgas xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfgas cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfgas xMun" json:"xMun"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/nfgas CEP" json:"CEP"`
	UF      TUFSemEX    `xml:"http://www.portalfiscal.inf.br/nfgas UF" json:"UF"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfgas fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nfgas email,omitempty" json:"email,omitempty"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfgas xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfgas nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfgas xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfgas xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfgas cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfgas xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/nfgas CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/nfgas UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/nfgas cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/nfgas xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfgas fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nfgas email,omitempty" json:"email,omitempty"`
}

type TEnteGov string

type TEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBSEstCred" json:"vCBSEstCred"`
}

type TFaixa string

type TFase string

type TFat string

type TFatNF string

type TFinNF string

type TGrContrat string

type TGrMed string

type TGrMedGas string

type TGrpTensao string

type TIBSCBSMonoTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSUF" json:"vIBSUF"`
}

type TIBSCBSMonoTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMun" json:"vIBSMun"`
}

type TIBSCBSMonoTotGIBS struct {
	GIBSUF           TIBSCBSMonoTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfgas gIBSUF" json:"gIBSUF"`
	GIBSMun          TIBSCBSMonoTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nfgas gIBSMun" json:"gIBSMun"`
	VIBS             TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfgas vIBS" json:"vIBS"`
	VCredPres        TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfgas vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfgas vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGCBS struct {
	VDif             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDif" json:"vDif"`
	VDevTrib         TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDevTrib" json:"vDevTrib"`
	VCBS             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBS" json:"vCBS"`
	VCredPres        TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGMono struct {
	VIBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMono" json:"vIBSMono"`
	VCBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBSMono" json:"vCBSMono"`
	VIBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMonoReten" json:"vIBSMonoReten"`
	VCBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBSMonoReten" json:"vCBSMonoReten"`
	VIBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMonoRet" json:"vIBSMonoRet"`
	VCBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBSMonoRet" json:"vCBSMonoRet"`
}

type TIBSCBSMonoTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSMonoTot struct {
	VBCIBSCBS    TDec1302RTC                 `xml:"http://www.portalfiscal.inf.br/nfgas vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         *TIBSCBSMonoTotGIBS         `xml:"http://www.portalfiscal.inf.br/nfgas gIBS,omitempty" json:"gIBS,omitempty"`
	GCBS         *TIBSCBSMonoTotGCBS         `xml:"http://www.portalfiscal.inf.br/nfgas gCBS,omitempty" json:"gCBS,omitempty"`
	GMono        *TIBSCBSMonoTotGMono        `xml:"http://www.portalfiscal.inf.br/nfgas gMono,omitempty" json:"gMono,omitempty"`
	GEstornoCred *TIBSCBSMonoTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIBSCBSTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSUF" json:"vIBSUF"`
}

type TIBSCBSTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMun" json:"vIBSMun"`
}

type TIBSCBSTotGIBS struct {
	GIBSUF  TIBSCBSTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfgas gIBSUF" json:"gIBSUF"`
	GIBSMun TIBSCBSTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nfgas gIBSMun" json:"gIBSMun"`
	VIBS    TDec1302RTC           `xml:"http://www.portalfiscal.inf.br/nfgas vIBS" json:"vIBS"`
}

type TIBSCBSTotGCBS struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vDevTrib" json:"vDevTrib"`
	VCBS     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBS" json:"vCBS"`
}

type TIBSCBSTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSTot struct {
	VBCIBSCBS    TDec1302RTC             `xml:"http://www.portalfiscal.inf.br/nfgas vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         TIBSCBSTotGIBS          `xml:"http://www.portalfiscal.inf.br/nfgas gIBS" json:"gIBS"`
	GCBS         TIBSCBSTotGCBS          `xml:"http://www.portalfiscal.inf.br/nfgas gCBS" json:"gCBS"`
	GEstornoCred *TIBSCBSTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIPv4 string

type TIS struct {
	CSTIS        TCST           `xml:"http://www.portalfiscal.inf.br/nfgas CSTIS" json:"CSTIS"`
	CClassTribIS TcClassTrib    `xml:"http://www.portalfiscal.inf.br/nfgas cClassTribIS" json:"cClassTribIS"`
	VBCIS        *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vBCIS,omitempty" json:"vBCIS,omitempty"`
	PIS          *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pIS,omitempty" json:"pIS,omitempty"`
	AdRemIS      *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas adRemIS,omitempty" json:"adRemIS,omitempty"`
	UTrib        *string        `xml:"http://www.portalfiscal.inf.br/nfgas uTrib,omitempty" json:"uTrib,omitempty"`
	QTrib        *TDec1104OpRTC `xml:"http://www.portalfiscal.inf.br/nfgas qTrib,omitempty" json:"qTrib,omitempty"`
	VIS          *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIS,omitempty" json:"vIS,omitempty"`
}

type TISTot struct {
	VIS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIS" json:"vIS"`
}

type TIe string

type TIeDest string

type TImpICMS00 struct {
	CST   string          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	VBC   TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	PICMS TDec0302        `xml:"http://www.portalfiscal.inf.br/nfgas pICMS" json:"pICMS"`
	VICMS TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vICMS" json:"vICMS"`
	PFCP  *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP  *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpICMS10 struct {
	CST     string          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	ModBC   string          `xml:"http://www.portalfiscal.inf.br/nfgas modBC" json:"modBC"`
	VBC     TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	PICMS   TDec0302        `xml:"http://www.portalfiscal.inf.br/nfgas pICMS" json:"pICMS"`
	VICMS   TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vICMS" json:"vICMS"`
	PFCP    *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP    *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vFCP,omitempty" json:"vFCP,omitempty"`
	ModBCST string          `xml:"http://www.portalfiscal.inf.br/nfgas modBCST" json:"modBCST"`
	PMVAST  *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pMVAST,omitempty" json:"pMVAST,omitempty"`
	VBCST   TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vBCST" json:"vBCST"`
	PICMSST TDec0302        `xml:"http://www.portalfiscal.inf.br/nfgas pICMSST" json:"pICMSST"`
	VICMSST TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vICMSST" json:"vICMSST"`
	PFCPST  *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pFCPST,omitempty" json:"pFCPST,omitempty"`
	VFCPST  *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vFCPST,omitempty" json:"vFCPST,omitempty"`
}

type TImpICMS20 struct {
	CST        string          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	PRedBC     TDec0302Opc     `xml:"http://www.portalfiscal.inf.br/nfgas pRedBC" json:"pRedBC"`
	VBC        TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	PICMS      TDec0302        `xml:"http://www.portalfiscal.inf.br/nfgas pICMS" json:"pICMS"`
	VICMS      TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vICMS" json:"vICMS"`
	VICMSDeson *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string         `xml:"http://www.portalfiscal.inf.br/nfgas cBenef,omitempty" json:"cBenef,omitempty"`
	PFCP       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpICMS40 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/nfgas cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS51 struct {
	CST      string   `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	VBC      TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	PICMS    TDec0302 `xml:"http://www.portalfiscal.inf.br/nfgas pICMS" json:"pICMS"`
	VICMS    TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vICMS" json:"vICMS"`
	PDif     TDec0302 `xml:"http://www.portalfiscal.inf.br/nfgas pDif" json:"pDif"`
	VICMSOp  TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vICMSOp" json:"vICMSOp"`
	VICMSDif TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vICMSDif" json:"vICMSDif"`
	CBenef   *string  `xml:"http://www.portalfiscal.inf.br/nfgas cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS60 struct {
	CST             string          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	VBCSTRet        *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vBCSTRet,omitempty" json:"vBCSTRet,omitempty"`
	PICMSSTRet      *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pICMSSTRet,omitempty" json:"pICMSSTRet,omitempty"`
	VICMSSubstituto *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMSSubstituto,omitempty" json:"vICMSSubstituto,omitempty"`
	VICMSSTRet      *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMSSTRet,omitempty" json:"vICMSSTRet,omitempty"`
	VBCFCPSTRet     *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vBCFCPSTRet,omitempty" json:"vBCFCPSTRet,omitempty"`
	PFCPSTRet       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pFCPSTRet,omitempty" json:"pFCPSTRet,omitempty"`
	VFCPSTRet       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vFCPSTRet,omitempty" json:"vFCPSTRet,omitempty"`
	PRedBCEfet      *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pRedBCEfet,omitempty" json:"pRedBCEfet,omitempty"`
	VBCEfet         *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vBCEfet,omitempty" json:"vBCEfet,omitempty"`
	PICMSEfet       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pICMSEfet,omitempty" json:"pICMSEfet,omitempty"`
	VICMSEfet       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMSEfet,omitempty" json:"vICMSEfet,omitempty"`
	VICMSDeson      *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef          *string         `xml:"http://www.portalfiscal.inf.br/nfgas cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS70 struct {
	CST           string          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	ModBC         string          `xml:"http://www.portalfiscal.inf.br/nfgas modBC" json:"modBC"`
	PRedBC        TDec0302a04     `xml:"http://www.portalfiscal.inf.br/nfgas pRedBC" json:"pRedBC"`
	VBC           TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	PICMS         TDec0302a04     `xml:"http://www.portalfiscal.inf.br/nfgas pICMS" json:"pICMS"`
	VICMS         TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vICMS" json:"vICMS"`
	VBCFCP        *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vBCFCP,omitempty" json:"vBCFCP,omitempty"`
	PFCP          *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP          *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vFCP,omitempty" json:"vFCP,omitempty"`
	ModBCST       string          `xml:"http://www.portalfiscal.inf.br/nfgas modBCST" json:"modBCST"`
	PMVAST        *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pMVAST,omitempty" json:"pMVAST,omitempty"`
	PRedBCST      *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pRedBCST,omitempty" json:"pRedBCST,omitempty"`
	VBCST         TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vBCST" json:"vBCST"`
	PICMSST       TDec0302a04     `xml:"http://www.portalfiscal.inf.br/nfgas pICMSST" json:"pICMSST"`
	VICMSST       TDec1302        `xml:"http://www.portalfiscal.inf.br/nfgas vICMSST" json:"vICMSST"`
	VBCFCPST      *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vBCFCPST,omitempty" json:"vBCFCPST,omitempty"`
	PFCPST        *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pFCPST,omitempty" json:"pFCPST,omitempty"`
	VFCPST        *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vFCPST,omitempty" json:"vFCPST,omitempty"`
	VICMSDeson    *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	MotDesICMS    *string         `xml:"http://www.portalfiscal.inf.br/nfgas motDesICMS,omitempty" json:"motDesICMS,omitempty"`
	IndDeduzDeson *string         `xml:"http://www.portalfiscal.inf.br/nfgas indDeduzDeson,omitempty" json:"indDeduzDeson,omitempty"`
	VICMSSTDeson  *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMSSTDeson,omitempty" json:"vICMSSTDeson,omitempty"`
	MotDesICMSST  *string         `xml:"http://www.portalfiscal.inf.br/nfgas motDesICMSST,omitempty" json:"motDesICMSST,omitempty"`
}

type TImpICMS90 struct {
	CST        string          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	VBC        *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vBC,omitempty" json:"vBC,omitempty"`
	PICMS      *TDec0302       `xml:"http://www.portalfiscal.inf.br/nfgas pICMS,omitempty" json:"pICMS,omitempty"`
	VICMS      *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMS,omitempty" json:"vICMS,omitempty"`
	VICMSDeson *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string         `xml:"http://www.portalfiscal.inf.br/nfgas cBenef,omitempty" json:"cBenef,omitempty"`
	PFCP       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nfgas pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nfgas vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpGICMSPrevistoPagtoAntecip struct {
	VICMSPrevisto TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vICMSPrevisto" json:"vICMSPrevisto"`
}

type TImpPIS struct {
	CST  string      `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	VBC  TDec1302    `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	PPIS TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfgas pPIS" json:"pPIS"`
	VPIS TDec1302    `xml:"http://www.portalfiscal.inf.br/nfgas vPIS" json:"vPIS"`
}

type TImpCOFINS struct {
	CST     string      `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	VBC     TDec1302    `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	PCOFINS TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfgas pCOFINS" json:"pCOFINS"`
	VCOFINS TDec1302    `xml:"http://www.portalfiscal.inf.br/nfgas vCOFINS" json:"vCOFINS"`
}

type TImpRetTrib struct {
	VRetPIS    TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfgas vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfgas vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfgas vRetCSLL" json:"vRetCSLL"`
	VIRRF      TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfgas vIRRF" json:"vIRRF"`
}

type TImpTxReg struct {
	VBC   TDec1302    `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	PTaxa TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfgas pTaxa" json:"pTaxa"`
	VTaxa TDec1302    `xml:"http://www.portalfiscal.inf.br/nfgas vTaxa" json:"vTaxa"`
}

type TImp struct {
	Orig                      *string                        `xml:"http://www.portalfiscal.inf.br/nfgas orig,omitempty" json:"orig,omitempty"`
	ICMS00                    *TImpICMS00                    `xml:"http://www.portalfiscal.inf.br/nfgas ICMS00,omitempty" json:"ICMS00,omitempty"`
	ICMS10                    *TImpICMS10                    `xml:"http://www.portalfiscal.inf.br/nfgas ICMS10,omitempty" json:"ICMS10,omitempty"`
	ICMS20                    *TImpICMS20                    `xml:"http://www.portalfiscal.inf.br/nfgas ICMS20,omitempty" json:"ICMS20,omitempty"`
	ICMS40                    *TImpICMS40                    `xml:"http://www.portalfiscal.inf.br/nfgas ICMS40,omitempty" json:"ICMS40,omitempty"`
	ICMS51                    *TImpICMS51                    `xml:"http://www.portalfiscal.inf.br/nfgas ICMS51,omitempty" json:"ICMS51,omitempty"`
	ICMS60                    *TImpICMS60                    `xml:"http://www.portalfiscal.inf.br/nfgas ICMS60,omitempty" json:"ICMS60,omitempty"`
	ICMS70                    *TImpICMS70                    `xml:"http://www.portalfiscal.inf.br/nfgas ICMS70,omitempty" json:"ICMS70,omitempty"`
	ICMS90                    *TImpICMS90                    `xml:"http://www.portalfiscal.inf.br/nfgas ICMS90,omitempty" json:"ICMS90,omitempty"`
	GICMSPrevistoPagtoAntecip *TImpGICMSPrevistoPagtoAntecip `xml:"http://www.portalfiscal.inf.br/nfgas gICMSPrevistoPagtoAntecip,omitempty" json:"gICMSPrevistoPagtoAntecip,omitempty"`
	IndSemCST                 *string                        `xml:"http://www.portalfiscal.inf.br/nfgas indSemCST,omitempty" json:"indSemCST,omitempty"`
	IBSCBS                    TTribNFGas                     `xml:"http://www.portalfiscal.inf.br/nfgas IBSCBS" json:"IBSCBS"`
	PIS                       *TImpPIS                       `xml:"http://www.portalfiscal.inf.br/nfgas PIS,omitempty" json:"PIS,omitempty"`
	COFINS                    *TImpCOFINS                    `xml:"http://www.portalfiscal.inf.br/nfgas COFINS,omitempty" json:"COFINS,omitempty"`
	RetTrib                   *TImpRetTrib                   `xml:"http://www.portalfiscal.inf.br/nfgas retTrib,omitempty" json:"retTrib,omitempty"`
	TxReg                     *TImpTxReg                     `xml:"http://www.portalfiscal.inf.br/nfgas TxReg,omitempty" json:"TxReg,omitempty"`
}

type TIndDoacao string

type TIndIEDest string

type TJust string

type TLatitude string

type TLigacao string

type TLongitude string

type TMed string

type TMedida struct {
	UMed TUMed        `xml:"http://www.portalfiscal.inf.br/nfgas uMed" json:"uMed"`
	VMed TDec13021304 `xml:"http://www.portalfiscal.inf.br/nfgas vMed" json:"vMed"`
}

type TModNF string

type TModTar string

type TMonofasiaGMonoPadrao struct {
	QBCMono  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfgas qBCMono" json:"qBCMono"`
	AdRemIBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas adRemIBS" json:"adRemIBS"`
	AdRemCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas adRemCBS" json:"adRemCBS"`
	VIBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMono" json:"vIBSMono"`
	VCBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vCBSMono" json:"vCBSMono"`
}

type TMonofasiaGMonoReten struct {
	QBCMonoReten  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfgas qBCMonoReten" json:"qBCMonoReten"`
	AdRemIBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas adRemIBSReten" json:"adRemIBSReten"`
	VIBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMonoReten" json:"vIBSMonoReten"`
	AdRemCBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas adRemCBSReten" json:"adRemCBSReten"`
	VCBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vCBSMonoReten" json:"vCBSMonoReten"`
}

type TMonofasiaGMonoRet struct {
	QBCMonoRet  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfgas qBCMonoRet" json:"qBCMonoRet"`
	AdRemIBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas adRemIBSRet" json:"adRemIBSRet"`
	VIBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMonoRet" json:"vIBSMonoRet"`
	AdRemCBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas adRemCBSRet" json:"adRemCBSRet"`
	VCBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vCBSMonoRet" json:"vCBSMonoRet"`
}

type TMonofasiaGMonoDif struct {
	PDifIBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pDifIBS" json:"pDifIBS"`
	VIBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMonoDif" json:"vIBSMonoDif"`
	PDifCBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pDifCBS" json:"pDifCBS"`
	VCBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vCBSMonoDif" json:"vCBSMonoDif"`
}

type TMonofasia struct {
	GMonoPadrao     *TMonofasiaGMonoPadrao `xml:"http://www.portalfiscal.inf.br/nfgas gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten      *TMonofasiaGMonoReten  `xml:"http://www.portalfiscal.inf.br/nfgas gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet        *TMonofasiaGMonoRet    `xml:"http://www.portalfiscal.inf.br/nfgas gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GMonoDif        *TMonofasiaGMonoDif    `xml:"http://www.portalfiscal.inf.br/nfgas gMonoDif,omitempty" json:"gMonoDif,omitempty"`
	VTotIBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/nfgas vTotIBSMonoItem" json:"vTotIBSMonoItem"`
	VTotCBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/nfgas vTotCBSMonoItem" json:"vTotCBSMonoItem"`
}

type TMotNaoLeitura string

type TMotSub string

type TMotivo string

type TNF string

type TNFGasInfNFGasIde struct {
	CUF          TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/nfgas cUF" json:"cUF"`
	TpAmb        TAmb                `xml:"http://www.portalfiscal.inf.br/nfgas tpAmb" json:"tpAmb"`
	Mod          TModNF              `xml:"http://www.portalfiscal.inf.br/nfgas mod" json:"mod"`
	Serie        TSerie              `xml:"http://www.portalfiscal.inf.br/nfgas serie" json:"serie"`
	NNF          TNF                 `xml:"http://www.portalfiscal.inf.br/nfgas nNF" json:"nNF"`
	CNF          string              `xml:"http://www.portalfiscal.inf.br/nfgas cNF" json:"cNF"`
	CDV          string              `xml:"http://www.portalfiscal.inf.br/nfgas cDV" json:"cDV"`
	DhEmi        TDateTimeUTC        `xml:"http://www.portalfiscal.inf.br/nfgas dhEmi" json:"dhEmi"`
	TpEmis       TtpEmis             `xml:"http://www.portalfiscal.inf.br/nfgas tpEmis" json:"tpEmis"`
	NSiteAutoriz string              `xml:"http://www.portalfiscal.inf.br/nfgas nSiteAutoriz" json:"nSiteAutoriz"`
	CMunFG       TCodMunIBGE         `xml:"http://www.portalfiscal.inf.br/nfgas cMunFG" json:"cMunFG"`
	FinNFGas     TFinNF              `xml:"http://www.portalfiscal.inf.br/nfgas finNFGas" json:"finNFGas"`
	TpFat        TFatNF              `xml:"http://www.portalfiscal.inf.br/nfgas tpFat" json:"tpFat"`
	VerProc      string              `xml:"http://www.portalfiscal.inf.br/nfgas verProc" json:"verProc"`
	DhCont       *TDateTimeUTC       `xml:"http://www.portalfiscal.inf.br/nfgas dhCont,omitempty" json:"dhCont,omitempty"`
	XJust        *string             `xml:"http://www.portalfiscal.inf.br/nfgas xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov   *TCompraGovReduzido `xml:"http://www.portalfiscal.inf.br/nfgas gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	TpPagAnt     *TPagAnt            `xml:"http://www.portalfiscal.inf.br/nfgas tpPagAnt,omitempty" json:"tpPagAnt,omitempty"`
}

type TNFGasInfNFGasEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/nfgas CNPJ" json:"CNPJ"`
	IE        TIe      `xml:"http://www.portalfiscal.inf.br/nfgas IE" json:"IE"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/nfgas xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/nfgas xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/nfgas enderEmit" json:"enderEmit"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/nfgas ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TNFGasInfNFGasDest struct {
	XNome          string    `xml:"http://www.portalfiscal.inf.br/nfgas xNome" json:"xNome"`
	CNPJ           *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/nfgas CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF            *TCPF     `xml:"http://www.portalfiscal.inf.br/nfgas CPF,omitempty" json:"CPF,omitempty"`
	IDOutros       *string   `xml:"http://www.portalfiscal.inf.br/nfgas idOutros,omitempty" json:"idOutros,omitempty"`
	IE             *TIeDest  `xml:"http://www.portalfiscal.inf.br/nfgas IE,omitempty" json:"IE,omitempty"`
	IM             *string   `xml:"http://www.portalfiscal.inf.br/nfgas IM,omitempty" json:"IM,omitempty"`
	CNIS           *string   `xml:"http://www.portalfiscal.inf.br/nfgas cNIS,omitempty" json:"cNIS,omitempty"`
	NB             *string   `xml:"http://www.portalfiscal.inf.br/nfgas NB,omitempty" json:"NB,omitempty"`
	XNomeAdicional *string   `xml:"http://www.portalfiscal.inf.br/nfgas xNomeAdicional,omitempty" json:"xNomeAdicional,omitempty"`
	EnderDest      TEndeEmi  `xml:"http://www.portalfiscal.inf.br/nfgas enderDest" json:"enderDest"`
}

type TNFGasInfNFGasInstalacao struct {
	IDInstalacao      string      `xml:"http://www.portalfiscal.inf.br/nfgas idInstalacao" json:"idInstalacao"`
	IDCodCliente      *string     `xml:"http://www.portalfiscal.inf.br/nfgas idCodCliente,omitempty" json:"idCodCliente,omitempty"`
	TpInstalacao      TLigacao    `xml:"http://www.portalfiscal.inf.br/nfgas tpInstalacao" json:"tpInstalacao"`
	NContrato         *string     `xml:"http://www.portalfiscal.inf.br/nfgas nContrato,omitempty" json:"nContrato,omitempty"`
	TpClasse          TClasse     `xml:"http://www.portalfiscal.inf.br/nfgas tpClasse" json:"tpClasse"`
	XClasse           *string     `xml:"http://www.portalfiscal.inf.br/nfgas xClasse,omitempty" json:"xClasse,omitempty"`
	LatGPS            *TLatitude  `xml:"http://www.portalfiscal.inf.br/nfgas latGPS,omitempty" json:"latGPS,omitempty"`
	LongGPS           *TLongitude `xml:"http://www.portalfiscal.inf.br/nfgas longGPS,omitempty" json:"longGPS,omitempty"`
	CodRoteiroLeitura *string     `xml:"http://www.portalfiscal.inf.br/nfgas codRoteiroLeitura,omitempty" json:"codRoteiroLeitura,omitempty"`
}

type TNFGasInfNFGasGSubGNF struct {
	CNPJ       TCNPJ   `xml:"http://www.portalfiscal.inf.br/nfgas CNPJ" json:"CNPJ"`
	Serie      string  `xml:"http://www.portalfiscal.inf.br/nfgas serie" json:"serie"`
	NNF        TNF     `xml:"http://www.portalfiscal.inf.br/nfgas nNF" json:"nNF"`
	CompetEmis TCompet `xml:"http://www.portalfiscal.inf.br/nfgas CompetEmis" json:"CompetEmis"`
	CompetApur TCompet `xml:"http://www.portalfiscal.inf.br/nfgas CompetApur" json:"CompetApur"`
	Hash115    *string `xml:"http://www.portalfiscal.inf.br/nfgas hash115,omitempty" json:"hash115,omitempty"`
}

type TNFGasInfNFGasGSub struct {
	ChNFGas *TChDFe                `xml:"http://www.portalfiscal.inf.br/nfgas chNFGas,omitempty" json:"chNFGas,omitempty"`
	GNF     *TNFGasInfNFGasGSubGNF `xml:"http://www.portalfiscal.inf.br/nfgas gNF,omitempty" json:"gNF,omitempty"`
	MotSub  TMotSub                `xml:"http://www.portalfiscal.inf.br/nfgas motSub" json:"motSub"`
}

type TNFGasInfNFGasGVolContrat struct {
	NContrat     string       `xml:"nContrat,attr" json:"nContrat"`
	TpVolContrat TGrContrat   `xml:"http://www.portalfiscal.inf.br/nfgas tpVolContrat" json:"tpVolContrat"`
	QUnidContrat TDec13021306 `xml:"http://www.portalfiscal.inf.br/nfgas qUnidContrat" json:"qUnidContrat"`
}

type TNFGasInfNFGasGMed struct {
	NMed      string       `xml:"nMed,attr" json:"nMed"`
	IDEqp     string       `xml:"http://www.portalfiscal.inf.br/nfgas idEqp" json:"idEqp"`
	DMedAnt   TData        `xml:"http://www.portalfiscal.inf.br/nfgas dMedAnt" json:"dMedAnt"`
	VMedAnt   TDec13021304 `xml:"http://www.portalfiscal.inf.br/nfgas vMedAnt" json:"vMedAnt"`
	DMedAtu   TData        `xml:"http://www.portalfiscal.inf.br/nfgas dMedAtu" json:"dMedAtu"`
	VMedAtu   TDec13021304 `xml:"http://www.portalfiscal.inf.br/nfgas vMedAtu" json:"vMedAtu"`
	TpEqp     string       `xml:"http://www.portalfiscal.inf.br/nfgas tpEqp" json:"tpEqp"`
	TpMedidor *string      `xml:"http://www.portalfiscal.inf.br/nfgas tpMedidor,omitempty" json:"tpMedidor,omitempty"`
}

type TNFGasInfNFGasDetGNormalGTarif struct {
	DIniTarif   TData  `xml:"http://www.portalfiscal.inf.br/nfgas dIniTarif" json:"dIniTarif"`
	DFimTarif   *TData `xml:"http://www.portalfiscal.inf.br/nfgas dFimTarif,omitempty" json:"dFimTarif,omitempty"`
	NAto        string `xml:"http://www.portalfiscal.inf.br/nfgas nAto" json:"nAto"`
	AnoAto      string `xml:"http://www.portalfiscal.inf.br/nfgas anoAto" json:"anoAto"`
	TpFaixaCons TFaixa `xml:"http://www.portalfiscal.inf.br/nfgas tpFaixaCons" json:"tpFaixaCons"`
	VTarifAplic string `xml:"http://www.portalfiscal.inf.br/nfgas vTarifAplic" json:"vTarifAplic"`
}

type TNFGasInfNFGasDetGNormalProdGMedicao struct {
	NMed            string          `xml:"http://www.portalfiscal.inf.br/nfgas nMed" json:"nMed"`
	NContrat        *string         `xml:"http://www.portalfiscal.inf.br/nfgas nContrat,omitempty" json:"nContrat,omitempty"`
	GMedida         *TMedida        `xml:"http://www.portalfiscal.inf.br/nfgas gMedida,omitempty" json:"gMedida,omitempty"`
	TpMotNaoLeitura *TMotNaoLeitura `xml:"http://www.portalfiscal.inf.br/nfgas tpMotNaoLeitura,omitempty" json:"tpMotNaoLeitura,omitempty"`
	XMotNaoLeitura  *string         `xml:"http://www.portalfiscal.inf.br/nfgas xMotNaoLeitura,omitempty" json:"xMotNaoLeitura,omitempty"`
}

type TNFGasInfNFGasDetGNormalProdGPagAntecipado struct {
	ChDFePagAnt TChDFe  `xml:"http://www.portalfiscal.inf.br/nfgas chDFePagAnt" json:"chDFePagAnt"`
	NItemPagAnt *string `xml:"http://www.portalfiscal.inf.br/nfgas nItemPagAnt,omitempty" json:"nItemPagAnt,omitempty"`
}

type TNFGasInfNFGasDetGNormalProd struct {
	IndOrigemQtd   TOrigemQtd                                  `xml:"http://www.portalfiscal.inf.br/nfgas indOrigemQtd" json:"indOrigemQtd"`
	GMedicao       *TNFGasInfNFGasDetGNormalProdGMedicao       `xml:"http://www.portalfiscal.inf.br/nfgas gMedicao,omitempty" json:"gMedicao,omitempty"`
	CProd          string                                      `xml:"http://www.portalfiscal.inf.br/nfgas cProd" json:"cProd"`
	XProd          string                                      `xml:"http://www.portalfiscal.inf.br/nfgas xProd" json:"xProd"`
	CClass         string                                      `xml:"http://www.portalfiscal.inf.br/nfgas cClass" json:"cClass"`
	CFOP           *TCFOP                                      `xml:"http://www.portalfiscal.inf.br/nfgas CFOP,omitempty" json:"CFOP,omitempty"`
	UMed           TUMedItem                                   `xml:"http://www.portalfiscal.inf.br/nfgas uMed" json:"uMed"`
	QFaturada      TDec11001104                                `xml:"http://www.portalfiscal.inf.br/nfgas qFaturada" json:"qFaturada"`
	VItem          TDec13021310                                `xml:"http://www.portalfiscal.inf.br/nfgas vItem" json:"vItem"`
	FatorPCS       *TDec11001104                               `xml:"http://www.portalfiscal.inf.br/nfgas fatorPCS,omitempty" json:"fatorPCS,omitempty"`
	FatorPTZ       *TDec11001104                               `xml:"http://www.portalfiscal.inf.br/nfgas fatorPTZ,omitempty" json:"fatorPTZ,omitempty"`
	FatorP         *TDec11001104                               `xml:"http://www.portalfiscal.inf.br/nfgas fatorP,omitempty" json:"fatorP,omitempty"`
	FatorT         *TDec11001104                               `xml:"http://www.portalfiscal.inf.br/nfgas fatorT,omitempty" json:"fatorT,omitempty"`
	VProd          TDec13021310                                `xml:"http://www.portalfiscal.inf.br/nfgas vProd" json:"vProd"`
	VItemLiq       *TDec13021310                               `xml:"http://www.portalfiscal.inf.br/nfgas vItemLiq,omitempty" json:"vItemLiq,omitempty"`
	VProdLiq       *TDec13021310                               `xml:"http://www.portalfiscal.inf.br/nfgas vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	IndDevolucao   *string                                     `xml:"http://www.portalfiscal.inf.br/nfgas indDevolucao,omitempty" json:"indDevolucao,omitempty"`
	GPagAntecipado *TNFGasInfNFGasDetGNormalProdGPagAntecipado `xml:"http://www.portalfiscal.inf.br/nfgas gPagAntecipado,omitempty" json:"gPagAntecipado,omitempty"`
}

type TNFGasInfNFGasDetGNormalGProcRefGIBSCBS struct {
	VBC     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	VIBSUF  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSUF" json:"vIBSUF"`
	VIBSMun TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSMun" json:"vIBSMun"`
	VCBS    TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBS" json:"vCBS"`
}

type TNFGasInfNFGasDetGNormalGProcRefGProc struct {
	TpProc    TProcesso `xml:"http://www.portalfiscal.inf.br/nfgas tpProc" json:"tpProc"`
	NProcesso string    `xml:"http://www.portalfiscal.inf.br/nfgas nProcesso" json:"nProcesso"`
}

type TNFGasInfNFGasDetGNormalGProcRef struct {
	VItem        TDec13021310                             `xml:"http://www.portalfiscal.inf.br/nfgas vItem" json:"vItem"`
	QFaturada    TDec11001104                             `xml:"http://www.portalfiscal.inf.br/nfgas qFaturada" json:"qFaturada"`
	VProd        TDec13021310                             `xml:"http://www.portalfiscal.inf.br/nfgas vProd" json:"vProd"`
	IndDevolucao *string                                  `xml:"http://www.portalfiscal.inf.br/nfgas indDevolucao,omitempty" json:"indDevolucao,omitempty"`
	GIBSCBS      *TNFGasInfNFGasDetGNormalGProcRefGIBSCBS `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GProc        []TNFGasInfNFGasDetGNormalGProcRefGProc  `xml:"http://www.portalfiscal.inf.br/nfgas gProc,omitempty" json:"gProc,omitempty"`
}

type TNFGasInfNFGasDetGNormal struct {
	GTarif    []TNFGasInfNFGasDetGNormalGTarif  `xml:"http://www.portalfiscal.inf.br/nfgas gTarif,omitempty" json:"gTarif,omitempty"`
	Prod      TNFGasInfNFGasDetGNormalProd      `xml:"http://www.portalfiscal.inf.br/nfgas prod" json:"prod"`
	Imposto   TImp                              `xml:"http://www.portalfiscal.inf.br/nfgas imposto" json:"imposto"`
	GProcRef  *TNFGasInfNFGasDetGNormalGProcRef `xml:"http://www.portalfiscal.inf.br/nfgas gProcRef,omitempty" json:"gProcRef,omitempty"`
	InfAdProd *string                           `xml:"http://www.portalfiscal.inf.br/nfgas infAdProd,omitempty" json:"infAdProd,omitempty"`
}

type TNFGasInfNFGasDetGAgregadora struct {
	CClass  string   `xml:"http://www.portalfiscal.inf.br/nfgas cClass" json:"cClass"`
	VTotDFe TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vTotDFe" json:"vTotDFe"`
}

type TNFGasInfNFGasDet struct {
	NItem       string                        `xml:"nItem,attr" json:"nItem"`
	ChNFGasAnt  *TChDFe                       `xml:"chNFGasAnt,attr,omitempty" json:"chNFGasAnt,omitempty"`
	NItemAnt    *string                       `xml:"nItemAnt,attr,omitempty" json:"nItemAnt,omitempty"`
	GNormal     *TNFGasInfNFGasDetGNormal     `xml:"http://www.portalfiscal.inf.br/nfgas gNormal,omitempty" json:"gNormal,omitempty"`
	GAgregadora *TNFGasInfNFGasDetGAgregadora `xml:"http://www.portalfiscal.inf.br/nfgas gAgregadora,omitempty" json:"gAgregadora,omitempty"`
}

type TNFGasInfNFGasTotalICMSTot struct {
	VBC        TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vBC" json:"vBC"`
	VICMS      TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vICMS" json:"vICMS"`
	VICMSDeson TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vICMSDeson" json:"vICMSDeson"`
	VFCP       TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vFCP" json:"vFCP"`
	VBCST      TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vBCST" json:"vBCST"`
	VST        TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vST" json:"vST"`
	VFCPST     TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vFCPST" json:"vFCPST"`
}

type TNFGasInfNFGasTotalVRetTribTot struct {
	VRetPIS    TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vRetCSLL" json:"vRetCSLL"`
	VIRRF      TDec1302 `xml:"http://www.portalfiscal.inf.br/nfgas vIRRF" json:"vIRRF"`
}

type TNFGasInfNFGasTotal struct {
	VProd       TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfgas vProd" json:"vProd"`
	VProdLiq    *TDec1302                      `xml:"http://www.portalfiscal.inf.br/nfgas vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	ICMSTot     TNFGasInfNFGasTotalICMSTot     `xml:"http://www.portalfiscal.inf.br/nfgas ICMSTot" json:"ICMSTot"`
	VRetTribTot TNFGasInfNFGasTotalVRetTribTot `xml:"http://www.portalfiscal.inf.br/nfgas vRetTribTot" json:"vRetTribTot"`
	VCOFINS     TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfgas vCOFINS" json:"vCOFINS"`
	VPIS        TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfgas vPIS" json:"vPIS"`
	VTxReg      TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfgas vTxReg" json:"vTxReg"`
	VNF         TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfgas vNF" json:"vNF"`
	IBSCBSTot   TIBSCBSTot                     `xml:"http://www.portalfiscal.inf.br/nfgas IBSCBSTot" json:"IBSCBSTot"`
	VTotDFe     TDec1302                       `xml:"http://www.portalfiscal.inf.br/nfgas vTotDFe" json:"vTotDFe"`
}

type TNFGasInfNFGasPgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/nfgas pgto,omitempty" json:"pgto,omitempty"`
}

type TNFGasInfNFGasGFatGPIX struct {
	URLQRCodePIX string `xml:"http://www.portalfiscal.inf.br/nfgas urlQRCodePIX" json:"urlQRCodePIX"`
}

type TNFGasInfNFGasGFat struct {
	CompetFat    TCompet                 `xml:"http://www.portalfiscal.inf.br/nfgas CompetFat" json:"CompetFat"`
	DVencFat     TData                   `xml:"http://www.portalfiscal.inf.br/nfgas dVencFat" json:"dVencFat"`
	DApresFat    *TData                  `xml:"http://www.portalfiscal.inf.br/nfgas dApresFat,omitempty" json:"dApresFat,omitempty"`
	DProxLeitura TData                   `xml:"http://www.portalfiscal.inf.br/nfgas dProxLeitura" json:"dProxLeitura"`
	NFat         *string                 `xml:"http://www.portalfiscal.inf.br/nfgas nFat,omitempty" json:"nFat,omitempty"`
	CodBarras    string                  `xml:"http://www.portalfiscal.inf.br/nfgas codBarras" json:"codBarras"`
	CodDebAuto   *string                 `xml:"http://www.portalfiscal.inf.br/nfgas codDebAuto,omitempty" json:"codDebAuto,omitempty"`
	CodBanco     *string                 `xml:"http://www.portalfiscal.inf.br/nfgas codBanco,omitempty" json:"codBanco,omitempty"`
	CodAgencia   *string                 `xml:"http://www.portalfiscal.inf.br/nfgas codAgencia,omitempty" json:"codAgencia,omitempty"`
	EnderCorresp *TEndeEmi               `xml:"http://www.portalfiscal.inf.br/nfgas enderCorresp,omitempty" json:"enderCorresp,omitempty"`
	GPIX         *TNFGasInfNFGasGFatGPIX `xml:"http://www.portalfiscal.inf.br/nfgas gPIX,omitempty" json:"gPIX,omitempty"`
	InfAdFat     *string                 `xml:"http://www.portalfiscal.inf.br/nfgas infAdFat,omitempty" json:"infAdFat,omitempty"`
}

type TNFGasInfNFGasGAgenciaGHistConsGCons struct {
	CompetFat TCompet       `xml:"http://www.portalfiscal.inf.br/nfgas CompetFat" json:"CompetFat"`
	UMed      TUMed         `xml:"http://www.portalfiscal.inf.br/nfgas uMed" json:"uMed"`
	QtdDias   string        `xml:"http://www.portalfiscal.inf.br/nfgas qtdDias" json:"qtdDias"`
	MedDiaria *TDec13021304 `xml:"http://www.portalfiscal.inf.br/nfgas medDiaria,omitempty" json:"medDiaria,omitempty"`
	Consumo   *TDec13021304 `xml:"http://www.portalfiscal.inf.br/nfgas consumo,omitempty" json:"consumo,omitempty"`
	VFat      TDec13021304  `xml:"http://www.portalfiscal.inf.br/nfgas vFat" json:"vFat"`
}

type TNFGasInfNFGasGAgenciaGHistCons struct {
	XHistorico string                                 `xml:"http://www.portalfiscal.inf.br/nfgas xHistorico" json:"xHistorico"`
	GCons      []TNFGasInfNFGasGAgenciaGHistConsGCons `xml:"http://www.portalfiscal.inf.br/nfgas gCons,omitempty" json:"gCons,omitempty"`
	MedMensal  TDec13021304                           `xml:"http://www.portalfiscal.inf.br/nfgas medMensal" json:"medMensal"`
}

type TNFGasInfNFGasGAgencia struct {
	NomeAgenciaAtend  *string                           `xml:"http://www.portalfiscal.inf.br/nfgas nomeAgenciaAtend,omitempty" json:"nomeAgenciaAtend,omitempty"`
	EnderAgenciaAtend *string                           `xml:"http://www.portalfiscal.inf.br/nfgas enderAgenciaAtend,omitempty" json:"enderAgenciaAtend,omitempty"`
	SitioAgenciaAtend *string                           `xml:"http://www.portalfiscal.inf.br/nfgas sitioAgenciaAtend,omitempty" json:"sitioAgenciaAtend,omitempty"`
	GHistCons         []TNFGasInfNFGasGAgenciaGHistCons `xml:"http://www.portalfiscal.inf.br/nfgas gHistCons,omitempty" json:"gHistCons,omitempty"`
	InfAdReg          *string                           `xml:"http://www.portalfiscal.inf.br/nfgas infAdReg,omitempty" json:"infAdReg,omitempty"`
}

type TNFGasInfNFGasAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/nfgas CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/nfgas CPF,omitempty" json:"CPF,omitempty"`
}

type TNFGasInfNFGasInfAdic struct {
	InfAdFisco *string  `xml:"http://www.portalfiscal.inf.br/nfgas infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     []string `xml:"http://www.portalfiscal.inf.br/nfgas infCpl,omitempty" json:"infCpl,omitempty"`
}

type TNFGasInfNFGas struct {
	Versao      TVerNFGas                   `xml:"versao,attr" json:"versao"`
	ID          string                      `xml:"Id,attr" json:"Id"`
	Ide         TNFGasInfNFGasIde           `xml:"http://www.portalfiscal.inf.br/nfgas ide" json:"ide"`
	Emit        TNFGasInfNFGasEmit          `xml:"http://www.portalfiscal.inf.br/nfgas emit" json:"emit"`
	Dest        TNFGasInfNFGasDest          `xml:"http://www.portalfiscal.inf.br/nfgas dest" json:"dest"`
	Instalacao  *TNFGasInfNFGasInstalacao   `xml:"http://www.portalfiscal.inf.br/nfgas instalacao,omitempty" json:"instalacao,omitempty"`
	GSub        *TNFGasInfNFGasGSub         `xml:"http://www.portalfiscal.inf.br/nfgas gSub,omitempty" json:"gSub,omitempty"`
	GVolContrat []TNFGasInfNFGasGVolContrat `xml:"http://www.portalfiscal.inf.br/nfgas gVolContrat,omitempty" json:"gVolContrat,omitempty"`
	GMed        []TNFGasInfNFGasGMed        `xml:"http://www.portalfiscal.inf.br/nfgas gMed,omitempty" json:"gMed,omitempty"`
	Det         []TNFGasInfNFGasDet         `xml:"http://www.portalfiscal.inf.br/nfgas det,omitempty" json:"det,omitempty"`
	Total       TNFGasInfNFGasTotal         `xml:"http://www.portalfiscal.inf.br/nfgas total" json:"total"`
	PgtoVinc    *TNFGasInfNFGasPgtoVinc     `xml:"http://www.portalfiscal.inf.br/nfgas pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	GFat        *TNFGasInfNFGasGFat         `xml:"http://www.portalfiscal.inf.br/nfgas gFat,omitempty" json:"gFat,omitempty"`
	GAgencia    *TNFGasInfNFGasGAgencia     `xml:"http://www.portalfiscal.inf.br/nfgas gAgencia,omitempty" json:"gAgencia,omitempty"`
	AutXML      []TNFGasInfNFGasAutXML      `xml:"http://www.portalfiscal.inf.br/nfgas autXML,omitempty" json:"autXML,omitempty"`
	InfAdic     *TNFGasInfNFGasInfAdic      `xml:"http://www.portalfiscal.inf.br/nfgas infAdic,omitempty" json:"infAdic,omitempty"`
	GRespTec    *TRespTec                   `xml:"http://www.portalfiscal.inf.br/nfgas gRespTec,omitempty" json:"gRespTec,omitempty"`
}

type TNFGasInfNFGasSupl struct {
	QRCodNFGas string `xml:"http://www.portalfiscal.inf.br/nfgas qrCodNFGas" json:"qrCodNFGas"`
}

type TNFGas struct {
	InfNFGas     TNFGasInfNFGas     `xml:"http://www.portalfiscal.inf.br/nfgas infNFGas" json:"infNFGas"`
	InfNFGasSupl TNFGasInfNFGasSupl `xml:"http://www.portalfiscal.inf.br/nfgas infNFGasSupl" json:"infNFGasSupl"`
	Signature    SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TNSU string

type TOperCompraGov string

type TOrigemQtd string

type TPagAnt string

type TPagRef struct {
	RefDFe []TChDFeRTC `xml:"http://www.portalfiscal.inf.br/nfgas refDFe,omitempty" json:"refDFe,omitempty"`
}

type TPagamentoRTC struct {
	NPag        string       `xml:"nPag,attr" json:"nPag"`
	IDTransacao string       `xml:"idTransacao,attr" json:"idTransacao"`
	TpMeioPgto  string       `xml:"http://www.portalfiscal.inf.br/nfgas tpMeioPgto" json:"tpMeioPgto"`
	CNPJReceb   TCNPJRTC     `xml:"http://www.portalfiscal.inf.br/nfgas CNPJReceb" json:"CNPJReceb"`
	CNPJBasePSP TCNPJBaseRTC `xml:"http://www.portalfiscal.inf.br/nfgas CNPJBasePSP" json:"CNPJBasePSP"`
}

type TPosTarSimples string

type TPosTarifCompleto string

type TProcEmi string

type TProcesso string

type TProt string

type TProtNFGasInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/nfgas tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/nfgas verAplic" json:"verAplic"`
	ChNFGas  TChDFe           `xml:"http://www.portalfiscal.inf.br/nfgas chNFGas" json:"chNFGas"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/nfgas dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/nfgas nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/nfgas digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/nfgas cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/nfgas xMotivo" json:"xMotivo"`
}

type TProtNFGasInfFisco struct {
	CMsg string  `xml:"http://www.portalfiscal.inf.br/nfgas cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/nfgas xMsg" json:"xMsg"`
}

type TProtNFGas struct {
	Versao    TVerNFGas           `xml:"versao,attr" json:"versao"`
	InfProt   TProtNFGasInfProt   `xml:"http://www.portalfiscal.inf.br/nfgas infProt" json:"infProt"`
	InfFisco  *TProtNFGasInfFisco `xml:"http://www.portalfiscal.inf.br/nfgas infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRBSN string

type TRec string

type TRed struct {
	PRedAliq  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqEfet" json:"pAliqEfet"`
}

type TRespTec struct {
	CNPJ     TCNPJ   `xml:"http://www.portalfiscal.inf.br/nfgas CNPJ" json:"CNPJ"`
	XContato string  `xml:"http://www.portalfiscal.inf.br/nfgas xContato" json:"xContato"`
	Email    TEmail  `xml:"http://www.portalfiscal.inf.br/nfgas email" json:"email"`
	Fone     string  `xml:"http://www.portalfiscal.inf.br/nfgas fone" json:"fone"`
	IDCSRT   *string `xml:"http://www.portalfiscal.inf.br/nfgas idCSRT,omitempty" json:"idCSRT,omitempty"`
	HashCSRT *string `xml:"http://www.portalfiscal.inf.br/nfgas hashCSRT,omitempty" json:"hashCSRT,omitempty"`
}

type TRetNFGas struct {
	Versao    TVerNFGas   `xml:"versao,attr" json:"versao"`
	TpAmb     TAmb        `xml:"http://www.portalfiscal.inf.br/nfgas tpAmb" json:"tpAmb"`
	CUF       TCodUFIBGE  `xml:"http://www.portalfiscal.inf.br/nfgas cUF" json:"cUF"`
	VerAplic  TVerAplic   `xml:"http://www.portalfiscal.inf.br/nfgas verAplic" json:"verAplic"`
	CStat     TStat       `xml:"http://www.portalfiscal.inf.br/nfgas cStat" json:"cStat"`
	XMotivo   TMotivo     `xml:"http://www.portalfiscal.inf.br/nfgas xMotivo" json:"xMotivo"`
	ProtNFGas *TProtNFGas `xml:"http://www.portalfiscal.inf.br/nfgas protNFGas,omitempty" json:"protNFGas,omitempty"`
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
	VRBSNTot          TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfgas vRBSNTot" json:"vRBSNTot"`
	VIBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSSNTot,omitempty" json:"vIBSSNTot,omitempty"`
	VIBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBSSNtotPendSusp,omitempty" json:"vIBSSNtotPendSusp,omitempty"`
	VCBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBSSNTot,omitempty" json:"vCBSSNTot,omitempty"`
	VCBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBSSNtotPendSusp,omitempty" json:"vCBSSNtotPendSusp,omitempty"`
}

type TTpCredPresIBSZFM string

type TTransfCred struct {
	VIBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vIBS" json:"vIBS"`
	VCBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfgas vCBS" json:"vCBS"`
}

type TTribBPe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfgas cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfgas indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCTe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfgas cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfgas indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCompraGov struct {
	PAliqIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vTribCBS" json:"vTribCBS"`
}

type TTribItemSN struct {
	NItem        string         `xml:"nItem,attr" json:"nItem"`
	VRBSNItem    TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfgas vRBSNItem" json:"vRBSNItem"`
	TpRBSN       TRBSN          `xml:"http://www.portalfiscal.inf.br/nfgas tpRBSN" json:"tpRBSN"`
	PIBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pIBSSN,omitempty" json:"pIBSSN,omitempty"`
	VIBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSSN,omitempty" json:"vIBSSN,omitempty"`
	PCBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pCBSSN,omitempty" json:"pCBSSN,omitempty"`
	VCBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vCBSSN,omitempty" json:"vCBSSN,omitempty"`
	VIBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vIBSPendSusp,omitempty" json:"vIBSPendSusp,omitempty"`
	VCBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vCBSPendSusp,omitempty" json:"vCBSPendSusp,omitempty"`
}

type TTribNF3e struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfgas cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfgas indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFAg struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfgas cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfgas indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFCe struct {
	CST         TCST        `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	CClassTrib  TcClassTrib `xml:"http://www.portalfiscal.inf.br/nfgas cClassTrib" json:"cClassTrib"`
	IndDoacao   *TIndDoacao `xml:"http://www.portalfiscal.inf.br/nfgas indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS     *TCIBSNFe   `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono *TMonofasia `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
}

type TTribNFCom struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfgas cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfgas indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFGas struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfgas cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfgas indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFe struct {
	CST             TCST             `xml:"http://www.portalfiscal.inf.br/nfgas CST" json:"CST"`
	CClassTrib      TcClassTrib      `xml:"http://www.portalfiscal.inf.br/nfgas cClassTrib" json:"cClassTrib"`
	IndDoacao       *TIndDoacao      `xml:"http://www.portalfiscal.inf.br/nfgas indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS         *TCIBSNFe        `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono     *TMonofasia      `xml:"http://www.portalfiscal.inf.br/nfgas gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
	GTransfCred     *TTransfCred     `xml:"http://www.portalfiscal.inf.br/nfgas gTransfCred,omitempty" json:"gTransfCred,omitempty"`
	GAjusteCompet   *TAjusteCompet   `xml:"http://www.portalfiscal.inf.br/nfgas gAjusteCompet,omitempty" json:"gAjusteCompet,omitempty"`
	GEstornoCred    *TEstornoCred    `xml:"http://www.portalfiscal.inf.br/nfgas gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
	GCredPresOper   *TCredPresOper   `xml:"http://www.portalfiscal.inf.br/nfgas gCredPresOper,omitempty" json:"gCredPresOper,omitempty"`
	GCredPresIBSZFM *TCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nfgas gCredPresIBSZFM,omitempty" json:"gCredPresIBSZFM,omitempty"`
}

type TTribRegular struct {
	CSTReg             TCST          `xml:"http://www.portalfiscal.inf.br/nfgas CSTReg" json:"CSTReg"`
	CClassTribReg      TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfgas cClassTribReg" json:"cClassTribReg"`
	PAliqEfetRegIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqEfetRegIBSUF" json:"pAliqEfetRegIBSUF"`
	VTribRegIBSUF      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vTribRegIBSUF" json:"vTribRegIBSUF"`
	PAliqEfetRegIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqEfetRegIBSMun" json:"pAliqEfetRegIBSMun"`
	VTribRegIBSMun     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vTribRegIBSMun" json:"vTribRegIBSMun"`
	PAliqEfetRegCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfgas pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS        TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfgas vTribRegCBS" json:"vTribRegCBS"`
}

type TUMed string

type TUMedItem string

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerNFGas string

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

type String string

// NFGas represents the global element {http://www.portalfiscal.inf.br/nfgas}NFGas.
type NFGas struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfgas NFGas" json:"-"`
	TNFGas
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
