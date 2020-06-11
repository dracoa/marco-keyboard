package microrobot

import (
	"github.com/tarm/serial"
	"log"
)

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

type MicroRobot struct {
	port *serial.Port
}

func Connect(name string) *MicroRobot {
	c := &serial.Config{Name: name, Baud: 9600, ReadTimeout: 1000}
	port, err := serial.OpenPort(c)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = writeLine(port, []byte{cmdEcho, 0x80})
	if err != nil {
		log.Println(err)
		return nil
	}
	bs := make([]byte, 5)
	_, err = port.Read(bs)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return &MicroRobot{port: port}
}

func (r *MicroRobot) Disconnect() error {
	return r.port.Close()
}

func (r *MicroRobot) KeyWrite(char byte) error {
	return writeLine(r.port, []byte{cmdKeyWrite, char})
}

func (r *MicroRobot) KeyPress(char byte) error {
	return writeLine(r.port, []byte{cmdKeyPress, char})
}

func (r *MicroRobot) KeyRelease(char byte) error {
	return writeLine(r.port, []byte{cmdKeyRelease, char})
}

func (r *MicroRobot) KeyReleaseAll() error {
	return writeLine(r.port, []byte{cmdReleaseAll})
}

func (r *MicroRobot) MouseClick(btn byte) error {
	return writeLine(r.port, []byte{cmdMouseClick, btn})
}

func (r *MicroRobot) MousePress(btn byte) error {
	return writeLine(r.port, []byte{cmdMousePress, btn})
}

func (r *MicroRobot) MouseRelease(btn byte) error {
	return writeLine(r.port, []byte{cmdMouseRelease, btn})
}

func (r *MicroRobot) MouseMove(x byte, y byte) error {
	return writeLine(r.port, []byte{cmdMouseMove, x, y})
}

func writeLine(port *serial.Port, bytes []byte) error {
	content := append(bytes, '\n')
	_, err := port.Write(content)
	if err != nil {
		return err
	}
	return nil
}
