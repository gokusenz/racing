package main

import "testing"

func TestGet(t *testing.T) {
	i := getNumber(5)
	if i != 5 {
		t.Error("it should be 5 but got", i)
	}
}
