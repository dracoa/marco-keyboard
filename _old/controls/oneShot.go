package controls

type OneShotAction struct {
	Control   Control `json:"control"`
	On        bool
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

func (a *OneShotAction) Status() bool {
	return a.On
}

func (a *OneShotAction) TurnOn() {
	a.On = true
}

func (a *OneShotAction) TurnOff() {
	a.On = false
}

func (a *OneShotAction) Trigger() {
	if a.Control.Action == "all_on" {
		for _, v := range actions {
			v.TurnOn()
		}
	} else if a.Control.Action == "all_off" {
		for _, v := range actions {
			v.TurnOff()
		}
	} else if a.Control.Action == "all_toggle" {
		if a.On {
			for _, v := range actions {
				v.TurnOff()
			}
		} else {
			for _, v := range actions {
				v.TurnOn()
			}
		}
	} else {
		in <- a.OnCommand
	}
}
