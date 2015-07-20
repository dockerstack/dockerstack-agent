package arch

import (
	"testing"
)

func TestDebian(t *testing.T) {

	k, _ := Debian()

	if k != true {
<<<<<<< HEAD
		t.Error("This  Functionality works on Debian Flavours")
	}
}

func TestRedhat(t *testing.T) {

	k, _ := Redhat()

	if k != true {
		t.Error("this Functionality works on Redhat Flavours")
=======
		t.Error("This  functionality works on debian flavours")
>>>>>>> 891b07febf61a7f3fab9c6aeb92d6e4584ceff89
	}
}
