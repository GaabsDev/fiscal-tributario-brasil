package nfsevia

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNFSeVia_Creation(t *testing.T) {
	doc := NewNFSeVia("VIA-20261024-001")
	doc.InfVia.Concess = Concessionaria{
		CNPJ:  "12345678000195",
		XNome: "Concessionária Rodovias do Tietê S/A",
	}
	doc.InfVia.Praca = Praca{
		CPraca: "01",
		XPraca: "Praça Itupeva",
		KM:     "72",
		Rod:    "SP-348",
	}
	doc.InfVia.Passag = Passagem{
		Sentido: "1",
		CatVeic: "1", // Passeio
	}
	doc.InfVia.Valores = ValoresVia{
		VTarifa: nfe.NewDecimal("11.80"),
		IBSCBS: &IBSCBSVia{
			CST:        "000",
			CClassTrib: "000002", // Exploração de via
			VBC:        nfe.NewDecimal("11.80"),
			PIBS:       nfe.NewDecimal("0.10"),
			VIBS:       nfe.NewDecimal("0.01"),
			PCBS:       nfe.NewDecimal("0.90"),
			VCBS:       nfe.NewDecimal("0.11"),
		},
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar NFS-e Via: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<NFSeVia") || !strings.Contains(xmlStr, "<cPraca>01</cPraca>") {
		t.Errorf("XML da NFS-e Via inválido: %s", xmlStr)
	}
}
