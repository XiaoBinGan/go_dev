package main

import (
	"testing"
)


func TestAdd(t *testing.T)  {
	r :=add(3,3)
	if r !=6 {
		t.Fatal("add(2,4) error ,exprct")
	}
	t.Logf("test add success")
}