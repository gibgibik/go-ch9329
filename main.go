package main

import (
	"log"

	"github.com/gibgibik/go-ch9329/pkg/ch9329"
	"go.bug.st/serial"
)

func main() {
	mode := &serial.Mode{
		BaudRate: 9600, // CH9329 стандартно працює на 9600
	}

	port, err := serial.Open("/dev/serial0", mode) // Заміни на свій порт, якщо інший
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()
	cl := ch9329.NewClient(port)
	cl.End()
}
