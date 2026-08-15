/*Itera sobre uma string com range e mostra os runes.*/

package main

import "fmt"

func main() {

	texto := "Olá, Mundo!"

	for i, r := range texto {

		fmt.Printf("Índice: %d, Rune: %c (unicode: U+%04X)\n", i, r, r)
	}

}
