// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package procnfag_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TALCZFMCBS struct {
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vTribRegCBS" json:"vTribRegCBS"`
}

type TALCZFMCBSNFe struct {
	TpALCZFMCBS     string        `xml:"http://www.portalfiscal.inf.br/nfag tpALCZFMCBS" json:"tpALCZFMCBS"`
	NProcSuframa    *string       `xml:"http://www.portalfiscal.inf.br/nfag nProcSuframa,omitempty" json:"nProcSuframa,omitempty"`
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vTribRegCBS" json:"vTribRegCBS"`
}

type TAcessante string

type TAjusteCompet struct {
	CompetApur TCompetApur `xml:"http://www.portalfiscal.inf.br/nfag competApur" json:"competApur"`
	VIBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBS" json:"vIBS"`
	VCBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBS" json:"vCBS"`
}

type TAmb string

type TBand string

type TCIBSGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfag gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfag gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfag gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSUF" json:"vIBSUF"`
}

type TCIBSGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfag gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfag gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfag gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSMun" json:"vIBSMun"`
}

type TCIBSGCBS struct {
	PCBS       TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pCBS" json:"pCBS"`
	GDif       *TDif         `xml:"http://www.portalfiscal.inf.br/nfag gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfag gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed         `xml:"http://www.portalfiscal.inf.br/nfag gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBS   `xml:"http://www.portalfiscal.inf.br/nfag gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vCBS" json:"vCBS"`
}

type TCIBS struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfag vBC" json:"vBC"`
	GIBSUF         TCIBSGIBSUF     `xml:"http://www.portalfiscal.inf.br/nfag gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSGIBSMun    `xml:"http://www.portalfiscal.inf.br/nfag gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfag vIBS" json:"vIBS"`
	GCBS           TCIBSGCBS       `xml:"http://www.portalfiscal.inf.br/nfag gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nfag gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfag gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCIBSNFeGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfag gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfag gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfag gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSUF" json:"vIBSUF"`
}

type TCIBSNFeGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nfag gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nfag gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nfag gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSMun" json:"vIBSMun"`
}

type TCIBSNFeGCBS struct {
	PCBS       TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfag pCBS" json:"pCBS"`
	GDif       *TDif          `xml:"http://www.portalfiscal.inf.br/nfag gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib      `xml:"http://www.portalfiscal.inf.br/nfag gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed          `xml:"http://www.portalfiscal.inf.br/nfag gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBSNFe `xml:"http://www.portalfiscal.inf.br/nfag gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfag vCBS" json:"vCBS"`
}

type TCIBSNFe struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfag vBC" json:"vBC"`
	GIBSUF         TCIBSNFeGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfag gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSNFeGIBSMun `xml:"http://www.portalfiscal.inf.br/nfag gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nfag vIBS" json:"vIBS"`
	GCBS           TCIBSNFeGCBS    `xml:"http://www.portalfiscal.inf.br/nfag gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nfag gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfag gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCOrgaoIBGE string

type TCST string

type TCategAG string

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
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nfag tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfag pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nfag tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nfag refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCompraGovReduzido struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nfag tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nfag pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nfag tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nfag refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCPF string

type TCPFVar string

type TCredPres struct {
	PCredPres        TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pCredPres" json:"pCredPres"`
	VCredPres        *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfag vCredPres,omitempty" json:"vCredPres,omitempty"`
	VCredPresCondSus *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfag vCredPresCondSus,omitempty" json:"vCredPresCondSus,omitempty"`
}

type TCredPresIBSZFM struct {
	CompetApur       TCompetApur       `xml:"http://www.portalfiscal.inf.br/nfag competApur" json:"competApur"`
	TpCredPresIBSZFM TTpCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nfag tpCredPresIBSZFM" json:"tpCredPresIBSZFM"`
	VCredPresIBSZFM  TDec1302RTC       `xml:"http://www.portalfiscal.inf.br/nfag vCredPresIBSZFM" json:"vCredPresIBSZFM"`
}

type TCredPresOper struct {
	VBCCredPres  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vBCCredPres" json:"vBCCredPres"`
	CCredPres    TcCredPres  `xml:"http://www.portalfiscal.inf.br/nfag cCredPres" json:"cCredPres"`
	GIBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nfag gIBSCredPres,omitempty" json:"gIBSCredPres,omitempty"`
	GCBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nfag gCBSCredPres,omitempty" json:"gCBSCredPres,omitempty"`
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
	PDevTrib *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pDevTrib,omitempty" json:"pDevTrib,omitempty"`
	VDevTrib TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfag vDevTrib" json:"vDevTrib"`
}

type TDif struct {
	PDif TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pDif" json:"pDif"`
	VDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vDif" json:"vDif"`
}

type TDoc string

type TEmail string

type TEmpresaSoft struct {
	CNPJ     *TCNPJ `xml:"http://www.portalfiscal.inf.br/nfag CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF      *TCPF  `xml:"http://www.portalfiscal.inf.br/nfag CPF,omitempty" json:"CPF,omitempty"`
	XContato string `xml:"http://www.portalfiscal.inf.br/nfag xContato" json:"xContato"`
	Email    TEmail `xml:"http://www.portalfiscal.inf.br/nfag email" json:"email"`
	Fone     string `xml:"http://www.portalfiscal.inf.br/nfag fone" json:"fone"`
}

type TEndeEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfag xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfag nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfag xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfag xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfag cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfag xMun" json:"xMun"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/nfag CEP" json:"CEP"`
	UF      TUFSemEX    `xml:"http://www.portalfiscal.inf.br/nfag UF" json:"UF"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfag fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nfag email,omitempty" json:"email,omitempty"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nfag xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nfag nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nfag xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nfag xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nfag cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nfag xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/nfag CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/nfag UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/nfag cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/nfag xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nfag fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nfag email,omitempty" json:"email,omitempty"`
}

type TEnteGov string

type TEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBSEstCred" json:"vCBSEstCred"`
}

type TFaixa string

type TFase string

type TFatNFAg string

type TFinNFAg string

type TGrContrat string

type TGrMed string

type TGrMedAG string

type TGrpTensao string

type TIBSCBSMonoTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSUF" json:"vIBSUF"`
}

type TIBSCBSMonoTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSMun" json:"vIBSMun"`
}

