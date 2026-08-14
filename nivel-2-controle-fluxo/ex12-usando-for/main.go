/*
Usa for das quatro formas:
clássico;
estilo while;
infinito;
com range.
*/

package main

import "fmt"

func main() {

	// Clássico
	print("========== CLÁSSICA ==========\n")
	for i := 0; i <= 10; i++ {

		fmt.Printf("%d\n", i)
	}

	print("========== ESTILO WILHE ==========\n")
	num := 10

	for num > 0 {

		fmt.Println(num)

		num -= 1
	}

	print("========== INFINITO ==========\n")

	infinito := 10

	for infinito > 0 {

		fmt.Println(infinito)

	}

	print("========== RANGE ==========\n")

	Range := []string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"}

	for _, semanas := range Range {

		fmt.Printf("%s\n", semanas)
	}
}
