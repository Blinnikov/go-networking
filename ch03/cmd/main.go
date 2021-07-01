package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	log.Printf("bound to %q", listener.Addr())

	var a string
	fmt.Scanln(&a)
}
