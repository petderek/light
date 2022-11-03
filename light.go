package light

import (
	"io"

	"go.bug.st/serial"
)

type command byte

const (
	RED_ON       command = 0x11
	YELLOW_ON    command = 0x12
	GREEN_ON     command = 0x14
	RED_BLINK    command = 0x41
	YELLOW_BLINK command = 0x42
	GREEN_BLINK  command = 0x44
	RED_OFF      command = 0x21
	YELLOW_OFF   command = 0x22
	GREEN_OFF    command = 0x24
	BUZZ_ON      command = 0x18
	BUZZ_BLINK   command = 0x48
	BUZZ_OFF     command = 0x28
)

func Send(portName string, data ...command) error {
	c, err := open(portName, &serial.Mode{BaudRate: 9600})
	if err != nil {
		return err
	}
	defer c.Close()
	msg := make([]byte, len(data))
	for i, v := range data {
		msg[i] = byte(v)
	}
	if _, err = c.Write(msg); err != nil {
		return err
	}
	return nil
}

func Off(portName string) error {
	return Send(portName, GREEN_OFF, RED_OFF, YELLOW_OFF, BUZZ_OFF)
}

var open = internalOpen

func internalOpen(port string, mode *serial.Mode) (io.WriteCloser, error) {
	return serial.Open(port, mode)
}
