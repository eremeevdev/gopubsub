package pubsub

import (
	"fmt"
	"log"
	"net"
)

// Server - pubsub server
type Server struct {
	Host string
	Port string
}

// NewServer - creates new pubsub server
func NewServer(host, port string) *Server {
	return &Server{host, port}
}

// Start - start pubsub server
func (server *Server) Start() {

	addr := fmt.Sprintf("%s:%s", server.Host, server.Port)

	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	pubsub := NewPubSub()
	go pubsub.Start()

	for {

		conn, err := listen.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go Handler(conn, pubsub)
	}
}
