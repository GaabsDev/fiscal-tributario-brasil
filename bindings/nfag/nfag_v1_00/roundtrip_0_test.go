package nfag_v1_00

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestExplicitCompatibilityRoundTrip0(t *testing.T) {
	data, err := os.ReadFile("testdata/structural_0.xml")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	originalError := ValidateXML(ctx, data)
	if originalError == nil || !strings.Contains(originalError.Error(), "Schemas validity error") || !strings.Contains(originalError.Error(), "det") {
		t.Fatalf("o defeito do original deve continuar explícito: %v", originalError)
	}
	var doc NFAg
	if err := ParseCompatible(ctx, data, &doc); err != nil {
		t.Fatal(err)
	}
	if _, err := MarshalCompatible(ctx, doc); err != nil {
		t.Fatal(err)
	}
	// An actual structural error must remain invalid even in compatibility mode.
	invalid := strings.ReplaceAll(strings.ReplaceAll(string(data), "<det ", "<invalidDet "), "</det>", "</invalidDet>")
	if invalid == string(data) {
		t.Fatal("controle de det não foi alterado")
	}
	if err := ValidateXMLCompatible(ctx, []byte(invalid)); err == nil || !strings.Contains(err.Error(), "Schemas validity error") {
		t.Fatalf("compatibilidade enfraqueceu a validação: %v", err)
	}
	// Compatibility must not mutate the embedded original for subsequent calls.
	if err := ValidateXML(ctx, data); err == nil {
		t.Fatal("validação oficial foi substituída silenciosamente")
	}
}
