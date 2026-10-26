// Package dce implementa a Declaração de Conteúdo Eletrônica (DC-e - Modelo 99).
// Layout oficial 1.00 com NT 2024.001 da SEFAZ para transporte de bens e encomendas sem finalidade comercial ou por não contribuintes.
package dce

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloDCe       = "99"
	VersaoLayoutDCe = "1.00"
)

// DCe representa a Declaração de Conteúdo Eletrônica (Modelo 99).
type DCe struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/dce DCe" json:"-"`
	InfDCe    InfDCe         `xml:"infDCe" json:"infDCe"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfDCe struct {
	Versao string  `xml:"versao,attr" json:"versao"`
	Id     string  `xml:"Id,attr" json:"id"`
	Ide    Ide     `xml:"ide" json:"ide"`
	Emit   Emit    `xml:"emit" json:"emit"`
	Dest   Dest    `xml:"dest" json:"dest"`
	Det    []Det   `xml:"det" json:"det"`
	Total  Total   `xml:"total" json:"total"`
	Transp *Transp `xml:"transp,omitempty" json:"transp,omitempty"`
}

type Ide struct {
	CUF    string       `xml:"cUF" json:"cUF"`
	TpAmb  string       `xml:"tpAmb" json:"tpAmb"`
	Mod    string       `xml:"mod" json:"mod"`
	Serie  string       `xml:"serie" json:"serie"`
	NDCe   string       `xml:"nDCe" json:"nDCe"`
	CDCe   string       `xml:"cDCe" json:"cDCe"`
	DhEmi  nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	TpEmis string       `xml:"tpEmis" json:"tpEmis"`
}

type Emit struct {
	CNPJ      *nfe.CNPJ     `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *nfe.CPF      `xml:"CPF,omitempty" json:"CPF,omitempty"`
	XNome     string        `xml:"xNome" json:"xNome"`
	EnderEmit nfe.EnderEmit `xml:"enderEmit" json:"enderEmit"`
}

type Dest struct {
	CNPJ      *nfe.CNPJ     `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *nfe.CPF      `xml:"CPF,omitempty" json:"CPF,omitempty"`
	XNome     string        `xml:"xNome" json:"xNome"`
	EnderDest nfe.EnderEmit `xml:"enderDest" json:"enderDest"`
}

type Det struct {
	NItem int  `xml:"nItem,attr" json:"nItem"`
	Prod  Prod `xml:"prod" json:"prod"`
}

type Prod struct {
	CProd string      `xml:"cProd" json:"cProd"`
	XProd string      `xml:"xProd" json:"xProd"`
	NCM   string      `xml:"NCM,omitempty" json:"NCM,omitempty"`
	QCom  nfe.Decimal `xml:"qCom" json:"qCom"`
	VProd nfe.Decimal `xml:"vProd" json:"vProd"`
}

type Total struct {
	VDCe nfe.Decimal `xml:"vDCe" json:"vDCe"`
}

type Transp struct {
	ModFrete string `xml:"modFrete" json:"modFrete"`
}

// NewDCe cria uma nova Declaração de Conteúdo Eletrônica (Modelo 99).
func NewDCe(chave nfe.ChaveAcesso) *DCe {
	id := fmt.Sprintf("DCe%s", string(chave))
	return &DCe{
		InfDCe: InfDCe{
			Versao: VersaoLayoutDCe,
			Id:     id,
			Ide: Ide{
				Mod: ModeloDCe,
			},
			Det: make([]Det, 0),
		},
	}
}

func (doc *DCe) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *DCe) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
