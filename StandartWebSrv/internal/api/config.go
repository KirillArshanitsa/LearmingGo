package api

import "StandartWebServer/storage"

type Config struct{
	Port string `toml:"bind_port"`
	LogLevel string `toml:"log_level"`
	ConfigDb *storage.ConfigDb
}

func NewConfig() *Config{
	return &Config{Port: "8080", LogLevel: "info", ConfigDb: storage.NewDbConfig()}
}