// Package nfcom implementa a Nota Fiscal Fatura de Serviços de Comunicação (NFCom - Modelo 62).
// Layout oficial 1.00 com regras completas de telecomunicações e Reforma Tributária (NT 2026.002 RTC 1.01).
package nfcom

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloNFCom       = "62"
	VersaoLayoutNFCom = "1.00"
)

// NFCom representa a Nota Fiscal Fatura de Serviços de Comunicação (Modelo 62).
type NFCom struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/nfcom NFCom" json:"-"`
	InfNFCom  InfNFCom       `xml:"infNFCom" json:"infNFCom"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfNFCom struct {
	Versao    string    `xml:"versao,attr" json:"versao"`
	Id        string    `xml:"Id,attr" json:"id"`
	Ide       Ide       `xml:"ide" json:"ide"`
	Emit      Emit      `xml:"emit" json:"emit"`
	Dest      Dest      `xml:"dest" json:"dest"`
	Assinante Assinante `xml:"assinante" json:"assinante"`
	Det       []Det     `xml:"det" json:"det"`
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

type Assinante struct {
	IContrato string `xml:"iContrato" json:"iContrato"`
	TpAssin   string `xml:"tpAssin" json:"tpAssin"` // 1-Comercial, 2-Industrial, 3-Residencial, etc.
}

type Det struct {
	NItem int  `xml:"nItem,attr" json:"nItem"`
	Prod  Prod `xml:"prod" json:"prod"`
	Imp   Imp  `xml:"imposto" json:"imposto"`
}

type Prod struct {
	CProd  string      `xml:"cProd" json:"cProd"`
	XProd  string      `xml:"xProd" json:"xProd"`
	CFOP   string      `xml:"CFOP" json:"CFOP"`
	UCom   string      `xml:"uCom" json:"uCom"`
	QCom   nfe.Decimal `xml:"qCom" json:"qCom"`
	VUnCom nfe.Decimal `xml:"vUnCom" json:"vUnCom"`
	VProd  nfe.Decimal `xml:"vProd" json:"vProd"`
}

type Imp struct {
	ICMS00 *ICMS00      `xml:"ICMS00,omitempty" json:"ICMS00,omitempty"`
	IBSCBS *IBSCBSNFCom `xml:"IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type ICMS00 struct {
	CST   string      `xml:"CST" json:"CST"`
	VBC   nfe.Decimal `xml:"vBC" json:"vBC"`
	PICMS nfe.Decimal `xml:"pICMS" json:"pICMS"`
	VICMS nfe.Decimal `xml:"vICMS" json:"vICMS"`
}

type IBSCBSNFCom struct {
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

// NewNFCom cria uma nova NFCom (Modelo 62).
func NewNFCom(chave nfe.ChaveAcesso) *NFCom {
	id := fmt.Sprintf("NFCom%s", string(chave))
	return &NFCom{
		InfNFCom: InfNFCom{
			Versao: VersaoLayoutNFCom,
			Id:     id,
			Ide: Ide{
				Mod: ModeloNFCom,
			},
			Det: make([]Det, 0),
		},
	}
}

func (doc *NFCom) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *NFCom) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
