// Package cte implementa o Conhecimento de Transporte Eletrônico (CT-e - Modelo 57, CT-e OS - Modelo 67 e GTV-e - Modelo 64).
// Layout oficial 4.00 consolidado com as Notas Técnicas vigentes da SEFAZ e Reforma Tributária (NT 2026.004 RTC 1.00).
package cte

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloCTe       = "57"
	ModeloCTeOS     = "67"
	ModeloGTVe      = "64"
	VersaoLayoutCTe = "4.00"
)

// CTe representa o Conhecimento de Transporte Eletrônico de Cargas (Modelo 57).
type CTe struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/cte CTe" json:"-"`
	InfCte    InfCte         `xml:"infCte" json:"infCte"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

// InfCte reúne os blocos de dados do CT-e 57.
type InfCte struct {
	Versao     string      `xml:"versao,attr" json:"versao"`
	Id         string      `xml:"Id,attr" json:"id"`
	Ide        Ide         `xml:"ide" json:"ide"`
	Compl      *Compl      `xml:"compl,omitempty" json:"compl,omitempty"`
	Emit       Emit        `xml:"emit" json:"emit"`
	Rem        *Rem        `xml:"rem,omitempty" json:"rem,omitempty"`
	Exped      *Exped      `xml:"exped,omitempty" json:"exped,omitempty"`
	Receb      *Receb      `xml:"receb,omitempty" json:"receb,omitempty"`
	Dest       *Dest       `xml:"dest,omitempty" json:"dest,omitempty"`
	VPrest     VPrest      `xml:"vPrest" json:"vPrest"`
	Imp        Imp         `xml:"imp" json:"imp"`
	InfCTeNorm *InfCTeNorm `xml:"infCTeNorm,omitempty" json:"infCTeNorm,omitempty"`
}

// Ide identifica as características da prestação do serviço de transporte.
type Ide struct {
	CUF       string       `xml:"cUF" json:"cUF"`
	CCT       string       `xml:"cCT" json:"cCT"`
	CFOP      string       `xml:"CFOP" json:"CFOP"`
	NatOp     string       `xml:"natOp" json:"natOp"`
	Mod       string       `xml:"mod" json:"mod"`
	Serie     string       `xml:"serie" json:"serie"`
	NCT       string       `xml:"nCT" json:"nCT"`
	DhEmi     nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	TpImp     string       `xml:"tpImp" json:"tpImp"`
	TpEmis    string       `xml:"tpEmis" json:"tpEmis"`
	CDV       string       `xml:"cDV" json:"cDV"`
	TpAmb     string       `xml:"tpAmb" json:"tpAmb"`
	TpCTe     string       `xml:"tpCTe" json:"tpCTe"`
	ProcEmi   string       `xml:"procEmi" json:"procEmi"`
	VerProc   string       `xml:"verProc" json:"verProc"`
	CMunIni   string       `xml:"cMunIni" json:"cMunIni"`
	XMunIni   string       `xml:"xMunIni" json:"xMunIni"`
	UFIni     nfe.UF       `xml:"UFIni" json:"UFIni"`
	CMunFim   string       `xml:"cMunFim" json:"cMunFim"`
	XMunFim   string       `xml:"xMunFim" json:"xMunFim"`
	UFFim     nfe.UF       `xml:"UFFim" json:"UFFim"`
	Retira    string       `xml:"retira" json:"retira"`
	IndIEToma string       `xml:"indIEToma" json:"indIEToma"`
	Toma3     *Toma3       `xml:"toma3,omitempty" json:"toma3,omitempty"`
	Toma4     *Toma4       `xml:"toma4,omitempty" json:"toma4,omitempty"`
}

type Toma3 struct {
	Toma string `xml:"toma" json:"toma"`
}

type Toma4 struct {
	Toma  string    `xml:"toma" json:"toma"`
	CNPJ  *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	IE    string    `xml:"IE,omitempty" json:"IE,omitempty"`
	XNome string    `xml:"xNome" json:"xNome"`
}

type Compl struct {
	XCaracAd  string `xml:"xCaracAd,omitempty" json:"xCaracAd,omitempty"`
	XCaracSer string `xml:"xCaracSer,omitempty" json:"xCaracSer,omitempty"`
	XEmi      string `xml:"xEmi,omitempty" json:"xEmi,omitempty"`
	OrigCalc  string `xml:"origCalc,omitempty" json:"origCalc,omitempty"`
	DestCalc  string `xml:"destCalc,omitempty" json:"destCalc,omitempty"`
	XObs      string `xml:"xObs,omitempty" json:"xObs,omitempty"`
}

// Emit identifica o transportador emitente do CT-e.
type Emit struct {
	CNPJ      nfe.CNPJ `xml:"CNPJ" json:"CNPJ"`
	IE        string   `xml:"IE" json:"IE"`
	XNome     string   `xml:"xNome" json:"xNome"`
	XFant     string   `xml:"xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit EnderCTe `xml:"enderEmit" json:"enderEmit"`
	CRT       string   `xml:"CRT,omitempty" json:"CRT,omitempty"`
}

