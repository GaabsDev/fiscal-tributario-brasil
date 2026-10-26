package conssitdce_v1_00

import (
	"context"
	"encoding/xml"
	"os"
	"testing"
)

func TestCapturedDocumentRoundTrip0(t *testing.T) {
	data, err := os.ReadFile("testdata/structural_0.xml")
	if err != nil {
		t.Fatal(err)
	}
	var doc ConsSitDCe
	if err := Parse(context.Background(), data, &doc); err != nil {
		t.Fatal(err)
	}
	if _, err := Marshal(context.Background(), doc); err != nil {
		data, _ := xml.Marshal(doc)
		t.Fatalf("%v\nXML:%s", err, data)
	}
}
