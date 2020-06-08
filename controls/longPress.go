package controls

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

func (a *LongPressAction) TurnOn() {
	if !a.On {
		a.Trigger()
	}
}

func (a *LongPressAction) TurnOff() {
	if a.On {
		a.Trigger()
	}
}

func (a *LongPressAction) Trigger() {
	a.On = !a.On
	if a.On {
		in <- a.OnCommand
	} else {
		in <- a.OffCommand
	}
}
