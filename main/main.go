package main

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"marco-keyboard/controls"
)

func main() {

	actions := controls.Register()
	trigger := func(key string) {
		if fn, ok := actions[key]; ok {
			fn.Trigger()
		}
	}
	/*
		go func(){
			for {
				for k, v := range actions {
					fmt.Printf("%s %t \t", k, v.Status())
				}
				fmt.Println()
				time.Sleep(time.Second)
			}
		}()
	*/
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyUp, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("quit hook")
		hook.End()
	})
	hook.Register(hook.KeyUp, []string{}, func(e hook.Event) {
		if e.Rawcode >= 112 && e.Rawcode <= 123 {
			trigger(fmt.Sprintf("F%d", e.Rawcode-111))
		}
	})
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		trigger(fmt.Sprintf("mouse+%d_%d", e.Button, e.Clicks))
	})

	s := hook.Start()
	<-hook.Process(s)

}
