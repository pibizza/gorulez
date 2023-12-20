package main

import "testing"

func TestTypeMatch(t *testing.T) {
	intcstrn := constr.istype[int]{}

	if !intcstrn.Match(3) {
		t.Error("intcstrn.Match(3) == false")
	}

	if intcstrn.Match(3.0) {
		t.Error("intcstrn.Match(3.0) == false")
	}

}
