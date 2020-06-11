package keyhook

import (
	hook "github.com/robotn/gohook"
	"log"
)

type Command struct{}

func Start() chan Command {
	out := make(chan Command)
	ctrl := false
	buffer := make([]byte, 0)
	go func() {
		EvChan := hook.Start()
		defer hook.End()
		for ev := range EvChan {
			if ev.Kind >= 3 && ev.Kind <= 5 {
				if ctrl {
					if ev.Rawcode == 162 && ev.Kind == 5 {
						log.Println(buffer)
						buffer = make([]byte, 0)
						ctrl = false
					}
				} else {
					if ev.Rawcode == 162 && ev.Kind == 4 {
						ctrl = true
					}
				}
				if ev.Kind == 5 && ev.Rawcode != 162 {
					if ctrl {
						buffer = append(buffer, byte(ev.Rawcode))
					}
				}
			}
		}
	}()
	return out
}
