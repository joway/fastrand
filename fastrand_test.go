package fastrand

import "testing"

func TestFastRand(t *testing.T) {
	rd := New()
	for i := 0; i < 100; i++ {
		if rd.Intn(1) != 0 {
			t.Fail()
		}
	}
}
