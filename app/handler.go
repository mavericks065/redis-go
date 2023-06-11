package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

func HandleRequests(request string) ([]byte, error) {
	requestArguments := strings.Split(strings.ToLower(request), "\r\n")
	isCommandPresent := slices.Contains(requestArguments, "command")
	isPingPresent := slices.Contains(requestArguments, "ping")
	if !isCommandPresent && !isPingPresent {
		return nil, fmt.Errorf("command not accepted")
	}

	return []byte("+PONG\r\n"), nil
}
