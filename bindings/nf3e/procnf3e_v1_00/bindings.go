// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package procnf3e_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TALCZFMCBS struct {
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vTribRegCBS" json:"vTribRegCBS"`
}

type TALCZFMCBSNFe struct {
	TpALCZFMCBS     string        `xml:"http://www.portalfiscal.inf.br/nf3e tpALCZFMCBS" json:"tpALCZFMCBS"`
	NProcSuframa    *string       `xml:"http://www.portalfiscal.inf.br/nf3e nProcSuframa,omitempty" json:"nProcSuframa,omitempty"`
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vTribRegCBS" json:"vTribRegCBS"`
}

type TANEELGHistFatGGrandFat struct {
	CompetFat TCompet      `xml:"http://www.portalfiscal.inf.br/nf3e CompetFat" json:"CompetFat"`
	VFat      TDec13021304 `xml:"http://www.portalfiscal.inf.br/nf3e vFat" json:"vFat"`
	UMed      TUMed        `xml:"http://www.portalfiscal.inf.br/nf3e uMed" json:"uMed"`
	QtdDias   string       `xml:"http://www.portalfiscal.inf.br/nf3e qtdDias" json:"qtdDias"`
}

type TANEELGHistFat struct {
	XGrandFat string                    `xml:"http://www.portalfiscal.inf.br/nf3e xGrandFat" json:"xGrandFat"`
	GGrandFat []TANEELGHistFatGGrandFat `xml:"http://www.portalfiscal.inf.br/nf3e gGrandFat,omitempty" json:"gGrandFat,omitempty"`
}

type TANEEL struct {
	GHistFat []TANEELGHistFat `xml:"http://www.portalfiscal.inf.br/nf3e gHistFat,omitempty" json:"gHistFat,omitempty"`
}

type TAcessante string

type TAjusteCompet struct {
	CompetApur TCompetApur `xml:"http://www.portalfiscal.inf.br/nf3e competApur" json:"competApur"`
	VIBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBS" json:"vIBS"`
	VCBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBS" json:"vCBS"`
}

type TAjusteNF3eAnt string

type TAjusteNF3eAntV2 string

type TAmb string

type TAtoANEEL string

type TBand string

type TCIBSGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nf3e gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nf3e gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nf3e gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSUF" json:"vIBSUF"`
}

type TCIBSGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nf3e gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nf3e gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nf3e gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMun" json:"vIBSMun"`
}

type TCIBSGCBS struct {
	PCBS       TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pCBS" json:"pCBS"`
	GDif       *TDif         `xml:"http://www.portalfiscal.inf.br/nf3e gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nf3e gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed         `xml:"http://www.portalfiscal.inf.br/nf3e gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBS   `xml:"http://www.portalfiscal.inf.br/nf3e gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vCBS" json:"vCBS"`
}

type TCIBS struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nf3e vBC" json:"vBC"`
	GIBSUF         TCIBSGIBSUF     `xml:"http://www.portalfiscal.inf.br/nf3e gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSGIBSMun    `xml:"http://www.portalfiscal.inf.br/nf3e gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nf3e vIBS" json:"vIBS"`
	GCBS           TCIBSGCBS       `xml:"http://www.portalfiscal.inf.br/nf3e gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nf3e gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nf3e gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCIBSNFeGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nf3e gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nf3e gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nf3e gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSUF" json:"vIBSUF"`
}

type TCIBSNFeGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/nf3e gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/nf3e gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/nf3e gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMun" json:"vIBSMun"`
}

type TCIBSNFeGCBS struct {
	PCBS       TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nf3e pCBS" json:"pCBS"`
	GDif       *TDif          `xml:"http://www.portalfiscal.inf.br/nf3e gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib      `xml:"http://www.portalfiscal.inf.br/nf3e gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed          `xml:"http://www.portalfiscal.inf.br/nf3e gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBSNFe `xml:"http://www.portalfiscal.inf.br/nf3e gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nf3e vCBS" json:"vCBS"`
}

type TCIBSNFe struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nf3e vBC" json:"vBC"`
	GIBSUF         TCIBSNFeGIBSUF  `xml:"http://www.portalfiscal.inf.br/nf3e gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSNFeGIBSMun `xml:"http://www.portalfiscal.inf.br/nf3e gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/nf3e vIBS" json:"vIBS"`
	GCBS           TCIBSNFeGCBS    `xml:"http://www.portalfiscal.inf.br/nf3e gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/nf3e gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/nf3e gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCOrgaoIBGE string

type TCST string

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
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nf3e tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nf3e pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nf3e tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nf3e refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCompraGovReduzido struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/nf3e tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/nf3e pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/nf3e tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/nf3e refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCPF string

type TCPFVar string

type TCredPres struct {
	PCredPres        TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pCredPres" json:"pCredPres"`
	VCredPres        *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nf3e vCredPres,omitempty" json:"vCredPres,omitempty"`
	VCredPresCondSus *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nf3e vCredPresCondSus,omitempty" json:"vCredPresCondSus,omitempty"`
}

type TCredPresIBSZFM struct {
	CompetApur       TCompetApur       `xml:"http://www.portalfiscal.inf.br/nf3e competApur" json:"competApur"`
	TpCredPresIBSZFM TTpCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nf3e tpCredPresIBSZFM" json:"tpCredPresIBSZFM"`
	VCredPresIBSZFM  TDec1302RTC       `xml:"http://www.portalfiscal.inf.br/nf3e vCredPresIBSZFM" json:"vCredPresIBSZFM"`
}

type TCredPresOper struct {
	VBCCredPres  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vBCCredPres" json:"vBCCredPres"`
	CCredPres    TcCredPres  `xml:"http://www.portalfiscal.inf.br/nf3e cCredPres" json:"cCredPres"`
	GIBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCredPres,omitempty" json:"gIBSCredPres,omitempty"`
	GCBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/nf3e gCBSCredPres,omitempty" json:"gCBSCredPres,omitempty"`
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
	PDevTrib *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pDevTrib,omitempty" json:"pDevTrib,omitempty"`
	VDevTrib TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nf3e vDevTrib" json:"vDevTrib"`
}

type TDif struct {
	PDif TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pDif" json:"pDif"`
	VDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vDif" json:"vDif"`
}

type TDoc string

type TEmail string

type TEmpresaSoft struct {
	CNPJ     *TCNPJ `xml:"http://www.portalfiscal.inf.br/nf3e CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF      *TCPF  `xml:"http://www.portalfiscal.inf.br/nf3e CPF,omitempty" json:"CPF,omitempty"`
	XContato string `xml:"http://www.portalfiscal.inf.br/nf3e xContato" json:"xContato"`
	Email    TEmail `xml:"http://www.portalfiscal.inf.br/nf3e email" json:"email"`
	Fone     string `xml:"http://www.portalfiscal.inf.br/nf3e fone" json:"fone"`
}

type TEndeEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nf3e xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nf3e nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nf3e xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nf3e xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nf3e cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nf3e xMun" json:"xMun"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/nf3e CEP" json:"CEP"`
	UF      TUFSemEX    `xml:"http://www.portalfiscal.inf.br/nf3e UF" json:"UF"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nf3e fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nf3e email,omitempty" json:"email,omitempty"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/nf3e xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/nf3e nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/nf3e xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/nf3e xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/nf3e cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/nf3e xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/nf3e CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/nf3e UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/nf3e cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/nf3e xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/nf3e fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/nf3e email,omitempty" json:"email,omitempty"`
}

type TEnteGov string

type TEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBSEstCred" json:"vCBSEstCred"`
}

type TFase string

type TFinNF3e string

type TFonteEnergia string

type TGrContrat string

type TGrMed string

type TGrpTensao string

type TIBSCBSMonoTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSUF" json:"vIBSUF"`
}

type TIBSCBSMonoTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMun" json:"vIBSMun"`
}

