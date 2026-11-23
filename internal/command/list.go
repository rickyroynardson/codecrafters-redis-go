package command

import "github.com/codecrafters-io/redis-starter-go/internal/resp"

func (h *Handler) rpush(args []string) string {
	length := h.store.RPush(args[1], args[2:]...)
	return resp.Integer(length)
}
