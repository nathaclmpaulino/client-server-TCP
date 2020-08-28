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

	for {
		socket, err := connection.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		netData, err := bufio.NewReader(socket).ReadString('\n')
		fmt.Print("-> ", strings.ToUpper(string(netData)))

		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
	}
	defer connection.Close()
}
