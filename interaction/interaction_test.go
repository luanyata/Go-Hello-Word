package interaction

import "testing"

func TestIRepeat(t *testing.T) {
	result := Repeat("a", 5)
	expected := "aaaaaa"

	if result != expected {
		t.Errorf("Expected '%s', Result '%s'", expected, result)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 20)
	}
}
