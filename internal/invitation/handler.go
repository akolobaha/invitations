package invitation

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	path = strings.TrimPrefix(path, "/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	fullPath := fmt.Sprintf("web/images/%s.jpg", path)

	if fileExists(fullPath) {
		w.Header().Set("Content-Type", "text/html")

		htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
		<title>Приглашение</title>
		<meta charset="utf-8">
		<style>
		  body {
			font-family: Arial, sans-serif;
			background-color: #f0f0f0;
			text-align: center;
		  }

		  h1 {
			color: #333;
		  }

		  p {
			font-size: 1.2em;
			color: #666;
		  }
		</style>
		</head>
		<body>
		  <img src="images/%s.jpg"></img>
		</body>
		</html>
		`, path)

		fmt.Fprintf(w, htmlContent)

		return
	}

	http.NotFound(w, r)
	return
}
