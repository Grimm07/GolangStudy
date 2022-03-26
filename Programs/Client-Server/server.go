package Client_Server

import (
	"fmt"
	"net"
)
import "bufio"

func start_server() {
	// start listening @ 127.0.0.1:8000
	ln, _ := net.Listen("tcp", ":8000")
	// accept a connection
	conn, _ := ln.Accept()

	for {
		// get message and send to output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message))
	}
}
