package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ahmed-bahaa/thirdwebapp/pkg/config"
	"github.com/ahmed-bahaa/thirdwebapp/pkg/handlers"
	"github.com/ahmed-bahaa/thirdwebapp/pkg/render"
)

const portNumber = ":8080"

func getToday() string {
	today := time.Now().Weekday().String()
	return today
}

func main() {

	var appConf config.AppConfig

	tc, err := render.CreateTemplateCach()
	if err != nil {
		log.Println("Error when we calling render test function", err)
	}

	appConf.CachedTemplate = tc
	appConf.UseCache = false

	render.NewTemplate(&appConf)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Println("Starting server on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
