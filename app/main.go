package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/internal/command"
	"github.com/codecrafters-io/redis-starter-go/internal/config"
	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/store"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	cfg := config.Init()

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	if err != nil {
		fmt.Printf("Failed to bind to port %d\n", cfg.Port)
		os.Exit(1)
	}
	defer l.Close()

	store := store.New()

	handler := command.NewHandler(store)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleConnection(conn, handler)
	}
}

func handleConnection(conn net.Conn, handler *command.Handler) {
	defer conn.Close()

	r := resp.NewReader(conn)

	for {
		args, err := r.ReadArray()
		if err != nil {
			break
		}
		if len(args) == 0 {
			continue
		}

		response, err := handler.Handle(args)
		if err != nil {
			conn.Write([]byte(resp.Error(err.Error())))
			continue
		}
		conn.Write([]byte(response))
	}
}
