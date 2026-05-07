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

	f, err := inst.context.GetFacade()
	if err != nil {
		panic(err)
	}
	return f
}

// Getter implements Holder.
func (inst *CommonHolder) Getter() Getter {

	v, err := inst.context.GetValues()
	if err != nil {
		panic(err)
	}
	return v
}

// Setter implements Holder.
func (inst *CommonHolder) Setter() Setter {

	v, err := inst.context.GetValues()
	if err != nil {
		panic(err)
	}
	return v
}

func (inst *CommonHolder) _impl() Holder {
	return inst
}
