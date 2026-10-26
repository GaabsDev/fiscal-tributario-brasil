// Code generated from captured official XSDs. DO NOT EDIT.
package mdfe

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
var allowedSchemas = map[string]bool{"PL_MDFe_300b_NT012025_1.05/consMDFeNaoEnc_v3.00.xsd": true,
	"PL_MDFe_300b_NT012025_1.05/consReciMDFe_v3.00.xsd":             true,
	"PL_MDFe_300b_NT012025_1.05/consSitMDFe_v3.00.xsd":              true,
	"PL_MDFe_300b_NT012025_1.05/consStatServMDFe_v3.00.xsd":         true,
	"PL_MDFe_300b_NT012025_1.05/distMDFe_v3.00.xsd":                 true,
	"PL_MDFe_300b_NT012025_1.05/enviMDFe_v3.00.xsd":                 true,
	"PL_MDFe_300b_NT012025_1.05/evAlteracaoPagtoServMDFe_v3.00.xsd": true,
	"PL_MDFe_300b_NT012025_1.05/evCancMDFe_v3.00.xsd":               true,
	"PL_MDFe_300b_NT012025_1.05/evConfirmaServMDFe_v3.00.xsd":       true,
	"PL_MDFe_300b_NT012025_1.05/evEncMDFe_v3.00.xsd":                true,
	"PL_MDFe_300b_NT012025_1.05/evIncCondutorMDFe_v3.00.xsd":        true,
	"PL_MDFe_300b_NT012025_1.05/evInclusaoDFeMDFe_v3.00.xsd":        true,
	"PL_MDFe_300b_NT012025_1.05/evPagtoOperMDFe_v3.00.xsd":          true,
	"PL_MDFe_300b_NT012025_1.05/eventoMDFe_v3.00.xsd":               true,
	"PL_MDFe_300b_NT012025_1.05/mdfeConsultaDFe_v3.00.xsd":          true,
	"PL_MDFe_300b_NT012025_1.05/mdfeModalAereo_v3.00.xsd":           true,
	"PL_MDFe_300b_NT012025_1.05/mdfeModalAquaviario_v3.00.xsd":      true,
	"PL_MDFe_300b_NT012025_1.05/mdfeModalFerroviario_v3.00.xsd":     true,
	"PL_MDFe_300b_NT012025_1.05/mdfeModalRodoviario_v3.00.xsd":      true,
	"PL_MDFe_300b_NT012025_1.05/mdfe_v3.00.xsd":                     true,
	"PL_MDFe_300b_NT012025_1.05/procEventoMDFe_v3.00.xsd":           true,
	"PL_MDFe_300b_NT012025_1.05/procMDFe_v3.00.xsd":                 true,
	"PL_MDFe_300b_NT012025_1.05/retConsMDFeNaoEnc_v3.00.xsd":        true,
	"PL_MDFe_300b_NT012025_1.05/retConsReciMDFe_v3.00.xsd":          true,
	"PL_MDFe_300b_NT012025_1.05/retConsSitMDFe_v3.00.xsd":           true,
	"PL_MDFe_300b_NT012025_1.05/retConsStatServMDFe_v3.00.xsd":      true,
	"PL_MDFe_300b_NT012025_1.05/retDistMDFe_v3.00.xsd":              true,
	"PL_MDFe_300b_NT012025_1.05/retEnviMDFe_v3.00.xsd":              true,
	"PL_MDFe_300b_NT012025_1.05/retEventoMDFe_v3.00.xsd":            true,
	"PL_MDFe_300b_NT012025_1.05/retMDFeConsultaDFe_v3.00.xsd":       true,
	"PL_MDFe_300b_NT012025_1.05/retMDFe_v3.00.xsd":                  true,
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
