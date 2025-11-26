package config

type App struct {
	Mode string `env:"APP_MODE,required"`
}

type config struct {
	App *App
}
