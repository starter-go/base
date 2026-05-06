package context2

import (
	"context"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////

// 这里提供面向 Console 应用程序的 Adapter 支持

type ConsoleAdapter struct {
}

// GetValue implements Adapter.
func (inst *ConsoleAdapter) GetValue(c *Context, name any) any {
	value := inst
	inst.innerCheckParams(c, name, value)
	return c.Facade.Value(name)
}

// SetValue implements Adapter.
func (inst *ConsoleAdapter) SetValue(c *Context, name any, value any) {

	inst.innerCheckParams(c, name, value)

	cc1 := c.Facade
	cc2 := context.WithValue(cc1, name, value)
	c.Facade = cc2
}

func (inst *ConsoleAdapter) innerCheckParams(c *Context, name, value any) {

	var err error

	if c == nil {
		err = fmt.Errorf("ConsoleAdapter: context2.Context is nil")
	}

	if name == nil {
		err = fmt.Errorf("ConsoleAdapter: name is nil")
	}

	if value == nil {
		err = fmt.Errorf("ConsoleAdapter: value is nil")
	}

	if err != nil {
		panic(err)
	}

}

func (inst *ConsoleAdapter) _impl() Adapter {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

func SetupConsoleContext(cc context.Context) context.Context {

	c1 := cc

	if c1 == nil {
		c1 = context.Background()
	}

	c2, err := Setup(c1, func(name string, ctx *Context) {
		ctx.Adapter = new(ConsoleAdapter)
		c3 := ctx.Facade
		c3 = context.WithValue(c3, name, ctx)
		ctx.Facade = c3
	})

	if err != nil {
		panic(err)
	}

	return c2
}

////////////////////////////////////////////////////////////////////////////////
