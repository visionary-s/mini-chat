package models

import (
	"strings"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID       uint          `json:"id"`
	ChatName string        `json:"chat_name"`
	Users    []uint        `json:"user_list" binding:"required"`
	History  []interface{} `json:"history"`
}

func NewChat(value interface{}) {
	var c Chat
	c.Users = value.(map[string]interface{})["users"].([]uint)
	if name := value.(map[string]interface{})["chatname"].(string); name != "" {
		c.ChatName = name
	} else {
		var nameList []string
		for _, id := range c.Users {
			var u User
			u = u.SearchUserById("id", id)
			nameList = append(nameList, u.Username)
		}
		c.ChatName = strings.Join(nameList, `','`)
	}
	db.Create(&c)
}

// save every sent msg to chat history
func (c *Chat) SaveHistory(chatID uint, sender string, msg string) {
	chat := c.FindChatByID(chatID)
	chat.History = append(c.History, map[string]string{
		"sender": sender,
		"msg":    msg,
	})
	db.Save(chat)
}

// update user list for sake of some one join or leave the chatroom
func (c *Chat) UpdateUsers(chatID uint, userId uint, operation string) {
	chat := c.FindChatByID(chatID)
	if operation == "join" {
		chat.Users = append(chat.Users, userId)
	} else {
		for index, id := range chat.Users {
			if id == userId {
				chat.Users = append(chat.Users[:index], chat.Users[index+1:]...)
				break
			}
		}
	}
	db.Save(chat)
}

func (c *Chat) FindChatByID(id uint) Chat {
	db.Where("id = ?", id).First(&c)
	return *c
}

func (c *Chat) FindChatByUser(userId uint) []uint {
	var allChats []map[string]interface{}
	var list []uint
	db.Find(&allChats)
	for _, e := range allChats {
		userList := e["user_list"]
		for _, id := range userList.([]uint) {
			if id == userId {
				list = append(list, id)
			}
		}
	}
	return list
}
