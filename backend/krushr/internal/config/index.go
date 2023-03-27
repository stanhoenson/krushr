package config

type DatabaseConfig struct {
	Name string
}

type Config struct {
	Port     int
	Database DatabaseConfig
}

func NewConfig() (*Config, error) {
	//MOCK FOR NOW EITHER USE https://github.com/joho/godotenv OR https://github.com/spf13/viper

	return &Config{
		Port: 8080,
		Database: DatabaseConfig{
			Name: "krushr.db",
		},
	}, nil
}
