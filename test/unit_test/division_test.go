package unit_test

import (
	divition "github.com/huangwei021230/api-gateway/microservices/division-service/kitex_gen/division/management"
	divImpl "github.com/huangwei021230/api-gateway/test/unit_test/utils"
	"golang.org/x/net/context"
	"strconv"
	"testing"
)

func Test_division1(t *testing.T) {
	divisionImpl := new(divImpl.DivisionManagementImpl)
	req := &divition.DivisionRequest{
		FirstNum:  "100",
		SecondNum: "20",
	}
	resp, err := divisionImpl.DivideNumbers(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	expectResult := 100.00 / 20
	result, _ := strconv.ParseFloat(resp.Quotient, 64)
	if result != expectResult {
		t.Errorf("expect result is %v, but get %v", expectResult, resp.String())
	}
}

func Test_division2(t *testing.T) {
	divisionImpl := new(divImpl.DivisionManagementImpl)
	req := &divition.DivisionRequest{
		FirstNum:  "0",
		SecondNum: "20",
	}
	resp, err := divisionImpl.DivideNumbers(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	expectResult := 0.00 / 20
	result, _ := strconv.ParseFloat(resp.Quotient, 64)
	if result != expectResult {
		t.Errorf("expect result is %v, but get %v", expectResult, resp.String())
	}
}
