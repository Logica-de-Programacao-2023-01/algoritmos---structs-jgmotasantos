package main

import "fmt"

type Viagem struct {
	origem  string
	destino string
	data    string
	preco   float64
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		return Viagem{} // Retorna uma viagem vazia se o slice estiver vazio
	}

	viagemMaisCara := viagens[0]
	for _, viagem := range viagens {
		if viagem.preco > viagemMaisCara.preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {

	viagens := []Viagem{
		{origem: "São Paulo", destino: "Rio de Janeiro", data: "2023-05-20", preco: 500.0},
		{origem: "Paris", destino: "Londres", data: "2023-06-15", preco: 800.0},
		{origem: "Nova York", destino: "Los Angeles", data: "2023-07-10", preco: 1000.0},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)

	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.origem)
	fmt.Println("Destino:", viagemMaisCara.destino)
	fmt.Println("Data:", viagemMaisCara.data)
	fmt.Println("Preço:", viagemMaisCara.preco)
}
