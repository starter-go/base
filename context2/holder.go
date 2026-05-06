package context2

import "context"

type Holder interface {
	Context() context.Context

	Setter() Setter

	Getter() Getter
}

////////////////////////////////////////////////////////////////////////////////

type CommonHolder struct {
	context *Context
}

// Context implements Holder.
func (inst *CommonHolder) Context() context.Context {
	panic("unimplemented")
}

// Getter implements Holder.
func (inst *CommonHolder) Getter() Getter {
	panic("unimplemented")
}

// Setter implements Holder.
func (inst *CommonHolder) Setter() Setter {
	panic("unimplemented")
}

func (inst *CommonHolder) _impl() Holder {
	return inst
}
