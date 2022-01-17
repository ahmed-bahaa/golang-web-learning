package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ahmed-bahaa/thirdwebapp/pkg/config"
	"github.com/ahmed-bahaa/thirdwebapp/pkg/handlers"
	"github.com/ahmed-bahaa/thirdwebapp/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var appConf config.AppConfig
var session *scs.SessionManager

func getToday() string {
	today := time.Now().Weekday().String()
	return today
}

func main() {

	appConf.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConf.InProduction

	appConf.Session = session

	tc, err := render.CreateTemplateCach()
	if err != nil {
		log.Println("Error when we calling render test function", err)
	}

	appConf.CachedTemplate = tc
	appConf.UseCache = false

	repo := handlers.NewRepo(&appConf)
	handlers.NewHandlers(repo)

	render.NewTemplate(&appConf)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	log.Println("Starting server on port:", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConf),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
