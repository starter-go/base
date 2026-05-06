package util

import (
	"math/rand"
	"testing"

	"github.com/starter-go/base"
)

func TestSorter(t *testing.T) {

	list := make([]int, 0)
	for i := 10; i > 0; i-- {
		n := rand.Int()
		list = append(list, n&0xff)
	}

	t.Log("before:")
	for i, n := range list {
		t.Logf("  int.list[%d]=%d", i, n)
	}

	sorter := &Sorter{
		OnLen:  func() int { return len(list) },
		OnLess: func(i1, i2 int) bool { return list[i1] < list[i2] },
		OnSwap: func(i1, i2 int) { list[i1], list[i2] = list[i2], list[i1] },
	}
	sorter.Sort()

	t.Log("after:")
	for i, n := range list {
		t.Logf("  int.list[%d]=%d", i, n)
	}
}

func TestThisModuleInfo(t *testing.T) {

	info := base.GetModuleInfoText()
	t.Logf("module.info =  %s", info)

}
