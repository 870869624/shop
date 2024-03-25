package main

import (
	"shop/api"
	"shop/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	db.Migrate() //建立数据库,但没有建立表格（迁移）
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	api.UserApi(r)
	api.CatalogueApi(r)
	r.Run(":8081")
}
