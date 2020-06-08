package controls

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
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
	_ = json.Unmarshal(bytes, &manifest)
	Listen(manifest.Port, in)
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
