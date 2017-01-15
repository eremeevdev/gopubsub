package pubsub

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Handler - handle new client connections
func Handler(conn net.Conn) {

	reader := bufio.NewReader(conn)

	data, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	command := strings.Split(strings.TrimSuffix(data, "\n"), " ")

	switch command[0] {

	case "PUBLISH":
		go Publish(conn, command)

	case "SUBSCRIBE":
		go Subscribe(conn, command)
	}

}

// Publish - handle pub command
func Publish(conn net.Conn, command []string) {
	defer conn.Close()
	fmt.Printf("PUBLISH %v TO: %v\n", command[1], command[2:])
}

// Subscribe - handle subscribe command
func Subscribe(conn net.Conn, command []string) {
	fmt.Println("SUBSCRIBE TO:", command[1:])
}
