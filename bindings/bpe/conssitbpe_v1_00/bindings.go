// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package conssitbpe_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TALCZFMCBS struct {
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vTribRegCBS" json:"vTribRegCBS"`
}

type TALCZFMCBSNFe struct {
	TpALCZFMCBS     string        `xml:"http://www.portalfiscal.inf.br/bpe tpALCZFMCBS" json:"tpALCZFMCBS"`
	NProcSuframa    *string       `xml:"http://www.portalfiscal.inf.br/bpe nProcSuframa,omitempty" json:"nProcSuframa,omitempty"`
	PAliqEfetRegCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vTribRegCBS" json:"vTribRegCBS"`
}

type TAjusteCompet struct {
	CompetApur TCompetApur `xml:"http://www.portalfiscal.inf.br/bpe competApur" json:"competApur"`
	VIBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBS" json:"vIBS"`
	VCBS       TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBS" json:"vCBS"`
}

type TAmb string

type TBPeInfBPeIde struct {
	CUF        TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/bpe cUF" json:"cUF"`
	TpAmb      TAmb                `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	Mod        TModBPe             `xml:"http://www.portalfiscal.inf.br/bpe mod" json:"mod"`
	Serie      string              `xml:"http://www.portalfiscal.inf.br/bpe serie" json:"serie"`
	NBP        TNF                 `xml:"http://www.portalfiscal.inf.br/bpe nBP" json:"nBP"`
	CBP        string              `xml:"http://www.portalfiscal.inf.br/bpe cBP" json:"cBP"`
	CDV        string              `xml:"http://www.portalfiscal.inf.br/bpe cDV" json:"cDV"`
	Modal      string              `xml:"http://www.portalfiscal.inf.br/bpe modal" json:"modal"`
	DhEmi      TDateTimeUTC        `xml:"http://www.portalfiscal.inf.br/bpe dhEmi" json:"dhEmi"`
	TpEmis     string              `xml:"http://www.portalfiscal.inf.br/bpe tpEmis" json:"tpEmis"`
	VerProc    string              `xml:"http://www.portalfiscal.inf.br/bpe verProc" json:"verProc"`
	TpBPe      string              `xml:"http://www.portalfiscal.inf.br/bpe tpBPe" json:"tpBPe"`
	IndPres    TIndPres            `xml:"http://www.portalfiscal.inf.br/bpe indPres" json:"indPres"`
	UFIni      TUFSemEX            `xml:"http://www.portalfiscal.inf.br/bpe UFIni" json:"UFIni"`
	CMunIni    TCodMunIBGE         `xml:"http://www.portalfiscal.inf.br/bpe cMunIni" json:"cMunIni"`
	UFFim      TUF                 `xml:"http://www.portalfiscal.inf.br/bpe UFFim" json:"UFFim"`
	CMunFim    TCodMunIBGE         `xml:"http://www.portalfiscal.inf.br/bpe cMunFim" json:"cMunFim"`
	DhCont     *TDateTimeUTC       `xml:"http://www.portalfiscal.inf.br/bpe dhCont,omitempty" json:"dhCont,omitempty"`
	XJust      *string             `xml:"http://www.portalfiscal.inf.br/bpe xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov *TCompraGovReduzido `xml:"http://www.portalfiscal.inf.br/bpe gCompraGov,omitempty" json:"gCompraGov,omitempty"`
}

type TBPeInfBPeEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/bpe CNPJ" json:"CNPJ"`
	IE        string   `xml:"http://www.portalfiscal.inf.br/bpe IE" json:"IE"`
	IEST      *string  `xml:"http://www.portalfiscal.inf.br/bpe IEST,omitempty" json:"IEST,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/bpe xFant,omitempty" json:"xFant,omitempty"`
	IM        *string  `xml:"http://www.portalfiscal.inf.br/bpe IM,omitempty" json:"IM,omitempty"`
	CNAE      *string  `xml:"http://www.portalfiscal.inf.br/bpe CNAE,omitempty" json:"CNAE,omitempty"`
	CRT       string   `xml:"http://www.portalfiscal.inf.br/bpe CRT" json:"CRT"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/bpe enderEmit" json:"enderEmit"`
	TAR       *string  `xml:"http://www.portalfiscal.inf.br/bpe TAR,omitempty" json:"TAR,omitempty"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/bpe ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TBPeInfBPeComp struct {
	XNome         string    `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	CNPJ          *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/bpe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF           *TCPF     `xml:"http://www.portalfiscal.inf.br/bpe CPF,omitempty" json:"CPF,omitempty"`
	IDEstrangeiro *string   `xml:"http://www.portalfiscal.inf.br/bpe idEstrangeiro,omitempty" json:"idEstrangeiro,omitempty"`
	IE            *string   `xml:"http://www.portalfiscal.inf.br/bpe IE,omitempty" json:"IE,omitempty"`
	EnderComp     TEndereco `xml:"http://www.portalfiscal.inf.br/bpe enderComp" json:"enderComp"`
}

type TBPeInfBPeAgencia struct {
	XNome        string    `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	CNPJ         TCNPJOpc  `xml:"http://www.portalfiscal.inf.br/bpe CNPJ" json:"CNPJ"`
	EnderAgencia TEndereco `xml:"http://www.portalfiscal.inf.br/bpe enderAgencia" json:"enderAgencia"`
}

type TBPeInfBPeInfBPeSub struct {
	ChBPe TChDFe            `xml:"http://www.portalfiscal.inf.br/bpe chBPe" json:"chBPe"`
	TpSub TTipoSubstituicao `xml:"http://www.portalfiscal.inf.br/bpe tpSub" json:"tpSub"`
}

type TBPeInfBPeInfPassagemInfPassageiro struct {
	XNome string  `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	CPF   *TCPF   `xml:"http://www.portalfiscal.inf.br/bpe CPF,omitempty" json:"CPF,omitempty"`
	TpDoc TDoc    `xml:"http://www.portalfiscal.inf.br/bpe tpDoc" json:"tpDoc"`
	NDoc  string  `xml:"http://www.portalfiscal.inf.br/bpe nDoc" json:"nDoc"`
	XDoc  *string `xml:"http://www.portalfiscal.inf.br/bpe xDoc,omitempty" json:"xDoc,omitempty"`
	DNasc *string `xml:"http://www.portalfiscal.inf.br/bpe dNasc,omitempty" json:"dNasc,omitempty"`
	Fone  *string `xml:"http://www.portalfiscal.inf.br/bpe fone,omitempty" json:"fone,omitempty"`
	Email *TEmail `xml:"http://www.portalfiscal.inf.br/bpe email,omitempty" json:"email,omitempty"`
}

type TBPeInfBPeInfPassagem struct {
	CLocOrig      string                              `xml:"http://www.portalfiscal.inf.br/bpe cLocOrig" json:"cLocOrig"`
	XLocOrig      string                              `xml:"http://www.portalfiscal.inf.br/bpe xLocOrig" json:"xLocOrig"`
	CLocDest      string                              `xml:"http://www.portalfiscal.inf.br/bpe cLocDest" json:"cLocDest"`
	XLocDest      string                              `xml:"http://www.portalfiscal.inf.br/bpe xLocDest" json:"xLocDest"`
	DhEmb         TDateTimeUTC                        `xml:"http://www.portalfiscal.inf.br/bpe dhEmb" json:"dhEmb"`
	DhValidade    TDateTimeUTC                        `xml:"http://www.portalfiscal.inf.br/bpe dhValidade" json:"dhValidade"`
	InfPassageiro *TBPeInfBPeInfPassagemInfPassageiro `xml:"http://www.portalfiscal.inf.br/bpe infPassageiro,omitempty" json:"infPassageiro,omitempty"`
}

type TBPeInfBPeInfViagemInfTravessia struct {
	TpVeiculo  string `xml:"http://www.portalfiscal.inf.br/bpe tpVeiculo" json:"tpVeiculo"`
	SitVeiculo string `xml:"http://www.portalfiscal.inf.br/bpe sitVeiculo" json:"sitVeiculo"`
}

type TBPeInfBPeInfViagem struct {
	CPercurso    string                           `xml:"http://www.portalfiscal.inf.br/bpe cPercurso" json:"cPercurso"`
	XPercurso    string                           `xml:"http://www.portalfiscal.inf.br/bpe xPercurso" json:"xPercurso"`
	TpViagem     string                           `xml:"http://www.portalfiscal.inf.br/bpe tpViagem" json:"tpViagem"`
	TpServ       string                           `xml:"http://www.portalfiscal.inf.br/bpe tpServ" json:"tpServ"`
	TpAcomodacao string                           `xml:"http://www.portalfiscal.inf.br/bpe tpAcomodacao" json:"tpAcomodacao"`
	TpTrecho     string                           `xml:"http://www.portalfiscal.inf.br/bpe tpTrecho" json:"tpTrecho"`
	DhViagem     TDateTimeUTC                     `xml:"http://www.portalfiscal.inf.br/bpe dhViagem" json:"dhViagem"`
	DhConexao    *TDateTimeUTC                    `xml:"http://www.portalfiscal.inf.br/bpe dhConexao,omitempty" json:"dhConexao,omitempty"`
	Prefixo      *string                          `xml:"http://www.portalfiscal.inf.br/bpe prefixo,omitempty" json:"prefixo,omitempty"`
	Poltrona     *string                          `xml:"http://www.portalfiscal.inf.br/bpe poltrona,omitempty" json:"poltrona,omitempty"`
	Plataforma   *string                          `xml:"http://www.portalfiscal.inf.br/bpe plataforma,omitempty" json:"plataforma,omitempty"`
	InfTravessia *TBPeInfBPeInfViagemInfTravessia `xml:"http://www.portalfiscal.inf.br/bpe infTravessia,omitempty" json:"infTravessia,omitempty"`
}

type TBPeInfBPeInfValorBPeComp struct {
	TpComp string   `xml:"http://www.portalfiscal.inf.br/bpe tpComp" json:"tpComp"`
	VComp  TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vComp" json:"vComp"`
}

