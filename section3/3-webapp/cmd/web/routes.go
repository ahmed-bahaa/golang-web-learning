package main

import (
	"net/http"

	"github.com/ahmed-bahaa/thirdwebapp/pkg/config"
	"github.com/ahmed-bahaa/thirdwebapp/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(appConf *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
