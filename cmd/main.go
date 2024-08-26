package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(database:3307)/")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Conn alive")
	}

	http.Handle("/", http.FileServer(http.Dir("./web/templates")))
	fmt.Println("deu certo")
	http.ListenAndServe(":8080", nil)

}
