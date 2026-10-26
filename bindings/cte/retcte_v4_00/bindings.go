// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package retcte_v4_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TALCZFMCBS struct {
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vTribRegCBS" json:"vTribRegCBS"`
}

type TALCZFMCBSNFe struct {
	TpALCZFMCBS     string        `xml:"http://www.portalfiscal.inf.br/cte tpALCZFMCBS" json:"tpALCZFMCBS"`
	NProcSuframa    *string       `xml:"http://www.portalfiscal.inf.br/cte nProcSuframa,omitempty" json:"nProcSuframa,omitempty"`
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vTribRegCBS" json:"vTribRegCBS"`
}

type TAjusteCompet struct {
	CompetApur TCompetApur `xml:"http://www.portalfiscal.inf.br/cte competApur" json:"competApur"`
	VIBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBS" json:"vIBS"`
	VCBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBS" json:"vCBS"`
}

type TAmb string

type TCIBSGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/cte gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/cte gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/cte gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSUF" json:"vIBSUF"`
}

type TCIBSGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/cte gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/cte gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/cte gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSMun" json:"vIBSMun"`
}

type TCIBSGCBS struct {
	PCBS       TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pCBS" json:"pCBS"`
	GDif       *TDif         `xml:"http://www.portalfiscal.inf.br/cte gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib     `xml:"http://www.portalfiscal.inf.br/cte gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed         `xml:"http://www.portalfiscal.inf.br/cte gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBS   `xml:"http://www.portalfiscal.inf.br/cte gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vCBS" json:"vCBS"`
}

type TCIBS struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	GIBSUF         TCIBSGIBSUF     `xml:"http://www.portalfiscal.inf.br/cte gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSGIBSMun    `xml:"http://www.portalfiscal.inf.br/cte gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/cte vIBS" json:"vIBS"`
	GCBS           TCIBSGCBS       `xml:"http://www.portalfiscal.inf.br/cte gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/cte gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/cte gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCIBSNFeGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/cte gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/cte gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/cte gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSUF" json:"vIBSUF"`
}

type TCIBSNFeGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/cte gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/cte gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/cte gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSMun" json:"vIBSMun"`
}

type TCIBSNFeGCBS struct {
	PCBS       TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/cte pCBS" json:"pCBS"`
	GDif       *TDif          `xml:"http://www.portalfiscal.inf.br/cte gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib      `xml:"http://www.portalfiscal.inf.br/cte gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed          `xml:"http://www.portalfiscal.inf.br/cte gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBSNFe `xml:"http://www.portalfiscal.inf.br/cte gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/cte vCBS" json:"vCBS"`
}

type TCIBSNFe struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	GIBSUF         TCIBSNFeGIBSUF  `xml:"http://www.portalfiscal.inf.br/cte gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSNFeGIBSMun `xml:"http://www.portalfiscal.inf.br/cte gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/cte vIBS" json:"vIBS"`
	GCBS           TCIBSNFeGCBS    `xml:"http://www.portalfiscal.inf.br/cte gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/cte gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/cte gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCIOT string

type TCListServ string

type TCOrgaoIBGE string

type TCRT string

type TCST string

type TCTeInfCTeIdeToma3 struct {
	Toma string `xml:"http://www.portalfiscal.inf.br/cte toma" json:"toma"`
}

type TCTeInfCTeIdeToma4 struct {
	Toma      string    `xml:"http://www.portalfiscal.inf.br/cte toma" json:"toma"`
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string   `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	Fone      *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	EnderToma TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderToma" json:"enderToma"`
	Email     *TEmail   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TCTeInfCTeIdeGPagAntecipado struct {
	ChDFePagAnt []TChDFe `xml:"http://www.portalfiscal.inf.br/cte chDFePagAnt,omitempty" json:"chDFePagAnt,omitempty"`
}

type TCTeInfCTeIde struct {
	CUF            TCodUFIBGE                   `xml:"http://www.portalfiscal.inf.br/cte cUF" json:"cUF"`
	CCT            string                       `xml:"http://www.portalfiscal.inf.br/cte cCT" json:"cCT"`
	CFOP           TCFOP                        `xml:"http://www.portalfiscal.inf.br/cte CFOP" json:"CFOP"`
	NatOp          string                       `xml:"http://www.portalfiscal.inf.br/cte natOp" json:"natOp"`
	Mod            TModCT                       `xml:"http://www.portalfiscal.inf.br/cte mod" json:"mod"`
	Serie          string                       `xml:"http://www.portalfiscal.inf.br/cte serie" json:"serie"`
	NCT            TNF                          `xml:"http://www.portalfiscal.inf.br/cte nCT" json:"nCT"`
	DhEmi          string                       `xml:"http://www.portalfiscal.inf.br/cte dhEmi" json:"dhEmi"`
	TpImp          string                       `xml:"http://www.portalfiscal.inf.br/cte tpImp" json:"tpImp"`
	TpEmis         string                       `xml:"http://www.portalfiscal.inf.br/cte tpEmis" json:"tpEmis"`
	CDV            string                       `xml:"http://www.portalfiscal.inf.br/cte cDV" json:"cDV"`
	TpAmb          TAmb                         `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	TpCTe          TFinCTe                      `xml:"http://www.portalfiscal.inf.br/cte tpCTe" json:"tpCTe"`
	ProcEmi        TProcEmiCTe                  `xml:"http://www.portalfiscal.inf.br/cte procEmi" json:"procEmi"`
	VerProc        string                       `xml:"http://www.portalfiscal.inf.br/cte verProc" json:"verProc"`
	IndGlobalizado *string                      `xml:"http://www.portalfiscal.inf.br/cte indGlobalizado,omitempty" json:"indGlobalizado,omitempty"`
	CMunEnv        TCodMunIBGE                  `xml:"http://www.portalfiscal.inf.br/cte cMunEnv" json:"cMunEnv"`
	XMunEnv        string                       `xml:"http://www.portalfiscal.inf.br/cte xMunEnv" json:"xMunEnv"`
	UFEnv          TUF                          `xml:"http://www.portalfiscal.inf.br/cte UFEnv" json:"UFEnv"`
	Modal          TModTransp                   `xml:"http://www.portalfiscal.inf.br/cte modal" json:"modal"`
	TpServ         string                       `xml:"http://www.portalfiscal.inf.br/cte tpServ" json:"tpServ"`
	CMunIni        TCodMunIBGE                  `xml:"http://www.portalfiscal.inf.br/cte cMunIni" json:"cMunIni"`
	XMunIni        string                       `xml:"http://www.portalfiscal.inf.br/cte xMunIni" json:"xMunIni"`
	UFIni          TUF                          `xml:"http://www.portalfiscal.inf.br/cte UFIni" json:"UFIni"`
	CMunFim        TCodMunIBGE                  `xml:"http://www.portalfiscal.inf.br/cte cMunFim" json:"cMunFim"`
	XMunFim        string                       `xml:"http://www.portalfiscal.inf.br/cte xMunFim" json:"xMunFim"`
	UFFim          TUF                          `xml:"http://www.portalfiscal.inf.br/cte UFFim" json:"UFFim"`
	Retira         string                       `xml:"http://www.portalfiscal.inf.br/cte retira" json:"retira"`
	XDetRetira     *string                      `xml:"http://www.portalfiscal.inf.br/cte xDetRetira,omitempty" json:"xDetRetira,omitempty"`
	IndIEToma      string                       `xml:"http://www.portalfiscal.inf.br/cte indIEToma" json:"indIEToma"`
	Toma3          *TCTeInfCTeIdeToma3          `xml:"http://www.portalfiscal.inf.br/cte toma3,omitempty" json:"toma3,omitempty"`
	Toma4          *TCTeInfCTeIdeToma4          `xml:"http://www.portalfiscal.inf.br/cte toma4,omitempty" json:"toma4,omitempty"`
	DhCont         *TDateTimeUTC                `xml:"http://www.portalfiscal.inf.br/cte dhCont,omitempty" json:"dhCont,omitempty"`
	XJust          *string                      `xml:"http://www.portalfiscal.inf.br/cte xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov     *TCompraGovReduzido          `xml:"http://www.portalfiscal.inf.br/cte gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	TpPagAnt       *TPagAnt                     `xml:"http://www.portalfiscal.inf.br/cte tpPagAnt,omitempty" json:"tpPagAnt,omitempty"`
	GPagAntecipado *TCTeInfCTeIdeGPagAntecipado `xml:"http://www.portalfiscal.inf.br/cte gPagAntecipado,omitempty" json:"gPagAntecipado,omitempty"`
}

type TCTeInfCTeComplFluxoPass struct {
	XPass *string `xml:"http://www.portalfiscal.inf.br/cte xPass,omitempty" json:"xPass,omitempty"`
}

type TCTeInfCTeComplFluxo struct {
	XOrig *string                    `xml:"http://www.portalfiscal.inf.br/cte xOrig,omitempty" json:"xOrig,omitempty"`
	Pass  []TCTeInfCTeComplFluxoPass `xml:"http://www.portalfiscal.inf.br/cte pass,omitempty" json:"pass,omitempty"`
	XDest *string                    `xml:"http://www.portalfiscal.inf.br/cte xDest,omitempty" json:"xDest,omitempty"`
	XRota *string                    `xml:"http://www.portalfiscal.inf.br/cte xRota,omitempty" json:"xRota,omitempty"`
}

type TCTeInfCTeComplEntregaSemData struct {
	TpPer string `xml:"http://www.portalfiscal.inf.br/cte tpPer" json:"tpPer"`
}

type TCTeInfCTeComplEntregaComData struct {
	TpPer string `xml:"http://www.portalfiscal.inf.br/cte tpPer" json:"tpPer"`
	DProg TData  `xml:"http://www.portalfiscal.inf.br/cte dProg" json:"dProg"`
}

type TCTeInfCTeComplEntregaNoPeriodo struct {
	TpPer string `xml:"http://www.portalfiscal.inf.br/cte tpPer" json:"tpPer"`
	DIni  TData  `xml:"http://www.portalfiscal.inf.br/cte dIni" json:"dIni"`
	DFim  TData  `xml:"http://www.portalfiscal.inf.br/cte dFim" json:"dFim"`
}

type TCTeInfCTeComplEntregaSemHora struct {
	TpHor string `xml:"http://www.portalfiscal.inf.br/cte tpHor" json:"tpHor"`
}

type TCTeInfCTeComplEntregaComHora struct {
	TpHor string `xml:"http://www.portalfiscal.inf.br/cte tpHor" json:"tpHor"`
	HProg TTime  `xml:"http://www.portalfiscal.inf.br/cte hProg" json:"hProg"`
}

type TCTeInfCTeComplEntregaNoInter struct {
	TpHor string `xml:"http://www.portalfiscal.inf.br/cte tpHor" json:"tpHor"`
	HIni  TTime  `xml:"http://www.portalfiscal.inf.br/cte hIni" json:"hIni"`
	HFim  TTime  `xml:"http://www.portalfiscal.inf.br/cte hFim" json:"hFim"`
}

type TCTeInfCTeComplEntrega struct {
	SemData   *TCTeInfCTeComplEntregaSemData   `xml:"http://www.portalfiscal.inf.br/cte semData,omitempty" json:"semData,omitempty"`
	ComData   *TCTeInfCTeComplEntregaComData   `xml:"http://www.portalfiscal.inf.br/cte comData,omitempty" json:"comData,omitempty"`
	NoPeriodo *TCTeInfCTeComplEntregaNoPeriodo `xml:"http://www.portalfiscal.inf.br/cte noPeriodo,omitempty" json:"noPeriodo,omitempty"`
	SemHora   *TCTeInfCTeComplEntregaSemHora   `xml:"http://www.portalfiscal.inf.br/cte semHora,omitempty" json:"semHora,omitempty"`
	ComHora   *TCTeInfCTeComplEntregaComHora   `xml:"http://www.portalfiscal.inf.br/cte comHora,omitempty" json:"comHora,omitempty"`
	NoInter   *TCTeInfCTeComplEntregaNoInter   `xml:"http://www.portalfiscal.inf.br/cte noInter,omitempty" json:"noInter,omitempty"`
}

type TCTeInfCTeComplObsCont struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/cte xTexto" json:"xTexto"`
}

type TCTeInfCTeComplObsFisco struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/cte xTexto" json:"xTexto"`
}

type TCTeInfCTeCompl struct {
	XCaracAd  *string                   `xml:"http://www.portalfiscal.inf.br/cte xCaracAd,omitempty" json:"xCaracAd,omitempty"`
	XCaracSer *string                   `xml:"http://www.portalfiscal.inf.br/cte xCaracSer,omitempty" json:"xCaracSer,omitempty"`
	XEmi      *string                   `xml:"http://www.portalfiscal.inf.br/cte xEmi,omitempty" json:"xEmi,omitempty"`
	Fluxo     *TCTeInfCTeComplFluxo     `xml:"http://www.portalfiscal.inf.br/cte fluxo,omitempty" json:"fluxo,omitempty"`
	Entrega   *TCTeInfCTeComplEntrega   `xml:"http://www.portalfiscal.inf.br/cte Entrega,omitempty" json:"Entrega,omitempty"`
	OrigCalc  *string                   `xml:"http://www.portalfiscal.inf.br/cte origCalc,omitempty" json:"origCalc,omitempty"`
	DestCalc  *string                   `xml:"http://www.portalfiscal.inf.br/cte destCalc,omitempty" json:"destCalc,omitempty"`
	XObs      *string                   `xml:"http://www.portalfiscal.inf.br/cte xObs,omitempty" json:"xObs,omitempty"`
	ObsCont   []TCTeInfCTeComplObsCont  `xml:"http://www.portalfiscal.inf.br/cte ObsCont,omitempty" json:"ObsCont,omitempty"`
	ObsFisco  []TCTeInfCTeComplObsFisco `xml:"http://www.portalfiscal.inf.br/cte ObsFisco,omitempty" json:"ObsFisco,omitempty"`
}