type TIBSCBSMonoTotGIBS struct {
	GIBSUF           TIBSCBSMonoTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfag gIBSUF" json:"gIBSUF"`
	GIBSMun          TIBSCBSMonoTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nfag gIBSMun" json:"gIBSMun"`
	VIBS             TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfag vIBS" json:"vIBS"`
	VCredPres        TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfag vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nfag vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGCBS struct {
	VDif             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDif" json:"vDif"`
	VDevTrib         TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDevTrib" json:"vDevTrib"`
	VCBS             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBS" json:"vCBS"`
	VCredPres        TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGMono struct {
	VIBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSMono" json:"vIBSMono"`
	VCBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBSMono" json:"vCBSMono"`
	VIBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSMonoReten" json:"vIBSMonoReten"`
	VCBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBSMonoReten" json:"vCBSMonoReten"`
	VIBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSMonoRet" json:"vIBSMonoRet"`
	VCBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBSMonoRet" json:"vCBSMonoRet"`
}

type TIBSCBSMonoTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSMonoTot struct {
	VBCIBSCBS    TDec1302RTC                 `xml:"http://www.portalfiscal.inf.br/nfag vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         *TIBSCBSMonoTotGIBS         `xml:"http://www.portalfiscal.inf.br/nfag gIBS,omitempty" json:"gIBS,omitempty"`
	GCBS         *TIBSCBSMonoTotGCBS         `xml:"http://www.portalfiscal.inf.br/nfag gCBS,omitempty" json:"gCBS,omitempty"`
	GMono        *TIBSCBSMonoTotGMono        `xml:"http://www.portalfiscal.inf.br/nfag gMono,omitempty" json:"gMono,omitempty"`
	GEstornoCred *TIBSCBSMonoTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIBSCBSTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSUF" json:"vIBSUF"`
}

type TIBSCBSTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSMun" json:"vIBSMun"`
}

type TIBSCBSTotGIBS struct {
	GIBSUF  TIBSCBSTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nfag gIBSUF" json:"gIBSUF"`
	GIBSMun TIBSCBSTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nfag gIBSMun" json:"gIBSMun"`
	VIBS    TDec1302RTC           `xml:"http://www.portalfiscal.inf.br/nfag vIBS" json:"vIBS"`
}

type TIBSCBSTotGCBS struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vDevTrib" json:"vDevTrib"`
	VCBS     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBS" json:"vCBS"`
}

type TIBSCBSTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSTot struct {
	VBCIBSCBS    TDec1302RTC             `xml:"http://www.portalfiscal.inf.br/nfag vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         TIBSCBSTotGIBS          `xml:"http://www.portalfiscal.inf.br/nfag gIBS" json:"gIBS"`
	GCBS         TIBSCBSTotGCBS          `xml:"http://www.portalfiscal.inf.br/nfag gCBS" json:"gCBS"`
	GEstornoCred *TIBSCBSTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIPv4 string

type TIS struct {
	CSTIS        TCST           `xml:"http://www.portalfiscal.inf.br/nfag CSTIS" json:"CSTIS"`
	CClassTribIS TcClassTrib    `xml:"http://www.portalfiscal.inf.br/nfag cClassTribIS" json:"cClassTribIS"`
	VBCIS        *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vBCIS,omitempty" json:"vBCIS,omitempty"`
	PIS          *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pIS,omitempty" json:"pIS,omitempty"`
	AdRemIS      *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag adRemIS,omitempty" json:"adRemIS,omitempty"`
	UTrib        *string        `xml:"http://www.portalfiscal.inf.br/nfag uTrib,omitempty" json:"uTrib,omitempty"`
	QTrib        *TDec1104OpRTC `xml:"http://www.portalfiscal.inf.br/nfag qTrib,omitempty" json:"qTrib,omitempty"`
	VIS          *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIS,omitempty" json:"vIS,omitempty"`
}

type TISTot struct {
	VIS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIS" json:"vIS"`
}

type TIe string

type TIeDest string

type TIndDoacao string

type TIndIEDest string

type TJust string

type TLatitude string

type TLigacao string

type TLongitude string

type TMed string

type TMedidaAG struct {
	TpGrMed      string        `xml:"http://www.portalfiscal.inf.br/nfag tpGrMed" json:"tpGrMed"`
	NUnidConsumo *string       `xml:"http://www.portalfiscal.inf.br/nfag nUnidConsumo,omitempty" json:"nUnidConsumo,omitempty"`
	VUnidConsumo *TDec13021304 `xml:"http://www.portalfiscal.inf.br/nfag vUnidConsumo,omitempty" json:"vUnidConsumo,omitempty"`
	UMed         TUMedItemAG   `xml:"http://www.portalfiscal.inf.br/nfag uMed" json:"uMed"`
	VMedAnt      TDec13021304  `xml:"http://www.portalfiscal.inf.br/nfag vMedAnt" json:"vMedAnt"`
	VMedAtu      TDec13021304  `xml:"http://www.portalfiscal.inf.br/nfag vMedAtu" json:"vMedAtu"`
	VConst       *TDec13021306 `xml:"http://www.portalfiscal.inf.br/nfag vConst,omitempty" json:"vConst,omitempty"`
	VMed         TDec13021304  `xml:"http://www.portalfiscal.inf.br/nfag vMed" json:"vMed"`
}

type TModNFAg string

type TModTar string

type TMonofasiaGMonoPadrao struct {
	QBCMono  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfag qBCMono" json:"qBCMono"`
	AdRemIBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag adRemIBS" json:"adRemIBS"`
	AdRemCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag adRemCBS" json:"adRemCBS"`
	VIBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSMono" json:"vIBSMono"`
	VCBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vCBSMono" json:"vCBSMono"`
}

type TMonofasiaGMonoReten struct {
	QBCMonoReten  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfag qBCMonoReten" json:"qBCMonoReten"`
	AdRemIBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag adRemIBSReten" json:"adRemIBSReten"`
	VIBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSMonoReten" json:"vIBSMonoReten"`
	AdRemCBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag adRemCBSReten" json:"adRemCBSReten"`
	VCBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vCBSMonoReten" json:"vCBSMonoReten"`
}

type TMonofasiaGMonoRet struct {
	QBCMonoRet  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nfag qBCMonoRet" json:"qBCMonoRet"`
	AdRemIBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag adRemIBSRet" json:"adRemIBSRet"`
	VIBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSMonoRet" json:"vIBSMonoRet"`
	AdRemCBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag adRemCBSRet" json:"adRemCBSRet"`
	VCBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vCBSMonoRet" json:"vCBSMonoRet"`
}

type TMonofasiaGMonoDif struct {
	PDifIBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pDifIBS" json:"pDifIBS"`
	VIBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSMonoDif" json:"vIBSMonoDif"`
	PDifCBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pDifCBS" json:"pDifCBS"`
	VCBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vCBSMonoDif" json:"vCBSMonoDif"`
}

