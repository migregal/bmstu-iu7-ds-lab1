package config

type Config struct {
	HTTPAddr string
}

func ReadConfig() *Config {
	return &Config{
		HTTPAddr: ":8080",
	}
}
