package api

import (
	"github.com/candbright/go-core/rest/handler"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(engine *gin.Engine) {
	engine.Use(handler.LogHandler())
	initServices()

	engine.GET("/register", RegisterService.GetRegisterCode)
	engine.POST("/register", RegisterService.RegisterUser)
	authGroup := RegisterAuthMiddleware(engine)
	authGroup.GET("/users/info", RegisterService.GetUserInfo)
	authGroup.PUT("/users/info", RegisterService.UpdateUserInfo)
	authGroup.DELETE("/users/info", RegisterService.DeleteUserInfo)
}
