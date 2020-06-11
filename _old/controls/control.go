package controls

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"log"
)

type Manifest struct {
	Port     string
	Controls []Control `json:"controls"`
}

type Control struct {
	Display   string      `json:"display"`
	Shortcut  string      `json:"shortcut"`
	Trigger   string      `json:"trigger"`
	Action    string      `json:"action"`
	Interval  uint16      `json:"interval"`
	Parameter interface{} `json:"parameter"`
	Icon      string      `json:"icon"`
}

var manifest Manifest
var in = make(chan *Command)

func init() {
	_ = godotenv.Load()
	bytes, _ := ioutil.ReadFile("./manifest.json")
	err := json.Unmarshal(bytes, &manifest)
	if err != nil {
		log.Panic(err)
	}
	Listen(manifest.Port, in)
}

func HookEvent() {
	actions := Register()
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

func Register() map[string]Action {
	actions = make(map[string]Action)
	for _, ctl := range manifest.Controls {
		var act Action
		switch ctl.Trigger {
		case "interval":
			act = NewIntervalAction(ctl)
		case "long_press":
			act = NewLongPressAction(ctl)
		case "one_shot":
			act = NewOneShotAction(ctl)
		default:
			panic("invalid trigger code " + ctl.Trigger)
		}
		actions[ctl.Shortcut] = act
	}
	return actions
}
