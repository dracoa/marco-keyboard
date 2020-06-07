package main

import (
	"encoding/json"
	"fmt"
	"github.com/robotn/gohook"
	"github.com/tarm/serial"
	"log"
)

type Command struct {
	Cmd  string `json:"cmd"`
	Char uint8 `json:"char"`
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
}

func main() {
	// add()
	// low()
	c := &serial.Config{Name: "COM6", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	testCmd := &Command{
		Cmd:  "mouse_move",
		Char: 0,
		X:    100,
		Y:    100,
	}
	bytes, _ := json.Marshal(&testCmd)
	_, err = s.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}

}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		if ev.Kind == hook.KeyUp {
			fmt.Println("hook: ", ev)
		}
	}
}