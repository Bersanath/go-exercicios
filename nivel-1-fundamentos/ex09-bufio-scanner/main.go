/*Lê uma linha completa com bufio.Scanner.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// // Cria um scanner que lê da entrada padrão (teclado)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite seu nome completo: ")

	// O programa pausa aqui e espera o utilizador pressionar Enter
	if scanner.Scan() {

		nomeCompleto := scanner.Text() //  // Captura o texto completo, incluindo os espaços

		fmt.Printf("O seu nome completo é: %q\n", nomeCompleto)

	}

}
