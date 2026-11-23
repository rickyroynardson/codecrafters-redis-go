package command

import (
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
)

func (h *Handler) rpush(args []string) string {
	length := h.store.RPush(args[1], args[2:]...)
	return resp.Integer(length)
}

func (h *Handler) lpush(args []string) string {
	length := h.store.LPush(args[1], args[2:]...)
	return resp.Integer(length)
}

func (h *Handler) lrange(args []string) string {
	key := args[1]
	start, _ := strconv.Atoi(args[2])
	stop, _ := strconv.Atoi(args[3])

	values := h.store.LRange(key, start, stop)
	return resp.Array(values)
}
