package repository

import (
	"github.com/genc-murat/go-blog-engine/models"
	"github.com/jinzhu/gorm"
)

//CategoryRepository ...
type CategoryRepository interface {
	Add(categgory models.Category)
	GetAll() []models.Category
}

//CategoryRepositorySQLite ...
type CategoryRepositorySQLite struct {
	db *gorm.DB
}

//NewCategoryRepositorySQLite ...
func NewCategoryRepositorySQLite(db *gorm.DB) *CategoryRepositorySQLite {
	return &CategoryRepositorySQLite{db}
}

//Add ...
func (c *CategoryRepositorySQLite) Add(category models.Category) {
	c.db.Create(&category)
}

//GetAll ...
func (c *CategoryRepositorySQLite) GetAll() []models.Category {
	categories := []models.Category{}
	c.db.Find(&categories)
	return categories
}
