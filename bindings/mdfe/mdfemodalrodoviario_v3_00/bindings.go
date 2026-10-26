// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package mdfemodalrodoviario_v3_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAmb string

type TCIOT string

type TCOrgaoIBGE string

type TChCTe string

type TChMDFe string

type TChNFe string

type TCNPJ string

type TCNPJOpc string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

type TCPF string

type TCPFVar string

type TData string

type TDateTimeUTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec0303 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec0906Opc string

type TDec1104 string

type TDec1104Opc string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec1302 string

type TDec1302Opc string

type TEmit string

type TIPv4 string

type TIe string

type TIeDest string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModMD string

type TMotivo string

type TNF string

type TNSU string

type TPlaca string

type TProt string

type TRNTRC string

type TRSAKeyValueType struct {
	Modulus  Base64Binary `xml:"http://www.portalfiscal.inf.br/mdfe Modulus" json:"Modulus"`
	Exponent Base64Binary `xml:"http://www.portalfiscal.inf.br/mdfe Exponent" json:"Exponent"`
}

type TRec string

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TTransp string

type TUF string

type TVerAplic string

type Tano string

type TtipoUnidCarga string

type TtipoUnidTransp string

type Base64Binary string

type String string

type RodoContentInfANTTInfCIOT struct {
	CIOT *string   `xml:"http://www.portalfiscal.inf.br/mdfe CIOT,omitempty" json:"CIOT,omitempty"`
	CPF  *TCPF     `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	CNPJ *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
}

type RodoContentInfANTTValePedDisp struct {
	CNPJForn  string    `xml:"http://www.portalfiscal.inf.br/mdfe CNPJForn" json:"CNPJForn"`
	CNPJPg    *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/mdfe CNPJPg,omitempty" json:"CNPJPg,omitempty"`
	CPFPg     *TCPF     `xml:"http://www.portalfiscal.inf.br/mdfe CPFPg,omitempty" json:"CPFPg,omitempty"`
	NCompra   *string   `xml:"http://www.portalfiscal.inf.br/mdfe nCompra,omitempty" json:"nCompra,omitempty"`
	VValePed  TDec1302  `xml:"http://www.portalfiscal.inf.br/mdfe vValePed" json:"vValePed"`
	TpValePed *string   `xml:"http://www.portalfiscal.inf.br/mdfe tpValePed,omitempty" json:"tpValePed,omitempty"`
}

type RodoContentInfANTTValePed struct {
	Disp          []RodoContentInfANTTValePedDisp `xml:"http://www.portalfiscal.inf.br/mdfe disp,omitempty" json:"disp,omitempty"`
	CategCombVeic *string                         `xml:"http://www.portalfiscal.inf.br/mdfe categCombVeic,omitempty" json:"categCombVeic,omitempty"`
}

type RodoContentInfANTTInfContratanteInfContrato struct {
	NroContrato     string      `xml:"http://www.portalfiscal.inf.br/mdfe NroContrato" json:"NroContrato"`
	VContratoGlobal TDec1302Opc `xml:"http://www.portalfiscal.inf.br/mdfe vContratoGlobal" json:"vContratoGlobal"`
}

type RodoContentInfANTTInfContratante struct {
	XNome         *string                                      `xml:"http://www.portalfiscal.inf.br/mdfe xNome,omitempty" json:"xNome,omitempty"`
	CPF           *TCPF                                        `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	CNPJ          *TCNPJ                                       `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	IDEstrangeiro *string                                      `xml:"http://www.portalfiscal.inf.br/mdfe idEstrangeiro,omitempty" json:"idEstrangeiro,omitempty"`
	InfContrato   *RodoContentInfANTTInfContratanteInfContrato `xml:"http://www.portalfiscal.inf.br/mdfe infContrato,omitempty" json:"infContrato,omitempty"`
}

type RodoContentInfANTTInfPagComp struct {
	TpComp string   `xml:"http://www.portalfiscal.inf.br/mdfe tpComp" json:"tpComp"`
	VComp  TDec1302 `xml:"http://www.portalfiscal.inf.br/mdfe vComp" json:"vComp"`
	XComp  *string  `xml:"http://www.portalfiscal.inf.br/mdfe xComp,omitempty" json:"xComp,omitempty"`
}

type RodoContentInfANTTInfPagInfPrazo struct {
	NParcela string      `xml:"http://www.portalfiscal.inf.br/mdfe nParcela" json:"nParcela"`
	DVenc    TData       `xml:"http://www.portalfiscal.inf.br/mdfe dVenc" json:"dVenc"`
	VParcela TDec1302Opc `xml:"http://www.portalfiscal.inf.br/mdfe vParcela" json:"vParcela"`
}

