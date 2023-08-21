/*
*
Handle error with Sentinel Error
*/
package main

import (
	"errors"
	"fmt"
)

var ErrLevel1 = errors.New("error level 1")
var ErrLevel2 = errors.New("error level 2")
var ErrLevel3 = errors.New("error level 3")

func giveMeErr(level int) error {
	if level == 1 {
		return ErrLevel1
	}

	if level == 2 {
		return ErrLevel2
	}

	if level == 3 {
		return ErrLevel3
	}

	return nil
}

func main() {
	err := giveMeErr(2)
	if err != nil {
		switch {
		case err == ErrLevel1:
			fmt.Println("handle error 1")
		case errors.Is(err, ErrLevel2):
			fmt.Println("handle error 2")
		case err.Error() == "error level 3": // not recommended using string
			fmt.Println("handle error 3")
		default:
			fmt.Println("unknown error")
		}
	}
}
