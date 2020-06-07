package main

import (
	"fmt"
	"github.com/robotn/gohook"
	"log"
)

func main() {
	// add()
	low()
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	kTargets := []string{"1", "2", "3", "4"}
	// mTargets := []string{}
	for _, t := range kTargets {
		hook.Register(hook.KeyDown, []string{t, "ctrl"}, func(e hook.Event) {
			fmt.Println(e)
		})
	}

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		log.Print("hook: ", ev, ev.Mask, ev.Button, ev.Clicks, ev.X, ev.Y, ev.Direction)
	}
}
