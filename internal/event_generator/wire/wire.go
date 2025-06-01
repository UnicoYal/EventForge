//go:build wireinject
// +build wireinject

package wire

import (
	"fmt"
	"os"

	"github.com/UnicoYal/EventForge/internal/event_generator/config"
	"github.com/google/wire"
	"github.com/stretchr/testify/assert/yaml"
)

func InitializeBox(configPath string) (*Box, error) {
	wire.Build(
		newBox,

		provideConfig,
	)

	return nil, nil
}

func newBox(conf *config.Config) *Box {
	return &Box{
		Config: *conf,
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