type TMonofasia struct {
	GMonoPadrao     *TMonofasiaGMonoPadrao `xml:"http://www.portalfiscal.inf.br/nfag gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten      *TMonofasiaGMonoReten  `xml:"http://www.portalfiscal.inf.br/nfag gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet        *TMonofasiaGMonoRet    `xml:"http://www.portalfiscal.inf.br/nfag gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GMonoDif        *TMonofasiaGMonoDif    `xml:"http://www.portalfiscal.inf.br/nfag gMonoDif,omitempty" json:"gMonoDif,omitempty"`
	VTotIBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/nfag vTotIBSMonoItem" json:"vTotIBSMonoItem"`
	VTotCBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/nfag vTotCBSMonoItem" json:"vTotCBSMonoItem"`
}

type TMotNaoLeitura string

type TMotSubAG string

type TMotivo string

type TNF string

type TNFAgInfNFAgIde struct {
	CUF          TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/nfag cUF" json:"cUF"`
	TpAmb        TAmb                `xml:"http://www.portalfiscal.inf.br/nfag tpAmb" json:"tpAmb"`
	Mod          TModNFAg            `xml:"http://www.portalfiscal.inf.br/nfag mod" json:"mod"`
	Serie        TSerie              `xml:"http://www.portalfiscal.inf.br/nfag serie" json:"serie"`
	NNF          TNF                 `xml:"http://www.portalfiscal.inf.br/nfag nNF" json:"nNF"`
	CNF          string              `xml:"http://www.portalfiscal.inf.br/nfag cNF" json:"cNF"`
	CDV          string              `xml:"http://www.portalfiscal.inf.br/nfag cDV" json:"cDV"`
	DhEmi        TDateTimeUTC        `xml:"http://www.portalfiscal.inf.br/nfag dhEmi" json:"dhEmi"`
	TpEmis       TtpEmis             `xml:"http://www.portalfiscal.inf.br/nfag tpEmis" json:"tpEmis"`
	NSiteAutoriz string              `xml:"http://www.portalfiscal.inf.br/nfag nSiteAutoriz" json:"nSiteAutoriz"`
	CMunFG       TCodMunIBGE         `xml:"http://www.portalfiscal.inf.br/nfag cMunFG" json:"cMunFG"`
	FinNFAg      TFinNFAg            `xml:"http://www.portalfiscal.inf.br/nfag finNFAg" json:"finNFAg"`
	TpFat        TFatNFAg            `xml:"http://www.portalfiscal.inf.br/nfag tpFat" json:"tpFat"`
	VerProc      string              `xml:"http://www.portalfiscal.inf.br/nfag verProc" json:"verProc"`
	DhCont       *TDateTimeUTC       `xml:"http://www.portalfiscal.inf.br/nfag dhCont,omitempty" json:"dhCont,omitempty"`
	XJust        *string             `xml:"http://www.portalfiscal.inf.br/nfag xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov   *TCompraGovReduzido `xml:"http://www.portalfiscal.inf.br/nfag gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	TpPagAnt     *TPagAnt            `xml:"http://www.portalfiscal.inf.br/nfag tpPagAnt,omitempty" json:"tpPagAnt,omitempty"`
}

type TNFAgInfNFAgEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/nfag CNPJ" json:"CNPJ"`
	IE        *TIe     `xml:"http://www.portalfiscal.inf.br/nfag IE,omitempty" json:"IE,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/nfag xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/nfag xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/nfag enderEmit" json:"enderEmit"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/nfag ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TNFAgInfNFAgDest struct {
	XNome          string    `xml:"http://www.portalfiscal.inf.br/nfag xNome" json:"xNome"`
	CNPJ           *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/nfag CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF            *TCPF     `xml:"http://www.portalfiscal.inf.br/nfag CPF,omitempty" json:"CPF,omitempty"`
	IDOutros       *string   `xml:"http://www.portalfiscal.inf.br/nfag idOutros,omitempty" json:"idOutros,omitempty"`
	IE             *TIeDest  `xml:"http://www.portalfiscal.inf.br/nfag IE,omitempty" json:"IE,omitempty"`
	IM             *string   `xml:"http://www.portalfiscal.inf.br/nfag IM,omitempty" json:"IM,omitempty"`
	CNIS           *string   `xml:"http://www.portalfiscal.inf.br/nfag cNIS,omitempty" json:"cNIS,omitempty"`
	NB             *string   `xml:"http://www.portalfiscal.inf.br/nfag NB,omitempty" json:"NB,omitempty"`
	XNomeAdicional *string   `xml:"http://www.portalfiscal.inf.br/nfag xNomeAdicional,omitempty" json:"xNomeAdicional,omitempty"`
	EnderDest      TEndeEmi  `xml:"http://www.portalfiscal.inf.br/nfag enderDest" json:"enderDest"`
}

type TNFAgInfNFAgLigacao struct {
	IDLigacao         string     `xml:"http://www.portalfiscal.inf.br/nfag idLigacao" json:"idLigacao"`
	IDCodCliente      *string    `xml:"http://www.portalfiscal.inf.br/nfag idCodCliente,omitempty" json:"idCodCliente,omitempty"`
	TpLigacao         TLigacao   `xml:"http://www.portalfiscal.inf.br/nfag tpLigacao" json:"tpLigacao"`
	LatGPS            TLatitude  `xml:"http://www.portalfiscal.inf.br/nfag latGPS" json:"latGPS"`
	LongGPS           TLongitude `xml:"http://www.portalfiscal.inf.br/nfag longGPS" json:"longGPS"`
	CodRoteiroLeitura *string    `xml:"http://www.portalfiscal.inf.br/nfag codRoteiroLeitura,omitempty" json:"codRoteiroLeitura,omitempty"`
}

type TNFAgInfNFAgGSub struct {
	ChNFAg TChDFe    `xml:"http://www.portalfiscal.inf.br/nfag chNFAg" json:"chNFAg"`
	MotSub TMotSubAG `xml:"http://www.portalfiscal.inf.br/nfag motSub" json:"motSub"`
}

type TNFAgInfNFAgGMed struct {
	NMed      string `xml:"nMed,attr" json:"nMed"`
	IDMedidor string `xml:"http://www.portalfiscal.inf.br/nfag idMedidor" json:"idMedidor"`
	DMedAnt   TData  `xml:"http://www.portalfiscal.inf.br/nfag dMedAnt" json:"dMedAnt"`
	DMedAtu   TData  `xml:"http://www.portalfiscal.inf.br/nfag dMedAtu" json:"dMedAtu"`
}

type TNFAgInfNFAgGFatConjunto struct {
	ChNFAgFat TChDFe `xml:"http://www.portalfiscal.inf.br/nfag chNFAgFat" json:"chNFAgFat"`
}

type TNFAgInfNFAgDetGTarif struct {
	DIniTarif   TData   `xml:"http://www.portalfiscal.inf.br/nfag dIniTarif" json:"dIniTarif"`
	DFimTarif   *TData  `xml:"http://www.portalfiscal.inf.br/nfag dFimTarif,omitempty" json:"dFimTarif,omitempty"`
	NAto        string  `xml:"http://www.portalfiscal.inf.br/nfag nAto" json:"nAto"`
	AnoAto      string  `xml:"http://www.portalfiscal.inf.br/nfag anoAto" json:"anoAto"`
	TpFaixaCons *TFaixa `xml:"http://www.portalfiscal.inf.br/nfag tpFaixaCons,omitempty" json:"tpFaixaCons,omitempty"`
}

type TNFAgInfNFAgDetProdGMedicao struct {
	NMed            string          `xml:"http://www.portalfiscal.inf.br/nfag nMed" json:"nMed"`
	GMedida         *TMedidaAG      `xml:"http://www.portalfiscal.inf.br/nfag gMedida,omitempty" json:"gMedida,omitempty"`
	TpMotNaoLeitura *TMotNaoLeitura `xml:"http://www.portalfiscal.inf.br/nfag tpMotNaoLeitura,omitempty" json:"tpMotNaoLeitura,omitempty"`
}

type TNFAgInfNFAgDetProdGPagAntecipado struct {
	ChDFePagAnt TChDFe  `xml:"http://www.portalfiscal.inf.br/nfag chDFePagAnt" json:"chDFePagAnt"`
	NItemPagAnt *string `xml:"http://www.portalfiscal.inf.br/nfag nItemPagAnt,omitempty" json:"nItemPagAnt,omitempty"`
}

type TNFAgInfNFAgDetProd struct {
	IndOrigemQtd   TOrigemQtdAG                       `xml:"http://www.portalfiscal.inf.br/nfag indOrigemQtd" json:"indOrigemQtd"`
	GMedicao       *TNFAgInfNFAgDetProdGMedicao       `xml:"http://www.portalfiscal.inf.br/nfag gMedicao,omitempty" json:"gMedicao,omitempty"`
	CProd          string                             `xml:"http://www.portalfiscal.inf.br/nfag cProd" json:"cProd"`
	XProd          string                             `xml:"http://www.portalfiscal.inf.br/nfag xProd" json:"xProd"`
	CClass         string                             `xml:"http://www.portalfiscal.inf.br/nfag cClass" json:"cClass"`
	TpCategoria    *TCategAG                          `xml:"http://www.portalfiscal.inf.br/nfag tpCategoria,omitempty" json:"tpCategoria,omitempty"`
	XCategoria     *string                            `xml:"http://www.portalfiscal.inf.br/nfag xCategoria,omitempty" json:"xCategoria,omitempty"`
	QEconomias     *string                            `xml:"http://www.portalfiscal.inf.br/nfag qEconomias,omitempty" json:"qEconomias,omitempty"`
	UMed           string                             `xml:"http://www.portalfiscal.inf.br/nfag uMed" json:"uMed"`
	QFaturada      TDec11001104                       `xml:"http://www.portalfiscal.inf.br/nfag qFaturada" json:"qFaturada"`
	VItem          TDec13021310                       `xml:"http://www.portalfiscal.inf.br/nfag vItem" json:"vItem"`
	FatorPoluicao  *TDec11001104                      `xml:"http://www.portalfiscal.inf.br/nfag fatorPoluicao,omitempty" json:"fatorPoluicao,omitempty"`
	VProd          TDec13021310                       `xml:"http://www.portalfiscal.inf.br/nfag vProd" json:"vProd"`
	VItemLiq       *TDec13021310                      `xml:"http://www.portalfiscal.inf.br/nfag vItemLiq,omitempty" json:"vItemLiq,omitempty"`
	VProdLiq       *TDec13021310                      `xml:"http://www.portalfiscal.inf.br/nfag vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	IndDevolucao   *string                            `xml:"http://www.portalfiscal.inf.br/nfag indDevolucao,omitempty" json:"indDevolucao,omitempty"`
	GPagAntecipado *TNFAgInfNFAgDetProdGPagAntecipado `xml:"http://www.portalfiscal.inf.br/nfag gPagAntecipado,omitempty" json:"gPagAntecipado,omitempty"`
}

type TNFAgInfNFAgDetImpostoPIS struct {
	CST  string      `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	VBC  TDec1302    `xml:"http://www.portalfiscal.inf.br/nfag vBC" json:"vBC"`
	PPIS TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfag pPIS" json:"pPIS"`
	VPIS TDec1302    `xml:"http://www.portalfiscal.inf.br/nfag vPIS" json:"vPIS"`
}

type TNFAgInfNFAgDetImpostoCOFINS struct {
	CST     string      `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	VBC     TDec1302    `xml:"http://www.portalfiscal.inf.br/nfag vBC" json:"vBC"`
	PCOFINS TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfag pCOFINS" json:"pCOFINS"`
	VCOFINS TDec1302    `xml:"http://www.portalfiscal.inf.br/nfag vCOFINS" json:"vCOFINS"`
}

type TNFAgInfNFAgDetImpostoRetTrib struct {
	VRetPIS    TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfag vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfag vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfag vRetCSLL" json:"vRetCSLL"`
	VIRRF      TDec13021308 `xml:"http://www.portalfiscal.inf.br/nfag vIRRF" json:"vIRRF"`
}

type TNFAgInfNFAgDetImpostoTFS struct {
	VBCTFS TDec1302    `xml:"http://www.portalfiscal.inf.br/nfag vBCTFS" json:"vBCTFS"`
	PTFS   TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfag pTFS" json:"pTFS"`
	VTFS   TDec1302    `xml:"http://www.portalfiscal.inf.br/nfag vTFS" json:"vTFS"`
}

type TNFAgInfNFAgDetImpostoTFU struct {
	VBCTFU TDec1302    `xml:"http://www.portalfiscal.inf.br/nfag vBCTFU" json:"vBCTFU"`
	PTFU   TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nfag pTFU" json:"pTFU"`
	VTFU   TDec1302    `xml:"http://www.portalfiscal.inf.br/nfag vTFU" json:"vTFU"`
}

