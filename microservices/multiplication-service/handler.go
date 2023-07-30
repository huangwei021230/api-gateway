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
	firstNum, err := strconv.Atoi(req.FirstNum[2 : len(req.FirstNum)-2])
	if err != nil {
		return nil, fmt.Errorf("invalid first number: %v", err)
	}

	secondNum, err := strconv.Atoi(req.SecondNum[2 : len(req.SecondNum)-2])
	if err != nil {
		return nil, fmt.Errorf("invalid second number: %v", err)
	}

	product := firstNum * secondNum

	return &management.MultiplicationResponse{
		Product: fmt.Sprintf("%d", product),
	}, nil

}
