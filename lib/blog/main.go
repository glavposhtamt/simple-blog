package blog

import (
	"go-cms/lib/model"
	"html/template"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

// IndexRender is a global render template namespace
type IndexRender struct {
	Posts []model.Post
	Title string
}

// PostRender is a global render template namespace
type PostRender struct {
	Post  model.Post
	Title string
}

// Title for <title> and <h1>
const Title = "Блог Сергея Семёнова"

// Index page handler
func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := model.SelectAllFromPost()
	pr := IndexRender{Posts: res, Title: Title}

	w.Header().Set("Content-type", "text/html")

	t, _ := template.ParseFiles("web/index.html")

	t.Execute(w, &pr)
}

// Post page handler
func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sID := ps.ByName("id")
	id := checkID(sID)

	res := model.SelectFromPostByID(id)
	pr := PostRender{Post: *res, Title: Title}

	w.Header().Set("Content-type", "text/html")

	t, _ := template.ParseFiles("web/post.html")

	t.Execute(w, &pr)
}

// PostJSON response
func PostJSON(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	sID := ps.ByName("id")
	id := checkID(sID)

	res := model.SelectFromPostByID(id)

	bytePost, _ := json.Marshal(res)

	w.Write(bytePost)
}

func checkID(id string) int {
	i, err := strconv.Atoi(id)

	if err != nil {
		panic(err.Error())
	}
	return i
}
