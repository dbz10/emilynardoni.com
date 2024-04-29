package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog       *log.Logger
	useCache      bool
	templateCache map[string]*template.Template
}

func main() {

	tc, err := newTemplateCache("./static/html")
	if err != nil {
		log.Fatalf("Could not instantiate template cache, %s", err)
	}

	app := application{
		infoLog:       log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		templateCache: tc,
		useCache:      true,
	}

	mux := app.routes()

	app.infoLog.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
