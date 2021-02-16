package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSum := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected: '%v', Result: %v", expected, result)
		}
	}

	t.Run("Array Dinâmico", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("Expected %d, Result %d, Data %v", expected, result, numbers)
		}
	})

	t.Run("Soma Tudo", func(t *testing.T) {
		result := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		checkSum(t, result, expected)

	})

	t.Run("Somar Todo o Resto", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSum(t, result, expected)
	})

	t.Run("Soma slices vazios de form segura", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSum(t, result, expected)
	})

}
