// Package cclasstrib provê a Tabela Oficial de Classificação Tributária do IBS e CBS da Reforma Tributária (NT 2025.002 / NT 2026.007).
// Preserva códigos, descrições, aplicabilidade e datas do catálogo utilizado na geração.
package cclasstrib

import (
	"sort"
	"strings"
)

// Definition representa uma entrada oficial da Tabela de Classificação Tributária do IBS e CBS.
type Definition struct {
	ValidFrom     string          `json:"valid_from,omitempty"`
	ValidTo       string          `json:"valid_to,omitempty"`
	Code          string          `json:"code"`
	CST           string          `json:"cst"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	IsFuel        bool            `json:"is_fuel"`
	Applicability map[string]bool `json:"applicability"` // modelo fiscal -> aplicável
}

// Registry armazena o catálogo completo indexado pelo código cClassTrib de 6 dígitos.
var Registry = map[string]Definition{}

// IsValid verifica se o código cClassTrib existe na tabela oficial da SEFAZ/RFB.
func IsValid(code string) bool {
	c := strings.TrimSpace(code)
	_, ok := Registry[c]
	return ok
}

// IsApplicableToModel verifica se o código cClassTrib é aplicável ao modelo do documento fiscal informado.
func IsApplicableToModel(code, model string) bool {
	c := strings.TrimSpace(code)
	def, ok := Registry[c]
	if !ok {
		return false
	}
	return def.Applicability[strings.TrimSpace(model)]
}

// IsFuel indica se a classificação tributária se refere a operações com combustíveis (sujeita à regra LA01-30).
func IsFuel(code string) bool {
	c := strings.TrimSpace(code)
	def, ok := Registry[c]
	if !ok {
		return false
	}
	return def.IsFuel
}

// Get retorna a definição oficial do cClassTrib informado.
func Get(code string) (Definition, bool) {
	c := strings.TrimSpace(code)
	def, ok := Registry[c]
	return def, ok
}

// List retorna todas as definições oficiais ordenadas pelo código.
func List() []Definition {
	codes := make([]string, 0, len(Registry))
	for k := range Registry {
		codes = append(codes, k)
	}
	sort.Strings(codes)

	list := make([]Definition, 0, len(codes))
	for _, k := range codes {
		list = append(list, Registry[k])
	}
	return list
}
