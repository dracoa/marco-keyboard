package controls

import (
	"fmt"
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
	switch cmd.Name {
	case "key_write":
		return []byte{cmdKeyWrite, code}
	case "key_press":
		return []byte{cmdKeyPress, code}
	case "key_release":
		return []byte{cmdKeyRelease, code}
	}
	return nil
}

func Listen(port string, in chan *Command) {
	c := &serial.Config{Name: port, Baud: 9600}
	var err error
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			c := <-in
			cmdBytes := processCommand(c)
			fmt.Println(cmdBytes)
			if cmdBytes != nil {
				_, err := s.Write(cmdBytes)
				if err != nil {
					panic(err)
				}
			}
		}
		s.Close()
	}()
}
