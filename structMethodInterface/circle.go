package structmethodinterface

import "math"

// Circle tipo criculo
type Circle struct {
	radiu float64
}

// Area Retorna a área de um circulo
func (c Circle) Area() float64 {
	return math.Pi * c.radiu * c.radiu
}
