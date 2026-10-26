// Code generated from captured official XSDs. DO NOT EDIT.
package evtarranjocreddest_v0_0_2

import (
	"context"
	"encoding/xml"
	"github.com/GaabsDev/fiscal-tributario-brasil/schemafiles/dere_previa"
)

// ValidateXML uses the unmodified captured official schema.
func ValidateXML(ctx context.Context, document []byte) error {
	return dere_previa.Validate(ctx, "evtArranjoCredDest-v0_0_2.xsd", document)
}

// Parse validates against the official schema before decoding XML.
func Parse(ctx context.Context, document []byte, target any) error {
	if err := ValidateXML(ctx, document); err != nil {
		return err
	}
	return xml.Unmarshal(document, target)
}

// Marshal serializes and validates against the official schema.
func Marshal(ctx context.Context, document any) ([]byte, error) {
	data, err := xml.Marshal(document)
	if err != nil {
		return nil, err
	}
	if err := ValidateXML(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}
