package user

import "github.com/gin-gonic/gin"

func InitUserGroup(e *gin.Engine) {
	// 用户模块
	userGroup := e.Group("/apis/v1/user")

	// 登录模块
	{
		userGroup.POST("/login")
		userGroup.POST("/register")
	}
}
