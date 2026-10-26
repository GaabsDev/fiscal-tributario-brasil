// Package nfgas implementa a Nota Fiscal Eletrônica de Gás Canalizado (NFGas - Modelo 76).
// Layout oficial 1.00 com medição volumétrica, pressão/temperatura e Reforma Tributária (NT 2026.002 RTC 1.01).
package nfgas

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloNFGas       = "76"
	VersaoLayoutNFGas = "1.00"
)

// NFGas representa a Nota Fiscal Eletrônica de Gás Canalizado (Modelo 76).
type NFGas struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/nfgas NFGas" json:"-"`
	InfNFGas  InfNFGas       `xml:"infNFGas" json:"infNFGas"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfNFGas struct {
	Versao string `xml:"versao,attr" json:"versao"`
	Id     string `xml:"Id,attr" json:"id"`
	Ide    Ide    `xml:"ide" json:"ide"`
	Emit   Emit   `xml:"emit" json:"emit"`
	Dest   Dest   `xml:"dest" json:"dest"`
	GMed   []GMed `xml:"gMed,omitempty" json:"gMed,omitempty"`
	Det    []Det  `xml:"det" json:"det"`
	Total  Total  `xml:"total" json:"total"`
}

type Ide struct {
	CUF    string       `xml:"cUF" json:"cUF"`
	TpAmb  string       `xml:"tpAmb" json:"tpAmb"`
	Mod    string       `xml:"mod" json:"mod"`
	Serie  string       `xml:"serie" json:"serie"`
	NNF    string       `xml:"nNF" json:"nNF"`
	CNF    string       `xml:"cNF" json:"cNF"`
	DhEmi  nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	TpEmis string       `xml:"tpEmis" json:"tpEmis"`
}

type Emit struct {
	CNPJ      nfe.CNPJ      `xml:"CNPJ" json:"CNPJ"`
	IE        string        `xml:"IE" json:"IE"`
	XNome     string        `xml:"xNome" json:"xNome"`
	EnderEmit nfe.EnderEmit `xml:"enderEmit" json:"enderEmit"`
}

type Dest struct {
	XNome string    `xml:"xNome" json:"xNome"`
	CNPJ  *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
}

type GMed struct {
	NMed    string      `xml:"nMed" json:"nMed"`
	DMedAnt nfe.Date    `xml:"dMedAnt" json:"dMedAnt"`
	DMedAct nfe.Date    `xml:"dMedAct" json:"dMedAct"`
	ConsM3  nfe.Decimal `xml:"consM3" json:"consM3"`
	FatorPT nfe.Decimal `xml:"fatorPT" json:"fatorPT"`
}

type Det struct {
	NItem int  `xml:"nItem,attr" json:"nItem"`
	Prod  Prod `xml:"prod" json:"prod"`
	Imp   Imp  `xml:"imposto" json:"imposto"`
}

type Prod struct {
	CProd string      `xml:"cProd" json:"cProd"`
	XProd string      `xml:"xProd" json:"xProd"`
	UCom  string      `xml:"uCom" json:"uCom"`
	QCom  nfe.Decimal `xml:"qCom" json:"qCom"`
	VProd nfe.Decimal `xml:"vProd" json:"vProd"`
}

type Imp struct {
	IBSCBS *IBSCBSNFGas `xml:"IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type IBSCBSNFGas struct {
	CST        string         `xml:"CST" json:"CST"`
	CClassTrib nfe.CClassTrib `xml:"cClassTrib" json:"cClassTrib"`
	VBC        nfe.Decimal    `xml:"vBC" json:"vBC"`
	PIBS       nfe.Decimal    `xml:"pIBS" json:"pIBS"`
	VIBS       nfe.Decimal    `xml:"vIBS" json:"vIBS"`
	PCBS       nfe.Decimal    `xml:"pCBS" json:"pCBS"`
	VCBS       nfe.Decimal    `xml:"vCBS" json:"vCBS"`
}

type Total struct {
	VProd nfe.Decimal `xml:"vProd" json:"vProd"`
	VNF   nfe.Decimal `xml:"vNF" json:"vNF"`
}

// NewNFGas cria uma nova NFGas (Modelo 76).
func NewNFGas(chave nfe.ChaveAcesso) *NFGas {
	id := fmt.Sprintf("NFGas%s", string(chave))
	return &NFGas{
		InfNFGas: InfNFGas{
			Versao: VersaoLayoutNFGas,
			Id:     id,
			Ide: Ide{
				Mod: ModeloNFGas,
			},
			Det: make([]Det, 0),
		},
	}
}

func (doc *NFGas) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *NFGas) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
