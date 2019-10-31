package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Printf("received: %s\n", input.Text())
		fmt.Fprintln(c, "Returning", input.Text())
	}
}
