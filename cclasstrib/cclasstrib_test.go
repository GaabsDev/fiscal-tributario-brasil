package cclasstrib

import (
	"testing"
)

func TestCClassTrib_Registry(t *testing.T) {
	if len(Registry) != 173 {
		t.Fatalf("esperado 173 códigos no catálogo oficial, obtido %d", len(Registry))
	}

	// Testa código padrão 000001
	if !IsValid("000001") {
		t.Errorf("000001 deveria ser válido")
	}
	if !IsApplicableToModel("000001", "55") {
		t.Errorf("000001 deveria ser aplicável ao modelo 55")
	}
	if IsFuel("000001") {
		t.Errorf("000001 não é combustível")
	}

	// LA01-30: 410013 está explicitamente incluído na regra de combustível.
	if !IsValid("410013") {
		t.Errorf("410013 deveria ser válido")
	}
	if !IsFuel("410013") {
		t.Errorf("410013 deveria ser combustível")
	}

	// Testa código inexistente
	if IsValid("999999") {
		t.Errorf("999999 não deveria ser válido")
	}
	if IsApplicableToModel("999999", "55") {
		t.Errorf("999999 não deveria ser aplicável")
	}

	// Testa List()
	list := List()
	if len(list) != 173 {
		t.Fatalf("List() deveria retornar 173 itens, retornou %d", len(list))
	}
}
