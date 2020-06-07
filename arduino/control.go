package arduino

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

var s *serial.Port

const (
	KEY_WRITE       = 0x00
	KEY_PRESS       = 0x01
	KEY_RELEASE     = 0x02
	KEY_RELEASE_ALL = 0x03
	MOUSE_CLICK     = 0x04
	MOUSE_PRESS     = 0x05
	MOUSE_RELEASE   = 0x06
	MOUSE_MOVE      = 0x07
	MOUSE_LEFT      = 0x01
	MOUSE_RIGHT     = 0x02
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
	s.Close()
}

func BaseTest() {
	Open("COM6")
	defer Close()

	KeyWrite(63)
	KeyPress(62)
	time.Sleep(1 * time.Second)
	KeyRelease(62)
	KeyPress(61)
	time.Sleep(1 * time.Second)
	KeyReleaseAll()
	time.Sleep(3 * time.Second)
	MousePress(MOUSE_LEFT)
	MouseRelease(MOUSE_LEFT)
	time.Sleep(3 * time.Second)
	MouseClick(MOUSE_RIGHT)
}

func KeyWrite(x byte) {
	write(KEY_WRITE, x)
}

func KeyPress(x byte) {
	write(KEY_PRESS, x)
}

func KeyRelease(x byte) {
	write(KEY_RELEASE, x)
}

func KeyReleaseAll() {
	write(KEY_RELEASE_ALL, 0)
}

func MouseClick(x byte) {
	write(MOUSE_CLICK, x)
}

func MousePress(x byte) {
	write(MOUSE_PRESS, x)
}

func MouseRelease(x byte) {
	write(MOUSE_RELEASE, x)
}

func MouseMove(x byte, y byte) {
	write2(MOUSE_MOVE, x, y)
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
