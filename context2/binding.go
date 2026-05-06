package context2

import (
	"context"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////

const theBindingName = "github.com/starter-go/base/context2/Context#binding"

type OnSetup func(name string, value *Context)

////////////////////////////////////////////////////////////////////////////////

func GetHolder(cc context.Context) (Holder, error) {
	ctx, err := innerTryGetContext(cc)
	if err != nil {
		return nil, err
	}
	return ctx.GetHolder()
}

func GetContext(cc context.Context) (*Context, error) {
	return innerTryGetContext(cc)
}

func Setup(cc context.Context, callback OnSetup) (context.Context, error) {

	// try get
	ctx, err := innerTryGetContext(cc)
	if err == nil && ctx != nil {
		return ctx.GetFacade()
	}

	// do setup
	const name = theBindingName
	ctx = new(Context)
	ctx.Facade = cc
	ctx.Raw = cc
	ctx.Loader = new(CommonLoader)

	callback(name, ctx)

	cc = ctx.Facade

	if ctx.Adapter == nil {
		return nil, fmt.Errorf("context2.Setup() : adapter is nil")
	}
	ctx, err = innerTryGetContext(cc)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		return nil, fmt.Errorf("context2.Setup() : context is nil")
	}
	return ctx.GetFacade()
}

func Bind(cc context.Context, callback OnSetup) (context.Context, error) {
	return Setup(cc, callback)
}

func innerTryGetContext(cc context.Context) (*Context, error) {

	const name = theBindingName
	value := cc.Value(name)
	ctx, ok := value.(*Context)

	if (!ok) || (ctx == nil) {
		return nil, fmt.Errorf("the context-value is not setup, name = '%s'", name)
	}

	return ctx, nil
}

////////////////////////////////////////////////////////////////////////////////
