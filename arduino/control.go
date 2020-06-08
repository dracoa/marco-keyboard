package arduino

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

var s *serial.Port

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

func Open(port string) {
	c := &serial.Config{Name: port, Baud: 9600}
	var err error
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	_ = s.Close()
}

func BaseTest() {
	Open("COM4?>")
	defer Close()

	KeyWrite(63)
	KeyPress(62)
	time.Sleep(1 * time.Second)
	KeyRelease(62)
	KeyPress(61)
	time.Sleep(1 * time.Second)
	KeyReleaseAll()
	time.Sleep(5 * time.Second)
	MousePress(MouseLeft)
	MouseRelease(MouseLeft)
	time.Sleep(5 * time.Second)
	MouseMove(30, 30)
	MouseClick(MouseRight)
}

func KeyWrite(x byte) {
	write(cmdKeyWrite, x)
}

func KeyPress(x byte) {
	write(cmdKeyPress, x)
}

func KeyRelease(x byte) {
	write(cmdKeyRelease, x)
}

func KeyReleaseAll() {
	write(cmdReleaseAll, 0)
}

func MouseClick(x byte) {
	write(cmdMouseClick, x)
}

func MousePress(x byte) {
	write(cmdMousePress, x)
}

func MouseRelease(x byte) {
	write(cmdMouseRelease, x)
}

func MouseMove(x byte, y byte) {
	write2(cmdMouseMove, x, y)
}

func write(cmd byte, p1 byte) {
	write2(cmd, p1, 0)
}

func write2(cmd byte, p1 byte, p2 byte) {
	_, err := s.Write([]byte{cmd, p1, p2, '\n'})
	if err != nil {
		log.Fatal(err)
	}
}
