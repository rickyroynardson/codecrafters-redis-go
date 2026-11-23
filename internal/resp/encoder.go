package resp

import "fmt"

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
