package component

type Components struct {
	nonBlockingComponents []Component
	blockingComponents    []Component
}

func (cs *Components) Register(nonBlocking, blocking Component) {}

func (cs *Components) List() []CompWithOptions {}