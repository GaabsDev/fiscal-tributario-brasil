// Package nfsevia implementa a Nota Fiscal de Serviços de Exploração de Vias e Pedágio (NFS-e Via).
// Layout oficial 1.01 para concessionárias de rodovias conforme diretrizes nacionais e Reforma Tributária (NT 2026.002 RTC).
package nfsevia

import (
	"encoding/xml"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	VersaoLayoutNFSeVia = "1.01"
)

// NFSeVia representa a Nota Fiscal Eletrônica de Exploração de Via (Pedágio).
type NFSeVia struct {
	XMLName   xml.Name       `xml:"http://www.sped.fazenda.gov.br/nfsevia NFSeVia" json:"-"`
	InfVia    InfVia         `xml:"infVia" json:"infVia"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfVia struct {
	Id      string         `xml:"Id,attr" json:"id"`
	Versao  string         `xml:"versao,attr" json:"versao"`
	DhEmi   nfe.DateTime   `xml:"dhEmi" json:"dhEmi"`
	Concess Concessionaria `xml:"concess" json:"concess"`
	Praca   Praca          `xml:"praca" json:"praca"`
	Passag  Passagem       `xml:"passag" json:"passag"`
	Valores ValoresVia     `xml:"valores" json:"valores"`
}

type Concessionaria struct {
	CNPJ  nfe.CNPJ `xml:"CNPJ" json:"CNPJ"`
	XNome string   `xml:"xNome" json:"xNome"`
}

type Praca struct {
	CPraca string `xml:"cPraca" json:"cPraca"`
	XPraca string `xml:"xPraca" json:"xPraca"`
	KM     string `xml:"km" json:"km"`
	Rod    string `xml:"rod" json:"rod"`
}

type Passagem struct {
	Sentido string `xml:"sentido" json:"sentido"` // 1-Norte/Sul, etc.
	Placa   string `xml:"placa,omitempty" json:"placa,omitempty"`
	CatVeic string `xml:"catVeic" json:"catVeic"`
}

type ValoresVia struct {
	VTarifa nfe.Decimal `xml:"vTarifa" json:"vTarifa"`
	IBSCBS  *IBSCBSVia  `xml:"IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type IBSCBSVia struct {
	CST        string         `xml:"CST" json:"CST"`
	CClassTrib nfe.CClassTrib `xml:"cClassTrib" json:"cClassTrib"`
	VBC        nfe.Decimal    `xml:"vBC" json:"vBC"`
	PIBS       nfe.Decimal    `xml:"pIBS" json:"pIBS"`
	VIBS       nfe.Decimal    `xml:"vIBS" json:"vIBS"`
	PCBS       nfe.Decimal    `xml:"pCBS" json:"pCBS"`
	VCBS       nfe.Decimal    `xml:"vCBS" json:"vCBS"`
}

// NewNFSeVia cria uma nova NFS-e Via.
func NewNFSeVia(id string) *NFSeVia {
	return &NFSeVia{
		InfVia: InfVia{
			Id:     id,
			Versao: VersaoLayoutNFSeVia,
		},
	}
}

func (doc *NFSeVia) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *NFSeVia) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
