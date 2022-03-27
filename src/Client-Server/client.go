package Client_Server

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
	id   int
}

func StartClient(id int) (client *Client) {
	client = &Client{}
	client.id = id
	var err error
	client.conn, err = net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		println("Error")
	}
	return
}

func (c *Client) CloseClient() {
	err := c.conn.Close()
	if err != nil {
		panic(err)
	}
	return
}

func (c *Client) SendMessage(msg string) {
	fmt.Fprintf(c.conn, msg)
}

func (c *Client) ReceiveMessage() (msg string) {
	s, _ := bufio.NewReader(c.conn).ReadString('\n')
	msg = fmt.Sprintf("[%d] Return Message From Server: "+s, c.id)
	return
}
