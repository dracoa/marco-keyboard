package microrobot

import (
	hook "github.com/robotn/gohook"
)

func Start() chan hook.Event {
	out := make(chan hook.Event)
	go func() {
		EvChan := hook.Start()
		defer hook.End()
		for ev := range EvChan {
			if ev.Kind == 5 || ev.Kind == 8 || ev.Kind == 9 {
				out <- ev
			}
		}
	}()
	return out
}