type TIBSCBSMonoTotGIBS struct {
	GIBSUF           TIBSCBSMonoTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nf3e gIBSUF" json:"gIBSUF"`
	GIBSMun          TIBSCBSMonoTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nf3e gIBSMun" json:"gIBSMun"`
	VIBS             TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nf3e vIBS" json:"vIBS"`
	VCredPres        TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nf3e vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/nf3e vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGCBS struct {
	VDif             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDif" json:"vDif"`
	VDevTrib         TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDevTrib" json:"vDevTrib"`
	VCBS             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBS" json:"vCBS"`
	VCredPres        TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGMono struct {
	VIBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMono" json:"vIBSMono"`
	VCBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBSMono" json:"vCBSMono"`
	VIBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMonoReten" json:"vIBSMonoReten"`
	VCBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBSMonoReten" json:"vCBSMonoReten"`
	VIBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMonoRet" json:"vIBSMonoRet"`
	VCBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBSMonoRet" json:"vCBSMonoRet"`
}

type TIBSCBSMonoTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSMonoTot struct {
	VBCIBSCBS    TDec1302RTC                 `xml:"http://www.portalfiscal.inf.br/nf3e vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         *TIBSCBSMonoTotGIBS         `xml:"http://www.portalfiscal.inf.br/nf3e gIBS,omitempty" json:"gIBS,omitempty"`
	GCBS         *TIBSCBSMonoTotGCBS         `xml:"http://www.portalfiscal.inf.br/nf3e gCBS,omitempty" json:"gCBS,omitempty"`
	GMono        *TIBSCBSMonoTotGMono        `xml:"http://www.portalfiscal.inf.br/nf3e gMono,omitempty" json:"gMono,omitempty"`
	GEstornoCred *TIBSCBSMonoTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIBSCBSTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSUF" json:"vIBSUF"`
}

type TIBSCBSTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMun" json:"vIBSMun"`
}

type TIBSCBSTotGIBS struct {
	GIBSUF  TIBSCBSTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/nf3e gIBSUF" json:"gIBSUF"`
	GIBSMun TIBSCBSTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/nf3e gIBSMun" json:"gIBSMun"`
	VIBS    TDec1302RTC           `xml:"http://www.portalfiscal.inf.br/nf3e vIBS" json:"vIBS"`
}

type TIBSCBSTotGCBS struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vDevTrib" json:"vDevTrib"`
	VCBS     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBS" json:"vCBS"`
}

type TIBSCBSTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSTot struct {
	VBCIBSCBS    TDec1302RTC             `xml:"http://www.portalfiscal.inf.br/nf3e vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         TIBSCBSTotGIBS          `xml:"http://www.portalfiscal.inf.br/nf3e gIBS" json:"gIBS"`
	GCBS         TIBSCBSTotGCBS          `xml:"http://www.portalfiscal.inf.br/nf3e gCBS" json:"gCBS"`
	GEstornoCred *TIBSCBSTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIPv4 string

type TIS struct {
	CSTIS        TCST           `xml:"http://www.portalfiscal.inf.br/nf3e CSTIS" json:"CSTIS"`
	CClassTribIS TcClassTrib    `xml:"http://www.portalfiscal.inf.br/nf3e cClassTribIS" json:"cClassTribIS"`
	VBCIS        *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vBCIS,omitempty" json:"vBCIS,omitempty"`
	PIS          *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pIS,omitempty" json:"pIS,omitempty"`
	AdRemIS      *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e adRemIS,omitempty" json:"adRemIS,omitempty"`
	UTrib        *string        `xml:"http://www.portalfiscal.inf.br/nf3e uTrib,omitempty" json:"uTrib,omitempty"`
	QTrib        *TDec1104OpRTC `xml:"http://www.portalfiscal.inf.br/nf3e qTrib,omitempty" json:"qTrib,omitempty"`
	VIS          *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIS,omitempty" json:"vIS,omitempty"`
}

type TISTot struct {
	VIS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIS" json:"vIS"`
}

type TIDLote string

type TIe string

type TIeDest string

