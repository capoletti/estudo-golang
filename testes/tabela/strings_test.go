package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"GoLang é show", "GoLang", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Nome", "o", 1},
	}

	for _, teste := range testes {
		t.Logf("Massa %v", teste)

		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
