package enums

type Action string

var ()

var actions = map[Action]struct{}{}

func (t Action) Valid() bool {
	_, ok := actions[t]
	return ok
}

func (t Action) Value() string {
	return string(t)
}
