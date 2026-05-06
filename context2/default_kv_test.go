package context2

import (
	"context"
	"sort"
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
	kv.SetValue("zzz", new(CommonLoader))

	// kv.GetValue("x")

	keys := kv.Keys()
	keys = append(keys, "x")
	sort.Strings(keys)

	for idx, key := range keys {
		value := kv.GetValue(key)
		t.Logf("  Context.value[%d]: %s = %s \n", idx, key, value)
	}

}
