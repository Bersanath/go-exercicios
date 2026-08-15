/*Usa break e continue.*/

package main

import "fmt"

func main() {

	for {

		var num int

		fmt.Println("Digite um número: ")
		fmt.Scan(&num)

		if num != 200 {

			continue
		} else {

			break
		}

	}
}
