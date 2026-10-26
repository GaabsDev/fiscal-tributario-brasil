// Code generated from captured official XSDs. DO NOT EDIT.
package procnfag_v1_00

import (
	"context"
	"encoding/xml"
	"github.com/GaabsDev/fiscal-tributario-brasil/schemafiles/nfag"
)

// ValidateXML uses the unmodified captured official schema.
func ValidateXML(ctx context.Context, document []byte) error {
	return nfag.Validate(ctx, "PL_NFAg_NT2026.004 RTC_1.00/procNFAg_v1.00.xsd", document)
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

// CompatibilityID identifies the hash-bound local correction. Strict APIs remain unchanged.
const CompatibilityID = "nfag-nt2026.004-empty-choice/v1"

// ValidateXMLCompatible explicitly uses the documented local derived schema.
// Successful validation does not imply acceptance by the fiscal authority.
func ValidateXMLCompatible(ctx context.Context, document []byte) error {
	return nfag.ValidateCompatible(ctx, "PL_NFAg_NT2026.004 RTC_1.00/procNFAg_v1.00.xsd", document)
}

// ParseCompatible validates using the explicit local compatibility mode before decoding.
func ParseCompatible(ctx context.Context, document []byte, target any) error {
	if err := ValidateXMLCompatible(ctx, document); err != nil {
		return err
	}
	return xml.Unmarshal(document, target)
}

// MarshalCompatible serializes and validates using the explicit local compatibility mode.
func MarshalCompatible(ctx context.Context, document any) ([]byte, error) {
	data, err := xml.Marshal(document)
	if err != nil {
		return nil, err
	}
	if err := ValidateXMLCompatible(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}
