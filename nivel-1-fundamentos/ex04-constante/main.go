/*Cria constantes simples e constantes com iota.*/

package main

import "fmt"

const num int = 20

const (
	
	domingo = iota
	segunda
	terca
	quarta
)

func main() {

	fmt.Println(num)
	fmt.Println(domingo, segunda, terca, quarta)
}