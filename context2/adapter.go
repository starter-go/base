package context2

type Adapter interface {
	GetValue(c *Context, name any) any

	SetValue(c *Context, name any, value any)
}
