// clock is a concurrent server that periodical prints the curren time
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port = flag.String("port", "8000", "port number")
	flag.Parse()

	// create a listening socket
	listener, err := net.Listen("tcp", "localhost:"+(*port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		// listen ot incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		// for each connection create a go routine
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	// close connection on return
	defer c.Close()
	// infinite for loop
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
