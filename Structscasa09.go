package main

import (
	"fmt"
	"math"
)

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.notas = append(a.notas, nota)
}

func (a *Aluno) removerNota(indice int) {
	if indice >= 0 && indice < len(a.notas) {
		a.notas = append(a.notas[:indice], a.notas[indice+1:]...)
	}
}

func (a Aluno) calcularMedia() float64 {
	soma := 0.0
	for _, nota := range a.notas {
		soma += nota
	}
	media := soma / float64(len(a.notas))
	return math.Round(media*100) / 100 // Arredondar para duas casas decimais
}

func (a Aluno) imprimirInformacoes() {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Média das notas:", a.calcularMedia())
}

func main() {

	aluno := Aluno{
		nome:  "João",
		idade: 20,
		notas: []float64{7.5, 8.0, 6.5},
	}

	aluno.adicionarNota(9.0)

	aluno.imprimirInformacoes()

	aluno.removerNota(1)

	aluno.imprimirInformacoes()
}
