package main

import "fmt"

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func increment(f Funcionario, percentual float64) Funcionario {
	f.salario = f.salario + (f.salario * percentual)
	return f
}

func decrement(f Funcionario, percentual float64) Funcionario {
	f.salario = f.salario - (f.salario * percentual)
	return f
}

func calcularTempoServico(f Funcionario) int {
	return f.idade - 18
}

func main() {

	funcionario := Funcionario{
		nome:    "João",
		salario: 2500.0,
		idade:   30,
	}

	funcionario = increment(funcionario, 0.10)

	fmt.Println("Nome:", funcionario.nome)
	fmt.Println("Salário:", funcionario.salario)
	fmt.Println("Idade:", funcionario.idade)

	funcionario = decrement(funcionario, 0.05)

	fmt.Println("Salário atualizado:", funcionario.salario)

	tempoServico := calcularTempoServico(funcionario)
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}
