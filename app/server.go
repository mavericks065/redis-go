package main

import (
	"fmt"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "6379"
)

func main() {
	fmt.Println("Logs from your program will appear here!")
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	defer listener.Close()
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := listener.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("Failed to accept", err.Error())
	}
	for {
		buf := make([]byte, 1024)
		err = checkDataIsReadable(err, conn, buf)

		request := string(buf)
		response, err := HandleRequests(request)

		handleRequestError(err)
		_, err = conn.Write(response)
		checkResponseError(err)
	}
}

func checkResponseError(err error) {
	if err != nil {
		fmt.Println("Error reading data: ", err.Error())
		os.Exit(1)
	}
}

func handleRequestError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func checkDataIsReadable(err error, conn net.Conn, buf []byte) error {
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading data: ", err.Error())
		os.Exit(1)
	}
	return err
}
