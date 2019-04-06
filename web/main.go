package main

import (
	"html/template"
	"log"
	"net/http"
	
	"goji.io"
	"goji.io/pat"
)

var (
	templates = template.Must(template.ParseFiles("templates/index.html"))
)

func main() {
	
	var err error
	
	log.Println("Starting web server...")
	
	mux := setupRoutes()
	
	err = http.ListenAndServe(":5000", mux)
	if err != nil {
		return
	}
}

func setupRoutes() *goji.Mux {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), index)
	mux.Handle(pat.Get("/static/*"), http.FileServer(http.Dir(".")))
	return mux
}