type TBPeInfBPeInfValorBPe struct {
	VBP        TDec1302                    `xml:"http://www.portalfiscal.inf.br/bpe vBP" json:"vBP"`
	VBPLiq     *TDec1302                   `xml:"http://www.portalfiscal.inf.br/bpe vBPLiq,omitempty" json:"vBPLiq,omitempty"`
	VDesconto  TDec1302                    `xml:"http://www.portalfiscal.inf.br/bpe vDesconto" json:"vDesconto"`
	VPgto      TDec1302                    `xml:"http://www.portalfiscal.inf.br/bpe vPgto" json:"vPgto"`
	VTroco     TDec1302                    `xml:"http://www.portalfiscal.inf.br/bpe vTroco" json:"vTroco"`
	TpDesconto *string                     `xml:"http://www.portalfiscal.inf.br/bpe tpDesconto,omitempty" json:"tpDesconto,omitempty"`
	XDesconto  *string                     `xml:"http://www.portalfiscal.inf.br/bpe xDesconto,omitempty" json:"xDesconto,omitempty"`
	CDesconto  *string                     `xml:"http://www.portalfiscal.inf.br/bpe cDesconto,omitempty" json:"cDesconto,omitempty"`
	Comp       []TBPeInfBPeInfValorBPeComp `xml:"http://www.portalfiscal.inf.br/bpe Comp,omitempty" json:"Comp,omitempty"`
}

type TBPeInfBPeImpICMSUFFim struct {
	VBCUFFim   TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vBCUFFim" json:"vBCUFFim"`
	PFCPUFFim  TDec0302 `xml:"http://www.portalfiscal.inf.br/bpe pFCPUFFim" json:"pFCPUFFim"`
	PICMSUFFim TDec0302 `xml:"http://www.portalfiscal.inf.br/bpe pICMSUFFim" json:"pICMSUFFim"`
	PICMSInter TDec0302 `xml:"http://www.portalfiscal.inf.br/bpe pICMSInter" json:"pICMSInter"`
	VFCPUFFim  TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vFCPUFFim" json:"vFCPUFFim"`
	VICMSUFFim TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vICMSUFFim" json:"vICMSUFFim"`
	VICMSUFIni TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vICMSUFIni" json:"vICMSUFIni"`
}

type TBPeInfBPeImp struct {
	ICMS       TImp                    `xml:"http://www.portalfiscal.inf.br/bpe ICMS" json:"ICMS"`
	VTotTrib   *TDec1302               `xml:"http://www.portalfiscal.inf.br/bpe vTotTrib,omitempty" json:"vTotTrib,omitempty"`
	InfAdFisco *string                 `xml:"http://www.portalfiscal.inf.br/bpe infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	ICMSUFFim  *TBPeInfBPeImpICMSUFFim `xml:"http://www.portalfiscal.inf.br/bpe ICMSUFFim,omitempty" json:"ICMSUFFim,omitempty"`
	IBSCBS     *TTribBPe               `xml:"http://www.portalfiscal.inf.br/bpe IBSCBS,omitempty" json:"IBSCBS,omitempty"`
	VTotDFe    *TDec1302               `xml:"http://www.portalfiscal.inf.br/bpe vTotDFe,omitempty" json:"vTotDFe,omitempty"`
}

type TBPeInfBPePagCard struct {
	TpIntegra string  `xml:"http://www.portalfiscal.inf.br/bpe tpIntegra" json:"tpIntegra"`
	CNPJ      *TCNPJ  `xml:"http://www.portalfiscal.inf.br/bpe CNPJ,omitempty" json:"CNPJ,omitempty"`
	TBand     *string `xml:"http://www.portalfiscal.inf.br/bpe tBand,omitempty" json:"tBand,omitempty"`
	XBand     *string `xml:"http://www.portalfiscal.inf.br/bpe xBand,omitempty" json:"xBand,omitempty"`
	CAut      *string `xml:"http://www.portalfiscal.inf.br/bpe cAut,omitempty" json:"cAut,omitempty"`
	NsuTrans  *string `xml:"http://www.portalfiscal.inf.br/bpe nsuTrans,omitempty" json:"nsuTrans,omitempty"`
	NsuHost   *string `xml:"http://www.portalfiscal.inf.br/bpe nsuHost,omitempty" json:"nsuHost,omitempty"`
	NParcelas *string `xml:"http://www.portalfiscal.inf.br/bpe nParcelas,omitempty" json:"nParcelas,omitempty"`
	InfAdCard *string `xml:"http://www.portalfiscal.inf.br/bpe infAdCard,omitempty" json:"infAdCard,omitempty"`
}

type TBPeInfBPePag struct {
	TPag    string             `xml:"http://www.portalfiscal.inf.br/bpe tPag" json:"tPag"`
	XPag    *string            `xml:"http://www.portalfiscal.inf.br/bpe xPag,omitempty" json:"xPag,omitempty"`
	NDocPag *string            `xml:"http://www.portalfiscal.inf.br/bpe nDocPag,omitempty" json:"nDocPag,omitempty"`
	VPag    TDec1302           `xml:"http://www.portalfiscal.inf.br/bpe vPag" json:"vPag"`
	Card    *TBPeInfBPePagCard `xml:"http://www.portalfiscal.inf.br/bpe card,omitempty" json:"card,omitempty"`
}

type TBPeInfBPePgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/bpe pgto,omitempty" json:"pgto,omitempty"`
}

type TBPeInfBPeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/bpe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/bpe CPF,omitempty" json:"CPF,omitempty"`
}

type TBPeInfBPeInfAdic struct {
	InfAdFisco *string `xml:"http://www.portalfiscal.inf.br/bpe infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     *string `xml:"http://www.portalfiscal.inf.br/bpe infCpl,omitempty" json:"infCpl,omitempty"`
}

type TBPeInfBPe struct {
	Versao      TVerBPe               `xml:"versao,attr" json:"versao"`
	ID          string                `xml:"Id,attr" json:"Id"`
	Ide         TBPeInfBPeIde         `xml:"http://www.portalfiscal.inf.br/bpe ide" json:"ide"`
	Emit        TBPeInfBPeEmit        `xml:"http://www.portalfiscal.inf.br/bpe emit" json:"emit"`
	Comp        *TBPeInfBPeComp       `xml:"http://www.portalfiscal.inf.br/bpe comp,omitempty" json:"comp,omitempty"`
	Agencia     *TBPeInfBPeAgencia    `xml:"http://www.portalfiscal.inf.br/bpe agencia,omitempty" json:"agencia,omitempty"`
	InfBPeSub   *TBPeInfBPeInfBPeSub  `xml:"http://www.portalfiscal.inf.br/bpe infBPeSub,omitempty" json:"infBPeSub,omitempty"`
	InfPassagem TBPeInfBPeInfPassagem `xml:"http://www.portalfiscal.inf.br/bpe infPassagem" json:"infPassagem"`
	InfViagem   []TBPeInfBPeInfViagem `xml:"http://www.portalfiscal.inf.br/bpe infViagem,omitempty" json:"infViagem,omitempty"`
	InfValorBPe TBPeInfBPeInfValorBPe `xml:"http://www.portalfiscal.inf.br/bpe infValorBPe" json:"infValorBPe"`
	Imp         TBPeInfBPeImp         `xml:"http://www.portalfiscal.inf.br/bpe imp" json:"imp"`
	Pag         []TBPeInfBPePag       `xml:"http://www.portalfiscal.inf.br/bpe pag,omitempty" json:"pag,omitempty"`
	PgtoVinc    *TBPeInfBPePgtoVinc   `xml:"http://www.portalfiscal.inf.br/bpe pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	AutXML      []TBPeInfBPeAutXML    `xml:"http://www.portalfiscal.inf.br/bpe autXML,omitempty" json:"autXML,omitempty"`
	InfAdic     *TBPeInfBPeInfAdic    `xml:"http://www.portalfiscal.inf.br/bpe infAdic,omitempty" json:"infAdic,omitempty"`
	InfRespTec  *TRespTec             `xml:"http://www.portalfiscal.inf.br/bpe infRespTec,omitempty" json:"infRespTec,omitempty"`
}

type TBPeInfBPeSupl struct {
	QRCodBPe     string  `xml:"http://www.portalfiscal.inf.br/bpe qrCodBPe" json:"qrCodBPe"`
	BoardPassBPe *string `xml:"http://www.portalfiscal.inf.br/bpe boardPassBPe,omitempty" json:"boardPassBPe,omitempty"`
}

