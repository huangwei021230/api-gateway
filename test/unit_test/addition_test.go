package unit_test

import (
	addition "github.com/huangwei021230/api-gateway/microservices/addition-service/kitex_gen/addition/management"
	addImpl "github.com/huangwei021230/api-gateway/test/unit_test/utils"
	"golang.org/x/net/context"
	"strconv"
	"testing"
)

func Test_addition1(t *testing.T) {
	additionImpl := new(addImpl.AdditionManagementImpl)
	req := &addition.AdditionRequest{
		FirstNum:  "100",
		SecondNum: "20",
	}
	resp, err := additionImpl.AddNumbers(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	expectResult := strconv.Itoa(100 + 20)
	if resp.Sum != expectResult {
		t.Errorf("expect result is %v, but get %v", expectResult, resp.String())
	}
}

func Test_addition2(t *testing.T) {
	additionImpl := new(addImpl.AdditionManagementImpl)
	req := &addition.AdditionRequest{
		FirstNum:  "10",
		SecondNum: "-10",
	}
	resp, err := additionImpl.AddNumbers(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	expectResult := strconv.Itoa(10 + (-10))
	if resp.Sum != expectResult {
		t.Errorf("expect result is %v, but get %v", expectResult, resp.String())
	}
}
