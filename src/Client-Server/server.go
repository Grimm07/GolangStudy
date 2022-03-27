/*

	Server code for simple HTTP server - accept a message and broadcast to other users

*/

package Client_Server

import (
	"bufio"
	"fmt"
	"net"
)

type Server struct {
	running     bool
	ln          net.Listener
	clients     [10]Client
	clientCount int
	msgCount    int
}

func (s *Server) HandleConn(conn net.Conn) {
	s.clients[s.clientCount] = Client{conn, s.clientCount}
	me := &s.clients[s.clientCount]
	s.clientCount++
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		s.msgCount++
		fmt.Printf("Message Received: ", string(msg))
		conn.Write([]byte(fmt.Sprintf("Message received from client: %d\nMessage count: %d\n", me.id, s.msgCount)))

	}
}

func (s *Server) StartServer() {
	s.running = true
	s.msgCount = 0
	s.clientCount = 0

	// start listening @ 127.0.0.1:8000
	s.ln, _ = net.Listen("tcp", ":8000")

	// accept a connection
	for i := 0; i < 10; i++ {
		conn, _ := s.ln.Accept()
		go s.HandleConn(conn)
	}

}

func (s *Server) CloseServer() {
	err := s.ln.Close()
	if err != nil {
		panic(err)
	}
}

func (s *Server) IsRunning() bool {
	return s.running
}
