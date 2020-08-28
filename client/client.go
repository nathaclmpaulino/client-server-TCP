package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// Change here the address value
	address := "127.0.0.1:3478"
	socket, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//Create a buffer to read an input event
		buffer := bufio.NewReader(os.Stdin)
		fmt.Print()
		text, _ := buffer.ReadString('\n')

		fmt.Fprintf(socket, text+"\n")

		message, _ := bufio.NewReader(socket).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
