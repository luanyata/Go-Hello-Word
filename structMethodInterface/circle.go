package structmethodinterface

import "math"

// Circle tipo criculo
type Circle struct {
	Radius float64
}

// Area Retorna a área de um circulo
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
