// Package nfeouro implementa a Nota Fiscal Eletrônica do Ouro Ativo Financeiro (NF-e Ouro) e Declaração de Aquisição de Ouro (DAO).
// Layout oficial 1.00 da Receita Federal do Brasil (RFB) e Banco Central do Brasil (BACEN) para operações com ouro como ativo financeiro / cambial.
package nfeouro

import (
	"encoding/xml"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	VersaoLayoutNFeOuro = "1.00"
	VersaoLayoutDAO     = "1.00"
)

// NFeOuro representa a Nota Fiscal Eletrônica do Ouro Ativo Financeiro.
type NFeOuro struct {
	XMLName   xml.Name       `xml:"http://www.receita.fazenda.gov.br/nfeouro NFeOuro" json:"-"`
	InfOuro   InfOuro        `xml:"infOuro" json:"infOuro"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfOuro struct {
	Id      string      `xml:"Id,attr" json:"id"`
	Versao  string      `xml:"versao,attr" json:"versao"`
	Ide     Ide         `xml:"ide" json:"ide"`
	InstFin InstFin     `xml:"instFin" json:"instFin"`
	Cliente Cliente     `xml:"cliente" json:"cliente"`
	Ouro    OuroDet     `xml:"ouro" json:"ouro"`
	Valores ValoresOuro `xml:"valores" json:"valores"`
}

type Ide struct {
	CUF    string       `xml:"cUF" json:"cUF"`
	NNF    string       `xml:"nNF" json:"nNF"`
	Serie  string       `xml:"serie" json:"serie"`
	DhEmi  nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	TpOper string       `xml:"tpOper" json:"tpOper"` // 1-Compra, 2-Venda
}

type InstFin struct {
	CNPJ   nfe.CNPJ `xml:"CNPJ" json:"CNPJ"`
	XNome  string   `xml:"xNome" json:"xNome"`
	CBacen string   `xml:"cBacen" json:"cBacen"`
}

type Cliente struct {
	CNPJ  *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	XNome string    `xml:"xNome" json:"xNome"`
}

type OuroDet struct {
	Teor    nfe.Decimal `xml:"teor" json:"teor"`       // Pureza (ex: 99.9)
	PesoBr  nfe.Decimal `xml:"pesoBr" json:"pesoBr"`   // Gramas brutas
	PesoLiq nfe.Decimal `xml:"pesoLiq" json:"pesoLiq"` // Gramas finas
	Lacre   string      `xml:"lacre,omitempty" json:"lacre,omitempty"`
}

type ValoresOuro struct {
	VGram  nfe.Decimal `xml:"vGram" json:"vGram"`
	VTotal nfe.Decimal `xml:"vTotal" json:"vTotal"`
	IOF    nfe.Decimal `xml:"IOF" json:"IOF"`
}

// DAO Declaração de Aquisição de Ouro.
type DAO struct {
	XMLName   xml.Name       `xml:"http://www.receita.fazenda.gov.br/dao DAO" json:"-"`
	InfDAO    InfDAO         `xml:"infDAO" json:"infDAO"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfDAO struct {
	Id         string       `xml:"Id,attr" json:"id"`
	Versao     string       `xml:"versao,attr" json:"versao"`
	DhEmi      nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	Garimpeiro Garimpeiro   `xml:"garimpeiro" json:"garimpeiro"`
	Origem     OrigemOuro   `xml:"origem" json:"origem"`
	Peso       nfe.Decimal  `xml:"peso" json:"peso"`
}

type Garimpeiro struct {
	CPF         nfe.CPF `xml:"CPF" json:"CPF"`
	XNome       string  `xml:"xNome" json:"xNome"`
	PermGarimpo string  `xml:"permGarimpo" json:"permGarimpo"`
}

type OrigemOuro struct {
	UFCons   nfe.UF `xml:"UFCons" json:"UFCons"`
	CMunCons string `xml:"cMunCons" json:"cMunCons"`
}

// NewNFeOuro inicializa uma nova NF-e Ouro.
func NewNFeOuro(id string) *NFeOuro {
	return &NFeOuro{
		InfOuro: InfOuro{
			Id:     id,
			Versao: VersaoLayoutNFeOuro,
		},
	}
}

// NewDAO inicializa uma nova DAO.
func NewDAO(id string) *DAO {
	return &DAO{
		InfDAO: InfDAO{
			Id:     id,
			Versao: VersaoLayoutDAO,
		},
	}
}

func (doc *NFeOuro) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *NFeOuro) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}

func (doc *DAO) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *DAO) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
