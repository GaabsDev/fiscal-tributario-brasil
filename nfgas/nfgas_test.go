package nfgas

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNFGas_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195760010000000011000000020")
	doc := NewNFGas(chave)

	doc.InfNFGas.GMed = append(doc.InfNFGas.GMed, GMed{
		NMed:    "MEDGAS-77",
		DMedAnt: nfe.NewDate(2026, 9, 10),
		DMedAct: nfe.NewDate(2026, 10, 10),
		ConsM3:  nfe.NewDecimal("125.40"),
		FatorPT: nfe.NewDecimal("1.0250"),
	})

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar NFGas: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<NFGas") || !strings.Contains(xmlStr, "<fatorPT>1.0250</fatorPT>") {
		t.Errorf("XML da NFGas deve conter medição com fator PT: %s", xmlStr)
	}
}
