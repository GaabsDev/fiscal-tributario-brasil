// Package bpe implementa o Bilhete de Passagem Eletrônico (BP-e - Modelo 63).
// Layout oficial 1.00 com modalidades Rodoviário, Ferroviário, Aquaviário, BP-e TM e BP-e TA com Reforma Tributária (NT 2026.004 RTC).
package bpe

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloBPe       = "63"
	VersaoLayoutBPe = "1.00"
)

// BPe representa o Bilhete de Passagem Eletrônico (Modelo 63).
type BPe struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/bpe BPe" json:"-"`
	InfBPe    InfBPe         `xml:"infBPe" json:"infBPe"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfBPe struct {
	Versao string `xml:"versao,attr" json:"versao"`
	Id     string `xml:"Id,attr" json:"id"`
	Ide    Ide    `xml:"ide" json:"ide"`
	Emit   Emit   `xml:"emit" json:"emit"`
	Passag Passag `xml:"passag" json:"passag"`
	Viagem Viagem `xml:"viagem" json:"viagem"`
	VPrest VPrest `xml:"vPrest" json:"vPrest"`
	Imp    Imp    `xml:"imp" json:"imp"`
}

type Ide struct {
	CUF   string       `xml:"cUF" json:"cUF"`
	CBPe  string       `xml:"cBPe" json:"cBPe"`
	CFOP  string       `xml:"CFOP" json:"CFOP"`
	NatOp string       `xml:"natOp" json:"natOp"`
	Mod   string       `xml:"mod" json:"mod"`
	Serie string       `xml:"serie" json:"serie"`
	NBPe  string       `xml:"nBPe" json:"nBPe"`
	DhEmi nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	TpAmb string       `xml:"tpAmb" json:"tpAmb"`
	Modal string       `xml:"modal" json:"modal"` // 1-Rodoviário, 2-Ferroviário, 3-Aquaviário
	UFIni nfe.UF       `xml:"UFIni" json:"UFIni"`
	UFFim nfe.UF       `xml:"UFFim" json:"UFFim"`
}

type Emit struct {
	CNPJ      nfe.CNPJ      `xml:"CNPJ" json:"CNPJ"`
	IE        string        `xml:"IE" json:"IE"`
	XNome     string        `xml:"xNome" json:"xNome"`
	EnderEmit nfe.EnderEmit `xml:"enderEmit" json:"enderEmit"`
}

type Passag struct {
	XNome string   `xml:"xNome" json:"xNome"`
	CPF   *nfe.CPF `xml:"CPF,omitempty" json:"CPF,omitempty"`
}

type Viagem struct {
	DViagem nfe.Date `xml:"dViagem" json:"dViagem"`
	HViagem string   `xml:"hViagem" json:"hViagem"`
}

type VPrest struct {
	VTPrest nfe.Decimal `xml:"vTPrest" json:"vTPrest"`
	VRec    nfe.Decimal `xml:"vRec" json:"vRec"`
}

type Imp struct {
	ICMS   ICMS       `xml:"ICMS" json:"ICMS"`
	IBSCBS *IBSCBSBPe `xml:"IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type ICMS struct {
	ICMS00 *ICMS00 `xml:"ICMS00,omitempty" json:"ICMS00,omitempty"`
}

type ICMS00 struct {
	CST   string      `xml:"CST" json:"CST"`
	VBC   nfe.Decimal `xml:"vBC" json:"vBC"`
	PICMS nfe.Decimal `xml:"pICMS" json:"pICMS"`
	VICMS nfe.Decimal `xml:"vICMS" json:"vICMS"`
}

type IBSCBSBPe struct {
	CST        string         `xml:"CST" json:"CST"`
	CClassTrib nfe.CClassTrib `xml:"cClassTrib" json:"cClassTrib"`
	VBC        nfe.Decimal    `xml:"vBC" json:"vBC"`
	PIBS       nfe.Decimal    `xml:"pIBS" json:"pIBS"`
	VIBS       nfe.Decimal    `xml:"vIBS" json:"vIBS"`
	PCBS       nfe.Decimal    `xml:"pCBS" json:"pCBS"`
	VCBS       nfe.Decimal    `xml:"vCBS" json:"vCBS"`
}

// NewBPe cria um novo BP-e (Modelo 63).
func NewBPe(chave nfe.ChaveAcesso) *BPe {
	id := fmt.Sprintf("BPe%s", string(chave))
	return &BPe{
		InfBPe: InfBPe{
			Versao: VersaoLayoutBPe,
			Id:     id,
			Ide: Ide{
				Mod: ModeloBPe,
			},
		},
	}
}

func (doc *BPe) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *BPe) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