type TBPe struct {
	InfBPe     TBPeInfBPe     `xml:"http://www.portalfiscal.inf.br/bpe infBPe" json:"infBPe"`
	InfBPeSupl TBPeInfBPeSupl `xml:"http://www.portalfiscal.inf.br/bpe infBPeSupl" json:"infBPeSupl"`
	Signature  SignatureType  `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TBPeTAInfBPeIde struct {
	CUF        TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/bpe cUF" json:"cUF"`
	TpAmb      TAmb                `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	Mod        TModBPe             `xml:"http://www.portalfiscal.inf.br/bpe mod" json:"mod"`
	Serie      string              `xml:"http://www.portalfiscal.inf.br/bpe serie" json:"serie"`
	NBP        TNF                 `xml:"http://www.portalfiscal.inf.br/bpe nBP" json:"nBP"`
	CBP        string              `xml:"http://www.portalfiscal.inf.br/bpe cBP" json:"cBP"`
	CDV        string              `xml:"http://www.portalfiscal.inf.br/bpe cDV" json:"cDV"`
	Modal      string              `xml:"http://www.portalfiscal.inf.br/bpe modal" json:"modal"`
	DhEmi      TDateTimeUTC        `xml:"http://www.portalfiscal.inf.br/bpe dhEmi" json:"dhEmi"`
	TpEmis     string              `xml:"http://www.portalfiscal.inf.br/bpe tpEmis" json:"tpEmis"`
	VerProc    string              `xml:"http://www.portalfiscal.inf.br/bpe verProc" json:"verProc"`
	TpBPe      string              `xml:"http://www.portalfiscal.inf.br/bpe tpBPe" json:"tpBPe"`
	TpCompra   string              `xml:"http://www.portalfiscal.inf.br/bpe tpCompra" json:"tpCompra"`
	IndPres    TIndPres            `xml:"http://www.portalfiscal.inf.br/bpe indPres" json:"indPres"`
	UFIni      TUFSemEX            `xml:"http://www.portalfiscal.inf.br/bpe UFIni" json:"UFIni"`
	CMunIni    TCodMunIBGE         `xml:"http://www.portalfiscal.inf.br/bpe cMunIni" json:"cMunIni"`
	UFFim      TUF                 `xml:"http://www.portalfiscal.inf.br/bpe UFFim" json:"UFFim"`
	CMunFim    TCodMunIBGE         `xml:"http://www.portalfiscal.inf.br/bpe cMunFim" json:"cMunFim"`
	DhCont     *TDateTimeUTC       `xml:"http://www.portalfiscal.inf.br/bpe dhCont,omitempty" json:"dhCont,omitempty"`
	XJust      *string             `xml:"http://www.portalfiscal.inf.br/bpe xJust,omitempty" json:"xJust,omitempty"`
	NroBPIATA  *string             `xml:"http://www.portalfiscal.inf.br/bpe nroBPIATA,omitempty" json:"nroBPIATA,omitempty"`
	GCompraGov *TCompraGovReduzido `xml:"http://www.portalfiscal.inf.br/bpe gCompraGov,omitempty" json:"gCompraGov,omitempty"`
}

type TBPeTAInfBPeEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/bpe CNPJ" json:"CNPJ"`
	IE        *string  `xml:"http://www.portalfiscal.inf.br/bpe IE,omitempty" json:"IE,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/bpe xFant,omitempty" json:"xFant,omitempty"`
	CRT       string   `xml:"http://www.portalfiscal.inf.br/bpe CRT" json:"CRT"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/bpe enderEmit" json:"enderEmit"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/bpe ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TBPeTAInfBPeComp struct {
	XNome         string     `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	CNPJ          *TCNPJOpc  `xml:"http://www.portalfiscal.inf.br/bpe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF           *TCPF      `xml:"http://www.portalfiscal.inf.br/bpe CPF,omitempty" json:"CPF,omitempty"`
	IDEstrangeiro *string    `xml:"http://www.portalfiscal.inf.br/bpe idEstrangeiro,omitempty" json:"idEstrangeiro,omitempty"`
	IE            *string    `xml:"http://www.portalfiscal.inf.br/bpe IE,omitempty" json:"IE,omitempty"`
	EnderComp     *TEndereco `xml:"http://www.portalfiscal.inf.br/bpe enderComp,omitempty" json:"enderComp,omitempty"`
}

type TBPeTAInfBPeAgencia struct {
	XNome        string    `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	CNPJ         TCNPJOpc  `xml:"http://www.portalfiscal.inf.br/bpe CNPJ" json:"CNPJ"`
	EnderAgencia TEndereco `xml:"http://www.portalfiscal.inf.br/bpe enderAgencia" json:"enderAgencia"`
}

type TBPeTAInfBPeInfBPeSub struct {
	ChBPe TChDFe              `xml:"http://www.portalfiscal.inf.br/bpe chBPe" json:"chBPe"`
	TpSub TTipoSubstituicaoTA `xml:"http://www.portalfiscal.inf.br/bpe tpSub" json:"tpSub"`
}

type TBPeTAInfBPeInfPassagemInfPassageiro struct {
	XNome string  `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	CPF   *TCPF   `xml:"http://www.portalfiscal.inf.br/bpe CPF,omitempty" json:"CPF,omitempty"`
	TpDoc TDoc    `xml:"http://www.portalfiscal.inf.br/bpe tpDoc" json:"tpDoc"`
	NDoc  string  `xml:"http://www.portalfiscal.inf.br/bpe nDoc" json:"nDoc"`
	XDoc  *string `xml:"http://www.portalfiscal.inf.br/bpe xDoc,omitempty" json:"xDoc,omitempty"`
	DNasc *string `xml:"http://www.portalfiscal.inf.br/bpe dNasc,omitempty" json:"dNasc,omitempty"`
	Fone  *string `xml:"http://www.portalfiscal.inf.br/bpe fone,omitempty" json:"fone,omitempty"`
	Email *TEmail `xml:"http://www.portalfiscal.inf.br/bpe email,omitempty" json:"email,omitempty"`
}

type TBPeTAInfBPeInfPassagem struct {
	DhEmb         TDateTimeUTC                          `xml:"http://www.portalfiscal.inf.br/bpe dhEmb" json:"dhEmb"`
	DhValidade    TDateTimeUTC                          `xml:"http://www.portalfiscal.inf.br/bpe dhValidade" json:"dhValidade"`
	InfPassageiro *TBPeTAInfBPeInfPassagemInfPassageiro `xml:"http://www.portalfiscal.inf.br/bpe infPassageiro,omitempty" json:"infPassageiro,omitempty"`
}

type TBPeTAInfBPeInfViagem struct {
	NroVoo          string        `xml:"http://www.portalfiscal.inf.br/bpe nroVoo" json:"nroVoo"`
	SiglaCiaOperVoo string        `xml:"http://www.portalfiscal.inf.br/bpe SiglaCiaOperVoo" json:"SiglaCiaOperVoo"`
	TpViagem        string        `xml:"http://www.portalfiscal.inf.br/bpe tpViagem" json:"tpViagem"`
	CAeroOrig       string        `xml:"http://www.portalfiscal.inf.br/bpe cAeroOrig" json:"cAeroOrig"`
	CAeroDest       string        `xml:"http://www.portalfiscal.inf.br/bpe cAeroDest" json:"cAeroDest"`
	TpServ          string        `xml:"http://www.portalfiscal.inf.br/bpe tpServ" json:"tpServ"`
	TpAcomodacao    string        `xml:"http://www.portalfiscal.inf.br/bpe tpAcomodacao" json:"tpAcomodacao"`
	TpTrecho        string        `xml:"http://www.portalfiscal.inf.br/bpe tpTrecho" json:"tpTrecho"`
	DhViagem        TDateTimeUTC  `xml:"http://www.portalfiscal.inf.br/bpe dhViagem" json:"dhViagem"`
	DhConexao       *TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/bpe dhConexao,omitempty" json:"dhConexao,omitempty"`
}

type TBPeTAInfBPeInfValorBPeComp struct {
	TpComp string   `xml:"http://www.portalfiscal.inf.br/bpe tpComp" json:"tpComp"`
	VComp  TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vComp" json:"vComp"`
}

type TBPeTAInfBPeInfValorBPe struct {
	VBP        TDec1302                      `xml:"http://www.portalfiscal.inf.br/bpe vBP" json:"vBP"`
	VBPLiq     *TDec1302                     `xml:"http://www.portalfiscal.inf.br/bpe vBPLiq,omitempty" json:"vBPLiq,omitempty"`
	VDesconto  TDec1302                      `xml:"http://www.portalfiscal.inf.br/bpe vDesconto" json:"vDesconto"`
	VPgto      TDec1302                      `xml:"http://www.portalfiscal.inf.br/bpe vPgto" json:"vPgto"`
	VTroco     TDec1302                      `xml:"http://www.portalfiscal.inf.br/bpe vTroco" json:"vTroco"`
	TpDesconto *string                       `xml:"http://www.portalfiscal.inf.br/bpe tpDesconto,omitempty" json:"tpDesconto,omitempty"`
	XDesconto  *string                       `xml:"http://www.portalfiscal.inf.br/bpe xDesconto,omitempty" json:"xDesconto,omitempty"`
	CDesconto  *string                       `xml:"http://www.portalfiscal.inf.br/bpe cDesconto,omitempty" json:"cDesconto,omitempty"`
	Comp       []TBPeTAInfBPeInfValorBPeComp `xml:"http://www.portalfiscal.inf.br/bpe Comp,omitempty" json:"Comp,omitempty"`
}

type TBPeTAInfBPeImp struct {
	IBSCBS     TTribBPe  `xml:"http://www.portalfiscal.inf.br/bpe IBSCBS" json:"IBSCBS"`
	VINSSRet   *TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vINSSRet,omitempty" json:"vINSSRet,omitempty"`
	VTotTrib   *TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vTotTrib,omitempty" json:"vTotTrib,omitempty"`
	InfAdFisco *string   `xml:"http://www.portalfiscal.inf.br/bpe infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	VTotDFe    TDec1302  `xml:"http://www.portalfiscal.inf.br/bpe vTotDFe" json:"vTotDFe"`
}

type TBPeTAInfBPePagCard struct {
	TpIntegra string  `xml:"http://www.portalfiscal.inf.br/bpe tpIntegra" json:"tpIntegra"`
	CNPJ      *TCNPJ  `xml:"http://www.portalfiscal.inf.br/bpe CNPJ,omitempty" json:"CNPJ,omitempty"`
	TBand     *string `xml:"http://www.portalfiscal.inf.br/bpe tBand,omitempty" json:"tBand,omitempty"`
	XBand     *string `xml:"http://www.portalfiscal.inf.br/bpe xBand,omitempty" json:"xBand,omitempty"`
	CAut      *string `xml:"http://www.portalfiscal.inf.br/bpe cAut,omitempty" json:"cAut,omitempty"`
	NsuTrans  *string `xml:"http://www.portalfiscal.inf.br/bpe nsuTrans,omitempty" json:"nsuTrans,omitempty"`
	NsuHost   *string `xml:"http://www.portalfiscal.inf.br/bpe nsuHost,omitempty" json:"nsuHost,omitempty"`
	NParcelas *string `xml:"http://www.portalfiscal.inf.br/bpe nParcelas,omitempty" json:"nParcelas,omitempty"`
	InfAdCard *string `xml:"http://www.portalfiscal.inf.br/bpe infAdCard,omitempty" json:"infAdCard,omitempty"`
}

type TBPeTAInfBPePag struct {
	TPag    string               `xml:"http://www.portalfiscal.inf.br/bpe tPag" json:"tPag"`
	XPag    *string              `xml:"http://www.portalfiscal.inf.br/bpe xPag,omitempty" json:"xPag,omitempty"`
	NDocPag *string              `xml:"http://www.portalfiscal.inf.br/bpe nDocPag,omitempty" json:"nDocPag,omitempty"`
	VPag    TDec1302             `xml:"http://www.portalfiscal.inf.br/bpe vPag" json:"vPag"`
	Card    *TBPeTAInfBPePagCard `xml:"http://www.portalfiscal.inf.br/bpe card,omitempty" json:"card,omitempty"`
}

type TBPeTAInfBPePgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/bpe pgto,omitempty" json:"pgto,omitempty"`
}

