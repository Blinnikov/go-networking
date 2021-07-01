package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	done := make(chan struct{})
	listener := createListener(done)
	conn := createDialer(listener)

	conn.Write([]byte("World"))

	conn.Close()
	<-done
	listener.Close()
	<-done

	var a string
	fmt.Scanln(&a)
}

func createListener(done chan struct{}) net.Listener {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer func() { done <- struct{}{} }()

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Dialed new connection from: %q", conn.LocalAddr())

			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
					}

					if n > 0 {
						log.Printf("received: %q", buf[:n])
					}
				}
			}(conn)
		}
	}()

	return listener
}

func createDialer(l net.Listener) net.Conn {
	conn, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
