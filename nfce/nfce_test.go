package nfce

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNFCe_CreationAndValidation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195650010000000011000000010")
	doc := NewNFCe(chave)

	if doc.InfNFe.Ide.Mod != "65" {
		t.Fatalf("esperado modelo 65, obtido: %s", doc.InfNFe.Ide.Mod)
	}

	doc.AddItem(nfe.Det{
		Prod: nfe.Prod{
			CProd:  "PROD-01",
			XProd:  "Produto Varejo",
			NCM:    "84713012",
			CFOP:   "5102",
			UCom:   "UN",
			QCom:   nfe.NewDecimal("1.0000"),
			VUnCom: nfe.NewDecimal("10.00"),
			VProd:  nfe.NewDecimal("10.00"),
		},
	})

	doc.SetQRCode("https://sefaz.sp.gov.br/nfce/qrcode?p=123", "https://sefaz.sp.gov.br/nfce/consulta")

	xmlData, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar NFC-e: %v", err)
	}

	xmlStr := string(xmlData)
	if !strings.Contains(xmlStr, "<infNFeSupl>") || !strings.Contains(xmlStr, "<qrCode>") {
		t.Errorf("XML da NFC-e deve conter nó suplementar do QR Code: %s", xmlStr)
	}
}

func TestNFCe_Regra_I05k10_Rejeicao1165(t *testing.T) {
	doc := NewNFCe("35260112345678000195650010000000011000000010")
	doc.AddItem(nfe.Det{
		Prod: nfe.Prod{
			TpCredPresIBSZFM: "1",
			CProd:            "PROD-ZFM",
			XProd:            "Item ZFM",
		},
		Imposto: nfe.Imposto{
			IBSCBS: &nfe.IBSCBS{
				CST:        "000",
				CClassTrib: "000001",
				GCredPresIBSZFM: &nfe.GCredPresIBSZFM{
					TpCredPresIBSZFM: "1",
					VCredPresIBSZFM:  nfe.NewDecimal("5.00"),
				},
			},
		},
	})

	err := doc.Validate()
	if err == nil {
		t.Fatalf("esperada Rejeição 1165 para NFC-e com tpCredPresIBSZFM")
	}
	if !strings.Contains(err.Error(), "1165") {
		t.Errorf("erro deve conter código 1165: %v", err)
	}
}
