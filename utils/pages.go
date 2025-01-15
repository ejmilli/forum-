package utils

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func InitializeHtml() {

	var err error
	tpl, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Printf("Error parsing templates: %v", err)
	}

}

func LoginPage(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Fatal("error executing login page ")
	}
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		log.Fatal("error executing login page ")
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "homepage.html", nil)
	if err != nil {
		log.Fatal("error executing home page ")
	}
}
