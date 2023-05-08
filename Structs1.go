package main

import "fmt"

func main() {

	var r ret
	fmt.Println("Qual a largura do retangulo?")
	fmt.Scanln(&r.largura)

	fmt.Println("Qual a altura do retangulo?")
	fmt.Scanln(&r.altura)

	area(r)

}

//Crie uma struct chamada Retângulo com os campos "largura" e "altura". Escreva uma
//função que receba um Retângulo como parâmetro e calcule a área do retângulo (área =
//largura * altura).

type ret struct {
	largura float64
	altura  float64
}

func area(ret ret) {

	fmt.Println("A area do retangulo e: ", ret.largura*ret.altura)

}
