package arrayslice

// Sum Soma os valores do array
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
}
