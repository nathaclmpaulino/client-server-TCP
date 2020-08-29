package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	// Specify the port to connect at the specific port
	port := ":3478"
	connection, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	socket, err := connection.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()

	// Running server until press ctrl c
	for {
		message, err := bufio.NewReader(socket).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(strings.ToUpper(string(message)))

		socket.Write([]byte(strings.ToUpper(string(message)) + "\n"))
	}
	defer connection.Close()
}
