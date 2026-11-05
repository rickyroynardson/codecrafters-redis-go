package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func readLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(line, "\r\n"), nil
}

func readCommand(r *bufio.Reader) ([]string, error) {
	line, err := readLine(r)
	if err != nil {
		return nil, err
	}

	if len(line) == 0 || line[0] != '*' {
		return nil, fmt.Errorf("expected array, got %q", line)
	}
	n, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, err
	}

	args := make([]string, 0, n)
	for range n {
		line, err := readLine(r)
		if err != nil {
			return nil, err
		}
		if len(line) == 0 || line[0] != '$' {
			return nil, fmt.Errorf("expected bulk string, got %q", line)
		}
		size, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		data := make([]byte, size+2)
		if _, err := io.ReadFull(r, data); err != nil {
			return nil, err
		}
		args = append(args, string(data[:size]))
	}
	return args, nil
}

func writeBulkString(w io.Writer, s string) error {
	_, err := fmt.Fprintf(w, "$%d\r\n%s\r\n", len(s), s)
	return err
}
