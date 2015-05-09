package crypt

import (
	"testing"
)

func TestExchange(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := Randomkey()
		b := Randomkey()
		A := DHExchange(a)
		B := DHExchange(b)
		secretA := DHSecret(a, B)
		secretB := DHSecret(b, A)
		if secretA.Cmp(secretB) != 0 {
			t.Errorf("test[%d] failed\n", i)
		}
	}
}