type TNFAgInfNFAgDetImposto struct {
	IBSCBS  TTribNFAg                      `xml:"http://www.portalfiscal.inf.br/nfag IBSCBS" json:"IBSCBS"`
	PIS     *TNFAgInfNFAgDetImpostoPIS     `xml:"http://www.portalfiscal.inf.br/nfag PIS,omitempty" json:"PIS,omitempty"`
	COFINS  *TNFAgInfNFAgDetImpostoCOFINS  `xml:"http://www.portalfiscal.inf.br/nfag COFINS,omitempty" json:"COFINS,omitempty"`
	RetTrib *TNFAgInfNFAgDetImpostoRetTrib `xml:"http://www.portalfiscal.inf.br/nfag retTrib,omitempty" json:"retTrib,omitempty"`
	TFS     *TNFAgInfNFAgDetImpostoTFS     `xml:"http://www.portalfiscal.inf.br/nfag TFS,omitempty" json:"TFS,omitempty"`
	TFU     *TNFAgInfNFAgDetImpostoTFU     `xml:"http://www.portalfiscal.inf.br/nfag TFU,omitempty" json:"TFU,omitempty"`
}

type TNFAgInfNFAgDetGProcRefGIBSCBS struct {
	VBC     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vBC" json:"vBC"`
	VIBSUF  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSUF" json:"vIBSUF"`
	VIBSMun TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSMun" json:"vIBSMun"`
	VCBS    TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBS" json:"vCBS"`
}

type TNFAgInfNFAgDetGProcRefGProc struct {
	TpProc    TProcessoAG `xml:"http://www.portalfiscal.inf.br/nfag tpProc" json:"tpProc"`
	NProcesso string      `xml:"http://www.portalfiscal.inf.br/nfag nProcesso" json:"nProcesso"`
}

type TNFAgInfNFAgDetGProcRef struct {
	VItem        TDec13021310                    `xml:"http://www.portalfiscal.inf.br/nfag vItem" json:"vItem"`
	QFaturada    TDec11001104                    `xml:"http://www.portalfiscal.inf.br/nfag qFaturada" json:"qFaturada"`
	VProd        TDec13021310                    `xml:"http://www.portalfiscal.inf.br/nfag vProd" json:"vProd"`
	IndDevolucao *string                         `xml:"http://www.portalfiscal.inf.br/nfag indDevolucao,omitempty" json:"indDevolucao,omitempty"`
	GIBSCBS      *TNFAgInfNFAgDetGProcRefGIBSCBS `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GProc        []TNFAgInfNFAgDetGProcRefGProc  `xml:"http://www.portalfiscal.inf.br/nfag gProc,omitempty" json:"gProc,omitempty"`
}

type TNFAgInfNFAgDetGAgregadora struct {
	CClass  string   `xml:"http://www.portalfiscal.inf.br/nfag cClass" json:"cClass"`
	VTotDFe TDec1302 `xml:"http://www.portalfiscal.inf.br/nfag vTotDFe" json:"vTotDFe"`
}

type TNFAgInfNFAgDet struct {
	NItem       string                      `xml:"nItem,attr" json:"nItem"`
	ChNFAgAnt   *TChDFe                     `xml:"chNFAgAnt,attr,omitempty" json:"chNFAgAnt,omitempty"`
	NItemAnt    *string                     `xml:"nItemAnt,attr,omitempty" json:"nItemAnt,omitempty"`
	GTarif      []TNFAgInfNFAgDetGTarif     `xml:"http://www.portalfiscal.inf.br/nfag gTarif,omitempty" json:"gTarif,omitempty"`
	Prod        *TNFAgInfNFAgDetProd        `xml:"http://www.portalfiscal.inf.br/nfag prod,omitempty" json:"prod,omitempty"`
	Imposto     *TNFAgInfNFAgDetImposto     `xml:"http://www.portalfiscal.inf.br/nfag imposto,omitempty" json:"imposto,omitempty"`
	GProcRef    *TNFAgInfNFAgDetGProcRef    `xml:"http://www.portalfiscal.inf.br/nfag gProcRef,omitempty" json:"gProcRef,omitempty"`
	InfAdProd   *string                     `xml:"http://www.portalfiscal.inf.br/nfag infAdProd,omitempty" json:"infAdProd,omitempty"`
	GAgregadora *TNFAgInfNFAgDetGAgregadora `xml:"http://www.portalfiscal.inf.br/nfag gAgregadora,omitempty" json:"gAgregadora,omitempty"`
}

type TNFAgInfNFAgTotalVRetTribTot struct {
	VRetPIS    TDec1302 `xml:"http://www.portalfiscal.inf.br/nfag vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec1302 `xml:"http://www.portalfiscal.inf.br/nfag vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec1302 `xml:"http://www.portalfiscal.inf.br/nfag vRetCSLL" json:"vRetCSLL"`
	VIRRF      TDec1302 `xml:"http://www.portalfiscal.inf.br/nfag vIRRF" json:"vIRRF"`
}

type TNFAgInfNFAgTotal struct {
	VProd       TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfag vProd" json:"vProd"`
	VProdLiq    *TDec1302                    `xml:"http://www.portalfiscal.inf.br/nfag vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	VRetTribTot TNFAgInfNFAgTotalVRetTribTot `xml:"http://www.portalfiscal.inf.br/nfag vRetTribTot" json:"vRetTribTot"`
	VCOFINS     TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfag vCOFINS" json:"vCOFINS"`
	VPIS        TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfag vPIS" json:"vPIS"`
	VTFS        TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfag vTFS" json:"vTFS"`
	VTFU        TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfag vTFU" json:"vTFU"`
	VNF         TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfag vNF" json:"vNF"`
	IBSCBSTot   TIBSCBSTot                   `xml:"http://www.portalfiscal.inf.br/nfag IBSCBSTot" json:"IBSCBSTot"`
	VTotDFe     TDec1302                     `xml:"http://www.portalfiscal.inf.br/nfag vTotDFe" json:"vTotDFe"`
}

type TNFAgInfNFAgPgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/nfag pgto,omitempty" json:"pgto,omitempty"`
}

type TNFAgInfNFAgGFatGPIX struct {
	URLQRCodePIX string `xml:"http://www.portalfiscal.inf.br/nfag urlQRCodePIX" json:"urlQRCodePIX"`
}