type TCTeInfCTeEmit struct {
	CNPJ      *TCNPJ   `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF    `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string  `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	IEST      *string  `xml:"http://www.portalfiscal.inf.br/cte IEST,omitempty" json:"IEST,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/cte enderEmit" json:"enderEmit"`
	CRT       TCRT     `xml:"http://www.portalfiscal.inf.br/cte CRT" json:"CRT"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/cte ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TCTeInfCTeRem struct {
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string   `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	Fone      *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	EnderReme TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderReme" json:"enderReme"`
	Email     *string   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TCTeInfCTeExped struct {
	CNPJ       *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF        *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE         *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome      string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	Fone       *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	EnderExped TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderExped" json:"enderExped"`
	Email      *TEmail   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TCTeInfCTeReceb struct {
	CNPJ       *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF        *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE         *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome      string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	Fone       *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	EnderReceb TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderReceb" json:"enderReceb"`
	Email      *TEmail   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TCTeInfCTeDest struct {
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	Fone      *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	ISUF      *string   `xml:"http://www.portalfiscal.inf.br/cte ISUF,omitempty" json:"ISUF,omitempty"`
	EnderDest TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderDest" json:"enderDest"`
	Email     *TEmail   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TCTeInfCTeVPrestComp struct {
	XNome string   `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	VComp TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vComp" json:"vComp"`
}

type TCTeInfCTeVPrest struct {
	VTPrest    TDec1302               `xml:"http://www.portalfiscal.inf.br/cte vTPrest" json:"vTPrest"`
	VTPrestLiq *TDec1302              `xml:"http://www.portalfiscal.inf.br/cte vTPrestLiq,omitempty" json:"vTPrestLiq,omitempty"`
	VRec       TDec1302               `xml:"http://www.portalfiscal.inf.br/cte vRec" json:"vRec"`
	Comp       []TCTeInfCTeVPrestComp `xml:"http://www.portalfiscal.inf.br/cte Comp,omitempty" json:"Comp,omitempty"`
}

type TCTeInfCTeImpICMSUFFim struct {
	VBCUFFim   TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vBCUFFim" json:"vBCUFFim"`
	PFCPUFFim  TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pFCPUFFim" json:"pFCPUFFim"`
	PICMSUFFim TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pICMSUFFim" json:"pICMSUFFim"`
	PICMSInter TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pICMSInter" json:"pICMSInter"`
	VFCPUFFim  TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vFCPUFFim" json:"vFCPUFFim"`
	VICMSUFFim TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSUFFim" json:"vICMSUFFim"`
	VICMSUFIni TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSUFIni" json:"vICMSUFIni"`
}

type TCTeInfCTeImp struct {
	ICMS       TImp                    `xml:"http://www.portalfiscal.inf.br/cte ICMS" json:"ICMS"`
	VTotTrib   *TDec1302               `xml:"http://www.portalfiscal.inf.br/cte vTotTrib,omitempty" json:"vTotTrib,omitempty"`
	InfAdFisco *string                 `xml:"http://www.portalfiscal.inf.br/cte infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	ICMSUFFim  *TCTeInfCTeImpICMSUFFim `xml:"http://www.portalfiscal.inf.br/cte ICMSUFFim,omitempty" json:"ICMSUFFim,omitempty"`
	IBSCBS     *TTribCTe               `xml:"http://www.portalfiscal.inf.br/cte IBSCBS,omitempty" json:"IBSCBS,omitempty"`
	VTotDFe    *TDec1302               `xml:"http://www.portalfiscal.inf.br/cte vTotDFe,omitempty" json:"vTotDFe,omitempty"`
}

type TCTeInfCTePgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/cte pgto,omitempty" json:"pgto,omitempty"`
}

type TCTeInfCTeInfCTeNormInfCargaInfQ struct {
	CUnid  string   `xml:"http://www.portalfiscal.inf.br/cte cUnid" json:"cUnid"`
	TpMed  string   `xml:"http://www.portalfiscal.inf.br/cte tpMed" json:"tpMed"`
	QCarga TDec1104 `xml:"http://www.portalfiscal.inf.br/cte qCarga" json:"qCarga"`
}

type TCTeInfCTeInfCTeNormInfCarga struct {
	VCarga      *TDec1302                          `xml:"http://www.portalfiscal.inf.br/cte vCarga,omitempty" json:"vCarga,omitempty"`
	ProPred     string                             `xml:"http://www.portalfiscal.inf.br/cte proPred" json:"proPred"`
	XOutCat     *string                            `xml:"http://www.portalfiscal.inf.br/cte xOutCat,omitempty" json:"xOutCat,omitempty"`
	InfQ        []TCTeInfCTeInfCTeNormInfCargaInfQ `xml:"http://www.portalfiscal.inf.br/cte infQ,omitempty" json:"infQ,omitempty"`
	VCargaAverb *TDec1302Opc                       `xml:"http://www.portalfiscal.inf.br/cte vCargaAverb,omitempty" json:"vCargaAverb,omitempty"`
}

type TCTeInfCTeInfCTeNormInfDocInfNF struct {
	NRoma         *string          `xml:"http://www.portalfiscal.inf.br/cte nRoma,omitempty" json:"nRoma,omitempty"`
	NPed          *string          `xml:"http://www.portalfiscal.inf.br/cte nPed,omitempty" json:"nPed,omitempty"`
	Mod           TModNF           `xml:"http://www.portalfiscal.inf.br/cte mod" json:"mod"`
	Serie         string           `xml:"http://www.portalfiscal.inf.br/cte serie" json:"serie"`
	NDoc          string           `xml:"http://www.portalfiscal.inf.br/cte nDoc" json:"nDoc"`
	DEmi          TData            `xml:"http://www.portalfiscal.inf.br/cte dEmi" json:"dEmi"`
	VBC           TDec1302         `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	VICMS         TDec1302         `xml:"http://www.portalfiscal.inf.br/cte vICMS" json:"vICMS"`
	VBCST         TDec1302         `xml:"http://www.portalfiscal.inf.br/cte vBCST" json:"vBCST"`
	VST           TDec1302         `xml:"http://www.portalfiscal.inf.br/cte vST" json:"vST"`
	VProd         TDec1302         `xml:"http://www.portalfiscal.inf.br/cte vProd" json:"vProd"`
	VNF           TDec1302         `xml:"http://www.portalfiscal.inf.br/cte vNF" json:"vNF"`
	NCFOP         TCFOP            `xml:"http://www.portalfiscal.inf.br/cte nCFOP" json:"nCFOP"`
	NPeso         *TDec1203Opc     `xml:"http://www.portalfiscal.inf.br/cte nPeso,omitempty" json:"nPeso,omitempty"`
	PIN           *string          `xml:"http://www.portalfiscal.inf.br/cte PIN,omitempty" json:"PIN,omitempty"`
	DPrev         *TData           `xml:"http://www.portalfiscal.inf.br/cte dPrev,omitempty" json:"dPrev,omitempty"`
	InfUnidCarga  []TUnidCarga     `xml:"http://www.portalfiscal.inf.br/cte infUnidCarga,omitempty" json:"infUnidCarga,omitempty"`
	InfUnidTransp []TUnidadeTransp `xml:"http://www.portalfiscal.inf.br/cte infUnidTransp,omitempty" json:"infUnidTransp,omitempty"`
}

type TCTeInfCTeInfCTeNormInfDocInfNFe struct {
	Chave         TChDFe           `xml:"http://www.portalfiscal.inf.br/cte chave" json:"chave"`
	PIN           *string          `xml:"http://www.portalfiscal.inf.br/cte PIN,omitempty" json:"PIN,omitempty"`
	DPrev         *TData           `xml:"http://www.portalfiscal.inf.br/cte dPrev,omitempty" json:"dPrev,omitempty"`
	InfUnidCarga  []TUnidCarga     `xml:"http://www.portalfiscal.inf.br/cte infUnidCarga,omitempty" json:"infUnidCarga,omitempty"`
	InfUnidTransp []TUnidadeTransp `xml:"http://www.portalfiscal.inf.br/cte infUnidTransp,omitempty" json:"infUnidTransp,omitempty"`
}

type TCTeInfCTeInfCTeNormInfDocInfOutros struct {
	TpDoc         string           `xml:"http://www.portalfiscal.inf.br/cte tpDoc" json:"tpDoc"`
	DescOutros    *string          `xml:"http://www.portalfiscal.inf.br/cte descOutros,omitempty" json:"descOutros,omitempty"`
	NDoc          *string          `xml:"http://www.portalfiscal.inf.br/cte nDoc,omitempty" json:"nDoc,omitempty"`
	DEmi          *TData           `xml:"http://www.portalfiscal.inf.br/cte dEmi,omitempty" json:"dEmi,omitempty"`
	VDocFisc      *TDec1302Opc     `xml:"http://www.portalfiscal.inf.br/cte vDocFisc,omitempty" json:"vDocFisc,omitempty"`
	DPrev         *TData           `xml:"http://www.portalfiscal.inf.br/cte dPrev,omitempty" json:"dPrev,omitempty"`
	InfUnidCarga  []TUnidCarga     `xml:"http://www.portalfiscal.inf.br/cte infUnidCarga,omitempty" json:"infUnidCarga,omitempty"`
	InfUnidTransp []TUnidadeTransp `xml:"http://www.portalfiscal.inf.br/cte infUnidTransp,omitempty" json:"infUnidTransp,omitempty"`
}

type TCTeInfCTeInfCTeNormInfDocInfDCe struct {
	Chave TChDFe `xml:"http://www.portalfiscal.inf.br/cte chave" json:"chave"`
}

type TCTeInfCTeInfCTeNormInfDoc struct {
	InfNF     []TCTeInfCTeInfCTeNormInfDocInfNF     `xml:"http://www.portalfiscal.inf.br/cte infNF,omitempty" json:"infNF,omitempty"`
	InfNFe    []TCTeInfCTeInfCTeNormInfDocInfNFe    `xml:"http://www.portalfiscal.inf.br/cte infNFe,omitempty" json:"infNFe,omitempty"`
	InfOutros []TCTeInfCTeInfCTeNormInfDocInfOutros `xml:"http://www.portalfiscal.inf.br/cte infOutros,omitempty" json:"infOutros,omitempty"`
	InfDCe    []TCTeInfCTeInfCTeNormInfDocInfDCe    `xml:"http://www.portalfiscal.inf.br/cte infDCe,omitempty" json:"infDCe,omitempty"`
}

type TCTeInfCTeInfCTeNormDocAntEmiDocAntIDDocAntIDDocAntPap struct {
	TpDoc  string  `xml:"http://www.portalfiscal.inf.br/cte tpDoc" json:"tpDoc"`
	Serie  string  `xml:"http://www.portalfiscal.inf.br/cte serie" json:"serie"`
	Subser *string `xml:"http://www.portalfiscal.inf.br/cte subser,omitempty" json:"subser,omitempty"`
	NDoc   string  `xml:"http://www.portalfiscal.inf.br/cte nDoc" json:"nDoc"`
	DEmi   TData   `xml:"http://www.portalfiscal.inf.br/cte dEmi" json:"dEmi"`
}

type TCTeInfCTeInfCTeNormDocAntEmiDocAntIDDocAntIDDocAntEle struct {
	ChCTe TChDFe `xml:"http://www.portalfiscal.inf.br/cte chCTe" json:"chCTe"`
}

type TCTeInfCTeInfCTeNormDocAntEmiDocAntIDDocAnt struct {
	IDDocAntPap []TCTeInfCTeInfCTeNormDocAntEmiDocAntIDDocAntIDDocAntPap `xml:"http://www.portalfiscal.inf.br/cte idDocAntPap,omitempty" json:"idDocAntPap,omitempty"`
	IDDocAntEle []TCTeInfCTeInfCTeNormDocAntEmiDocAntIDDocAntIDDocAntEle `xml:"http://www.portalfiscal.inf.br/cte idDocAntEle,omitempty" json:"idDocAntEle,omitempty"`
}

type TCTeInfCTeInfCTeNormDocAntEmiDocAnt struct {
	CNPJ     *TCNPJOpc                                     `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF      *TCPF                                         `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE       *TIe                                          `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	UF       *TUF                                          `xml:"http://www.portalfiscal.inf.br/cte UF,omitempty" json:"UF,omitempty"`
	XNome    string                                        `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	IDDocAnt []TCTeInfCTeInfCTeNormDocAntEmiDocAntIDDocAnt `xml:"http://www.portalfiscal.inf.br/cte idDocAnt,omitempty" json:"idDocAnt,omitempty"`
}

type TCTeInfCTeInfCTeNormDocAnt struct {
	EmiDocAnt []TCTeInfCTeInfCTeNormDocAntEmiDocAnt `xml:"http://www.portalfiscal.inf.br/cte emiDocAnt,omitempty" json:"emiDocAnt,omitempty"`
}

type TCTeInfCTeInfCTeNormInfModal struct {
	VersaoModal string       `xml:"versaoModal,attr" json:"versaoModal"`
	Any         []RawElement `xml:",any" json:"any,omitempty"`
}

type TCTeInfCTeInfCTeNormVeicNovos struct {
	Chassi string   `xml:"http://www.portalfiscal.inf.br/cte chassi" json:"chassi"`
	CCor   string   `xml:"http://www.portalfiscal.inf.br/cte cCor" json:"cCor"`
	XCor   string   `xml:"http://www.portalfiscal.inf.br/cte xCor" json:"xCor"`
	CMod   string   `xml:"http://www.portalfiscal.inf.br/cte cMod" json:"cMod"`
	VUnit  TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vUnit" json:"vUnit"`
	VFrete TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vFrete" json:"vFrete"`
}

type TCTeInfCTeInfCTeNormCobrFat struct {
	NFat  *string      `xml:"http://www.portalfiscal.inf.br/cte nFat,omitempty" json:"nFat,omitempty"`
	VOrig *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vOrig,omitempty" json:"vOrig,omitempty"`
	VDesc *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vDesc,omitempty" json:"vDesc,omitempty"`
	VLiq  *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vLiq,omitempty" json:"vLiq,omitempty"`
}

type TCTeInfCTeInfCTeNormCobrDup struct {
	NDup  *string      `xml:"http://www.portalfiscal.inf.br/cte nDup,omitempty" json:"nDup,omitempty"`
	DVenc *TData       `xml:"http://www.portalfiscal.inf.br/cte dVenc,omitempty" json:"dVenc,omitempty"`
	VDup  *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vDup,omitempty" json:"vDup,omitempty"`
}

type TCTeInfCTeInfCTeNormCobr struct {
	Fat *TCTeInfCTeInfCTeNormCobrFat  `xml:"http://www.portalfiscal.inf.br/cte fat,omitempty" json:"fat,omitempty"`
	Dup []TCTeInfCTeInfCTeNormCobrDup `xml:"http://www.portalfiscal.inf.br/cte dup,omitempty" json:"dup,omitempty"`
}

type TCTeInfCTeInfCTeNormInfCTeSub struct {
	ChCTe         TChDFe  `xml:"http://www.portalfiscal.inf.br/cte chCte" json:"chCte"`
	IndAlteraToma *string `xml:"http://www.portalfiscal.inf.br/cte indAlteraToma,omitempty" json:"indAlteraToma,omitempty"`
}

type TCTeInfCTeInfCTeNormInfGlobalizado struct {
	XObs string `xml:"http://www.portalfiscal.inf.br/cte xObs" json:"xObs"`
}

type TCTeInfCTeInfCTeNormInfServVincInfCTeMultimodal struct {
	ChCTeMultimodal TChDFe `xml:"http://www.portalfiscal.inf.br/cte chCTeMultimodal" json:"chCTeMultimodal"`
}

type TCTeInfCTeInfCTeNormInfServVinc struct {
	InfCTeMultimodal []TCTeInfCTeInfCTeNormInfServVincInfCTeMultimodal `xml:"http://www.portalfiscal.inf.br/cte infCTeMultimodal,omitempty" json:"infCTeMultimodal,omitempty"`
}

type TCTeInfCTeInfCTeNorm struct {
	InfCarga       TCTeInfCTeInfCTeNormInfCarga        `xml:"http://www.portalfiscal.inf.br/cte infCarga" json:"infCarga"`
	InfDoc         *TCTeInfCTeInfCTeNormInfDoc         `xml:"http://www.portalfiscal.inf.br/cte infDoc,omitempty" json:"infDoc,omitempty"`
	DocAnt         *TCTeInfCTeInfCTeNormDocAnt         `xml:"http://www.portalfiscal.inf.br/cte docAnt,omitempty" json:"docAnt,omitempty"`
	InfModal       TCTeInfCTeInfCTeNormInfModal        `xml:"http://www.portalfiscal.inf.br/cte infModal" json:"infModal"`
	VeicNovos      []TCTeInfCTeInfCTeNormVeicNovos     `xml:"http://www.portalfiscal.inf.br/cte veicNovos,omitempty" json:"veicNovos,omitempty"`
	Cobr           *TCTeInfCTeInfCTeNormCobr           `xml:"http://www.portalfiscal.inf.br/cte cobr,omitempty" json:"cobr,omitempty"`
	InfCTeSub      *TCTeInfCTeInfCTeNormInfCTeSub      `xml:"http://www.portalfiscal.inf.br/cte infCteSub,omitempty" json:"infCteSub,omitempty"`
	InfGlobalizado *TCTeInfCTeInfCTeNormInfGlobalizado `xml:"http://www.portalfiscal.inf.br/cte infGlobalizado,omitempty" json:"infGlobalizado,omitempty"`
	InfServVinc    *TCTeInfCTeInfCTeNormInfServVinc    `xml:"http://www.portalfiscal.inf.br/cte infServVinc,omitempty" json:"infServVinc,omitempty"`
}

type TCTeInfCTeInfCTeComp struct {
	ChCTe TChDFe `xml:"http://www.portalfiscal.inf.br/cte chCTe" json:"chCTe"`
}

type TCTeInfCTeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
}

type TCTeInfCTeInfSolicNFF struct {
	XSolic string `xml:"http://www.portalfiscal.inf.br/cte xSolic" json:"xSolic"`
}

type TCTeInfCTeInfPAAPAASignature struct {
	SignatureValue Base64Binary     `xml:"http://www.portalfiscal.inf.br/cte SignatureValue" json:"SignatureValue"`
	RSAKeyValue    TRSAKeyValueType `xml:"http://www.portalfiscal.inf.br/cte RSAKeyValue" json:"RSAKeyValue"`
}

type TCTeInfCTeInfPAA struct {
	CNPJPAA      TCNPJ                        `xml:"http://www.portalfiscal.inf.br/cte CNPJPAA" json:"CNPJPAA"`
	PAASignature TCTeInfCTeInfPAAPAASignature `xml:"http://www.portalfiscal.inf.br/cte PAASignature" json:"PAASignature"`
}

type TCTeInfCTe struct {
	Versao      string                 `xml:"versao,attr" json:"versao"`
	ID          string                 `xml:"Id,attr" json:"Id"`
	Ide         TCTeInfCTeIde          `xml:"http://www.portalfiscal.inf.br/cte ide" json:"ide"`
	Compl       *TCTeInfCTeCompl       `xml:"http://www.portalfiscal.inf.br/cte compl,omitempty" json:"compl,omitempty"`
	Emit        TCTeInfCTeEmit         `xml:"http://www.portalfiscal.inf.br/cte emit" json:"emit"`
	Rem         *TCTeInfCTeRem         `xml:"http://www.portalfiscal.inf.br/cte rem,omitempty" json:"rem,omitempty"`
	Exped       *TCTeInfCTeExped       `xml:"http://www.portalfiscal.inf.br/cte exped,omitempty" json:"exped,omitempty"`
	Receb       *TCTeInfCTeReceb       `xml:"http://www.portalfiscal.inf.br/cte receb,omitempty" json:"receb,omitempty"`
	Dest        *TCTeInfCTeDest        `xml:"http://www.portalfiscal.inf.br/cte dest,omitempty" json:"dest,omitempty"`
	VPrest      TCTeInfCTeVPrest       `xml:"http://www.portalfiscal.inf.br/cte vPrest" json:"vPrest"`
	Imp         TCTeInfCTeImp          `xml:"http://www.portalfiscal.inf.br/cte imp" json:"imp"`
	PgtoVinc    *TCTeInfCTePgtoVinc    `xml:"http://www.portalfiscal.inf.br/cte pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	InfCTeNorm  *TCTeInfCTeInfCTeNorm  `xml:"http://www.portalfiscal.inf.br/cte infCTeNorm,omitempty" json:"infCTeNorm,omitempty"`
	InfCTeComp  []TCTeInfCTeInfCTeComp `xml:"http://www.portalfiscal.inf.br/cte infCteComp,omitempty" json:"infCteComp,omitempty"`
	AutXML      []TCTeInfCTeAutXML     `xml:"http://www.portalfiscal.inf.br/cte autXML,omitempty" json:"autXML,omitempty"`
	InfRespTec  *TRespTec              `xml:"http://www.portalfiscal.inf.br/cte infRespTec,omitempty" json:"infRespTec,omitempty"`
	InfSolicNFF *TCTeInfCTeInfSolicNFF `xml:"http://www.portalfiscal.inf.br/cte infSolicNFF,omitempty" json:"infSolicNFF,omitempty"`
	InfPAA      *TCTeInfCTeInfPAA      `xml:"http://www.portalfiscal.inf.br/cte infPAA,omitempty" json:"infPAA,omitempty"`
}

type TCTeInfCTeSupl struct {
	QRCodCTe string `xml:"http://www.portalfiscal.inf.br/cte qrCodCTe" json:"qrCodCTe"`
}

type TCTe struct {
	InfCTe     TCTeInfCTe      `xml:"http://www.portalfiscal.inf.br/cte infCte" json:"infCte"`
	InfCTeSupl *TCTeInfCTeSupl `xml:"http://www.portalfiscal.inf.br/cte infCTeSupl,omitempty" json:"infCTeSupl,omitempty"`
	Signature  SignatureType   `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCTeOSInfCTeIdeInfPercurso struct {
	UFPer TUF `xml:"http://www.portalfiscal.inf.br/cte UFPer" json:"UFPer"`
}

type TCTeOSInfCTeIdeGPagAntecipado struct {
	ChDFePagAnt []TChDFe `xml:"http://www.portalfiscal.inf.br/cte chDFePagAnt,omitempty" json:"chDFePagAnt,omitempty"`
}

type TCTeOSInfCTeIde struct {
	CUF            TCodUFIBGE                     `xml:"http://www.portalfiscal.inf.br/cte cUF" json:"cUF"`
	CCT            string                         `xml:"http://www.portalfiscal.inf.br/cte cCT" json:"cCT"`
	CFOP           TCFOP                          `xml:"http://www.portalfiscal.inf.br/cte CFOP" json:"CFOP"`
	NatOp          string                         `xml:"http://www.portalfiscal.inf.br/cte natOp" json:"natOp"`
	Mod            TModCTOS                       `xml:"http://www.portalfiscal.inf.br/cte mod" json:"mod"`
	Serie          string                         `xml:"http://www.portalfiscal.inf.br/cte serie" json:"serie"`
	NCT            TNF                            `xml:"http://www.portalfiscal.inf.br/cte nCT" json:"nCT"`
	DhEmi          string                         `xml:"http://www.portalfiscal.inf.br/cte dhEmi" json:"dhEmi"`
	TpImp          string                         `xml:"http://www.portalfiscal.inf.br/cte tpImp" json:"tpImp"`
	TpEmis         string                         `xml:"http://www.portalfiscal.inf.br/cte tpEmis" json:"tpEmis"`
	CDV            string                         `xml:"http://www.portalfiscal.inf.br/cte cDV" json:"cDV"`
	TpAmb          TAmb                           `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	TpCTe          TFinCTe                        `xml:"http://www.portalfiscal.inf.br/cte tpCTe" json:"tpCTe"`
	ProcEmi        TProcEmi                       `xml:"http://www.portalfiscal.inf.br/cte procEmi" json:"procEmi"`
	VerProc        string                         `xml:"http://www.portalfiscal.inf.br/cte verProc" json:"verProc"`
	CMunEnv        TCodMunIBGE                    `xml:"http://www.portalfiscal.inf.br/cte cMunEnv" json:"cMunEnv"`
	XMunEnv        string                         `xml:"http://www.portalfiscal.inf.br/cte xMunEnv" json:"xMunEnv"`
	UFEnv          TUF                            `xml:"http://www.portalfiscal.inf.br/cte UFEnv" json:"UFEnv"`
	Modal          TModTranspOS                   `xml:"http://www.portalfiscal.inf.br/cte modal" json:"modal"`
	TpServ         string                         `xml:"http://www.portalfiscal.inf.br/cte tpServ" json:"tpServ"`
	IndIEToma      string                         `xml:"http://www.portalfiscal.inf.br/cte indIEToma" json:"indIEToma"`
	CMunIni        *TCodMunIBGE                   `xml:"http://www.portalfiscal.inf.br/cte cMunIni,omitempty" json:"cMunIni,omitempty"`
	XMunIni        *string                        `xml:"http://www.portalfiscal.inf.br/cte xMunIni,omitempty" json:"xMunIni,omitempty"`
	UFIni          *TUF                           `xml:"http://www.portalfiscal.inf.br/cte UFIni,omitempty" json:"UFIni,omitempty"`
	CMunFim        *TCodMunIBGE                   `xml:"http://www.portalfiscal.inf.br/cte cMunFim,omitempty" json:"cMunFim,omitempty"`
	XMunFim        *string                        `xml:"http://www.portalfiscal.inf.br/cte xMunFim,omitempty" json:"xMunFim,omitempty"`
	UFFim          *TUF                           `xml:"http://www.portalfiscal.inf.br/cte UFFim,omitempty" json:"UFFim,omitempty"`
	InfPercurso    []TCTeOSInfCTeIdeInfPercurso   `xml:"http://www.portalfiscal.inf.br/cte infPercurso,omitempty" json:"infPercurso,omitempty"`
	DhCont         *TDateTimeUTC                  `xml:"http://www.portalfiscal.inf.br/cte dhCont,omitempty" json:"dhCont,omitempty"`
	XJust          *string                        `xml:"http://www.portalfiscal.inf.br/cte xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov     *TCompraGovReduzido            `xml:"http://www.portalfiscal.inf.br/cte gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	TpPagAnt       *TPagAnt                       `xml:"http://www.portalfiscal.inf.br/cte tpPagAnt,omitempty" json:"tpPagAnt,omitempty"`
	GPagAntecipado *TCTeOSInfCTeIdeGPagAntecipado `xml:"http://www.portalfiscal.inf.br/cte gPagAntecipado,omitempty" json:"gPagAntecipado,omitempty"`
}

type TCTeOSInfCTeComplObsCont struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/cte xTexto" json:"xTexto"`
}

type TCTeOSInfCTeComplObsFisco struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/cte xTexto" json:"xTexto"`
}

type TCTeOSInfCTeCompl struct {
	XCaracAd  *string                     `xml:"http://www.portalfiscal.inf.br/cte xCaracAd,omitempty" json:"xCaracAd,omitempty"`
	XCaracSer *string                     `xml:"http://www.portalfiscal.inf.br/cte xCaracSer,omitempty" json:"xCaracSer,omitempty"`
	XEmi      *string                     `xml:"http://www.portalfiscal.inf.br/cte xEmi,omitempty" json:"xEmi,omitempty"`
	XObs      *string                     `xml:"http://www.portalfiscal.inf.br/cte xObs,omitempty" json:"xObs,omitempty"`
	ObsCont   []TCTeOSInfCTeComplObsCont  `xml:"http://www.portalfiscal.inf.br/cte ObsCont,omitempty" json:"ObsCont,omitempty"`
	ObsFisco  []TCTeOSInfCTeComplObsFisco `xml:"http://www.portalfiscal.inf.br/cte ObsFisco,omitempty" json:"ObsFisco,omitempty"`
}

type TCTeOSInfCTeEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/cte CNPJ" json:"CNPJ"`
	IE        string   `xml:"http://www.portalfiscal.inf.br/cte IE" json:"IE"`
	IEST      *string  `xml:"http://www.portalfiscal.inf.br/cte IEST,omitempty" json:"IEST,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/cte enderEmit" json:"enderEmit"`
	CRT       TCRT     `xml:"http://www.portalfiscal.inf.br/cte CRT" json:"CRT"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/cte ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TCTeOSInfCTeToma struct {
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string   `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	Fone      *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	EnderToma TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderToma" json:"enderToma"`
	Email     *string   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TCTeOSInfCTeVPrestComp struct {
	XNome string   `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	VComp TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vComp" json:"vComp"`
}

type TCTeOSInfCTeVPrest struct {
	VTPrest    TDec1302                 `xml:"http://www.portalfiscal.inf.br/cte vTPrest" json:"vTPrest"`
	VTPrestLiq *TDec1302                `xml:"http://www.portalfiscal.inf.br/cte vTPrestLiq,omitempty" json:"vTPrestLiq,omitempty"`
	VRec       TDec1302                 `xml:"http://www.portalfiscal.inf.br/cte vRec" json:"vRec"`
	Comp       []TCTeOSInfCTeVPrestComp `xml:"http://www.portalfiscal.inf.br/cte Comp,omitempty" json:"Comp,omitempty"`
}

type TCTeOSInfCTeImpICMSUFFim struct {
	VBCUFFim   TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vBCUFFim" json:"vBCUFFim"`
	PFCPUFFim  TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pFCPUFFim" json:"pFCPUFFim"`
	PICMSUFFim TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pICMSUFFim" json:"pICMSUFFim"`
	PICMSInter TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pICMSInter" json:"pICMSInter"`
	VFCPUFFim  TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vFCPUFFim" json:"vFCPUFFim"`
	VICMSUFFim TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSUFFim" json:"vICMSUFFim"`
	VICMSUFIni TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSUFIni" json:"vICMSUFIni"`
}

type TCTeOSInfCTeImpInfTribFed struct {
	VPIS    *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vPIS,omitempty" json:"vPIS,omitempty"`
	VCOFINS *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vCOFINS,omitempty" json:"vCOFINS,omitempty"`
	VIR     *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vIR,omitempty" json:"vIR,omitempty"`
	VINSS   *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vINSS,omitempty" json:"vINSS,omitempty"`
	VCSLL   *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vCSLL,omitempty" json:"vCSLL,omitempty"`
}

type TCTeOSInfCTeImp struct {
	ICMS       TImpOS                     `xml:"http://www.portalfiscal.inf.br/cte ICMS" json:"ICMS"`
	VTotTrib   *TDec1302                  `xml:"http://www.portalfiscal.inf.br/cte vTotTrib,omitempty" json:"vTotTrib,omitempty"`
	InfAdFisco *string                    `xml:"http://www.portalfiscal.inf.br/cte infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	ICMSUFFim  *TCTeOSInfCTeImpICMSUFFim  `xml:"http://www.portalfiscal.inf.br/cte ICMSUFFim,omitempty" json:"ICMSUFFim,omitempty"`
	InfTribFed *TCTeOSInfCTeImpInfTribFed `xml:"http://www.portalfiscal.inf.br/cte infTribFed,omitempty" json:"infTribFed,omitempty"`
	IBSCBS     *TTribCTe                  `xml:"http://www.portalfiscal.inf.br/cte IBSCBS,omitempty" json:"IBSCBS,omitempty"`
	VTotDFe    *TDec1302                  `xml:"http://www.portalfiscal.inf.br/cte vTotDFe,omitempty" json:"vTotDFe,omitempty"`
}

type TCTeOSInfCTePgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/cte pgto,omitempty" json:"pgto,omitempty"`
}

type TCTeOSInfCTeInfCTeNormInfServicoInfQ struct {
	QCarga TDec1104 `xml:"http://www.portalfiscal.inf.br/cte qCarga" json:"qCarga"`
}

type TCTeOSInfCTeInfCTeNormInfServico struct {
	XDescServ string                                `xml:"http://www.portalfiscal.inf.br/cte xDescServ" json:"xDescServ"`
	InfQ      *TCTeOSInfCTeInfCTeNormInfServicoInfQ `xml:"http://www.portalfiscal.inf.br/cte infQ,omitempty" json:"infQ,omitempty"`
}

type TCTeOSInfCTeInfCTeNormInfDocRef struct {
	NDoc     *string   `xml:"http://www.portalfiscal.inf.br/cte nDoc,omitempty" json:"nDoc,omitempty"`
	Serie    *string   `xml:"http://www.portalfiscal.inf.br/cte serie,omitempty" json:"serie,omitempty"`
	Subserie *string   `xml:"http://www.portalfiscal.inf.br/cte subserie,omitempty" json:"subserie,omitempty"`
	DEmi     *TData    `xml:"http://www.portalfiscal.inf.br/cte dEmi,omitempty" json:"dEmi,omitempty"`
	VDoc     *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vDoc,omitempty" json:"vDoc,omitempty"`
	ChBPe    *string   `xml:"http://www.portalfiscal.inf.br/cte chBPe,omitempty" json:"chBPe,omitempty"`
}

type TCTeOSInfCTeInfCTeNormSeg struct {
	RespSeg string  `xml:"http://www.portalfiscal.inf.br/cte respSeg" json:"respSeg"`
	XSeg    *string `xml:"http://www.portalfiscal.inf.br/cte xSeg,omitempty" json:"xSeg,omitempty"`
	NApol   *string `xml:"http://www.portalfiscal.inf.br/cte nApol,omitempty" json:"nApol,omitempty"`
}

type TCTeOSInfCTeInfCTeNormInfModal struct {
	VersaoModal string       `xml:"versaoModal,attr" json:"versaoModal"`
	Any         []RawElement `xml:",any" json:"any,omitempty"`
}

type TCTeOSInfCTeInfCTeNormInfCTeSub struct {
	ChCTe TChDFe `xml:"http://www.portalfiscal.inf.br/cte chCte" json:"chCte"`
}

type TCTeOSInfCTeInfCTeNormCobrFat struct {
	NFat  *string      `xml:"http://www.portalfiscal.inf.br/cte nFat,omitempty" json:"nFat,omitempty"`
	VOrig *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vOrig,omitempty" json:"vOrig,omitempty"`
	VDesc *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vDesc,omitempty" json:"vDesc,omitempty"`
	VLiq  *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vLiq,omitempty" json:"vLiq,omitempty"`
}

type TCTeOSInfCTeInfCTeNormCobrDup struct {
	NDup  *string      `xml:"http://www.portalfiscal.inf.br/cte nDup,omitempty" json:"nDup,omitempty"`
	DVenc *TData       `xml:"http://www.portalfiscal.inf.br/cte dVenc,omitempty" json:"dVenc,omitempty"`
	VDup  *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vDup,omitempty" json:"vDup,omitempty"`
}

type TCTeOSInfCTeInfCTeNormCobr struct {
	Fat *TCTeOSInfCTeInfCTeNormCobrFat  `xml:"http://www.portalfiscal.inf.br/cte fat,omitempty" json:"fat,omitempty"`
	Dup []TCTeOSInfCTeInfCTeNormCobrDup `xml:"http://www.portalfiscal.inf.br/cte dup,omitempty" json:"dup,omitempty"`
}

type TCTeOSInfCTeInfCTeNormInfGTVeComp struct {
	TpComp string   `xml:"http://www.portalfiscal.inf.br/cte tpComp" json:"tpComp"`
	VComp  TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vComp" json:"vComp"`
	XComp  *string  `xml:"http://www.portalfiscal.inf.br/cte xComp,omitempty" json:"xComp,omitempty"`
}

type TCTeOSInfCTeInfCTeNormInfGTVe struct {
	ChCTe TChDFe                              `xml:"http://www.portalfiscal.inf.br/cte chCTe" json:"chCTe"`
	Comp  []TCTeOSInfCTeInfCTeNormInfGTVeComp `xml:"http://www.portalfiscal.inf.br/cte Comp,omitempty" json:"Comp,omitempty"`
}

type TCTeOSInfCTeInfCTeNorm struct {
	InfServico TCTeOSInfCTeInfCTeNormInfServico  `xml:"http://www.portalfiscal.inf.br/cte infServico" json:"infServico"`
	InfDocRef  []TCTeOSInfCTeInfCTeNormInfDocRef `xml:"http://www.portalfiscal.inf.br/cte infDocRef,omitempty" json:"infDocRef,omitempty"`
	Seg        []TCTeOSInfCTeInfCTeNormSeg       `xml:"http://www.portalfiscal.inf.br/cte seg,omitempty" json:"seg,omitempty"`
	InfModal   *TCTeOSInfCTeInfCTeNormInfModal   `xml:"http://www.portalfiscal.inf.br/cte infModal,omitempty" json:"infModal,omitempty"`
	InfCTeSub  *TCTeOSInfCTeInfCTeNormInfCTeSub  `xml:"http://www.portalfiscal.inf.br/cte infCteSub,omitempty" json:"infCteSub,omitempty"`
	RefCTeCanc *string                           `xml:"http://www.portalfiscal.inf.br/cte refCTeCanc,omitempty" json:"refCTeCanc,omitempty"`
	Cobr       *TCTeOSInfCTeInfCTeNormCobr       `xml:"http://www.portalfiscal.inf.br/cte cobr,omitempty" json:"cobr,omitempty"`
	InfGTVe    []TCTeOSInfCTeInfCTeNormInfGTVe   `xml:"http://www.portalfiscal.inf.br/cte infGTVe,omitempty" json:"infGTVe,omitempty"`
}

type TCTeOSInfCTeInfCTeComp struct {
	ChCTe TChDFe `xml:"http://www.portalfiscal.inf.br/cte chCTe" json:"chCTe"`
}

type TCTeOSInfCTeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
}

type TCTeOSInfCTe struct {
	Versao     string                   `xml:"versao,attr" json:"versao"`
	ID         string                   `xml:"Id,attr" json:"Id"`
	Ide        TCTeOSInfCTeIde          `xml:"http://www.portalfiscal.inf.br/cte ide" json:"ide"`
	Compl      *TCTeOSInfCTeCompl       `xml:"http://www.portalfiscal.inf.br/cte compl,omitempty" json:"compl,omitempty"`
	Emit       TCTeOSInfCTeEmit         `xml:"http://www.portalfiscal.inf.br/cte emit" json:"emit"`
	Toma       *TCTeOSInfCTeToma        `xml:"http://www.portalfiscal.inf.br/cte toma,omitempty" json:"toma,omitempty"`
	VPrest     TCTeOSInfCTeVPrest       `xml:"http://www.portalfiscal.inf.br/cte vPrest" json:"vPrest"`
	Imp        TCTeOSInfCTeImp          `xml:"http://www.portalfiscal.inf.br/cte imp" json:"imp"`
	PgtoVinc   *TCTeOSInfCTePgtoVinc    `xml:"http://www.portalfiscal.inf.br/cte pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	InfCTeNorm *TCTeOSInfCTeInfCTeNorm  `xml:"http://www.portalfiscal.inf.br/cte infCTeNorm,omitempty" json:"infCTeNorm,omitempty"`
	InfCTeComp []TCTeOSInfCTeInfCTeComp `xml:"http://www.portalfiscal.inf.br/cte infCteComp,omitempty" json:"infCteComp,omitempty"`
	AutXML     []TCTeOSInfCTeAutXML     `xml:"http://www.portalfiscal.inf.br/cte autXML,omitempty" json:"autXML,omitempty"`
	InfRespTec *TRespTec                `xml:"http://www.portalfiscal.inf.br/cte infRespTec,omitempty" json:"infRespTec,omitempty"`
}

type TCTeOSInfCTeSupl struct {
	QRCodCTe string `xml:"http://www.portalfiscal.inf.br/cte qrCodCTe" json:"qrCodCTe"`
}

type TCTeOS struct {
	Versao     string            `xml:"versao,attr" json:"versao"`
	InfCTe     TCTeOSInfCTe      `xml:"http://www.portalfiscal.inf.br/cte infCte" json:"infCte"`
	InfCTeSupl *TCTeOSInfCTeSupl `xml:"http://www.portalfiscal.inf.br/cte infCTeSupl,omitempty" json:"infCTeSupl,omitempty"`
	Signature  SignatureType     `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCTeSimpInfCTeIdeGPagAntecipado struct {
	ChDFePagAnt []TChDFe `xml:"http://www.portalfiscal.inf.br/cte chDFePagAnt,omitempty" json:"chDFePagAnt,omitempty"`
}

type TCTeSimpInfCTeIde struct {
	CUF            TCodUFIBGE                       `xml:"http://www.portalfiscal.inf.br/cte cUF" json:"cUF"`
	CCT            string                           `xml:"http://www.portalfiscal.inf.br/cte cCT" json:"cCT"`
	CFOP           TCFOP                            `xml:"http://www.portalfiscal.inf.br/cte CFOP" json:"CFOP"`
	NatOp          string                           `xml:"http://www.portalfiscal.inf.br/cte natOp" json:"natOp"`
	Mod            TModCT                           `xml:"http://www.portalfiscal.inf.br/cte mod" json:"mod"`
	Serie          string                           `xml:"http://www.portalfiscal.inf.br/cte serie" json:"serie"`
	NCT            TNF                              `xml:"http://www.portalfiscal.inf.br/cte nCT" json:"nCT"`
	DhEmi          string                           `xml:"http://www.portalfiscal.inf.br/cte dhEmi" json:"dhEmi"`
	TpImp          string                           `xml:"http://www.portalfiscal.inf.br/cte tpImp" json:"tpImp"`
	TpEmis         string                           `xml:"http://www.portalfiscal.inf.br/cte tpEmis" json:"tpEmis"`
	CDV            string                           `xml:"http://www.portalfiscal.inf.br/cte cDV" json:"cDV"`
	TpAmb          TAmb                             `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	TpCTe          TFinCTeSimp                      `xml:"http://www.portalfiscal.inf.br/cte tpCTe" json:"tpCTe"`
	ProcEmi        TProcEmiCTe                      `xml:"http://www.portalfiscal.inf.br/cte procEmi" json:"procEmi"`
	VerProc        string                           `xml:"http://www.portalfiscal.inf.br/cte verProc" json:"verProc"`
	CMunEnv        TCodMunIBGE                      `xml:"http://www.portalfiscal.inf.br/cte cMunEnv" json:"cMunEnv"`
	XMunEnv        string                           `xml:"http://www.portalfiscal.inf.br/cte xMunEnv" json:"xMunEnv"`
	UFEnv          TUF                              `xml:"http://www.portalfiscal.inf.br/cte UFEnv" json:"UFEnv"`
	Modal          TModTranspSimp                   `xml:"http://www.portalfiscal.inf.br/cte modal" json:"modal"`
	TpServ         string                           `xml:"http://www.portalfiscal.inf.br/cte tpServ" json:"tpServ"`
	UFIni          TUF                              `xml:"http://www.portalfiscal.inf.br/cte UFIni" json:"UFIni"`
	UFFim          TUF                              `xml:"http://www.portalfiscal.inf.br/cte UFFim" json:"UFFim"`
	Retira         string                           `xml:"http://www.portalfiscal.inf.br/cte retira" json:"retira"`
	XDetRetira     *string                          `xml:"http://www.portalfiscal.inf.br/cte xDetRetira,omitempty" json:"xDetRetira,omitempty"`
	DhCont         *TDateTimeUTC                    `xml:"http://www.portalfiscal.inf.br/cte dhCont,omitempty" json:"dhCont,omitempty"`
	XJust          *string                          `xml:"http://www.portalfiscal.inf.br/cte xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov     *TCompraGovReduzido              `xml:"http://www.portalfiscal.inf.br/cte gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	TpPagAnt       *TPagAnt                         `xml:"http://www.portalfiscal.inf.br/cte tpPagAnt,omitempty" json:"tpPagAnt,omitempty"`
	GPagAntecipado *TCTeSimpInfCTeIdeGPagAntecipado `xml:"http://www.portalfiscal.inf.br/cte gPagAntecipado,omitempty" json:"gPagAntecipado,omitempty"`
}

type TCTeSimpInfCTeComplFluxoPass struct {
	XPass *string `xml:"http://www.portalfiscal.inf.br/cte xPass,omitempty" json:"xPass,omitempty"`
}

type TCTeSimpInfCTeComplFluxo struct {
	XOrig *string                        `xml:"http://www.portalfiscal.inf.br/cte xOrig,omitempty" json:"xOrig,omitempty"`
	Pass  []TCTeSimpInfCTeComplFluxoPass `xml:"http://www.portalfiscal.inf.br/cte pass,omitempty" json:"pass,omitempty"`
	XDest *string                        `xml:"http://www.portalfiscal.inf.br/cte xDest,omitempty" json:"xDest,omitempty"`
	XRota *string                        `xml:"http://www.portalfiscal.inf.br/cte xRota,omitempty" json:"xRota,omitempty"`
}

type TCTeSimpInfCTeComplObsCont struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/cte xTexto" json:"xTexto"`
}

type TCTeSimpInfCTeComplObsFisco struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/cte xTexto" json:"xTexto"`
}

type TCTeSimpInfCTeCompl struct {
	XCaracAd  *string                       `xml:"http://www.portalfiscal.inf.br/cte xCaracAd,omitempty" json:"xCaracAd,omitempty"`
	XCaracSer *string                       `xml:"http://www.portalfiscal.inf.br/cte xCaracSer,omitempty" json:"xCaracSer,omitempty"`
	Fluxo     *TCTeSimpInfCTeComplFluxo     `xml:"http://www.portalfiscal.inf.br/cte fluxo,omitempty" json:"fluxo,omitempty"`
	XObs      *string                       `xml:"http://www.portalfiscal.inf.br/cte xObs,omitempty" json:"xObs,omitempty"`
	ObsCont   []TCTeSimpInfCTeComplObsCont  `xml:"http://www.portalfiscal.inf.br/cte ObsCont,omitempty" json:"ObsCont,omitempty"`
	ObsFisco  []TCTeSimpInfCTeComplObsFisco `xml:"http://www.portalfiscal.inf.br/cte ObsFisco,omitempty" json:"ObsFisco,omitempty"`
}

type TCTeSimpInfCTeEmit struct {
	CNPJ      *TCNPJ   `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF    `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string  `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	IEST      *string  `xml:"http://www.portalfiscal.inf.br/cte IEST,omitempty" json:"IEST,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/cte enderEmit" json:"enderEmit"`
	CRT       TCRT     `xml:"http://www.portalfiscal.inf.br/cte CRT" json:"CRT"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/cte ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TCTeSimpInfCTeToma struct {
	Toma      string    `xml:"http://www.portalfiscal.inf.br/cte toma" json:"toma"`
	IndIEToma string    `xml:"http://www.portalfiscal.inf.br/cte indIEToma" json:"indIEToma"`
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	ISUF      *string   `xml:"http://www.portalfiscal.inf.br/cte ISUF,omitempty" json:"ISUF,omitempty"`
	Fone      *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	EnderToma TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderToma" json:"enderToma"`
	Email     *TEmail   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TCTeSimpInfCTeInfCargaInfQ struct {
	CUnid  string   `xml:"http://www.portalfiscal.inf.br/cte cUnid" json:"cUnid"`
	TpMed  string   `xml:"http://www.portalfiscal.inf.br/cte tpMed" json:"tpMed"`
	QCarga TDec1104 `xml:"http://www.portalfiscal.inf.br/cte qCarga" json:"qCarga"`
}

type TCTeSimpInfCTeInfCarga struct {
	VCarga      TDec1302                     `xml:"http://www.portalfiscal.inf.br/cte vCarga" json:"vCarga"`
	ProPred     string                       `xml:"http://www.portalfiscal.inf.br/cte proPred" json:"proPred"`
	XOutCat     *string                      `xml:"http://www.portalfiscal.inf.br/cte xOutCat,omitempty" json:"xOutCat,omitempty"`
	InfQ        []TCTeSimpInfCTeInfCargaInfQ `xml:"http://www.portalfiscal.inf.br/cte infQ,omitempty" json:"infQ,omitempty"`
	VCargaAverb *TDec1302Opc                 `xml:"http://www.portalfiscal.inf.br/cte vCargaAverb,omitempty" json:"vCargaAverb,omitempty"`
}

type TCTeSimpInfCTeDetComp struct {
	XNome string   `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	VComp TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vComp" json:"vComp"`
}

type TCTeSimpInfCTeDetInfNFe struct {
	ChNFe         string           `xml:"http://www.portalfiscal.inf.br/cte chNFe" json:"chNFe"`
	PIN           *string          `xml:"http://www.portalfiscal.inf.br/cte PIN,omitempty" json:"PIN,omitempty"`
	DPrev         *TData           `xml:"http://www.portalfiscal.inf.br/cte dPrev,omitempty" json:"dPrev,omitempty"`
	InfUnidCarga  []TUnidCarga     `xml:"http://www.portalfiscal.inf.br/cte infUnidCarga,omitempty" json:"infUnidCarga,omitempty"`
	InfUnidTransp []TUnidadeTransp `xml:"http://www.portalfiscal.inf.br/cte infUnidTransp,omitempty" json:"infUnidTransp,omitempty"`
}

type TCTeSimpInfCTeDetInfDocAntInfNFeTranspParcial struct {
	ChNFe TChDFe `xml:"http://www.portalfiscal.inf.br/cte chNFe" json:"chNFe"`
}

type TCTeSimpInfCTeDetInfDocAnt struct {
	ChCTe               TChDFe                                          `xml:"http://www.portalfiscal.inf.br/cte chCTe" json:"chCTe"`
	TpPrest             string                                          `xml:"http://www.portalfiscal.inf.br/cte tpPrest" json:"tpPrest"`
	InfNFeTranspParcial []TCTeSimpInfCTeDetInfDocAntInfNFeTranspParcial `xml:"http://www.portalfiscal.inf.br/cte infNFeTranspParcial,omitempty" json:"infNFeTranspParcial,omitempty"`
}

type TCTeSimpInfCTeDet struct {
	NItem     string                       `xml:"nItem,attr" json:"nItem"`
	CMunIni   TCodMunIBGE                  `xml:"http://www.portalfiscal.inf.br/cte cMunIni" json:"cMunIni"`
	XMunIni   string                       `xml:"http://www.portalfiscal.inf.br/cte xMunIni" json:"xMunIni"`
	CMunFim   TCodMunIBGE                  `xml:"http://www.portalfiscal.inf.br/cte cMunFim" json:"cMunFim"`
	XMunFim   string                       `xml:"http://www.portalfiscal.inf.br/cte xMunFim" json:"xMunFim"`
	VPrest    TDec1302                     `xml:"http://www.portalfiscal.inf.br/cte vPrest" json:"vPrest"`
	VPrestLiq *TDec1302                    `xml:"http://www.portalfiscal.inf.br/cte vPrestLiq,omitempty" json:"vPrestLiq,omitempty"`
	VRec      TDec1302                     `xml:"http://www.portalfiscal.inf.br/cte vRec" json:"vRec"`
	Comp      []TCTeSimpInfCTeDetComp      `xml:"http://www.portalfiscal.inf.br/cte Comp,omitempty" json:"Comp,omitempty"`
	IBSCBS    *TTribCTe                    `xml:"http://www.portalfiscal.inf.br/cte IBSCBS,omitempty" json:"IBSCBS,omitempty"`
	InfNFe    []TCTeSimpInfCTeDetInfNFe    `xml:"http://www.portalfiscal.inf.br/cte infNFe,omitempty" json:"infNFe,omitempty"`
	InfDocAnt []TCTeSimpInfCTeDetInfDocAnt `xml:"http://www.portalfiscal.inf.br/cte infDocAnt,omitempty" json:"infDocAnt,omitempty"`
}

type TCTeSimpInfCTeInfModal struct {
	VersaoModal string       `xml:"versaoModal,attr" json:"versaoModal"`
	Any         []RawElement `xml:",any" json:"any,omitempty"`
}

type TCTeSimpInfCTeCobrFat struct {
	NFat  *string      `xml:"http://www.portalfiscal.inf.br/cte nFat,omitempty" json:"nFat,omitempty"`
	VOrig *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vOrig,omitempty" json:"vOrig,omitempty"`
	VDesc *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vDesc,omitempty" json:"vDesc,omitempty"`
	VLiq  *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vLiq,omitempty" json:"vLiq,omitempty"`
}

type TCTeSimpInfCTeCobrDup struct {
	NDup  *string      `xml:"http://www.portalfiscal.inf.br/cte nDup,omitempty" json:"nDup,omitempty"`
	DVenc *TData       `xml:"http://www.portalfiscal.inf.br/cte dVenc,omitempty" json:"dVenc,omitempty"`
	VDup  *TDec1302Opc `xml:"http://www.portalfiscal.inf.br/cte vDup,omitempty" json:"vDup,omitempty"`
}

type TCTeSimpInfCTeCobr struct {
	Fat *TCTeSimpInfCTeCobrFat  `xml:"http://www.portalfiscal.inf.br/cte fat,omitempty" json:"fat,omitempty"`
	Dup []TCTeSimpInfCTeCobrDup `xml:"http://www.portalfiscal.inf.br/cte dup,omitempty" json:"dup,omitempty"`
}

type TCTeSimpInfCTeInfCTeSub struct {
	ChCTe         TChDFe  `xml:"http://www.portalfiscal.inf.br/cte chCte" json:"chCte"`
	IndAlteraToma *string `xml:"http://www.portalfiscal.inf.br/cte indAlteraToma,omitempty" json:"indAlteraToma,omitempty"`
}

type TCTeSimpInfCTeImpICMSUFFim struct {
	VBCUFFim   TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vBCUFFim" json:"vBCUFFim"`
	PFCPUFFim  TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pFCPUFFim" json:"pFCPUFFim"`
	PICMSUFFim TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pICMSUFFim" json:"pICMSUFFim"`
	PICMSInter TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pICMSInter" json:"pICMSInter"`
	VFCPUFFim  TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vFCPUFFim" json:"vFCPUFFim"`
	VICMSUFFim TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSUFFim" json:"vICMSUFFim"`
	VICMSUFIni TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSUFIni" json:"vICMSUFIni"`
}

type TCTeSimpInfCTeImp struct {
	ICMS       TImp                        `xml:"http://www.portalfiscal.inf.br/cte ICMS" json:"ICMS"`
	VTotTrib   *TDec1302                   `xml:"http://www.portalfiscal.inf.br/cte vTotTrib,omitempty" json:"vTotTrib,omitempty"`
	InfAdFisco *string                     `xml:"http://www.portalfiscal.inf.br/cte infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	ICMSUFFim  *TCTeSimpInfCTeImpICMSUFFim `xml:"http://www.portalfiscal.inf.br/cte ICMSUFFim,omitempty" json:"ICMSUFFim,omitempty"`
}

type TCTeSimpInfCTePgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/cte pgto,omitempty" json:"pgto,omitempty"`
}

type TCTeSimpInfCTeTotal struct {
	VTPrest    TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vTPrest" json:"vTPrest"`
	VTPrestLiq *TDec1302   `xml:"http://www.portalfiscal.inf.br/cte vTPrestLiq,omitempty" json:"vTPrestLiq,omitempty"`
	VTRec      TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vTRec" json:"vTRec"`
	IBSCBSTot  *TIBSCBSTot `xml:"http://www.portalfiscal.inf.br/cte IBSCBSTot,omitempty" json:"IBSCBSTot,omitempty"`
	VTotDFe    *TDec1302   `xml:"http://www.portalfiscal.inf.br/cte vTotDFe,omitempty" json:"vTotDFe,omitempty"`
}

type TCTeSimpInfCTeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
}

type TCTeSimpInfCTeInfSolicNFF struct {
	XSolic string `xml:"http://www.portalfiscal.inf.br/cte xSolic" json:"xSolic"`
}

type TCTeSimpInfCTeInfPAAPAASignature struct {
	SignatureValue Base64Binary     `xml:"http://www.portalfiscal.inf.br/cte SignatureValue" json:"SignatureValue"`
	RSAKeyValue    TRSAKeyValueType `xml:"http://www.portalfiscal.inf.br/cte RSAKeyValue" json:"RSAKeyValue"`
}

type TCTeSimpInfCTeInfPAA struct {
	CNPJPAA      TCNPJ                            `xml:"http://www.portalfiscal.inf.br/cte CNPJPAA" json:"CNPJPAA"`
	PAASignature TCTeSimpInfCTeInfPAAPAASignature `xml:"http://www.portalfiscal.inf.br/cte PAASignature" json:"PAASignature"`
}

type TCTeSimpInfCTe struct {
	Versao      string                     `xml:"versao,attr" json:"versao"`
	ID          string                     `xml:"Id,attr" json:"Id"`
	Ide         TCTeSimpInfCTeIde          `xml:"http://www.portalfiscal.inf.br/cte ide" json:"ide"`
	Compl       *TCTeSimpInfCTeCompl       `xml:"http://www.portalfiscal.inf.br/cte compl,omitempty" json:"compl,omitempty"`
	Emit        TCTeSimpInfCTeEmit         `xml:"http://www.portalfiscal.inf.br/cte emit" json:"emit"`
	Toma        TCTeSimpInfCTeToma         `xml:"http://www.portalfiscal.inf.br/cte toma" json:"toma"`
	InfCarga    TCTeSimpInfCTeInfCarga     `xml:"http://www.portalfiscal.inf.br/cte infCarga" json:"infCarga"`
	Det         []TCTeSimpInfCTeDet        `xml:"http://www.portalfiscal.inf.br/cte det,omitempty" json:"det,omitempty"`
	InfModal    TCTeSimpInfCTeInfModal     `xml:"http://www.portalfiscal.inf.br/cte infModal" json:"infModal"`
	Cobr        *TCTeSimpInfCTeCobr        `xml:"http://www.portalfiscal.inf.br/cte cobr,omitempty" json:"cobr,omitempty"`
	InfCTeSub   *TCTeSimpInfCTeInfCTeSub   `xml:"http://www.portalfiscal.inf.br/cte infCteSub,omitempty" json:"infCteSub,omitempty"`
	Imp         TCTeSimpInfCTeImp          `xml:"http://www.portalfiscal.inf.br/cte imp" json:"imp"`
	PgtoVinc    *TCTeSimpInfCTePgtoVinc    `xml:"http://www.portalfiscal.inf.br/cte pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	Total       TCTeSimpInfCTeTotal        `xml:"http://www.portalfiscal.inf.br/cte total" json:"total"`
	AutXML      []TCTeSimpInfCTeAutXML     `xml:"http://www.portalfiscal.inf.br/cte autXML,omitempty" json:"autXML,omitempty"`
	InfRespTec  *TRespTec                  `xml:"http://www.portalfiscal.inf.br/cte infRespTec,omitempty" json:"infRespTec,omitempty"`
	InfSolicNFF *TCTeSimpInfCTeInfSolicNFF `xml:"http://www.portalfiscal.inf.br/cte infSolicNFF,omitempty" json:"infSolicNFF,omitempty"`
	InfPAA      *TCTeSimpInfCTeInfPAA      `xml:"http://www.portalfiscal.inf.br/cte infPAA,omitempty" json:"infPAA,omitempty"`
}

type TCTeSimpInfCTeSupl struct {
	QRCodCTe string `xml:"http://www.portalfiscal.inf.br/cte qrCodCTe" json:"qrCodCTe"`
}

type TCTeSimp struct {
	InfCTe     TCTeSimpInfCTe      `xml:"http://www.portalfiscal.inf.br/cte infCte" json:"infCte"`
	InfCTeSupl *TCTeSimpInfCTeSupl `xml:"http://www.portalfiscal.inf.br/cte infCTeSupl,omitempty" json:"infCTeSupl,omitempty"`
	Signature  SignatureType       `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCFOP string

type TChDFe string

type TChDFeRTC string

type TCNPJ string

type TCNPJBaseRTC string

type TCNPJOpc string

type TCNPJRTC string

type TCodMunIBGE string

type TCodUFIBGE string

type TCompetApur string

type TCompraGov struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/cte tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/cte pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/cte tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/cte refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCompraGovReduzido struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/cte tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/cte pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/cte tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/cte refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TContainer string

type TCPF string

type TCPFVar string

type TCredPres struct {
	PCredPres        TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pCredPres" json:"pCredPres"`
	VCredPres        *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/cte vCredPres,omitempty" json:"vCredPres,omitempty"`
	VCredPresCondSus *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/cte vCredPresCondSus,omitempty" json:"vCredPresCondSus,omitempty"`
}

type TCredPresIBSZFM struct {
	CompetApur       TCompetApur       `xml:"http://www.portalfiscal.inf.br/cte competApur" json:"competApur"`
	TpCredPresIBSZFM TTpCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/cte tpCredPresIBSZFM" json:"tpCredPresIBSZFM"`
	VCredPresIBSZFM  TDec1302RTC       `xml:"http://www.portalfiscal.inf.br/cte vCredPresIBSZFM" json:"vCredPresIBSZFM"`
}

type TCredPresOper struct {
	VBCCredPres  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vBCCredPres" json:"vBCCredPres"`
	CCredPres    TcCredPres  `xml:"http://www.portalfiscal.inf.br/cte cCredPres" json:"cCredPres"`
	GIBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/cte gIBSCredPres,omitempty" json:"gIBSCredPres,omitempty"`
	GCBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/cte gCBSCredPres,omitempty" json:"gCBSCredPres,omitempty"`
}

type TData string

type TDateTimeUTC string

type TDec1104RTC string

type TDec1302RTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec030204RTC string

type TDec0303 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec0906Opc string

type TDec1104 string

type TDec1104OpRTC string

type TDec1104Opc string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec1302 string

type TDec1302Opc string

type TDevTrib struct {
	PDevTrib *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pDevTrib,omitempty" json:"pDevTrib,omitempty"`
	VDevTrib TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/cte vDevTrib" json:"vDevTrib"`
}

type TDif struct {
	PDif TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pDif" json:"pDif"`
	VDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vDif" json:"vDif"`
}

type TDocAssoc string

type TEmail string

type TEndOrg struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/cte xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/cte nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/cte xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/cte xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/cte cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/cte xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/cte CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/cte UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/cte cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/cte xPais,omitempty" json:"xPais,omitempty"`
	Fone    *TFone      `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
}

type TEndReEnt struct {
	CNPJ    *TCNPJ      `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TCPF       `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	XNome   string      `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/cte xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/cte nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/cte xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/cte xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/cte cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/cte xMun" json:"xMun"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/cte UF" json:"UF"`
}

type TEndeEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/cte xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/cte nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/cte xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/cte xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/cte cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/cte xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/cte CEP,omitempty" json:"CEP,omitempty"`
	UF      TUFSemEX    `xml:"http://www.portalfiscal.inf.br/cte UF" json:"UF"`
	Fone    *TFone      `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/cte xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/cte nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/cte xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/cte xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/cte cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/cte xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/cte CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/cte UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/cte cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/cte xPais,omitempty" json:"xPais,omitempty"`
}

type TEndernac struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/cte xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/cte nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/cte xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/cte xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/cte cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/cte xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/cte CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/cte UF" json:"UF"`
}

type TEnteGov string

type TEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBSEstCred" json:"vCBSEstCred"`
}

