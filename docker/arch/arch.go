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

<<<<<<< HEAD
func Redhat() (value bool, err error) {

	if _, err := os.Stat("/etc/redhat-release"); os.IsNotExist(err) {
		fmt.Println(err.Error())
		return false, errors.New("No Such file : /etc/redhat-release")
	}

	return true, nil

}
=======
/*func Redhat() (value bool, err error) {

}*/
>>>>>>> 891b07febf61a7f3fab9c6aeb92d6e4584ceff89
