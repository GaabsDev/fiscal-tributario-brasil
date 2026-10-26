// Package nfse implementa a Nota Fiscal de Serviços Eletrônica de Padrão Nacional (NFS-e) e a Declaração de Prestação de Serviços (DPS).
// Layout oficial 1.01 do Sistema Nacional da NFS-e (SEFAZ/RFB/Comitê Gestor da NFS-e).
package nfse

import (
	"encoding/xml"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	VersaoLayoutNFSe = "1.01"
	VersaoLayoutDPS  = "1.01"
)

// NFSe é o elemento raiz da Nota Fiscal de Serviços Eletrônica Padrão Nacional.
type NFSe struct {
	XMLName   xml.Name       `xml:"http://www.sped.fazenda.gov.br/nfse NFSe" json:"-"`
	InfNFSe   InfNFSe        `xml:"infNFSe" json:"infNFSe"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfNFSe struct {
	Id            string       `xml:"Id,attr" json:"id"`
	Versao        string       `xml:"versao,attr" json:"versao"`
	XLocEmi       string       `xml:"xLocEmi" json:"xLocEmi"`
	XLocPrestacao string       `xml:"xLocPrestacao" json:"xLocPrestacao"`
	NNFSe         string       `xml:"nNFSe" json:"nNFSe"`
	CDFS          string       `xml:"cDFS" json:"cDFS"`
	DhEmi         nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	DPS           DPS          `xml:"DPS" json:"DPS"`
	Valores       ValoresNFSe  `xml:"valores" json:"valores"`
}

// DPS representa a Declaração de Prestação de Serviços do emitente.
type DPS struct {
	XMLName xml.Name `xml:"http://www.sped.fazenda.gov.br/nfse DPS" json:"-"`
	InfDPS  InfDPS   `xml:"infDPS" json:"infDPS"`
}

type InfDPS struct {
	Id      string       `xml:"Id,attr" json:"id"`
	Versao  string       `xml:"versao,attr" json:"versao"`
	DhEmi   nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	DCompet nfe.Date     `xml:"dCompet" json:"dCompet"`
	Prest   Prestador    `xml:"prest" json:"prest"`
	Toma    Tomador      `xml:"toma" json:"toma"`
	Serv    Servico      `xml:"serv" json:"serv"`
	Valores ValoresDPS   `xml:"valores" json:"valores"`
}

type Prestador struct {
	CNPJ nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	IM   string   `xml:"IM,omitempty" json:"IM,omitempty"`
}

type Tomador struct {
	CNPJ  *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	XNome string    `xml:"xNome" json:"xNome"`
}

type Servico struct {
	CLocPrestacao string `xml:"cLocPrestacao" json:"cLocPrestacao"`
	CServ         CServ  `xml:"cServ" json:"cServ"`
}

type CServ struct {
	CTribNac   string         `xml:"cTribNac" json:"cTribNac"`
	CClassTrib nfe.CClassTrib `xml:"cClassTrib,omitempty" json:"cClassTrib,omitempty"`
	XDescServ  string         `xml:"xDescServ" json:"xDescServ"`
}

type ValoresDPS struct {
	VServPrest VServPrest  `xml:"vServPrest" json:"vServPrest"`
	Trib       TributosDPS `xml:"trib" json:"trib"`
}

type VServPrest struct {
	VServ nfe.Decimal `xml:"vServ" json:"vServ"`
}

type TributosDPS struct {
	ISSQN  ISSQNDPS   `xml:"ISSQN" json:"ISSQN"`
	IBSCBS *IBSCBSDPS `xml:"IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type ISSQNDPS struct {
	CSTISSQN string      `xml:"cstISSQN" json:"cstISSQN"`
	VBCISSQN nfe.Decimal `xml:"vBCISSQN" json:"vBCISSQN"`
	PAliq    nfe.Decimal `xml:"pAliq" json:"pAliq"`
	VISSQN   nfe.Decimal `xml:"vISSQN" json:"vISSQN"`
}

type IBSCBSDPS struct {
	VBC  nfe.Decimal `xml:"vBC" json:"vBC"`
	VIBS nfe.Decimal `xml:"vIBS" json:"vIBS"`
	VCBS nfe.Decimal `xml:"vCBS" json:"vCBS"`
}

type ValoresNFSe struct {
	VLiq nfe.Decimal `xml:"vLiq" json:"vLiq"`
}

// NewDPS inicializa uma nova Declaração de Prestação de Serviços.
func NewDPS(id string) *DPS {
	return &DPS{
		InfDPS: InfDPS{
			Id:     id,
			Versao: VersaoLayoutDPS,
		},
	}
}

// NewNFSe inicializa uma nova NFS-e Nacional.
func NewNFSe(id string) *NFSe {
	return &NFSe{
		InfNFSe: InfNFSe{
			Id:     id,
			Versao: VersaoLayoutNFSe,
		},
	}
}

func (doc *NFSe) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *NFSe) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}

func (doc *DPS) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *DPS) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
