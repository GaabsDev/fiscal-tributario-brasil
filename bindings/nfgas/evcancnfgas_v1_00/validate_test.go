package evcancnfgas_v1_00

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
}
