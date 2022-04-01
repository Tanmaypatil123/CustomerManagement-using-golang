package main

import (
	"net/http"
	"text/template"
	"log"
)

func Home(w http.ResponseWriter, r *http.Request){
	var template_html *template.Template
	template_html = template.Must(template.ParseFiles("main.html"))
	template_html.Execute(w,nil)
}

func main(){
	log.Println("Server started on : http://localhost:8000")
	http.HandleFunc("/",Home)
	http.ListenAndServe(":8000",nil)
}