type TBPeTAInfBPeInfBPeVinculado struct {
	ChBPeMultiplo TChDFe `xml:"http://www.portalfiscal.inf.br/bpe chBPeMultiplo" json:"chBPeMultiplo"`
}

type TBPeTAInfBPeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/bpe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/bpe CPF,omitempty" json:"CPF,omitempty"`
}

type TBPeTAInfBPeInfAdic struct {
	InfAdFisco *string `xml:"http://www.portalfiscal.inf.br/bpe infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     *string `xml:"http://www.portalfiscal.inf.br/bpe infCpl,omitempty" json:"infCpl,omitempty"`
}

type TBPeTAInfBPe struct {
	Versao          TVerBPe                      `xml:"versao,attr" json:"versao"`
	ID              string                       `xml:"Id,attr" json:"Id"`
	Ide             TBPeTAInfBPeIde              `xml:"http://www.portalfiscal.inf.br/bpe ide" json:"ide"`
	Emit            TBPeTAInfBPeEmit             `xml:"http://www.portalfiscal.inf.br/bpe emit" json:"emit"`
	Comp            *TBPeTAInfBPeComp            `xml:"http://www.portalfiscal.inf.br/bpe comp,omitempty" json:"comp,omitempty"`
	Agencia         *TBPeTAInfBPeAgencia         `xml:"http://www.portalfiscal.inf.br/bpe agencia,omitempty" json:"agencia,omitempty"`
	InfBPeSub       *TBPeTAInfBPeInfBPeSub       `xml:"http://www.portalfiscal.inf.br/bpe infBPeSub,omitempty" json:"infBPeSub,omitempty"`
	InfPassagem     TBPeTAInfBPeInfPassagem      `xml:"http://www.portalfiscal.inf.br/bpe infPassagem" json:"infPassagem"`
	InfViagem       []TBPeTAInfBPeInfViagem      `xml:"http://www.portalfiscal.inf.br/bpe infViagem,omitempty" json:"infViagem,omitempty"`
	InfValorBPe     *TBPeTAInfBPeInfValorBPe     `xml:"http://www.portalfiscal.inf.br/bpe infValorBPe,omitempty" json:"infValorBPe,omitempty"`
	Imp             *TBPeTAInfBPeImp             `xml:"http://www.portalfiscal.inf.br/bpe imp,omitempty" json:"imp,omitempty"`
	Pag             []TBPeTAInfBPePag            `xml:"http://www.portalfiscal.inf.br/bpe pag,omitempty" json:"pag,omitempty"`
	PgtoVinc        *TBPeTAInfBPePgtoVinc        `xml:"http://www.portalfiscal.inf.br/bpe pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	InfBPeVinculado *TBPeTAInfBPeInfBPeVinculado `xml:"http://www.portalfiscal.inf.br/bpe infBPeVinculado,omitempty" json:"infBPeVinculado,omitempty"`
	AutXML          []TBPeTAInfBPeAutXML         `xml:"http://www.portalfiscal.inf.br/bpe autXML,omitempty" json:"autXML,omitempty"`
	InfAdic         *TBPeTAInfBPeInfAdic         `xml:"http://www.portalfiscal.inf.br/bpe infAdic,omitempty" json:"infAdic,omitempty"`
	InfRespTec      *TRespTec                    `xml:"http://www.portalfiscal.inf.br/bpe infRespTec,omitempty" json:"infRespTec,omitempty"`
}

type TBPeTAInfBPeSupl struct {
	QRCodBPe string `xml:"http://www.portalfiscal.inf.br/bpe qrCodBPe" json:"qrCodBPe"`
}

type TBPeTA struct {
	InfBPe     TBPeTAInfBPe     `xml:"http://www.portalfiscal.inf.br/bpe infBPe" json:"infBPe"`
	InfBPeSupl TBPeTAInfBPeSupl `xml:"http://www.portalfiscal.inf.br/bpe infBPeSupl" json:"infBPeSupl"`
	Signature  SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TBPeTMInfBPeIde struct {
	CUF        TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/bpe cUF" json:"cUF"`
	TpAmb      TAmb                `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	Mod        TModBPe             `xml:"http://www.portalfiscal.inf.br/bpe mod" json:"mod"`
	Serie      string              `xml:"http://www.portalfiscal.inf.br/bpe serie" json:"serie"`
	NBP        TNF                 `xml:"http://www.portalfiscal.inf.br/bpe nBP" json:"nBP"`
	CBP        string              `xml:"http://www.portalfiscal.inf.br/bpe cBP" json:"cBP"`
	CDV        string              `xml:"http://www.portalfiscal.inf.br/bpe cDV" json:"cDV"`
	Modal      string              `xml:"http://www.portalfiscal.inf.br/bpe modal" json:"modal"`
	DhEmi      TDateTimeUTC        `xml:"http://www.portalfiscal.inf.br/bpe dhEmi" json:"dhEmi"`
	DCompet    TData               `xml:"http://www.portalfiscal.inf.br/bpe dCompet" json:"dCompet"`
	TpEmis     string              `xml:"http://www.portalfiscal.inf.br/bpe tpEmis" json:"tpEmis"`
	VerProc    string              `xml:"http://www.portalfiscal.inf.br/bpe verProc" json:"verProc"`
	TpBPe      TTipoBPeTM          `xml:"http://www.portalfiscal.inf.br/bpe tpBPe" json:"tpBPe"`
	CFOP       TCFOP               `xml:"http://www.portalfiscal.inf.br/bpe CFOP" json:"CFOP"`
	DhCont     *TDateTimeUTC       `xml:"http://www.portalfiscal.inf.br/bpe dhCont,omitempty" json:"dhCont,omitempty"`
	XJust      *string             `xml:"http://www.portalfiscal.inf.br/bpe xJust,omitempty" json:"xJust,omitempty"`
	GCompraGov *TCompraGovReduzido `xml:"http://www.portalfiscal.inf.br/bpe gCompraGov,omitempty" json:"gCompraGov,omitempty"`
}

type TBPeTMInfBPeEmit struct {
	CNPJ      TCNPJ    `xml:"http://www.portalfiscal.inf.br/bpe CNPJ" json:"CNPJ"`
	IE        string   `xml:"http://www.portalfiscal.inf.br/bpe IE" json:"IE"`
	IEST      *string  `xml:"http://www.portalfiscal.inf.br/bpe IEST,omitempty" json:"IEST,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/bpe xFant,omitempty" json:"xFant,omitempty"`
	IM        *string  `xml:"http://www.portalfiscal.inf.br/bpe IM,omitempty" json:"IM,omitempty"`
	CNAE      *string  `xml:"http://www.portalfiscal.inf.br/bpe CNAE,omitempty" json:"CNAE,omitempty"`
	CRT       string   `xml:"http://www.portalfiscal.inf.br/bpe CRT" json:"CRT"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/bpe enderEmit" json:"enderEmit"`
	TAR       *string  `xml:"http://www.portalfiscal.inf.br/bpe TAR,omitempty" json:"TAR,omitempty"`
	ISUFEmit  *string  `xml:"http://www.portalfiscal.inf.br/bpe ISUFEmit,omitempty" json:"ISUFEmit,omitempty"`
}

type TBPeTMInfBPeDetBPeTMDetImp struct {
	ICMS       TImp      `xml:"http://www.portalfiscal.inf.br/bpe ICMS" json:"ICMS"`
	InfAdFisco *string   `xml:"http://www.portalfiscal.inf.br/bpe infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	IBSCBS     *TTribBPe `xml:"http://www.portalfiscal.inf.br/bpe IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type TBPeTMInfBPeDetBPeTMDetComp struct {
	XNome string `xml:"http://www.portalfiscal.inf.br/bpe xNome" json:"xNome"`
	QComp string `xml:"http://www.portalfiscal.inf.br/bpe qComp" json:"qComp"`
}

type TBPeTMInfBPeDetBPeTMDet struct {
	NViagem     string                        `xml:"nViagem,attr" json:"nViagem"`
	CMunIni     TCodMunIBGE                   `xml:"http://www.portalfiscal.inf.br/bpe cMunIni" json:"cMunIni"`
	CMunFim     *TCodMunIBGE                  `xml:"http://www.portalfiscal.inf.br/bpe cMunFim,omitempty" json:"cMunFim,omitempty"`
	NContInicio *string                       `xml:"http://www.portalfiscal.inf.br/bpe nContInicio,omitempty" json:"nContInicio,omitempty"`
	NContFim    *string                       `xml:"http://www.portalfiscal.inf.br/bpe nContFim,omitempty" json:"nContFim,omitempty"`
	QPass       string                        `xml:"http://www.portalfiscal.inf.br/bpe qPass" json:"qPass"`
	VBP         TDec1302                      `xml:"http://www.portalfiscal.inf.br/bpe vBP" json:"vBP"`
	VBPLiq      *TDec1302                     `xml:"http://www.portalfiscal.inf.br/bpe vBPLiq,omitempty" json:"vBPLiq,omitempty"`
	Imp         TBPeTMInfBPeDetBPeTMDetImp    `xml:"http://www.portalfiscal.inf.br/bpe imp" json:"imp"`
	Comp        []TBPeTMInfBPeDetBPeTMDetComp `xml:"http://www.portalfiscal.inf.br/bpe Comp,omitempty" json:"Comp,omitempty"`
}

type TBPeTMInfBPeDetBPeTM struct {
	IDEqpCont   string                    `xml:"idEqpCont,attr" json:"idEqpCont"`
	UFIniViagem TUFSemEX                  `xml:"http://www.portalfiscal.inf.br/bpe UFIniViagem" json:"UFIniViagem"`
	UFFimViagem *TUFSemEX                 `xml:"http://www.portalfiscal.inf.br/bpe UFFimViagem,omitempty" json:"UFFimViagem,omitempty"`
	Placa       *TPlaca                   `xml:"http://www.portalfiscal.inf.br/bpe placa,omitempty" json:"placa,omitempty"`
	Prefixo     *string                   `xml:"http://www.portalfiscal.inf.br/bpe prefixo,omitempty" json:"prefixo,omitempty"`
	Det         []TBPeTMInfBPeDetBPeTMDet `xml:"http://www.portalfiscal.inf.br/bpe det,omitempty" json:"det,omitempty"`
}

type TBPeTMInfBPeTotalICMSTot struct {
	VBC   TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vBC" json:"vBC"`
	VICMS TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vICMS" json:"vICMS"`
}

type TBPeTMInfBPeTotal struct {
	QPass     string                   `xml:"http://www.portalfiscal.inf.br/bpe qPass" json:"qPass"`
	VBP       TDec1302                 `xml:"http://www.portalfiscal.inf.br/bpe vBP" json:"vBP"`
	VBPLiq    *TDec1302                `xml:"http://www.portalfiscal.inf.br/bpe vBPLiq,omitempty" json:"vBPLiq,omitempty"`
	ICMSTot   TBPeTMInfBPeTotalICMSTot `xml:"http://www.portalfiscal.inf.br/bpe ICMSTot" json:"ICMSTot"`
	IBSCBSTot *TIBSCBSTot              `xml:"http://www.portalfiscal.inf.br/bpe IBSCBSTot,omitempty" json:"IBSCBSTot,omitempty"`
	VTotDFe   *TDec1302                `xml:"http://www.portalfiscal.inf.br/bpe vTotDFe,omitempty" json:"vTotDFe,omitempty"`
}

type TBPeTMInfBPePgtoVinc struct {
	Pgto []TPagamentoRTC `xml:"http://www.portalfiscal.inf.br/bpe pgto,omitempty" json:"pgto,omitempty"`
}

type TBPeTMInfBPeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/bpe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/bpe CPF,omitempty" json:"CPF,omitempty"`
}

