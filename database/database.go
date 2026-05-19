package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(db)

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	row, err := db.Query("select * from usuarios")
	if err != nil {
		fmt.Sprintf("error to execute query %v+", err)
		log.Fatal(err)
	}
	fmt.Sprintf("query result %v", row)
}
