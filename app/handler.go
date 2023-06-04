package main

import (
	"fmt"
	"strings"
)

func HandleRequests(request string) ([]byte, error) {
	if !strings.Contains(request, "ping") {
		return nil, fmt.Errorf("command not accepted")
	}

	return []byte("+PONG\r\n"), nil
}
