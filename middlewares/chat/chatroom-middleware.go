package chat

import (
	"mini-chat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {

	// should include values of user list, which contains ids of selected users and the current user
	models.NewChat(map[string]interface{}{
		"users": c.PostForm("users"),
	})
}

func LoadAllChat(userId uint) []*models.Chat {
	var chats []*models.Chat
	chat := models.Chat{}

	list := chat.FindChatByUser(userId)
	for _, id := range list {
		single := chat.FindChatByID(id)
		chats = append(chats, &single)
	}
	return chats
}

func loadChatList(c *gin.Context) {
	// here need to add session store to get the current user info, to check if the user is in this chat
	currentUser, _ := strconv.ParseInt(c.PostForm("userId"), 10, 64)
	u := uint(currentUser)
	LoadAllChat(u)
}

func loadSingleChat(c *gin.Context) {
	chatID := c.Param("chatId")
}

func UpdateChatname(c *gin.Context) {

}

func sendMsg(c *gin.Context) {

}
