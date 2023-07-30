package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	management "github.com/huangwei021230/api-gateway/microservices/multiplication-service/kitex_gen/multiplication/management/multiplicationmanagement"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	// initate new etcd registry at port 2379
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := management.NewServer(
		new(MultiplicationManagementImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "multiplication-server",
		}),
		server.WithRegistry(r),
		server.WithServiceAddr(&net.TCPAddr{Port: 8890}),
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
