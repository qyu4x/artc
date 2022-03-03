package web

type ArticleUpdateRequest struct {
	Id int `validate:"required" json:"id"`
	Title string `validate:"required" min:"1" json:"title"`
	Author string `validate:"required" min:"1" json:"author"`
	Content string `validate:"required" min:"1" json:"content"`
}
