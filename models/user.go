package models

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"mini-chat/services/crypto"
	"time"

	"gorm.io/gorm"
)

// const (
// 	heartbeatTimeout = 3 * 60
// )

var db = New().InitDB()

type User struct {
	gorm.Model
	ID          uint            `json:"id"`
	Username    string          `json:"username"	binding:"required,max=16,min=2"`
	Password    string          `json:"password"	binding:"required,max=32,min=8"`
	Privkey     *rsa.PrivateKey `json:"privkey"	binding:"required"`
	Pubkey      rsa.PublicKey   `json:"pubkey"	binding:"required"`
	SigninTime  uint64          `json:"signin_time"`
	SignoutTime uint64          `json:"signout_time"`
	IsLogoff    bool            `json:"islogoff"`
	// HeartbeatTime uint64 `json:"heartbeatTime"`
}

func CreateUser(value interface{}) {
	var u User

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln(err)
	}
	u.Privkey = privateKey
	u.Pubkey = privateKey.PublicKey

	u.Username = value.(map[string]interface{})["username"].(string)
	pwd := value.(map[string]interface{})["password"].(string)
	ePwd := crypto.Encrypt(pwd, u.Pubkey)
	u.Password = ePwd

	db.Create(&u)
}

func Signin(id uint, username string, password string, signinTime uint64) (user *User) {
	user = &User{
		ID:         id,
		Username:   username,
		Password:   password,
		SigninTime: signinTime,
		// HeartbeatTime: signinTime,
		IsLogoff: false,
	}
	return
}

// func (u *User) UpdateHeartbeat(currentTime uint64) {
// 	u.HeartbeatTime = currentTime
// 	u.IsLogoff = false
// }

func (u *User) Signout() {
	user := u.SearchUser("username", c.PostForm("username"))
	currentTime := uint64(time.Now().Unix())
	user.SignoutTime = currentTime
	user.IsLogoff = true
	db.Save(user)
}

func (u *User) IsOnline() (online bool) {
	if u.IsLogoff {
		return
	}

	// currentTime := uint64(time.Now().Unix())
	// if u.HeartbeatTime < (currentTime - heartbeatTimeout) {
	// 	fmt.Printf("IsOnline: heatbeat timeout!\n %d \t %s \t %d", u.ID, u.Username, u.HeartbeatTime)
	// 	return
	// }

	if u.IsLogoff {
		fmt.Printf("IsOnline: user has been logged off!\n %d \t %s", u.ID, u.Username)
		return
	}
	return true
}

func (u *User) SearchUser(field, value string) User {
	if field == "username" {
		db.Where(field+" = ?", value).First(&u)
	}
	return *u
}

func (u *User) GetOnlineUsers() []map[string]interface{} {
	var results []map[string]interface{}
	db.Where("isLogoff = ?", false).Find(&u)
	return results
}
