package main

import (
	"fmt"
	"time"
)

/*
Canal possui 2 operações, Enviar e Receber dados
Essas operações são bloqueantes
Caso os canais não estejam mais sendo usados é necessário fechar o canal
*/

func main() {
	ch := make(chan string)
	go escrever("Olá mundo", ch)

	/*for {
		mensagem, aberto := <-ch
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}*/
	for mensagem := range ch {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do loop")
}

func escrever(texto string, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- texto
		time.Sleep(time.Second)
	}
	close(ch)
}
