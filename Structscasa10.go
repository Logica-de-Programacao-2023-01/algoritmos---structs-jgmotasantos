package main

import (
	"fmt"
	"math"
)

type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

func (f *Filme) adicionarAvaliacao(avaliacao int) {
	f.avaliacoes = append(f.avaliacoes, avaliacao)
}

func (f *Filme) removerAvaliacao(indice int) {
	if indice >= 0 && indice < len(f.avaliacoes) {
		f.avaliacoes = append(f.avaliacoes[:indice], f.avaliacoes[indice+1:]...)
	}
}

func (f Filme) calcularMediaAvaliacoes() float64 {
	soma := 0
	for _, avaliacao := range f.avaliacoes {
		soma += avaliacao
	}
	media := float64(soma) / float64(len(f.avaliacoes))
	return math.Round(media*100) / 100 // Arredondar para duas casas decimais
}

func (f Filme) imprimirInformacoes() {
	fmt.Println("Título:", f.titulo)
	fmt.Println("Diretor:", f.diretor)
	fmt.Println("Ano:", f.ano)
	fmt.Println("Média das avaliações:", f.calcularMediaAvaliacoes())
}

func main() {

	filme := Filme{
		titulo:     "Inception",
		diretor:    "Christopher Nolan",
		ano:        2010,
		avaliacoes: []int{8, 9, 7},
	}

	filme.adicionarAvaliacao(8)

	filme.imprimirInformacoes()

	filme.removerAvaliacao(1)

	filme.imprimirInformacoes()
}
