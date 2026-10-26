// Package cte - GTV-e (Modelo 64)
// Layout oficial conforme GTVe_v4.00.xsd da SEFAZ para Guia de Transporte de Valores Eletrônica.
package cte

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

// GTVe é o elemento raiz da Guia de Transporte de Valores Eletrônica (Modelo 64).
type GTVe struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/cte GTVe" json:"-"`
	InfCte    InfCteGTVe     `xml:"infCte" json:"infCte"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

// InfCteGTVe reúne as informações da GTV-e.
type InfCteGTVe struct {
	Versao  string  `xml:"versao,attr" json:"versao"`
	Id      string  `xml:"Id,attr" json:"id"`
	Ide     IdeGTVe `xml:"ide" json:"ide"`
	Compl   *Compl  `xml:"compl,omitempty" json:"compl,omitempty"`
	Emit    Emit    `xml:"emit" json:"emit"`
	Rem     *Rem    `xml:"rem,omitempty" json:"rem,omitempty"`
	Dest    *Dest   `xml:"dest,omitempty" json:"dest,omitempty"`
	InfGTVe InfGTVe `xml:"infGTVe" json:"infGTVe"`
}

type IdeGTVe struct {
	CUF     string       `xml:"cUF" json:"cUF"`
	CCT     string       `xml:"cCT" json:"cCT"`
	CFOP    string       `xml:"CFOP" json:"CFOP"`
	NatOp   string       `xml:"natOp" json:"natOp"`
	Mod     string       `xml:"mod" json:"mod"`
	Serie   string       `xml:"serie" json:"serie"`
	NCT     string       `xml:"nCT" json:"nCT"`
	DhEmi   nfe.DateTime `xml:"dhEmi" json:"dhEmi"`
	TpImp   string       `xml:"tpImp" json:"tpImp"`
	TpEmis  string       `xml:"tpEmis" json:"tpEmis"`
	CDV     string       `xml:"cDV" json:"cDV"`
	TpAmb   string       `xml:"tpAmb" json:"tpAmb"`
	TpCTe   string       `xml:"tpCTe" json:"tpCTe"`
	ProcEmi string       `xml:"procEmi" json:"procEmi"`
	VerProc string       `xml:"verProc" json:"verProc"`
}

// InfGTVe detalha as informações específicas do transporte de valores.
type InfGTVe struct {
	InfEspecie []InfEspecie `xml:"infEspecie" json:"infEspecie"`
	RemOrig    *RemOrig     `xml:"remOrig,omitempty" json:"remOrig,omitempty"`
	DestFim    *DestFim     `xml:"destFim,omitempty" json:"destFim,omitempty"`
}

type InfEspecie struct {
	TpEspecie string      `xml:"tpEspecie" json:"tpEspecie"` // 1-Numerário, 2-Cheque, 3-Moeda, 4-Outros
	VItem     nfe.Decimal `xml:"vItem" json:"vItem"`
}

type RemOrig struct {
	CNPJ  *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	XNome string    `xml:"xNome" json:"xNome"`
}

type DestFim struct {
	CNPJ  *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	XNome string    `xml:"xNome" json:"xNome"`
}

// NewGTVe inicializa uma nova GTV-e (Modelo 64).
func NewGTVe(chave nfe.ChaveAcesso) *GTVe {
	id := fmt.Sprintf("CTe%s", string(chave))
	return &GTVe{
		InfCte: InfCteGTVe{
			Versao: VersaoLayoutCTe,
			Id:     id,
			Ide: IdeGTVe{
				Mod:   ModeloGTVe,
				TpCTe: "0",
			},
		},
	}
}

func (doc *GTVe) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *GTVe) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
