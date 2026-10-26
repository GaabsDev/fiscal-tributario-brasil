// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package consstatservnfag_v1_00

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

type TCodUFIBGEEX string

type TConsStatServ struct {
	Versao TVerConsStat `xml:"versao,attr" json:"versao"`
	TpAmb  TAmb         `xml:"http://www.portalfiscal.inf.br/nfag tpAmb" json:"tpAmb"`
	XServ  TServ        `xml:"http://www.portalfiscal.inf.br/nfag xServ" json:"xServ"`
}

type TCPF string

type TCPFVar string

type TData string

type TDateTimeUTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec0302a04 string

type TDec0302a04Max100 string

type TDec0302a04Opc string

type TDec0303 string

type TDec0303Opc string

type TDec04020408 string

type TDec0503 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec0906 string

type TDec0906Opc string

type TDec0908 string

type TDec0908Opc string

type TDec11001104 string

type TDec1104 string

type TDec1104Opc string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec1302 string

type TDec1302Opc string

type TDec13021304 string

type TDec13021306 string

type TDec13021308 string

type TDec13021310 string

type TIPv4 string

type TIe string

type TIeDest string

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModNFAg string

type TMotivo string

type TNF string

type TNSU string

type TProt string

type TRec string

type TRetConsStatServ struct {
	Versao    TVerConsStat  `xml:"versao,attr" json:"versao"`
	TpAmb     TAmb          `xml:"http://www.portalfiscal.inf.br/nfag tpAmb" json:"tpAmb"`
	VerAplic  string        `xml:"http://www.portalfiscal.inf.br/nfag verAplic" json:"verAplic"`
	CStat     TStat         `xml:"http://www.portalfiscal.inf.br/nfag cStat" json:"cStat"`
	XMotivo   TMotivo       `xml:"http://www.portalfiscal.inf.br/nfag xMotivo" json:"xMotivo"`
	CUF       TCodUFIBGE    `xml:"http://www.portalfiscal.inf.br/nfag cUF" json:"cUF"`
	DhRecbto  TDateTimeUTC  `xml:"http://www.portalfiscal.inf.br/nfag dhRecbto" json:"dhRecbto"`
	TMed      *string       `xml:"http://www.portalfiscal.inf.br/nfag tMed,omitempty" json:"tMed,omitempty"`
	DhRetorno *TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfag dhRetorno,omitempty" json:"dhRetorno,omitempty"`
	XObs      *string       `xml:"http://www.portalfiscal.inf.br/nfag xObs,omitempty" json:"xObs,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerConsStat string

type Tano string

type Integer string

type String string

// ConsStatServNFAg represents the global element {http://www.portalfiscal.inf.br/nfag}consStatServNFAg.
type ConsStatServNFAg struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfag consStatServNFAg" json:"-"`
	TConsStatServ
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
