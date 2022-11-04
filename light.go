package light

import (
	"errors"
	"io"
	"strings"

	"go.bug.st/serial"
)

type command byte

const (
	ZERO         command = 0x00
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

var enableMap = map[string]command{
	"red":    RED_ON,
	"yellow": YELLOW_ON,
	"green":  GREEN_ON,
}

var disableMap = map[string]command{
	"red":    RED_OFF,
	"yellow": YELLOW_OFF,
	"green":  GREEN_OFF,
}

// Infer turns a human-readable word like 'green','on' into a serial command
func Infer(name string, toggle string) (command, error) {
	var m map[string]command
	switch {
	case strings.EqualFold(toggle, "on"):
		m = enableMap
	case strings.EqualFold(toggle, "off"):
		m = disableMap
	default:
		return ZERO, errors.New("toggle not recognized: " + toggle)
	}

	for k, v := range m {
		if strings.EqualFold(k, name) {
			return v, nil
		}
	}
	return ZERO, errors.New("name not recognized: " + name)
}

func Send(portName string, data ...command) error {
	c, err := open(portName, &serial.Mode{BaudRate: 9600})
	if err != nil {
		return err
	}
	defer c.Close()
	msg := make([]byte, len(data))
	for i, v := range data {
		if v == ZERO {
			return errors.New("zero value sent in list of commands")
		}
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
