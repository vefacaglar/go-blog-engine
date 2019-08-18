package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/genc-murat/go-blog-engine/controllers"
	"github.com/genc-murat/go-blog-engine/models"
	"github.com/genc-murat/go-blog-engine/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := gorm.Open("sqlite3", "blog.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Tag{})
	defer db.Close()

	router := gin.Default()
	router.Static("/static", "./static")

	gv := goview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{"partials/_menu", "partials/_footer", "partials/_header", "partials/_mobilemenu"},
		DisableCache: true,
	})

	goview.Use(gv)
	router.HTMLRender = ginview.Default()

	postRepo := repository.NewPostRepositorySQLite(db)
	tagRepo := repository.NewTagRepositorySQLite(db)
	categoryRepo := repository.NewCategoryRepositorySQLite(db)

	mainController := controllers.NewMainController(postRepo, tagRepo, categoryRepo)

	router.GET("/", mainController.Index)

	router.Run(":8080")
}
