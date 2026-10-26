package nfse

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestNFSe_And_DPS(t *testing.T) {
	dps := NewDPS("DPS123456789")
	dps.InfDPS.Prest = Prestador{
		CNPJ: "12345678000195",
	}
	dps.InfDPS.Serv = Servico{
		CLocPrestacao: "3550308",
		CServ: CServ{
			CTribNac:   "010101",
			CClassTrib: "000001",
			XDescServ:  "Serviços de consultoria em TI",
		},
	}
	dps.InfDPS.Valores = ValoresDPS{
		VServPrest: VServPrest{
			VServ: nfe.NewDecimal("5000.00"),
		},
		Trib: TributosDPS{
			ISSQN: ISSQNDPS{
				CSTISSQN: "1",
				VBCISSQN: nfe.NewDecimal("5000.00"),
				PAliq:    nfe.NewDecimal("5.00"),
				VISSQN:   nfe.NewDecimal("250.00"),
			},
			IBSCBS: &IBSCBSDPS{
				VBC:  nfe.NewDecimal("5000.00"),
				VIBS: nfe.NewDecimal("5.00"),
				VCBS: nfe.NewDecimal("45.00"),
			},
		},
	}

	xmlBytes, err := dps.ToXMLIndent("", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar DPS: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<DPS") || !strings.Contains(xmlStr, "<cTribNac>010101</cTribNac>") {
		t.Errorf("XML da DPS inválido: %s", xmlStr)
	}
}
