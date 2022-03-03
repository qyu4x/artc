package web

import "time"

type ArticleResponse struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
