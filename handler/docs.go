package handler

import (
	"html/template"
	"log"
	"net/http"
)

var swaggerTemplate = loadTemplate("swagger/swagger.html")

func DocsHandler(w http.ResponseWriter, r *http.Request) {
	if err := swaggerTemplate.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, "server error", 500)
	}
}

func loadTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		panic(err)
	}
	return t
}
