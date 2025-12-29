package main

import (
	"cmd_udemy/app"
	"fmt"
	"os"
)

/*
	Para executar o programa deve-se executar o seguinte comando no terminal
	go run main.go ip --host <host>
	go run main.go servidores --host <host>

	Outra maneira é seguindo os seguintes comandos
	go build
	./cmd_udemy ip --host <host>
	./cmd_udemy servidores --host <host>
*/

func main() {
	fmt.Println("Iniciando...")

	aplicacao := app.Gerar()
	aplicacao.Run(os.Args)
}
