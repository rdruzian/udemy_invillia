package main

import (
	"fmt"
	"time"
)

func main() {
	ch := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func escrever(texto string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	return c
}
