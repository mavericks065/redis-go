package main

import (
	"fmt"
	"strings"
)

func HandleRequests(request string) ([]byte, error) {
	args := strings.Split(request, "\n")

	if args[2] != "ping" {
		return nil, fmt.Errorf("command not accepted")
	}

	return []byte("+PONG\r\n"), nil
}
