package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("stating the web application..")
	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/incident", incident)
	http.HandleFunc("/status", statusPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "PulseTracker is up and running")
	if err != nil {
		return
	}
}

func incident(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "incident details will be visible here..")
	if err != nil {
		return
	}
}

func statusPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("status.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}