type TFinCTe string

type TFinCTeSimp string

type TFinGTVe string

type TFone string

type TGTVeInfCTeIdeToma struct {
	Toma string `xml:"http://www.portalfiscal.inf.br/cte toma" json:"toma"`
}

type TGTVeInfCTeIdeTomaTerceiro struct {
	Toma      string    `xml:"http://www.portalfiscal.inf.br/cte toma" json:"toma"`
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string   `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	Fone      *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	EnderToma TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderToma" json:"enderToma"`
	Email     *TEmail   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TGTVeInfCTeIde struct {
	CUF           TCodUFIBGE                  `xml:"http://www.portalfiscal.inf.br/cte cUF" json:"cUF"`
	CCT           string                      `xml:"http://www.portalfiscal.inf.br/cte cCT" json:"cCT"`
	CFOP          TCFOP                       `xml:"http://www.portalfiscal.inf.br/cte CFOP" json:"CFOP"`
	NatOp         string                      `xml:"http://www.portalfiscal.inf.br/cte natOp" json:"natOp"`
	Mod           TModGTVe                    `xml:"http://www.portalfiscal.inf.br/cte mod" json:"mod"`
	Serie         string                      `xml:"http://www.portalfiscal.inf.br/cte serie" json:"serie"`
	NCT           TNF                         `xml:"http://www.portalfiscal.inf.br/cte nCT" json:"nCT"`
	DhEmi         string                      `xml:"http://www.portalfiscal.inf.br/cte dhEmi" json:"dhEmi"`
	TpImp         string                      `xml:"http://www.portalfiscal.inf.br/cte tpImp" json:"tpImp"`
	TpEmis        string                      `xml:"http://www.portalfiscal.inf.br/cte tpEmis" json:"tpEmis"`
	CDV           string                      `xml:"http://www.portalfiscal.inf.br/cte cDV" json:"cDV"`
	TpAmb         TAmb                        `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	TpCTe         string                      `xml:"http://www.portalfiscal.inf.br/cte tpCTe" json:"tpCTe"`
	VerProc       string                      `xml:"http://www.portalfiscal.inf.br/cte verProc" json:"verProc"`
	CMunEnv       TCodMunIBGE                 `xml:"http://www.portalfiscal.inf.br/cte cMunEnv" json:"cMunEnv"`
	XMunEnv       string                      `xml:"http://www.portalfiscal.inf.br/cte xMunEnv" json:"xMunEnv"`
	UFEnv         TUF                         `xml:"http://www.portalfiscal.inf.br/cte UFEnv" json:"UFEnv"`
	Modal         string                      `xml:"http://www.portalfiscal.inf.br/cte modal" json:"modal"`
	TpServ        string                      `xml:"http://www.portalfiscal.inf.br/cte tpServ" json:"tpServ"`
	IndIEToma     string                      `xml:"http://www.portalfiscal.inf.br/cte indIEToma" json:"indIEToma"`
	DhSaidaOrig   string                      `xml:"http://www.portalfiscal.inf.br/cte dhSaidaOrig" json:"dhSaidaOrig"`
	DhChegadaDest string                      `xml:"http://www.portalfiscal.inf.br/cte dhChegadaDest" json:"dhChegadaDest"`
	Toma          *TGTVeInfCTeIdeToma         `xml:"http://www.portalfiscal.inf.br/cte toma,omitempty" json:"toma,omitempty"`
	TomaTerceiro  *TGTVeInfCTeIdeTomaTerceiro `xml:"http://www.portalfiscal.inf.br/cte tomaTerceiro,omitempty" json:"tomaTerceiro,omitempty"`
	DhCont        *TDateTimeUTC               `xml:"http://www.portalfiscal.inf.br/cte dhCont,omitempty" json:"dhCont,omitempty"`
	XJust         *string                     `xml:"http://www.portalfiscal.inf.br/cte xJust,omitempty" json:"xJust,omitempty"`
}

type TGTVeInfCTeComplObsCont struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/cte xTexto" json:"xTexto"`
}

type TGTVeInfCTeComplObsFisco struct {
	XCampo string `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/cte xTexto" json:"xTexto"`
}

type TGTVeInfCTeCompl struct {
	XCaracAd  *string                    `xml:"http://www.portalfiscal.inf.br/cte xCaracAd,omitempty" json:"xCaracAd,omitempty"`
	XCaracSer *string                    `xml:"http://www.portalfiscal.inf.br/cte xCaracSer,omitempty" json:"xCaracSer,omitempty"`
	XEmi      *string                    `xml:"http://www.portalfiscal.inf.br/cte xEmi,omitempty" json:"xEmi,omitempty"`
	XObs      *string                    `xml:"http://www.portalfiscal.inf.br/cte xObs,omitempty" json:"xObs,omitempty"`
	ObsCont   []TGTVeInfCTeComplObsCont  `xml:"http://www.portalfiscal.inf.br/cte ObsCont,omitempty" json:"ObsCont,omitempty"`
	ObsFisco  []TGTVeInfCTeComplObsFisco `xml:"http://www.portalfiscal.inf.br/cte ObsFisco,omitempty" json:"ObsFisco,omitempty"`
}

type TGTVeInfCTeEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/cte CNPJ" json:"CNPJ"`
	IE        string   `xml:"http://www.portalfiscal.inf.br/cte IE" json:"IE"`
	IEST      *string  `xml:"http://www.portalfiscal.inf.br/cte IEST,omitempty" json:"IEST,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/cte enderEmit" json:"enderEmit"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/cte ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TGTVeInfCTeRem struct {
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	XFant     *string   `xml:"http://www.portalfiscal.inf.br/cte xFant,omitempty" json:"xFant,omitempty"`
	Fone      *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	EnderReme TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderReme" json:"enderReme"`
	Email     *string   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TGTVeInfCTeDest struct {
	CNPJ      *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
	IE        *string   `xml:"http://www.portalfiscal.inf.br/cte IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/cte xNome" json:"xNome"`
	Fone      *TFone    `xml:"http://www.portalfiscal.inf.br/cte fone,omitempty" json:"fone,omitempty"`
	ISUF      *string   `xml:"http://www.portalfiscal.inf.br/cte ISUF,omitempty" json:"ISUF,omitempty"`
	EnderDest TEndereco `xml:"http://www.portalfiscal.inf.br/cte enderDest" json:"enderDest"`
	Email     *TEmail   `xml:"http://www.portalfiscal.inf.br/cte email,omitempty" json:"email,omitempty"`
}

type TGTVeInfCTeDetGTVInfEspecie struct {
	TpEspecie   string   `xml:"http://www.portalfiscal.inf.br/cte tpEspecie" json:"tpEspecie"`
	VEspecie    TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vEspecie" json:"vEspecie"`
	TpNumerario string   `xml:"http://www.portalfiscal.inf.br/cte tpNumerario" json:"tpNumerario"`
	XMoedaEstr  *string  `xml:"http://www.portalfiscal.inf.br/cte xMoedaEstr,omitempty" json:"xMoedaEstr,omitempty"`
}

type TGTVeInfCTeDetGTVInfVeiculo struct {
	Placa TPlaca  `xml:"http://www.portalfiscal.inf.br/cte placa" json:"placa"`
	UF    *TUF    `xml:"http://www.portalfiscal.inf.br/cte UF,omitempty" json:"UF,omitempty"`
	RNTRC *string `xml:"http://www.portalfiscal.inf.br/cte RNTRC,omitempty" json:"RNTRC,omitempty"`
}

type TGTVeInfCTeDetGTV struct {
	InfEspecie []TGTVeInfCTeDetGTVInfEspecie `xml:"http://www.portalfiscal.inf.br/cte infEspecie,omitempty" json:"infEspecie,omitempty"`
	QCarga     TDec1104                      `xml:"http://www.portalfiscal.inf.br/cte qCarga" json:"qCarga"`
	InfVeiculo []TGTVeInfCTeDetGTVInfVeiculo `xml:"http://www.portalfiscal.inf.br/cte infVeiculo,omitempty" json:"infVeiculo,omitempty"`
}

type TGTVeInfCTeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/cte CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/cte CPF,omitempty" json:"CPF,omitempty"`
}

