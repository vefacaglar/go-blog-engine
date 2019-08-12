package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/genc-murat/go-blog-engine/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")

	gv := goview.New(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/_menu", "partials/_footer", "partials/_header", "partials/_mobilemenu"},
		/*Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},*/
		DisableCache: true,
	})

	goview.Use(gv)

	router.HTMLRender = ginview.Default()
	mainController := controllers.MainController{} //new constructor ı kullanmak için

	router.GET("/", mainController.Index)

	router.Run(":8080")
}