type TNFAgInfNFAgGFat struct {
	CompetFat    TCompet               `xml:"http://www.portalfiscal.inf.br/nfag CompetFat" json:"CompetFat"`
	DVencFat     TData                 `xml:"http://www.portalfiscal.inf.br/nfag dVencFat" json:"dVencFat"`
	DApresFat    *TData                `xml:"http://www.portalfiscal.inf.br/nfag dApresFat,omitempty" json:"dApresFat,omitempty"`
	DProxLeitura TData                 `xml:"http://www.portalfiscal.inf.br/nfag dProxLeitura" json:"dProxLeitura"`
	NFat         *string               `xml:"http://www.portalfiscal.inf.br/nfag nFat,omitempty" json:"nFat,omitempty"`
	CodBarras    string                `xml:"http://www.portalfiscal.inf.br/nfag codBarras" json:"codBarras"`
	CodDebAuto   *string               `xml:"http://www.portalfiscal.inf.br/nfag codDebAuto,omitempty" json:"codDebAuto,omitempty"`
	CodBanco     *string               `xml:"http://www.portalfiscal.inf.br/nfag codBanco,omitempty" json:"codBanco,omitempty"`
	CodAgencia   *string               `xml:"http://www.portalfiscal.inf.br/nfag codAgencia,omitempty" json:"codAgencia,omitempty"`
	EnderCorresp *TEndeEmi             `xml:"http://www.portalfiscal.inf.br/nfag enderCorresp,omitempty" json:"enderCorresp,omitempty"`
	GPIX         *TNFAgInfNFAgGFatGPIX `xml:"http://www.portalfiscal.inf.br/nfag gPIX,omitempty" json:"gPIX,omitempty"`
}

type TNFAgInfNFAgGAgenciaGHistConsGCons struct {
	CompetFat TCompet       `xml:"http://www.portalfiscal.inf.br/nfag CompetFat" json:"CompetFat"`
	UMed      TUMedAG       `xml:"http://www.portalfiscal.inf.br/nfag uMed" json:"uMed"`
	QtdDias   string        `xml:"http://www.portalfiscal.inf.br/nfag qtdDias" json:"qtdDias"`
	MedDiaria *TDec13021304 `xml:"http://www.portalfiscal.inf.br/nfag medDiaria,omitempty" json:"medDiaria,omitempty"`
	Consumo   *TDec1302     `xml:"http://www.portalfiscal.inf.br/nfag consumo,omitempty" json:"consumo,omitempty"`
	VolFat    TDec1302      `xml:"http://www.portalfiscal.inf.br/nfag volFat" json:"volFat"`
}

type TNFAgInfNFAgGAgenciaGHistCons struct {
	XHistorico string                               `xml:"http://www.portalfiscal.inf.br/nfag xHistorico" json:"xHistorico"`
	GCons      []TNFAgInfNFAgGAgenciaGHistConsGCons `xml:"http://www.portalfiscal.inf.br/nfag gCons,omitempty" json:"gCons,omitempty"`
	MedMensal  TDec13021304                         `xml:"http://www.portalfiscal.inf.br/nfag medMensal" json:"medMensal"`
}

type TNFAgInfNFAgGAgencia struct {
	Econ              *string                         `xml:"http://www.portalfiscal.inf.br/nfag econ,omitempty" json:"econ,omitempty"`
	EconAcumulada     *string                         `xml:"http://www.portalfiscal.inf.br/nfag econAcumulada,omitempty" json:"econAcumulada,omitempty"`
	SPrestador        *string                         `xml:"http://www.portalfiscal.inf.br/nfag sPrestador,omitempty" json:"sPrestador,omitempty"`
	DEmissSelo        *TData                          `xml:"http://www.portalfiscal.inf.br/nfag dEmissSelo,omitempty" json:"dEmissSelo,omitempty"`
	SRegulador        *string                         `xml:"http://www.portalfiscal.inf.br/nfag sRegulador,omitempty" json:"sRegulador,omitempty"`
	NAgenciaAtend     string                          `xml:"http://www.portalfiscal.inf.br/nfag nAgenciaAtend" json:"nAgenciaAtend"`
	EnderAgenciaAtend string                          `xml:"http://www.portalfiscal.inf.br/nfag enderAgenciaAtend" json:"enderAgenciaAtend"`
	GHistCons         []TNFAgInfNFAgGAgenciaGHistCons `xml:"http://www.portalfiscal.inf.br/nfag gHistCons,omitempty" json:"gHistCons,omitempty"`
}

type TNFAgInfNFAgGQualiAguaGAnalise struct {
	XItemAnalisado    string  `xml:"http://www.portalfiscal.inf.br/nfag xItemAnalisado" json:"xItemAnalisado"`
	NAmostraMinima    *string `xml:"http://www.portalfiscal.inf.br/nfag nAmostraMinima,omitempty" json:"nAmostraMinima,omitempty"`
	NAmostraAnalisada *string `xml:"http://www.portalfiscal.inf.br/nfag nAmostraAnalisada,omitempty" json:"nAmostraAnalisada,omitempty"`
	NAmostraFPadrao   *string `xml:"http://www.portalfiscal.inf.br/nfag nAmostraFPadrao,omitempty" json:"nAmostraFPadrao,omitempty"`
	NAmostraDPadrao   *string `xml:"http://www.portalfiscal.inf.br/nfag nAmostraDPadrao,omitempty" json:"nAmostraDPadrao,omitempty"`
	NMediaMensal      *string `xml:"http://www.portalfiscal.inf.br/nfag nMediaMensal,omitempty" json:"nMediaMensal,omitempty"`
	XValorReferencia  *string `xml:"http://www.portalfiscal.inf.br/nfag xValorReferencia,omitempty" json:"xValorReferencia,omitempty"`
}

type TNFAgInfNFAgGQualiAgua struct {
	CompetAnalise TCompet                          `xml:"http://www.portalfiscal.inf.br/nfag CompetAnalise" json:"CompetAnalise"`
	GAnalise      []TNFAgInfNFAgGQualiAguaGAnalise `xml:"http://www.portalfiscal.inf.br/nfag gAnalise,omitempty" json:"gAnalise,omitempty"`
	Conclusao     *string                          `xml:"http://www.portalfiscal.inf.br/nfag Conclusao,omitempty" json:"Conclusao,omitempty"`
	CProcesso     *string                          `xml:"http://www.portalfiscal.inf.br/nfag cProcesso,omitempty" json:"cProcesso,omitempty"`
	SistemaAbast  *string                          `xml:"http://www.portalfiscal.inf.br/nfag SistemaAbast,omitempty" json:"SistemaAbast,omitempty"`
}

type TNFAgInfNFAgAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/nfag CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/nfag CPF,omitempty" json:"CPF,omitempty"`
}

