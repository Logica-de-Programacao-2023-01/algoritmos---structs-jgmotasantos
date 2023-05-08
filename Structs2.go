package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
	ano    int
}

func LivroX(livro Livro) {

	fmt.Printf("Seu livro foi: %s\n", livro.titulo)
	fmt.Printf("O autor foi: %s\n", livro.autor)
	fmt.Printf("O ano de publicaçao foi: %d\n", livro.ano)

}

func main() {
	livro := Livro{("Le Petit Prince"), ("Antoine de Saint-Exupéry"), (1943)}

	LivroX(livro)

}
