package main

import (
	"errors"
	"fmt"
	"lession2/week2/code"

	code_errors "github.com/pkg/errors"
)

func main() {
	if err := getUser(); err != nil {

		fmt.Printf("%#-v\n\n", err)

	}
	fmt.Println("service ok")
}

func getUser() error {
	err := queryDatabase()
	// Step2: Wrap the error with a new error message.
	if errors.Is(err, code.NoExists) {
		return err

	}

	return nil
}

func queryDatabase() error {
	// Step1. Create error with specified error code.
	sql := "SELECT userName,userCode FROM user WHERE id = 1999"
	return code_errors.Wrapf(code.NoExists, fmt.Sprintf("sql: %s err: %v"), sql, errors.New("test"))

}
