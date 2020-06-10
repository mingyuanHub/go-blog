package main

import (
	"net/http"
	ctr "go-blog/web/controller"
)

func BindRoute(mux *http.ServeMux) {
	mux.HandleFunc("/", ctr.Home)
	mux.HandleFunc("/test", ctr.Test)
}
