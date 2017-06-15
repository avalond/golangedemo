package error

import (
	"net/http"
	"html/template"
	"golangedemo/code/libs"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	t, err := template.ParseFiles("template/html/404.html")
	if (err != nil) {
		libs.Print(err)
	}
	t.Execute(w, nil)
}
