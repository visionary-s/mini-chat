package main

import (
	"flag"
	"log"
	"mini-chat/server"
	"net"
)

func main() {
	var addr = flag.String("addr", ":8090", "Address for the app")
	flag.Parse()

	s := server.NewServer()
	log.Printf("New server created!")

	go s.Run()

	listener, err := net.Listen("tcp", *addr)

	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	} else {
		log.Printf("listening to port : %s", *addr)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}

		go s.NewClient(conn)
		log.Printf("Add new client : %s", conn.RemoteAddr().String())
	}
}
