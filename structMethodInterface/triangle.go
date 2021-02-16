package structmethodinterface

// Triangle tipo de triangulo
type Triangle struct {
	Base   float64
	Height float64
}

// Area Retorna a área do triangulo
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
