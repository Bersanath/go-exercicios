/*
Trabalha com tipos básicos:
int
int8, int16, int32, int64
uint
float32, float64
bool
string
byte
rune
*/

package main

import "fmt"

func main() {

	// ==========================================
	// 1. BOOLEANO (bool)
	// ==========================================

	var ligado bool = true

	// ==========================================
	// 2. INTEIROS COM SINAL (int, int8 a int64)
	// ==========================================

	var i int = -1000 // Tamanho automático (32 ou 64 bits)
	var i8 int8 = -128 // Limite: -128 a 127
	var i16 int16 = -32768 // Limite: -32768 a 32767
	var i32 int32 = -2147483648
	var i64 int64 = -9223372036854775808

	// ==========================================
	// 3. INTEIROS SEM SINAL (uint, uint8 a uint64, uintptr)
	// ==========================================

	var u uint = 5000          // Apenas positivos. Tamanho automático.
	var u8 uint8 = 255         // Limite: 0 até 255
	var u16 uint16 = 65535     // Limite: 0 até 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var ptr uintptr = 0xDEADC0DE // Usado para guardar endereços de memória

	// ==========================================
	// 4. DECIMAIS / PONTO FLUTUANTE (float32, float64)
	// ==========================================
	var f32 float32 = 3.141592
	var f64 float64 = 3.141592653589793 // Padrão e mais preciso

	// ==========================================
	// 5. TEXTO E CARACTERES (string, byte, rune)
	// ==========================================
	var texto string = "Golang em português 🇧🇷"
	var b byte = 'A'          // Apelido para uint8 (Guarda o código ASCII 65)
	var r rune = '💻'         // Apelido para int32 (Guarda o código Unicode do emoji)

	// ==========================================
	// 6. NÚMEROS COMPLEXOS (complex64, complex128)
	// ==========================================
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 4 + 7i

	// ==========================================
	// EXIBINDO OS VALORES
	// ==========================================
	fmt.Println("--- Booleano ---")
	fmt.Printf("bool: %v\n\n", ligado)

	fmt.Println("--- Inteiros com Sinal ---")
	fmt.Printf("int: %v, int8: %v, int16: %v, int32: %v, int64: %v\n\n", i, i8, i16, i32, i64)

	fmt.Println("--- Inteiros sem Sinal ---")
	fmt.Printf("uint: %v, uint8: %v, uint16: %v, uint32: %v, uint64: %v, uintptr: %v\n\n", u, u8, u16, u32, u64, ptr)

	fmt.Println("--- Decimais ---")
	fmt.Printf("float32: %v, float64: %v\n\n", f32, f64)

	fmt.Println("--- Texto e Caracteres ---")
	fmt.Printf("string: %s\n", texto)
	fmt.Printf("byte: %v (caractere impresso: %c)\n", b, b)
	fmt.Printf("rune: %v (caractere impresso: %c)\n\n", r, r)

	fmt.Println("--- Complexos ---")
	fmt.Printf("complex64: %v, complex128: %v\n", c64, c128)
}