type EnderCTe struct {
	XLgr    string              `xml:"xLgr" json:"xLgr"`
	Nro     string              `xml:"nro" json:"nro"`
	XCpl    string              `xml:"xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string              `xml:"xBairro" json:"xBairro"`
	CMun    nfe.CodigoMunicipio `xml:"cMun" json:"cMun"`
	XMun    string              `xml:"xMun" json:"xMun"`
	CEP     string              `xml:"CEP,omitempty" json:"CEP,omitempty"`
	UF      nfe.UF              `xml:"UF" json:"UF"`
	Fone    string              `xml:"fone,omitempty" json:"fone,omitempty"`
}

// Rem remetente da carga.
type Rem struct {
	CNPJ     *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF      *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	IE       string    `xml:"IE,omitempty" json:"IE,omitempty"`
	XNome    string    `xml:"xNome" json:"xNome"`
	EnderRme EnderCTe  `xml:"enderRme" json:"enderRme"`
}

// Exped expedidor da carga.
type Exped struct {
	CNPJ       *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF        *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	IE         string    `xml:"IE,omitempty" json:"IE,omitempty"`
	XNome      string    `xml:"xNome" json:"xNome"`
	EnderExped EnderCTe  `xml:"enderExped" json:"enderExped"`
}

// Receb recebedor da carga.
type Receb struct {
	CNPJ       *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF        *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	IE         string    `xml:"IE,omitempty" json:"IE,omitempty"`
	XNome      string    `xml:"xNome" json:"xNome"`
	EnderReceb EnderCTe  `xml:"enderReceb" json:"enderReceb"`
}

// Dest destinatário da carga.
type Dest struct {
	CNPJ      *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	IE        string    `xml:"IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"xNome" json:"xNome"`
	EnderDest EnderCTe  `xml:"enderDest" json:"enderDest"`
}

// VPrest valores da prestação do serviço de transporte.
type VPrest struct {
	VTPrest nfe.Decimal `xml:"vTPrest" json:"vTPrest"`
	VRec    nfe.Decimal `xml:"vRec" json:"vRec"`
	Comp    []Comp      `xml:"Comp,omitempty" json:"Comp,omitempty"`
}

type Comp struct {
	XNome string      `xml:"xNome" json:"xNome"`
	VComp nfe.Decimal `xml:"vComp" json:"vComp"`
}

