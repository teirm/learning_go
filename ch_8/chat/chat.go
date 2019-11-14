package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel
type ClientInfo struct {
	channel client
	name    string
}

var (
	entering = make(chan ClientInfo)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	clientNames := make(map[client]string)
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli <- msg:
					// non-blocking send
				default:
					// do nothing
				}
			}
		case cli := <-entering:
			clients[cli.channel] = true
			clientNames[cli.channel] = cli.name
			for _, name := range clientNames {
				cli.channel <- name
			}
		case cli := <-leaving:
			delete(clients, cli)
			delete(clientNames, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 20) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ClientInfo{ch, who}

	period := 30 * time.Second
	tick := time.NewTicker(period)

	go func() {
		for {
			select {
			case <-tick.C:
				tick.Stop()
				conn.Close()
				return
			}
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		tick = time.NewTicker(period)
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
