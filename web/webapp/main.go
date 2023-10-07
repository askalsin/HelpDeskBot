package main

import (
	"html/template"
	"net/http"
	"os"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}	

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	
	mux.HandleFunc("/price_list", indexHandler)
	log.Error.Fatalln(http.ListenAndServeTLS(":" + port, "localhost.crt", "localhost.key", mux))
}
