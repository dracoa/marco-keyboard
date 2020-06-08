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
	cmdKeyWrite     = 0x00
	cmdKeyPress     = 0x01
	cmdKeyRelease   = 0x02
	cmdReleaseAll   = 0x03
	cmdMouseClick   = 0x04
	cmdMousePress   = 0x05
	cmdMouseRelease = 0x06
	cmdMouseMove    = 0x07
	MouseLeft       = 0x01
	MouseRight      = 0x02
)

func processCommand(cmd *Command) []byte {
	log.Println(cmd)
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
			writeLine(s, []byte{cmdReleaseAll})
			_ = s.Close()
		}()
		for {
			c := <-in
			cmdBytes := processCommand(c)
			writeLine(s, cmdBytes)
		}
	}()
}
