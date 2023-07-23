package unit_test

import (
	multiplication "github.com/huangwei021230/api-gateway/microservices/multiplication-service/kitex_gen/multiplication/management"
	mulImpl "github.com/huangwei021230/api-gateway/test/unit_test/utils"
	"golang.org/x/net/context"

	"strconv"
	"testing"
)

func Test_multiplication1(t *testing.T) {
	multiplicationImpl := new(mulImpl.MultiplicationManagementImpl)
	req := &multiplication.MultiplicationRequest{
		FirstNum:  "100",
		SecondNum: "20",
	}
	resp, err := multiplicationImpl.MultiplyNumbers(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	expectResult := strconv.Itoa(100 * 20)
	if resp.Product != expectResult {
		t.Errorf("expect result is %v, but get %v", expectResult, resp.String())
	}
}

func Test_multiplication2(t *testing.T) {
	multiplicationImpl := new(mulImpl.MultiplicationManagementImpl)
	req := &multiplication.MultiplicationRequest{
		FirstNum:  "0",
		SecondNum: "20",
	}
	resp, err := multiplicationImpl.MultiplyNumbers(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	expectResult := strconv.Itoa(0 * 20)
	if resp.Product != expectResult {
		t.Errorf("expect result is %v, but get %v", expectResult, resp.String())
	}
}
