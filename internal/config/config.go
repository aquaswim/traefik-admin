package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Addr   string `env:"APP_ADDRESS" envDefault:":3000"`
	DBPath string `env:"DB_PATH,required"`
}

func LoadConfig() (Config, error) {
	return env.ParseAs[Config]()
}

func MustLoadConfig() Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
