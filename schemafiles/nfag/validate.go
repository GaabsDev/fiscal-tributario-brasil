// Code generated from captured official XSDs. DO NOT EDIT.
package nfag

import (
	"bytes"
	"context"
	"crypto/sha256"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed all:data
var schemas embed.FS
var allowedSchemas = map[string]bool{"PL_NFAg_NT2026.004 RTC_1.00/consSitNFAg_v1.00.xsd": true,
	"PL_NFAg_NT2026.004 RTC_1.00/consStatServNFAg_v1.00.xsd":    true,
	"PL_NFAg_NT2026.004 RTC_1.00/evCancNFAg_v1.00.xsd":          true,
	"PL_NFAg_NT2026.004 RTC_1.00/evCancVincPgto_v1.00.xsd":      true,
	"PL_NFAg_NT2026.004 RTC_1.00/evVincPgto_v1.00.xsd":          true,
	"PL_NFAg_NT2026.004 RTC_1.00/eventoNFAg_v1.00.xsd":          true,
	"PL_NFAg_NT2026.004 RTC_1.00/nfag_v1.00.xsd":                true,
	"PL_NFAg_NT2026.004 RTC_1.00/procEventoNFAg_v1.00.xsd":      true,
	"PL_NFAg_NT2026.004 RTC_1.00/procNFAg_v1.00.xsd":            true,
	"PL_NFAg_NT2026.004 RTC_1.00/retConsSitNFAg_v1.00.xsd":      true,
	"PL_NFAg_NT2026.004 RTC_1.00/retConsStatServNFAg_v1.00.xsd": true,
	"PL_NFAg_NT2026.004 RTC_1.00/retEventoNFAg_v1.00.xsd":       true,
	"PL_NFAg_NT2026.004 RTC_1.00/retNFAg_v1.00.xsd":             true,
}

func digest(data []byte) string { return fmt.Sprintf("%x", sha256.Sum256(data)) }

// Validate uses the captured official schema. Compatibility is never selected implicitly.
func Validate(ctx context.Context, entry string, document []byte) error {
	return validate(ctx, entry, document, false)
}
func validate(ctx context.Context, entry string, document []byte, compatible bool) error {
	if !allowedSchemas[entry] {
		return fmt.Errorf("schema não registrado: %s", entry)
	}
	if len(document) > 50*1024*1024 {
		return fmt.Errorf("XML excede 50 MiB")
	}
	upper := bytes.ToUpper(document)
	if bytes.Contains(upper, []byte("<!DOCTYPE")) || bytes.Contains(upper, []byte("<!ENTITY")) {
		return fmt.Errorf("DTD e declarações de entidades não são aceitas")
	}
	dir, err := os.MkdirTemp("", "fiscal-tributario-brasil-xsd-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	if err := fs.WalkDir(schemas, "data", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		content, err := schemas.ReadFile(path)
		if err != nil {
			return err
		}
		target := filepath.Join(dir, filepath.FromSlash(path))
		if err := os.MkdirAll(filepath.Dir(target), 0700); err != nil {
			return err
		}
		return os.WriteFile(target, content, 0600)
	}); err != nil {
		return err
	}
	if compatible {
		if err := applyCompatibility(dir); err != nil {
			return err
		}
	}
	cmd := exec.CommandContext(ctx, "xmllint", "--nonet", "--noout", "--schema", filepath.Join(dir, "data", filepath.FromSlash(entry)), "-")
	cmd.Stdin = bytes.NewReader(document)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("validação XSD falhou: %w: %s", err, output)
	}
	return nil
}

//go:embed all:compatibility
var compatibilitySchemas embed.FS
var overlayFiles = []struct{ path, officialSHA, derivedSHA string }{{"PL_NFAg_NT2026.004 RTC_1.00/nfagTiposBasico_v1.00.xsd", "e72acd32937077c33831aaaf87dae91f96fef65206accadb4eee9f948826e7a4", "dd30fd9d456c28d92b2ca73c5d12d1588a13e2ffcfa3374da7424257d4fb8b9b"}}

func applyCompatibility(dir string) error {
	for _, overlay := range overlayFiles {
		original, err := schemas.ReadFile("data/" + overlay.path)
		if err != nil {
			return err
		}
		if digest(original) != overlay.officialSHA {
			return fmt.Errorf("schema oficial fora do escopo da compatibilidade: %s", overlay.path)
		}
		derived, err := compatibilitySchemas.ReadFile("compatibility/" + overlay.path)
		if err != nil {
			return err
		}
		if digest(derived) != overlay.derivedSHA {
			return fmt.Errorf("schema derivado corrompido: %s", overlay.path)
		}
		if err := os.WriteFile(filepath.Join(dir, "data", filepath.FromSlash(overlay.path)), derived, 0600); err != nil {
			return err
		}
	}
	return nil
}

// ValidateCompatible explicitly selects the documented local compatibility overlay.
// It does not certify conformity with the unchanged official schema or authorization.
func ValidateCompatible(ctx context.Context, entry string, document []byte) error {
	return validate(ctx, entry, document, true)
}
