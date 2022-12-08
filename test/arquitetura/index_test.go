package arquitetura

func TestDependente(t *testing.T) {
	// t.Fatal("not implemented")
	if runtime.GOOS == "windows" {
		t.Skip("Não funciona no windows")
	}

	t.Fail()

}
