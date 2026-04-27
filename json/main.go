package main

import (
	"bytes"
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
	c := Cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cachorroJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Struct: ", cachorroJson)
	fmt.Println("bytes.NewBuffer: ", bytes.NewBuffer(cachorroJson))
	fmt.Println("Casting para string: ", string(cachorroJson))

	c2 := map[string]string{
		"nome":  "Rex",
		"raca":  "Dalmata",
		"idade": "Rex",
	}
	fmt.Println(c2)
	cachorro2Json, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Map: ", cachorro2Json)
	fmt.Println("bytes.NewBuffer: ", bytes.NewBuffer(cachorro2Json))
	fmt.Println("Casting para string: ", string(cachorro2Json))
}