type TImpICMS00 struct {
	CST   string          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	VBC   TDec1302        `xml:"http://www.portalfiscal.inf.br/nf3e vBC" json:"vBC"`
	PICMS TDec0302        `xml:"http://www.portalfiscal.inf.br/nf3e pICMS" json:"pICMS"`
	VICMS TDec1302        `xml:"http://www.portalfiscal.inf.br/nf3e vICMS" json:"vICMS"`
	PFCP  *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nf3e pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP  *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpICMS10 struct {
	CST     string          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	VBCST   TDec1302        `xml:"http://www.portalfiscal.inf.br/nf3e vBCST" json:"vBCST"`
	PICMSST TDec0302        `xml:"http://www.portalfiscal.inf.br/nf3e pICMSST" json:"pICMSST"`
	VICMSST TDec1302        `xml:"http://www.portalfiscal.inf.br/nf3e vICMSST" json:"vICMSST"`
	PFCPST  *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nf3e pFCPST,omitempty" json:"pFCPST,omitempty"`
	VFCPST  *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vFCPST,omitempty" json:"vFCPST,omitempty"`
}

type TImpICMS20 struct {
	CST        string          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	PRedBC     TDec0302Opc     `xml:"http://www.portalfiscal.inf.br/nf3e pRedBC" json:"pRedBC"`
	VBC        TDec1302        `xml:"http://www.portalfiscal.inf.br/nf3e vBC" json:"vBC"`
	PICMS      TDec0302        `xml:"http://www.portalfiscal.inf.br/nf3e pICMS" json:"pICMS"`
	VICMS      TDec1302        `xml:"http://www.portalfiscal.inf.br/nf3e vICMS" json:"vICMS"`
	VICMSDeson *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string         `xml:"http://www.portalfiscal.inf.br/nf3e cBenef,omitempty" json:"cBenef,omitempty"`
	PFCP       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nf3e pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpICMS40 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/nf3e cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS51 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/nf3e cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS60 struct {
	CST             string          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	VBCSTRet        *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vBCSTRet,omitempty" json:"vBCSTRet,omitempty"`
	PICMSSTRet      *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nf3e pICMSSTRet,omitempty" json:"pICMSSTRet,omitempty"`
	VICMSSubstituto *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vICMSSubstituto,omitempty" json:"vICMSSubstituto,omitempty"`
	VICMSSTRet      *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vICMSSTRet,omitempty" json:"vICMSSTRet,omitempty"`
	VBCFCPSTRet     *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vBCFCPSTRet,omitempty" json:"vBCFCPSTRet,omitempty"`
	PFCPSTRet       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nf3e pFCPSTRet,omitempty" json:"pFCPSTRet,omitempty"`
	VFCPSTRet       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vFCPSTRet,omitempty" json:"vFCPSTRet,omitempty"`
	PRedBCEfet      *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nf3e pRedBCEfet,omitempty" json:"pRedBCEfet,omitempty"`
	VBCEfet         *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vBCEfet,omitempty" json:"vBCEfet,omitempty"`
	PICMSEfet       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nf3e pICMSEfet,omitempty" json:"pICMSEfet,omitempty"`
	VICMSEfet       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vICMSEfet,omitempty" json:"vICMSEfet,omitempty"`
	VICMSDeson      *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef          *string         `xml:"http://www.portalfiscal.inf.br/nf3e cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS90 struct {
	CST        string          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	VBC        *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vBC,omitempty" json:"vBC,omitempty"`
	PICMS      *TDec0302       `xml:"http://www.portalfiscal.inf.br/nf3e pICMS,omitempty" json:"pICMS,omitempty"`
	VICMS      *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vICMS,omitempty" json:"vICMS,omitempty"`
	VICMSDeson *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string         `xml:"http://www.portalfiscal.inf.br/nf3e cBenef,omitempty" json:"cBenef,omitempty"`
	PFCP       *TDec0302a04Opc `xml:"http://www.portalfiscal.inf.br/nf3e pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP       *TDec1302       `xml:"http://www.portalfiscal.inf.br/nf3e vFCP,omitempty" json:"vFCP,omitempty"`
}

type TImpGICMSPrevistoPagtoAntecip struct {
	VICMSPrevisto TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vICMSPrevisto" json:"vICMSPrevisto"`
}

type TImpPIS struct {
	CST  string      `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	VBC  TDec1302    `xml:"http://www.portalfiscal.inf.br/nf3e vBC" json:"vBC"`
	PPIS TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nf3e pPIS" json:"pPIS"`
	VPIS TDec1302    `xml:"http://www.portalfiscal.inf.br/nf3e vPIS" json:"vPIS"`
}

type TImpPISEfet struct {
	VBCPISEfet TDec1302    `xml:"http://www.portalfiscal.inf.br/nf3e vBCPISEfet" json:"vBCPISEfet"`
	PPISEfet   TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nf3e pPISEfet" json:"pPISEfet"`
	VPISEfet   TDec1302    `xml:"http://www.portalfiscal.inf.br/nf3e vPISEfet" json:"vPISEfet"`
}

type TImpCOFINS struct {
	CST     string      `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	VBC     TDec1302    `xml:"http://www.portalfiscal.inf.br/nf3e vBC" json:"vBC"`
	PCOFINS TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nf3e pCOFINS" json:"pCOFINS"`
	VCOFINS TDec1302    `xml:"http://www.portalfiscal.inf.br/nf3e vCOFINS" json:"vCOFINS"`
}

type TImpCOFINSEfet struct {
	VBCCOFINSEfet TDec1302    `xml:"http://www.portalfiscal.inf.br/nf3e vBCCOFINSEfet" json:"vBCCOFINSEfet"`
	PCOFINSEfet   TDec0302a04 `xml:"http://www.portalfiscal.inf.br/nf3e pCOFINSEfet" json:"pCOFINSEfet"`
	VCOFINSEfet   TDec1302    `xml:"http://www.portalfiscal.inf.br/nf3e vCOFINSEfet" json:"vCOFINSEfet"`
}

type TImpRetTrib struct {
	VRetPIS    TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vRetCSLL" json:"vRetCSLL"`
	VBCIRRF    TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vBCIRRF" json:"vBCIRRF"`
	VIRRF      TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vIRRF" json:"vIRRF"`
}

type TImp struct {
	ICMS00                    *TImpICMS00                    `xml:"http://www.portalfiscal.inf.br/nf3e ICMS00,omitempty" json:"ICMS00,omitempty"`
	ICMS10                    *TImpICMS10                    `xml:"http://www.portalfiscal.inf.br/nf3e ICMS10,omitempty" json:"ICMS10,omitempty"`
	ICMS20                    *TImpICMS20                    `xml:"http://www.portalfiscal.inf.br/nf3e ICMS20,omitempty" json:"ICMS20,omitempty"`
	ICMS40                    *TImpICMS40                    `xml:"http://www.portalfiscal.inf.br/nf3e ICMS40,omitempty" json:"ICMS40,omitempty"`
	ICMS51                    *TImpICMS51                    `xml:"http://www.portalfiscal.inf.br/nf3e ICMS51,omitempty" json:"ICMS51,omitempty"`
	ICMS60                    *TImpICMS60                    `xml:"http://www.portalfiscal.inf.br/nf3e ICMS60,omitempty" json:"ICMS60,omitempty"`
	ICMS90                    *TImpICMS90                    `xml:"http://www.portalfiscal.inf.br/nf3e ICMS90,omitempty" json:"ICMS90,omitempty"`
	GICMSPrevistoPagtoAntecip *TImpGICMSPrevistoPagtoAntecip `xml:"http://www.portalfiscal.inf.br/nf3e gICMSPrevistoPagtoAntecip,omitempty" json:"gICMSPrevistoPagtoAntecip,omitempty"`
	IndSemCST                 *string                        `xml:"http://www.portalfiscal.inf.br/nf3e indSemCST,omitempty" json:"indSemCST,omitempty"`
	PIS                       *TImpPIS                       `xml:"http://www.portalfiscal.inf.br/nf3e PIS,omitempty" json:"PIS,omitempty"`
	PISEfet                   *TImpPISEfet                   `xml:"http://www.portalfiscal.inf.br/nf3e PISEfet,omitempty" json:"PISEfet,omitempty"`
	COFINS                    *TImpCOFINS                    `xml:"http://www.portalfiscal.inf.br/nf3e COFINS,omitempty" json:"COFINS,omitempty"`
	COFINSEfet                *TImpCOFINSEfet                `xml:"http://www.portalfiscal.inf.br/nf3e COFINSEfet,omitempty" json:"COFINSEfet,omitempty"`
	RetTrib                   *TImpRetTrib                   `xml:"http://www.portalfiscal.inf.br/nf3e retTrib,omitempty" json:"retTrib,omitempty"`
	IBSCBS                    *TTribNF3e                     `xml:"http://www.portalfiscal.inf.br/nf3e IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type TIndDoacao string

type TIndIEDest string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TMedida struct {
	TpGrMed       TGrMed            `xml:"http://www.portalfiscal.inf.br/nf3e tpGrMed" json:"tpGrMed"`
	CPosTarif     TPosTarifCompleto `xml:"http://www.portalfiscal.inf.br/nf3e cPosTarif" json:"cPosTarif"`
	UMed          TUMed             `xml:"http://www.portalfiscal.inf.br/nf3e uMed" json:"uMed"`
	VMedAnt       TDec13021304      `xml:"http://www.portalfiscal.inf.br/nf3e vMedAnt" json:"vMedAnt"`
	VMedAtu       TDec13021304      `xml:"http://www.portalfiscal.inf.br/nf3e vMedAtu" json:"vMedAtu"`
	VConst        TDec13021306      `xml:"http://www.portalfiscal.inf.br/nf3e vConst" json:"vConst"`
	VMed          TDec13021304      `xml:"http://www.portalfiscal.inf.br/nf3e vMed" json:"vMed"`
	PPerdaTran    *TDec0302         `xml:"http://www.portalfiscal.inf.br/nf3e pPerdaTran,omitempty" json:"pPerdaTran,omitempty"`
	VMedPerdaTran *TDec13021304     `xml:"http://www.portalfiscal.inf.br/nf3e vMedPerdaTran,omitempty" json:"vMedPerdaTran,omitempty"`
	VMedPerdaTec  *TDec13021304     `xml:"http://www.portalfiscal.inf.br/nf3e vMedPerdaTec,omitempty" json:"vMedPerdaTec,omitempty"`
}

type TModNF3e string

type TModTar string

type TMonofasiaGMonoPadrao struct {
	QBCMono  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nf3e qBCMono" json:"qBCMono"`
	AdRemIBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e adRemIBS" json:"adRemIBS"`
	AdRemCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e adRemCBS" json:"adRemCBS"`
	VIBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMono" json:"vIBSMono"`
	VCBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vCBSMono" json:"vCBSMono"`
}

type TMonofasiaGMonoReten struct {
	QBCMonoReten  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nf3e qBCMonoReten" json:"qBCMonoReten"`
	AdRemIBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e adRemIBSReten" json:"adRemIBSReten"`
	VIBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMonoReten" json:"vIBSMonoReten"`
	AdRemCBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e adRemCBSReten" json:"adRemCBSReten"`
	VCBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vCBSMonoReten" json:"vCBSMonoReten"`
}

type TMonofasiaGMonoRet struct {
	QBCMonoRet  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/nf3e qBCMonoRet" json:"qBCMonoRet"`
	AdRemIBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e adRemIBSRet" json:"adRemIBSRet"`
	VIBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMonoRet" json:"vIBSMonoRet"`
	AdRemCBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e adRemCBSRet" json:"adRemCBSRet"`
	VCBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vCBSMonoRet" json:"vCBSMonoRet"`
}

type TMonofasiaGMonoDif struct {
	PDifIBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pDifIBS" json:"pDifIBS"`
	VIBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMonoDif" json:"vIBSMonoDif"`
	PDifCBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pDifCBS" json:"pDifCBS"`
	VCBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vCBSMonoDif" json:"vCBSMonoDif"`
}

type TMonofasia struct {
	GMonoPadrao     *TMonofasiaGMonoPadrao `xml:"http://www.portalfiscal.inf.br/nf3e gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten      *TMonofasiaGMonoReten  `xml:"http://www.portalfiscal.inf.br/nf3e gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet        *TMonofasiaGMonoRet    `xml:"http://www.portalfiscal.inf.br/nf3e gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GMonoDif        *TMonofasiaGMonoDif    `xml:"http://www.portalfiscal.inf.br/nf3e gMonoDif,omitempty" json:"gMonoDif,omitempty"`
	VTotIBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/nf3e vTotIBSMonoItem" json:"vTotIBSMonoItem"`
	VTotCBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/nf3e vTotCBSMonoItem" json:"vTotCBSMonoItem"`
}

type TMotAjuste string

type TMotNaoLeitura string

type TMotSub string

type TMotivo string

type TNF string

type TNF3eInfNF3eIde struct {
	CUF          TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/nf3e cUF" json:"cUF"`
	TpAmb        TAmb                `xml:"http://www.portalfiscal.inf.br/nf3e tpAmb" json:"tpAmb"`
	Mod          TModNF3e            `xml:"http://www.portalfiscal.inf.br/nf3e mod" json:"mod"`
	Serie        TSerie              `xml:"http://www.portalfiscal.inf.br/nf3e serie" json:"serie"`
	NNF          TNF                 `xml:"http://www.portalfiscal.inf.br/nf3e nNF" json:"nNF"`
	CNF          string              `xml:"http://www.portalfiscal.inf.br/nf3e cNF" json:"cNF"`
	CDV          string              `xml:"http://www.portalfiscal.inf.br/nf3e cDV" json:"cDV"`
	DhEmi        TDateTimeUTC        `xml:"http://www.portalfiscal.inf.br/nf3e dhEmi" json:"dhEmi"`
	TpEmis       TtpEmis             `xml:"http://www.portalfiscal.inf.br/nf3e tpEmis" json:"tpEmis"`
	NSiteAutoriz string              `xml:"http://www.portalfiscal.inf.br/nf3e nSiteAutoriz" json:"nSiteAutoriz"`
	CMunFG       TCodMunIBGE         `xml:"http://www.portalfiscal.inf.br/nf3e cMunFG" json:"cMunFG"`
	FinNF3e      TFinNF3e            `xml:"http://www.portalfiscal.inf.br/nf3e finNF3e" json:"finNF3e"`
	VerProc      string              `xml:"http://www.portalfiscal.inf.br/nf3e verProc" json:"verProc"`
	DhCont       *TDateTimeUTC       `xml:"http://www.portalfiscal.inf.br/nf3e dhCont,omitempty" json:"dhCont,omitempty"`
	XJust        *string             `xml:"http://www.portalfiscal.inf.br/nf3e xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov   *TCompraGovReduzido `xml:"http://www.portalfiscal.inf.br/nf3e gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	TpPagAnt     *TPagAnt            `xml:"http://www.portalfiscal.inf.br/nf3e tpPagAnt,omitempty" json:"tpPagAnt,omitempty"`
}

type TNF3eInfNF3eEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/nf3e CNPJ" json:"CNPJ"`
	IE        TIe      `xml:"http://www.portalfiscal.inf.br/nf3e IE" json:"IE"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/nf3e xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/nf3e xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/nf3e enderEmit" json:"enderEmit"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/nf3e ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TNF3eInfNF3eDest struct {
	XNome          string     `xml:"http://www.portalfiscal.inf.br/nf3e xNome" json:"xNome"`
	CNPJ           *TCNPJOpc  `xml:"http://www.portalfiscal.inf.br/nf3e CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF            *TCPF      `xml:"http://www.portalfiscal.inf.br/nf3e CPF,omitempty" json:"CPF,omitempty"`
	IDOutros       *string    `xml:"http://www.portalfiscal.inf.br/nf3e idOutros,omitempty" json:"idOutros,omitempty"`
	IndIEDest      TIndIEDest `xml:"http://www.portalfiscal.inf.br/nf3e indIEDest" json:"indIEDest"`
	IE             *TIeDest   `xml:"http://www.portalfiscal.inf.br/nf3e IE,omitempty" json:"IE,omitempty"`
	IM             *string    `xml:"http://www.portalfiscal.inf.br/nf3e IM,omitempty" json:"IM,omitempty"`
	CNIS           *string    `xml:"http://www.portalfiscal.inf.br/nf3e cNIS,omitempty" json:"cNIS,omitempty"`
	NB             *string    `xml:"http://www.portalfiscal.inf.br/nf3e NB,omitempty" json:"NB,omitempty"`
	XNomeAdicional *string    `xml:"http://www.portalfiscal.inf.br/nf3e xNomeAdicional,omitempty" json:"xNomeAdicional,omitempty"`
	EnderDest      TEndeEmi   `xml:"http://www.portalfiscal.inf.br/nf3e enderDest" json:"enderDest"`
}

type TNF3eInfNF3eAcessante struct {
	IDAcesso          string      `xml:"http://www.portalfiscal.inf.br/nf3e idAcesso" json:"idAcesso"`
	IDCodCliente      *string     `xml:"http://www.portalfiscal.inf.br/nf3e idCodCliente,omitempty" json:"idCodCliente,omitempty"`
	TpAcesso          TAcessante  `xml:"http://www.portalfiscal.inf.br/nf3e tpAcesso" json:"tpAcesso"`
	XNomeUC           *string     `xml:"http://www.portalfiscal.inf.br/nf3e xNomeUC,omitempty" json:"xNomeUC,omitempty"`
	TpClasse          *TClasse    `xml:"http://www.portalfiscal.inf.br/nf3e tpClasse,omitempty" json:"tpClasse,omitempty"`
	TpSubClasse       *TSubClasse `xml:"http://www.portalfiscal.inf.br/nf3e tpSubClasse,omitempty" json:"tpSubClasse,omitempty"`
	TpFase            TFase       `xml:"http://www.portalfiscal.inf.br/nf3e tpFase" json:"tpFase"`
	TpGrpTensao       TGrpTensao  `xml:"http://www.portalfiscal.inf.br/nf3e tpGrpTensao" json:"tpGrpTensao"`
	TpModTar          TModTar     `xml:"http://www.portalfiscal.inf.br/nf3e tpModTar" json:"tpModTar"`
	LatGPS            TLatitude   `xml:"http://www.portalfiscal.inf.br/nf3e latGPS" json:"latGPS"`
	LongGPS           TLongitude  `xml:"http://www.portalfiscal.inf.br/nf3e longGPS" json:"longGPS"`
	CodRoteiroLeitura *string     `xml:"http://www.portalfiscal.inf.br/nf3e codRoteiroLeitura,omitempty" json:"codRoteiroLeitura,omitempty"`
}

type TNF3eInfNF3eGSubGNF struct {
	CNPJ       TCNPJ   `xml:"http://www.portalfiscal.inf.br/nf3e CNPJ" json:"CNPJ"`
	Serie      string  `xml:"http://www.portalfiscal.inf.br/nf3e serie" json:"serie"`
	NNF        TNF     `xml:"http://www.portalfiscal.inf.br/nf3e nNF" json:"nNF"`
	CompetEmis TCompet `xml:"http://www.portalfiscal.inf.br/nf3e CompetEmis" json:"CompetEmis"`
	CompetApur TCompet `xml:"http://www.portalfiscal.inf.br/nf3e CompetApur" json:"CompetApur"`
	Hash115    *string `xml:"http://www.portalfiscal.inf.br/nf3e hash115,omitempty" json:"hash115,omitempty"`
}

type TNF3eInfNF3eGSub struct {
	ChNF3e *TChDFe              `xml:"http://www.portalfiscal.inf.br/nf3e chNF3e,omitempty" json:"chNF3e,omitempty"`
	GNF    *TNF3eInfNF3eGSubGNF `xml:"http://www.portalfiscal.inf.br/nf3e gNF,omitempty" json:"gNF,omitempty"`
	MotSub TMotSub              `xml:"http://www.portalfiscal.inf.br/nf3e motSub" json:"motSub"`
}

type TNF3eInfNF3eGJudic struct {
	ChNF3e TChDFe `xml:"http://www.portalfiscal.inf.br/nf3e chNF3e" json:"chNF3e"`
}

type TNF3eInfNF3eGGrContrat struct {
	NContrat     string         `xml:"nContrat,attr" json:"nContrat"`
	TpGrContrat  TGrContrat     `xml:"http://www.portalfiscal.inf.br/nf3e tpGrContrat" json:"tpGrContrat"`
	TpPosTar     TPosTarSimples `xml:"http://www.portalfiscal.inf.br/nf3e tpPosTar" json:"tpPosTar"`
	QUnidContrat TDec13021306   `xml:"http://www.portalfiscal.inf.br/nf3e qUnidContrat" json:"qUnidContrat"`
}

type TNF3eInfNF3eGMed struct {
	NMed      string `xml:"nMed,attr" json:"nMed"`
	IDMedidor string `xml:"http://www.portalfiscal.inf.br/nf3e idMedidor" json:"idMedidor"`
	DMedAnt   TData  `xml:"http://www.portalfiscal.inf.br/nf3e dMedAnt" json:"dMedAnt"`
	DMedAtu   TData  `xml:"http://www.portalfiscal.inf.br/nf3e dMedAtu" json:"dMedAtu"`
}

type TNF3eInfNF3eGSCEEGConsumidor struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type TNF3eInfNF3eGSCEEGSaldoCred struct {
	TpPosTar      TPosTarSimples `xml:"http://www.portalfiscal.inf.br/nf3e tpPosTar" json:"tpPosTar"`
	VSaldAnt      TDec11001104   `xml:"http://www.portalfiscal.inf.br/nf3e vSaldAnt" json:"vSaldAnt"`
	VCredExpirado TDec11001104   `xml:"http://www.portalfiscal.inf.br/nf3e vCredExpirado" json:"vCredExpirado"`
	VSaldAtual    TDec11001104   `xml:"http://www.portalfiscal.inf.br/nf3e vSaldAtual" json:"vSaldAtual"`
	VCredExpirar  *TDec11001104  `xml:"http://www.portalfiscal.inf.br/nf3e vCredExpirar,omitempty" json:"vCredExpirar,omitempty"`
	CompetExpirar *TCompet       `xml:"http://www.portalfiscal.inf.br/nf3e CompetExpirar,omitempty" json:"CompetExpirar,omitempty"`
}

type TNF3eInfNF3eGSCEEGTipoSaldoGSaldoCred struct {
	TpPosTar      TPosTarSimples `xml:"http://www.portalfiscal.inf.br/nf3e tpPosTar" json:"tpPosTar"`
	VSaldAnt      TDec11001104   `xml:"http://www.portalfiscal.inf.br/nf3e vSaldAnt" json:"vSaldAnt"`
	VCredExpirado TDec11001104   `xml:"http://www.portalfiscal.inf.br/nf3e vCredExpirado" json:"vCredExpirado"`
	VSaldAtual    TDec11001104   `xml:"http://www.portalfiscal.inf.br/nf3e vSaldAtual" json:"vSaldAtual"`
	VCredExpirar  *TDec11001104  `xml:"http://www.portalfiscal.inf.br/nf3e vCredExpirar,omitempty" json:"vCredExpirar,omitempty"`
	CompetExpirar *TCompet       `xml:"http://www.portalfiscal.inf.br/nf3e CompetExpirar,omitempty" json:"CompetExpirar,omitempty"`
}

type TNF3eInfNF3eGSCEEGTipoSaldo struct {
	NTipoSaldo string                                  `xml:"nTipoSaldo,attr" json:"nTipoSaldo"`
	GSaldoCred []TNF3eInfNF3eGSCEEGTipoSaldoGSaldoCred `xml:"http://www.portalfiscal.inf.br/nf3e gSaldoCred,omitempty" json:"gSaldoCred,omitempty"`
}

type TNF3eInfNF3eGSCEE struct {
	TpPartComp  TPartComp                      `xml:"http://www.portalfiscal.inf.br/nf3e tpPartComp" json:"tpPartComp"`
	GConsumidor []TNF3eInfNF3eGSCEEGConsumidor `xml:"http://www.portalfiscal.inf.br/nf3e gConsumidor,omitempty" json:"gConsumidor,omitempty"`
	GSaldoCred  []TNF3eInfNF3eGSCEEGSaldoCred  `xml:"http://www.portalfiscal.inf.br/nf3e gSaldoCred,omitempty" json:"gSaldoCred,omitempty"`
	GTipoSaldo  []TNF3eInfNF3eGSCEEGTipoSaldo  `xml:"http://www.portalfiscal.inf.br/nf3e gTipoSaldo,omitempty" json:"gTipoSaldo,omitempty"`
}

type TNF3eInfNF3eNFdetDetGAjusteNF3eAnt struct {
	TpAjuste  TAjusteNF3eAnt `xml:"http://www.portalfiscal.inf.br/nf3e tpAjuste" json:"tpAjuste"`
	MotAjuste TMotAjuste     `xml:"http://www.portalfiscal.inf.br/nf3e motAjuste" json:"motAjuste"`
}

type TNF3eInfNF3eNFdetDetDetItemAntRetTrib struct {
	VRetPIS    TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vRetCSLL" json:"vRetCSLL"`
	VBCIRRF    TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vBCIRRF" json:"vBCIRRF"`
	VIRRF      TDec13021308 `xml:"http://www.portalfiscal.inf.br/nf3e vIRRF" json:"vIRRF"`
}

type TNF3eInfNF3eNFdetDetDetItemAnt struct {
	NItemAnt     string                                 `xml:"nItemAnt,attr" json:"nItemAnt"`
	VItem        TDec13021310                           `xml:"http://www.portalfiscal.inf.br/nf3e vItem" json:"vItem"`
	QFaturada    TDec11001104                           `xml:"http://www.portalfiscal.inf.br/nf3e qFaturada" json:"qFaturada"`
	VProd        TDec13021310                           `xml:"http://www.portalfiscal.inf.br/nf3e vProd" json:"vProd"`
	CClass       string                                 `xml:"http://www.portalfiscal.inf.br/nf3e cClass" json:"cClass"`
	VBC          *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vBC,omitempty" json:"vBC,omitempty"`
	PICMS        *TDec0302                              `xml:"http://www.portalfiscal.inf.br/nf3e pICMS,omitempty" json:"pICMS,omitempty"`
	VICMS        *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vICMS,omitempty" json:"vICMS,omitempty"`
	VFCP         *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vFCP,omitempty" json:"vFCP,omitempty"`
	VBCST        *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vBCST,omitempty" json:"vBCST,omitempty"`
	VICMSST      *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vICMSST,omitempty" json:"vICMSST,omitempty"`
	VFCPST       *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vFCPST,omitempty" json:"vFCPST,omitempty"`
	VPIS         *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vPIS,omitempty" json:"vPIS,omitempty"`
	VPISEfet     *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vPISEfet,omitempty" json:"vPISEfet,omitempty"`
	VCOFINS      *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vCOFINS,omitempty" json:"vCOFINS,omitempty"`
	VCOFINSEfet  *TDec1302                              `xml:"http://www.portalfiscal.inf.br/nf3e vCOFINSEfet,omitempty" json:"vCOFINSEfet,omitempty"`
	RetTrib      *TNF3eInfNF3eNFdetDetDetItemAntRetTrib `xml:"http://www.portalfiscal.inf.br/nf3e retTrib,omitempty" json:"retTrib,omitempty"`
	IndDevolucao *string                                `xml:"http://www.portalfiscal.inf.br/nf3e indDevolucao,omitempty" json:"indDevolucao,omitempty"`
}

type TNF3eInfNF3eNFdetDetDetItemGTarif struct {
	DIniTarif   TData             `xml:"http://www.portalfiscal.inf.br/nf3e dIniTarif" json:"dIniTarif"`
	DFimTarif   TData             `xml:"http://www.portalfiscal.inf.br/nf3e dFimTarif" json:"dFimTarif"`
	TpAto       TAtoANEEL         `xml:"http://www.portalfiscal.inf.br/nf3e tpAto" json:"tpAto"`
	NAto        string            `xml:"http://www.portalfiscal.inf.br/nf3e nAto" json:"nAto"`
	AnoAto      string            `xml:"http://www.portalfiscal.inf.br/nf3e anoAto" json:"anoAto"`
	TpTarif     TTarifa           `xml:"http://www.portalfiscal.inf.br/nf3e tpTarif" json:"tpTarif"`
	CPosTarif   TPosTarifCompleto `xml:"http://www.portalfiscal.inf.br/nf3e cPosTarif" json:"cPosTarif"`
	UMed        string            `xml:"http://www.portalfiscal.inf.br/nf3e uMed" json:"uMed"`
	VTarifHom   string            `xml:"http://www.portalfiscal.inf.br/nf3e vTarifHom" json:"vTarifHom"`
	VTarifAplic *string           `xml:"http://www.portalfiscal.inf.br/nf3e vTarifAplic,omitempty" json:"vTarifAplic,omitempty"`
	MotDifTarif *string           `xml:"http://www.portalfiscal.inf.br/nf3e motDifTarif,omitempty" json:"motDifTarif,omitempty"`
}

type TNF3eInfNF3eNFdetDetDetItemGAdBand struct {
	DIniAdBand   TData         `xml:"http://www.portalfiscal.inf.br/nf3e dIniAdBand" json:"dIniAdBand"`
	DFimAdBand   TData         `xml:"http://www.portalfiscal.inf.br/nf3e dFimAdBand" json:"dFimAdBand"`
	TpBand       TBand         `xml:"http://www.portalfiscal.inf.br/nf3e tpBand" json:"tpBand"`
	VAdBand      TDec13021310  `xml:"http://www.portalfiscal.inf.br/nf3e vAdBand" json:"vAdBand"`
	VAdBandAplic *TDec13021310 `xml:"http://www.portalfiscal.inf.br/nf3e vAdBandAplic,omitempty" json:"vAdBandAplic,omitempty"`
	MotDifBand   *string       `xml:"http://www.portalfiscal.inf.br/nf3e motDifBand,omitempty" json:"motDifBand,omitempty"`
}

type TNF3eInfNF3eNFdetDetDetItemProdGMedicao struct {
	NMed            string          `xml:"http://www.portalfiscal.inf.br/nf3e nMed" json:"nMed"`
	NContrat        *string         `xml:"http://www.portalfiscal.inf.br/nf3e nContrat,omitempty" json:"nContrat,omitempty"`
	GMedida         *TMedida        `xml:"http://www.portalfiscal.inf.br/nf3e gMedida,omitempty" json:"gMedida,omitempty"`
	TpMotNaoLeitura *TMotNaoLeitura `xml:"http://www.portalfiscal.inf.br/nf3e tpMotNaoLeitura,omitempty" json:"tpMotNaoLeitura,omitempty"`
}

type TNF3eInfNF3eNFdetDetDetItemProdGPagAntecipado struct {
	ChDFePagAnt TChDFe  `xml:"http://www.portalfiscal.inf.br/nf3e chDFePagAnt" json:"chDFePagAnt"`
	NItemPagAnt *string `xml:"http://www.portalfiscal.inf.br/nf3e nItemPagAnt,omitempty" json:"nItemPagAnt,omitempty"`
}

type TNF3eInfNF3eNFdetDetDetItemProd struct {
	IndOrigemQtd   TOrigemQtd                                     `xml:"http://www.portalfiscal.inf.br/nf3e indOrigemQtd" json:"indOrigemQtd"`
	GMedicao       *TNF3eInfNF3eNFdetDetDetItemProdGMedicao       `xml:"http://www.portalfiscal.inf.br/nf3e gMedicao,omitempty" json:"gMedicao,omitempty"`
	CProd          string                                         `xml:"http://www.portalfiscal.inf.br/nf3e cProd" json:"cProd"`
	XProd          string                                         `xml:"http://www.portalfiscal.inf.br/nf3e xProd" json:"xProd"`
	CClass         string                                         `xml:"http://www.portalfiscal.inf.br/nf3e cClass" json:"cClass"`
	CFOP           *TCFOP                                         `xml:"http://www.portalfiscal.inf.br/nf3e CFOP,omitempty" json:"CFOP,omitempty"`
	UMed           TUMedItem                                      `xml:"http://www.portalfiscal.inf.br/nf3e uMed" json:"uMed"`
	QFaturada      TDec11001104                                   `xml:"http://www.portalfiscal.inf.br/nf3e qFaturada" json:"qFaturada"`
	VItem          TDec13021310                                   `xml:"http://www.portalfiscal.inf.br/nf3e vItem" json:"vItem"`
	VProd          TDec13021310                                   `xml:"http://www.portalfiscal.inf.br/nf3e vProd" json:"vProd"`
	VItemLiq       *TDec13021310                                  `xml:"http://www.portalfiscal.inf.br/nf3e vItemLiq,omitempty" json:"vItemLiq,omitempty"`
	VProdLiq       *TDec13021310                                  `xml:"http://www.portalfiscal.inf.br/nf3e vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	IndDevolucao   *string                                        `xml:"http://www.portalfiscal.inf.br/nf3e indDevolucao,omitempty" json:"indDevolucao,omitempty"`
	IndPrecoACL    *string                                        `xml:"http://www.portalfiscal.inf.br/nf3e indPrecoACL,omitempty" json:"indPrecoACL,omitempty"`
	GPagAntecipado *TNF3eInfNF3eNFdetDetDetItemProdGPagAntecipado `xml:"http://www.portalfiscal.inf.br/nf3e gPagAntecipado,omitempty" json:"gPagAntecipado,omitempty"`
}

type TNF3eInfNF3eNFdetDetDetItemGProcRefGIBSCBS struct {
	VBC     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vBC" json:"vBC"`
	VIBSUF  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSUF" json:"vIBSUF"`
	VIBSMun TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSMun" json:"vIBSMun"`
	VCBS    TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBS" json:"vCBS"`
}

type TNF3eInfNF3eNFdetDetDetItemGProcRefGProc struct {
	TpProc    TProcesso `xml:"http://www.portalfiscal.inf.br/nf3e tpProc" json:"tpProc"`
	NProcesso string    `xml:"http://www.portalfiscal.inf.br/nf3e nProcesso" json:"nProcesso"`
}

type TNF3eInfNF3eNFdetDetDetItemGProcRef struct {
	VItem        TDec13021310                                `xml:"http://www.portalfiscal.inf.br/nf3e vItem" json:"vItem"`
	QFaturada    TDec11001104                                `xml:"http://www.portalfiscal.inf.br/nf3e qFaturada" json:"qFaturada"`
	VProd        TDec13021310                                `xml:"http://www.portalfiscal.inf.br/nf3e vProd" json:"vProd"`
	IndDevolucao *string                                     `xml:"http://www.portalfiscal.inf.br/nf3e indDevolucao,omitempty" json:"indDevolucao,omitempty"`
	VBC          *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vBC,omitempty" json:"vBC,omitempty"`
	PICMS        *TDec0302                                   `xml:"http://www.portalfiscal.inf.br/nf3e pICMS,omitempty" json:"pICMS,omitempty"`
	VICMS        *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vICMS,omitempty" json:"vICMS,omitempty"`
	PFCP         *TDec0302a04Opc                             `xml:"http://www.portalfiscal.inf.br/nf3e pFCP,omitempty" json:"pFCP,omitempty"`
	VFCP         *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vFCP,omitempty" json:"vFCP,omitempty"`
	VBCST        *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vBCST,omitempty" json:"vBCST,omitempty"`
	PICMSST      *TDec0302                                   `xml:"http://www.portalfiscal.inf.br/nf3e pICMSST,omitempty" json:"pICMSST,omitempty"`
	VICMSST      *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vICMSST,omitempty" json:"vICMSST,omitempty"`
	PFCPST       *TDec0302a04Opc                             `xml:"http://www.portalfiscal.inf.br/nf3e pFCPST,omitempty" json:"pFCPST,omitempty"`
	VFCPST       *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vFCPST,omitempty" json:"vFCPST,omitempty"`
	VPIS         *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vPIS,omitempty" json:"vPIS,omitempty"`
	VPISEfet     *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vPISEfet,omitempty" json:"vPISEfet,omitempty"`
	VCOFINS      *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vCOFINS,omitempty" json:"vCOFINS,omitempty"`
	VCOFINSEfet  *TDec1302                                   `xml:"http://www.portalfiscal.inf.br/nf3e vCOFINSEfet,omitempty" json:"vCOFINSEfet,omitempty"`
	GIBSCBS      *TNF3eInfNF3eNFdetDetDetItemGProcRefGIBSCBS `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GProc        []TNF3eInfNF3eNFdetDetDetItemGProcRefGProc  `xml:"http://www.portalfiscal.inf.br/nf3e gProc,omitempty" json:"gProc,omitempty"`
}

type TNF3eInfNF3eNFdetDetDetItemGContab struct {
	CContab string   `xml:"http://www.portalfiscal.inf.br/nf3e cContab" json:"cContab"`
	XContab string   `xml:"http://www.portalfiscal.inf.br/nf3e xContab" json:"xContab"`
	VContab TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vContab" json:"vContab"`
	TpLanc  string   `xml:"http://www.portalfiscal.inf.br/nf3e tpLanc" json:"tpLanc"`
}

type TNF3eInfNF3eNFdetDetDetItem struct {
	NItemAnt  *string                              `xml:"nItemAnt,attr,omitempty" json:"nItemAnt,omitempty"`
	GTarif    []TNF3eInfNF3eNFdetDetDetItemGTarif  `xml:"http://www.portalfiscal.inf.br/nf3e gTarif,omitempty" json:"gTarif,omitempty"`
	GAdBand   []TNF3eInfNF3eNFdetDetDetItemGAdBand `xml:"http://www.portalfiscal.inf.br/nf3e gAdBand,omitempty" json:"gAdBand,omitempty"`
	Prod      TNF3eInfNF3eNFdetDetDetItemProd      `xml:"http://www.portalfiscal.inf.br/nf3e prod" json:"prod"`
	Imposto   TImp                                 `xml:"http://www.portalfiscal.inf.br/nf3e imposto" json:"imposto"`
	GProcRef  *TNF3eInfNF3eNFdetDetDetItemGProcRef `xml:"http://www.portalfiscal.inf.br/nf3e gProcRef,omitempty" json:"gProcRef,omitempty"`
	GContab   []TNF3eInfNF3eNFdetDetDetItemGContab `xml:"http://www.portalfiscal.inf.br/nf3e gContab,omitempty" json:"gContab,omitempty"`
	InfAdProd *string                              `xml:"http://www.portalfiscal.inf.br/nf3e infAdProd,omitempty" json:"infAdProd,omitempty"`
}

type TNF3eInfNF3eNFdetDet struct {
	NItem          string                              `xml:"nItem,attr" json:"nItem"`
	GAjusteNF3eAnt *TNF3eInfNF3eNFdetDetGAjusteNF3eAnt `xml:"http://www.portalfiscal.inf.br/nf3e gAjusteNF3eAnt,omitempty" json:"gAjusteNF3eAnt,omitempty"`
	DetItemAnt     *TNF3eInfNF3eNFdetDetDetItemAnt     `xml:"http://www.portalfiscal.inf.br/nf3e detItemAnt,omitempty" json:"detItemAnt,omitempty"`
	DetItem        *TNF3eInfNF3eNFdetDetDetItem        `xml:"http://www.portalfiscal.inf.br/nf3e detItem,omitempty" json:"detItem,omitempty"`
}

type TNF3eInfNF3eNFdet struct {
	ChNF3eAnt   *TChDFe                `xml:"chNF3eAnt,attr,omitempty" json:"chNF3eAnt,omitempty"`
	Mod6HashAnt *Base64Binary          `xml:"mod6HashAnt,attr,omitempty" json:"mod6HashAnt,omitempty"`
	Det         []TNF3eInfNF3eNFdetDet `xml:"http://www.portalfiscal.inf.br/nf3e det,omitempty" json:"det,omitempty"`
}

type TNF3eInfNF3eTotalICMSTot struct {
	VBC        TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vBC" json:"vBC"`
	VICMS      TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vICMS" json:"vICMS"`
	VICMSDeson TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vICMSDeson" json:"vICMSDeson"`
	VFCP       TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vFCP" json:"vFCP"`
	VBCST      TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vBCST" json:"vBCST"`
	VST        TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vST" json:"vST"`
	VFCPST     TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vFCPST" json:"vFCPST"`
}

type TNF3eInfNF3eTotalVRetTribTot struct {
	VRetPIS    TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vRetPIS" json:"vRetPIS"`
	VRetCOFINS TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vRetCofins" json:"vRetCofins"`
	VRetCSLL   TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vRetCSLL" json:"vRetCSLL"`
	VIRRF      TDec1302 `xml:"http://www.portalfiscal.inf.br/nf3e vIRRF" json:"vIRRF"`
}

type TNF3eInfNF3eTotal struct {
	VProd       TDec1302                     `xml:"http://www.portalfiscal.inf.br/nf3e vProd" json:"vProd"`
	VProdLiq    *TDec1302                    `xml:"http://www.portalfiscal.inf.br/nf3e vProdLiq,omitempty" json:"vProdLiq,omitempty"`
	ICMSTot     TNF3eInfNF3eTotalICMSTot     `xml:"http://www.portalfiscal.inf.br/nf3e ICMSTot" json:"ICMSTot"`
	VRetTribTot TNF3eInfNF3eTotalVRetTribTot `xml:"http://www.portalfiscal.inf.br/nf3e vRetTribTot" json:"vRetTribTot"`
	VCOFINS     TDec1302                     `xml:"http://www.portalfiscal.inf.br/nf3e vCOFINS" json:"vCOFINS"`
	VCOFINSEfet TDec1302                     `xml:"http://www.portalfiscal.inf.br/nf3e vCOFINSEfet" json:"vCOFINSEfet"`
	VPIS        TDec1302                     `xml:"http://www.portalfiscal.inf.br/nf3e vPIS" json:"vPIS"`
	VPISEfet    TDec1302                     `xml:"http://www.portalfiscal.inf.br/nf3e vPISEfet" json:"vPISEfet"`
	VNF         TDec1302                     `xml:"http://www.portalfiscal.inf.br/nf3e vNF" json:"vNF"`
	IBSCBSTot   *TIBSCBSTot                  `xml:"http://www.portalfiscal.inf.br/nf3e IBSCBSTot,omitempty" json:"IBSCBSTot,omitempty"`
	VTotDFe     *TDec1302                    `xml:"http://www.portalfiscal.inf.br/nf3e vTotDFe,omitempty" json:"vTotDFe,omitempty"`
}

type TNF3eInfNF3ePgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/nf3e pgto,omitempty" json:"pgto,omitempty"`
}