type RodoContentInfANTTInfPagInfBanc struct {
	CodBanco   *string   `xml:"http://www.portalfiscal.inf.br/mdfe codBanco,omitempty" json:"codBanco,omitempty"`
	CodAgencia *string   `xml:"http://www.portalfiscal.inf.br/mdfe codAgencia,omitempty" json:"codAgencia,omitempty"`
	CNPJIPEF   *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/mdfe CNPJIPEF,omitempty" json:"CNPJIPEF,omitempty"`
	PIX        *string   `xml:"http://www.portalfiscal.inf.br/mdfe PIX,omitempty" json:"PIX,omitempty"`
}

type RodoContentInfANTTInfPag struct {
	XNome             *string                            `xml:"http://www.portalfiscal.inf.br/mdfe xNome,omitempty" json:"xNome,omitempty"`
	CPF               *TCPF                              `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	CNPJ              *TCNPJOpc                          `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	IDEstrangeiro     *string                            `xml:"http://www.portalfiscal.inf.br/mdfe idEstrangeiro,omitempty" json:"idEstrangeiro,omitempty"`
	Comp              []RodoContentInfANTTInfPagComp     `xml:"http://www.portalfiscal.inf.br/mdfe Comp,omitempty" json:"Comp,omitempty"`
	VContrato         TDec1302                           `xml:"http://www.portalfiscal.inf.br/mdfe vContrato" json:"vContrato"`
	IndAltoDesemp     *string                            `xml:"http://www.portalfiscal.inf.br/mdfe indAltoDesemp,omitempty" json:"indAltoDesemp,omitempty"`
	IndPag            string                             `xml:"http://www.portalfiscal.inf.br/mdfe indPag" json:"indPag"`
	VAdiant           *TDec1302                          `xml:"http://www.portalfiscal.inf.br/mdfe vAdiant,omitempty" json:"vAdiant,omitempty"`
	IndAntecipaAdiant *string                            `xml:"http://www.portalfiscal.inf.br/mdfe indAntecipaAdiant,omitempty" json:"indAntecipaAdiant,omitempty"`
	InfPrazo          []RodoContentInfANTTInfPagInfPrazo `xml:"http://www.portalfiscal.inf.br/mdfe infPrazo,omitempty" json:"infPrazo,omitempty"`
	TpAntecip         *string                            `xml:"http://www.portalfiscal.inf.br/mdfe tpAntecip,omitempty" json:"tpAntecip,omitempty"`
	InfBanc           RodoContentInfANTTInfPagInfBanc    `xml:"http://www.portalfiscal.inf.br/mdfe infBanc" json:"infBanc"`
}

type RodoContentInfANTT struct {
	RNTRC          *TRNTRC                            `xml:"http://www.portalfiscal.inf.br/mdfe RNTRC,omitempty" json:"RNTRC,omitempty"`
	InfCIOT        []RodoContentInfANTTInfCIOT        `xml:"http://www.portalfiscal.inf.br/mdfe infCIOT,omitempty" json:"infCIOT,omitempty"`
	ValePed        *RodoContentInfANTTValePed         `xml:"http://www.portalfiscal.inf.br/mdfe valePed,omitempty" json:"valePed,omitempty"`
	InfContratante []RodoContentInfANTTInfContratante `xml:"http://www.portalfiscal.inf.br/mdfe infContratante,omitempty" json:"infContratante,omitempty"`
	InfPag         []RodoContentInfANTTInfPag         `xml:"http://www.portalfiscal.inf.br/mdfe infPag,omitempty" json:"infPag,omitempty"`
}

