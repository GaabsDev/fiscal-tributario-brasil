// Code generated from captured official XSDs. DO NOT EDIT.
package evregmultimodal_v4_00

import (
	"context"
	"encoding/xml"
	"github.com/GaabsDev/fiscal-tributario-brasil/schemafiles/gtve"
)

// ValidateXML uses the unmodified captured official schema.
func ValidateXML(ctx context.Context, document []byte) error {
	return gtve.Validate(ctx, "PL_CTe_400_NT2026.004 RTC_1.00/evRegMultimodal_v4.00.xsd", document)
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
