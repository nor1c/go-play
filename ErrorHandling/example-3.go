package main

import (
	"errors"
	"fmt"
)

type CustomErr struct {
	Code    string
	Message string
}

func (ec CustomErr) Error() string {
	return ec.Message
}

func NewCustomErr(code string, message string) *CustomErr {
	return &CustomErr{
		Code:    code,
		Message: message,
	}
}

func doSomething(level int) error {
	if level == 1 {
		return NewCustomErr("ERR01", "error level 1")
	}

	if level == 2 {
		return errors.New("regular error")
	}

	return nil
}

func main() {
	err := doSomething(1)
	if err != nil {
		var myCustomErr *CustomErr

		if errors.As(err, &myCustomErr) {
			fmt.Println("query error", myCustomErr.Message)
		} else {
			fmt.Println(err)
		}
	}
}
