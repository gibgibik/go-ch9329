package ch9329

import "go.bug.st/serial"

type Client struct {
	Port *serial.Port
}

func (c *Client) End() (n int, err error) {
	stopPacket := []byte{
		0x57,
		0xAB,
		0x00,
		0x02,
		0x08,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x0C,
	}
	return c.Port.Write(stopPacket)
}

func NewClient(port *serial.Port) *Client {
	return &Client{
		Port: port,
	}
}
