package config

import "time"

type Config struct {
	MinRate           int              `yaml:"min_rate"`
	MinDuration       time.Duration    `yaml:"min_duration"`
	IngestionEndpoint string           `yaml:"ingestion_endpoint"`
	LoadBalancePolicy string           `yaml:"load_balance_policy"`
	EventType         string           `yaml:"event_type"`
	WorkerPool        WorkerPoolConfig `yaml:"worker_pool"`
}

type WorkerPoolConfig struct {
	Size           int           `yaml:"size"`
	ReleaseTimeout time.Duration `yaml:"release_timeout"`
}
