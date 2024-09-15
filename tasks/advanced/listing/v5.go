package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test2() *customError {
	{

	}

	return nil
}

func main() {
	var err error
	err = test2()
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println("ok")
}
