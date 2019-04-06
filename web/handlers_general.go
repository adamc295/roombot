package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	templateData := make(map[string]interface{})
	
	err := templates.ExecuteTemplate(w, "index", templateData)
	if err != nil {
		log.Println("Failed to execute template: ", err)
	}
}