/*Faz conversões explícitas entre tipos numéricos.*/

package main

import "fmt"

func main() {

	var salario float64 = 350.560

	// 1. O problema: Tipos diferentes não se misturam sozinhos
	var precoInteiro int = 10
	var taxaDecimal float64 = 5.5

	// O código abaixo daria ERRO de compilação:
	// var total = precoInteiro * taxaDecimal 

	// 2. A solução: Conversão Explícita (Manual)
	// Transformamos o 'int' em 'float64' para fazer a conta
	var total float64 = float64(precoInteiro) * taxaDecimal
	fmt.Println("Total da conta:", total) // Resultado: 55

	// 3. Cuidado com a perda de precisão!
	// Ao converter um decimal para inteiro, Go corta as casas decimais (não 
	var salarioInt int = int(salario)

	fmt.Printf("Salario em inteiro: %d", salarioInt)
}