type TNF3eInfNF3eGFatGPIX struct {
	URLQRCodePIX string `xml:"http://www.portalfiscal.inf.br/nf3e urlQRCodePIX" json:"urlQRCodePIX"`
}

type TNF3eInfNF3eGFat struct {
	CompetFat    TCompet               `xml:"http://www.portalfiscal.inf.br/nf3e CompetFat" json:"CompetFat"`
	DVencFat     TData                 `xml:"http://www.portalfiscal.inf.br/nf3e dVencFat" json:"dVencFat"`
	DApresFat    *TData                `xml:"http://www.portalfiscal.inf.br/nf3e dApresFat,omitempty" json:"dApresFat,omitempty"`
	DProxLeitura TData                 `xml:"http://www.portalfiscal.inf.br/nf3e dProxLeitura" json:"dProxLeitura"`
	NFat         *string               `xml:"http://www.portalfiscal.inf.br/nf3e nFat,omitempty" json:"nFat,omitempty"`
	CodBarras    string                `xml:"http://www.portalfiscal.inf.br/nf3e codBarras" json:"codBarras"`
	CodDebAuto   *string               `xml:"http://www.portalfiscal.inf.br/nf3e codDebAuto,omitempty" json:"codDebAuto,omitempty"`
	CodBanco     *string               `xml:"http://www.portalfiscal.inf.br/nf3e codBanco,omitempty" json:"codBanco,omitempty"`
	CodAgencia   *string               `xml:"http://www.portalfiscal.inf.br/nf3e codAgencia,omitempty" json:"codAgencia,omitempty"`
	EnderCorresp *TEndeEmi             `xml:"http://www.portalfiscal.inf.br/nf3e enderCorresp,omitempty" json:"enderCorresp,omitempty"`
	GPIX         *TNF3eInfNF3eGFatGPIX `xml:"http://www.portalfiscal.inf.br/nf3e gPIX,omitempty" json:"gPIX,omitempty"`
}

