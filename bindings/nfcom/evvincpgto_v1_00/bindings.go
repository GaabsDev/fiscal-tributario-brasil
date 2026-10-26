// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package evvincpgto_v1_00

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

type TAmb string

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

type TCST string

type TChDFe string

type TChDFeRTC string

type TCNPJ string

type TCNPJBaseRTC string

type TCNPJOpc string

type TCNPJRTC string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

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

type TEnteGov string

type TEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfcom vCBSEstCred" json:"vCBSEstCred"`
}

type TEventoInfEventoDetEvento struct {
	VersaoEvento string       `xml:"versaoEvento,attr" json:"versaoEvento"`
	Any          []RawElement `xml:",any" json:"any,omitempty"`
}

type TEventoInfEvento struct {
	ID         string                    `xml:"Id,attr" json:"Id"`
	COrgao     TCOrgaoIBGE               `xml:"http://www.portalfiscal.inf.br/nfcom cOrgao" json:"cOrgao"`
	TpAmb      TAmb                      `xml:"http://www.portalfiscal.inf.br/nfcom tpAmb" json:"tpAmb"`
	CNPJ       TCNPJ                     `xml:"http://www.portalfiscal.inf.br/nfcom CNPJ" json:"CNPJ"`
	ChNFCom    TChDFe                    `xml:"http://www.portalfiscal.inf.br/nfcom chNFCom" json:"chNFCom"`
	DhEvento   TDateTimeUTC              `xml:"http://www.portalfiscal.inf.br/nfcom dhEvento" json:"dhEvento"`
	TpEvento   string                    `xml:"http://www.portalfiscal.inf.br/nfcom tpEvento" json:"tpEvento"`
	NSeqEvento string                    `xml:"http://www.portalfiscal.inf.br/nfcom nSeqEvento" json:"nSeqEvento"`
	DetEvento  TEventoInfEventoDetEvento `xml:"http://www.portalfiscal.inf.br/nfcom detEvento" json:"detEvento"`
}

type TEvento struct {
	Versao    TVerEvento       `xml:"versao,attr" json:"versao"`
	InfEvento TEventoInfEvento `xml:"http://www.portalfiscal.inf.br/nfcom infEvento" json:"infEvento"`
	Signature SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

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

type TIe string

type TIeDest string

type TIndDoacao string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModNFCom string

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

type TMotivo string

type TNF string

type TNSU string

type TOperCompraGov string

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

type TProcEvento struct {
	Versao         TVerEvento    `xml:"versao,attr" json:"versao"`
	IpTransmissor  *TIPv4        `xml:"ipTransmissor,attr,omitempty" json:"ipTransmissor,omitempty"`
	NPortaCon      *string       `xml:"nPortaCon,attr,omitempty" json:"nPortaCon,omitempty"`
	DhConexao      *TDateTimeUTC `xml:"dhConexao,attr,omitempty" json:"dhConexao,omitempty"`
	EventoNFCom    TEvento       `xml:"http://www.portalfiscal.inf.br/nfcom eventoNFCom" json:"eventoNFCom"`
	RetEventoNFCom TRetEvento    `xml:"http://www.portalfiscal.inf.br/nfcom retEventoNFCom" json:"retEventoNFCom"`
}

type TProt string

type TRBSN string

type TRec string

type TRed struct {
	PRedAliq  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfcom pAliqEfet" json:"pAliqEfet"`
}

type TRetEventoInfEvento struct {
	ID          *string       `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb       TAmb          `xml:"http://www.portalfiscal.inf.br/nfcom tpAmb" json:"tpAmb"`
	VerAplic    TVerAplic     `xml:"http://www.portalfiscal.inf.br/nfcom verAplic" json:"verAplic"`
	COrgao      TCOrgaoIBGE   `xml:"http://www.portalfiscal.inf.br/nfcom cOrgao" json:"cOrgao"`
	CStat       TStat         `xml:"http://www.portalfiscal.inf.br/nfcom cStat" json:"cStat"`
	XMotivo     TMotivo       `xml:"http://www.portalfiscal.inf.br/nfcom xMotivo" json:"xMotivo"`
	ChNFCom     *TChDFe       `xml:"http://www.portalfiscal.inf.br/nfcom chNFCom,omitempty" json:"chNFCom,omitempty"`
	TpEvento    *string       `xml:"http://www.portalfiscal.inf.br/nfcom tpEvento,omitempty" json:"tpEvento,omitempty"`
	XEvento     *string       `xml:"http://www.portalfiscal.inf.br/nfcom xEvento,omitempty" json:"xEvento,omitempty"`
	NSeqEvento  *string       `xml:"http://www.portalfiscal.inf.br/nfcom nSeqEvento,omitempty" json:"nSeqEvento,omitempty"`
	DhRegEvento *TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfcom dhRegEvento,omitempty" json:"dhRegEvento,omitempty"`
	NProt       *TProt        `xml:"http://www.portalfiscal.inf.br/nfcom nProt,omitempty" json:"nProt,omitempty"`
}

type TRetEvento struct {
	Versao    string              `xml:"versao,attr" json:"versao"`
	InfEvento TRetEventoInfEvento `xml:"http://www.portalfiscal.inf.br/nfcom infEvento" json:"infEvento"`
	Signature *SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TStringRTC string

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

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerEvento string

type Tano string

type TcClassTrib string

type TcCredPres string

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

type EvVincPgtoContent struct {
	DescEvento string        `xml:"http://www.portalfiscal.inf.br/nfcom descEvento" json:"descEvento"`
	NProt      TProt         `xml:"http://www.portalfiscal.inf.br/nfcom nProt" json:"nProt"`
	Pgto       TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/nfcom pgto" json:"pgto"`
}

// EvVincPgto represents the global element {http://www.portalfiscal.inf.br/nfcom}evVincPgto.
type EvVincPgto struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfcom evVincPgto" json:"-"`
	EvVincPgtoContent
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