type TBPeTMInfBPeInfAdic struct {
	InfAdFisco *string `xml:"http://www.portalfiscal.inf.br/bpe infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     *string `xml:"http://www.portalfiscal.inf.br/bpe infCpl,omitempty" json:"infCpl,omitempty"`
}

type TBPeTMInfBPe struct {
	Versao     TVerBPe                `xml:"versao,attr" json:"versao"`
	ID         string                 `xml:"Id,attr" json:"Id"`
	Ide        TBPeTMInfBPeIde        `xml:"http://www.portalfiscal.inf.br/bpe ide" json:"ide"`
	Emit       TBPeTMInfBPeEmit       `xml:"http://www.portalfiscal.inf.br/bpe emit" json:"emit"`
	DetBPeTM   []TBPeTMInfBPeDetBPeTM `xml:"http://www.portalfiscal.inf.br/bpe detBPeTM,omitempty" json:"detBPeTM,omitempty"`
	Total      TBPeTMInfBPeTotal      `xml:"http://www.portalfiscal.inf.br/bpe total" json:"total"`
	PgtoVinc   *TBPeTMInfBPePgtoVinc  `xml:"http://www.portalfiscal.inf.br/bpe pgtoVinc,omitempty" json:"pgtoVinc,omitempty"`
	AutXML     []TBPeTMInfBPeAutXML   `xml:"http://www.portalfiscal.inf.br/bpe autXML,omitempty" json:"autXML,omitempty"`
	InfAdic    *TBPeTMInfBPeInfAdic   `xml:"http://www.portalfiscal.inf.br/bpe infAdic,omitempty" json:"infAdic,omitempty"`
	InfRespTec *TRespTec              `xml:"http://www.portalfiscal.inf.br/bpe infRespTec,omitempty" json:"infRespTec,omitempty"`
}

type TBPeTM struct {
	InfBPe    TBPeTMInfBPe  `xml:"http://www.portalfiscal.inf.br/bpe infBPe" json:"infBPe"`
	Signature SignatureType `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCIBSGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/bpe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/bpe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/bpe gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSUF" json:"vIBSUF"`
}

type TCIBSGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/bpe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/bpe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/bpe gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSMun" json:"vIBSMun"`
}

type TCIBSGCBS struct {
	PCBS       TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pCBS" json:"pCBS"`
	GDif       *TDif         `xml:"http://www.portalfiscal.inf.br/bpe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib     `xml:"http://www.portalfiscal.inf.br/bpe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed         `xml:"http://www.portalfiscal.inf.br/bpe gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBS   `xml:"http://www.portalfiscal.inf.br/bpe gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vCBS" json:"vCBS"`
}

type TCIBS struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/bpe vBC" json:"vBC"`
	GIBSUF         TCIBSGIBSUF     `xml:"http://www.portalfiscal.inf.br/bpe gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSGIBSMun    `xml:"http://www.portalfiscal.inf.br/bpe gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/bpe vIBS" json:"vIBS"`
	GCBS           TCIBSGCBS       `xml:"http://www.portalfiscal.inf.br/bpe gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/bpe gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/bpe gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCIBSNFeGIBSUF struct {
	PIBSUF   TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pIBSUF" json:"pIBSUF"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/bpe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/bpe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/bpe gRed,omitempty" json:"gRed,omitempty"`
	VIBSUF   TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSUF" json:"vIBSUF"`
}

type TCIBSNFeGIBSMun struct {
	PIBSMun  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pIBSMun" json:"pIBSMun"`
	GDif     *TDif         `xml:"http://www.portalfiscal.inf.br/bpe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib *TDevTrib     `xml:"http://www.portalfiscal.inf.br/bpe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed     *TRed         `xml:"http://www.portalfiscal.inf.br/bpe gRed,omitempty" json:"gRed,omitempty"`
	VIBSMun  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSMun" json:"vIBSMun"`
}

type TCIBSNFeGCBS struct {
	PCBS       TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/bpe pCBS" json:"pCBS"`
	GDif       *TDif          `xml:"http://www.portalfiscal.inf.br/bpe gDif,omitempty" json:"gDif,omitempty"`
	GDevTrib   *TDevTrib      `xml:"http://www.portalfiscal.inf.br/bpe gDevTrib,omitempty" json:"gDevTrib,omitempty"`
	GRed       *TRed          `xml:"http://www.portalfiscal.inf.br/bpe gRed,omitempty" json:"gRed,omitempty"`
	GALCZFMCBS *TALCZFMCBSNFe `xml:"http://www.portalfiscal.inf.br/bpe gALCZFMCBS,omitempty" json:"gALCZFMCBS,omitempty"`
	VCBS       TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/bpe vCBS" json:"vCBS"`
}

type TCIBSNFe struct {
	VBC            TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/bpe vBC" json:"vBC"`
	GIBSUF         TCIBSNFeGIBSUF  `xml:"http://www.portalfiscal.inf.br/bpe gIBSUF" json:"gIBSUF"`
	GIBSMun        TCIBSNFeGIBSMun `xml:"http://www.portalfiscal.inf.br/bpe gIBSMun" json:"gIBSMun"`
	VIBS           TDec1302RTC     `xml:"http://www.portalfiscal.inf.br/bpe vIBS" json:"vIBS"`
	GCBS           TCIBSNFeGCBS    `xml:"http://www.portalfiscal.inf.br/bpe gCBS" json:"gCBS"`
	GTribRegular   *TTribRegular   `xml:"http://www.portalfiscal.inf.br/bpe gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TTribCompraGov `xml:"http://www.portalfiscal.inf.br/bpe gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCOrgaoIBGE string

type TCST string

type TCFOP string

type TChDFe string

type TChDFeRTC string

type TCNPJ string

type TCNPJBaseRTC string

type TCNPJOpc string

type TCNPJRTC string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

type TCompetApur string

type TCompraGov struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/bpe tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/bpe pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/bpe tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/bpe refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TCompraGovReduzido struct {
	TpEnteGov TEnteGov       `xml:"http://www.portalfiscal.inf.br/bpe tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204RTC  `xml:"http://www.portalfiscal.inf.br/bpe pRedutor" json:"pRedutor"`
	TpOperGov TOperCompraGov `xml:"http://www.portalfiscal.inf.br/bpe tpOperGov" json:"tpOperGov"`
	RefDFeAnt []TChDFeRTC    `xml:"http://www.portalfiscal.inf.br/bpe refDFeAnt,omitempty" json:"refDFeAnt,omitempty"`
}

type TConsSitBPe struct {
	Versao TVerConsSitBPe `xml:"versao,attr" json:"versao"`
	TpAmb  TAmb           `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	XServ  TServ          `xml:"http://www.portalfiscal.inf.br/bpe xServ" json:"xServ"`
	ChBPe  TChDFe         `xml:"http://www.portalfiscal.inf.br/bpe chBPe" json:"chBPe"`
}

type TCPF string

type TCPFVar string

type TCredPres struct {
	PCredPres        TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pCredPres" json:"pCredPres"`
	VCredPres        *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/bpe vCredPres,omitempty" json:"vCredPres,omitempty"`
	VCredPresCondSus *TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/bpe vCredPresCondSus,omitempty" json:"vCredPresCondSus,omitempty"`
}

type TCredPresIBSZFM struct {
	CompetApur       TCompetApur       `xml:"http://www.portalfiscal.inf.br/bpe competApur" json:"competApur"`
	TpCredPresIBSZFM TTpCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/bpe tpCredPresIBSZFM" json:"tpCredPresIBSZFM"`
	VCredPresIBSZFM  TDec1302RTC       `xml:"http://www.portalfiscal.inf.br/bpe vCredPresIBSZFM" json:"vCredPresIBSZFM"`
}

type TCredPresOper struct {
	VBCCredPres  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vBCCredPres" json:"vBCCredPres"`
	CCredPres    TcCredPres  `xml:"http://www.portalfiscal.inf.br/bpe cCredPres" json:"cCredPres"`
	GIBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/bpe gIBSCredPres,omitempty" json:"gIBSCredPres,omitempty"`
	GCBSCredPres *TCredPres  `xml:"http://www.portalfiscal.inf.br/bpe gCBSCredPres,omitempty" json:"gCBSCredPres,omitempty"`
}

type TData string

type TDateTimeUTC string

type TDec1104RTC string

type TDec1302RTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec030204RTC string

type TDec0303 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec0906Opc string

type TDec1104 string

type TDec1104OpRTC string

type TDec1104Opc string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec1302 string

type TDec1302Opc string

type TDevTrib struct {
	PDevTrib *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pDevTrib,omitempty" json:"pDevTrib,omitempty"`
	VDevTrib TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/bpe vDevTrib" json:"vDevTrib"`
}

type TDif struct {
	PDif TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pDif" json:"pDif"`
	VDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vDif" json:"vDif"`
}

type TDoc string

type TEmail string

type TEndeEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/bpe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/bpe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/bpe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/bpe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/bpe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/bpe xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/bpe CEP,omitempty" json:"CEP,omitempty"`
	UF      TUFSemEX    `xml:"http://www.portalfiscal.inf.br/bpe UF" json:"UF"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/bpe fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/bpe email,omitempty" json:"email,omitempty"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/bpe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/bpe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/bpe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/bpe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/bpe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/bpe xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/bpe CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/bpe UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/bpe cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/bpe xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/bpe fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/bpe email,omitempty" json:"email,omitempty"`
}

