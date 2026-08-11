/*
Usa fmt.Printf com vários verbos:
%v
%+v
%#v
%T
%d
%s
%q
%t
%f
*/
package main

import "fmt"

type Usuario struct {
	Nome  string
	Idade int
}

func main() {

	u := Usuario{Nome: "João", Idade: 23}

	nome := "Eduardo"
	inteiro := 10
	flutuante := 10.50
	booleano := true
	letra := 'E'

	fmt.Printf("Nome: %s\n", nome)        // O verbo %s: serve para texto (string)
	fmt.Printf("Inteiro: %d\n", inteiro)  // O verbo %d: serve para números inteiros (int, inteiro de struct)
	fmt.Printf("Decimais: %f\n", flutuante) // O verbo %f: serve para números decimais (float64 ou float32)
	fmt.Printf("Bool: %t\n", booleano)    // O verbo %t: serve para os valores lógicos (true e false ou 0 e 1)
	fmt.Printf("Coringa: %v\n", u)        // O verbo %v: o famoso coringa ele mostra o valor de qualquer tipo
	fmt.Printf("Usuario: %+v\n", u)       // O verbo %+v: mostra apensas os nomes dos campos juntos os valores
	fmt.Printf("Usuario: %#v\n", u)         // O verbo %#v: mostra a representação exata em sintaxe go (tipo,campos e valores)
	fmt.Printf("Letra: %q\n", letra)         // O verbo %q: Adiciona aspas e escapa caracteres invisíveis.
	fmt.Printf("Tipo do usuario: %T\n", u)

}
