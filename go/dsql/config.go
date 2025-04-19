package dsql

import (
	"context"

	"github.com/sethvargo/go-envconfig"

	"github.com/tusmasoma/go-microservice-k8s/go/pkg/log"
)

const (
	serverPrefix = "SERVER_"
)

type DBConfig struct {
	Host            string `env:"HOST, required"`
	Port            string `env:"PORT, required"`
	User            string `env:"USER, required"`
	Password        string `env:"PASSWORD, required"`
	DBName          string `env:"DB_NAME, required"`
	ClusterEndpoint string `env:"CLUSTER_ENDPOINT, required"`
}

func NewDBConfig(ctx context.Context, dbPrefix string) (*DBConfig, error) {
	conf := &DBConfig{}
	pl := envconfig.PrefixLookuper(dbPrefix, envconfig.OsLookuper())
	if err := envconfig.ProcessWith(ctx, &envconfig.Config{
		Target:   conf,
		Lookuper: pl,
	}); err != nil {
		log.Error("Failed to load database config", log.Ferror(err))
		return nil, err
	}
	return conf, nil

}