type TGTVeInfCTe struct {
	Versao     string              `xml:"versao,attr" json:"versao"`
	ID         string              `xml:"Id,attr" json:"Id"`
	Ide        TGTVeInfCTeIde      `xml:"http://www.portalfiscal.inf.br/cte ide" json:"ide"`
	Compl      *TGTVeInfCTeCompl   `xml:"http://www.portalfiscal.inf.br/cte compl,omitempty" json:"compl,omitempty"`
	Emit       TGTVeInfCTeEmit     `xml:"http://www.portalfiscal.inf.br/cte emit" json:"emit"`
	Rem        TGTVeInfCTeRem      `xml:"http://www.portalfiscal.inf.br/cte rem" json:"rem"`
	Dest       TGTVeInfCTeDest     `xml:"http://www.portalfiscal.inf.br/cte dest" json:"dest"`
	Origem     *TEndeEmi           `xml:"http://www.portalfiscal.inf.br/cte origem,omitempty" json:"origem,omitempty"`
	Destino    *TEndeEmi           `xml:"http://www.portalfiscal.inf.br/cte destino,omitempty" json:"destino,omitempty"`
	DetGTV     TGTVeInfCTeDetGTV   `xml:"http://www.portalfiscal.inf.br/cte detGTV" json:"detGTV"`
	AutXML     []TGTVeInfCTeAutXML `xml:"http://www.portalfiscal.inf.br/cte autXML,omitempty" json:"autXML,omitempty"`
	InfRespTec *TRespTec           `xml:"http://www.portalfiscal.inf.br/cte infRespTec,omitempty" json:"infRespTec,omitempty"`
}

