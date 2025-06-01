//go:build wireinject
// +build wireinject

package wire

import (
	"fmt"
	"os"

	"github.com/UnicoYal/EventForge/internal/event_generator/config"
	"github.com/google/wire"
	"github.com/panjf2000/ants/v2"
	"github.com/stretchr/testify/assert/yaml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitializeBox(configPath string) (*Box, func(), error) {
	wire.Build(
		newBox,

		provideConfig,
		provideIngestionConn,
		provideWorkerPool,
	)

	return nil, nil, nil
}

func newBox(
	conf *config.Config,
	ingestionConn *grpc.ClientConn,
	wp *ants.Pool,
) *Box {
	return &Box{
		Config:        *conf,
		IngestionConn: ingestionConn,
		WorkerPool:    wp,
	}
}

func provideConfig(configPath string) (*config.Config, error) {
	cfg := &config.Config{}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	if err = yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("cannot unmarshall config: %w", err)
	}

	return cfg, nil
}

func provideIngestionClient(config *config.Config) (interface{}, func(), error) {
	conn, err := grpc.NewClient(
		config.IngestionEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy": "%s"}`, config.LoadBalancePolicy)),
	)

	if err != nil {
		return nil, nil, fmt.Errorf("cannot establish conn with ingestion service: %w", err)
	}

	cleanup := func() {
		conn.Close()
	}

	return pb.NewIngestionClient(conn), cleanup, nil
}

func provideWorkerPool(config *config.Config) (*ants.Pool, func(), error) {
	pool, err := ants.NewPool(
		config.WorkerPool.Size,
		ants.WithMaxBlockingTasks(0),
		ants.WithNonblocking(true),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot initialize worker pool: %w", err)
	}

	cleanup := func() {
		pool.ReleaseTimeout(config.WorkerPool.ReleaseTimeout)
	}

	return pool, cleanup, nil
}
