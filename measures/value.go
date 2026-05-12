package measures

import "fmt"

type Value struct {
	Number float64 // 数值
	Unit   Unit    // 单位
}

func (v Value) String() string {
	n := v.Number
	u := v.Unit
	return fmt.Sprintf("%f(%s)", n, u)
}
