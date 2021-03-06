package router

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/b3kt/account-srv/controller"
	"github.com/b3kt/account-srv/middleware"
	"github.com/b3kt/account-srv/model"
)

const prefix = "/api/v1"

// Route makes the routing
func Route(app *gin.Engine) {

	indexController := new(controller.IndexController)
	app.GET(
		"/", indexController.GetIndex,
	)

	auth := app.Group(prefix + "/auth")
	authMiddleware := middleware.Auth()
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get("email")
			c.JSON(200, gin.H{
				"email": claims["email"],
				"name":  user.(*model.User).Username,
				"text":  "Hello World.",
			})
		})
	}

	userController := new(controller.UserController)
	app.GET(prefix+"/user/:id", userController.GetUser)
	app.POST(prefix+"/auth/register", userController.Signup)
	app.POST(prefix+"/auth/login", userController.Signin)
	app.POST(prefix+"/recovery", userController.Recovery)
	app.POST(prefix+"/resetpass", userController.ResetPass)

	api := app.Group(prefix + "/api")
	{
		api.GET("/version", indexController.GetVersion)
	}
}