type TNFAgInfNFAgInfAdic struct {
	InfAdFisco *string  `xml:"http://www.portalfiscal.inf.br/nfag infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     []string `xml:"http://www.portalfiscal.inf.br/nfag infCpl,omitempty" json:"infCpl,omitempty"`
}

type TNFAgInfNFAgInfPAA struct {
	CNPJPAA TCNPJ `xml:"http://www.portalfiscal.inf.br/nfag CNPJPAA" json:"CNPJPAA"`
}

type TNFAgInfNFAg struct {
	Versao       TVerNFAg                  `xml:"versao,attr" json:"versao"`
	ID           string                    `xml:"Id,attr" json:"Id"`
	Ide          TNFAgInfNFAgIde           `xml:"http://www.portalfiscal.inf.br/nfag ide" json:"ide"`
	Emit         TNFAgInfNFAgEmit          `xml:"http://www.portalfiscal.inf.br/nfag emit" json:"emit"`
	Dest         TNFAgInfNFAgDest          `xml:"http://www.portalfiscal.inf.br/nfag dest" json:"dest"`
	Ligacao      *TNFAgInfNFAgLigacao      `xml:"http://www.portalfiscal.inf.br/nfag ligacao,omitempty" json:"ligacao,omitempty"`
	GSub         *TNFAgInfNFAgGSub         `xml:"http://www.portalfiscal.inf.br/nfag gSub,omitempty" json:"gSub,omitempty"`
	GMed         []TNFAgInfNFAgGMed        `xml:"http://www.portalfiscal.inf.br/nfag gMed,omitempty" json:"gMed,omitempty"`
	GFatConjunto *TNFAgInfNFAgGFatConjunto `xml:"http://www.portalfiscal.inf.br/nfag gFatConjunto,omitempty" json:"gFatConjunto,omitempty"`
	Det          []TNFAgInfNFAgDet         `xml:"http://www.portalfiscal.inf.br/nfag det,omitempty" json:"det,omitempty"`
	Total        TNFAgInfNFAgTotal         `xml:"http://www.portalfiscal.inf.br/nfag total" json:"total"`
	PgtoVinc     *TNFAgInfNFAgPgtoVinc     `xml:"http://www.portalfiscal.inf.br/nfag pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	GFat         *TNFAgInfNFAgGFat         `xml:"http://www.portalfiscal.inf.br/nfag gFat,omitempty" json:"gFat,omitempty"`
	GAgencia     *TNFAgInfNFAgGAgencia     `xml:"http://www.portalfiscal.inf.br/nfag gAgencia,omitempty" json:"gAgencia,omitempty"`
	GQualiAgua   *TNFAgInfNFAgGQualiAgua   `xml:"http://www.portalfiscal.inf.br/nfag gQualiAgua,omitempty" json:"gQualiAgua,omitempty"`
	AutXML       []TNFAgInfNFAgAutXML      `xml:"http://www.portalfiscal.inf.br/nfag autXML,omitempty" json:"autXML,omitempty"`
	InfAdic      *TNFAgInfNFAgInfAdic      `xml:"http://www.portalfiscal.inf.br/nfag infAdic,omitempty" json:"infAdic,omitempty"`
	InfPAA       *TNFAgInfNFAgInfPAA       `xml:"http://www.portalfiscal.inf.br/nfag infPAA,omitempty" json:"infPAA,omitempty"`
	GRespTec     *TRespTec                 `xml:"http://www.portalfiscal.inf.br/nfag gRespTec,omitempty" json:"gRespTec,omitempty"`
}

type TNFAgInfNFAgSupl struct {
	QRCodNFAg string `xml:"http://www.portalfiscal.inf.br/nfag qrCodNFAg" json:"qrCodNFAg"`
}

type TNFAg struct {
	InfNFAg     TNFAgInfNFAg     `xml:"http://www.portalfiscal.inf.br/nfag infNFAg" json:"infNFAg"`
	InfNFAgSupl TNFAgInfNFAgSupl `xml:"http://www.portalfiscal.inf.br/nfag infNFAgSupl" json:"infNFAgSupl"`
	Signature   SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TNSU string

type TOperCompraGov string

type TOrigemQtdAG string

type TPagAnt string

type TPagRef struct {
	RefDFe []TChDFeRTC `xml:"http://www.portalfiscal.inf.br/nfag refDFe,omitempty" json:"refDFe,omitempty"`
}

type TPagamentoRTC struct {
	NPag        string       `xml:"nPag,attr" json:"nPag"`
	IDTransacao string       `xml:"idTransacao,attr" json:"idTransacao"`
	TpMeioPgto  string       `xml:"http://www.portalfiscal.inf.br/nfag tpMeioPgto" json:"tpMeioPgto"`
	CNPJReceb   TCNPJRTC     `xml:"http://www.portalfiscal.inf.br/nfag CNPJReceb" json:"CNPJReceb"`
	CNPJBasePSP TCNPJBaseRTC `xml:"http://www.portalfiscal.inf.br/nfag CNPJBasePSP" json:"CNPJBasePSP"`
}

type TPosTarSimples string

type TPosTarifCompleto string

type TProcEmi string

type TProcessoAG string

type TProt string

type TProtNFAgInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/nfag tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/nfag verAplic" json:"verAplic"`
	ChNFAg   TChDFe           `xml:"http://www.portalfiscal.inf.br/nfag chNFAg" json:"chNFAg"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/nfag dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/nfag nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/nfag digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/nfag cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/nfag xMotivo" json:"xMotivo"`
}

type TProtNFAgInfFisco struct {
	CMsg string  `xml:"http://www.portalfiscal.inf.br/nfag cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/nfag xMsg" json:"xMsg"`
}

