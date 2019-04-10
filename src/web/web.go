package web

import (
	log "github.com/sirupsen/logrus"
	"goji.io"
	"goji.io/pat"
	"html/template"
	"net/http"
	"time"
)

var (
	Templates *template.Template
	
	RootMux *goji.Mux
)

func Run() {
	Templates = template.Must(Templates.ParseFiles("templates/index.html"))
	
	mux := SetupRoutes()
	log.Info("Running web server...")
	RunServer(mux)
}

func Stop() {
	log.Info("web.Stop() was called!")
}

func RunServer(MainMux *goji.Mux) {
	log.Info("Starting web server at :5000")
	
	server := &http.Server {
		Addr: ":5000",
		Handler: MainMux,
		IdleTimeout: time.Minute,
	}
	
	err := server.ListenAndServe()
	if err != nil {
		log.Error("Failed to ListenAndServe(), ", err)
	}
}

func SetupRoutes() *goji.Mux {
	SetupRootMux()
	
	return RootMux
}

func SetupRootMux() {
	mux := goji.NewMux()
	RootMux = mux
	
	mux.Handle(pat.Get("/static/*"), http.FileServer(http.Dir(".")))
	
	// Routes
	mux.HandleFunc(pat.Get("/"), index)
}

func index(w http.ResponseWriter, r *http.Request) {
	TemplateData := make(map[string]interface {})
	
	err := Templates.ExecuteTemplate(w, "index", TemplateData)
	if err != nil {
		log.Error("Failed to execute template: ", err)
	}
}