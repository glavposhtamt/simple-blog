package model

import (
	"database/sql"
	"fmt"
)

// MysqlConnect return *sql.DB connection
func MysqlConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:mysql@/go-cms")

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return db
}

// SelectAllFromPost return slice Post
func SelectAllFromPost() []Post {
	db := MysqlConnect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM post")
	if err != nil {
		panic(err.Error())
	}

	var result []Post

	for results.Next() {

		var post Post

		err = results.Scan(&post.ID, &post.Title, &post.Description, &post.Post,
			&post.Date, &post.Author, &post.Thumbnail, &post.Categories)

		if err != nil {
			panic(err.Error())
		}
		result = append(result, post)
	}

	return result
}

// SelectFromPostByID return one row by Id
func SelectFromPostByID(id int) *Post {
	db := MysqlConnect()
	defer db.Close()

	var post Post

	err := db.QueryRow("SELECT * FROM post where id = ?", id).Scan(&post.ID,
		&post.Title, &post.Description, &post.Post, &post.Date, &post.Author,
		&post.Thumbnail, &post.Categories)

	if err != nil {
		panic(err.Error())
	}

	return &post
}
