package main

import (
	"log"
	"marco-keyboard/keyhook"
	"marco-keyboard/microrobot"
	"strings"
)

const Version = "0.0.0"

func init() {}

type Macro struct {
	Interval byte
	KeyCode  byte
	On       bool
	Robot    *microrobot.MicroRobot
}

func (m *Macro) on() {
	if m.Interval == 0 {
		_ = m.Robot.KeyPress(m.KeyCode)
		log.Printf("start %d", m.KeyCode)
	}
	m.On = true
}

func (m *Macro) off() {
	if m.Interval == 0 {
		_ = m.Robot.KeyRelease(m.KeyCode)
		log.Printf("stop %d", m.KeyCode)
	}
	m.On = false
}

func (m *Macro) Toggle() {
	if m.On {
		m.off()
	} else {
		m.on()
	}
}

func (m *Macro) Config(interval []byte) {
	log.Printf("%d - %v", m.KeyCode, interval)
}

var macros = make(map[byte]*Macro)

func main() {
	log.Printf("Version: %s", Version)
	robot := robot()
	if robot == nil {
		log.Fatal("No micro keyboard found.")
	}
	defer robot.Disconnect()
	startHook(robot)
}

func startHook(robot *microrobot.MicroRobot) {
	out := keyhook.Start()
	for {
		select {
		case cmd := <-out:
			key := cmd.Parameter[0]
			if _, ok := macros[key]; !ok {
				macros[key] = &Macro{
					Interval: 0,
					KeyCode:  key,
					On:       false,
					Robot:    robot,
				}
			}
			if cmd.Config {
				macros[key].Config(cmd.Parameter[1:])
			} else {
				macros[key].Toggle()
			}
		}
	}
}

func robot() *microrobot.MicroRobot {
	var robot *microrobot.MicroRobot
	for com, name := range ListSerial() {
		if strings.Contains(name, "USB") {
			log.Printf("Testing connect to %s: %s", com, name)
			robot = microrobot.Connect(com)
		}
		if robot != nil {
			return robot
		}
	}
	return nil
}
