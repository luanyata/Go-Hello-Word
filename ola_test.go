package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola()
	expect := "Olá Mundo"

	if result != expect {
		t.Errorf("Result '%s', expected '%s'", result, expect)
	}
}
