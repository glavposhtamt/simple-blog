package main

import (
	"go-cms/lib/blog"
	"go-cms/lib/dashboard"
	"net/http"

	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := httprouter.New()
	router.GET("/", blog.Index)
	router.GET("/admin/", dashboard.Admin)
	router.GET("/post/:id", blog.Post)
	router.GET("/spa/:id", blog.PostJSON)
	router.ServeFiles("/css/*filepath", http.Dir("web/css/"))
	http.ListenAndServe(":8080", router)
}
