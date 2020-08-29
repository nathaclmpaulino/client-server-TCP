package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// Change here the address value
	address := "127.0.0.1:3478"
	socket, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	for {
		//Create a buffer to read an input event
		buffer := bufio.NewReader(os.Stdin)
		text, _ := buffer.ReadString('\n')

		fmt.Fprintf(socket, text)
		count++
		message, _ := bufio.NewReader(socket).ReadString('\n')
		fmt.Print("Resposta da mensagem " + fmt.Sprintf("%d", count) + ":" + message)
	}
}
