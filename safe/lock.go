package safe

import "sync"

// Lock 代表一个同步锁
type Lock interface {
	Lock()
	Unlock()
}

////////////////////////////////////////////////////////////////////////////////

type lockFast struct{}

func (inst *lockFast) Lock() {}

func (inst *lockFast) Unlock() {}

////////////////////////////////////////////////////////////////////////////////

type lockSafe struct {
	m sync.Mutex
}

func (inst *lockSafe) Lock() {
	inst.m.Lock()
}

func (inst *lockSafe) Unlock() {
	inst.m.Unlock()
}

////////////////////////////////////////////////////////////////////////////////