type TGTVeInfCTeSupl struct {
	QRCodCTe string `xml:"http://www.portalfiscal.inf.br/cte qrCodCTe" json:"qrCodCTe"`
}

type TGTVe struct {
	Versao     string           `xml:"versao,attr" json:"versao"`
	InfCTe     TGTVeInfCTe      `xml:"http://www.portalfiscal.inf.br/cte infCte" json:"infCte"`
	InfCTeSupl *TGTVeInfCTeSupl `xml:"http://www.portalfiscal.inf.br/cte infCTeSupl,omitempty" json:"infCTeSupl,omitempty"`
	Signature  SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TIBSCBSMonoTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSUF" json:"vIBSUF"`
}

type TIBSCBSMonoTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSMun" json:"vIBSMun"`
}

type TIBSCBSMonoTotGIBS struct {
	GIBSUF           TIBSCBSMonoTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/cte gIBSUF" json:"gIBSUF"`
	GIBSMun          TIBSCBSMonoTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/cte gIBSMun" json:"gIBSMun"`
	VIBS             TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/cte vIBS" json:"vIBS"`
	VCredPres        TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/cte vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/cte vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGCBS struct {
	VDif             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDif" json:"vDif"`
	VDevTrib         TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDevTrib" json:"vDevTrib"`
	VCBS             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBS" json:"vCBS"`
	VCredPres        TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGMono struct {
	VIBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSMono" json:"vIBSMono"`
	VCBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBSMono" json:"vCBSMono"`
	VIBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSMonoReten" json:"vIBSMonoReten"`
	VCBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBSMonoReten" json:"vCBSMonoReten"`
	VIBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSMonoRet" json:"vIBSMonoRet"`
	VCBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBSMonoRet" json:"vCBSMonoRet"`
}

type TIBSCBSMonoTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSMonoTot struct {
	VBCIBSCBS    TDec1302RTC                 `xml:"http://www.portalfiscal.inf.br/cte vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         *TIBSCBSMonoTotGIBS         `xml:"http://www.portalfiscal.inf.br/cte gIBS,omitempty" json:"gIBS,omitempty"`
	GCBS         *TIBSCBSMonoTotGCBS         `xml:"http://www.portalfiscal.inf.br/cte gCBS,omitempty" json:"gCBS,omitempty"`
	GMono        *TIBSCBSMonoTotGMono        `xml:"http://www.portalfiscal.inf.br/cte gMono,omitempty" json:"gMono,omitempty"`
	GEstornoCred *TIBSCBSMonoTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIBSCBSTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSUF" json:"vIBSUF"`
}

type TIBSCBSTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSMun" json:"vIBSMun"`
}

type TIBSCBSTotGIBS struct {
	GIBSUF  TIBSCBSTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/cte gIBSUF" json:"gIBSUF"`
	GIBSMun TIBSCBSTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/cte gIBSMun" json:"gIBSMun"`
	VIBS    TDec1302RTC           `xml:"http://www.portalfiscal.inf.br/cte vIBS" json:"vIBS"`
}

type TIBSCBSTotGCBS struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vDevTrib" json:"vDevTrib"`
	VCBS     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBS" json:"vCBS"`
}

type TIBSCBSTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSTot struct {
	VBCIBSCBS    TDec1302RTC             `xml:"http://www.portalfiscal.inf.br/cte vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         TIBSCBSTotGIBS          `xml:"http://www.portalfiscal.inf.br/cte gIBS" json:"gIBS"`
	GCBS         TIBSCBSTotGCBS          `xml:"http://www.portalfiscal.inf.br/cte gCBS" json:"gCBS"`
	GEstornoCred *TIBSCBSTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIPv4 string

type TIS struct {
	CSTIS        TCST           `xml:"http://www.portalfiscal.inf.br/cte CSTIS" json:"CSTIS"`
	CClassTribIS TcClassTrib    `xml:"http://www.portalfiscal.inf.br/cte cClassTribIS" json:"cClassTribIS"`
	VBCIS        *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vBCIS,omitempty" json:"vBCIS,omitempty"`
	PIS          *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pIS,omitempty" json:"pIS,omitempty"`
	AdRemIS      *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte adRemIS,omitempty" json:"adRemIS,omitempty"`
	UTrib        *string        `xml:"http://www.portalfiscal.inf.br/cte uTrib,omitempty" json:"uTrib,omitempty"`
	QTrib        *TDec1104OpRTC `xml:"http://www.portalfiscal.inf.br/cte qTrib,omitempty" json:"qTrib,omitempty"`
	VIS          *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIS,omitempty" json:"vIS,omitempty"`
}

