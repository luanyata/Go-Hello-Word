package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Luan")
	expect := "Olá Luan"

	if result != expect {
		t.Errorf("Result '%s', expected '%s'", result, expect)
	}
}
