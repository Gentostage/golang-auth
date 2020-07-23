package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DataBaseUrl string `toml:"data_base"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
	}
}
