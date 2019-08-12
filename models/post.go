package models

import "time"

type Post struct {
	Title    string
	SubTitle string
	Content  string
	Created  time.Time
	Slug     string
	Tags     []string
}
