package main

import (
	"fmt"
	"strings"
)

func HandleRequests(request string) ([]byte, error) {
	fmt.Println("request")
	fmt.Println(request)
	args := strings.Split(request, " ")
	fmt.Println("args")
	fmt.Println(args)
	if args[2] != "ping" {
		return nil, fmt.Errorf("command not accepted")
	}

	return []byte("+PONG\r\n"), nil
}
