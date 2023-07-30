package main

import (
	"context"
	"fmt"
	"strconv"

	management "github.com/huangwei021230/api-gateway/microservices/addition-service/kitex_gen/addition/management"
)

// AdditionManagementImpl implements the last service interface defined in the IDL.
type AdditionManagementImpl struct{}

// AddNumbers implements the AdditionManagementImpl interface.
func (s *AdditionManagementImpl) AddNumbers(ctx context.Context, req *management.AdditionRequest) (resp *management.AdditionResponse, err error) {

	// parse int from string of First Number
	firstNumInt, err := strconv.Atoi(req.FirstNum[2 : len(req.FirstNum)-2])
	if err != nil {
		panic(err)
	}

	// parse int from string of Second Number
	secondNumInt, err := strconv.Atoi(req.SecondNum[2 : len(req.SecondNum)-2])
	if err != nil {
		panic(err)
	}

	// add two numbers together
	finalSum := firstNumInt + secondNumInt

	// convert finalSum to string and return response of type AdditionResponse and error = nil
	return &management.AdditionResponse{
		Sum: fmt.Sprintf("%d", finalSum),
	}, nil

}
