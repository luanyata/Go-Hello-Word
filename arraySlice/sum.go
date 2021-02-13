package arrayslice

// Sum Soma os valores do array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll Soma os valores dos slice e retorna um novo slice
func SumAll(numbersAll ...[]int) []int {

	var sum []int

	for _, number := range numbersAll {
		sum = append(sum, Sum(number))
	}

	return sum
}
