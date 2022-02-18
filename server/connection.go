package server

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"log"
	"net"
	"strings"
)

type connection struct {
	conn    net.Conn
	name    string
	contact string
	cmds    chan<- cmd
	private *rsa.PrivateKey
	public  rsa.PublicKey
}

func (c *connection) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		command := strings.TrimSpace(args[0])

		switch command {
		case "/name":
			c.cmds <- cmd{
				id:         cmdName,
				connection: c,
				args:       args,
			}
		case "/join":
			c.cmds <- cmd{
				id:         cmdJoin,
				connection: c,
				args:       args,
			}
		case "/list":
			c.cmds <- cmd{
				id:         cmdList,
				connection: c,
			}
		case "/msg":
			c.cmds <- cmd{
				id:         cmdMsg,
				connection: c,
				args:       args,
			}
		case "/quit":
			c.cmds <- cmd{
				id:         cmdQuit,
				connection: c,
			}
		case "/help":
			c.cmds <- cmd{
				id:         cmdHelp,
				connection: c,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", command))
			c.msg(c, "* use '/help' to show all available commands")
		}
	}
}

func (c *connection) err(err error) {
	_, e := c.conn.Write([]byte("err: " + err.Error() + "\n"))
	if e != nil {
		log.Fatalln("Failed to write to the connection", e)
	}
}

func (c *connection) msg(x *connection, msg string) {
	if c.private != x.private {
		dMsg := decrypt(msg, *x.private)

		_, e := x.conn.Write([]byte("> " + dMsg + "\n"))
		if e != nil {
			log.Fatalln("Failed to write through connection")
		}
	} else {
		_, e := x.conn.Write([]byte("> " + msg + "\n"))
		if e != nil {
			log.Fatalln("Failed to write through connection")
		}
	}
}
