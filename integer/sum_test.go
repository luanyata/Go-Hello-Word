package integer

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("Expected '%d', Result '%d'", expected, result)
	}
}
