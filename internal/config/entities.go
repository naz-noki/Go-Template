package config

type App struct {
	Mode string `env:"APP_MODE,required"`
}

type Postgres struct {
	Host            string `env:"POSTGRES_HOST,required"`
	Port            int    `env:"POSTGRES_PORT,required"`
	Database        string `env:"POSTGRES_DATABASE,required"`
	MaxConns        *int   `env:"POSTGRES_MAX_CONNS"`
	MinConns        *int   `env:"POSTGRES_MIN_CONNS"`
	MaxConnLifetime *int   `env:"POSTGRES_MAX_CONN_LIFETIME"`
	MaxConnIdleTime *int   `env:"POSTGRES_MAX_CONN_IDLE_TIME"`
}

type config struct {
	App      *App
	Postgres *Postgres
}
