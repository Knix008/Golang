package fibonacci

import "testing"

func TestFibonacci_1(t *testing.T) {
	if fibo(0) != 0 {
		t.Error(`fibo(0) ! = 0`)
	}
	if fibo(1) != 1 {
		t.Error(`fibo(1) ! = 1`)
	}
	if fibo(2) != 1 {
		t.Error(`fibo(2) ! = 1`)
	}
	if fibo(3) != 2 {
		t.Error(`fibo(3) ! = 2`)
	}
}

func TestFibonacci_2(t *testing.T) {
	if fibo(4) != 3 {
		t.Error(`fibo(4 ! = 3)`)
	}
	if fibo(5) != 4 {
		// Error
		t.Error(`fibo(=5 ! = 4)`)
	}
}
