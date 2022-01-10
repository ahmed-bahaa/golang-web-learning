package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmp string) {
	_, err := renderTemplateTest(w)
	if err != nil {
		log.Println("Error when we calling render test function", err)
	}
	parsedTemp, _ := template.ParseFiles("./templates/" + tmp)
	err = parsedTemp.Execute(w, nil)
	if err != nil {
		log.Println("err with parsing the template")
		return
	}
}

func renderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {

	myCach := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Println("Error Will fetching the page templates", err)
		return myCach, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("we are working on: ", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("Error while creating new templates", err)
			return myCach, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("Error while matching layouts", err)
			return myCach, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println("Error while parsing layouts", err)
				return myCach, err
			}

		}

		myCach[name] = ts

	}

	return myCach, nil
}
