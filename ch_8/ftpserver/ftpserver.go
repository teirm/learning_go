package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
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

	// keep state of current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	input := bufio.NewScanner(c)
	for input.Scan() {
		command := input.Text()
		arguments := strings.Split(command, " ")
		switch arguments[0] {
		case "ls":
			listCurrentDirectory(currentDir, c)
		case "pwd":
			getCurrentDirectory(currentDir, c)
		case "chdir":
			if len(arguments) != 2 {
				badCommand(command, c)
			} else {
				fmt.Printf("newDir: %s\n", arguments[1])
				changeCurrentDirectory(arguments[1], &currentDir, c)
				fmt.Printf("currentDir: %s\n", currentDir)
			}
		case "get":
		default:
			badCommand(command, c)
		}
	}
}

// write the current working directory to the
// connection
func getCurrentDirectory(cwd string, c net.Conn) {
	fmt.Fprintf(c, "%s\n", cwd)
}

// write current directory contents to connection
func listCurrentDirectory(cwd string, c net.Conn) {
	entries, err := ioutil.ReadDir(cwd)
	if err != nil {
		fmt.Fprintf(c, "%s\n", err.Error())
	} else {
		for _, entry := range entries {
			fmt.Fprintf(c, "%s\n", entry.Name())
		}
	}
}

// print an error if invalid argument count
func badCommand(command string, c net.Conn) {
	fmt.Fprintf(c, "Invalid command: %s\n", command)
}

// change the current working directory
func changeCurrentDirectory(newDir string, currDir *string, c net.Conn) {
	fmt.Printf("in chdir\n")
	if _, err := os.Stat(newDir); err != nil {
		fmt.Fprintf(c, "%s\n", err.Error())
	} else {
		*currDir = newDir
	}
	fmt.Fprintf(c, "\n")
}
