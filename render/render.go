package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, templateFile string) {
	parsedTemplate, templateErr := template.ParseFiles("./templates/" + templateFile)

	if templateErr != nil {
		fmt.Println("Templating file read failed", templateErr)
	}

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Rendering template failed", err)
	}
}
