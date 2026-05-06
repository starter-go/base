package context2

import (
	"context"
	"fmt"
)

type Context struct {
	Adapter Adapter

	Holder Holder

	Loader Loader

	Facade context.Context

	Raw context.Context

	Getter Getter

	Setter Setter

	Values Values

	Loaded bool

	Keys map[string]bool
}

func (inst *Context) GetAdapter() (Adapter, error) {

	if inst == nil {
		return nil, fmt.Errorf("context2.Context.GetAdapter() : self (Context) is nil")
	}

	ada := inst.Adapter
	if ada == nil {
		return nil, fmt.Errorf("context2.Context.GetAdapter() : adapter is nil")
	}

	return ada, nil
}

func (inst *Context) GetHolder() (Holder, error) {

	holder := inst.Holder
	if holder == nil {
		return nil, fmt.Errorf("context2.Context.GetHolder() : holder is nil")
	}
	return holder, nil
}

func (inst *Context) GetLoader() (Loader, error) {

	loader := inst.Loader
	if loader == nil {
		return nil, fmt.Errorf("context2.Context.GetLoader() : loader is nil")
	}
	return loader, nil
}

func (inst *Context) Load() error {

	if inst.Loaded {
		return nil
	}

	// do load

	loader, err := inst.GetLoader()
	if err != nil {
		return err
	}

	err = loader.Load(inst)
	if err != nil {
		return err
	}

	inst.Loaded = true
	return nil
}

func (inst *Context) GetFacade() (context.Context, error) {

	err := inst.Load()
	if err != nil {
		return nil, err
	}

	facade := inst.Facade
	if facade == nil {
		return nil, fmt.Errorf("context2.Context.GetFacade() : facade is nil")
	}

	return facade, nil
}

func (inst *Context) GetValues() (Values, error) {

	err := inst.Load()
	if err != nil {
		return nil, err
	}

	values := inst.Values
	if values == nil {
		return nil, fmt.Errorf("context2.Context.GetValues() : values is nil")
	}

	return values, nil
}

func (inst *Context) PutKey(key any) {
	if key == nil {
		return
	}
	table := inst.Keys
	if table == nil {
		table = make(map[string]bool)
		inst.Keys = table
	}
	keyStr, ok := key.(string)
	if !ok {
		keyStr = fmt.Sprint(key)
	}
	table[keyStr] = true
}
