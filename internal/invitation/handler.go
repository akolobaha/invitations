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

	fName := path + ".jpg"
	fullPath := fmt.Sprintf("web/images/%s", fName)

	if fileExists(fullPath) {

		tmpl, err := template.ParseFiles("web/templates/invitation.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := Data{Url: "images/" + fName}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	http.NotFound(w, r)
	return
}
