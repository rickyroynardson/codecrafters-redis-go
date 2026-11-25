package command

import (
	"strconv"
	"time"

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

func (h *Handler) blpop(args []string) string {
	key := args[1]

	if h.store.Length(key) > 0 {
		result := h.store.LPop(key, 1)
		return resp.Array([]string{key, result[0]})
	}

	timeout, _ := strconv.ParseFloat(args[2], 64)

	ch := make(chan []string)
	h.store.AddBlockedClient(key, ch)

	if timeout == 0 {
		result := <-ch
		return resp.Array(result)
	}

	select {
	case result := <-ch:
		return resp.Array(result)
	case <-time.After(time.Duration(timeout * float64(time.Second))):
		h.store.RemoveBlockedClient(key, ch)
		return resp.NullArray()
	}
}
