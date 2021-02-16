package structmethodinterface

import "testing"

func TestGeometry(t *testing.T) {

	checkResult := func(t *testing.T, result, expected float64) {
		t.Helper()
		if result != expected {
			t.Errorf("Expected: %.2f, result: %.2f ", result, expected)
		}
	}

	verifyArea := func(t *testing.T, name string, form Form, expected float64) {
		t.Helper()
		result := form.Area()

		checkResult(t, result, expected)
	}

	t.Run("Perimetro de um retangulo", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		result := Perimeter(rectangle)
		expected := 40.0

		checkResult(t, result, expected)
	})

	t.Run("Retângulo", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		expected := 72.0

		verifyArea(t, "Teste Simples Retangulo", rectangle, expected)

	})

	t.Run("Circulo", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		verifyArea(t, "Teste Simples Circulo", circle, expected)
	})

	// Table Drive Tests: Uniao dos testes de Retangulo e Cirtculo acima
	t.Run("Geometry Table Drive Tests", func(t *testing.T) {
		testArea := []struct {
			name     string
			form     Form
			expected float64
		}{
			{name: "Retangulo", form: Rectangle{Width: 12, Height: 6}, expected: 72.0},
			{name: "Circlo", form: Circle{Radius: 10}, expected: 314.1592653589793},
			{name: "Triangulo", form: Triangle{Base: 12, Height: 6}, expected: 36.0},
		}

		for _, tt := range testArea {
			verifyArea(t, tt.name, tt.form, tt.expected)
		}
	})

}
