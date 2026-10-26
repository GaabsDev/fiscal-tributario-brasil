// Code generated from captured official XSDs. DO NOT EDIT.
package gtve

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
var allowedSchemas = map[string]bool{"PL_CTe_400_NT2026.004 RTC_1.00/GTVe_v4.00.xsd": true,
	"PL_CTe_400_NT2026.004 RTC_1.00/consSitCTe_v4.00.xsd":           true,
	"PL_CTe_400_NT2026.004 RTC_1.00/consStatServCTe_v4.00.xsd":      true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteModalAereo_v4.00.xsd":        true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteModalAquaviario_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteModalDutoviario_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteModalFerroviario_v4.00.xsd":  true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteModalRodoviarioOS_v4.00.xsd": true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteModalRodoviario_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteMultiModal_v4.00.xsd":        true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteOS_v4.00.xsd":                true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cteSimp_v4.00.xsd":              true,
	"PL_CTe_400_NT2026.004 RTC_1.00/cte_v4.00.xsd":                  true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evApropriaCredPres_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evCCeCTe_v4.00.xsd":             true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evCECTe_v4.00.xsd":              true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evCancCECTe_v4.00.xsd":          true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evCancCTe_v4.00.xsd":            true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evCancIECTe_v4.00.xsd":          true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evCancPrestDesacordo_v4.00.xsd": true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evCancVincPgto_v4.00.xsd":       true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evEPECCTe_v4.00.xsd":            true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evGTV_v4.00.xsd":                true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evIECTe_v4.00.xsd":              true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evPrestDesacordo_v4.00.xsd":     true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evRegMultimodal_v4.00.xsd":      true,
	"PL_CTe_400_NT2026.004 RTC_1.00/evVincPgto_v4.00.xsd":           true,
	"PL_CTe_400_NT2026.004 RTC_1.00/eventoCTe_v4.00.xsd":            true,
	"PL_CTe_400_NT2026.004 RTC_1.00/procCTeOS_v4.00.xsd":            true,
	"PL_CTe_400_NT2026.004 RTC_1.00/procCTeSimp_v4.00.xsd":          true,
	"PL_CTe_400_NT2026.004 RTC_1.00/procCTe_v4.00.xsd":              true,
	"PL_CTe_400_NT2026.004 RTC_1.00/procEventoCTe_v4.00.xsd":        true,
	"PL_CTe_400_NT2026.004 RTC_1.00/procGTVe_v4.00.xsd":             true,
	"PL_CTe_400_NT2026.004 RTC_1.00/retCTeOS_v4.00.xsd":             true,
	"PL_CTe_400_NT2026.004 RTC_1.00/retCTeSimp_v4.00.xsd":           true,
	"PL_CTe_400_NT2026.004 RTC_1.00/retCTe_v4.00.xsd":               true,
	"PL_CTe_400_NT2026.004 RTC_1.00/retConsSitCTe_v4.00.xsd":        true,
	"PL_CTe_400_NT2026.004 RTC_1.00/retConsStatServCTe_v4.00.xsd":   true,
	"PL_CTe_400_NT2026.004 RTC_1.00/retEventoCTe_v4.00.xsd":         true,
	"PL_CTe_400_NT2026.004 RTC_1.00/retGTVe_v4.00.xsd":              true,
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
