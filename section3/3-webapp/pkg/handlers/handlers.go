package handlers

import (
	"net/http"

	"github.com/ahmed-bahaa/thirdwebapp/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "Hello, Gopher Boy!")
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, fmt.Sprintf("We are Egyptian Gophers, ready to GO! Today is : %s", getToday()))
	render.RenderTemplate(w, "about.page.tmpl")

}
