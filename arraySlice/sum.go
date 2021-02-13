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
func SumAll(numbersAll ...[]int) (sum []int) {

	qtdNumber := len(numbersAll)
	sum = make([]int, qtdNumber)

	for i, number := range numbersAll {
		sum[i] = Sum(number)
	}

	return
}
