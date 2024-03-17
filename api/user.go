package api

import (
	"net/http"
	"shop/entities"
	"shop/midleware"
	"shop/models"
	"shop/service"

	"github.com/gin-gonic/gin"
)

func UserApi(r *gin.Engine) {
	user_WithoutAuthgroup := r.Group("/users")
	{
		// 注册
		user_WithoutAuthgroup.POST("/register", func(ctx *gin.Context) {
			var userRegisterPayload entities.UserRegisterPayload
			if err := ctx.ShouldBindJSON(&userRegisterPayload); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// 检验user的值是不是符合要求
			if err := userRegisterPayload.Validate(); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			//创建用户
			err := service.CreateUser(userRegisterPayload)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "创建成功",
			})
		})
	}
	//登录,调用AUthentication
	r.POST("/users/signin", func(ctx *gin.Context) {
		var user models.Authentication
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "参数错误",
			})
			return
		}
		login := entities.UserLoginRequestPayload{
			Username: user.Username,
			Password: user.Password,
		}
		if err1 := login.Validate(); err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err1.Error(),
			})
			return
		}

		token, err2 := user.Signin()
		if err2 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err2.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"token": token,
		})
	})

	user_WithAuthGroup := r.Group("/users")
	user_WithAuthGroup.Use(midleware.ParaseJwt())
	{
		//注销
		user_WithAuthGroup.DELETE("/logout", func(ctx *gin.Context) {
			var user entities.UserLoginRequestPayload
			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := service.DeleteUser(user); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "注销成功",
			})
		})
		user_WithAuthGroup.PATCH("/updateuser", func(ctx *gin.Context) {
			var user entities.UserRegisterPayload
			user_Information := ctx.Request.Header.Get("username")
			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := user.Validate(); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			if err := service.UpdataUser(user, user_Information); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "更新成功",
			})
		})
	}
}
