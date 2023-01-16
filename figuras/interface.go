package figuras

import "fmt"

type funcionesGeometricas interface {
	getArea() float64
	getPerimetro() float64
}

func Medidas(g funcionesGeometricas) {
	fmt.Println(g)
	fmt.Println("Area: ", g.getArea())
	fmt.Println("Perimetro:", g.getPerimetro())
}
