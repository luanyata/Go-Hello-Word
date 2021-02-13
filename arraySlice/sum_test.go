package arrayslice

import "testing"

func TestSum(t *testing.T) {

	t.Run("Array Dinâmico", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("Expected %d, Result %d, Data %v", expected, result, numbers)
		}
	})

}
