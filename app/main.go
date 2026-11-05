package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type record struct {
	key   string
	value string
	exp   time.Time
}

var store = make(map[string]record)

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
			if len(args) != 3 && len(args) != 5 {
				conn.Write([]byte("-ERR invalid argument\r\n"))
				continue
			}

			rec := record{key: args[1], value: args[2]}
			if len(args) == 5 {
				n, err := strconv.Atoi(args[4])
				if err != nil {
					conn.Write([]byte("-ERR invalid expiration time\r\n"))
					continue
				}
				switch strings.ToUpper(args[3]) {
				case "EX":
					rec.exp = time.Now().Add(time.Duration(n) * time.Second)
				case "PX":
					rec.exp = time.Now().Add(time.Duration(n) * time.Millisecond)
				}
			}
			store[args[1]] = rec
			conn.Write([]byte("+OK\r\n"))
		case "GET":
			if len(args) < 2 {
				conn.Write([]byte("-ERR invalid argument\r\n"))
				continue
			}
			v, ok := store[args[1]]
			if ok && !v.exp.IsZero() && time.Now().After(v.exp) {
				delete(store, args[1])
				ok = false
			}
			if !ok {
				conn.Write([]byte("$-1\r\n"))
			} else {
				writeBulkString(conn, v.value)
			}
		default:
			conn.Write([]byte("-ERR unknown command\r\n"))
		}
	}
}
