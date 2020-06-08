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
	Stop      bool
	Interval  time.Duration
	OnCommand *Command
}

func NewIntervalAction(ctl Control) Action {
	var interval = ctl.Interval
	if interval == 0 {
		interval = 1
	}
	action := &IntervalAction{
		Control:  ctl,
		Stop:     true,
		Interval: time.Duration(interval) * time.Second,
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
		a.Stop = false
		go func() {
			for !a.Stop {
				time.Sleep(a.Interval)
				in <- a.OnCommand
			}
		}()
	} else {
		a.Stop = true
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

type OneShotAction struct {
	Control   Control `json:"control"`
	OnCommand *Command
}

func NewOneShotAction(ctl Control) Action {
	action := &OneShotAction{
		Control: ctl,
		OnCommand: &Command{
			Name:      ctl.Action,
			Parameter: ctl.Parameter,
		},
	}
	return action
}

func (a *OneShotAction) Trigger() {
	in <- a.OnCommand
}

func InverseAction(action string) string {
	if strings.HasSuffix(action, "_press") {
		parts := strings.Split(action, "_")
		return fmt.Sprintf("%s_release", parts[0])
	}
	return action
}
