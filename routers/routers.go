package routers

import (
	"mini-chat/middlewares/home"
	"mini-chat/middlewares/user"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	v1 := router.Group("/v1")
	{
		v1.GET("/home", home.Index)

		userRouter := router.Group("/user")
		{
			userRouter.POST("/signup", user.Signup)
			userRouter.POST("/signin", user.Signin)
			userRouter.GET("/signout", user.Signout)
			userRouter.GET("/all", user.ListOnline)
		}
	}

	return router
}
