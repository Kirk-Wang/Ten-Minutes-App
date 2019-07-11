package component

// Base implements a default component for Component.
type Base struct{}

func (c *Base) Init(ComponentContext) {}

func (c *Base) Setup(ComponentContext) {}

func (c *Base) Start(ComponentContext) {}

func (c *Base) StartBlocking(ComponentContext) {}

func (c *Base) Shutdown(ComponentContext) {}
