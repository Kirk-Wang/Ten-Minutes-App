package component

import (
	log "github.com/sirupsen/logrus"
)

type Components struct {
	nonBlockingComponents []Component
	blockingComponents    []Component
}

// Register registers a component to nonBlockingComponents or blockingComponents
func (cs *Components) Register(component Component) {
	if component.StartBlocking() {
		cs.blockingComponents = append(cs.blockingComponents, component)
	} else {
		cs.nonBlockingComponents = append(cs.nonBlockingComponents, component)
	}

	typ := reflect.TypeOf(component)
	log.Infof("Registers a component: %s", typ.String())
}

// List returns all components
func (cs *Components) List() []Components {
	components := make([]components, 0)
	components = append(components, cs.nonBlockingComponents...)
	components = append(components, cs.blockingComponents...)
	return components
}