// Code generated from captured official XSDs. DO NOT EDIT.
package cte

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
var allowedSchemas = map[string]bool{"PL_CTe_400_NT2026.004 RTC_1.01/GTVe_v4.00.xsd": true,
	"PL_CTe_400_NT2026.004 RTC_1.01/consSitCTe_v4.00.xsd":           true,
	"PL_CTe_400_NT2026.004 RTC_1.01/consStatServCTe_v4.00.xsd":      true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteModalAereo_v4.00.xsd":        true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteModalAquaviario_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteModalDutoviario_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteModalFerroviario_v4.00.xsd":  true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteModalRodoviarioOS_v4.00.xsd": true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteModalRodoviario_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteMultiModal_v4.00.xsd":        true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteOS_v4.00.xsd":                true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cteSimp_v4.00.xsd":              true,
	"PL_CTe_400_NT2026.004 RTC_1.01/cte_v4.00.xsd":                  true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evApropriaCredPres_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evCCeCTe_v4.00.xsd":             true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evCECTe_v4.00.xsd":              true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evCancCECTe_v4.00.xsd":          true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evCancCTe_v4.00.xsd":            true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evCancIECTe_v4.00.xsd":          true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evCancPrestDesacordo_v4.00.xsd": true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evCancVincPgto_v4.00.xsd":       true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evEPECCTe_v4.00.xsd":            true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evGTV_v4.00.xsd":                true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evIECTe_v4.00.xsd":              true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evPrestDesacordo_v4.00.xsd":     true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evRegMultimodal_v4.00.xsd":      true,
	"PL_CTe_400_NT2026.004 RTC_1.01/evVincPgto_v4.00.xsd":           true,
	"PL_CTe_400_NT2026.004 RTC_1.01/eventoCTe_v4.00.xsd":            true,
	"PL_CTe_400_NT2026.004 RTC_1.01/procCTeOS_v4.00.xsd":            true,
	"PL_CTe_400_NT2026.004 RTC_1.01/procCTeSimp_v4.00.xsd":          true,
	"PL_CTe_400_NT2026.004 RTC_1.01/procCTe_v4.00.xsd":              true,
	"PL_CTe_400_NT2026.004 RTC_1.01/procEventoCTe_v4.00.xsd":        true,
	"PL_CTe_400_NT2026.004 RTC_1.01/procGTVe_v4.00.xsd":             true,
	"PL_CTe_400_NT2026.004 RTC_1.01/retCTeOS_v4.00.xsd":             true,
	"PL_CTe_400_NT2026.004 RTC_1.01/retCTeSimp_v4.00.xsd":           true,
	"PL_CTe_400_NT2026.004 RTC_1.01/retCTe_v4.00.xsd":               true,
	"PL_CTe_400_NT2026.004 RTC_1.01/retConsSitCTe_v4.00.xsd":        true,
	"PL_CTe_400_NT2026.004 RTC_1.01/retConsStatServCTe_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.01/retEventoCTe_v4.00.xsd":         true,
	"PL_CTe_400_NT2026.004 RTC_1.01/retGTVe_v4.00.xsd":              true,
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
