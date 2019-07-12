package engines

// Engine is the interface that represent a engine.
type Engine interface {
	Init(EngineContext)
	Setup(EngineContext)
	Start(EngineContext)
	StartBlocking() bool
	Shutdown(EngineContext)
}
