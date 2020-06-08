package controls

import (
	"fmt"
	"strings"
)

type Action interface {
	Trigger()
	TurnOn()
	TurnOff()
}

var actions = make(map[string]Action)

func InverseAction(action string) string {
	if strings.HasSuffix(action, "_press") {
		parts := strings.Split(action, "_")
		return fmt.Sprintf("%s_release", parts[0])
	}
	return action
}
