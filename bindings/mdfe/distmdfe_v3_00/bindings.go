// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package distmdfe_v3_00

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

type TDistDFe struct {
	Versao     TVerDFe   `xml:"versao,attr" json:"versao"`
	TpAmb      TAmb      `xml:"http://www.portalfiscal.inf.br/mdfe tpAmb" json:"tpAmb"`
	VerAplic   TVerAplic `xml:"http://www.portalfiscal.inf.br/mdfe verAplic" json:"verAplic"`
	IndDFe     string    `xml:"http://www.portalfiscal.inf.br/mdfe indDFe" json:"indDFe"`
	IndCompRet string    `xml:"http://www.portalfiscal.inf.br/mdfe indCompRet" json:"indCompRet"`
	UltNSU     TNSU      `xml:"http://www.portalfiscal.inf.br/mdfe ultNSU" json:"ultNSU"`
}

type TEmit string

type TIPv4 string

type TIe string

type TIeDest string

type TJust string

type TLatitude string

type TLongitude string

type TLoteDistDFeProc struct {
	Schema        String       `xml:"schema,attr" json:"schema"`
	NSU           TNSU         `xml:"NSU,attr" json:"NSU"`
	IpTransmissor *TIPv4       `xml:"ipTransmissor,attr,omitempty" json:"ipTransmissor,omitempty"`
	Any           []RawElement `xml:",any" json:"any,omitempty"`
}

type TLoteDistDFe struct {
	Versao TVerDFe            `xml:"versao,attr" json:"versao"`
	Proc   []TLoteDistDFeProc `xml:"http://www.portalfiscal.inf.br/mdfe proc,omitempty" json:"proc,omitempty"`
}

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

type TRetDistDFe struct {
	Versao           TVerDFe       `xml:"versao,attr" json:"versao"`
	TpAmb            TAmb          `xml:"http://www.portalfiscal.inf.br/mdfe tpAmb" json:"tpAmb"`
	VerAplic         TVerAplic     `xml:"http://www.portalfiscal.inf.br/mdfe verAplic" json:"verAplic"`
	CStat            TStat         `xml:"http://www.portalfiscal.inf.br/mdfe cStat" json:"cStat"`
	XMotivo          TMotivo       `xml:"http://www.portalfiscal.inf.br/mdfe xMotivo" json:"xMotivo"`
	UltNSU           *TNSU         `xml:"http://www.portalfiscal.inf.br/mdfe ultNSU,omitempty" json:"ultNSU,omitempty"`
	LoteDistMDFeComp *Base64Binary `xml:"http://www.portalfiscal.inf.br/mdfe loteDistMDFeComp,omitempty" json:"loteDistMDFeComp,omitempty"`
	LoteDistMDFe     *TLoteDistDFe `xml:"http://www.portalfiscal.inf.br/mdfe loteDistMDFe,omitempty" json:"loteDistMDFe,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TTransp string

type TUF string

type TVerAplic string

type TVerDFe string

type Tano string

type TtipoUnidCarga string

type TtipoUnidTransp string

type Base64Binary string

type String string

type UnsignedInt string

// DistMDFe represents the global element {http://www.portalfiscal.inf.br/mdfe}distMDFe.
type DistMDFe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/mdfe distMDFe" json:"-"`
	TDistDFe
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
