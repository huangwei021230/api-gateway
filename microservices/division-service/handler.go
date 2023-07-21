package main

import (
	"context"
	"fmt"
	"strconv"

	api "github.com/huangwei021230/api-gateway/microservices/division-service/kitex_gen/division/api"
)

// DivisionManagementImpl implements the last service interface defined in the IDL.
type DivisionManagementImpl struct{}

// DivideNumbers implements the DivisionManagementImpl interface.
func (s *DivisionManagementImpl) DivideNumbers(ctx context.Context, req *api.DivisionRequest) (resp *api.DivisionResponse, err error) {
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

	return &api.DivisionResponse{
		Quotient: fmt.Sprintf("%d", finalSum),
	}, nil

}
