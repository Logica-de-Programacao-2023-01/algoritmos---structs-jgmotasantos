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

func encontrarMusicas(playlists []Playlist, nomeMusica string) []Playlist {
	var resultado []Playlist
	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == nomeMusica {
				resultado = append(resultado, playlist)
			}
		}
	}
	return resultado
}

func main() {
	musica1 := Musica{titulo: "Música 1", artista: "Artista 1", duracao: 180}
	musica2 := Musica{titulo: "Música 2", artista: "Artista 2", duracao: 240}
	musica3 := Musica{titulo: "Música 3", artista: "Artista 3", duracao: 200}
	musica4 := Musica{titulo: "Música 4", artista: "Artista 4", duracao: 300}

	playlist1 := Playlist{nome: "Playlist 1", musicas: []Musica{musica1, musica2}}
	playlist2 := Playlist{nome: "Playlist 2", musicas: []Musica{musica3, musica4}}
	playlist3 := Playlist{nome: "Playlist 3", musicas: []Musica{musica1, musica4}}

	playlists := []Playlist{playlist1, playlist2, playlist3}

	resultado := encontrarMusicas(playlists, "Música 1")

	for _, playlist := range resultado {
		fmt.Println("Nome da playlist:", playlist.nome)
		fmt.Println("Músicas:")
		for _, musica := range playlist.musicas {
			fmt.Println("Título:", musica.titulo)
			fmt.Println("Artista:", musica.artista)
			fmt.Println("Duração:", musica.duracao)
			fmt.Println("---------------------")
		}
	}
}
