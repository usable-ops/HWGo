package main

import "fmt"

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: Code: %d, Message: %s", e.Code, e.Message)
}

func main() {
	errChan := make(chan error, 1)

	myErr := &MyError{
		Message: "Low pressure",
		Code:    5432,
	}

	errChan <- myErr
	close(errChan)
	receivedErr := <-errChan

	if receivedErr != nil {
		fmt.Println("Received error:", receivedErr)
	} else {
		fmt.Println("No error received from the channel.")
	}
}
