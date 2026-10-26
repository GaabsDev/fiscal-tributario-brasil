// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package ctemodalaereo_v4_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAmb string

type TCOrgaoIBGE string

type TChDFe string

type TCNPJ string

type TCNPJOpc string

type TCodMunIBGE string

type TCodUFIBGE string

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

type TFone string

type TIPv4 string

type TIe string

type TIeDest string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModCT string

type TModCTOS string

type TModCTCargaOS string

type TModGTVe string

type TModNF string

type TMotivo string

type TNF string

type TPlaca string

type TProt string

type TRSAKeyValueType struct {
	Modulus  Base64Binary `xml:"http://www.portalfiscal.inf.br/cte Modulus" json:"Modulus"`
	Exponent Base64Binary `xml:"http://www.portalfiscal.inf.br/cte Exponent" json:"Exponent"`
}

type TRec string

type TSerie string

type TServ string

type TStat string

type TString string

type TUFSemEX string

type TUF string

type TVerAplic string

type Tano string

type TtipoUnidCarga string

type TtipoUnidTransp string

type Base64Binary string

type String string

type AereoContentNatCarga struct {
	XDime    *string  `xml:"http://www.portalfiscal.inf.br/cte xDime,omitempty" json:"xDime,omitempty"`
	CInfManu []string `xml:"http://www.portalfiscal.inf.br/cte cInfManu,omitempty" json:"cInfManu,omitempty"`
}

type AereoContentTarifa struct {
	CL   string   `xml:"http://www.portalfiscal.inf.br/cte CL" json:"CL"`
	CTar *string  `xml:"http://www.portalfiscal.inf.br/cte cTar,omitempty" json:"cTar,omitempty"`
	VTar TDec1302 `xml:"http://www.portalfiscal.inf.br/cte vTar" json:"vTar"`
}

type AereoContentPeriInfTotAP struct {
	QTotProd TDec1104 `xml:"http://www.portalfiscal.inf.br/cte qTotProd" json:"qTotProd"`
	UniAP    string   `xml:"http://www.portalfiscal.inf.br/cte uniAP" json:"uniAP"`
}

type AereoContentPeri struct {
	NONU     string                   `xml:"http://www.portalfiscal.inf.br/cte nONU" json:"nONU"`
	QTotEmb  string                   `xml:"http://www.portalfiscal.inf.br/cte qTotEmb" json:"qTotEmb"`
	InfTotAP AereoContentPeriInfTotAP `xml:"http://www.portalfiscal.inf.br/cte infTotAP" json:"infTotAP"`
}

type AereoContent struct {
	NMinu      *string              `xml:"http://www.portalfiscal.inf.br/cte nMinu,omitempty" json:"nMinu,omitempty"`
	NOCA       *string              `xml:"http://www.portalfiscal.inf.br/cte nOCA,omitempty" json:"nOCA,omitempty"`
	DPrevAereo TData                `xml:"http://www.portalfiscal.inf.br/cte dPrevAereo" json:"dPrevAereo"`
	NatCarga   AereoContentNatCarga `xml:"http://www.portalfiscal.inf.br/cte natCarga" json:"natCarga"`
	Tarifa     AereoContentTarifa   `xml:"http://www.portalfiscal.inf.br/cte tarifa" json:"tarifa"`
	Peri       []AereoContentPeri   `xml:"http://www.portalfiscal.inf.br/cte peri,omitempty" json:"peri,omitempty"`
}

// Aereo represents the global element {http://www.portalfiscal.inf.br/cte}aereo.
type Aereo struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/cte aereo" json:"-"`
	AereoContent
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
