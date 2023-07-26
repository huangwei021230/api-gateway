package main

import (
	"context"
	"fmt"
	api "github.com/huangwei021230/api-gateway/microservices/division-service/kitex_gen/division/management"
	"strconv"
)

// DivisionManagementImpl implements the last service interface defined in the IDL.
type DivisionManagementImpl struct{}

// DivideNumbers implements the DivisionManagementImpl interface.
func (s *DivisionManagementImpl) DivideNumbers(ctx context.Context, req *api.DivisionRequest) (resp *api.DivisionResponse, err error) {
	firstNumFloat, err := strconv.ParseFloat(req.FirstNum[2:len(req.FirstNum)-2], 64)
	if err != nil {
		return nil, err
	}

	secondNumFloat, err := strconv.ParseFloat(req.SecondNum[2:len(req.SecondNum)-2], 64)
	if err != nil {
		return nil, err
	}

	finalQuotient := firstNumFloat / secondNumFloat

	return &api.DivisionResponse{
		Quotient: fmt.Sprintf("%.2f", finalQuotient),
	}, nil
}
