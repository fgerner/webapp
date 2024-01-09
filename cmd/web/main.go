package main

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	app := application{}

	app.Session = getSession()

	log.Println("starting server...")

	err := http.ListenAndServe("localhost:8080", app.routes())
	if err != nil {
		log.Println("Error", err)
	}
}
