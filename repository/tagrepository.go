package repository

import (
	"github.com/genc-murat/go-blog-engine/models"
	"github.com/jinzhu/gorm"
)

//TagRepository ...
type TagRepository interface {
	Add(tag models.Tag)
	GetAll() []models.Tag
}

//TagRepositorySQLite ...
type TagRepositorySQLite struct {
	db *gorm.DB
}

//NewTagRepositorySQLite ...
func NewTagRepositorySQLite(db *gorm.DB) *TagRepositorySQLite {
	return &TagRepositorySQLite{db}
}

//Add ...
func (t *TagRepositorySQLite) Add(tag models.Tag) {
	t.db.Create(&tag)
}

//GetAll ...
func (t *TagRepositorySQLite) GetAll() []models.Tag {
	tags := []models.Tag{}
	t.db.Find(&tags)
	return tags
}
