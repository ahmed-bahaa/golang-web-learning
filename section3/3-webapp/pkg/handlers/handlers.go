package handlers

import (
	"net/http"

	"github.com/ahmed-bahaa/thirdwebapp/pkg/config"
	"github.com/ahmed-bahaa/thirdwebapp/pkg/models"
	"github.com/ahmed-bahaa/thirdwebapp/pkg/render"
)

// The repo Used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Function to create a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// function sets the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "Hello, Gopher Boy!")
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, fmt.Sprintf("We are Egyptian Gophers, ready to GO! Today is : %s", getToday()))
	myMessage := make(map[string]string)
	myMessage["AboutMessage"] = "Temp hello! :)"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: myMessage,
	})

}
