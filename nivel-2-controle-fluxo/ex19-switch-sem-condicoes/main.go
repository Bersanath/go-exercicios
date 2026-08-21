/*Usa switch sem condição para classificar notas.*/

package main

import "fmt"

func main() {

	nota := 95

	switch {

	case nota >= 80:

		fmt.Println("Nota: B")

	case nota >= 90:

		fmt.Println("Nota: A")

	default:

		fmt.Println("Nota: C")
	}

}
