package controllers

import (
	"net/http"

	"github.com/genc-murat/go-blog-engine/repository"
	"github.com/gin-gonic/gin"
)

//MainController main controller
type MainController struct {
	posts      repository.PostRepository
	tags       repository.TagRepository
	categories repository.CategoryRepository
}

//NewMainController ...
func NewMainController(posts repository.PostRepository, tags repository.TagRepository, categories repository.CategoryRepository) *MainController {
	return &MainController{posts, tags, categories}
}

//Index anasayfa
func (c *MainController) Index(ctx *gin.Context) {

	posts := c.posts.GetLastPosts(10, 0)
	categories := c.categories.GetAll()
	//ctx.Bind(posts)
	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":      "Index title!",
		"posts":      posts,
		"categories": categories,
	})

}
