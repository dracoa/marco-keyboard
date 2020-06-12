package main

import (
	"encoding/json"
	"errors"
	"log"
	"marco-keyboard/microrobot"
	"marco-keyboard/screen"
	"marco-keyboard/websvr"
	"strings"
)

const Version = "0.0.0"

type WsCommand struct {
	Cmd     string `json:"cmd"`
	Screen  int    `json:"screen"`
	Quality int    `json:"quality"`
	KeyCode byte   `json:"key_code"`
	PosX    byte   `json:"pos_x"`
	PosY    byte   `json:"pos_y"`
}

func init() {}

func main() {
	log.Printf("Version: %s", Version)
	server := websvr.Start(":8088")
	robot := robot()
	if robot == nil {
		log.Fatal("No micro keyboard found.")
	}
	defer robot.Disconnect()
	go startHook(server)
	for {
		select {
		case income := <-server.In:
			req := &WsCommand{}
			_ = json.Unmarshal(income, &req)
			if req.Cmd == "screen_capture" {
				server.SendBinary(screen.CaptureScreen(req.Screen, req.Quality))
			} else if req.Cmd == "mouse_move" {
				_ = robot.MouseMove(req.PosX, req.PosY)
			} else {
				_ = executeEvent(robot, req.Cmd, req.KeyCode)
			}
		}
	}
}

func executeEvent(robot *microrobot.MicroRobot, event string, keyCode byte) error {
	switch event {
	case "key_write":
		return robot.KeyWrite(keyCode)
	case "key_press":
		return robot.KeyPress(keyCode)
	case "key_release":
		return robot.KeyRelease(keyCode)
	case "key_release_all":
		return robot.KeyReleaseAll()
	case "mouse_click":
		return robot.MouseClick(keyCode)
	case "mouse_press":
		return robot.MousePress(keyCode)
	case "mouse_release":
		return robot.MouseRelease(keyCode)
	}
	return errors.New("invalid command")
}

func startHook(server *websvr.Server) {
	out := microrobot.Start()
	for {
		select {
		case cmd := <-out:
			server.SendJson(cmd)
		}
	}
}

func robot() *microrobot.MicroRobot {
	var robot *microrobot.MicroRobot
	for com, name := range microrobot.ListSerial() {
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
