package main

import (
	"task/controllers"
	_ "task/docs"
	"task/helper"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	helper.ConnectDatabase()
}

// @title Blog and News API
// @version 1.0
// @description A CRUD API for Blog and News

// @host 	localhost 8080
// @Basepath /api
func main() {
	r := gin.Default()
	// endpoint for swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// blog endpoints
	r.POST("/blogs", controllers.CreateBlogController)
	r.GET("/blogs", controllers.GetBlogsController)
	r.GET("/blogs/:id", controllers.GetBlogByIdController)
	r.PUT("/blogs/:id", controllers.UpdateBlogByIdController)
	r.DELETE("/blogs/:id", controllers.DeleteBlogByIdController)

	// news endpoints
	r.POST("/news", controllers.CreateNewsController)
	r.GET("/news", controllers.GetNewsController)
	r.GET("/news/:id", controllers.GetNewsByIdController)
	r.PUT("/news/:id", controllers.UpdateNewsByIdController)
	r.DELETE("/news/:id", controllers.DeleteNewsByIdController)

	r.Run(":8080")
}
