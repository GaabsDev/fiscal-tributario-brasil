package nfcom

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNFCom_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195620010000000011000000016")
	doc := NewNFCom(chave)

	doc.InfNFCom.Det = append(doc.InfNFCom.Det, Det{
		NItem: 1,
		Prod: Prod{
			CProd:  "INT-100MB",
			XProd:  "Banda Larga Fibra 100MB",
			CFOP:   "5303",
			UCom:   "UN",
			QCom:   nfe.NewDecimal("1.0000"),
			VUnCom: nfe.NewDecimal("99.90"),
			VProd:  nfe.NewDecimal("99.90"),
		},
		Imp: Imp{
			IBSCBS: &IBSCBSNFCom{
				CST:        "000",
				CClassTrib: "000001",
				VBC:        nfe.NewDecimal("99.90"),
				PIBS:       nfe.NewDecimal("0.10"),
				VIBS:       nfe.NewDecimal("0.10"),
				PCBS:       nfe.NewDecimal("0.90"),
				VCBS:       nfe.NewDecimal("0.90"),
			},
		},
	})

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar NFCom: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<NFCom") || !strings.Contains(xmlStr, "<cProd>INT-100MB</cProd>") {
		t.Errorf("XML da NFCom deve conter tag raiz e item: %s", xmlStr)
	}
}
