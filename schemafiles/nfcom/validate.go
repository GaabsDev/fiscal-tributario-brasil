// Code generated from captured official XSDs. DO NOT EDIT.
package nfcom

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
var allowedSchemas = map[string]bool{"PL_NFCOM_1.00_NT2026.004 RTC_1.00/consSitNFCom_v1.00.xsd": true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/consStatServNFCom_v1.00.xsd":    true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/evCancNFCom_v1.00.xsd":          true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/evCancVincPgto_v1.00.xsd":       true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/evVincPgto_v1.00.xsd":           true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/eventoNFCom_v1.00.xsd":          true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/nfcom_v1.00.xsd":                true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/procEventoNFCom_v1.00.xsd":      true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/procNFCom_v1.00.xsd":            true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/retConsSitNFCom_v1.00.xsd":      true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/retConsStatServNFCom_v1.00.xsd": true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/retEventoNFCom_v1.00.xsd":       true,
	"PL_NFCOM_1.00_NT2026.004 RTC_1.00/retNFCom_v1.00.xsd":             true,
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
func applyCompatibility(dir string) error {
	return fmt.Errorf("nenhuma compatibilidade registrada para este perfil")
}