type TNF3eInfNF3eAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/nf3e CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/nf3e CPF,omitempty" json:"CPF,omitempty"`
}

type TNF3eInfNF3eInfAdic struct {
	InfAdFisco *string  `xml:"http://www.portalfiscal.inf.br/nf3e infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     []string `xml:"http://www.portalfiscal.inf.br/nf3e infCpl,omitempty" json:"infCpl,omitempty"`
}

type TNF3eInfNF3e struct {
	Versao     TVerNF3e                 `xml:"versao,attr" json:"versao"`
	ID         string                   `xml:"Id,attr" json:"Id"`
	Ide        TNF3eInfNF3eIde          `xml:"http://www.portalfiscal.inf.br/nf3e ide" json:"ide"`
	Emit       TNF3eInfNF3eEmit         `xml:"http://www.portalfiscal.inf.br/nf3e emit" json:"emit"`
	Dest       TNF3eInfNF3eDest         `xml:"http://www.portalfiscal.inf.br/nf3e dest" json:"dest"`
	Acessante  TNF3eInfNF3eAcessante    `xml:"http://www.portalfiscal.inf.br/nf3e acessante" json:"acessante"`
	GSub       *TNF3eInfNF3eGSub        `xml:"http://www.portalfiscal.inf.br/nf3e gSub,omitempty" json:"gSub,omitempty"`
	GJudic     *TNF3eInfNF3eGJudic      `xml:"http://www.portalfiscal.inf.br/nf3e gJudic,omitempty" json:"gJudic,omitempty"`
	GGrContrat []TNF3eInfNF3eGGrContrat `xml:"http://www.portalfiscal.inf.br/nf3e gGrContrat,omitempty" json:"gGrContrat,omitempty"`
	GMed       []TNF3eInfNF3eGMed       `xml:"http://www.portalfiscal.inf.br/nf3e gMed,omitempty" json:"gMed,omitempty"`
	GSCEE      *TNF3eInfNF3eGSCEE       `xml:"http://www.portalfiscal.inf.br/nf3e gSCEE,omitempty" json:"gSCEE,omitempty"`
	NFdet      []TNF3eInfNF3eNFdet      `xml:"http://www.portalfiscal.inf.br/nf3e NFdet,omitempty" json:"NFdet,omitempty"`
	Total      TNF3eInfNF3eTotal        `xml:"http://www.portalfiscal.inf.br/nf3e total" json:"total"`
	PgtoVinc   *TNF3eInfNF3ePgtoVinc    `xml:"http://www.portalfiscal.inf.br/nf3e pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	GFat       TNF3eInfNF3eGFat         `xml:"http://www.portalfiscal.inf.br/nf3e gFat" json:"gFat"`
	GANEEL     TANEEL                   `xml:"http://www.portalfiscal.inf.br/nf3e gANEEL" json:"gANEEL"`
	AutXML     []TNF3eInfNF3eAutXML     `xml:"http://www.portalfiscal.inf.br/nf3e autXML,omitempty" json:"autXML,omitempty"`
	InfAdic    *TNF3eInfNF3eInfAdic     `xml:"http://www.portalfiscal.inf.br/nf3e infAdic,omitempty" json:"infAdic,omitempty"`
	GRespTec   *TRespTec                `xml:"http://www.portalfiscal.inf.br/nf3e gRespTec,omitempty" json:"gRespTec,omitempty"`
}

