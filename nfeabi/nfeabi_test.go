package nfeabi

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNFeABI_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195770010000000011000000021")
	doc := NewNFeABI(chave)

	doc.InfNFeABI.Imovel = Imovel{
		Matricula: "123.456",
		Cartorio:  "1º Cartório de Registro de Imóveis de SP",
		CIB:       "BR-87654321",
		XLgr:      "Av. Paulista",
		Nro:       "1000",
		XBairro:   "Bela Vista",
		CMun:      "3550308",
		UF:        "SP",
	}

	doc.InfNFeABI.Operacao = Operacao{
		TpOperacao: "1",
		VOperacao:  nfe.NewDecimal("1200000.00"),
		IBSCBS: &IBSCBSNFeABI{
			CST:        "000",
			CClassTrib: "000001",
			VBC:        nfe.NewDecimal("1200000.00"),
			PIBS:       nfe.NewDecimal("0.10"),
			VIBS:       nfe.NewDecimal("1200.00"),
			PCBS:       nfe.NewDecimal("0.90"),
			VCBS:       nfe.NewDecimal("10800.00"),
		},
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar NF-e ABI: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<NFeABI") || !strings.Contains(xmlStr, "<matricula>123.456</matricula>") {
		t.Errorf("XML da NF-e ABI deve conter matrícula do imóvel: %s", xmlStr)
	}
}