type TEnteGov string

type TEnviBPe struct {
	Versao TVerBPe `xml:"versao,attr" json:"versao"`
	IDLote TIDLote `xml:"http://www.portalfiscal.inf.br/bpe idLote" json:"idLote"`
	BPe    TBPe    `xml:"http://www.portalfiscal.inf.br/bpe BPe" json:"BPe"`
}

type TEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBSEstCred" json:"vCBSEstCred"`
}

type TEventoInfEventoDetEvento struct {
	VersaoEvento string       `xml:"versaoEvento,attr" json:"versaoEvento"`
	Any          []RawElement `xml:",any" json:"any,omitempty"`
}

type TEventoInfEvento struct {
	ID         string                    `xml:"Id,attr" json:"Id"`
	COrgao     TCOrgaoIBGE               `xml:"http://www.portalfiscal.inf.br/bpe cOrgao" json:"cOrgao"`
	TpAmb      TAmb                      `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	CNPJ       TCNPJ                     `xml:"http://www.portalfiscal.inf.br/bpe CNPJ" json:"CNPJ"`
	ChBPe      TChDFe                    `xml:"http://www.portalfiscal.inf.br/bpe chBPe" json:"chBPe"`
	DhEvento   TDateTimeUTC              `xml:"http://www.portalfiscal.inf.br/bpe dhEvento" json:"dhEvento"`
	TpEvento   string                    `xml:"http://www.portalfiscal.inf.br/bpe tpEvento" json:"tpEvento"`
	NSeqEvento string                    `xml:"http://www.portalfiscal.inf.br/bpe nSeqEvento" json:"nSeqEvento"`
	DetEvento  TEventoInfEventoDetEvento `xml:"http://www.portalfiscal.inf.br/bpe detEvento" json:"detEvento"`
}

type TEvento struct {
	Versao    TVerEvento       `xml:"versao,attr" json:"versao"`
	InfEvento TEventoInfEvento `xml:"http://www.portalfiscal.inf.br/bpe infEvento" json:"infEvento"`
	Signature SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TForPag string

type TIBSCBSMonoTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSUF" json:"vIBSUF"`
}

type TIBSCBSMonoTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSMun" json:"vIBSMun"`
}

type TIBSCBSMonoTotGIBS struct {
	GIBSUF           TIBSCBSMonoTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/bpe gIBSUF" json:"gIBSUF"`
	GIBSMun          TIBSCBSMonoTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/bpe gIBSMun" json:"gIBSMun"`
	VIBS             TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/bpe vIBS" json:"vIBS"`
	VCredPres        TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/bpe vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC               `xml:"http://www.portalfiscal.inf.br/bpe vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGCBS struct {
	VDif             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDif" json:"vDif"`
	VDevTrib         TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDevTrib" json:"vDevTrib"`
	VCBS             TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBS" json:"vCBS"`
	VCredPres        TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCredPres" json:"vCredPres"`
	VCredPresCondSus TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCredPresCondSus" json:"vCredPresCondSus"`
}

type TIBSCBSMonoTotGMono struct {
	VIBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSMono" json:"vIBSMono"`
	VCBSMono      TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBSMono" json:"vCBSMono"`
	VIBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSMonoReten" json:"vIBSMonoReten"`
	VCBSMonoReten TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBSMonoReten" json:"vCBSMonoReten"`
	VIBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSMonoRet" json:"vIBSMonoRet"`
	VCBSMonoRet   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBSMonoRet" json:"vCBSMonoRet"`
}

type TIBSCBSMonoTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSMonoTot struct {
	VBCIBSCBS    TDec1302RTC                 `xml:"http://www.portalfiscal.inf.br/bpe vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         *TIBSCBSMonoTotGIBS         `xml:"http://www.portalfiscal.inf.br/bpe gIBS,omitempty" json:"gIBS,omitempty"`
	GCBS         *TIBSCBSMonoTotGCBS         `xml:"http://www.portalfiscal.inf.br/bpe gCBS,omitempty" json:"gCBS,omitempty"`
	GMono        *TIBSCBSMonoTotGMono        `xml:"http://www.portalfiscal.inf.br/bpe gMono,omitempty" json:"gMono,omitempty"`
	GEstornoCred *TIBSCBSMonoTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIBSCBSTotGIBSGIBSUF struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDevTrib" json:"vDevTrib"`
	VIBSUF   TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSUF" json:"vIBSUF"`
}

type TIBSCBSTotGIBSGIBSMun struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDevTrib" json:"vDevTrib"`
	VIBSMun  TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSMun" json:"vIBSMun"`
}

type TIBSCBSTotGIBS struct {
	GIBSUF  TIBSCBSTotGIBSGIBSUF  `xml:"http://www.portalfiscal.inf.br/bpe gIBSUF" json:"gIBSUF"`
	GIBSMun TIBSCBSTotGIBSGIBSMun `xml:"http://www.portalfiscal.inf.br/bpe gIBSMun" json:"gIBSMun"`
	VIBS    TDec1302RTC           `xml:"http://www.portalfiscal.inf.br/bpe vIBS" json:"vIBS"`
}

type TIBSCBSTotGCBS struct {
	VDif     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDif" json:"vDif"`
	VDevTrib TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vDevTrib" json:"vDevTrib"`
	VCBS     TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBS" json:"vCBS"`
}

type TIBSCBSTotGEstornoCred struct {
	VIBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBSEstCred" json:"vCBSEstCred"`
}

type TIBSCBSTot struct {
	VBCIBSCBS    TDec1302RTC             `xml:"http://www.portalfiscal.inf.br/bpe vBCIBSCBS" json:"vBCIBSCBS"`
	GIBS         TIBSCBSTotGIBS          `xml:"http://www.portalfiscal.inf.br/bpe gIBS" json:"gIBS"`
	GCBS         TIBSCBSTotGCBS          `xml:"http://www.portalfiscal.inf.br/bpe gCBS" json:"gCBS"`
	GEstornoCred *TIBSCBSTotGEstornoCred `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIPv4 string

type TIS struct {
	CSTIS        TCST           `xml:"http://www.portalfiscal.inf.br/bpe CSTIS" json:"CSTIS"`
	CClassTribIS TcClassTrib    `xml:"http://www.portalfiscal.inf.br/bpe cClassTribIS" json:"cClassTribIS"`
	VBCIS        *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vBCIS,omitempty" json:"vBCIS,omitempty"`
	PIS          *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pIS,omitempty" json:"pIS,omitempty"`
	AdRemIS      *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe adRemIS,omitempty" json:"adRemIS,omitempty"`
	UTrib        *string        `xml:"http://www.portalfiscal.inf.br/bpe uTrib,omitempty" json:"uTrib,omitempty"`
	QTrib        *TDec1104OpRTC `xml:"http://www.portalfiscal.inf.br/bpe qTrib,omitempty" json:"qTrib,omitempty"`
	VIS          *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIS,omitempty" json:"vIS,omitempty"`
}

type TISTot struct {
	VIS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIS" json:"vIS"`
}

type TIDLote string

type TIe string

type TIeDest string

type TImpICMS00 struct {
	CST   string   `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	VBC   TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vBC" json:"vBC"`
	PICMS TDec0302 `xml:"http://www.portalfiscal.inf.br/bpe pICMS" json:"pICMS"`
	VICMS TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vICMS" json:"vICMS"`
}

type TImpICMS20 struct {
	CST        string      `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	PRedBC     TDec0302Opc `xml:"http://www.portalfiscal.inf.br/bpe pRedBC" json:"pRedBC"`
	VBC        TDec1302    `xml:"http://www.portalfiscal.inf.br/bpe vBC" json:"vBC"`
	PICMS      TDec0302    `xml:"http://www.portalfiscal.inf.br/bpe pICMS" json:"pICMS"`
	VICMS      TDec1302    `xml:"http://www.portalfiscal.inf.br/bpe vICMS" json:"vICMS"`
	VICMSDeson *TDec1302   `xml:"http://www.portalfiscal.inf.br/bpe vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string     `xml:"http://www.portalfiscal.inf.br/bpe cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS45 struct {
	CST        string    `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	VICMSDeson *TDec1302 `xml:"http://www.portalfiscal.inf.br/bpe vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string   `xml:"http://www.portalfiscal.inf.br/bpe cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMS90 struct {
	CST        string       `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	PRedBC     *TDec0302Opc `xml:"http://www.portalfiscal.inf.br/bpe pRedBC,omitempty" json:"pRedBC,omitempty"`
	VBC        TDec1302     `xml:"http://www.portalfiscal.inf.br/bpe vBC" json:"vBC"`
	PICMS      TDec0302     `xml:"http://www.portalfiscal.inf.br/bpe pICMS" json:"pICMS"`
	VICMS      TDec1302     `xml:"http://www.portalfiscal.inf.br/bpe vICMS" json:"vICMS"`
	VCred      *TDec1302    `xml:"http://www.portalfiscal.inf.br/bpe vCred,omitempty" json:"vCred,omitempty"`
	VICMSDeson *TDec1302    `xml:"http://www.portalfiscal.inf.br/bpe vICMSDeson,omitempty" json:"vICMSDeson,omitempty"`
	CBenef     *string      `xml:"http://www.portalfiscal.inf.br/bpe cBenef,omitempty" json:"cBenef,omitempty"`
}

type TImpICMSSN struct {
	CST   string `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	IndSN string `xml:"http://www.portalfiscal.inf.br/bpe indSN" json:"indSN"`
}

