package service

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/candbright/go-core/rest"
	"github.com/candbright/go-core/rest/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"piano-server/config"
	"piano-server/server/domain"
	"time"
)

var identityKey = "phone_number"

var AdminUser = &domain.User{
	Name:        "admin",
	Password:    "admin@123456",
	PhoneNumber: "15888888888",
}

func RegisterAuthMiddleware(eng *gin.Engine) *gin.RouterGroup {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       config.Get("application.name"),
		Key:         []byte("secret key"),
		Timeout:     24 * time.Hour,
		MaxRefresh:  24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*domain.User); ok {
				return jwt.MapClaims{
					identityKey: v.PhoneNumber,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			phoneNumber := claims[identityKey].(string)
			if phoneNumber == AdminUser.PhoneNumber {
				return AdminUser
			}
			user, _ := userDao.GetUserByPhoneNumber(phoneNumber)
			return &user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVal domain.User
			if err := c.ShouldBind(&loginVal); err != nil {
				return "", errors.New("missing phone number")
			}
			username := loginVal.Name
			password := loginVal.Password
			phoneNumber := loginVal.PhoneNumber

			if username == "admin" && password == "admin@123456" {
				return AdminUser, nil
			}
			user, err := userDao.GetUserByPhoneNumber(phoneNumber)
			if err != nil {
				return nil, errors.New("incorrect phone number")
			}
			return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			rest.Error(c, rest.StatusError(errors.New(message), code))
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		panic("JWT Error:" + err.Error())
	}
	eng.POST("/login", authMiddleware.LoginHandler)
	eng.POST("/logout", authMiddleware.LogoutHandler)
	// Refresh time can be longer than token timeout
	eng.GET("/refresh_token", authMiddleware.RefreshHandler)
	eng.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		rest.Error(c, rest.StatusError(errors.New("page not found"), http.StatusNotFound))
	})
	authGroup := eng.Group("")
	authGroup.Use(authMiddleware.MiddlewareFunc())
	return authGroup
}

func RegisterHandlers(engine *gin.Engine) {
	engine.Use(handler.LogHandler())
	engine.GET("/register", GetRegisterCode)
	engine.POST("/register", RegisterUser)
	authGroup := RegisterAuthMiddleware(engine)
	authGroup.GET("/users/info", GetUserInfo)
	authGroup.PUT("/users/info", UpdateUserInfo)
	authGroup.DELETE("/users/info", DeleteUserInfo)
}
