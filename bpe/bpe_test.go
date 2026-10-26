package bpe

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestBPe_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195630010000000011000000015")
	doc := NewBPe(chave)

	doc.InfBPe.Passag = Passag{
		XNome: "Fulano de Tal",
	}

	doc.InfBPe.Imp = Imp{
		IBSCBS: &IBSCBSBPe{
			CST:        "000",
			CClassTrib: "000001",
			VBC:        nfe.NewDecimal("150.00"),
			PIBS:       nfe.NewDecimal("0.10"),
			VIBS:       nfe.NewDecimal("0.15"),
			PCBS:       nfe.NewDecimal("0.90"),
			VCBS:       nfe.NewDecimal("1.35"),
		},
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar BP-e: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<BPe") || !strings.Contains(xmlStr, "<cClassTrib>000001</cClassTrib>") {
		t.Errorf("XML do BP-e deve conter tag raiz e cClassTrib: %s", xmlStr)
	}
}
