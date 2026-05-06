package context2

import (
	"context"
	"testing"
)

func TestValues(t *testing.T) {

	// setup

	cc := context.Background()
	cc = SetupConsoleContext(cc)

	// use (get)

	kv, err := GetValues(cc)
	if err != nil {
		t.Error(err)
	}

	kv.SetValue("a", "1")
	kv.SetValue("b", "2")
	kv.SetValue("c", "3")

	kv.GetValue("x")
}
