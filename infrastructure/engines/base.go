package engines

// Base implements a default engine for Engine.
type Base struct{}

func (c *Base) Init(EngineContext) {}

func (c *Base) Setup(EngineContext) {}

func (c *Base) Start(EngineContext) {}

func (s *Base) StartBlocking() bool { return false }

func (c *Base) Shutdown(EngineContext) {}
