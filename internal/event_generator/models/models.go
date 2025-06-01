package models

import "time"

type Event struct {
	Type      string
	Payload   map[string]any
	Timestamp time.Time
}
