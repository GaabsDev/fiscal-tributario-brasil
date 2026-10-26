package mdfe

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestMDFe_Modais_RodoEAereo(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195580010000000011000000014")
	doc := NewMDFe(chave)

	// Modal Aéreo
	doc.InfMDFe.Ide.Modal = "2"
	doc.InfMDFe.InfModal = InfModal{
		VersaoModal: "3.00",
		Aereo: &Aereo{
			Nac:     "BR",
			Matr:    "PTXYZ",
			NVoo:    "AZ1234",
			CAerEmb: "GRU",
			CAerDes: "BSB",
			DVoo:    nfe.NewDate(2026, 10, 24),
		},
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar MDF-e Aéreo: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<aereo>") || !strings.Contains(xmlStr, "<nVoo>AZ1234</nVoo>") {
		t.Errorf("XML do MDF-e deve conter nó modal aéreo: %s", xmlStr)
	}
}
