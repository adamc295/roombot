package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	templates = template.Must(template.ParseFiles("templates/index.html"))
)

func main() {
	
	var err error
	
	log.Println("Starting web server...")
	
	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		return
	}
}