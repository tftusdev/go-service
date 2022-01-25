package config

import "github.com/Netflix/go-env"

type Config struct {
	RestPort int `env:"REST_PORT,default=3000"`
	GRPCPort int `env:"GRPC_PORT,default=3001"`
}

func ReadConfigFromEnv() (cfg Config, err error) {
	cfg = Config{}
	_, err = env.UnmarshalFromEnviron(&cfg)
	return
}
