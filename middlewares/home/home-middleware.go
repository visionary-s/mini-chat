package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "websockets.html", gin.H{
		"title": "Main Page",
	})
}

// func Echo(w http.ResponseWriter, r *http.Request) {
// 	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

// 	for {
// 		// Read message from browser
// 		msgType, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			return
// 		}

// 		// Print the message to the console
// 		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

// 		// Write message back to browser
// 		if err = conn.WriteMessage(msgType, msg); err != nil {
// 			return
// 		}
// 	}
// }
