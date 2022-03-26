package Client_Server

import "net"
import "fmt"

func send() {
	// connect to server
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	fmt.Print("Connected to server on port 8000.")
	for {
		// get message from user
		fmt.Print("Please enter a message to send: ")
		var message string
		fmt.Scanln(&message)
		fmt.Print("Sending message to server.")
		// send message
		fmt.Fprintf(conn, message)
		fmt.Print("Message sent.")
	}
}
