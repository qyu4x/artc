package domain

import "time"

type Article struct {
	Id int
	Title string
	Author string
	Content string
	CreatedAt time.Time
}
