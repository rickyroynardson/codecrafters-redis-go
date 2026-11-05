package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		args, err := readCommand(r)
		if err != nil {
			return
		}
		if len(args) == 0 {
			continue
		}

		switch strings.ToUpper(args[0]) {
		case "PING":
			conn.Write([]byte("+PONG\r\n"))
		case "ECHO":
			writeBulkString(conn, args[1])
		default:
			conn.Write([]byte("-ERR unknown command\r\n"))
		}
	}
}
