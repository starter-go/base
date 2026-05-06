package context2

import "fmt"

type Loader interface {
	Load(c *Context) error
}

////////////////////////////////////////////////////////////////////////////////

type CommonLoader struct {
}

// Load implements Loader.
func (inst *CommonLoader) Load(c *Context) error {

	ada := c.Adapter
	loader := c.Loader
	cc1 := c.Facade
	cc2 := c.Raw

	if ada == nil {
		return fmt.Errorf("context2.CommonLoader : adapter is nil")
	}
	if loader == nil {
		return fmt.Errorf("context2.CommonLoader : loader is nil")
	}
	if cc1 == nil {
		return fmt.Errorf("context2.CommonLoader : facade (Context) is nil")
	}
	if cc2 == nil {
		return fmt.Errorf("context2.CommonLoader : raw (Context) is nil")
	}

	holder := c.Holder
	getter := c.Getter
	setter := c.Setter
	values := c.Values

	if holder == nil {
		holder = &CommonHolder{
			context: c,
		}
		c.Holder = holder
	}

	if values == nil {
		values = &CommonValues{
			context: c,
		}
		c.Values = values
	}

	if getter == nil {
		getter = values
		c.Getter = getter
	}

	if setter == nil {
		setter = values
		c.Setter = setter
	}

	c.Loaded = true
	return nil
}

func (inst *CommonLoader) _impl() Loader {
	return inst
}
