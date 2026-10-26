package cte

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestCTe_57_ModaisAndRTC(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195570010000000011000000011")
	doc := NewCTe(chave)

	if doc.InfCte.Ide.Mod != "57" {
		t.Fatalf("esperado modelo 57, obtido %s", doc.InfCte.Ide.Mod)
	}

	doc.InfCte.InfCTeNorm = &InfCTeNorm{
		InfCarga: InfCarga{
			VCarga:  nfe.NewDecimal("15000.00"),
			ProPred: "Carga Geral",
		},
		InfModal: InfModal{
			VersaoModal: "4.00",
			Rodo: &Rodo{
				RNTRC: "12345678",
			},
		},
	}

	doc.InfCte.Imp = Imp{
		ICMS: ICMS{
			ICMS00: &ICMS00{
				CST:   "00",
				VBC:   nfe.NewDecimal("1000.00"),
				PICMS: nfe.NewDecimal("12.00"),
				VICMS: nfe.NewDecimal("120.00"),
			},
		},
		IBSCBS: &IBSCBSCTe{
			CST:        "000",
			CClassTrib: "000001",
			VBC:        nfe.NewDecimal("1000.00"),
			PIBS:       nfe.NewDecimal("0.10"),
			VIBS:       nfe.NewDecimal("1.00"),
			PCBS:       nfe.NewDecimal("0.90"),
			VCBS:       nfe.NewDecimal("9.00"),
		},
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar CT-e: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<RNTRC>12345678</RNTRC>") {
		t.Errorf("XML do CT-e deve conter RNTRC: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<IBSCBS>") || !strings.Contains(xmlStr, "<cClassTrib>000001</cClassTrib>") {
		t.Errorf("XML do CT-e deve conter grupo IBSCBS da Reforma Tributária: %s", xmlStr)
	}
}

func TestCTeOS_67_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195670010000000011000000012")
	doc := NewCTeOS(chave)

	if doc.InfCte.Ide.Mod != "67" {
		t.Fatalf("esperado modelo 67, obtido %s", doc.InfCte.Ide.Mod)
	}

	doc.InfCte.InfCTeNorm = InfCTeNormOS{
		InfServico: InfServico{
			XDescServ: "Transporte Intermunicipal de Passageiros",
			QPassag:   42,
		},
		InfModal: InfModalOS{
			VersaoModal: "4.00",
			RodoOS: &RodoOS{
				TAF: "123456",
			},
		},
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar CT-e OS: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<CTeOS") {
		t.Errorf("XML deve conter tag raiz CTeOS: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<TAF>123456</TAF>") {
		t.Errorf("XML deve conter TAF do modal rodoviário OS: %s", xmlStr)
	}
}

func TestGTVe_64_Creation(t *testing.T) {
	chave := nfe.ChaveAcesso("35260112345678000195640010000000011000000013")
	doc := NewGTVe(chave)

	if doc.InfCte.Ide.Mod != "64" {
		t.Fatalf("esperado modelo 64, obtido %s", doc.InfCte.Ide.Mod)
	}

	doc.InfCte.InfGTVe = InfGTVe{
		InfEspecie: []InfEspecie{
			{
				TpEspecie: "1",
				VItem:     nfe.NewDecimal("50000.00"),
			},
		},
	}

	xmlBytes, err := doc.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar GTV-e: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<GTVe") {
		t.Errorf("XML deve conter tag raiz GTVe: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<tpEspecie>1</tpEspecie>") {
		t.Errorf("XML deve conter espécie de valor: %s", xmlStr)
	}
}