type TNF3eInfNF3eSupl struct {
	QRCodNF3e string `xml:"http://www.portalfiscal.inf.br/nf3e qrCodNF3e" json:"qrCodNF3e"`
}

type TNF3e struct {
	InfNF3e     TNF3eInfNF3e     `xml:"http://www.portalfiscal.inf.br/nf3e infNF3e" json:"infNF3e"`
	InfNF3eSupl TNF3eInfNF3eSupl `xml:"http://www.portalfiscal.inf.br/nf3e infNF3eSupl" json:"infNF3eSupl"`
	Signature   SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TNSU string

type TOperCompraGov string

type TOrigemQtd string

type TPagAnt string

type TPagRef struct {
	RefDFe []TChDFeRTC `xml:"http://www.portalfiscal.inf.br/nf3e refDFe,omitempty" json:"refDFe,omitempty"`
}

type TPagamentoRTC struct {
	NPag        string       `xml:"nPag,attr" json:"nPag"`
	IDTransacao string       `xml:"idTransacao,attr" json:"idTransacao"`
	TpMeioPgto  string       `xml:"http://www.portalfiscal.inf.br/nf3e tpMeioPgto" json:"tpMeioPgto"`
	CNPJReceb   TCNPJRTC     `xml:"http://www.portalfiscal.inf.br/nf3e CNPJReceb" json:"CNPJReceb"`
	CNPJBasePSP TCNPJBaseRTC `xml:"http://www.portalfiscal.inf.br/nf3e CNPJBasePSP" json:"CNPJBasePSP"`
}

type TPartComp string

type TPosTarSimples string

type TPosTarifCompleto string

type TProcEmi string

type TProcesso string

type TProt string

type TProtNF3eInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/nf3e tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/nf3e verAplic" json:"verAplic"`
	ChNF3e   TChDFe           `xml:"http://www.portalfiscal.inf.br/nf3e chNF3e" json:"chNF3e"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/nf3e dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/nf3e nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/nf3e digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/nf3e cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/nf3e xMotivo" json:"xMotivo"`
}

