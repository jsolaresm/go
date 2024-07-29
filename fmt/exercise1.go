package main

import (
	"fmt"
	"math"
)

func calcularhipotenusa(a, b float64) float64 {

	//return math.sqrt((a * a) + (b * b))
	return math.Sqrt((a * a) + (b * b))

}

func calculararea(a, b float64) float64 {

	//return math.sqrt((a * a) + (b * b))
	return (a * b) / 2

}

func calcularPerimetro(a, b float64) float64 {

	//return math.sqrt((a * a) + (b * b))
	c := calcularhipotenusa(a, b)
	return (a + b + c)

}

func main() {
	const presicion = 2
	var (
		a float64
		b float64
	)

	fmt.Println("ingrese el lado 1: ")
	fmt.Scanln(&a)
	fmt.Println("ingrese el lado 2")
	fmt.Scanln(&b)

	hipotenusa := calcularhipotenusa(a, b)
	fmt.Printf("La hipotenusa del triángulo es: %.*f\n", presicion, hipotenusa)

	area := calculararea(a, b)
	fmt.Printf("El  area del triángulo es: %.*f\n", presicion, area)

	perimetro := calcularPerimetro(a, b)
	fmt.Printf("El  perimetro del triángulo es: %.*f\n", presicion, perimetro)

}
