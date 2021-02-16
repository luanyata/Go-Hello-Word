package structmethodinterface

// Rectangle Dados de um Retangulo
type Rectangle struct {
	Width  float64
	Height float64
}

// Area Retorna a área de um retangulo
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
