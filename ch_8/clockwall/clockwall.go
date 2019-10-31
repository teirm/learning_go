package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// parseArgs parses the command line arguments
// expecting Location=hostname:port
//
// returns a map of locations to hostname:port strings
func parseArgs(args []string) (map[string]string, error) {
	wallMap := make(map[string]string)

	for _, arg := range args {
		information := strings.Split(arg, "=")
		if len(information) != 2 {
			return nil, fmt.Errorf("Invalid argument %s\n", arg)
		}
		location := information[0]
		netInfo := information[1]
		wallMap[location] = netInfo
	}
	return wallMap, nil
}

func main() {
	wallMap, err := parseArgs(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	var connections []net.Conn
	for location, netInfo := range wallMap {
		fmt.Printf("%s\t\t", location)
		conn, err := net.Dial("tcp", netInfo)
		if err != nil {
			log.Fatal(err)
		}
		connections = append(connections, conn)
	}
	fmt.Printf("\n")
	handleConnections(connections)
}

func handleConnections(connections []net.Conn) {

	var scanners []*bufio.Scanner
	for _, conn := range connections {
		output := bufio.NewScanner(conn)
		scanners = append(scanners, output)
	}

	for {
		for _, scanner := range scanners {
			scanner.Scan()
			fmt.Printf("%s\t", scanner.Text())
		}
		fmt.Printf("\n")
	}
}
