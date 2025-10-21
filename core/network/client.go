package network

import (
	"net"
)

type Client struct {
	socket net.Conn
	highId int
	lowId  int
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		socket: conn,
	}
}

func (c *Client) GetSocket() net.Conn {
	return c.socket
}

func (c *Client) GetHighId() int {
	return c.highId
}

func (c *Client) GetLowId() int {
	return c.lowId
}

func (c *Client) SetHighId(id int) {
	c.highId = id
}

func (c *Client) SetLowId(id int) {
	c.lowId = id
}
