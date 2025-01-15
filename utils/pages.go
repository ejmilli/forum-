package utils

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

<<<<<<< HEAD
func InitializeHtml() {

	var err error
	tpl, err = template.ParseGlob("template/*.html")
=======
// Initialize HTML templates
func InitializeHtml() {
	var err error
	tpl, err = template.ParseGlob("template/*.html") // adjust path based on folder structure
	if err != nil {
		log.Fatal("Error parsing templates: ", err)
	}
}

// LoginPage handler
func LoginPage(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD

	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Fatal("error executing login page ")
	}
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "signup.html", nil)
=======
	err := tpl.ExecuteTemplate(w, "login.html", nil)
>>>>>>> fa668d48cbc8fbb03f8ec97b03e7e410053aa341
	if err != nil {
		log.Fatal("Error executing login page: ", err)
	}
}

<<<<<<< HEAD
func HomePage(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "homepage.html", nil)
	if err != nil {
		log.Fatal("error executing home page ")
	}
}
=======
// SignUpPage handler
func SignUpPage(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		log.Fatal("Error executing signup page: ", err)
	}
}


>>>>>>> fa668d48cbc8fbb03f8ec97b03e7e410053aa341
