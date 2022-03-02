package models

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"mini-chat/services/crypto"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID       uint            `json:"id"`
	Sender   string          `json:"sender"`
	TimeSent time.Time       `json:"time_sent"`
	Content  string          `json:"msg"		binding:"required"`
	ChatID   uint            `json:"chat_id"	binding:"required"`
	Privkey  *rsa.PrivateKey `json:"privkey"	binding:"required"`
	Pubkey   rsa.PublicKey   `json:"pubkey"		binding:"required"`
}

func NewMsg(value interface{}) {
	var m Message

	m.Sender = value.(map[string]interface{})["sender"].(string)

	m.TimeSent = time.Now().Local()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln(err)
	}
	m.Privkey = privateKey
	m.Pubkey = privateKey.PublicKey

	msg := value.(map[string]interface{})["msg"].(string)
	eMsg := crypto.Encrypt(msg, m.Pubkey)
	m.Content = eMsg

	m.ChatID = value.(map[string]interface{})["room"].(uint)

	db.Create(&m)
}
