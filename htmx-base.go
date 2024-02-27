package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func renderStatic(res http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("src/public/" + tmpl + ".html")
	if err != nil {
		fmt.Println("Error executing template")
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(res, nil)
}

func renderComponent(res http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("src/components/" + tmpl + ".html")
	if err != nil {
		fmt.Println("Error executing template")
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(res, nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	renderStatic(res, "index")
}

func clickHandler(res http.ResponseWriter, req *http.Request) {
	renderComponent(res, "clicked")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/clicked", clickHandler)
	fs := http.FileServer(http.Dir("./src/public/css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))
	log.Fatal(http.ListenAndServe(":8777", nil))
}