// Imp detalha os tributos incidentes (ICMS, Reforma Tributária IBS/CBS).
type Imp struct {
	ICMS     ICMS         `xml:"ICMS" json:"ICMS"`
	VTotTrib *nfe.Decimal `xml:"vTotTrib,omitempty" json:"vTotTrib,omitempty"`
	IBSCBS   *IBSCBSCTe   `xml:"IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type ICMS struct {
	ICMS00 *ICMS00 `xml:"ICMS00,omitempty" json:"ICMS00,omitempty"`
	ICMS20 *ICMS20 `xml:"ICMS20,omitempty" json:"ICMS20,omitempty"`
	ICMS45 *ICMS45 `xml:"ICMS45,omitempty" json:"ICMS45,omitempty"`
	ICMS90 *ICMS90 `xml:"ICMS90,omitempty" json:"ICMS90,omitempty"`
	ICMSSN *ICMSSN `xml:"ICMSSN,omitempty" json:"ICMSSN,omitempty"`
}

type ICMS00 struct {
	CST   string      `xml:"CST" json:"CST"`
	VBC   nfe.Decimal `xml:"vBC" json:"vBC"`
	PICMS nfe.Decimal `xml:"pICMS" json:"pICMS"`
	VICMS nfe.Decimal `xml:"vICMS" json:"vICMS"`
}

type ICMS20 struct {
	CST    string      `xml:"CST" json:"CST"`
	PRedBC nfe.Decimal `xml:"pRedBC" json:"pRedBC"`
	VBC    nfe.Decimal `xml:"vBC" json:"vBC"`
	PICMS  nfe.Decimal `xml:"pICMS" json:"pICMS"`
	VICMS  nfe.Decimal `xml:"vICMS" json:"vICMS"`
}

type ICMS45 struct {
	CST string `xml:"CST" json:"CST"`
}

type ICMS90 struct {
	CST   string      `xml:"CST" json:"CST"`
	VBC   nfe.Decimal `xml:"vBC" json:"vBC"`
	PICMS nfe.Decimal `xml:"pICMS" json:"pICMS"`
	VICMS nfe.Decimal `xml:"vICMS" json:"vICMS"`
}

type ICMSSN struct {
	CST   string `xml:"CST" json:"CST"`
	IndSN string `xml:"indSN" json:"indSN"`
}

// IBSCBSCTe representa a tributação da Reforma Tributária no CT-e (NT 2026.004 RTC).
type IBSCBSCTe struct {
	CST        string         `xml:"CST" json:"CST"`
	CClassTrib nfe.CClassTrib `xml:"cClassTrib" json:"cClassTrib"`
	VBC        nfe.Decimal    `xml:"vBC" json:"vBC"`
	PIBS       nfe.Decimal    `xml:"pIBS" json:"pIBS"`
	VIBS       nfe.Decimal    `xml:"vIBS" json:"vIBS"`
	PCBS       nfe.Decimal    `xml:"pCBS" json:"pCBS"`
	VCBS       nfe.Decimal    `xml:"vCBS" json:"vCBS"`
}

// InfCTeNorm informações do CT-e Normal e Substituto com modais completos.
type InfCTeNorm struct {
	InfCarga InfCarga `xml:"infCarga" json:"infCarga"`
	InfDoc   *InfDoc  `xml:"infDoc,omitempty" json:"infDoc,omitempty"`
	InfModal InfModal `xml:"infModal" json:"infModal"`
}

type InfCarga struct {
	VCarga  nfe.Decimal `xml:"vCarga" json:"vCarga"`
	ProPred string      `xml:"proPred" json:"proPred"`
	XOutCat string      `xml:"xOutCat,omitempty" json:"xOutCat,omitempty"`
	InfQ    []InfQ      `xml:"infQ" json:"infQ"`
}

type InfQ struct {
	CUnid  string      `xml:"cUnid" json:"cUnid"`
	TpMed  string      `xml:"tpMed" json:"tpMed"`
	QCarga nfe.Decimal `xml:"qCarga" json:"qCarga"`
}

type InfDoc struct {
	InfNFe []InfNFeRef `xml:"infNFe,omitempty" json:"infNFe,omitempty"`
}

type InfNFeRef struct {
	Chave nfe.ChaveAcesso `xml:"chave" json:"chave"`
}

// InfModal suporta todos os modais oficiais do CT-e: Rodoviário, Aéreo, Aquaviário, Ferroviário, Dutoviário, Multimodal.
type InfModal struct {
	VersaoModal string      `xml:"versaoModal,attr" json:"versaoModal"`
	Rodo        *Rodo       `xml:"rodo,omitempty" json:"rodo,omitempty"`
	Aereo       *Aereo      `xml:"aereo,omitempty" json:"aereo,omitempty"`
	Aquav       *Aquav      `xml:"aquav,omitempty" json:"aquav,omitempty"`
	Ferrov      *Ferrov     `xml:"ferrov,omitempty" json:"ferrov,omitempty"`
	Duto        *Duto       `xml:"duto,omitempty" json:"duto,omitempty"`
	MultiModal  *MultiModal `xml:"multimodal,omitempty" json:"multimodal,omitempty"`
}

// Rodo modal rodoviário (cteModalRodoviario_v4.00.xsd).
type Rodo struct {
	RNTRC string `xml:"RNTRC" json:"RNTRC"`
	Occ   []Occ  `xml:"occ,omitempty" json:"occ,omitempty"`
}

type Occ struct {
	NOcc int      `xml:"nOcc" json:"nOcc"`
	DOcc nfe.Date `xml:"dOcc" json:"dOcc"`
}

// Aereo modal aéreo (cteModalAereo_v4.00.xsd).
type Aereo struct {
	NMinu      string   `xml:"nMinu,omitempty" json:"nMinu,omitempty"`
	NOCA       string   `xml:"nOCA,omitempty" json:"nOCA,omitempty"`
	DPrevAereo nfe.Date `xml:"dPrevAereo" json:"dPrevAereo"`
	NatCarga   NatCarga `xml:"natCarga" json:"natCarga"`
	Tarifa     Tarifa   `xml:"tarifa" json:"tarifa"`
}

type NatCarga struct {
	XDime    string   `xml:"xDime,omitempty" json:"xDime,omitempty"`
	CInfManu []string `xml:"cInfManu,omitempty" json:"cInfManu,omitempty"`
}

type Tarifa struct {
	CL   string      `xml:"CL" json:"CL"`
	CCod string      `xml:"cCod,omitempty" json:"cCod,omitempty"`
	VTar nfe.Decimal `xml:"vTar" json:"vTar"`
}

// Aquav modal aquaviário (cteModalAquaviario_v4.00.xsd).
type Aquav struct {
	VPrest   nfe.Decimal `xml:"vPrest" json:"vPrest"`
	VAFRMM   nfe.Decimal `xml:"vAFRMM" json:"vAFRMM"`
	NBooking string      `xml:"nBooking,omitempty" json:"nBooking,omitempty"`
	NViag    string      `xml:"nViag,omitempty" json:"nViag,omitempty"`
	Direc    string      `xml:"direc" json:"direc"`
	Irin     string      `xml:"irin" json:"irin"`
	PrtEmb   string      `xml:"prtEmb" json:"prtEmb"`
	PrtDest  string      `xml:"prtDest" json:"prtDest"`
	TpNav    string      `xml:"tpNav" json:"tpNav"`
}

// Ferrov modal ferroviário (cteModalFerroviario_v4.00.xsd).
type Ferrov struct {
	TpTraf string   `xml:"tpTraf" json:"tpTraf"`
	Fluxo  string   `xml:"fluxo" json:"fluxo"`
	DetVag []DetVag `xml:"detVag" json:"detVag"`
}

type DetVag struct {
	NVag   int         `xml:"nVag" json:"nVag"`
	Cap    nfe.Decimal `xml:"cap" json:"cap"`
	TpVag  string      `xml:"tpVag,omitempty" json:"tpVag,omitempty"`
	PesoR  nfe.Decimal `xml:"pesoR" json:"pesoR"`
	PesoBC nfe.Decimal `xml:"pesoBC" json:"pesoBC"`
}

// Duto modal dutoviário (cteModalDutoviario_v4.00.xsd).
type Duto struct {
	VTar nfe.Decimal `xml:"vTar,omitempty" json:"vTar,omitempty"`
	DIni nfe.Date    `xml:"dIni" json:"dIni"`
	DFim nfe.Date    `xml:"dFim" json:"dFim"`
}

// MultiModal modal multimodal (cteMultiModal_v4.00.xsd).
type MultiModal struct {
	COTM          string `xml:"COTM" json:"COTM"`
	IndNegociavel string `xml:"indNegociavel" json:"indNegociavel"`
}

// NewCTe cria um novo CT-e (Modelo 57).
func NewCTe(chave nfe.ChaveAcesso) *CTe {
	id := fmt.Sprintf("CTe%s", string(chave))
	return &CTe{
		InfCte: InfCte{
			Versao: VersaoLayoutCTe,
			Id:     id,
			Ide: Ide{
				Mod:   ModeloCTe,
				TpCTe: "0", // Normal
			},
		},
	}
}

func (doc *CTe) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *CTe) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
