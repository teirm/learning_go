package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// display output
func displayOutput(output *bufio.Scanner) {
	for output.Scan() {
		fmt.Println(output.Text())
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewScanner(conn)
	go displayOutput(output)
	for input.Scan() {
		fmt.Fprintln(conn, input.Text())
	}
}
