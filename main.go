package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/project.html", project)
	http.HandleFunc("/blog.html", blog)

	http.Handle("/img/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.Handle("/css/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.Handle("/file/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tpl := template.Must(template.ParseFiles("./template/index.html"))
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func project(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./template/project.html"))
	err := tpl.ExecuteTemplate(w, "project.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func blog(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./template/blog.html"))
	err := tpl.ExecuteTemplate(w, "blog.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
