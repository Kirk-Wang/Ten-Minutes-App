package infrastructure

import (
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/lotteryjs/Ten-Minutes-App/infrastructure/engines"
)

// internal
type manager struct {
	nonBlockingEngines []engines.Engine
	blockingEngines    []engines.Engine
}

func (m *manager) Register(engine engines.Engine) {
	if engine.StartBlocking() {
		m.blockingEngines = append(m.blockingEngines, engine)
	} else {
		m.nonBlockingEngines = append(m.nonBlockingEngines, engine)
	}

	typ := reflect.TypeOf(engine)
	log.Infof("Registers a engine: %s", typ.String())
}

func (m *manager) List() []engines.Engine {
	engines := make([]engines.Engine, 0)
	engines = append(engines, m.nonBlockingEngines...)
	engines = append(engines, m.blockingEngines...)
	return engines
}

var Manager *manager = &manager{}

// Register registers a engine to nonBlockingEngines or blockingEngines
func Register(e engines.Engine) {
	Manager.Register(e)
}

// GetEngines returns all engines
func GetEngines() []engines.Engine {
	return Manager.List()
}
