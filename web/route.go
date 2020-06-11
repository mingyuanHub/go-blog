package main

import (
	"net/http"
	ctr "go-blog/web/controller"
)

func BindRoute(mux *http.ServeMux) {
	mux.HandleFunc("/Index", ctr.Index)
	mux.HandleFunc("/home", ctr.Home)
	mux.HandleFunc("/list", ctr.GetBlogList)
	mux.HandleFunc("/blog/", ctr.Blog)
	mux.HandleFunc("/detail/", ctr.GetBlog)
}
