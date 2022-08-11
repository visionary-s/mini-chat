package routers

import (
	"mini-chat/middlewares/chat"
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

		chatRouter := router.Group("/chat")
		{
			chatRouter.GET("/chats", chat.loadChatList)
			chatRouter.GET("/chats/:chatId")
			chatRouter.POST("/new", chat.CreateChat)
			chatRouter.POST("/message")
		}
	}

	return router
}
