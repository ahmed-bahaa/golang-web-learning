package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/ahmed-bahaa/thirdwebapp/pkg/config"
	"github.com/ahmed-bahaa/thirdwebapp/pkg/models"
)

var functions = template.FuncMap{}
var conf *config.AppConfig

func NewTemplate(app *config.AppConfig) {
	conf = app
}

func RenderTemplate(w http.ResponseWriter, tmp string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if conf.UseCache {
		tc = conf.CachedTemplate
	} else {
		tc, _ = CreateTemplateCach()
	}

	t, ok := tc[tmp]
	if !ok {
		log.Fatal("Couldn't find the template")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error when writting to the writter", err)
	}
}

func CreateTemplateCach() (map[string]*template.Template, error) {

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
