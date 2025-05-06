package invitation

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Data struct {
	Url string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	path = strings.TrimPrefix(path, "/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	fullPathJpg := fmt.Sprintf("web/images/%s.jpg", path)
	fullPathPng := fmt.Sprintf("web/images/%s.png", path)

	filePath, err := existingFile([]string{fullPathPng, fullPathJpg})

	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/invitation.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := Data{Url: strings.TrimPrefix(filePath, "web/")}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return

}
