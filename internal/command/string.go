package command

import (
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
)

func (h *Handler) set(args []string) string {
	var expTime time.Time

	if len(args) == 5 {
		n, err := strconv.Atoi(args[4])
		if err != nil {
			return resp.Error("invalid expiration time")
		}
		switch strings.ToUpper(args[3]) {
		case "EX":
			expTime = time.Now().Add(time.Duration(n) * time.Second)
		case "PX":
			expTime = time.Now().Add(time.Duration(n) * time.Millisecond)
		}
	}

	h.store.Set(args[1], args[2], expTime)
	return resp.SimpleString("OK")
}

func (h *Handler) get(key string) string {
	value, ok := h.store.Get(key)
	if !ok {
		return resp.NullBulkString()
	}

	return resp.BulkString(value)
}
