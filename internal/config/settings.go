package config

type Config struct {
	AppName string
	Server  ServerConfig
	Logger  LoggerConfig
	MongoDB MongoConfig
	Redis   RedisConfig
}

type MongoConfig struct {
	URI      string
	Database string
}

type ServerConfig struct {
	Port string
}

type LoggerConfig struct {
	Level      string
	OutputPath string
	MaxAge     int
	Production bool
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}
