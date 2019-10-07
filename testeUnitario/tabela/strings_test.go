package strings

import (
	"strings"
	"testing"
)

const msg = "%s (parte %s) - indices: esperado (%d) <> encontrado (%d)."

func TextStrings(t *testing.T) {

	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Golang é incrivel", "Golang", 0},
		{"", "", 0},
		{"Teste", "teste", -1},
		{"Joao", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)

		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msg, teste.texto, teste.parte, teste.esperado, atual)
		}
	}

}
