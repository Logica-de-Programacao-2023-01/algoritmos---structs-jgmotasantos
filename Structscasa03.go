package main

import (
	"fmt"
)

type Triângulo struct {
	base   float64
	altura float64
}

func areaTriangulo(triângulo Triângulo) {

	fmt.Println("A area do triangulo e :", (triângulo.base*triângulo.altura)/2)

}

func main() {

	var t Triângulo

	fmt.Println("Qual o valor da base do seu triangulo?")
	fmt.Scan(&t.base)

	fmt.Println("Qual o valor da altura do seu triangulo?")
	fmt.Scan(&t.altura)

	areaTriangulo(t)

}
