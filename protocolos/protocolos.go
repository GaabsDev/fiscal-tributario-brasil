// Package protocolos implementa os envelopes e contratos oficiais de distribuição, eventos e protocolos SEFAZ.
// Inclui procNFe, procCTe, procCTeOS, procGTVe, procMDFe, retEnviNFe, retConsSitNFe, retConsStatServ e TEvento.
package protocolos

import (
	"encoding/xml"

	"github.com/GaabsDev/fiscal-tributario-brasil/cte"
	"github.com/GaabsDev/fiscal-tributario-brasil/mdfe"
	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

// ProcNFe representa o XML completo distribuído ao destinatário (NF-e + ProtNFe) conforme procNFe_v4.00.xsd.
type ProcNFe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe nfeProc" json:"-"`
	Versao  string   `xml:"versao,attr" json:"versao"`
	NFe     nfe.NFe  `xml:"NFe" json:"NFe"`
	ProtNFe ProtNFe  `xml:"protNFe" json:"protNFe"`
}

// ProtNFe contém os dados do protocolo de autorização de uso da NF-e.
type ProtNFe struct {
	Versao  string     `xml:"versao,attr" json:"versao"`
	InfProt InfProtNFe `xml:"infProt" json:"infProt"`
}

type InfProtNFe struct {
	TpAmb    string       `xml:"tpAmb" json:"tpAmb"`
	VerAplic string       `xml:"verAplic" json:"verAplic"`
	ChNFe    string       `xml:"chNFe" json:"chNFe"`
	DhRecbto nfe.DateTime `xml:"dhRecbto" json:"dhRecbto"`
	NProt    string       `xml:"nProt" json:"nProt"`
	DigVal   string       `xml:"digVal" json:"digVal"`
	CStat    string       `xml:"cStat" json:"cStat"`
	XMotivo  string       `xml:"xMotivo" json:"xMotivo"`
}

// ProcCTe representa o documento de distribuição do CT-e (procCTe_v4.00.xsd).
type ProcCTe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/cte cteProc" json:"-"`
	Versao  string   `xml:"versao,attr" json:"versao"`
	CTe     cte.CTe  `xml:"CTe" json:"CTe"`
	ProtCTe ProtCTe  `xml:"protCTe" json:"protCTe"`
}

// ProcCTeOS representa o documento de distribuição do CT-e OS (procCTeOS_v4.00.xsd).
type ProcCTeOS struct {
	XMLName xml.Name  `xml:"http://www.portalfiscal.inf.br/cte cteOSProc" json:"-"`
	Versao  string    `xml:"versao,attr" json:"versao"`
	CTeOS   cte.CTeOS `xml:"CTeOS" json:"CTeOS"`
	ProtCTe ProtCTe   `xml:"protCTe" json:"protCTe"`
}

// ProcGTVe representa o documento de distribuição da GTV-e (procGTVe_v4.00.xsd).
type ProcGTVe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/cte gtveProc" json:"-"`
	Versao  string   `xml:"versao,attr" json:"versao"`
	GTVe    cte.GTVe `xml:"GTVe" json:"GTVe"`
	ProtCTe ProtCTe  `xml:"protCTe" json:"protCTe"`
}

type ProtCTe struct {
	Versao  string     `xml:"versao,attr" json:"versao"`
	InfProt InfProtCTe `xml:"infProt" json:"infProt"`
}

type InfProtCTe struct {
	TpAmb    string       `xml:"tpAmb" json:"tpAmb"`
	VerAplic string       `xml:"verAplic" json:"verAplic"`
	ChCTe    string       `xml:"chCTe" json:"chCTe"`
	DhRecbto nfe.DateTime `xml:"dhRecbto" json:"dhRecbto"`
	NProt    string       `xml:"nProt" json:"nProt"`
	CStat    string       `xml:"cStat" json:"cStat"`
	XMotivo  string       `xml:"xMotivo" json:"xMotivo"`
}

// ProcMDFe representa o documento de distribuição do MDF-e (procMDFe_v3.00.xsd).
type ProcMDFe struct {
	XMLName  xml.Name  `xml:"http://www.portalfiscal.inf.br/mdfe mdfeProc" json:"-"`
	Versao   string    `xml:"versao,attr" json:"versao"`
	MDFe     mdfe.MDFe `xml:"MDFe" json:"MDFe"`
	ProtMDFe ProtMDFe  `xml:"protMDFe" json:"protMDFe"`
}

