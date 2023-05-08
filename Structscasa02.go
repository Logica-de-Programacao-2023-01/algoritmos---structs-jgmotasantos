package main

import "fmt"

// Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço".
//O campo "endereço" deve ser uma outra struct com os campos "rua", "número", "cidade" e "estado".
//Escreva uma função que receba uma Pessoa como parâmetro e imprima seu endereço completo.

type Pessoa struct {
	nome     string
	idade    int
	endereço Endereço
}

type Endereço struct {
	rua    string
	cidade string
	estado string
	numero int
}

func parametro(pessoa Pessoa) {

}

func main() {
	p := Pessoa{
		"João",
		18,
		Endereço{
			"RuaX",
			"Sao Paulo",
			"SP",
			517,
		},
	}
	fmt.Printf("A pessoa foi: %s\n", p.nome)
	fmt.Printf("A idade dele(a) é: %d\n", p.idade)
	fmt.Printf("A pessoa mora na rua: %s\n", p.endereço.rua)
	fmt.Printf("A pessoa mora no numero: %d\n", p.endereço.numero)
	fmt.Printf("A pessoa mora na cidade: %s\n", p.endereço.cidade)
	fmt.Printf("A pessoa mora no estado: %s\n", p.endereço.estado)

	parametro(p)
}
