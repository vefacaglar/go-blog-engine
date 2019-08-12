package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//MainController main controller
type MainController struct{}

//Index anasayfa
func (c MainController) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title": "Index title!",
	})

}
