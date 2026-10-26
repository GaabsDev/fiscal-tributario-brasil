// Package nf3e implementa a Nota Fiscal de Energia Elétrica Eletrônica (NF3e - Modelo 66).
// Layout oficial 1.00 com suporte a medição de grandezas elétricas, bandeiras tarifárias e Reforma Tributária (NT 2026.004 RTC).
package nf3e

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloNF3e       = "66"
	VersaoLayoutNF3e = "1.00"
)

// NF3e representa a Nota Fiscal de Energia Elétrica Eletrônica (Modelo 66).
type NF3e struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/nf3e NF3e" json:"-"`
	InfNF3e   InfNF3e        `xml:"infNF3e" json:"infNF3e"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfNF3e struct {
	Versao    string    `xml:"versao,attr" json:"versao"`
	Id        string    `xml:"Id,attr" json:"id"`
	Ide       Ide       `xml:"ide" json:"ide"`
	Emit      Emit      `xml:"emit" json:"emit"`
	Dest      Dest      `xml:"dest" json:"dest"`
	Acessante Acessante `xml:"acessante" json:"acessante"`
	GMed      []GMed    `xml:"gMed,omitempty" json:"gMed,omitempty"`
	Total     Total     `xml:"total" json:"total"`
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

type Acessante struct {
	IdAcesso  string `xml:"idAcesso" json:"idAcesso"`
	CodClasse string `xml:"codClasse" json:"codClasse"` // 01-Residencial, 02-Comercial, 03-Industrial, etc.
}

type GMed struct {
	NMed    string      `xml:"nMed" json:"nMed"`
	DMedAnt nfe.Date    `xml:"dMedAnt" json:"dMedAnt"`
	DMedAct nfe.Date    `xml:"dMedAct" json:"dMedAct"`
	ConsKWh nfe.Decimal `xml:"consKWh" json:"consKWh"`
}

type Total struct {
	VProd  nfe.Decimal `xml:"vProd" json:"vProd"`
	VNF    nfe.Decimal `xml:"vNF" json:"vNF"`
	IBSCBS *IBSCBSNF3e `xml:"IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type IBSCBSNF3e struct {
	CST        string         `xml:"CST" json:"CST"`
	CClassTrib nfe.CClassTrib `xml:"cClassTrib" json:"cClassTrib"`
	VBC        nfe.Decimal    `xml:"vBC" json:"vBC"`
	VIBS       nfe.Decimal    `xml:"vIBS" json:"vIBS"`
	VCBS       nfe.Decimal    `xml:"vCBS" json:"vCBS"`
}

// NewNF3e cria uma nova NF3e (Modelo 66).
func NewNF3e(chave nfe.ChaveAcesso) *NF3e {
	id := fmt.Sprintf("NF3e%s", string(chave))
	return &NF3e{
		InfNF3e: InfNF3e{
			Versao: VersaoLayoutNF3e,
			Id:     id,
			Ide: Ide{
				Mod: ModeloNF3e,
			},
		},
	}
}

func (doc *NF3e) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *NF3e) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
