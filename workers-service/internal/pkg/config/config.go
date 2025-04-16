package config

import "github.com/spf13/viper"

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type Config struct {
	Postgres Postgres

	ServiceHost string
	ServicePort int
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Postgres: Postgres{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			DbName:   viper.GetString("postgres.dbname"),
		},

		ServiceHost: viper.GetString("service.host"),
		ServicePort: viper.GetInt("service.port"),
	}

	return &cfg, nil
}
