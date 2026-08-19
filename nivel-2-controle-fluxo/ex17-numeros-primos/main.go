/*Cria um programa que verifica se um número é primo.*/

package main

import (
	"fmt"
	"math"
)

func main(){

	var num int

	fmt.Println("Digite um Número: ")
	
	fmt.Scan(&num)  

	primo := true // assume que é primo

	if num > 1 {

		for i := 2; i <= int(math.Sqrt(float64(num))); i++{

			if num % i == 0 {

				primo = false  // achou um divisor, não é primo
				break  // sai do for
			}
		}
	}else {

		primo = false // números menores ou iguais a 1 não são primos
	}
	if primo {

		fmt.Printf("O número: %d é Primo!", num)

	}else{

		fmt.Printf("O número: %d Não é Primo!", num)
	}

}