package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog *log.Logger
}

func main() {
	app := application{
		log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}

	mux := app.routes()

	app.infoLog.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
