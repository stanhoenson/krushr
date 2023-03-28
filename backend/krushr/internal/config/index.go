package config

type DatabaseConfig struct {
	Name string
}

type Config struct {
	Port     int
	Database DatabaseConfig
}

// TODO maybe change the way this is done it is a little bit weird to be creating Configs
// MOCK FOR NOW EITHER USE https://github.com/joho/godotenv OR https://github.com/spf13/viper
func NewConfig() (*Config, error) {

	return &Config{
		Port: 8080,
		Database: DatabaseConfig{
			Name: "krushr.db",
		},
	}, nil
}
