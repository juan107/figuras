package figuras

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) getArea() float64 {
	return c.Lado * c.Lado
}

func (c *Cuadrado) getPerimetro() float64 {
	return c.Lado * 4
}
