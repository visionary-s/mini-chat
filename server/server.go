package server

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	contacts map[string]*connection
	commands chan cmd
}

func NewServer() *Server {
	return &Server{
		contacts: make(map[string]*connection),
		commands: make(chan cmd),
	}
}

func (s *Server) Run() {
	log.Printf("Running server...")

	for cmd := range s.commands {
		switch cmd.id {
		case cmdName:
			s.name(cmd.connection, cmd.args[1])
		case cmdJoin:
			s.join(cmd.connection, cmd.args[1])
		case cmdList:
			s.list(cmd.connection)
		case cmdMsg:
			s.msg(cmd.connection, cmd.args)
		case cmdQuit:
			s.quit(cmd.connection)
		case cmdHelp:
			s.help(cmd.connection)
		}
	}
}

func (s *Server) NewClient(conn net.Conn) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln(err)
	}

	c := &connection{
		conn:    conn,
		name:    "anonymous",
		cmds:    s.commands,
		contact: "",
		private: privateKey,
		public:  privateKey.PublicKey,
	}

	log.Printf("new client has joined : %s", conn.RemoteAddr().String())

	c.readInput()
}

func (s *Server) name(c *connection, name string) {
	c.name = name
	s.contacts[name] = c
	c.msg(c, fmt.Sprintf("You will be known as %s", name))
}

func (s *Server) join(c *connection, contactName string) {
	_, ok := s.contacts[contactName]

	if ok && contactName != "" {
		c.contact = contactName
		c.msg(c, fmt.Sprintf("You are now talking to :%s", c.contact))
	} else {
		c.msg(c, "No such user. please check again. Or you can list all connected users with `/list`")
	}
}

func (s *Server) list(c *connection) {
	var contacts []string

	for name := range s.contacts {
		if name != c.name {
			contacts = append(contacts, name)
		}
	}

	c.msg(c, fmt.Sprintf("available users: %s", strings.Join(contacts, ", ")))
}

func (s *Server) msg(c *connection, args []string) {
	_, ok := s.contacts[c.contact]

	if ok && c.contact != "" {
		msg := strings.Join(args[1:], " ")
		msg = c.name + " : " + msg

		publicKey := s.contacts[c.contact].public

		eMsg := encrypt(msg, publicKey)
		log.Printf("Encrypting messages...")

		c.msg(s.contacts[c.contact], eMsg)
		log.Printf("Sending message to %s", c.contact)

	} else {
		c.msg(c, "No one hears you. Please use /list` to list all available users and try again.")
	}
}

func (s *Server) quit(c *connection) {
	log.Printf("%s left the chat.", c.conn.RemoteAddr().String())

	_, ok := s.contacts[c.name]
	if ok {
		delete(s.contacts, c.name)
	}

	c.msg(c, "Closing connection...")
	c.conn.Close()
}

func (s *Server) help(c *connection) {
	c.msg(c, "Usage : /<command> [arguments]\n\n* name : Specify your name.\n* list : List all connected users.\n* join : Specify message recepient.\n* msg  : Send message to recepient.\n* quit : Exit the App.\n* help : List help commands.\n")
}