type ProtMDFe struct {
	Versao  string      `xml:"versao,attr" json:"versao"`
	InfProt InfProtMDFe `xml:"infProt" json:"infProt"`
}

type InfProtMDFe struct {
	TpAmb    string       `xml:"tpAmb" json:"tpAmb"`
	VerAplic string       `xml:"verAplic" json:"verAplic"`
	ChMDFe   string       `xml:"chMDFe" json:"chMDFe"`
	DhRecbto nfe.DateTime `xml:"dhRecbto" json:"dhRecbto"`
	NProt    string       `xml:"nProt" json:"nProt"`
	CStat    string       `xml:"cStat" json:"cStat"`
	XMotivo  string       `xml:"xMotivo" json:"xMotivo"`
}

// RetEnviNFe resultado do processamento síncrono ou assíncrono do lote de NF-e.
type RetEnviNFe struct {
	XMLName  xml.Name `xml:"http://www.portalfiscal.inf.br/nfe retEnviNFe" json:"-"`
	Versao   string   `xml:"versao,attr" json:"versao"`
	TpAmb    string   `xml:"tpAmb" json:"tpAmb"`
	VerAplic string   `xml:"verAplic" json:"verAplic"`
	CStat    string   `xml:"cStat" json:"cStat"`
	XMotivo  string   `xml:"xMotivo" json:"xMotivo"`
	ProtNFe  *ProtNFe `xml:"protNFe,omitempty" json:"protNFe,omitempty"`
}

// RetConsStatServ resultado da consulta ao status de serviço da SEFAZ.
type RetConsStatServ struct {
	XMLName  xml.Name     `xml:"http://www.portalfiscal.inf.br/nfe retConsStatServ" json:"-"`
	Versao   string       `xml:"versao,attr" json:"versao"`
	TpAmb    string       `xml:"tpAmb" json:"tpAmb"`
	VerAplic string       `xml:"verAplic" json:"verAplic"`
	CStat    string       `xml:"cStat" json:"cStat"`
	XMotivo  string       `xml:"xMotivo" json:"xMotivo"`
	CUF      string       `xml:"cUF" json:"cUF"`
	DhRecbto nfe.DateTime `xml:"dhRecbto" json:"dhRecbto"`
	TMed     int          `xml:"tMed" json:"tMed"`
}

// TEvento envelope para eventos fiscais (Cancelamento, CC-e, Manifestação do Destinatário, etc.).
type TEvento struct {
	XMLName   xml.Name       `xml:"http://www.portalfiscal.inf.br/nfe evento" json:"-"`
	Versao    string         `xml:"versao,attr" json:"versao"`
	InfEvento InfEvento      `xml:"infEvento" json:"infEvento"`
	Signature *nfe.Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfEvento struct {
	Id         string       `xml:"Id,attr" json:"id"`
	COrgao     string       `xml:"cOrgao" json:"cOrgao"`
	TpAmb      string       `xml:"tpAmb" json:"tpAmb"`
	CNPJ       string       `xml:"CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF        string       `xml:"CPF,omitempty" json:"CPF,omitempty"`
	ChNFe      string       `xml:"chNFe" json:"chNFe"`
	DhEvento   nfe.DateTime `xml:"dhEvento" json:"dhEvento"`
	TpEvento   string       `xml:"tpEvento" json:"tpEvento"`
	NSeqEvento int          `xml:"nSeqEvento" json:"nSeqEvento"`
	VerEvento  string       `xml:"verEvento" json:"verEvento"`
	DetEvento  DetEvento    `xml:"detEvento" json:"detEvento"`
}

type DetEvento struct {
	Versao     string `xml:"versao,attr" json:"versao"`
	DescEvento string `xml:"descEvento" json:"descEvento"`
	XCorrecao  string `xml:"xCorrecao,omitempty" json:"xCorrecao,omitempty"`
	XJust      string `xml:"xJust,omitempty" json:"xJust,omitempty"`
	NProt      string `xml:"nProt,omitempty" json:"nProt,omitempty"`
}
