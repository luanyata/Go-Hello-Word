package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

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

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected: '%v', Result: %v", expected, result)
		}
	})

}
