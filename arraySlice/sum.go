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

// SumAllRest Somta tudo meno o primeiro numero
func SumAllRest(numbersSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			finale := numbers[1:]
			sums = append(sums, Sum(finale))
		}
	}

	return sums
}
