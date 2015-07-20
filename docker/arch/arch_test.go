package arch

import (
	"testing"
)

func TestDebian(t *testing.T) {

	k, _ := Debian()

	if k != true {
		t.Error("This  functionality works on debian flavours")
	}
}
