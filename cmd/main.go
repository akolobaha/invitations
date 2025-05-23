package main

import (
	"fmt"
	"invitation/internal/invitation"
	"net/http"
)

const imageDir = "web/images"
const cssDir = "web/css"
const faviconDir = "web"
const host = "localhost"
const port = "8081"

func main() {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(imageDir))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(cssDir))))
	http.Handle("/favicon.ico", http.FileServer(http.Dir(faviconDir)))
	http.HandleFunc("/", invitation.Handler)
	fmt.Printf("Server listening on %s:%s", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
