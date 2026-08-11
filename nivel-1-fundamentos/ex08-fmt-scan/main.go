/*Lê dados do utilizador com fmt.Scan.*/

package main

import "fmt"

func main() {

	var nome string
	var idade int

	fmt.Print("Digite o seu nome e a sua idade: ")
	fmt.Scan(&nome, &idade) // A função fmt.Scan lê texto inserido no terminal, separando os valores por espaços em branco ou quebras de linha. Ela exige que você passe o endereço de memória das variáveis (usando o operador &) para conseguir salvar os dados digitados

	fmt.Printf("O seu nome é: %s e a sua idade é: %d", nome, idade)
}
