package controls

import "time"

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

func (a *IntervalAction) TurnOn() {
	if !a.On {
		a.Trigger()
	}
}

func (a *IntervalAction) TurnOff() {
	if a.On {
		a.Trigger()
	}
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