type TProtNF3eInfFisco struct {
	CMsg string  `xml:"http://www.portalfiscal.inf.br/nf3e cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/nf3e xMsg" json:"xMsg"`
}

type TProtNF3e struct {
	Versao    TVerNF3e           `xml:"versao,attr" json:"versao"`
	InfProt   TProtNF3eInfProt   `xml:"http://www.portalfiscal.inf.br/nf3e infProt" json:"infProt"`
	InfFisco  *TProtNF3eInfFisco `xml:"http://www.portalfiscal.inf.br/nf3e infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType     `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRBSN string

type TRec string

type TRed struct {
	PRedAliq  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqEfet" json:"pAliqEfet"`
}

type TRespTec struct {
	CNPJ     TCNPJ   `xml:"http://www.portalfiscal.inf.br/nf3e CNPJ" json:"CNPJ"`
	XContato string  `xml:"http://www.portalfiscal.inf.br/nf3e xContato" json:"xContato"`
	Email    TEmail  `xml:"http://www.portalfiscal.inf.br/nf3e email" json:"email"`
	Fone     string  `xml:"http://www.portalfiscal.inf.br/nf3e fone" json:"fone"`
	IDCSRT   *string `xml:"http://www.portalfiscal.inf.br/nf3e idCSRT,omitempty" json:"idCSRT,omitempty"`
	HashCSRT *string `xml:"http://www.portalfiscal.inf.br/nf3e hashCSRT,omitempty" json:"hashCSRT,omitempty"`
}

type TRetEnviNF3eInfRec struct {
	NRec     TRec         `xml:"http://www.portalfiscal.inf.br/nf3e nRec" json:"nRec"`
	DhRecbto TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nf3e dhRecbto" json:"dhRecbto"`
	TMed     string       `xml:"http://www.portalfiscal.inf.br/nf3e tMed" json:"tMed"`
}

type TRetEnviNF3e struct {
	Versao   TVerNF3e            `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb                `xml:"http://www.portalfiscal.inf.br/nf3e tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/nf3e cUF" json:"cUF"`
	VerAplic TVerAplic           `xml:"http://www.portalfiscal.inf.br/nf3e verAplic" json:"verAplic"`
	CStat    TStat               `xml:"http://www.portalfiscal.inf.br/nf3e cStat" json:"cStat"`
	XMotivo  TMotivo             `xml:"http://www.portalfiscal.inf.br/nf3e xMotivo" json:"xMotivo"`
	InfRec   *TRetEnviNF3eInfRec `xml:"http://www.portalfiscal.inf.br/nf3e infRec,omitempty" json:"infRec,omitempty"`
}

type TRetNF3e struct {
	Versao   TVerNF3e   `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb       `xml:"http://www.portalfiscal.inf.br/nf3e tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/nf3e cUF" json:"cUF"`
	VerAplic TVerAplic  `xml:"http://www.portalfiscal.inf.br/nf3e verAplic" json:"verAplic"`
	CStat    TStat      `xml:"http://www.portalfiscal.inf.br/nf3e cStat" json:"cStat"`
	XMotivo  TMotivo    `xml:"http://www.portalfiscal.inf.br/nf3e xMotivo" json:"xMotivo"`
	ProtNF3e *TProtNF3e `xml:"http://www.portalfiscal.inf.br/nf3e protNF3e,omitempty" json:"protNF3e,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TStringRTC string

type TSubClasse string

type TTarifa string

type TTime string

type TTotalSN struct {
	VRBSNTot          TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/nf3e vRBSNTot" json:"vRBSNTot"`
	VIBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSSNTot,omitempty" json:"vIBSSNTot,omitempty"`
	VIBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBSSNtotPendSusp,omitempty" json:"vIBSSNtotPendSusp,omitempty"`
	VCBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBSSNTot,omitempty" json:"vCBSSNTot,omitempty"`
	VCBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBSSNtotPendSusp,omitempty" json:"vCBSSNtotPendSusp,omitempty"`
}

type TTpCredPresIBSZFM string

type TTransfCred struct {
	VIBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vIBS" json:"vIBS"`
	VCBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/nf3e vCBS" json:"vCBS"`
}

type TTribBPe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nf3e cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nf3e indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCTe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nf3e cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nf3e indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCompraGov struct {
	PAliqIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vTribCBS" json:"vTribCBS"`
}

type TTribItemSN struct {
	NItem        string         `xml:"nItem,attr" json:"nItem"`
	VRBSNItem    TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/nf3e vRBSNItem" json:"vRBSNItem"`
	TpRBSN       TRBSN          `xml:"http://www.portalfiscal.inf.br/nf3e tpRBSN" json:"tpRBSN"`
	PIBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pIBSSN,omitempty" json:"pIBSSN,omitempty"`
	VIBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSSN,omitempty" json:"vIBSSN,omitempty"`
	PCBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pCBSSN,omitempty" json:"pCBSSN,omitempty"`
	VCBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vCBSSN,omitempty" json:"vCBSSN,omitempty"`
	VIBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vIBSPendSusp,omitempty" json:"vIBSPendSusp,omitempty"`
	VCBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vCBSPendSusp,omitempty" json:"vCBSPendSusp,omitempty"`
}

type TTribNF3e struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nf3e cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nf3e indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFAg struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nf3e cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nf3e indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFCe struct {
	CST         TCST        `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	CClassTrib  TcClassTrib `xml:"http://www.portalfiscal.inf.br/nf3e cClassTrib" json:"cClassTrib"`
	IndDoacao   *TIndDoacao `xml:"http://www.portalfiscal.inf.br/nf3e indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS     *TCIBSNFe   `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono *TMonofasia `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
}

type TTribNFCom struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nf3e cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nf3e indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFGas struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nf3e cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/nf3e indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFe struct {
	CST             TCST             `xml:"http://www.portalfiscal.inf.br/nf3e CST" json:"CST"`
	CClassTrib      TcClassTrib      `xml:"http://www.portalfiscal.inf.br/nf3e cClassTrib" json:"cClassTrib"`
	IndDoacao       *TIndDoacao      `xml:"http://www.portalfiscal.inf.br/nf3e indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS         *TCIBSNFe        `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono     *TMonofasia      `xml:"http://www.portalfiscal.inf.br/nf3e gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
	GTransfCred     *TTransfCred     `xml:"http://www.portalfiscal.inf.br/nf3e gTransfCred,omitempty" json:"gTransfCred,omitempty"`
	GAjusteCompet   *TAjusteCompet   `xml:"http://www.portalfiscal.inf.br/nf3e gAjusteCompet,omitempty" json:"gAjusteCompet,omitempty"`
	GEstornoCred    *TEstornoCred    `xml:"http://www.portalfiscal.inf.br/nf3e gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
	GCredPresOper   *TCredPresOper   `xml:"http://www.portalfiscal.inf.br/nf3e gCredPresOper,omitempty" json:"gCredPresOper,omitempty"`
	GCredPresIBSZFM *TCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/nf3e gCredPresIBSZFM,omitempty" json:"gCredPresIBSZFM,omitempty"`
}

type TTribRegular struct {
	CSTReg             TCST          `xml:"http://www.portalfiscal.inf.br/nf3e CSTReg" json:"CSTReg"`
	CClassTribReg      TcClassTrib   `xml:"http://www.portalfiscal.inf.br/nf3e cClassTribReg" json:"cClassTribReg"`
	PAliqEfetRegIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqEfetRegIBSUF" json:"pAliqEfetRegIBSUF"`
	VTribRegIBSUF      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vTribRegIBSUF" json:"vTribRegIBSUF"`
	PAliqEfetRegIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqEfetRegIBSMun" json:"pAliqEfetRegIBSMun"`
	VTribRegIBSMun     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vTribRegIBSMun" json:"vTribRegIBSMun"`
	PAliqEfetRegCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/nf3e pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS        TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/nf3e vTribRegCBS" json:"vTribRegCBS"`
}

type TUMed string

type TUMedItem string

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerNF3e string

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

type NF3eProcContent struct {
	Versao        TVerNF3e      `xml:"versao,attr" json:"versao"`
	IpTransmissor *TIPv4        `xml:"ipTransmissor,attr,omitempty" json:"ipTransmissor,omitempty"`
	NPortaCon     *string       `xml:"nPortaCon,attr,omitempty" json:"nPortaCon,omitempty"`
	DhConexao     *TDateTimeUTC `xml:"dhConexao,attr,omitempty" json:"dhConexao,omitempty"`
	NF3e          TNF3e         `xml:"http://www.portalfiscal.inf.br/nf3e NF3e" json:"NF3e"`
	ProtNF3e      TProtNF3e     `xml:"http://www.portalfiscal.inf.br/nf3e protNF3e" json:"protNF3e"`
}

// NF3eProc represents the global element {http://www.portalfiscal.inf.br/nf3e}nf3eProc.
type NF3eProc struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nf3e nf3eProc" json:"-"`
	NF3eProcContent
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
