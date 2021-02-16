package structmethodinterface

// Perimeter Retorna perimetor de um retangulo
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area retorna a area de um triangulo
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
