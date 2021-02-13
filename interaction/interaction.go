package interaction

// Repeat : Retorna o caracter repetido 5 vezes
func Repeat(chart string, qtdLoop int) string {
	var repeats string

	for i := 0; i <= qtdLoop; i++ {
		repeats += chart
	}

	return repeats
}
