package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var store = make(map[string]string)

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
			if len(args) < 2 {
				conn.Write([]byte("-ERR invalid argument\r\n"))
				continue
			}
			writeBulkString(conn, args[1])
		case "SET":
			if len(args) < 3 {
				conn.Write([]byte("-ERR invalid argument\r\n"))
				continue
			}
			store[args[1]] = args[2]
			conn.Write([]byte("+OK\r\n"))
		case "GET":
			if len(args) < 2 {
				conn.Write([]byte("-ERR invalid argument\r\n"))
				continue
			}
			v, ok := store[args[1]]
			if !ok {
				conn.Write([]byte("$-1\r\n"))
			} else {
				writeBulkString(conn, v)
			}
		default:
			conn.Write([]byte("-ERR unknown command\r\n"))
		}
	}
}
