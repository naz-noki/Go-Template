package config

import (
	"github.com/caarlos0/env/v11"
)

var (
	Config = new(config)
)

func Load() error {
	// App configuration
	app := new(App)
	if err := env.Parse(app); err != nil {
		return err
	}
	Config.App = app
	// PostgreSQL configuration
	postgres := new(Postgres)
	if err := env.Parse(postgres); err != nil {
		return err
	}
	Config.Postgres = postgres

	return nil
}
