package main

import "testing"

func TestHello(t *testing.T) {

	checkMessage := func(t *testing.T, result, expect string) {
		t.Helper()
		if result != expect {
			t.Errorf("Result '%s', expected '%s'", result, expect)
		}
	}

	t.Run("Diz Olá para as pessoas", func(t *testing.T) {
		result := Hello("Luan")
		expect := "Olá Luan"

		checkMessage(t, result, expect)

	})

	t.Run("Diz 'Olá mundo' quando uma string vazia for passada", func(t *testing.T) {
		result := Hello("")
		expect := "Olá mundo"

		checkMessage(t, result, expect)
	})
}
