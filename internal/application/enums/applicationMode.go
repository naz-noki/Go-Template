package enums

type ApplicationMode string

const (
	LocalApplicationMode ApplicationMode = "local"
	DevApplicationMode   ApplicationMode = "dev"
	ProdApplicationMode  ApplicationMode = "prod"
)

var applicationModes = map[ApplicationMode]struct{}{
	LocalApplicationMode: {},
	DevApplicationMode:   {},
	ProdApplicationMode:  {},
}

func (t ApplicationMode) Valid() bool {
	_, ok := applicationModes[t]
	return ok
}

func (t ApplicationMode) Value() string {
	return string(t)
}
