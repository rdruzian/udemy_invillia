package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Paralelismo -> são 2 tarefas/processos sendo executadas ao mesmo tempo.
	Concorrência -> são 2 tarefas/processos acessando o mesmo recurso, ou rodando na mesma thread.
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		escrever("olá mundo") // goroutine, nova thread
		wg.Done()
	}()
	go func() {
		escrever("programando em go")
		wg.Done()
	}()

	wg.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
