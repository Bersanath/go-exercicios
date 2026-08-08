/*
Mostra o zero value de cada tipo básico:
int
float64
bool
string
pointer
slice
map
channel
interface
*/

package main

import "fmt"



func main() {

	var ch chan int
	var slice []int
	var i interface{}
	var m map[int]string
	var n1 int
	var n2 float64
	var n3 bool
	var n4 string
	var n5 *int

	fmt.Println("=== Zero Values ===")
	fmt.Println("int:", n1)           // 0
	fmt.Println("float64:", n2)      // 0
	fmt.Println("bool:", n3)         // false
	fmt.Println("string:", n4)       // "" (vazio)
	fmt.Println("pointer:", n5)      // <nil>
	fmt.Println("channel:", ch)      // <nil>
	fmt.Println("slice:", slice)     // []
	fmt.Println("map:", m)           // map[]
	fmt.Println("interface:", i)     // <nil>
}
