package keyhook

import (
	hook "github.com/robotn/gohook"
)

type Command struct {
	Config    bool
	Parameter []byte
}

func Start() chan Command {
	out := make(chan Command)
	buffer := make([]byte, 0)
	go func() {
		EvChan := hook.Start()
		defer hook.End()
		for ev := range EvChan {
			if ev.Kind == 5 {
				if ev.Mask == 10 {
					buffer = append(buffer, byte(ev.Rawcode))
				} else {
					if len(buffer) > 0 {
						out <- Command{
							Config:    true,
							Parameter: buffer,
						}
						buffer = make([]byte, 0)
					} else if ev.Rawcode != 162 {
						if ev.Mask == 2 {
							out <- Command{
								Parameter: []byte{byte(ev.Rawcode)},
							}
						}
					}
				}
			}
		}
	}()
	return out
}
