package dere

import (
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestDeRE_ModelagemNormativa(t *testing.T) {
	doc := NewDeclaracao("12345678000195", "2026-10")
	doc.Regimes = append(doc.Regimes, ApuracaoRegime{
		Tipo:            RegimeCombustiveisMono,
		CClassTrib:      "000005",
		BaseCalculo:     nfe.NewDecimal("100000.00"),
		ValorApuradoIBS: nfe.NewDecimal("1500.00"),
		ValorApuradoCBS: nfe.NewDecimal("13500.00"),
	})
	doc.Totalizacao = TotalizacaoDeRE{
		TotalIBS: nfe.NewDecimal("1500.00"),
		TotalCBS: nfe.NewDecimal("13500.00"),
	}

	jsonData, err := doc.ToJSON()
	if err != nil {
		t.Fatalf("falha ao serializar DeRE para JSON: %v", err)
	}
	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, "COMBUSTIVEIS_MONOFASICO") || !strings.Contains(jsonStr, "Aguardando publicação pública do pacote XSD") {
		t.Errorf("DeRE JSON deve conter regime e aviso de regulamentação: %s", jsonStr)
	}
}
