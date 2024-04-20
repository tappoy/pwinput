package pwinput

import (
	"testing"
)

func TestDummyPasswordInput(t *testing.T) {
	pwi := NewDummyPasswordInput("dummy")
	if password, err := pwi.InputPassword(); err != nil {
		t.Error(err)
	} else if password != "dummy" {
		t.Error("Password is not correct")
	}
}
