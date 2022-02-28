package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/atharva-art/bookings-project/internals/config"
	"github.com/atharva-art/bookings-project/internals/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets config for the template pkg
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request)*models.TemplateData{
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache{
		tc = app.TemplateCache
	}else{
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok{
		log.Fatal("Could not get template from cache")
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buff, td)

	_, err := buff.WriteTo(w)
	if err != nil{
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cahe as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// get template cache from config
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}
