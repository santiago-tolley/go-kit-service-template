package service

import (
	"context"
	"fmt"
)

func NewServer() Service {
	return Service{}
}

type Service struct{}

func (p Service) DoStuff(ctx context.Context, value int) (int, error) {
	// Do Stuff: call a client service, check and log errors, etc
	fmt.Printf("Hello! the value recieved was %v\n", value)
	result := value * 2
	fmt.Printf("The result of the calculations is %v\n", result)
	return result, nil
}
