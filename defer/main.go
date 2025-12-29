package main

import "fmt"

func teste1() {
	fmt.Println("Teste 1")
}

func teste2() {
	fmt.Println("Teste 2")
}

func main() {
	// defer vai executar da última chamada para a primeira nesse caso a saída é:
	// Teste 2 -> foi a última chamada do defer
	// Teste 1 -> primeira chamada do defer
	// as funções defer são LIFO
	defer teste1()
	teste2()
}
