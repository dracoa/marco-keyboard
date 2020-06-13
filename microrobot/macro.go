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
			out <- ev
		}
	}()
	return out
}
