package controller

import (
	"net/http"
	"html/template"
	// "fmt"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html", "views/home.html")
	t.ExecuteTemplate(w, "layout", nil)

}

func Test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html")
	t.ExecuteTemplate(w, "layout", nil)

}



