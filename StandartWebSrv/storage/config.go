package storage

type ConfigDb struct{
	DbUri string `toml:"database_uri" env:"database_uri"`
}

func NewDbConfig() *ConfigDb {
	return &ConfigDb{ }
}