type RodoContentVeicTracaoProp struct {
	CPF    *TCPF     `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	CNPJ   *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	RNTRC  TRNTRC    `xml:"http://www.portalfiscal.inf.br/mdfe RNTRC" json:"RNTRC"`
	XNome  string    `xml:"http://www.portalfiscal.inf.br/mdfe xNome" json:"xNome"`
	IE     *string   `xml:"http://www.portalfiscal.inf.br/mdfe IE,omitempty" json:"IE,omitempty"`
	UF     *TUF      `xml:"http://www.portalfiscal.inf.br/mdfe UF,omitempty" json:"UF,omitempty"`
	TpProp string    `xml:"http://www.portalfiscal.inf.br/mdfe tpProp" json:"tpProp"`
}

type RodoContentVeicTracaoCondutor struct {
	XNome string `xml:"http://www.portalfiscal.inf.br/mdfe xNome" json:"xNome"`
	CPF   TCPF   `xml:"http://www.portalfiscal.inf.br/mdfe CPF" json:"CPF"`
}

type RodoContentVeicTracao struct {
	CInt     *string                         `xml:"http://www.portalfiscal.inf.br/mdfe cInt,omitempty" json:"cInt,omitempty"`
	Placa    string                          `xml:"http://www.portalfiscal.inf.br/mdfe placa" json:"placa"`
	RENAVAM  *string                         `xml:"http://www.portalfiscal.inf.br/mdfe RENAVAM,omitempty" json:"RENAVAM,omitempty"`
	Tara     string                          `xml:"http://www.portalfiscal.inf.br/mdfe tara" json:"tara"`
	CapKG    *string                         `xml:"http://www.portalfiscal.inf.br/mdfe capKG,omitempty" json:"capKG,omitempty"`
	CapM3    *string                         `xml:"http://www.portalfiscal.inf.br/mdfe capM3,omitempty" json:"capM3,omitempty"`
	Prop     *RodoContentVeicTracaoProp      `xml:"http://www.portalfiscal.inf.br/mdfe prop,omitempty" json:"prop,omitempty"`
	Condutor []RodoContentVeicTracaoCondutor `xml:"http://www.portalfiscal.inf.br/mdfe condutor,omitempty" json:"condutor,omitempty"`
	TpRod    string                          `xml:"http://www.portalfiscal.inf.br/mdfe tpRod" json:"tpRod"`
	TpCar    string                          `xml:"http://www.portalfiscal.inf.br/mdfe tpCar" json:"tpCar"`
	UF       *TUF                            `xml:"http://www.portalfiscal.inf.br/mdfe UF,omitempty" json:"UF,omitempty"`
}

type RodoContentVeicReboqueProp struct {
	CPF    *TCPF     `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	CNPJ   *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	RNTRC  TRNTRC    `xml:"http://www.portalfiscal.inf.br/mdfe RNTRC" json:"RNTRC"`
	XNome  string    `xml:"http://www.portalfiscal.inf.br/mdfe xNome" json:"xNome"`
	IE     *string   `xml:"http://www.portalfiscal.inf.br/mdfe IE,omitempty" json:"IE,omitempty"`
	UF     *TUF      `xml:"http://www.portalfiscal.inf.br/mdfe UF,omitempty" json:"UF,omitempty"`
	TpProp string    `xml:"http://www.portalfiscal.inf.br/mdfe tpProp" json:"tpProp"`
}

type RodoContentVeicReboque struct {
	CInt    *string                     `xml:"http://www.portalfiscal.inf.br/mdfe cInt,omitempty" json:"cInt,omitempty"`
	Placa   string                      `xml:"http://www.portalfiscal.inf.br/mdfe placa" json:"placa"`
	RENAVAM *string                     `xml:"http://www.portalfiscal.inf.br/mdfe RENAVAM,omitempty" json:"RENAVAM,omitempty"`
	Tara    string                      `xml:"http://www.portalfiscal.inf.br/mdfe tara" json:"tara"`
	CapKG   string                      `xml:"http://www.portalfiscal.inf.br/mdfe capKG" json:"capKG"`
	CapM3   *string                     `xml:"http://www.portalfiscal.inf.br/mdfe capM3,omitempty" json:"capM3,omitempty"`
	Prop    *RodoContentVeicReboqueProp `xml:"http://www.portalfiscal.inf.br/mdfe prop,omitempty" json:"prop,omitempty"`
	TpCar   string                      `xml:"http://www.portalfiscal.inf.br/mdfe tpCar" json:"tpCar"`
	UF      *TUF                        `xml:"http://www.portalfiscal.inf.br/mdfe UF,omitempty" json:"UF,omitempty"`
}

type RodoContentLacRodo struct {
	NLacre string `xml:"http://www.portalfiscal.inf.br/mdfe nLacre" json:"nLacre"`
}

type RodoContent struct {
	InfANTT     *RodoContentInfANTT      `xml:"http://www.portalfiscal.inf.br/mdfe infANTT,omitempty" json:"infANTT,omitempty"`
	VeicTracao  RodoContentVeicTracao    `xml:"http://www.portalfiscal.inf.br/mdfe veicTracao" json:"veicTracao"`
	VeicReboque []RodoContentVeicReboque `xml:"http://www.portalfiscal.inf.br/mdfe veicReboque,omitempty" json:"veicReboque,omitempty"`
	CodAgPorto  *string                  `xml:"http://www.portalfiscal.inf.br/mdfe codAgPorto,omitempty" json:"codAgPorto,omitempty"`
	LacRodo     []RodoContentLacRodo     `xml:"http://www.portalfiscal.inf.br/mdfe lacRodo,omitempty" json:"lacRodo,omitempty"`
}

// Rodo represents the global element {http://www.portalfiscal.inf.br/mdfe}rodo.
type Rodo struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/mdfe rodo" json:"-"`
	RodoContent
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
