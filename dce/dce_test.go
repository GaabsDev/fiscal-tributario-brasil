package dce

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestDCe_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195990010000000011000000018")
	doc := NewDCe(chave)

	doc.InfDCe.Det = append(doc.InfDCe.Det, Det{
		NItem: 1,
		Prod: Prod{
			CProd: "ENCOMENDA-01",
			XProd: "Livros e Documentos Pessoais",
			QCom:  nfe.NewDecimal("2.0000"),
			VProd: nfe.NewDecimal("150.00"),
		},
	})
	doc.InfDCe.Total = Total{
		VDCe: nfe.NewDecimal("150.00"),
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar DC-e: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<DCe") || !strings.Contains(xmlStr, "<xProd>Livros e Documentos Pessoais</xProd>") {
		t.Errorf("XML da DC-e inválido: %s", xmlStr)
	}
}