type TImp struct {
	ICMS00 *TImpICMS00 `xml:"http://www.portalfiscal.inf.br/bpe ICMS00,omitempty" json:"ICMS00,omitempty"`
	ICMS20 *TImpICMS20 `xml:"http://www.portalfiscal.inf.br/bpe ICMS20,omitempty" json:"ICMS20,omitempty"`
	ICMS45 *TImpICMS45 `xml:"http://www.portalfiscal.inf.br/bpe ICMS45,omitempty" json:"ICMS45,omitempty"`
	ICMS90 *TImpICMS90 `xml:"http://www.portalfiscal.inf.br/bpe ICMS90,omitempty" json:"ICMS90,omitempty"`
	ICMSSN *TImpICMSSN `xml:"http://www.portalfiscal.inf.br/bpe ICMSSN,omitempty" json:"ICMSSN,omitempty"`
}

type TIndDoacao string

type TIndPres string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModBPe string

type TModal string

type TMonofasiaGMonoPadrao struct {
	QBCMono  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/bpe qBCMono" json:"qBCMono"`
	AdRemIBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe adRemIBS" json:"adRemIBS"`
	AdRemCBS TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe adRemCBS" json:"adRemCBS"`
	VIBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSMono" json:"vIBSMono"`
	VCBSMono TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vCBSMono" json:"vCBSMono"`
}

type TMonofasiaGMonoReten struct {
	QBCMonoReten  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/bpe qBCMonoReten" json:"qBCMonoReten"`
	AdRemIBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe adRemIBSReten" json:"adRemIBSReten"`
	VIBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSMonoReten" json:"vIBSMonoReten"`
	AdRemCBSReten TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe adRemCBSReten" json:"adRemCBSReten"`
	VCBSMonoReten TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vCBSMonoReten" json:"vCBSMonoReten"`
}

type TMonofasiaGMonoRet struct {
	QBCMonoRet  TDec1104RTC   `xml:"http://www.portalfiscal.inf.br/bpe qBCMonoRet" json:"qBCMonoRet"`
	AdRemIBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe adRemIBSRet" json:"adRemIBSRet"`
	VIBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSMonoRet" json:"vIBSMonoRet"`
	AdRemCBSRet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe adRemCBSRet" json:"adRemCBSRet"`
	VCBSMonoRet TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vCBSMonoRet" json:"vCBSMonoRet"`
}

type TMonofasiaGMonoDif struct {
	PDifIBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pDifIBS" json:"pDifIBS"`
	VIBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSMonoDif" json:"vIBSMonoDif"`
	PDifCBS     TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pDifCBS" json:"pDifCBS"`
	VCBSMonoDif TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vCBSMonoDif" json:"vCBSMonoDif"`
}

type TMonofasia struct {
	GMonoPadrao     *TMonofasiaGMonoPadrao `xml:"http://www.portalfiscal.inf.br/bpe gMonoPadrao,omitempty" json:"gMonoPadrao,omitempty"`
	GMonoReten      *TMonofasiaGMonoReten  `xml:"http://www.portalfiscal.inf.br/bpe gMonoReten,omitempty" json:"gMonoReten,omitempty"`
	GMonoRet        *TMonofasiaGMonoRet    `xml:"http://www.portalfiscal.inf.br/bpe gMonoRet,omitempty" json:"gMonoRet,omitempty"`
	GMonoDif        *TMonofasiaGMonoDif    `xml:"http://www.portalfiscal.inf.br/bpe gMonoDif,omitempty" json:"gMonoDif,omitempty"`
	VTotIBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/bpe vTotIBSMonoItem" json:"vTotIBSMonoItem"`
	VTotCBSMonoItem TDec1302RTC            `xml:"http://www.portalfiscal.inf.br/bpe vTotCBSMonoItem" json:"vTotCBSMonoItem"`
}

type TMotivo string

type TNF string

type TNSU string

type TOperCompraGov string

type TPagRef struct {
	RefDFe []TChDFeRTC `xml:"http://www.portalfiscal.inf.br/bpe refDFe,omitempty" json:"refDFe,omitempty"`
}

type TPagamentoRTC struct {
	NPag        string       `xml:"nPag,attr" json:"nPag"`
	IDTransacao string       `xml:"idTransacao,attr" json:"idTransacao"`
	TpMeioPgto  string       `xml:"http://www.portalfiscal.inf.br/bpe tpMeioPgto" json:"tpMeioPgto"`
	CNPJReceb   TCNPJRTC     `xml:"http://www.portalfiscal.inf.br/bpe CNPJReceb" json:"CNPJReceb"`
	CNPJBasePSP TCNPJBaseRTC `xml:"http://www.portalfiscal.inf.br/bpe CNPJBasePSP" json:"CNPJBasePSP"`
}

type TPlaca string

type TProcEmi string

type TProcEvento struct {
	Versao        TVerEvento    `xml:"versao,attr" json:"versao"`
	IpTransmissor *TIPv4        `xml:"ipTransmissor,attr,omitempty" json:"ipTransmissor,omitempty"`
	NPortaCon     *string       `xml:"nPortaCon,attr,omitempty" json:"nPortaCon,omitempty"`
	DhConexao     *TDateTimeUTC `xml:"dhConexao,attr,omitempty" json:"dhConexao,omitempty"`
	EventoBPe     TEvento       `xml:"http://www.portalfiscal.inf.br/bpe eventoBPe" json:"eventoBPe"`
	RetEventoBPe  TRetEvento    `xml:"http://www.portalfiscal.inf.br/bpe retEventoBPe" json:"retEventoBPe"`
}

type TProt string

type TProtBPeInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/bpe verAplic" json:"verAplic"`
	ChBPe    TChDFe           `xml:"http://www.portalfiscal.inf.br/bpe chBPe" json:"chBPe"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/bpe dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/bpe nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/bpe digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/bpe cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/bpe xMotivo" json:"xMotivo"`
}

type TProtBPeInfFisco struct {
	CMsg string  `xml:"http://www.portalfiscal.inf.br/bpe cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/bpe xMsg" json:"xMsg"`
}

