package utils

import (
	"context"
	"encoding/json"

	kitexClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func GenerateClient(serviceName string) (genericclient.Client, error) {

	// inital declarations
	var err error

	// initating loadbalancer
	lb := loadbalance.NewWeightedBalancer()

	// initating etcs resolver (for service discovery)
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	// importing idl for reference(generic call)
	// 本地文件 idl 解析
	// YOUR_IDL_PATH thrift 文件路径: 举例 ./idl/example.thrift
	// includeDirs: 指定 include 路径，默认用当前文件的相对路径寻找 include
	p, err := generic.NewThriftFileProvider("../thrift-idl/gateway_api.thrift")
	if err != nil {
		panic(err)
	}

	// convert to thrift generic form
	// 构造 JSON 请求和返回类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	// create new generic client
	client, err := genericclient.NewClient(
		serviceName,
		g,
		kitexClient.WithResolver(r),
		kitexClient.WithLoadBalancer(lb),
	)
	if err != nil {
		panic(err)
	}

	return client, nil
}

func jsonStringify(item any) (string, error) {
	// convert to request struct to JSON format (so it can be converted to json string)
	jsonForm, err := json.Marshal(&item)
	if err != nil {
		panic(err)
	}

	return string(jsonForm), nil
}

func MakeRpcRequest(ctx context.Context, kitexClient genericclient.Client, methodName string, request interface{}, response interface{}) error {
	stringedReq, err := jsonStringify(request)
	if err != nil {
		panic(err)
	}

	// making generic call to addNumbers method of client
	respRpc, err := kitexClient.GenericCall(ctx, methodName, stringedReq)
	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(respRpc.(string)), response)

	return nil
}

func RunAllServices() {

}
