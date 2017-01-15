package main

import (
	"fmt"

	"flag"

	"github.com/eremeevdev/gopubsub/pubsub"
)

var host = flag.String("host", "", "Host to bind")
var port = flag.String("port", "9999", "Port to listen on")

func main() {

	flag.Parse()

	fmt.Printf("Starting at %v:%v\n", *host, *port)

	server := pubsub.NewServer(*host, *port)

	server.Start()
}
