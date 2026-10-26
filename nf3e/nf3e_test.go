package nf3e

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNF3e_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195660010000000011000000017")
	doc := NewNF3e(chave)

	doc.InfNF3e.GMed = append(doc.InfNF3e.GMed, GMed{
		NMed:    "MED-987654",
		DMedAnt: nfe.NewDate(2026, 9, 20),
		DMedAct: nfe.NewDate(2026, 10, 20),
		ConsKWh: nfe.NewDecimal("350.50"),
	})

	doc.InfNF3e.Total = Total{
		VProd: nfe.NewDecimal("280.40"),
		VNF:   nfe.NewDecimal("280.40"),
		IBSCBS: &IBSCBSNF3e{
			CST:        "000",
			CClassTrib: "000001",
			VBC:        nfe.NewDecimal("280.40"),
			VIBS:       nfe.NewDecimal("2.80"),
			VCBS:       nfe.NewDecimal("25.24"),
		},
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar NF3e: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<NF3e") || !strings.Contains(xmlStr, "<consKWh>350.50</consKWh>") {
		t.Errorf("XML da NF3e deve conter medição de consumo: %s", xmlStr)
	}
}
