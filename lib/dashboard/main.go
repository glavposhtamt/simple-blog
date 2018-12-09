package dashboard

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type page struct {
	Title string
	Msg   string
}

// Post result from database
type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Post        string `json:"post"`
	Date        string `json:"date"`
	Author      int    `json:"author"`
	Thumbnail   string `json:"thumbnail"`
	Categories  int    `json:"categories"`
}

// Admin page handler
func Admin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-type", "text/html")

	t, _ := template.ParseFiles("admin/admin.html")
	t.Execute(w, &page{Title: "Portfolio", Msg: ""})
}
