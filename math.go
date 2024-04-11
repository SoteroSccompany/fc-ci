package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 2))

}

// Soma é uma função que soma dois números inteiros e retorna o resultado
func Soma(a int, b int) int {
	return a + b
}
