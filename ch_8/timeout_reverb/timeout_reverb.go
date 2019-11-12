// reverb server that will disconnect users if silent for 10s

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	timeoutCounter := 0

	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			timeoutCounter = 0
			go echo(c, input.Text(), 1*time.Second)
		}
	}()

	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			timeoutCounter++
			fmt.Println("Counter is at:", timeoutCounter)
			if timeoutCounter == 10 {
				tick.Stop()
				return
			}
		}
	}

}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
