// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package proceventovia_v1_01

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TCEventoVia struct {
	Versao       TSVerNFSeVia   `xml:"versao,attr" json:"versao"`
	InfEventoVia TCInfEventoVia `xml:"http://www.sped.fazenda.gov.br/nfsevia infEventoVia" json:"infEventoVia"`
	Signature    SignatureType  `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCEventoViaProc struct {
	Evento        TCEventoVia        `xml:"http://www.sped.fazenda.gov.br/nfsevia evento" json:"evento"`
	ProtEventoVia TCInfProtEventoVia `xml:"http://www.sped.fazenda.gov.br/nfsevia protEventoVia" json:"protEventoVia"`
}

type TCInfEventoVia struct {
	ID         TSIDEvento      `xml:"id,attr" json:"id"`
	TpAmb      TSTipoAmbiente  `xml:"http://www.sped.fazenda.gov.br/nfsevia tpAmb" json:"tpAmb"`
	VerAplic   TSVersaoEmissor `xml:"http://www.sped.fazenda.gov.br/nfsevia verAplic" json:"verAplic"`
	DhEvento   TSDateTimeUTC   `xml:"http://www.sped.fazenda.gov.br/nfsevia dhEvento" json:"dhEvento"`
	ChNFSeVia  TSChaveNFSeVia  `xml:"http://www.sped.fazenda.gov.br/nfsevia chNFSeVia" json:"chNFSeVia"`
	NSeqEvento TSNum3Dig       `xml:"http://www.sped.fazenda.gov.br/nfsevia nSeqEvento" json:"nSeqEvento"`
	E101101    *TE101101       `xml:"http://www.sped.fazenda.gov.br/nfsevia e101101,omitempty" json:"e101101,omitempty"`
	E105102    *TE105102       `xml:"http://www.sped.fazenda.gov.br/nfsevia e105102,omitempty" json:"e105102,omitempty"`
	E203202    *TE203202       `xml:"http://www.sped.fazenda.gov.br/nfsevia e203202,omitempty" json:"e203202,omitempty"`
	E203204    *TE203204       `xml:"http://www.sped.fazenda.gov.br/nfsevia e203204,omitempty" json:"e203204,omitempty"`
}

type TCInfProtEventoVia struct {
	InfProt TCProtEventoVia `xml:"http://www.sped.fazenda.gov.br/nfsevia infProt" json:"infProt"`
}

type TCProtEventoVia struct {
	TpAmb       TSTipoAmbiente  `xml:"http://www.sped.fazenda.gov.br/nfsevia tpAmb" json:"tpAmb"`
	VerAplic    TSVersaoEmissor `xml:"http://www.sped.fazenda.gov.br/nfsevia verAplic" json:"verAplic"`
	ChNFSeVia   TSChaveNFSeVia  `xml:"http://www.sped.fazenda.gov.br/nfsevia chNFSeVia" json:"chNFSeVia"`
	TpEventoVia TSTipoEventoVia `xml:"http://www.sped.fazenda.gov.br/nfsevia tpEventoVia" json:"tpEventoVia"`
	NSeqEvento  TSNum3Dig       `xml:"http://www.sped.fazenda.gov.br/nfsevia nSeqEvento" json:"nSeqEvento"`
	DhRecbto    TSDateTimeUTC   `xml:"http://www.sped.fazenda.gov.br/nfsevia dhRecbto" json:"dhRecbto"`
	NProt       TSNumProtocolo  `xml:"http://www.sped.fazenda.gov.br/nfsevia nProt" json:"nProt"`
}

type TE101101 struct {
	XDesc   string        `xml:"http://www.sped.fazenda.gov.br/nfsevia xDesc" json:"xDesc"`
	CMotivo TSCodJustCanc `xml:"http://www.sped.fazenda.gov.br/nfsevia cMotivo" json:"cMotivo"`
	XMotivo TSMotivo      `xml:"http://www.sped.fazenda.gov.br/nfsevia xMotivo" json:"xMotivo"`
}

type TE105102 struct {
	XDesc        string         `xml:"http://www.sped.fazenda.gov.br/nfsevia xDesc" json:"xDesc"`
	CMotivo      TSCodJustSubst `xml:"http://www.sped.fazenda.gov.br/nfsevia cMotivo" json:"cMotivo"`
	XMotivo      *TSMotivo      `xml:"http://www.sped.fazenda.gov.br/nfsevia xMotivo,omitempty" json:"xMotivo,omitempty"`
	ChSubstituta TSChaveNFSeVia `xml:"http://www.sped.fazenda.gov.br/nfsevia chSubstituta" json:"chSubstituta"`
}

type TE203202 struct {
	XDesc        string `xml:"http://www.sped.fazenda.gov.br/nfsevia xDesc" json:"xDesc"`
	TpNi         string `xml:"http://www.sped.fazenda.gov.br/nfsevia tpNi" json:"tpNi"`
	CNPJCPF      string `xml:"http://www.sped.fazenda.gov.br/nfsevia CNPJCPF" json:"CNPJCPF"`
	XRazaoSocial string `xml:"http://www.sped.fazenda.gov.br/nfsevia xRazaoSocial" json:"xRazaoSocial"`
}

type TE203204 struct {
	XDesc              string                   `xml:"http://www.sped.fazenda.gov.br/nfsevia xDesc" json:"xDesc"`
	TpNi               string                   `xml:"http://www.sped.fazenda.gov.br/nfsevia tpNi" json:"tpNi"`
	CNPJCPF            string                   `xml:"http://www.sped.fazenda.gov.br/nfsevia CNPJCPF" json:"CNPJCPF"`
	XRazaoSocial       string                   `xml:"http://www.sped.fazenda.gov.br/nfsevia xRazaoSocial" json:"xRazaoSocial"`
	CCodMotivoRejeicao TSCodMotivoRejeicaoManif `xml:"http://www.sped.fazenda.gov.br/nfsevia cCodMotivoRejeicao" json:"cCodMotivoRejeicao"`
	XMotivoRejeita     *string                  `xml:"http://www.sped.fazenda.gov.br/nfsevia xMotivoRejeita,omitempty" json:"xMotivoRejeita,omitempty"`
}

type TSAliqCOFINSDesc2V2 string

type TSAliqDesc2V2 string

type TSAliqEfetMunDesc2V2 string

type TSAliqEfetUFDesc2V2 string

type TSAliqPISDesc2V2 string

type TSBCDesc15V2 string

type TSBCIBSCBSDesc15V2 string

type TSBCPISCOFINSDesc15V2 string

type TSCNPJ string

type TSCNPJFornec string

type TSCPF string

type TSCategoriaVeiculo string

type TSChaveNFSeVia string

type TSCodAcessoPedagio string

type TSCodCST string

type TSCodClassTrib string

type TSCodContrato string

type TSCodJustCanc string

type TSCodJustSubst string

type TSCodMotivoRejeicaoManif string

type TSCodMunIBGE string

type TSCodNaoNIF string

type TSCodTribNac string

type TSCOFINSDesc15V2 string

type TSCompetenciaDateTimeUTC string

type TSData string

type TSDateTimeUTC string

type TSDec15V2 string

type TSDec2V2 string

type TSDec3V2 string

type TSDec7V3 string

type TSDec7V6 string

type TSDesc150 string

type TSDesc2000 string

type TSDesc600 string

type TSDescCondDesc15V2 string

type TSDescIncondDesc15V2 string

type TSDescricaoMotivoSubstituicao string

type TSEmissaoDateTimeUTC string

type TSIBSTotDesc15V2 string

type TSISSQNDesc15V2 string

type TSIDEvento string

type TSIDNFSeVia string

type TSIndOperacao string

type TSKmPracaEmi string

type TSLiqIBSCBSDesc15V2 string

type TSLocEmiCodMunIBGE string

type TSLocEmiDesc150 string

type TSLocOperCodMunIBGE string

type TSLocOperDesc150 string

type TSModPagamento string

type TSModPassagem string

type TSMotivo string

type TSMotivoSubstituicao string

type TSMunFornec string

type TSNBS string

type TSNIF string

type TSNomeFornecDesc150 string

type TSNum3Dig string

type TSNumNFSeVia string

type TSNumProcExigSuspensa string

type TSNumProtocolo string

type TSOpExigSuspensa string

type TSPercCBSDesc2V2 string

type TSPISDesc15V2 string

type TSPlaca string

type TSPracaEmiCodigo string

type TSPracaEmiDesc150 string

type TSRetCPDesc15V2 string

type TSRetCSLLDesc15V2 string

type TSRetIRRFDesc15V2 string

type TSRodagem string

type TSSentido string

type TSSerie string

type TSServDesc15V2 string

type TSString string

type TSStringComQuebraDeLinha string

type TSTipoAmbiente string

type TSTipoCST string

type TSTipoEventoVia string

type TSTipoRetPISCOFINS string

type TSTotNFIBSCBSDesc15V2 string

type TSTribISSQN string

type TSValorCBSDesc15V2 string

type TSValorIBSMunDesc15V2 string

type TSValorIBSUFDesc15V2 string

type TSVerNFSeVia string

type TSVersaoEmissor string

type TStat string

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

type String string

// EventoViaProc represents the global element {http://www.sped.fazenda.gov.br/nfsevia}eventoViaProc.
type EventoViaProc struct {
	XMLName xml.Name `xml:"http://www.sped.fazenda.gov.br/nfsevia eventoViaProc" json:"-"`
	TCEventoViaProc
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
