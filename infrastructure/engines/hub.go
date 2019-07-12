package engines

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

// internal
type engines struct {
	nonBlockingEngines []Engine
	blockingEngines    []Engine
}

func (cs *engines) Register(engine Engine) {
	if engine.StartBlocking() {
		cs.blockingEngines = append(cs.blockingEngines, engine)
	} else {
		cs.nonBlockingEngines = append(cs.nonBlockingEngines, engine)
	}

	typ := reflect.TypeOf(engine)
	log.Infof("Registers a engine: %s", typ.String())
}

func (cs *engines) List() []Engine {
	engines := make([]Engine, 0)
	engines = append(engines, cs.nonBlockingEngines...)
	engines = append(engines, cs.blockingEngines...)
	return engines
}

var manager *engines = &engines{}

// Register registers a engine to nonBlockingEngines or blockingEngines
func Register(c Engine) {
	manager.Register(c)
}

// GetEngines returns all engines
func GetEngines() []Engine {
	return manager.List()
}
