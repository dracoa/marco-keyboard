package controls

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

func (a *OneShotAction) TurnOn() {
}

func (a *OneShotAction) TurnOff() {
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
	} else {
		in <- a.OnCommand
	}
}
