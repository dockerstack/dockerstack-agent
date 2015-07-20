package arch

import (
	"testing"
)

func TestDebian(t *testing.T) {

	k, _ := Debian()

	if k != true {
		t.Error("This  Functionality works on Debian Flavours")
	}
}

func TestRedhat(t *testing.T) {

	k, _ := Redhat()

	if k != true {
		t.Error("this Functionality works on Redhat Flavours")
	}
}
