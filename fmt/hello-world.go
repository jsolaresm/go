package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// Declaracion de constantes a nivel de paquetes
const pi float32 = 3.1416
const (
	x = 100    //
	y = 0b1010 //binario
	z = 0o12   //octal
	w = 0xFF   //hexadecimal
)
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(quote.Go())

	//Declaración de variables

	firstName, lastName := "Javier", "Solares"
	age := 31
	age = 37
	fullName := "Javier Solares \t (alias \"javiercode\")\n"
	var a byte = 'a'
	var r rune = '❤'

	var (
		defaultint    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
		integer16     int16  = 50
		integer32     int32  = 100
		s             string = "10"
		name          string
		name2         string
		name3         string
		age2          int
	)
	i, _ := strconv.Atoi(s)

	fmt.Println(defaultBool, defaultFloat, defaultString, defaultUint, defaultint)

	fmt.Println(firstName, lastName, age, pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)
	fmt.Println(math.MinInt64, math.MaxInt64)
	fmt.Println(fullName)
	fmt.Println(a)
	fmt.Println(s[0])
	fmt.Println(r)
	fmt.Println(integer16 + int16(integer32) + int16(i))
	fmt.Print("hola")
	fmt.Printf("Hola, me llamo %s y tengo %d años, \n", name, age)

	fmt.Println("Ingrese su nombre :")
	fmt.Scanln(&name, &name2, &name3)
	fmt.Println("Ingrese su edad :")
	fmt.Scanln(&age2)

	fmt.Printf("Hola, me llamo %s %s %s y tengo %d años, \n", name, name2, name3, age2)
}
