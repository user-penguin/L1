package p8

import "testing"

func Test_pow2_0(t *testing.T) {
	if pow2(0) != 1 {
		t.Error("pow2(0) != 1")
	}
}
func Test_pow2_1(t *testing.T) {
	if pow2(1) != 2 {
		t.Error("pow2(1) != 2")
	}
}

func Test_pow2_10(t *testing.T) {
	if pow2(10) != 1024 {
		t.Error("pow2(1) != 2")
	}
}
