package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/sindrishtepani/bookings/internal/config"
	"github.com/sindrishtepani/bookings/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// Sets the config from the templates pkg
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// create a template cache
	// get the template cache from config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested tempate from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	// render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

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

	return myCache, err
}
