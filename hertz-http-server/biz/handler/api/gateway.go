// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	addManagement "github.com/huangwei021230/api-gateway/hertz-http-server/biz/model/addition/management"
	divManagement "github.com/huangwei021230/api-gateway/hertz-http-server/biz/model/division/management"
	mulManagement "github.com/huangwei021230/api-gateway/hertz-http-server/biz/model/multiplication/management"
	additionService "github.com/huangwei021230/api-gateway/microservices/addition-service/kitex_gen/addition/management"
	divisionService "github.com/huangwei021230/api-gateway/microservices/division-service/kitex_gen/division/management"
	multiplicationService "github.com/huangwei021230/api-gateway/microservices/multiplication-service/kitex_gen/multiplication/management"
	"github.com/huangwei021230/api-gateway/utils"
)

// AddNumbers .
// @router /add [POST]
func AddNumbers(ctx context.Context, c *app.RequestContext) {

	// inital declarations (pre-generated)
	var err error
	var req addManagement.AdditionRequest

	// bind error params to req (pre-generated)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// create new client (with loadbalancing, service discovery capabilities) using utils.GenerateClient feature
	additionClient, err := utils.GenerateClient("Addition")
	if err != nil {
		panic(err)
	}

	// binding req params to RPC reqest struct (following the request format declared in RPC service IDL)
	reqRpc := &additionService.AdditionRequest{
		FirstNum:  req.FirstNum,
		SecondNum: req.SecondNum,
	}

	// initate new RPC response struct (as declared in RPC service IDL). This response variable will be populated by MakeRpcRequst function
	var respRpc additionService.AdditionResponse

	// calling MakeRpcRequest method declared in the utils package
	err = utils.MakeRpcRequest(ctx, additionClient, "addNumbers", reqRpc, &respRpc)
	if err != nil {
		panic(err)
	}

	// initating and repackaging RPC response into new HTTP AdditionResponse
	resp := &addManagement.AdditionResponse{
		Sum: respRpc.Sum,
	}

	// return to client as JSON HTTP response
	c.JSON(consts.StatusOK, resp)
}

// MultiplyNumbers .
// @router /multiply [POST]
func MultiplyNumbers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req mulManagement.MultiplicationRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// create new client (with loadbalancing, service discovery capabilities) using utils.GenerateClient feature
	multiplicationClient, err := utils.GenerateClient("Multiplication")
	if err != nil {
		panic(err)
	}

	// binding req params to RPC reqest struct (following the request format declared in RPC service IDL)
	reqRpc := &multiplicationService.MultiplicationRequest{
		FirstNum:  req.FirstNum,
		SecondNum: req.SecondNum,
	}

	// initate new RPC response struct (as declared in RPC service IDL). This response variable will be populated by MakeRpcRequst function
	var respRpc multiplicationService.MultiplicationResponse

	// calling MakeRpcRequest method declared in the utils package
	err = utils.MakeRpcRequest(ctx, multiplicationClient, "multiplyNumbers", reqRpc, &respRpc)
	if err != nil {
		panic(err)
	}

	// initating and repackaging RPC response into new HTTP MultiplicationResponse
	resp := &mulManagement.MultiplicationResponse{
		Product: respRpc.Product,
	}

	// return to client as JSON HTTP response
	c.JSON(consts.StatusOK, resp)
}

// DivideNumbers .
// @router /divide [POST]
func DivideNumbers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req divManagement.DivisionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// create new client (with loadbalancing, service discovery capabilities) using utils.GenerateClient feature
	divisionClient, err := utils.GenerateClient("Division")
	if err != nil {
		panic(err)
	}

	// binding req params to RPC reqest struct (following the request format declared in RPC service IDL)
	reqRpc := &divisionService.DivisionRequest{
		FirstNum:  req.FirstNum,
		SecondNum: req.SecondNum,
	}

	// initate new RPC response struct (as declared in RPC service IDL). This response variable will be populated by MakeRpcRequst function
	var respRpc divisionService.DivisionResponse

	// calling MakeRpcRequest method declared in the utils package
	err = utils.MakeRpcRequest(ctx, divisionClient, "divideNumbers", reqRpc, &respRpc)
	if err != nil {
		panic(err)
	}

	// initating and repackaging RPC response into new HTTP DivisionResponse
	resp := &divManagement.DivisionResponse{
		Quotient: respRpc.Quotient,
	}

	// return to client as JSON HTTP response
	c.JSON(consts.StatusOK, resp)
}
