package nfeouro

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNFeOuro_And_DAO(t *testing.T) {
	ouro := NewNFeOuro("OURO-20261024-001")
	ouro.InfOuro.InstFin = InstFin{
		CNPJ:   "12345678000195",
		XNome:  "DTVM Distribuidora de Títulos e Valores Mobiliários S/A",
		CBacen: "12345",
	}
	ouro.InfOuro.Ouro = OuroDet{
		Teor:    nfe.NewDecimal("99.90"),
		PesoBr:  nfe.NewDecimal("100.00"),
		PesoLiq: nfe.NewDecimal("99.90"),
		Lacre:   "LACRE-999",
	}
	ouro.InfOuro.Valores = ValoresOuro{
		VGram:  nfe.NewDecimal("450.00"),
		VTotal: nfe.NewDecimal("44955.00"),
		IOF:    nfe.NewDecimal("449.55"),
	}

	xmlBytes, err := ouro.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar NF-e Ouro: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<NFeOuro") || !strings.Contains(xmlStr, "<teor>99.90</teor>") {
		t.Errorf("XML da NF-e Ouro inválido: %s", xmlStr)
	}
}
