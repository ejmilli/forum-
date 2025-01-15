package utils

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

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
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Fatal("Error executing login page: ", err)
	}
}

// SignUpPage handler
func SignUpPage(w http.ResponseWriter, r *http.Request) {
	// Handle sign-up form submission logic here if needed
	// For now, we're just serving the same login page template
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Fatal("Error executing signup page: ", err)
	}
}


