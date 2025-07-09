package ch9329

import (
	"errors"

	"go.bug.st/serial"
)

type Client struct {
	Port serial.Port
}

func (c *Client) SendKey(modifier byte, key string) (n int, err error) {
	val, ok := HidKeycodes[key]
	if !ok {
		return 0, errors.New("no such key")
	}
	packet := []byte{
		0x57,            //default header
		0xAB,            //default header
		0x00,            // constant
		CmdSendKeyboard, // command
		0x08,            //length, constant for keyboard
		modifier,
		0x00,
		val,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
	}
	var checkSum int
	for _, k := range packet {
		checkSum += int(k)
	}
	packet = append(packet, byte(checkSum%256))
	return c.Port.Write(packet)
}

func (c *Client) EndKey() (n int, err error) {
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

func NewClient(port serial.Port) *Client {
	return &Client{
		Port: port,
	}
}
