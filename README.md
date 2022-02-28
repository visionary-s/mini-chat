# Mini Chat
This is currently a basic implementation for the encrypted chatting system, where clients can only connect to server via tcp. For specific please refer to `server` folder.  
I am turning the whole project to a new mode. More functions are in development and will soon apply.

A chat app via WebSocket using Golang with `Gin` & `Gorilla WebSocket`

# TO DO
1. Intergration for user registeration and login part
2. Update with group info
3. Frontend UI
4. Upload file(picture, video less than 5M, etc) in chat room


# Running the server
1. Clone this repository
2. Mount the repository & run this command to install dependencies
```bash
$ go get
```
3. Run the websocket server
```bash
$ go run main.go -addr <:port>
```
Websocket server will be running on localhost:<port>