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
type UserService struct {
	Host string
	Port int
}

type Services struct {
	WorkersService WorkersService
	ApiGateway     ApiGateway
	UserService    UserService
}

type TLS struct {
	KeyFile  string
	CertFile string
}

type Config struct {
	Services Services
	TLS      TLS
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
			UserService: UserService{
				Host: viper.GetString("services.user-service.host"),
				Port: viper.GetInt("services.user-service.port"),
			},
		},
		TLS: TLS{
			KeyFile:  viper.GetString("tls.key_file"),
			CertFile: viper.GetString("tls.cert_file"),
		},
	}

	return &cfg, nil
}
