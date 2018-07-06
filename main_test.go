package main

import "testing"

func TestGet(t *testing.T) {
	i := getNumber(5)
	if i != 0 {
		t.Error("it should be 0 but got", i)
	}
}
