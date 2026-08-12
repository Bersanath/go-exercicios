/*Cria uma CLI simples que recebe argumentos via os.Args*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Se não houver argumentos, dá ajuda
	if len(os.Args) < 2 {
		fmt.Println("Uso: programa <nome> [sobrenome]")
		fmt.Println("Exemplo: go run . Eduardo Tecacunda")
		return
	}

	// Junta todos os argumentos a partir do índice 1
	nomeCompleto := strings.Join(os.Args[1:], " ")
	fmt.Printf("Olá, %s! Bem-vindo ao Go.\n", nomeCompleto)
}
