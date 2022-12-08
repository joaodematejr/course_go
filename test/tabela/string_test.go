package strings

const msgIndex = "%s (parte: %s) - Indice: esperando (%d), <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto  string
		parte  string
		indice int
	}{
		{"Code.education Rocks!", "Code", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"João", "0", 1},

	for _, tt := range testes {
		t.Logf("Massa: %v", tt)
		atual := strings.Index(tt.texto, tt.parte)

		if atual != tt.indice {
			t.Errorf(msgIndex, tt.texto, tt.parte, tt.indice, atual)
		}

	}
}