type TProtNFAg struct {
	Versao    TVerNFAg           `xml:"versao,attr" json:"versao"`
	InfProt   TProtNFAgInfProt   `xml:"http://www.portalfiscal.inf.br/nfag infProt" json:"infProt"`
	InfFisco  *TProtNFAgInfFisco `xml:"http://www.portalfiscal.inf.br/nfag infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType     `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRBSN string

type TRec string

type TRed struct {
	PRedAliq  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqEfet" json:"pAliqEfet"`
}

type TRespTec struct {
	CNPJ     TCNPJ   `xml:"http://www.portalfiscal.inf.br/nfag CNPJ" json:"CNPJ"`
	XContato string  `xml:"http://www.portalfiscal.inf.br/nfag xContato" json:"xContato"`
	Email    TEmail  `xml:"http://www.portalfiscal.inf.br/nfag email" json:"email"`
	Fone     string  `xml:"http://www.portalfiscal.inf.br/nfag fone" json:"fone"`
	IDCSRT   *string `xml:"http://www.portalfiscal.inf.br/nfag idCSRT,omitempty" json:"idCSRT,omitempty"`
	HashCSRT *string `xml:"http://www.portalfiscal.inf.br/nfag hashCSRT,omitempty" json:"hashCSRT,omitempty"`
}

type TRetNFAg struct {
	Versao   TVerNFAg   `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb       `xml:"http://www.portalfiscal.inf.br/nfag tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/nfag cUF" json:"cUF"`
	VerAplic TVerAplic  `xml:"http://www.portalfiscal.inf.br/nfag verAplic" json:"verAplic"`
	CStat    TStat      `xml:"http://www.portalfiscal.inf.br/nfag cStat" json:"cStat"`
	XMotivo  TMotivo    `xml:"http://www.portalfiscal.inf.br/nfag xMotivo" json:"xMotivo"`
	ProtNFAg *TProtNFAg `xml:"http://www.portalfiscal.inf.br/nfag protNFAg,omitempty" json:"protNFAg,omitempty"`
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
	VRBSNTot          TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nfag vRBSNTot" json:"vRBSNTot"`
	VIBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSSNTot,omitempty" json:"vIBSSNTot,omitempty"`
	VIBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBSSNtotPendSusp,omitempty" json:"vIBSSNtotPendSusp,omitempty"`
	VCBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBSSNTot,omitempty" json:"vCBSSNTot,omitempty"`
	VCBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBSSNtotPendSusp,omitempty" json:"vCBSSNtotPendSusp,omitempty"`
}

type TTpCredPresIBSZFM string

type TTransfCred struct {
	VIBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vIBS" json:"vIBS"`
	VCBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nfag vCBS" json:"vCBS"`
}

type TTribBPe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfag cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfag indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCTe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfag cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfag indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCompraGov struct {
	PAliqIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vTribCBS" json:"vTribCBS"`
}

type TTribItemSN struct {
	NItem        string         `xml:"nItem,attr" json:"nItem"`
	VRBSNItem    TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nfag vRBSNItem" json:"vRBSNItem"`
	TpRBSN       TRBSN          `xml:"http://www.portalfiscal.inf.br/nfag tpRBSN" json:"tpRBSN"`
	PIBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pIBSSN,omitempty" json:"pIBSSN,omitempty"`
	VIBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSSN,omitempty" json:"vIBSSN,omitempty"`
	PCBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pCBSSN,omitempty" json:"pCBSSN,omitempty"`
	VCBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vCBSSN,omitempty" json:"vCBSSN,omitempty"`
	VIBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vIBSPendSusp,omitempty" json:"vIBSPendSusp,omitempty"`
	VCBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vCBSPendSusp,omitempty" json:"vCBSPendSusp,omitempty"`
}

type TTribNF3e struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfag cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfag indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFAg struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfag cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfag indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFCe struct {
	CST         TCST        `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	CClassTrib  TcClassTrib `xml:"http://www.portalfiscal.inf.br/nfag cClassTrib" json:"cClassTrib"`
	IndDoacao   *TIndDoacao `xml:"http://www.portalfiscal.inf.br/nfag indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS     *TCIBSNFe   `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono *TMonofasia `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
}

type TTribNFCom struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfag cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfag indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFGas struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfag cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nfag indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFe struct {
	CST             TCST             `xml:"http://www.portalfiscal.inf.br/nfag CST" json:"CST"`
	CClassTrib      TcClassTrib      `xml:"http://www.portalfiscal.inf.br/nfag cClassTrib" json:"cClassTrib"`
	IndDoacao       *TIndDoacao      `xml:"http://www.portalfiscal.inf.br/nfag indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS         *TCIBSNFe        `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono     *TMonofasia      `xml:"http://www.portalfiscal.inf.br/nfag gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
	GTransfCred     *TTransfCred     `xml:"http://www.portalfiscal.inf.br/nfag gTransfCred,omitempty" json:"gTransfCred,omitempty"`
	GAjusteCompet   *TAjusteCompet   `xml:"http://www.portalfiscal.inf.br/nfag gAjusteCompet,omitempty" json:"gAjusteCompet,omitempty"`
	GEstornoCred    *TEstornoCred    `xml:"http://www.portalfiscal.inf.br/nfag gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
	GCredPresOper   *TCredPresOper   `xml:"http://www.portalfiscal.inf.br/nfag gCredPresOper,omitempty" json:"gCredPresOper,omitempty"`
	GCredPresIBSZFM *TCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nfag gCredPresIBSZFM,omitempty" json:"gCredPresIBSZFM,omitempty"`
}

type TTribRegular struct {
	CSTReg             TCST          `xml:"http://www.portalfiscal.inf.br/nfag CSTReg" json:"CSTReg"`
	CClassTribReg      TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nfag cClassTribReg" json:"cClassTribReg"`
	PAliqEfetRegIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqEfetRegIBSUF" json:"pAliqEfetRegIBSUF"`
	VTribRegIBSUF      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vTribRegIBSUF" json:"vTribRegIBSUF"`
	PAliqEfetRegIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqEfetRegIBSMun" json:"pAliqEfetRegIBSMun"`
	VTribRegIBSMun     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vTribRegIBSMun" json:"vTribRegIBSMun"`
	PAliqEfetRegCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nfag pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS        TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nfag vTribRegCBS" json:"vTribRegCBS"`
}

type TUMedAG string

type TUMedItemAG string

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerNFAg string

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

type NFAgProcContent struct {
	Versao        TVerNFAg      `xml:"versao,attr" json:"versao"`
	IpTransmissor *TIPv4        `xml:"ipTransmissor,attr,omitempty" json:"ipTransmissor,omitempty"`
	NPortaCon     *string       `xml:"nPortaCon,attr,omitempty" json:"nPortaCon,omitempty"`
	DhConexao     *TDateTimeUTC `xml:"dhConexao,attr,omitempty" json:"dhConexao,omitempty"`
	NFAg          TNFAg         `xml:"http://www.portalfiscal.inf.br/nfag NFAg" json:"NFAg"`
	ProtNFAg      TProtNFAg     `xml:"http://www.portalfiscal.inf.br/nfag protNFAg" json:"protNFAg"`
}

// NFAgProc represents the global element {http://www.portalfiscal.inf.br/nfag}nfagProc.
type NFAgProc struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfag nfagProc" json:"-"`
	NFAgProcContent
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
