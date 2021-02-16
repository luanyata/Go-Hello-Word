package structmethodinterface

import "testing"

func TestPerimeter(t *testing.T) {

	checkResult := func(t *testing.T, result, expected float64) {
		t.Helper()
		if result != expected {
			t.Errorf("Expected: %.2f, result: %.2f ", result, expected)
		}
	}

	t.Run("Perimetro de um retangulo", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		result := Perimeter(rectangle)
		expected := 40.0

		checkResult(t, result, expected)
	})

	t.Run("Retângulo", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		result := rectangle.Area()
		expected := 72.0

		checkResult(t, result, expected)

	})

	t.Run("Circulo", func(t *testing.T) {
		circle := Circle{10}
		result := circle.Area()
		expected := 314.1592653589793

		checkResult(t, result, expected)
	})

}
