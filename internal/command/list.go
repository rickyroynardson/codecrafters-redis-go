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

func (h *Handler) llen(args []string) string {
	return resp.Integer(h.store.LLen(args[1]))
}

func (h *Handler) lpop(args []string) string {
	count := 1
	if len(args) == 3 {
		count, _ = strconv.Atoi(args[2])
	}

	values := h.store.LPop(args[1], count)
	if len(values) == 0 {
		return resp.NullBulkString()
	} else if len(values) == 1 {
		return resp.BulkString(values[0])
	}
	return resp.Array(values)
}
