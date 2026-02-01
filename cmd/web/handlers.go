package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"
)

type GenericData struct {
	Now time.Time
}

// render executes a template into a bytes buffer and returns an error if anything goes wrong in the process.
func render(ts *template.Template) (*bytes.Buffer, error) {
	var data = GenericData{
		Now: time.Now(),
	}
	buf := new(bytes.Buffer)
	err := ts.Execute(buf, data)
	if err != nil {
		log.Printf("Encountered an error when trying to execute template, %s", err)
		return nil, err
	}
	return buf, nil

}


func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var (
		ts  *template.Template
		ok  bool
		err error
	)
	if app.useCache {
		ts, ok = app.templateCache["home.page.tmpl"]
		if !ok {
			log.Println("Couldn't find the `home.page.tmpl` template in the template cache")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		ts, err = loadTemplate("./static/html/home.page.tmpl", "./static/html")
		if err != nil {
			log.Printf("Failed to load home page template with %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	buf, err := render(ts)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) Research(w http.ResponseWriter, r *http.Request) {
	var (
		ts  *template.Template
		ok  bool
		err error
	)
	if app.useCache {
		ts, ok = app.templateCache["research.page.tmpl"]
		if !ok {
			log.Println("Couldn't find the `research.page.tmpl` template in the template cache")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		ts, err = loadTemplate("./static/html/research.page.tmpl", "./static/html")
		if err != nil {
			log.Printf("Failed to load research page template with %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	buf, err := render(ts)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) Group(w http.ResponseWriter, r *http.Request) {
	var (
		ts  *template.Template
		ok  bool
		err error
	)
	if app.useCache {
		ts, ok = app.templateCache["group.page.tmpl"]
		if !ok {
			log.Println("Couldn't find the `cv.page.tmpl` template in the template cache")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		ts, err = loadTemplate("./static/html/group.page.tmpl", "./static/html")
		if err != nil {
			log.Printf("Failed to load cv page template with %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	buf, err := render(ts)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) Teaching(w http.ResponseWriter, r *http.Request) {
	var (
		ts  *template.Template
		ok  bool
		err error
	)
	if app.useCache {
		ts, ok = app.templateCache["teaching.page.tmpl"]
		if !ok {
			log.Println("Couldn't find the `teaching.page.tmpl` template in the template cache")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		ts, err = loadTemplate("./static/html/teaching.page.tmpl", "./static/html")
		if err != nil {
			log.Printf("Failed to load teaching page template with %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	buf, err := render(ts)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) Cv(w http.ResponseWriter, r *http.Request) {
	var (
		ts  *template.Template
		ok  bool
		err error
	)
	if app.useCache {
		ts, ok = app.templateCache["cv.page.tmpl"]
		if !ok {
			log.Println("Couldn't find the `cv.page.tmpl` template in the template cache")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		ts, err = loadTemplate("./static/html/cv.page.tmpl", "./static/html")
		if err != nil {
			log.Printf("Failed to load cv page template with %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	buf, err := render(ts)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
