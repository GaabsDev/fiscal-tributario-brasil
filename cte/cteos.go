// Package cte - CT-e OS (Modelo 67)
// Layout oficial conforme cteOS_v4.00.xsd da SEFAZ para transporte de passageiros, valores e excesso de bagagem.
package cte

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

// CTeOS é o elemento raiz do Conhecimento de Transporte Eletrônico para Outros Serviços (Modelo 67).
type CTeOS struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/cte CTeOS" json:"-"`
	InfCte    InfCteOS       `xml:"infCte" json:"infCte"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

// InfCteOS reúne as informações do CT-e OS.
type InfCteOS struct {
	Versao     string       `xml:"versao,attr" json:"versao"`
	Id         string       `xml:"Id,attr" json:"id"`
	Ide        IdeOS        `xml:"ide" json:"ide"`
	Compl      *Compl       `xml:"compl,omitempty" json:"compl,omitempty"`
	Emit       Emit         `xml:"emit" json:"emit"`
	Toma       TomaOS       `xml:"toma" json:"toma"`
	VPrest     VPrest       `xml:"vPrest" json:"vPrest"`
	Imp        Imp          `xml:"imp" json:"imp"`
	InfCTeNorm InfCTeNormOS `xml:"infCTeNorm" json:"infCTeNorm"`
}

type IdeOS struct {
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
	CMunIni string       `xml:"cMunIni" json:"cMunIni"`
	XMunIni string       `xml:"xMunIni" json:"xMunIni"`
	UFIni   nfe.UF       `xml:"UFIni" json:"UFIni"`
	CMunFim string       `xml:"cMunFim" json:"cMunFim"`
	XMunFim string       `xml:"xMunFim" json:"xMunFim"`
	UFFim   nfe.UF       `xml:"UFFim" json:"UFFim"`
	TpServ  string       `xml:"tpServ" json:"tpServ"` // 6-Transporte de Pessoas, 7-Valores, 8-Excesso Bagagem
}

type TomaOS struct {
	CNPJ      *nfe.CNPJ `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *nfe.CPF  `xml:"CPF,omitempty" json:"CPF,omitempty"`
	IE        string    `xml:"IE,omitempty" json:"IE,omitempty"`
	XNome     string    `xml:"xNome" json:"xNome"`
	EnderToma EnderCTe  `xml:"enderToma" json:"enderToma"`
}

type InfCTeNormOS struct {
	InfServico InfServico `xml:"infServico" json:"infServico"`
	InfModal   InfModalOS `xml:"infModal" json:"infModal"`
}

type InfServico struct {
	XDescServ string `xml:"xDescServ" json:"xDescServ"`
	QPassag   int    `xml:"qPassag,omitempty" json:"qPassag,omitempty"`
}

// InfModalOS modal rodoviário para outros serviços (cteModalRodoviarioOS_v4.00.xsd).
type InfModalOS struct {
	VersaoModal string  `xml:"versaoModal,attr" json:"versaoModal"`
	RodoOS      *RodoOS `xml:"rodoOS,omitempty" json:"rodoOS,omitempty"`
}

type RodoOS struct {
	TAF            string  `xml:"TAF,omitempty" json:"TAF,omitempty"`
	NroRegEstadual string  `xml:"NroRegEstadual,omitempty" json:"NroRegEstadual,omitempty"`
	Veic           *VeicOS `xml:"veic,omitempty" json:"veic,omitempty"`
}

type VeicOS struct {
	Placa   string `xml:"placa" json:"placa"`
	RENAVAM string `xml:"RENAVAM,omitempty" json:"RENAVAM,omitempty"`
	UF      nfe.UF `xml:"UF" json:"UF"`
}

// NewCTeOS inicializa um novo CT-e OS (Modelo 67).
func NewCTeOS(chave nfe.ChaveAcesso) *CTeOS {
	id := fmt.Sprintf("CTe%s", string(chave))
	return &CTeOS{
		InfCte: InfCteOS{
			Versao: VersaoLayoutCTe,
			Id:     id,
			Ide: IdeOS{
				Mod:   ModeloCTeOS,
				TpCTe: "0",
			},
		},
	}
}

func (doc *CTeOS) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

func (doc *CTeOS) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
