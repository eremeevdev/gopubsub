package pubsub

import "net"

// Handler - handle new client connections
func Handler(conn net.Conn) {
	conn.Write([]byte("ok\n"))
	conn.Close()
}
