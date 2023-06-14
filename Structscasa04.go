package main

import "fmt"

type Musica struct {
	titulo  string
	artista string
	duracao int
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func imprimirPlaylist(p Playlist) {
	fmt.Println("Nome da playlist:", p.nome)

	tempoTotal := 0
	for _, musica := range p.musicas {
		fmt.Println("Título:", musica.titulo)
		fmt.Println("Artista:", musica.artista)
		fmt.Println("Duração:", musica.duracao)
		fmt.Println("---------------------")

		tempoTotal += musica.duracao
	}

	fmt.Println("Tempo total da playlist:", tempoTotal, "segundos")
}

func main() {

	playlist := Playlist{
		nome: "Minha Playlist",
		musicas: []Musica{
			{titulo: "Música 1", artista: "Artista 1", duracao: 180},
			{titulo: "Música 2", artista: "Artista 2", duracao: 240},
			{titulo: "Música 3", artista: "Artista 3", duracao: 200},
		},
	}

	imprimirPlaylist(playlist)
}
