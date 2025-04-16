package config

import "github.com/spf13/viper"

type WorkersService struct {
	Host string
	Port int
}
type ApiGateway struct {
	Host string
	Port int
}

type Services struct {
	WorkersService WorkersService
	ApiGateway     ApiGateway
}

type Config struct {
	Services Services
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Services: Services{
			WorkersService: WorkersService{
				Host: viper.GetString("services.workers-service.host"),
				Port: viper.GetInt("services.workers-service.port"),
			},
			ApiGateway: ApiGateway{
				Host: viper.GetString("services.api-gateway.host"),
				Port: viper.GetInt("services.api-gateway.port"),
			},
		},
	}

	return &cfg, nil
}
