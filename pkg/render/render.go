package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(rw http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(rw, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
