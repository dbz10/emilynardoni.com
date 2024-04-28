package main

import (
	"html/template"
	"log"
	"path/filepath"
)

// loadTemplate loads a template from disk
func loadTemplate(path string, root string) (*template.Template, error) {

	partials, err := filepath.Glob(filepath.Join(root, "*.partial.tmpl"))
	if err != nil {
		return nil, err
	}

	layouts, err := filepath.Glob(filepath.Join(root, "*.layout.tmpl"))
	if err != nil {
		return nil, err
	}

	files := []string{path}
	files = append(files, partials...)
	files = append(files, layouts...)

	ts, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	return ts, nil

}

// newTemplateCache parses the `root` directory and builds a map containing
// each *.page.tmlp file combined together with its dependencies
func newTemplateCache(root string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(root, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	loadedTemplates := []string{}
	for _, page := range pages {
		templateName := filepath.Base(page)

		ts, err := loadTemplate(page, root)
		if err != nil {
			log.Printf("Encountered an error when trying to parse template %s, %s", page, err)
			continue
		}

		cache[templateName] = ts
		loadedTemplates = append(loadedTemplates, templateName)
	}

	log.Printf("Loaded the template cache with pages %s", loadedTemplates)

	return cache, nil

}
