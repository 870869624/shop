package api

import (
	"net/http"
	"shop/models"
	"shop/service"

	"github.com/gin-gonic/gin"
)

func CatalogueApi(r *gin.Engine) {
	var catalogue models.Catalogue
	r.POST("/catalogueAdd", func(ctx *gin.Context) {
		if err := ctx.ShouldBindJSON(&catalogue); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "类目信息错误",
			})
			return
		}
		service.CreateCatalogue(&catalogue)
		ctx.Abort()

	})
}
