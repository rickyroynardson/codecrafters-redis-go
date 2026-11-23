package resp

import (
	"fmt"
	"strings"
)

func Error(raw string) string {
	return fmt.Sprintf("-ERR %s\r\n", raw)
}

func SimpleString(raw string) string {
	return fmt.Sprintf("+%s\r\n", raw)
}

func NullBulkString() string {
	return "$-1\r\n"
}

func BulkString(raw string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(raw), raw)
}

func Integer(raw int) string {
	return fmt.Sprintf(":%d\r\n", raw)
}

func EmptyArray() string {
	return "*0\r\n"
}

func Array(values []string) string {
	if len(values) == 0 {
		return EmptyArray()
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("*%d\r\n", len(values)))

	for _, v := range values {
		builder.WriteString(BulkString(v))
	}

	return builder.String()
}

func RawArray(values []string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("*%d\r\n", len(values)))

	for _, v := range values {
		builder.WriteString(v)
	}

	return builder.String()
}
