/*Usa switch simples.*/

package main

import "fmt"

func main() {

	dia := "Sabado"

	switch dia {

	case "Segunda", "Terça", "Quarta", "Quinta", "Sexta":

		fmt.Println("Dia de Semana!")

	case "Sabado", "Domingo":

		fmt.Println("Final de Semana")

	default:

		fmt.Println("Dia Inválido!")
	}

}
