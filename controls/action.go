package controls

import (
	"fmt"
	"strings"
	"time"
)

type Action interface {
	Trigger()
}

type IntervalAction struct {
	Control   Control `json:"control"`
	On        bool
	StopChan  chan bool
	Ticker    *time.Ticker
	OnCommand *Command
}

func NewIntervalAction(ctl Control) Action {
	var interval = ctl.Interval
	if interval == 0 {
		interval = 1
	}
	action := &IntervalAction{
		Control:  ctl,
		Ticker:   time.NewTicker(time.Duration(interval) * time.Second),
		StopChan: make(chan bool),
		OnCommand: &Command{
			Name:      ctl.Action,
			Parameter: ctl.Parameter,
		},
	}
	return action
}

func (a *IntervalAction) Trigger() {
	a.On = !a.On
	if a.On {
		go func() {
			for {
				select {
				case <-a.Ticker.C:
					in <- a.OnCommand
				case stop := <-a.StopChan:
					if stop {
						return
					}
				}
			}

		}()
	} else {
		a.StopChan <- true
	}
}

type LongPressAction struct {
	Control    Control `json:"control"`
	On         bool
	OnCommand  *Command
	OffCommand *Command
}

func NewLongPressAction(ctl Control) Action {
	action := &LongPressAction{
		Control: ctl,
		OnCommand: &Command{
			Name:      ctl.Action,
			Parameter: ctl.Parameter,
		},
		OffCommand: &Command{
			Name:      InverseAction(ctl.Action),
			Parameter: ctl.Parameter,
		},
	}
	return action
}

func (a *LongPressAction) Trigger() {
	a.On = !a.On
	if a.On {
		in <- a.OnCommand
	} else {
		in <- a.OffCommand
	}
}

func InverseAction(action string) string {
	if strings.HasSuffix(action, "_press") {
		parts := strings.Split(action, "_")
		return fmt.Sprintf("%s_release", parts[0])
	}
	return action
}
