package controller

import (
	"net/http"
	"html/template"
	"go-blog/common/model"
	"go-blog/common/core"
	"path"
	// "fmt"
	
)

func Blog(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html", "views/blog.html")
	id := core.StrToInt(path.Base(r.URL.Path))
	t.ExecuteTemplate(w, "layout", id)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Blog model.Blog
	}
	id := core.StrToInt(path.Base(r.URL.Path))

	res := data{}
	if id > 0 {
		res.Blog = model.GetBlog(id)
	}

	core.JsonCode(w, 200, res)
}