type TISTot struct {
	VIS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIS" json:"vIS"`
}

type TIDLote string

type TIe string

type TIeDest string

type TImpICMS00 struct {
	CST   string   `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	VBC   TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	PICMS TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pICMS" json:"pICMS"`
	VICMS TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMS" json:"vICMS"`
}

type TImpICMS20 struct {
	CST        string      `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	PRedBC     TDec0302Opc `xml:"http://www.portalfiscal.inf.br/cte pRedBC" json:"pRedBC"`
	VBC        TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	PICMS      TDec0302    `xml:"http://www.portalfiscal.inf.br/cte pICMS" json:"pICMS"`
	VICMS      TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vICMS" json:"vICMS"`
	VICMSDeson *TDec1302   `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string     `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS45 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS60 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	VBCSTRet   TDec1302  `xml:"http://www.portalfiscal.inf.br/cte vBCSTRet" json:"vBCSTRet"`
	VICMSSTRet TDec1302  `xml:"http://www.portalfiscal.inf.br/cte vICMSSTRet" json:"vICMSSTRet"`
	PICMSSTRet TDec0302  `xml:"http://www.portalfiscal.inf.br/cte pICMSSTRet" json:"pICMSSTRet"`
	VCred      *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vCred,omitempty" json:"vCred,omitempty"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS90 struct {
	CST        string       `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	PRedBC     *TDec0302Opc `xml:"http://www.portalfiscal.inf.br/cte pRedBC,omitempty" json:"pRedBC,omitempty"`
	VBC        TDec1302     `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	PICMS      TDec0302     `xml:"http://www.portalfiscal.inf.br/cte pICMS" json:"pICMS"`
	VICMS      TDec1302     `xml:"http://www.portalfiscal.inf.br/cte vICMS" json:"vICMS"`
	VCred      *TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vCred,omitempty" json:"vCred,omitempty"`
	VICMSDeson *TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string      `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMSOutraUF struct {
	CST           string       `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	PRedBCOutraUF *TDec0302Opc `xml:"http://www.portalfiscal.inf.br/cte pRedBCOutraUF,omitempty" json:"pRedBCOutraUF,omitempty"`
	VBCOutraUF    TDec1302     `xml:"http://www.portalfiscal.inf.br/cte vBCOutraUF" json:"vBCOutraUF"`
	PICMSOutraUF  TDec0302     `xml:"http://www.portalfiscal.inf.br/cte pICMSOutraUF" json:"pICMSOutraUF"`
	VICMSOutraUF  TDec1302     `xml:"http://www.portalfiscal.inf.br/cte vICMSOutraUF" json:"vICMSOutraUF"`
	VICMSDeson    *TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef        *string      `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMSSN struct {
	CST   string `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	IndSN string `xml:"http://www.portalfiscal.inf.br/cte indSN" json:"indSN"`
}

type TImpGICMSPrevistoPagtoAntecip struct {
	VICMSPrevisto TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSPrevisto" json:"vICMSPrevisto"`
}

type TImp struct {
	ICMS00                    *TImpICMS00                    `xml:"http://www.portalfiscal.inf.br/cte ICMS00,omitempty" json:"ICMS00,omitempty"`
	ICMS20                    *TImpICMS20                    `xml:"http://www.portalfiscal.inf.br/cte ICMS20,omitempty" json:"ICMS20,omitempty"`
	ICMS45                    *TImpICMS45                    `xml:"http://www.portalfiscal.inf.br/cte ICMS45,omitempty" json:"ICMS45,omitempty"`
	ICMS60                    *TImpICMS60                    `xml:"http://www.portalfiscal.inf.br/cte ICMS60,omitempty" json:"ICMS60,omitempty"`
	ICMS90                    *TImpICMS90                    `xml:"http://www.portalfiscal.inf.br/cte ICMS90,omitempty" json:"ICMS90,omitempty"`
	ICMSOutraUF               *TImpICMSOutraUF               `xml:"http://www.portalfiscal.inf.br/cte ICMSOutraUF,omitempty" json:"ICMSOutraUF,omitempty"`
	ICMSSN                    *TImpICMSSN                    `xml:"http://www.portalfiscal.inf.br/cte ICMSSN,omitempty" json:"ICMSSN,omitempty"`
	GICMSPrevistoPagtoAntecip *TImpGICMSPrevistoPagtoAntecip `xml:"http://www.portalfiscal.inf.br/cte gICMSPrevistoPagtoAntecip,omitempty" json:"gICMSPrevistoPagtoAntecip,omitempty"`
}

type TImpOSICMS00 struct {
	CST   string   `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	VBC   TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	PICMS TDec0302 `xml:"http://www.portalfiscal.inf.br/cte pICMS" json:"pICMS"`
	VICMS TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMS" json:"vICMS"`
}

type TImpOSICMS20 struct {
	CST        string      `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	PRedBC     TDec0302Opc `xml:"http://www.portalfiscal.inf.br/cte pRedBC" json:"pRedBC"`
	VBC        TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	PICMS      TDec0302    `xml:"http://www.portalfiscal.inf.br/cte pICMS" json:"pICMS"`
	VICMS      TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vICMS" json:"vICMS"`
	VICMSDeson *TDec1302   `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string     `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpOSICMS45 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpOSICMS90 struct {
	CST        string       `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	PRedBC     *TDec0302Opc `xml:"http://www.portalfiscal.inf.br/cte pRedBC,omitempty" json:"pRedBC,omitempty"`
	VBC        TDec1302     `xml:"http://www.portalfiscal.inf.br/cte vBC" json:"vBC"`
	PICMS      TDec0302     `xml:"http://www.portalfiscal.inf.br/cte pICMS" json:"pICMS"`
	VICMS      TDec1302     `xml:"http://www.portalfiscal.inf.br/cte vICMS" json:"vICMS"`
	VCred      *TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vCred,omitempty" json:"vCred,omitempty"`
	VICMSDeson *TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string      `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpOSICMSOutraUF struct {
	CST           string       `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	PRedBCOutraUF *TDec0302Opc `xml:"http://www.portalfiscal.inf.br/cte pRedBCOutraUF,omitempty" json:"pRedBCOutraUF,omitempty"`
	VBCOutraUF    TDec1302     `xml:"http://www.portalfiscal.inf.br/cte vBCOutraUF" json:"vBCOutraUF"`
	PICMSOutraUF  TDec0302     `xml:"http://www.portalfiscal.inf.br/cte pICMSOutraUF" json:"pICMSOutraUF"`
	VICMSOutraUF  TDec1302     `xml:"http://www.portalfiscal.inf.br/cte vICMSOutraUF" json:"vICMSOutraUF"`
	VICMSDeson    *TDec1302    `xml:"http://www.portalfiscal.inf.br/cte vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef        *string      `xml:"http://www.portalfiscal.inf.br/cte cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpOSICMSSN struct {
	CST   string `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	IndSN string `xml:"http://www.portalfiscal.inf.br/cte indSN" json:"indSN"`
}

type TImpOSGICMSPrevistoPagtoAntecip struct {
	VICMSPrevisto TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vICMSPrevisto" json:"vICMSPrevisto"`
}

type TImpOS struct {
	ICMS00                    *TImpOSICMS00                    `xml:"http://www.portalfiscal.inf.br/cte ICMS00,omitempty" json:"ICMS00,omitempty"`
	ICMS20                    *TImpOSICMS20                    `xml:"http://www.portalfiscal.inf.br/cte ICMS20,omitempty" json:"ICMS20,omitempty"`
	ICMS45                    *TImpOSICMS45                    `xml:"http://www.portalfiscal.inf.br/cte ICMS45,omitempty" json:"ICMS45,omitempty"`
	ICMS90                    *TImpOSICMS90                    `xml:"http://www.portalfiscal.inf.br/cte ICMS90,omitempty" json:"ICMS90,omitempty"`
	ICMSOutraUF               *TImpOSICMSOutraUF               `xml:"http://www.portalfiscal.inf.br/cte ICMSOutraUF,omitempty" json:"ICMSOutraUF,omitempty"`
	ICMSSN                    *TImpOSICMSSN                    `xml:"http://www.portalfiscal.inf.br/cte ICMSSN,omitempty" json:"ICMSSN,omitempty"`
	GICMSPrevistoPagtoAntecip *TImpOSGICMSPrevistoPagtoAntecip `xml:"http://www.portalfiscal.inf.br/cte gICMSPrevistoPagtoAntecip,omitempty" json:"gICMSPrevistoPagtoAntecip,omitempty"`
}

type TIndDoacao string

type TJust string

type TLatitude string

type TLocal struct {
	CMun TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/cte cMun" json:"cMun"`
	XMun string      `xml:"http://www.portalfiscal.inf.br/cte xMun" json:"xMun"`
	UF   TUF         `xml:"http://www.portalfiscal.inf.br/cte UF" json:"UF"`
}

type TLongitude string

type TMed string

type TModCT string

type TModCTOS string

type TModCTCargaOS string

type TModDoc string

type TModGTVe string

type TModNF string

type TModTransp string

type TModTranspGTVe string

type TModTranspOS string

type TModTranspSimp string

type TMonofasiaGMonoPadrao struct {
	QBCMono  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/cte qBCMono" json:"qBCMono"`
	AdRemIBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte adRemIBS" json:"adRemIBS"`
	AdRemCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte adRemCBS" json:"adRemCBS"`
	VIBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSMono" json:"vIBSMono"`
	VCBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vCBSMono" json:"vCBSMono"`
}

type TMonofasiaGMonoReten struct {
	QBCMonoReten  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/cte qBCMonoReten" json:"qBCMonoReten"`
	AdRemIBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte adRemIBSReten" json:"adRemIBSReten"`
	VIBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSMonoReten" json:"vIBSMonoReten"`
	AdRemCBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte adRemCBSReten" json:"adRemCBSReten"`
	VCBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vCBSMonoReten" json:"vCBSMonoReten"`
}

type TMonofasiaGMonoRet struct {
	QBCMonoRet  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/cte qBCMonoRet" json:"qBCMonoRet"`
	AdRemIBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte adRemIBSRet" json:"adRemIBSRet"`
	VIBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSMonoRet" json:"vIBSMonoRet"`
	AdRemCBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte adRemCBSRet" json:"adRemCBSRet"`
	VCBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vCBSMonoRet" json:"vCBSMonoRet"`
}

type TMonofasiaGMonoDif struct {
	PDifIBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pDifIBS" json:"pDifIBS"`
	VIBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSMonoDif" json:"vIBSMonoDif"`
	PDifCBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pDifCBS" json:"pDifCBS"`
	VCBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vCBSMonoDif" json:"vCBSMonoDif"`
}

type TMonofasia struct {
	GMonoPadrao     *TMonofasiaGMonoPadrao `xml:"http://www.portalfiscal.inf.br/cte gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten      *TMonofasiaGMonoReten  `xml:"http://www.portalfiscal.inf.br/cte gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet        *TMonofasiaGMonoRet    `xml:"http://www.portalfiscal.inf.br/cte gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GMonoDif        *TMonofasiaGMonoDif    `xml:"http://www.portalfiscal.inf.br/cte gMonoDif,omitempty" json:"gMonoDif,omitempty"`
	VTotIBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/cte vTotIBSMonoItem" json:"vTotIBSMonoItem"`
	VTotCBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/cte vTotCBSMonoItem" json:"vTotCBSMonoItem"`
}

type TMotivo string

type TNF string

type TOperCompraGov string

type TPagAnt string

type TPagRef struct {
	RefDFe []TChDFeRTC `xml:"http://www.portalfiscal.inf.br/cte refDFe,omitempty" json:"refDFe,omitempty"`
}

type TPagamentoRTC struct {
	NPag        string       `xml:"nPag,attr" json:"nPag"`
	IDTransacao string       `xml:"idTransacao,attr" json:"idTransacao"`
	TpMeioPgto  string       `xml:"http://www.portalfiscal.inf.br/cte tpMeioPgto" json:"tpMeioPgto"`
	CNPJReceb   TCNPJRTC     `xml:"http://www.portalfiscal.inf.br/cte CNPJReceb" json:"CNPJReceb"`
	CNPJBasePSP TCNPJBaseRTC `xml:"http://www.portalfiscal.inf.br/cte CNPJBasePSP" json:"CNPJBasePSP"`
}

type TPlaca string

type TProcEmi string

type TProcEmiCTe string

type TProt string

type TProtCTeInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/cte verAplic" json:"verAplic"`
	ChCTe    TChDFe           `xml:"http://www.portalfiscal.inf.br/cte chCTe" json:"chCTe"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/cte dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/cte nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/cte digVal,omitempty" json:"digVal,omitempty"`
	CStat    TStat            `xml:"http://www.portalfiscal.inf.br/cte cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/cte xMotivo" json:"xMotivo"`
}

