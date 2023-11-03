package main

import "testing"

func TestFoo(t *testing.T) {
	c := foo(1, 2, 3, 4)

	if c != 6 {
		t.Error("test failed c is", c)
	}
}
