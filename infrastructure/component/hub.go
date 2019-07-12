package component

import (
	log "github.com/sirupsen/logrus"
)

// internal
type components struct {
	nonBlockingComponents []Component
	blockingComponents    []Component
}

func (cs *components) Register(component Component) {
	if component.StartBlocking() {
		cs.blockingComponents = append(cs.blockingComponents, component)
	} else {
		cs.nonBlockingComponents = append(cs.nonBlockingComponents, component)
	}

	typ := reflect.TypeOf(component)
	log.Infof("Registers a component: %s", typ.String())
}

func (cs *components) List() []Components {
	components := make([]components, 0)
	components = append(components, cs.nonBlockingComponents...)
	components = append(components, cs.blockingComponents...)
	return components
}

var manager *components = &components{}

// Register registers a component to nonBlockingComponents or blockingComponents
func Register(c Component) {
	manager.Register(c)
}

// GetComponents returns all components
func GetComponents() []Starter {
	return manager.List()
}