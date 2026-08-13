/*
Mini projeto do nível:
Cria uma calculadora CLI simples que recebe dois números e uma operação via argumentos.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Valida número de argumentos
	if len(os.Args) != 4 {

		fmt.Println("Uso: programa <num1> <op> <num2>")
		return

	}

	// Lê os argumentos (índices corretos!)
	textonum1 := os.Args[1]
	textoOp := os.Args[2]
	textonum2 := os.Args[3]

	// Converte primeiro número
	n1, err := strconv.ParseFloat(textonum1, 64)
	if err != nil {

		fmt.Printf("erro: '%s'não é um número válido\n", textonum1)
		return
	}

	// Converte segundo número
	n2, err := strconv.ParseFloat(textonum2, 64)
	if err != nil {

		fmt.Printf("erro: '%s'não é um número válido\n", textonum2)
		return
	}

	// Executa a operação
	var resultado float64

	switch textoOp {
	case "+":
		resultado = n1 + n2
	case "-":
		resultado = n1 - n2
	case "*":
		resultado = n1 * n2
	case "/":
		if n2 == 0 {
			fmt.Println("erro: divisão por zero")
			return
		}
		resultado = n1 / n2
	default:
		fmt.Printf("erro: operação '%s' não suportada\n", textoOp)
		return
	}

	// Mostra o resultado formatado
	fmt.Printf("%.2f %s %.2f = %.2f\n", n1, textoOp, n2, resultado)

}
