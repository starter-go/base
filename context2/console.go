package context2

import (
	"context"
	"fmt"
	"time"
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

	f1, err := c.GetFacade()
	if err != nil {
		panic(err)
	}

	f2 := f1.(*innerConsoleContextFacade)
	f2.putValue(name, value)
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

type innerConsoleContextFacade struct {
	head context.Context
	ctx2 *Context
}

// Deadline implements context.Context.
func (inst *innerConsoleContextFacade) Deadline() (deadline time.Time, ok bool) {
	h := inst.head
	if h == nil {
		t0 := time.Unix(0, 0)
		return t0, false
	}
	return h.Deadline()
}

// Done implements context.Context.
func (inst *innerConsoleContextFacade) Done() <-chan struct{} {
	h := inst.head
	if h == nil {
		return nil
	}
	return h.Done()
}

// Err implements context.Context.
func (inst *innerConsoleContextFacade) Err() error {
	h := inst.head
	if h == nil {
		return nil
	}
	return h.Err()
}

// Value implements context.Context.
func (inst *innerConsoleContextFacade) Value(key any) any {
	h := inst.head
	if h == nil {
		return nil
	}
	return h.Value(key)
}

func (inst *innerConsoleContextFacade) putValue(name, value any) {

	if name == nil || value == nil {
		return
	}

	cc := inst.head
	if cc == nil {
		cc = context.Background()
	}

	cc = context.WithValue(cc, name, value)

	inst.head = cc
	inst.ctx2.PutKey(name)
}

func (inst *innerConsoleContextFacade) _impl() context.Context {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

func SetupConsoleContext(cc context.Context) context.Context {

	c1 := cc
	if c1 == nil {
		c1 = context.Background()
	}

	c2, err := Setup(c1, func(name string, ctx *Context) {

		ada := new(ConsoleAdapter)
		facade := new(innerConsoleContextFacade)

		facade.ctx2 = ctx
		facade.head = cc

		ctx.Adapter = ada
		ctx.Facade = facade

		facade.putValue(name, ctx)
	})
	if err != nil {
		panic(err)
	}

	return c2
}

////////////////////////////////////////////////////////////////////////////////
