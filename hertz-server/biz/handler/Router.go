package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	manager "github.com/huangwei021230/api-gateway/hertz-server/biz/handler/idl"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func FindServiceAndCall(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interface{}
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	serviceName := c.Param("service")
	methodName := c.Param("method")

	//将请求req参数转换为 JSON
	jsonReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("jsonReq:", string(jsonReq))

	p, err := generic.NewThriftFileProvider(manager.GetIdlPath(serviceName))
	if err != nil {
		panic(err)
	}

	// 构造 JSON 请求和返回类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	// 在 etcd 中服务发现
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	cli, err := genericclient.NewClient(serviceName, g, client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()))
	if err != nil {
		panic(err)
	}

	// 泛化调用
	resp, err := cli.GenericCall(ctx, methodName, string(jsonReq))

	c.JSON(consts.StatusOK, resp)
}
