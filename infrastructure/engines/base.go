package engines

// Base implements a default engine for Engine.
type Base struct{}

func (c *Base) Init(EngineContext) {}

func (c *Base) Setup(EngineContext) {}

func (c *Base) Start(EngineContext) {}

func (c *Base) StartBlocking(EngineContext) {}

func (c *Base) Shutdown(EngineContext) {}
