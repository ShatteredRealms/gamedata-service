package config

import (
	"context"

	cconfig "github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/sirupsen/logrus"
)

var (
	Version         = "v1.0.0"
	ServiceName     = "GameDataService"
	GlobalConfig    *GamedataConfig
	GlobalConfigErr error
)

func init() {
	GlobalConfig, GlobalConfigErr = NewGamedataConfig(context.Background())
}

type GamedataConfig struct {
	cconfig.BaseConfig `yaml:",inline" mapstructure:",squash"`
	Postgres           cconfig.DBConfig     `yaml:"postgres"`
	Redis              cconfig.DBPoolConfig `yaml:"redis"`
}

func NewGamedataConfig(ctx context.Context) (*GamedataConfig, error) {
	config := &GamedataConfig{
		BaseConfig: cconfig.BaseConfig{
			Server: cconfig.ServerAddress{
				Host: "localhost",
				Port: "8084",
			},
			Keycloak: cconfig.KeycloakConfig{
				BaseURL:      "http://localhost:8080",
				Realm:        "default",
				Id:           "361c9378-9ae3-49e9-96fb-ec03f764766d",
				ClientId:     "sro-gamedata-service",
				ClientSecret: "**********",
			},
			Mode:                "local",
			LogLevel:            logrus.InfoLevel,
			OpenTelemtryAddress: "localhost:4317",
			Kafka: cconfig.ServerAddresses{
				{
					Host: "localhost",
					Port: "29092",
				},
			},
			MigrationPath: "migrations",
		},
		Postgres: cconfig.DBConfig{
			ServerAddress: cconfig.ServerAddress{
				Host: "localhost",
				Port: "5432",
			},
			Name:     "gamedata_service",
			Username: "postgres",
			Password: "password",
		},
		Redis: cconfig.DBPoolConfig{
			Master: cconfig.DBConfig{
				ServerAddress: cconfig.ServerAddress{
					Host: "localhost",
					Port: "7000",
				},
			},
		},
	}

	err := cconfig.BindConfigEnvs(ctx, "sro-gamedata", config)
	return config, err
}
