// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package enviolotedere_v1_0_1

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TEventoDere struct {
	ID  ID           `xml:"id,attr" json:"id"`
	Any []RawElement `xml:",any" json:"any,omitempty"`
}

type ID string

type String string

type DeREContentLoteEventosIdeContrib struct {
	NrInsc string `xml:"http://www.dere.gov.br/schemas/envioLoteDere/v1_0_1 nrInsc" json:"nrInsc"`
}

type DeREContentLoteEventosEventos struct {
	Evento []TEventoDere `xml:"http://www.dere.gov.br/schemas/envioLoteDere/v1_0_1 evento,omitempty" json:"evento,omitempty"`
}

type DeREContentLoteEventos struct {
	IdeContrib DeREContentLoteEventosIdeContrib `xml:"http://www.dere.gov.br/schemas/envioLoteDere/v1_0_1 ideContrib" json:"ideContrib"`
	Eventos    DeREContentLoteEventosEventos    `xml:"http://www.dere.gov.br/schemas/envioLoteDere/v1_0_1 eventos" json:"eventos"`
}

type DeREContent struct {
	LoteEventos DeREContentLoteEventos `xml:"http://www.dere.gov.br/schemas/envioLoteDere/v1_0_1 loteEventos" json:"loteEventos"`
}

// DeRE represents the global element {http://www.dere.gov.br/schemas/envioLoteDere/v1_0_1}DeRE.
type DeRE struct {
	XMLName xml.Name `xml:"http://www.dere.gov.br/schemas/envioLoteDere/v1_0_1 DeRE" json:"-"`
	DeREContent
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
