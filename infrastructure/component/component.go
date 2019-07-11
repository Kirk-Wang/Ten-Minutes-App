package component

// Context
type ComponentContext map[string]interface{}

// Component is the interface that represent a component.
type Component interface {
	Init(ComponentContext)
	Setup(ComponentContext)
	Start(ComponentContext)
	StartBlocking() bool
	Stop(ComponentContext)
}
