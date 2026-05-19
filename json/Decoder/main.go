package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	var c Cachorro
	cachorroEmJSON := `{"nome": "Rex", "raca": "Dálmata", "idade": 5}`

	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome": "Toby", "raca": "Poodle"}`
	c2 := make(map[string]string)
	if err := json.Unmarshal([]byte(cachorro2EmJSON), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)
}
