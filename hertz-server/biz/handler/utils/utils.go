package utils

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/generic"
	manager "github.com/huangwei021230/api-gateway/hertz-server/biz/handler/idl"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

func NewResolver() discovery.Resolver {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		log.Fatal("Error: fail to new etcd resolver---" + err.Error())
	}
	return r
}

func NewProvider(serviceName string) *generic.ThriftContentProvider {

	path := manager.GetIdlPath(serviceName)
	content, _ := manager.GetIdlContent(serviceName)
	includes := map[string]string{
		path: content,
	}
	p, err := generic.NewThriftContentProvider(content, includes)

	if err != nil {
		panic("Error: fail to load the thrift file " + err.Error())
	} else {
		go func() {
			err = p.UpdateIDL(content, includes)
			if err != nil {
				panic(err)
			}
			time.Sleep(30 * time.Second)
		}()
	}
	return p
}

func NewClient(destServiceName string, provider *generic.ThriftContentProvider, resolver discovery.Resolver) genericclient.Client {
	g, err := generic.HTTPThriftGeneric(provider)
	if err != nil {
		panic("Error: fail to generic thrift " + err.Error())
	}
	var opts []client.Option
	opts = append(opts, client.WithResolver(resolver))
	opts = append(opts, client.WithLongConnection(connpool.IdleConfig{
		MaxIdlePerAddress: 1000,
		MaxIdleGlobal:     1000 * 10,
	}))
	opts = append(opts, client.WithTag("Cluster", destServiceName+"Cluster"))

	cli, err := genericclient.NewClient(destServiceName, g, opts...)
	if err != nil {
		panic("Error: fail to establish the client " + err.Error())
	}
	return cli
}

func GetHTTPGenericResponse(ctx context.Context, c *app.RequestContext, methodName string, cli genericclient.Client) *generic.HTTPResponse {
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic("Error: fail to generic from hertz request " + err.Error())
	}
	// customReq *generic.HttpRequest
	// 由于 hertz 泛化的 method 是通过 bam 规则从 hertz request 中获取的，所以填空就行
	fmt.Println(string(c.GetRequest().Body()))
	fmt.Println(cli)
	resp, err := cli.GenericCall(ctx, methodName, customReq)
	if err != nil {
		panic("Error: fail to generic call " + err.Error())
	}

	realResp := resp.(*generic.HTTPResponse)
	return realResp
}
