package wire

import (
	"github.com/UnicoYal/EventForge/internal/event_generator/config"
	"github.com/panjf2000/ants/v2"
)

type Box struct {
	Config          config.Config
	IngestionClient interface{}
	WorkerPool      *ants.Pool
}
