// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package mdfemodalferroviario_v3_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAmb string

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

type FerrovContentTrem struct {
	XPref  string        `xml:"http://www.portalfiscal.inf.br/mdfe xPref" json:"xPref"`
	DhTrem *TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/mdfe dhTrem,omitempty" json:"dhTrem,omitempty"`
	XOri   string        `xml:"http://www.portalfiscal.inf.br/mdfe xOri" json:"xOri"`
	XDest  string        `xml:"http://www.portalfiscal.inf.br/mdfe xDest" json:"xDest"`
	QVag   string        `xml:"http://www.portalfiscal.inf.br/mdfe qVag" json:"qVag"`
}

type FerrovContentVag struct {
	PesoBC TDec0303 `xml:"http://www.portalfiscal.inf.br/mdfe pesoBC" json:"pesoBC"`
	PesoR  TDec0303 `xml:"http://www.portalfiscal.inf.br/mdfe pesoR" json:"pesoR"`
	TpVag  *string  `xml:"http://www.portalfiscal.inf.br/mdfe tpVag,omitempty" json:"tpVag,omitempty"`
	Serie  string   `xml:"http://www.portalfiscal.inf.br/mdfe serie" json:"serie"`
	NVag   string   `xml:"http://www.portalfiscal.inf.br/mdfe nVag" json:"nVag"`
	NSeq   *string  `xml:"http://www.portalfiscal.inf.br/mdfe nSeq,omitempty" json:"nSeq,omitempty"`
	TU     string   `xml:"http://www.portalfiscal.inf.br/mdfe TU" json:"TU"`
}

type FerrovContent struct {
	Trem FerrovContentTrem  `xml:"http://www.portalfiscal.inf.br/mdfe trem" json:"trem"`
	Vag  []FerrovContentVag `xml:"http://www.portalfiscal.inf.br/mdfe vag,omitempty" json:"vag,omitempty"`
}

// Ferrov represents the global element {http://www.portalfiscal.inf.br/mdfe}ferrov.
type Ferrov struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/mdfe ferrov" json:"-"`
	FerrovContent
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
