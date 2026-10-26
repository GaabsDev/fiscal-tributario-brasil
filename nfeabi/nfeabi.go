// Package nfeabi implementa a Nota Fiscal Eletrônica de Alienação de Bens Imóveis (NF-e ABI - Modelo 77).
// Layout oficial 1.00 com matrícula cartorária, identificação imobiliária e regras da Reforma Tributária (NT 2026.002 RTC).
package nfeabi

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloNFeABI       = "77"
	VersaoLayoutNFeABI = "1.00"
)

// NFeABI representa a NF-e de Alienação de Bens Imóveis (Modelo 77).
type NFeABI struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/nfeabi NFeABI" json:"-"`
	InfNFeABI InfNFeABI      `xml:"infNFeABI" json:"infNFeABI"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfNFeABI struct {
	Versao   string     `xml:"versao,attr" json:"versao"`
	Id       string     `xml:"Id,attr" json:"id"`
	Ide      Ide        `xml:"ide" json:"ide"`
	Emit     Emit       `xml:"emit" json:"emit"`
	Adquir   Adquirente `xml:"adquir" json:"adquir"`
	Imovel   Imovel     `xml:"imovel" json:"imovel"`
	Operacao Operacao   `xml:"operacao" json:"operacao"`
	Total    Total      `xml:"total" json:"total"`
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
	CNPJ      *nfe.CNPJ     `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *nfe.CPF      `xml:"CPF,omitempty" json:"CPF,omitempty"`
	XNome     string        `xml:"xNome" json:"xNome"`
	EnderEmit nfe.EnderEmit `xml:"enderEmit" json:"enderEmit"`
}

type Adquirente struct {
	CNPJ  *nfe.CNPJ     `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *nfe.CPF      `xml:"CPF,omitempty" json:"CPF,omitempty"`
	XNome string        `xml:"xNome" json:"xNome"`
	Ender nfe.EnderEmit `xml:"ender" json:"ender"`
}

type Imovel struct {
	Matricula string `xml:"matricula" json:"matricula"`
	Cartorio  string `xml:"cartorio" json:"cartorio"`
	Rip       string `xml:"rip,omitempty" json:"rip,omitempty"`
	CIB       string `xml:"CIB,omitempty" json:"CIB,omitempty"` // Cadastro Imobiliário Brasileiro
	XLgr      string `xml:"xLgr" json:"xLgr"`
	Nro       string `xml:"nro" json:"nro"`
	XBairro   string `xml:"xBairro" json:"xBairro"`
	CMun      string `xml:"cMun" json:"cMun"`
	UF        nfe.UF `xml:"UF" json:"UF"`
}

type Operacao struct {
	TpOperacao string        `xml:"tpOperacao" json:"tpOperacao"` // 1-Venda, 2-Permuta, 3-Cessão, etc.
	VOperacao  nfe.Decimal   `xml:"vOperacao" json:"vOperacao"`
	IBSCBS     *IBSCBSNFeABI `xml:"IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type IBSCBSNFeABI struct {
	CST        string         `xml:"CST" json:"CST"`
	CClassTrib nfe.CClassTrib `xml:"cClassTrib" json:"cClassTrib"`
	VBC        nfe.Decimal    `xml:"vBC" json:"vBC"`
	PIBS       nfe.Decimal    `xml:"pIBS" json:"pIBS"`
	VIBS       nfe.Decimal    `xml:"vIBS" json:"vIBS"`
	PCBS       nfe.Decimal    `xml:"pCBS" json:"pCBS"`
	VCBS       nfe.Decimal    `xml:"vCBS" json:"vCBS"`
}

type Total struct {
	VNF nfe.Decimal `xml:"vNF" json:"vNF"`
}

// NewNFeABI cria uma nova NF-e ABI (Modelo 77).
func NewNFeABI(chave nfe.ChaveAcesso) *NFeABI {
	id := fmt.Sprintf("NFeABI%s", string(chave))
	return &NFeABI{
		InfNFeABI: InfNFeABI{
			Versao: VersaoLayoutNFeABI,
			Id:     id,
			Ide: Ide{
				Mod: ModeloNFeABI,
			},
		},
	}
}

func (doc *NFeABI) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *NFeABI) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
