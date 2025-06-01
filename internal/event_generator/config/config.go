package config

import "time"

type Config struct {
	MinRate           int           `yaml:"min_rate"`
	MinDuration       time.Duration `yaml:"min_duration"`
	DnsResolver       string        `yaml:"dns_resolver"`
	LoadBalancePolicy string        `yaml:"load_balance_policy"`
	EventType         string        `yaml:"event_type"`
}
