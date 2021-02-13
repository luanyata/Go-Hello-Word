package helloworld

import "testing"

func TestHello(t *testing.T) {

	checkMessage := func(t *testing.T, result string, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("Result '%s', expecteded '%s'", result, expected)
		}
	}

	t.Run("Diz Olá para as pessoas", func(t *testing.T) {
		result := Hello("Luan", "")
		expected := "Olá Luan"

		checkMessage(t, result, expected)

	})

	t.Run("Diz 'Olá mundo' quando uma string vazia for passada", func(t *testing.T) {
		result := Hello("", "")
		expected := "Olá mundo"

		checkMessage(t, result, expected)
	})

	t.Run("Em Espanhol", func(t *testing.T) {
		result := Hello("Fernanda", "es")
		expected := "Hola Fernanda"

		checkMessage(t, result, expected)
	})

	t.Run("Em Frances", func(t *testing.T) {
		result := Hello("Lucky", "fr")
		expected := "Bonjour Lucky"

		checkMessage(t, result, expected)
	})
}
