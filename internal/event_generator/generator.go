package eventgenerator

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/UnicoYal/EventForge/internal/event_generator/models"
	"github.com/UnicoYal/EventForge/internal/event_generator/wire"
)

func Generate(box *wire.Box, eventType string, rate int, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		default:
			err := box.WorkerPool.Submit(func() {
				sendEvent(box.IngestionClient, eventType)
			})

			if err != nil {
				fmt.Printf("cannot submit task: %v", err)
			}
		}
	}
}

func sendEvent(client interface{}, eventType string) {
	payload := generatePayload(eventType)
	event := models.Event{
		Type:      eventType,
		Payload:   payload,
		Timestamp: time.Now(),
	}

	fmt.Printf("Send request with payload: %v", event)
}

func generatePayload(eventType string) map[string]any {
	switch eventType {
	case "click":
		return map[string]any{
			"x":      rand.Intn(500),
			"y":      rand.Intn(800),
			"button": "buy",
		}
	case "temperature":
		return map[string]any{
			"value": rand.Float64() * 40,
			"unit":  "C",
		}
	case "transaction":
		return map[string]any{
			"amount":   rand.Float64() * 1000,
			"currency": "RUB",
			"status":   "approved",
		}
	default:
		return map[string]any{}
	}
}
