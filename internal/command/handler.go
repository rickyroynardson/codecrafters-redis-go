package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/store"
)

type Handler struct {
	store *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) Handle(args []string) (string, error) {
	cmd := strings.ToUpper(args[0])

	switch cmd {
	case "PING":
		return resp.SimpleString("PONG"), nil
	case "ECHO":
		if len(args) < 2 {
			return "", errors.New("invalid argument")
		}
		return resp.BulkString(args[1]), nil
	case "SET":
		if len(args) != 3 && len(args) != 5 {
			return "", errors.New("invalid argument")
		}
		return h.set(args), nil
	case "GET":
		if len(args) < 2 {
			return "", errors.New("invalid argument")
		}
		return h.get(args[1]), nil
	case "RPUSH":
		if len(args) < 2 {
			return "", errors.New("invalid argument")
		}
		return h.rpush(args), nil
	}
	return "", fmt.Errorf("unknown command %q", cmd)
}
