package main

import "fmt"

type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func (a *Animal) definirSom(novoSom string) {
	a.som = novoSom
}

func (a Animal) imprimirInformacoes() {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Espécie:", a.especie)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Som:", a.som)
}

func main() {

	animal := Animal{
		nome:    "Rex",
		especie: "Cachorro",
		idade:   5,
		som:     "Au Au",
	}

	animal.imprimirInformacoes()

	animal.definirSom("Guau Guau")

	animal.imprimirInformacoes()
}