type TProtCTeInfFisco struct {
	CMsg TStat   `xml:"http://www.portalfiscal.inf.br/cte cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/cte xMsg" json:"xMsg"`
}

type TProtCTe struct {
	Versao    string            `xml:"versao,attr" json:"versao"`
	InfProt   TProtCTeInfProt   `xml:"http://www.portalfiscal.inf.br/cte infProt" json:"infProt"`
	InfFisco  *TProtCTeInfFisco `xml:"http://www.portalfiscal.inf.br/cte infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TProtCTeOSInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/cte verAplic" json:"verAplic"`
	ChCTe    TChDFe           `xml:"http://www.portalfiscal.inf.br/cte chCTe" json:"chCTe"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/cte dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/cte nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/cte digVal,omitempty" json:"digVal,omitempty"`
	CStat    TStat            `xml:"http://www.portalfiscal.inf.br/cte cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/cte xMotivo" json:"xMotivo"`
}

type TProtCTeOSInfFisco struct {
	CMsg TStat   `xml:"http://www.portalfiscal.inf.br/cte cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/cte xMsg" json:"xMsg"`
}

type TProtCTeOS struct {
	Versao    string              `xml:"versao,attr" json:"versao"`
	InfProt   TProtCTeOSInfProt   `xml:"http://www.portalfiscal.inf.br/cte infProt" json:"infProt"`
	InfFisco  *TProtCTeOSInfFisco `xml:"http://www.portalfiscal.inf.br/cte infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TProtGTVeInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/cte verAplic" json:"verAplic"`
	ChCTe    TChDFe           `xml:"http://www.portalfiscal.inf.br/cte chCTe" json:"chCTe"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/cte dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/cte nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/cte digVal,omitempty" json:"digVal,omitempty"`
	CStat    TStat            `xml:"http://www.portalfiscal.inf.br/cte cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/cte xMotivo" json:"xMotivo"`
}

type TProtGTVeInfFisco struct {
	CMsg TStat   `xml:"http://www.portalfiscal.inf.br/cte cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/cte xMsg" json:"xMsg"`
}

type TProtGTVe struct {
	Versao    string             `xml:"versao,attr" json:"versao"`
	InfProt   TProtGTVeInfProt   `xml:"http://www.portalfiscal.inf.br/cte infProt" json:"infProt"`
	InfFisco  *TProtGTVeInfFisco `xml:"http://www.portalfiscal.inf.br/cte infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType     `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRBSN string

type TRNTRC string

type TRSAKeyValueType struct {
	Modulus  Base64Binary `xml:"http://www.portalfiscal.inf.br/cte Modulus" json:"Modulus"`
	Exponent Base64Binary `xml:"http://www.portalfiscal.inf.br/cte Exponent" json:"Exponent"`
}

type TRec string

type TRed struct {
	PRedAliq  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqEfet" json:"pAliqEfet"`
}

type TRespTec struct {
	CNPJ     TCNPJ   `xml:"http://www.portalfiscal.inf.br/cte CNPJ" json:"CNPJ"`
	XContato string  `xml:"http://www.portalfiscal.inf.br/cte xContato" json:"xContato"`
	Email    TEmail  `xml:"http://www.portalfiscal.inf.br/cte email" json:"email"`
	Fone     string  `xml:"http://www.portalfiscal.inf.br/cte fone" json:"fone"`
	IDCSRT   *string `xml:"http://www.portalfiscal.inf.br/cte idCSRT,omitempty" json:"idCSRT,omitempty"`
	HashCSRT *string `xml:"http://www.portalfiscal.inf.br/cte hashCSRT,omitempty" json:"hashCSRT,omitempty"`
}

type TRetCTe struct {
	Versao   TVerCTe    `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb       `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/cte cUF" json:"cUF"`
	VerAplic TVerAplic  `xml:"http://www.portalfiscal.inf.br/cte verAplic" json:"verAplic"`
	CStat    TStat      `xml:"http://www.portalfiscal.inf.br/cte cStat" json:"cStat"`
	XMotivo  TMotivo    `xml:"http://www.portalfiscal.inf.br/cte xMotivo" json:"xMotivo"`
	ProtCTe  *TProtCTe  `xml:"http://www.portalfiscal.inf.br/cte protCTe,omitempty" json:"protCTe,omitempty"`
}

type TRetCTeOS struct {
	Versao   TVerCTe     `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb        `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE  `xml:"http://www.portalfiscal.inf.br/cte cUF" json:"cUF"`
	VerAplic TVerAplic   `xml:"http://www.portalfiscal.inf.br/cte verAplic" json:"verAplic"`
	CStat    TStat       `xml:"http://www.portalfiscal.inf.br/cte cStat" json:"cStat"`
	XMotivo  TMotivo     `xml:"http://www.portalfiscal.inf.br/cte xMotivo" json:"xMotivo"`
	ProtCTe  *TProtCTeOS `xml:"http://www.portalfiscal.inf.br/cte protCTe,omitempty" json:"protCTe,omitempty"`
}

type TRetCTeSimp struct {
	Versao   TVerCTe    `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb       `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/cte cUF" json:"cUF"`
	VerAplic TVerAplic  `xml:"http://www.portalfiscal.inf.br/cte verAplic" json:"verAplic"`
	CStat    TStat      `xml:"http://www.portalfiscal.inf.br/cte cStat" json:"cStat"`
	XMotivo  TMotivo    `xml:"http://www.portalfiscal.inf.br/cte xMotivo" json:"xMotivo"`
	ProtCTe  *TProtCTe  `xml:"http://www.portalfiscal.inf.br/cte protCTe,omitempty" json:"protCTe,omitempty"`
}

type TRetGTVe struct {
	Versao   TVerCTe    `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb       `xml:"http://www.portalfiscal.inf.br/cte tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/cte cUF" json:"cUF"`
	VerAplic TVerAplic  `xml:"http://www.portalfiscal.inf.br/cte verAplic" json:"verAplic"`
	CStat    TStat      `xml:"http://www.portalfiscal.inf.br/cte cStat" json:"cStat"`
	XMotivo  TMotivo    `xml:"http://www.portalfiscal.inf.br/cte xMotivo" json:"xMotivo"`
	ProtCTe  *TProtGTVe `xml:"http://www.portalfiscal.inf.br/cte protCTe,omitempty" json:"protCTe,omitempty"`
}

type TSerie string

type TServ string

type TStat string

type TString string

type TStringRTC string

type TTime string

type TTotalSN struct {
	VRBSNTot          TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/cte vRBSNTot" json:"vRBSNTot"`
	VIBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSSNTot,omitempty" json:"vIBSSNTot,omitempty"`
	VIBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBSSNtotPendSusp,omitempty" json:"vIBSSNtotPendSusp,omitempty"`
	VCBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBSSNTot,omitempty" json:"vCBSSNTot,omitempty"`
	VCBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBSSNtotPendSusp,omitempty" json:"vCBSSNtotPendSusp,omitempty"`
}

type TTpCredPresIBSZFM string

type TTransfCred struct {
	VIBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vIBS" json:"vIBS"`
	VCBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/cte vCBS" json:"vCBS"`
}

type TTribBPe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/cte cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/cte indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/cte gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCTe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/cte cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/cte indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/cte gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCompraGov struct {
	PAliqIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vTribCBS" json:"vTribCBS"`
}

type TTribItemSN struct {
	NItem        string         `xml:"nItem,attr" json:"nItem"`
	VRBSNItem    TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/cte vRBSNItem" json:"vRBSNItem"`
	TpRBSN       TRBSN          `xml:"http://www.portalfiscal.inf.br/cte tpRBSN" json:"tpRBSN"`
	PIBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pIBSSN,omitempty" json:"pIBSSN,omitempty"`
	VIBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSSN,omitempty" json:"vIBSSN,omitempty"`
	PCBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pCBSSN,omitempty" json:"pCBSSN,omitempty"`
	VCBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vCBSSN,omitempty" json:"vCBSSN,omitempty"`
	VIBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vIBSPendSusp,omitempty" json:"vIBSPendSusp,omitempty"`
	VCBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vCBSPendSusp,omitempty" json:"vCBSPendSusp,omitempty"`
}

type TTribNF3e struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/cte cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/cte indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/cte gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFAg struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/cte cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/cte indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/cte gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFCe struct {
	CST         TCST        `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	CClassTrib  TcClassTrib `xml:"http://www.portalfiscal.inf.br/cte cClassTrib" json:"cClassTrib"`
	IndDoacao   *TIndDoacao `xml:"http://www.portalfiscal.inf.br/cte indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS     *TCIBSNFe   `xml:"http://www.portalfiscal.inf.br/cte gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono *TMonofasia `xml:"http://www.portalfiscal.inf.br/cte gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
}

type TTribNFCom struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/cte cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/cte indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/cte gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFGas struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/cte cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/cte indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/cte gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFe struct {
	CST             TCST             `xml:"http://www.portalfiscal.inf.br/cte CST" json:"CST"`
	CClassTrib      TcClassTrib      `xml:"http://www.portalfiscal.inf.br/cte cClassTrib" json:"cClassTrib"`
	IndDoacao       *TIndDoacao      `xml:"http://www.portalfiscal.inf.br/cte indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS         *TCIBSNFe        `xml:"http://www.portalfiscal.inf.br/cte gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono     *TMonofasia      `xml:"http://www.portalfiscal.inf.br/cte gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
	GTransfCred     *TTransfCred     `xml:"http://www.portalfiscal.inf.br/cte gTransfCred,omitempty" json:"gTransfCred,omitempty"`
	GAjusteCompet   *TAjusteCompet   `xml:"http://www.portalfiscal.inf.br/cte gAjusteCompet,omitempty" json:"gAjusteCompet,omitempty"`
	GEstornoCred    *TEstornoCred    `xml:"http://www.portalfiscal.inf.br/cte gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
	GCredPresOper   *TCredPresOper   `xml:"http://www.portalfiscal.inf.br/cte gCredPresOper,omitempty" json:"gCredPresOper,omitempty"`
	GCredPresIBSZFM *TCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/cte gCredPresIBSZFM,omitempty" json:"gCredPresIBSZFM,omitempty"`
}

type TTribRegular struct {
	CSTReg             TCST          `xml:"http://www.portalfiscal.inf.br/cte CSTReg" json:"CSTReg"`
	CClassTribReg      TcClassTrib   `xml:"http://www.portalfiscal.inf.br/cte cClassTribReg" json:"cClassTribReg"`
	PAliqEfetRegIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqEfetRegIBSUF" json:"pAliqEfetRegIBSUF"`
	VTribRegIBSUF      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vTribRegIBSUF" json:"vTribRegIBSUF"`
	PAliqEfetRegIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqEfetRegIBSMun" json:"pAliqEfetRegIBSMun"`
	VTribRegIBSMun     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vTribRegIBSMun" json:"vTribRegIBSMun"`
	PAliqEfetRegCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/cte pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS        TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/cte vTribRegCBS" json:"vTribRegCBS"`
}

type TUFSemEX string

type TUF string

type TUnidCargaLacUnidCarga struct {
	NLacre string `xml:"http://www.portalfiscal.inf.br/cte nLacre" json:"nLacre"`
}

type TUnidCarga struct {
	TpUnidCarga  TtipoUnidCarga           `xml:"http://www.portalfiscal.inf.br/cte tpUnidCarga" json:"tpUnidCarga"`
	IDUnidCarga  TContainer               `xml:"http://www.portalfiscal.inf.br/cte idUnidCarga" json:"idUnidCarga"`
	LacUnidCarga []TUnidCargaLacUnidCarga `xml:"http://www.portalfiscal.inf.br/cte lacUnidCarga,omitempty" json:"lacUnidCarga,omitempty"`
	QtdRat       *TDec03020303            `xml:"http://www.portalfiscal.inf.br/cte qtdRat,omitempty" json:"qtdRat,omitempty"`
}

type TUnidadeTranspLacUnidTransp struct {
	NLacre string `xml:"http://www.portalfiscal.inf.br/cte nLacre" json:"nLacre"`
}

type TUnidadeTransp struct {
	TpUnidTransp  TtipoUnidTransp               `xml:"http://www.portalfiscal.inf.br/cte tpUnidTransp" json:"tpUnidTransp"`
	IDUnidTransp  TContainer                    `xml:"http://www.portalfiscal.inf.br/cte idUnidTransp" json:"idUnidTransp"`
	LacUnidTransp []TUnidadeTranspLacUnidTransp `xml:"http://www.portalfiscal.inf.br/cte lacUnidTransp,omitempty" json:"lacUnidTransp,omitempty"`
	InfUnidCarga  []TUnidCarga                  `xml:"http://www.portalfiscal.inf.br/cte infUnidCarga,omitempty" json:"infUnidCarga,omitempty"`
	QtdRat        *TDec03020303                 `xml:"http://www.portalfiscal.inf.br/cte qtdRat,omitempty" json:"qtdRat,omitempty"`
}

type TVerAplic string

type TVerCTe string

type Tano string

type TcClassTrib string

type TcCredPres string

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

type GYearMonth string

type String string

// RetCTe represents the global element {http://www.portalfiscal.inf.br/cte}retCTe.
type RetCTe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/cte retCTe" json:"-"`
	TRetCTe
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
