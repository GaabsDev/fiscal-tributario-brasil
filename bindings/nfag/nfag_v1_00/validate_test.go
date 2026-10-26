package nfag_v1_00

import (
	"context"
	"strings"
	"testing"
)

func TestCapturedSchemaRejectsUnknownRoot(t *testing.T) {
	err := ValidateXML(context.Background(), []byte("<invalid/>"))
	if err == nil || !strings.Contains(err.Error(), "Schemas validity error") {
		t.Fatalf("schema não validou o documento de controle: %v", err)
	}

	if err := ValidateXMLCompatible(context.Background(), []byte("<invalid/>")); err == nil || !strings.Contains(err.Error(), "Schemas validity error") {
		t.Fatalf("compatibilidade aceitou raiz inválida: %v", err)
	}
}
