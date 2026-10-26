// Package mdfe implementa o Manifesto Eletrônico de Documentos Fiscais (MDF-e - Modelo 58).
// Layout oficial 3.00 consolidado com todos os 4 modais: Rodoviário, Aéreo, Aquaviário e Ferroviário.
package mdfe

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloMDFe       = "58"
	VersaoLayoutMDFe = "3.00"
)

// MDFe representa o Manifesto Eletrônico de Documentos Fiscais (Modelo 58).
type MDFe struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/mdfe MDFe" json:"-"`
	InfMDFe   InfMDFe        `xml:"infMDFe" json:"infMDFe"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

// InfMDFe reúne os blocos de dados do manifesto.
type InfMDFe struct {
	Versao   string   `xml:"versao,attr" json:"versao"`
	Id       string   `xml:"Id,attr" json:"id"`
	Ide      Ide      `xml:"ide" json:"ide"`
	Emit     Emit     `xml:"emit" json:"emit"`
	InfModal InfModal `xml:"infModal" json:"infModal"`
	InfDoc   InfDoc   `xml:"infDoc" json:"infDoc"`
	Tot      Tot      `xml:"tot" json:"tot"`
}

type Ide struct {
	CUF     string       `xml:"cUF" json:"cUF"`
	TpAmb   string       `xml:"tpAmb" json:"tpAmb"`
	TpEmit  string       `xml:"tpEmit" json:"tpEmit"`
	Mod     string       `xml:"mod" json:"mod"`
	Serie   string       `xml:"serie" json:"serie"`
	NMDF    string       `xml:"nMDF" json:"nMDF"`
	CMDF    string       `xml:"cMDF" json:"cMDF"`
	CDV     string       `xml:"cDV" json:"cDV"`
	Modal   string       `xml:"modal" json:"modal"` // 1-Rodoviário, 2-Aéreo, 3-Aquaviário, 4-Ferroviário
	DhEmi   nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	TpEmis  string       `xml:"tpEmis" json:"tpEmis"`
	ProcEmi string       `xml:"procEmi" json:"procEmi"`
	VerProc string       `xml:"verProc" json:"verProc"`
	UFIni   nfe.UF       `xml:"UFIni" json:"UFIni"`
	UFFim   nfe.UF       `xml:"UFFim" json:"UFFim"`
}

type Emit struct {
	CNPJ      nfe.CNPJ      `xml:"CNPJ" json:"CNPJ"`
	IE        string        `xml:"IE" json:"IE"`
	XNome     string        `xml:"xNome" json:"xNome"`
	EnderEmit nfe.EnderEmit `xml:"enderEmit" json:"enderEmit"`
}

// InfModal contempla os 4 modais do MDF-e.
type InfModal struct {
	VersaoModal string  `xml:"versaoModal,attr" json:"versaoModal"`
	Rodo        *Rodo   `xml:"rodo,omitempty" json:"rodo,omitempty"`
	Aereo       *Aereo  `xml:"aereo,omitempty" json:"aereo,omitempty"`
	Aquav       *Aquav  `xml:"aquav,omitempty" json:"aquav,omitempty"`
	Ferrov      *Ferrov `xml:"ferrov,omitempty" json:"ferrov,omitempty"`
}

// Rodo modal rodoviário (mdfeModalRodoviario_v3.00.xsd).
type Rodo struct {
	InfANTT     *InfANTT   `xml:"infANTT,omitempty" json:"infANTT,omitempty"`
	VeicTracao  VeicTracao `xml:"veicTracao" json:"veicTracao"`
	VeicReboque []VeicReb  `xml:"veicReboque,omitempty" json:"veicReboque,omitempty"`
}

type InfANTT struct {
	RNTRC string `xml:"RNTRC,omitempty" json:"RNTRC,omitempty"`
}

type VeicTracao struct {
	Placa string `xml:"placa" json:"placa"`
	Tara  int    `xml:"tara" json:"tara"`
	CapKG int    `xml:"capKG,omitempty" json:"capKG,omitempty"`
	UF    nfe.UF `xml:"UF" json:"UF"`
}

type VeicReb struct {
	Placa string `xml:"placa" json:"placa"`
	Tara  int    `xml:"tara" json:"tara"`
	UF    nfe.UF `xml:"UF" json:"UF"`
}

// Aereo modal aéreo (mdfeModalAereo_v3.00.xsd).
type Aereo struct {
	Nac     string   `xml:"nac" json:"nac"`
	Matr    string   `xml:"matr" json:"matr"`
	NVoo    string   `xml:"nVoo" json:"nVoo"`
	CAerEmb string   `xml:"cAerEmb" json:"cAerEmb"`
	CAerDes string   `xml:"cAerDes" json:"cAerDes"`
	DVoo    nfe.Date `xml:"dVoo" json:"dVoo"`
}

// Aquav modal aquaviário (mdfeModalAquaviario_v3.00.xsd).
type Aquav struct {
	Irin    string `xml:"irin" json:"irin"`
	TpEmb   string `xml:"tpEmb" json:"tpEmb"`
	CEmb    string `xml:"cEmb" json:"cEmb"`
	XEmb    string `xml:"xEmb" json:"xEmb"`
	NViag   string `xml:"nViag" json:"nViag"`
	CPrtEmb string `xml:"cPrtEmb" json:"cPrtEmb"`
	CPrtDes string `xml:"cPrtDes" json:"cPrtDes"`
}

// Ferrov modal ferroviário (mdfeModalFerroviario_v3.00.xsd).
type Ferrov struct {
	Trem Trem  `xml:"trem" json:"trem"`
	Vag  []Vag `xml:"vag" json:"vag"`
}

type Trem struct {
	XPref string `xml:"xPref" json:"xPref"`
}

type Vag struct {
	PesoBC nfe.Decimal `xml:"pesoBC" json:"pesoBC"`
	PesoR  nfe.Decimal `xml:"pesoR" json:"pesoR"`
	TpVag  string      `xml:"tpVag,omitempty" json:"tpVag,omitempty"`
	Serie  string      `xml:"serie" json:"serie"`
	NVag   int         `xml:"nVag" json:"nVag"`
}

type InfDoc struct {
	InfMunDescarga []InfMunDescarga `xml:"infMunDescarga" json:"infMunDescarga"`
}

type InfMunDescarga struct {
	CMunDescarga string      `xml:"cMunDescarga" json:"cMunDescarga"`
	XMunDescarga string      `xml:"xMunDescarga" json:"xMunDescarga"`
	InfCTe       []InfCTeRef `xml:"infCTe,omitempty" json:"infCTe,omitempty"`
	InfNFe       []InfNFeRef `xml:"infNFe,omitempty" json:"infNFe,omitempty"`
}

type InfCTeRef struct {
	ChCTe nfe.ChaveAcesso `xml:"chCTe" json:"chCTe"`
}

type InfNFeRef struct {
	ChNFe nfe.ChaveAcesso `xml:"chNFe" json:"chNFe"`
}

type Tot struct {
	QCTe   int         `xml:"qCTe,omitempty" json:"qCTe,omitempty"`
	QNFe   int         `xml:"qNFe,omitempty" json:"qNFe,omitempty"`
	VCarga nfe.Decimal `xml:"vCarga" json:"vCarga"`
	CUnid  string      `xml:"cUnid" json:"cUnid"`
	QCarga nfe.Decimal `xml:"qCarga" json:"qCarga"`
}

// NewMDFe cria um novo MDF-e (Modelo 58).
func NewMDFe(chave nfe.ChaveAcesso) *MDFe {
	id := fmt.Sprintf("MDFe%s", string(chave))
	return &MDFe{
		InfMDFe: InfMDFe{
			Versao: VersaoLayoutMDFe,
			Id:     id,
			Ide: Ide{
				Mod:    ModeloMDFe,
				TpEmit: "1", // Prestador de serviço de transporte
			},
		},
	}
}

func (doc *MDFe) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *MDFe) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