type TProtBPe struct {
	Versao    TVerBPe           `xml:"versao,attr" json:"versao"`
	InfProt   TProtBPeInfProt   `xml:"http://www.portalfiscal.inf.br/bpe infProt" json:"infProt"`
	InfFisco  *TProtBPeInfFisco `xml:"http://www.portalfiscal.inf.br/bpe infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType    `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRBSN string

type TRec string

type TRed struct {
	PRedAliq  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqEfet" json:"pAliqEfet"`
}

type TRespTec struct {
	CNPJ     TCNPJ   `xml:"http://www.portalfiscal.inf.br/bpe CNPJ" json:"CNPJ"`
	XContato string  `xml:"http://www.portalfiscal.inf.br/bpe xContato" json:"xContato"`
	Email    TEmail  `xml:"http://www.portalfiscal.inf.br/bpe email" json:"email"`
	Fone     string  `xml:"http://www.portalfiscal.inf.br/bpe fone" json:"fone"`
	IDCSRT   *string `xml:"http://www.portalfiscal.inf.br/bpe idCSRT,omitempty" json:"idCSRT,omitempty"`
	HashCSRT *string `xml:"http://www.portalfiscal.inf.br/bpe hashCSRT,omitempty" json:"hashCSRT,omitempty"`
}

type TRetBPe struct {
	Versao   TVerBPe    `xml:"versao,attr" json:"versao"`
	TpAmb    TAmb       `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/bpe cUF" json:"cUF"`
	VerAplic TVerAplic  `xml:"http://www.portalfiscal.inf.br/bpe verAplic" json:"verAplic"`
	CStat    TStat      `xml:"http://www.portalfiscal.inf.br/bpe cStat" json:"cStat"`
	XMotivo  TMotivo    `xml:"http://www.portalfiscal.inf.br/bpe xMotivo" json:"xMotivo"`
	ProtBPe  *TProtBPe  `xml:"http://www.portalfiscal.inf.br/bpe protBPe,omitempty" json:"protBPe,omitempty"`
}

type TRetConsSitBPe struct {
	Versao        TVerConsSitBPe `xml:"versao,attr" json:"versao"`
	TpAmb         TAmb           `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	VerAplic      TVerAplic      `xml:"http://www.portalfiscal.inf.br/bpe verAplic" json:"verAplic"`
	CStat         TStat          `xml:"http://www.portalfiscal.inf.br/bpe cStat" json:"cStat"`
	XMotivo       TMotivo        `xml:"http://www.portalfiscal.inf.br/bpe xMotivo" json:"xMotivo"`
	CUF           TCodUFIBGE     `xml:"http://www.portalfiscal.inf.br/bpe cUF" json:"cUF"`
	ProtBPe       []TProtBPe     `xml:"http://www.portalfiscal.inf.br/bpe protBPe,omitempty" json:"protBPe,omitempty"`
	ProcEventoBPe []TProcEvento  `xml:"http://www.portalfiscal.inf.br/bpe procEventoBPe,omitempty" json:"procEventoBPe,omitempty"`
}

type TRetEventoInfEvento struct {
	ID          *string       `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb       TAmb          `xml:"http://www.portalfiscal.inf.br/bpe tpAmb" json:"tpAmb"`
	VerAplic    TVerAplic     `xml:"http://www.portalfiscal.inf.br/bpe verAplic" json:"verAplic"`
	COrgao      TCOrgaoIBGE   `xml:"http://www.portalfiscal.inf.br/bpe cOrgao" json:"cOrgao"`
	CStat       TStat         `xml:"http://www.portalfiscal.inf.br/bpe cStat" json:"cStat"`
	XMotivo     TMotivo       `xml:"http://www.portalfiscal.inf.br/bpe xMotivo" json:"xMotivo"`
	ChBPe       *TChDFe       `xml:"http://www.portalfiscal.inf.br/bpe chBPe,omitempty" json:"chBPe,omitempty"`
	TpEvento    *string       `xml:"http://www.portalfiscal.inf.br/bpe tpEvento,omitempty" json:"tpEvento,omitempty"`
	XEvento     *string       `xml:"http://www.portalfiscal.inf.br/bpe xEvento,omitempty" json:"xEvento,omitempty"`
	NSeqEvento  *string       `xml:"http://www.portalfiscal.inf.br/bpe nSeqEvento,omitempty" json:"nSeqEvento,omitempty"`
	DhRegEvento *TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/bpe dhRegEvento,omitempty" json:"dhRegEvento,omitempty"`
	NProt       *TProt        `xml:"http://www.portalfiscal.inf.br/bpe nProt,omitempty" json:"nProt,omitempty"`
}

type TRetEvento struct {
	Versao    string              `xml:"versao,attr" json:"versao"`
	InfEvento TRetEventoInfEvento `xml:"http://www.portalfiscal.inf.br/bpe infEvento" json:"infEvento"`
	Signature *SignatureType      `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TStringRTC string

type TTime string

type TTipoBPe string

type TTipoBPeTM string

type TTipoOper string

type TTipoServ string

type TTipoSubstituicao string

type TTipoSubstituicaoTA string

type TTotalSN struct {
	VRBSNTot          TDec1302RTC  `xml:"http://www.portalfiscal.inf.br/bpe vRBSNTot" json:"vRBSNTot"`
	VIBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSSNTot,omitempty" json:"vIBSSNTot,omitempty"`
	VIBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBSSNtotPendSusp,omitempty" json:"vIBSSNtotPendSusp,omitempty"`
	VCBSSNTot         *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBSSNTot,omitempty" json:"vCBSSNTot,omitempty"`
	VCBSSNtotPendSusp *TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBSSNtotPendSusp,omitempty" json:"vCBSSNtotPendSusp,omitempty"`
}

type TTpCredPresIBSZFM string

type TTransfCred struct {
	VIBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vIBS" json:"vIBS"`
	VCBS TDec1302RTC `xml:"http://www.portalfiscal.inf.br/bpe vCBS" json:"vCBS"`
}

type TTribBPe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/bpe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/bpe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCTe struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/bpe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/bpe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribCompraGov struct {
	PAliqIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vTribCBS" json:"vTribCBS"`
}

type TTribItemSN struct {
	NItem        string         `xml:"nItem,attr" json:"nItem"`
	VRBSNItem    TDec1302RTC    `xml:"http://www.portalfiscal.inf.br/bpe vRBSNItem" json:"vRBSNItem"`
	TpRBSN       TRBSN          `xml:"http://www.portalfiscal.inf.br/bpe tpRBSN" json:"tpRBSN"`
	PIBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pIBSSN,omitempty" json:"pIBSSN,omitempty"`
	VIBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSSN,omitempty" json:"vIBSSN,omitempty"`
	PCBSSN       *TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pCBSSN,omitempty" json:"pCBSSN,omitempty"`
	VCBSSN       *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vCBSSN,omitempty" json:"vCBSSN,omitempty"`
	VIBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vIBSPendSusp,omitempty" json:"vIBSPendSusp,omitempty"`
	VCBSPendSusp *TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vCBSPendSusp,omitempty" json:"vCBSPendSusp,omitempty"`
}

type TTribNF3e struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/bpe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/bpe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFAg struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/bpe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/bpe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFCe struct {
	CST         TCST        `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	CClassTrib  TcClassTrib `xml:"http://www.portalfiscal.inf.br/bpe cClassTrib" json:"cClassTrib"`
	IndDoacao   *TIndDoacao `xml:"http://www.portalfiscal.inf.br/bpe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS     *TCIBSNFe   `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono *TMonofasia `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
}

type TTribNFCom struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/bpe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/bpe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFGas struct {
	CST          TCST          `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	CClassTrib   TcClassTrib   `xml:"http://www.portalfiscal.inf.br/bpe cClassTrib" json:"cClassTrib"`
	IndDoacao    *TIndDoacao   `xml:"http://www.portalfiscal.inf.br/bpe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS      *TCIBS        `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TEstornoCred `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TTribNFe struct {
	CST             TCST             `xml:"http://www.portalfiscal.inf.br/bpe CST" json:"CST"`
	CClassTrib      TcClassTrib      `xml:"http://www.portalfiscal.inf.br/bpe cClassTrib" json:"cClassTrib"`
	IndDoacao       *TIndDoacao      `xml:"http://www.portalfiscal.inf.br/bpe indDoacao,omitempty" json:"indDoacao,omitempty"`
	GIBSCBS         *TCIBSNFe        `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GIBSCBSMono     *TMonofasia      `xml:"http://www.portalfiscal.inf.br/bpe gIBSCBSMono,omitempty" json:"gIBSCBSMono,omitempty"`
	GTransfCred     *TTransfCred     `xml:"http://www.portalfiscal.inf.br/bpe gTransfCred,omitempty" json:"gTransfCred,omitempty"`
	GAjusteCompet   *TAjusteCompet   `xml:"http://www.portalfiscal.inf.br/bpe gAjusteCompet,omitempty" json:"gAjusteCompet,omitempty"`
	GEstornoCred    *TEstornoCred    `xml:"http://www.portalfiscal.inf.br/bpe gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
	GCredPresOper   *TCredPresOper   `xml:"http://www.portalfiscal.inf.br/bpe gCredPresOper,omitempty" json:"gCredPresOper,omitempty"`
	GCredPresIBSZFM *TCredPresIBSZFM `xml:"http://www.portalfiscal.inf.br/bpe gCredPresIBSZFM,omitempty" json:"gCredPresIBSZFM,omitempty"`
}

type TTribRegular struct {
	CSTReg             TCST          `xml:"http://www.portalfiscal.inf.br/bpe CSTReg" json:"CSTReg"`
	CClassTribReg      TcClassTrib   `xml:"http://www.portalfiscal.inf.br/bpe cClassTribReg" json:"cClassTribReg"`
	PAliqEfetRegIBSUF  TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqEfetRegIBSUF" json:"pAliqEfetRegIBSUF"`
	VTribRegIBSUF      TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vTribRegIBSUF" json:"vTribRegIBSUF"`
	PAliqEfetRegIBSMun TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqEfetRegIBSMun" json:"pAliqEfetRegIBSMun"`
	VTribRegIBSMun     TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vTribRegIBSMun" json:"vTribRegIBSMun"`
	PAliqEfetRegCBS    TDec030204RTC `xml:"http://www.portalfiscal.inf.br/bpe pAliqEfetRegCBS" json:"pAliqEfetRegCBS"`
	VTribRegCBS        TDec1302RTC   `xml:"http://www.portalfiscal.inf.br/bpe vTribRegCBS" json:"vTribRegCBS"`
}

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerBPe string

type TVerConsSitBPe string

type TVerEvento string

type Tano string

type TcClassTrib string

type TcCredPres string

type DigestValueType string

type KeyInfoType struct {
	ID       *ID          `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	X509Data X509DataType `xml:"http://www.w3.org/2000/09/xmldsig# X509Data" json:"X509Data"`
}

type ReferenceTypeDigestMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type ReferenceType struct {
	ID           *ID                       `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	URI          string                    `xml:"URI,attr" json:"URI"`
	TypeField    *AnyURI                   `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	Transforms   TransformsType            `xml:"http://www.w3.org/2000/09/xmldsig# Transforms" json:"Transforms"`
	DigestMethod ReferenceTypeDigestMethod `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod" json:"DigestMethod"`
	DigestValue  DigestValueType           `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue" json:"DigestValue"`
}

type SignatureType struct {
	ID             *ID                `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	SignedInfo     SignedInfoType     `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo" json:"SignedInfo"`
	SignatureValue SignatureValueType `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue" json:"SignatureValue"`
	KeyInfo        KeyInfoType        `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo" json:"KeyInfo"`
}

type SignatureValueType struct {
	Value Base64Binary `xml:",chardata" json:"value"`
	ID    *ID          `xml:"Id,attr,omitempty" json:"Id,omitempty"`
}

type SignedInfoTypeCanonicalizationMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type SignedInfoTypeSignatureMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type SignedInfoType struct {
	ID                     *ID                                  `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	CanonicalizationMethod SignedInfoTypeCanonicalizationMethod `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod" json:"CanonicalizationMethod"`
	SignatureMethod        SignedInfoTypeSignatureMethod        `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod" json:"SignatureMethod"`
	Reference              ReferenceType                        `xml:"http://www.w3.org/2000/09/xmldsig# Reference" json:"Reference"`
}

type TTransformURI string

type TransformType struct {
	Algorithm TTransformURI `xml:"Algorithm,attr" json:"Algorithm"`
	XPath     []String      `xml:"http://www.w3.org/2000/09/xmldsig# XPath,omitempty" json:"XPath,omitempty"`
}

type TransformsType struct {
	Transform []TransformType `xml:"http://www.w3.org/2000/09/xmldsig# Transform,omitempty" json:"Transform,omitempty"`
}

type X509DataType struct {
	X509Certificate Base64Binary `xml:"http://www.w3.org/2000/09/xmldsig# X509Certificate" json:"X509Certificate"`
}

type ID string

type AnyURI string

type Base64Binary string

type GYearMonth string

type String string

// ConsSitBPe represents the global element {http://www.portalfiscal.inf.br/bpe}consSitBPe.
type ConsSitBPe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/bpe consSitBPe" json:"-"`
	TConsSitBPe
}

// Signature represents the global element {http://www.w3.org/2000/09/xmldsig#}Signature.
type Signature struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"-"`
	SignatureType
}

// MarshalXML avoids re-emitting a default xmlns already supplied by encoding/xml.
func (value RawElement) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	if value.XMLName.Local != "" {
		start.Name = value.XMLName
	}
	for _, attr := range value.Attrs {
		if attr.Name.Space == "" && attr.Name.Local == "xmlns" && attr.Value == start.Name.Space {
			continue
		}
		if attr.Name.Space == "xmlns" {
			attr.Name = xml.Name{Local: "xmlns:" + attr.Name.Local}
		}
		start.Attr = append(start.Attr, attr)
	}
	body := struct {
		InnerXML string `xml:",innerxml"`
	}{value.InnerXML}
	return encoder.EncodeElement(body, start)
}
