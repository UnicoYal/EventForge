package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	eventgenerator "github.com/UnicoYal/EventForge/internal/event_generator"
	"github.com/UnicoYal/EventForge/internal/event_generator/wire"
)

var defaultConfigPath = "/usr/local/etc/event_generator.yaml"

func main() {
	var (
		eventType string
		rate      int
		duration  time.Duration
	)

	box, err := wire.InitializeBox(defaultConfigPath)
	if err != nil {
		fmt.Printf("failed to InitializeBox: %v", err)
		os.Exit(1)
	}

	flag.StringVar(&eventType, "event_type", box.Config.EventType, "The type of event to brute force")
	flag.IntVar(&rate, "rate", box.Config.MinRate, "The request per second amount")
	flag.DurationVar(&duration, "duration", box.Config.MinDuration, "Duration to brute force")

	flag.Parse()

	eventgenerator.Generate(eventType, rate, duration)
}
