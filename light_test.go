package light

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"go.bug.st/serial"
)

func TestSend(t *testing.T) {
	s := &stub{}
	open = s.testOpen

	err := Send("portName", GREEN_OFF, GREEN_ON)
	if err != nil {
		t.Fatal("expected err to be nil")
	}

	data := s.data.Bytes()
	if command(data[0]) != GREEN_OFF {
		t.Fatal("expected first command to be green off")
	}
	if command(data[1]) != GREEN_ON {
		t.Fatal("expected second command to be green on")
	}
}

func TestOff(t *testing.T) {
	s := &stub{}
	open = s.testOpen

	err := Off("port")
	if err != nil {
		t.Fatal("expected err to be nil")
	}

	data := s.data.Bytes()
	if len(data) != 4 {
		t.Fatal("expected 4 elements")
	}

	needle := make([]byte, 1)
	for _, cmd := range []command{GREEN_OFF, YELLOW_OFF, RED_OFF, BUZZ_OFF} {
		needle[0] = byte(cmd)
		if !bytes.Contains(data, needle) {
			t.Fatal("expected to contain: ", needle)
		}
	}
}

func TestLight(t *testing.T) {
	port := os.Getenv("GO_LIGHT_SERIAL_PORT")
	if port == "" {
		t.Skip("integ test missing GO_LIGHT_SERIAL_PORT envvar")
	}
	open = internalOpen
	if err := Send(port, GREEN_ON); err != nil {
		t.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	if err := Off(port); err != nil {
		t.Fatal(err)
	}
}

type stub struct {
	data bytes.Buffer
}

func (s *stub) testOpen(_ string, _ *serial.Mode) (io.WriteCloser, error) {
	return s, nil
}

func (s *stub) Write(p []byte) (int, error) {
	return s.data.Write(p)
}

func (s *stub) Close() error {
	return nil
}
