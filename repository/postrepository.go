package repository

import (
	"github.com/genc-murat/go-blog-engine/models"
	"github.com/jinzhu/gorm"
)

//PostRepository ...
type PostRepository interface {
	Add(post models.Post) uint
	Update(post models.Post)
	GetByID(id uint) models.Post
	GetLastPosts(take int, skip int) []models.Post
}

//PostRepositorySQLite ...
type PostRepositorySQLite struct {
	db *gorm.DB
}

//NewPostRepositorySQLite ...
func NewPostRepositorySQLite(db *gorm.DB) *PostRepositorySQLite {
	return &PostRepositorySQLite{db}
}

//Add ...
func (p *PostRepositorySQLite) Add(post models.Post) uint {
	p.db.NewRecord(post)
	p.db.Create(&post)

	return post.ID
}

//Update ...
func (p *PostRepositorySQLite) Update(post models.Post) {

	currentPost := new(models.Post)
	p.db.Where("ID=?", post.ID).Find(&currentPost)
	p.db.Model(&currentPost).Updates(post)

}

//GetByID ...
func (p *PostRepositorySQLite) GetByID(id uint) models.Post {

	post := new(models.Post)
	p.db.First(&post, id)
	return *post
}

//GetLastPosts ...
func (p *PostRepositorySQLite) GetLastPosts(take int, skip int) []models.Post {

	posts := []models.Post{}
	p.db.Order("created_at DESC").Limit(take).Offset(skip).Find(&posts)
	return posts
}
