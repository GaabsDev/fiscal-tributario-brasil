// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package e110300_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAmb string

type TCOrgaoIBGE string

type TChNFe string

type TCNPJ string

type TCNPJBase string

type TCNPJOpc string

type TCNPJVar string

type TCodMunIBGE string

type TCodUFIBGE string

type TCPF string

type TCPFVar string

type TData string

type TDateTimeUTC string

type TDec1302 string

type TDec0302 string

type TDec0302Opc string

type TDec030204 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec1104 string

type TDec1104Neg string

type TDec1104Opc string

type TDec1104v string

type TDec1110 string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec13022 string

type TDec1302Opc string

type TEventoPgto struct {
	IDTransacao string    `xml:"idTransacao,attr" json:"idTransacao"`
	TpMeioPgto  string    `xml:"http://www.portalfiscal.inf.br/nfe tpMeioPgto" json:"tpMeioPgto"`
	CNPJReceb   TCNPJ     `xml:"http://www.portalfiscal.inf.br/nfe CNPJReceb" json:"CNPJReceb"`
	CNPJBasePSP TCNPJBase `xml:"http://www.portalfiscal.inf.br/nfe CNPJBasePSP" json:"CNPJBasePSP"`
}

type TIe string

type TIeDest string

type TIeST string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TMod string

type TMotivo string

type TNF string

type TPlaca string

type TProt string

type TRSAKeyValueType struct {
	Modulus  Base64Binary `xml:"http://www.portalfiscal.inf.br/nfe Modulus" json:"Modulus"`
	Exponent Base64Binary `xml:"http://www.portalfiscal.inf.br/nfe Exponent" json:"Exponent"`
}

type TRec string

type TSerie string

type TServ string

type TStat string

type TString string

type TTime string

type TUF string

type TUFEmi string

type TVerAplic string

type Tano string

type TnItem string

type Tpais string

type Base64Binary string

type String string

type DetEventoContent struct {
	Versao      string      `xml:"versao,attr" json:"versao"`
	DescEvento  string      `xml:"http://www.portalfiscal.inf.br/nfe descEvento" json:"descEvento"`
	COrgaoAutor TCodUFIBGE  `xml:"http://www.portalfiscal.inf.br/nfe cOrgaoAutor" json:"cOrgaoAutor"`
	TpAutor     string      `xml:"http://www.portalfiscal.inf.br/nfe tpAutor" json:"tpAutor"`
	VerAplic    TVerAplic   `xml:"http://www.portalfiscal.inf.br/nfe verAplic" json:"verAplic"`
	NProt       TProt       `xml:"http://www.portalfiscal.inf.br/nfe nProt" json:"nProt"`
	GPgto       TEventoPgto `xml:"http://www.portalfiscal.inf.br/nfe gPgto" json:"gPgto"`
}

// DetEvento represents the global element {http://www.portalfiscal.inf.br/nfe}detEvento.
type DetEvento struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe detEvento" json:"-"`
	DetEventoContent
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
