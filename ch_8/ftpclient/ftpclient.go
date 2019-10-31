package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Fprintln(conn, input.Text())
		output.Scan()
		fmt.Println(output.Text())
	}
}
