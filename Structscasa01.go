package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func areaCirculo(circulo Circulo) {

	fmt.Println("A area do circulo e: ", circulo.raio*circulo.raio*math.Pi)

}

func main() {

	var c Circulo

	fmt.Println("Qual o valor do raio do seu circulo?")
	fmt.Scan(&c.raio)

	areaCirculo(c)

}
