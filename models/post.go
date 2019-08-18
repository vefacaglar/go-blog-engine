package models

import (
	"github.com/jinzhu/gorm"
)

//Post ...
type Post struct {
	gorm.Model
	Title      string
	Summary    string
	Author     string
	Content    string
	Slug       string
	Tags       []Tag      `gorm:"association_autoupdate:false"`
	Categories []Category `gorm:"association_autoupdate:false"`
}

//Tag ...
type Tag struct {
	gorm.Model
	Name   string
	PostID uint
}

//Category ...
type Category struct {
	gorm.Model
	Name   string
	PostID uint
}
