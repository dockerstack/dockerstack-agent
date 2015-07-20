package arch

import (
	"errors"
	"fmt"
	"os"
)

func Debian() (value bool, err error) {

	if _, err := os.Stat("/etc/lsb-release"); os.IsNotExist(err) {
		fmt.Println(err.Error())
		return false, errors.New("No such file  : /etc/lsb-release")
	}

	return true, nil

}

/*func Redhat() (value bool, err error) {

}*/
