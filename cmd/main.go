package main

import (
	_ "database/sql"
	"fmt"
	_ "log"
	"net/http"

	"github.com/a-h/templ"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luisnogit/learningGO/views"
)

func main() {
	http.Handle("/", templ.Handler(views.Landingpage()))

	fmt.Println("deu certo")
	http.ListenAndServe(":8080", nil)

}
