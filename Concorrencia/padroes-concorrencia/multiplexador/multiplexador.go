package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := multiplexar(escrever("Ola mundo"), escrever("Programando em GO."))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func multiplexar(ch1, ch2 <-chan string) <-chan string {
	chSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-ch1:
				chSaida <- mensagem
			case mensagem := <-ch2:
				chSaida <- mensagem
			}
		}
	}()

	return chSaida
}

func escrever(texto string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		}
	}()

	return c
}
