package enums

type ApplicationMode string

const (
	LocalApplicationMode   ApplicationMode = "LOCAL"
	DevApplicationMode     ApplicationMode = "DEV"
	StagingApplicationMode ApplicationMode = "STAGING"
	ProdApplicationMode    ApplicationMode = "PROD"
)

var applicationModes = map[ApplicationMode]string{
	LocalApplicationMode:   "local",
	DevApplicationMode:     "dev",
	StagingApplicationMode: "staging",
	ProdApplicationMode:    "prod",
}

func (t ApplicationMode) Valid() bool {
	_, ok := applicationModes[t]
	return ok
}

func (t ApplicationMode) Key() string {
	return string(t)
}

func (t ApplicationMode) Value() string {
	return applicationModes[t]
}
