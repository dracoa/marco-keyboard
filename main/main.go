package main

import (
	"encoding/json"
	"log"
	"marco-keyboard/screen"
	"marco-keyboard/websvr"
)

type WebRequest struct {
	Command   string      `json:"command"`
	Parameter interface{} `json:"parameter"`
}

type MouseEvent struct {
	EventType  string `json:"event_type"`
	ButtonCode string `json:"button_code"`
}

type KeyEvent struct {
	EventType string `json:"event_type"`
	KeyCode   string `json:"key_code"`
}

var handlers = make(map[string]func(interface{}) []byte)

func main() {
	log.Println("Keyboard Control Server v1.0.0")
	server := websvr.Start("localhost:2303")
	handlers["KeyboardAction"] = KeyboardAction
	handlers["MouseAction"] = MouseAction
	handlers["ScreenCapture"] = ScreenCapture
	for {
		select {
		case income := <-server.In:
			req := &WebRequest{}
			_ = json.Unmarshal(income, &req)
			if val, ok := handlers[req.Command]; ok {
				server.Out <- val(req.Parameter)
			} else {
				panic("command not found: " + req.Command)
			}
		}
	}
}

func KeyboardAction(para interface{}) []byte {
	keyEvent := &KeyEvent{}
	bytes, err := json.Marshal(para)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &keyEvent)
	if err != nil {
		panic(err)
	}
	return []byte(`{}`)
}

func MouseAction(para interface{}) []byte {
	mouseEvent := &MouseEvent{}
	bytes, err := json.Marshal(para)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &mouseEvent)
	if err != nil {
		panic(err)
	}
	return []byte(`{}`)
}

func ScreenCapture(para interface{}) []byte {
	return screen.Capture()
}
