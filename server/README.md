# Deprecated
This is a deprecated version of a basic message system with tcp connection only.

### Usage
1. Start a new server with `go run main.go -addr :<port>`  
   e.g.
   ```go
   $ go run main.go -addr :8090
    2022/02/21 15:39:02 New server created!
    2022/02/21 15:39:02 Running server...
    2022/02/21 15:39:02 listening to port :8090
   ```
2. Connection using telent:  
   Note that you have to enable `telnet` on windows first.
   ```
    telnet localhost 8090
   ```
3. Use `/` + cmd to execute:
   ```
    /help           see all available cmds
    /name           specify your presented name
    /list           list all available users can connect to
    /join <name>    join to someone's chat
    /msg <word>     send words to users you have joined
    /quit           connection break
   ```
   ```bash
   $ go run main.go -addr :8090
   2022/02/21 15:39:02 New server created!
   2022/02/21 15:39:02 Running server...
   2022/02/21 15:39:02 listening to port :8090
   2022/02/21 15:47:15 Add new client : [::1]:61751
   2022/02/21 15:47:15 new client has joined : [::1]:61751
   2022/02/21 15:48:26 Add new client : [::1]:61874
   2022/02/21 15:48:26 new client has joined : [::1]:61874
   2022/02/21 15:49:00 Encrypting messages...
   2022/02/21 15:49:00 Sending message to orianna
   2022/02/21 15:49:33 Encrypting messages...
   2022/02/21 15:49:33 Sending message to eliza
   2022/02/21 15:51:07 Encrypting messages...
   2022/02/21 15:51:07 Sending message to orianna
   2022/02/21 15:51:17 [::1]:61874 left the chat.
   2022/02/21 15:51:17 read tcp [::1]:8090->[::1]:61874: use of closed network connection
   ```