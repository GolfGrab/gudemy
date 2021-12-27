package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/GolfGrab/gudemy/pkg/config"
	"github.com/GolfGrab/gudemy/pkg/models"
)

var function = template.FuncMap{}

var app *config.AppConfig

// NewTemplates is a helper function to sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate is a helper function to render a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("Error finding template:", tmpl)

		return
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template:", err)
		log.Fatalln(err)
		return
	}

}

// CreateTemplateCache is a helper function to create a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		fmt.Println("Error finding templates:", err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// fmt.Println("page is currently: ", page)
		ts, err := template.New(name).Funcs(function).ParseFiles(page)

		if err != nil {
			fmt.Println("Error parsing template:", err)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			fmt.Println("Error finding templates:", err)
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				fmt.Println("Error parsing template:", err)
				return myCache, err
			}

		}
		myCache[name] = ts

	}
	return myCache, nil
}
