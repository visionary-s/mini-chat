package user

import (
	"mini-chat/models"
	"mini-chat/services/crypto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	models.CreateUser(map[string]interface{}{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}

func Signin(c *gin.Context) {
	user := models.User{}
	logintime, _ := strconv.ParseInt(c.PostForm("signin_time"), 10, 64)
	user.Username = c.PostForm("username")
	user.SigninTime = uint64(logintime)
	user.Password = c.PostForm("password")
	user.IsLogoff = false

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	u := user.SearchUserByName("username", c.PostForm("username"))

	ePwd := crypto.Encrypt(user.Password, u.Pubkey)
	if u.ID > 0 {
		if u.Password != ePwd {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "Wrong password! Please check or signup first!",
			})
			return
		}
	}
}

func Signout(c *gin.Context) {
	user := models.User{}
	user.Username = c.PostForm("username")
	user.Signout()
	c.Redirect(http.StatusFound, "/home")
}

func ListOnline(c *gin.Context) {
	list := models.GetOnlineUsers()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Set("content-type", "application/json")
	c.JSON(http.StatusOK, list)
}
