package fibonacci

import (
	"testing"
)

func TestFibonacci_1(t *testing.T) {
	if fibo(0) != 0 {
		t.Error("fibo(0) is not 0")
	}
	if fibo(1) != 1 {
		t.Error("fibo(1) is not 1")
	}
	if fibo(2) != 1 {
		t.Error("fibo(2) is not 1")
	}
	if fibo(3) != 2 {
		t.Error("fibo(3) is not 2")
	}
}

func TestFibonacci_2(t *testing.T) {
	if fibo(10) != 55 {
		t.Error("fibo(10) is not 55")
	}
}
