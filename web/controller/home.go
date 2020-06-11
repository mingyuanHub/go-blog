package controller

import (
	"net/http"
	"html/template"
	"go-blog/common/model"
	"go-blog/common/core"
	// "fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html", "views/home.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func GetBlogList(w http.ResponseWriter, r *http.Request) {
	type data struct {
		List []model.Blog
	}
	res := data{}
	res.List = model.ListBlog()
	core.JsonCode(w, 200, res)
}





