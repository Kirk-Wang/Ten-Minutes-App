package component

type Components struct {
	nonBlockingComponents []Component
	blockingComponents    []Component
}

// Register registers a component to nonBlockingComponents or blockingComponents
func (cs *Components) Register(nonBlocking, blocking Component) {}

// List returns all components
func (cs *Components) List() []Components {
	components := make([]components, 0)
	components = append(components, cs.nonBlockingComponents...)
	components = append(components, cs.blockingComponents...)
	return components
}