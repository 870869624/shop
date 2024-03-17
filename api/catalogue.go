package api

import "github.com/gin-gonic/gin"

func CatalogueApi(r *gin.Engine) {
	r.POST("/catalogueAdd")
}
