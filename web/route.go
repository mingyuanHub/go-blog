package main

import (
	"net/http"
	ctr "go-blog/web/controller"
)

func BindRoute(mux *http.ServeMux) {
	mux.HandleFunc("/home", ctr.Home)
	mux.HandleFunc("/list", ctr.GetBlogList)

	mux.HandleFunc("/test", ctr.Test)
}
