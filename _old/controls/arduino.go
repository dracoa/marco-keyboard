package controls

import (
	"github.com/tarm/serial"
	"log"
	"strings"
)

type Command struct {
	Name      string
	Parameter interface{}
}

const (
	cmdEcho         = 0x00
	cmdKeyWrite     = 0x01
	cmdKeyPress     = 0x02
	cmdKeyRelease   = 0x03
	cmdReleaseAll   = 0x04
	cmdMouseClick   = 0x05
	cmdMousePress   = 0x06
	cmdMouseRelease = 0x07
	cmdMouseMove    = 0x08
	MouseLeft       = 0x01
	MouseRight      = 0x02
)

func processCommand(cmd *Command) []byte {
	var code byte
	if strings.HasPrefix(cmd.Name, "key") {
		ch := cmd.Parameter.(string)
		code = ch[0]
	}
	if strings.HasPrefix(cmd.Name, "mouse") {
		ch := cmd.Parameter.(string)
		switch ch {
		case "right":
			code = MouseRight
		default:
			code = MouseLeft
		}
	}
	switch cmd.Name {
	case "key_write":
		return []byte{cmdKeyWrite, code}
	case "key_press":
		return []byte{cmdKeyPress, code}
	case "key_release":
		return []byte{cmdKeyRelease, code}
	case "mouse_click":
		return []byte{cmdMouseClick, code}
	case "mouse_press":
		return []byte{cmdMousePress, code}
	case "mouse_release":
		return []byte{cmdMouseRelease, code}
	case "reset":
		return []byte{cmdReleaseAll, '\n', cmdMouseRelease, MouseLeft, '\n', cmdMouseRelease, MouseRight}
	}
	return nil
}

func writeLine(port *serial.Port, bytes []byte) {
	content := append(bytes, '\n')
	_, err := port.Write(content)
	if err != nil {
		panic(err)
	}
}

func Listen(port string, in chan *Command) {
	c := &serial.Config{Name: port, Baud: 9600}
	var err error
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer func() {
			writeLine(s, processCommand(&Command{
				Name:      "reset",
				Parameter: nil,
			}))
			_ = s.Close()
		}()
		for {
			c := <-in
			cmdBytes := processCommand(c)
			if cmdBytes != nil {
				writeLine(s, cmdBytes)
			}
		}
	}()
}
