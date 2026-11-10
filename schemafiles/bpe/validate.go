// Code generated from captured official XSDs. DO NOT EDIT.
package bpe

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
var allowedSchemas = map[string]bool{"PL_BPe_100b_NT2026.004 RTC 1.01/bpeTA_v1.00.xsd": true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/bpeTM_v1.00.xsd":               true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/bpe_v1.00.xsd":                 true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/consSitBPe_v1.00.xsd":          true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/consStatServBPe_v1.00.xsd":     true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/evAlteracaoPoltrona_v1.00.xsd": true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/evCancBPe_v1.00.xsd":           true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/evCancVincPgto_v1.00.xsd":      true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/evExcessoBagagem_v1.00.xsd":    true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/evNaoEmbBPe_v1.00.xsd":         true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/evVincPgto_v1.00.xsd":          true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/eventoBPe_v1.00.xsd":           true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/procBPeTM_v1.00.xsd":           true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/procBPe_v1.00.xsd":             true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/procEventoBPe_v1.00.xsd":       true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/retBPe_v1.00.xsd":              true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/retConsSitBPe_v1.00.xsd":       true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/retConsStatServBPe_v1.00.xsd":  true,
	"PL_BPe_100b_NT2026.004 RTC 1.01/retEventoBPe_v1.00.xsd":        true,
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
