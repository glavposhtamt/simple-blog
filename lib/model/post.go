package model

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