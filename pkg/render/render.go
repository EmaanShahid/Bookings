package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/EmaanShahid/bookings/pkg/config"
	"github.com/EmaanShahid/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// sets the config for new template package
func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

//RenderTemplate renders template using html/template

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// if we get past this then we have template that we want to use
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td) // take the template execute it dont pass any data and store it in a buff variable

	_, err := buf.WriteTo(w) // write the buffer to the response writer
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// CreateTemplateCache: creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//go to the templates folder and get all the files having name page.tmpl

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page) // get the name of page to create a template set

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// now we're going to check: should i use a layout that's defined for this template for this template
		//so look for the existence of layout

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
