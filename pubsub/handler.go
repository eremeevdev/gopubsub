package pubsub

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Handler - handle new client connections
func Handler(conn net.Conn, pubsub *PubSub) {

	reader := bufio.NewReader(conn)

	data, err := reader.ReadString('\n')

	if err != nil {
		//log.Fatal(err)
		return
	}

	command := strings.Split(strings.TrimSuffix(data, "\n"), " ")

	switch command[0] {

	case "PUBLISH":
		go Publish(conn, command, pubsub)

	case "SUBSCRIBE":
		go Subscribe(conn, command, pubsub)
	}

}

// Publish - handle pub command
func Publish(conn net.Conn, command []string, pubsub *PubSub) {
	defer conn.Close()
	//fmt.Printf("PUBLISH %v TO: %v\n", command[1], command[2:])
	//pubsub.Broadcast(command[1], command[2])
	pubsub.Broadcast <- BroadcastEvent{command[1], command[2]}
}

// Subscribe - handle subscribe command
func Subscribe(conn net.Conn, command []string, pubsub *PubSub) {

	fmt.Println("SUBSCRIBE TO:", command[1:])

	ch := make(chan string)

	defer func() {
		conn.Close()
		pubsub.Unsubscribe <- UnsubscribeEvent{command[1], ch}
	}()

	pubsub.Subscribe <- SubscribeEvent{command[1], ch}

	for msg := range ch {
		fmt.Fprintf(conn, "%s\n", msg)
	}
}
