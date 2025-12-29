package main

import (
	"fmt"
	"time"
)

/*
	Paralelismo -> são 2 tarefas/processos sendo executadas ao mesmo tempo.
	Concorrência -> são 2 tarefas/processos acessando o mesmo recurso, ou rodando na mesma thread.
*/

func main() {
	go escrever("olá mundo") // goroutine, nova thread
	escrever("programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
