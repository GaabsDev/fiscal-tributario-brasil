package nfag

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNFAg_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195750010000000011000000019")
	doc := NewNFAg(chave)

	doc.InfNFAg.GHidro = append(doc.InfNFAg.GHidro, GHidro{
		NHidro:  "HIDRO-4321",
		DMedAnt: nfe.NewDate(2026, 9, 15),
		DMedAct: nfe.NewDate(2026, 10, 15),
		ConsM3:  nfe.NewDecimal("18.5"),
	})

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar NFAg: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<NFAg") || !strings.Contains(xmlStr, "<nHidro>HIDRO-4321</nHidro>") {
		t.Errorf("XML da NFAg deve conter hidrômetro: %s", xmlStr)
	}
}
