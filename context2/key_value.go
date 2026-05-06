package context2

import (
	"context"
)

// const theValuesKey = "base/context2.Values#binding"

// Getter 是作用于 context.Context 的 key-value 取值接口
type Getter interface {
	GetValue(key any) any
}

// Setter 是作用于 context.Context 的 key-value 设置接口
type Setter interface {
	SetValue(key, value any)
}

// Values 是作用于 context.Context 的 key-value 存取接口
type Values interface {
	Getter
	Setter
	Context() context.Context
}

// GetValues 从上下文中获取 Values 接口
func GetValues(c context.Context) (Values, error) {

	ctx, err := GetContext(c)
	if err != nil {
		return nil, err
	}
	return ctx.GetValues()
}

////////////////////////////////////////////////////////////////////////////////

type CommonValues struct {
	context *Context
}

// Context implements Values.
func (inst *CommonValues) Context() context.Context {

	facade, err := inst.context.GetFacade()
	if err != nil {
		panic(err)
	}
	return facade
}

// GetValue implements Values.
func (inst *CommonValues) GetValue(key any) any {

	ctx := inst.context
	ada, err := ctx.GetAdapter()
	if err != nil {
		panic(err)
	}
	return ada.GetValue(ctx, key)
}

// SetValue implements Values.
func (inst *CommonValues) SetValue(key any, value any) {

	ctx := inst.context
	ada, err := ctx.GetAdapter()
	if err != nil {
		panic(err)
	}
	ada.SetValue(ctx, key, value)
}

func (inst *CommonValues) _impl() Values {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
