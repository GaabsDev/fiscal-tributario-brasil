// Package nfce implementa a Nota Fiscal de Consumidor Eletrônica (NFC-e - Modelo 65).
// Leiaute 4.00; SetQRCode recebe a URL calculada externamente.
// A regra I05k-10 rejeita prod.tpCredPresIBSZFM no modelo 65.
package nfce

import (
	"encoding/xml"
	"fmt"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	ModeloNFCe       = "65"
	VersaoLayoutNFCe = "4.00"
)

// NFCe é o elemento raiz da Nota Fiscal de Consumidor Eletrônica.
type NFCe struct {
	XMLName    xml.Name        `xml:"http://www.portalfiscal.inf.br/nfe NFe" json:"-"`
	InfNFe     nfe.InfNFe      `xml:"infNFe" json:"infNFe"`
	InfNFeSupl *nfe.InfNFeSupl `xml:"infNFeSupl,omitempty" json:"infNFeSupl,omitempty"`
	Signature  *nfe.Signature  `xml:"Signature,omitempty" json:"signature,omitempty"`
}

// NewNFCe inicializa uma nova NFC-e garantindo modelo 65 e identificador oficial.
func NewNFCe(chave nfe.ChaveAcesso) *NFCe {
	id := fmt.Sprintf("NFe%s", string(chave))
	return &NFCe{
		InfNFe: nfe.InfNFe{
			Versao: VersaoLayoutNFCe,
			Id:     id,
			Ide: nfe.Ide{
				Mod:   ModeloNFCe,
				TpNF:  "1",
				TpImp: "4", // DANFE NFC-e
			},
			Det: make([]nfe.Det, 0),
			Pag: nfe.Pag{
				DetPag: make([]nfe.DetPag, 0),
			},
		},
	}
}

// SetEmitente define os dados do emitente da NFC-e.
func (doc *NFCe) SetEmitente(emit nfe.Emit) {
	doc.InfNFe.Emit = emit
}

// AddItem anexa um item de venda à NFC-e, calculando o sequencial nItem.
func (doc *NFCe) AddItem(det nfe.Det) {
	det.NItem = len(doc.InfNFe.Det) + 1
	doc.InfNFe.Det = append(doc.InfNFe.Det, det)
}

// SetQRCode define a URL e parâmetros do QR Code de consulta pública (infNFeSupl).
func (doc *NFCe) SetQRCode(qrCodeURL, urlChave string) {
	doc.InfNFeSupl = &nfe.InfNFeSupl{
		QrCode:   qrCodeURL,
		UrlChave: urlChave,
	}
}

// Validate executa a validação determinística de regras SEFAZ para NFC-e (modelo 65).
func (doc *NFCe) Validate() error {
	var errs nfe.ValidationErrors

	if doc.InfNFe.Ide.Mod != ModeloNFCe {
		errs.AddStructural("infNFe.ide.mod", "INVALID_MODEL", fmt.Sprintf("modelo inválido para NFC-e: esperado '%s', informado '%s'", ModeloNFCe, doc.InfNFe.Ide.Mod))
	}

	if len(doc.InfNFe.Det) == 0 {
		errs.AddStructural("infNFe.det", "MIN_OCCURS", "NFC-e deve conter pelo menos um item (det)")
	}

	// Regra I05k-10 (NT 2025.002 / Rejeição 1165): NFC-e não aceita crédito presumido de IBS ZFM
	for idx, item := range doc.InfNFe.Det {
		if item.Prod.TpCredPresIBSZFM != "" {
			errs.AddBusinessRule("infNFe.det.prod.tpCredPresIBSZFM", "I05k-10", "1165",
				fmt.Sprintf("Item %d: crédito presumido de IBS ZFM não é permitido para NFC-e (Modelo 65)", idx+1),
				"NT 2025.002", nfe.SeverityError)
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// ToXML serializa a NFC-e para XML compacto oficial.
func (doc *NFCe) ToXML() ([]byte, error) {
	return xml.Marshal(doc)
}

// ToXMLIndent serializa a NFC-e para XML formatado com indentação.
func (doc *NFCe) ToXMLIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(doc, prefix, indent)
}
