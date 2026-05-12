package measures

import "testing"

func TestLength(t *testing.T) {

	min := float64(0.00000000001)
	max := float64(10000000000)

	for n := min; n < max; n = n * 2 {

		value := new(Value)
		value.Unit = UnitM
		value.Number = n
		t.Logf("value = %s", value.String())

	}

}
