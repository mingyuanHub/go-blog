package main

import (
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	str, _ := os.Getwd()
	files := http.FileServer(http.Dir(str + "/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	BindRoute(mux)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}