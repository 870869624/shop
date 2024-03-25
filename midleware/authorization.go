package midleware

import (
	"net/http"
	"shop/db"
	"shop/models"

	"github.com/gin-gonic/gin"
)

// 从token获取用户信息然后和数据库用户信息对比，检验是否存在该用户,校验接口
func ParaseJwt() func(ctx *gin.Context) {
	db := db.Connect()
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "请登录",
			})
			ctx.Abort()
		}
		GetUser, err := models.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "登陆信息错误",
			})
			return
		}
		//获得token却在数据库没有该用户
		request := db.Where(map[string]interface{}{"username": GetUser.Username, "password": GetUser.Password})
		if request.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "权限不足",
			})
			return
		}
		ctx.Request.Header.Set("user", GetUser.Username)
	}
}
