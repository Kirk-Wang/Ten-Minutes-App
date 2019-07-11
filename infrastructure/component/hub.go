package component

type CompWithOptions struct {
	nonBlockingComponents []Component
	blockingComponents    []Component
}

type Components struct {
	comps []CompWithOptions
}

func (cs *Components) Register(nonBlocking, blocking Component) {
	cs.comps = append(cs.comps, CompWithOptions{c, options})
}