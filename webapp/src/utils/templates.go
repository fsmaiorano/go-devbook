package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("src/views/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) {
	templates.ExecuteTemplate(w, name, data)
}
