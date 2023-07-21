package main

import (
	"context"
	"fmt"
	"strconv"

	management "github.com/huangwei021230/api-gateway/microservices/multiplication-service/kitex_gen/multiplication/management"
)

// MultiplicationManagementImpl implements the last service interface defined in the IDL.
type MultiplicationManagementImpl struct{}

// MultiplyNumbers implements the MultiplicationManagementImpl interface.
func (s *MultiplicationManagementImpl) MultiplyNumbers(ctx context.Context, req *management.MultiplicationRequest) (resp *management.MultiplicationResponse, err error) {
	// TODO: Your code here...
	firstNumInt, err := strconv.Atoi(req.FirstNum)
	if err != nil {
		panic(err)
	}

	secondNumInt, err := strconv.Atoi(req.SecondNum)
	if err != nil {
		panic(err)
	}

	// add two numbers together
	finalSum := firstNumInt / secondNumInt

	return &management.MultiplicationResponse{
		Product: fmt.Sprintf("%d", finalSum),
	}, nil

}
