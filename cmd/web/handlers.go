package main

import (
	"log"
	"net/http"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	files := []string{
		"./static/html/home.page.tmpl",
		"./static/html/base.layout.tmpl",
		"./static/html/footer.partial.tmpl",
		"./static/html/navbar.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("This is the about page"))
}

func (app *application) research(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("This is the reserach page"))
}

func (app *application) downloadCv(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/content/emily_nardoni.pdf")
}
