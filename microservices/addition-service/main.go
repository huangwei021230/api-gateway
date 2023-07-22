package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	management "github.com/huangwei021230/api-gateway/microservices/addition-service/kitex_gen/addition/management/additionmanagement"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	// initate new etcd registry at port 2379
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	// create new Kitex server for Addition Service
	svr := management.NewServer(
		new(AdditionManagementImpl), // Follow AdditionManagementImpl as defined in ./handler.go
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Addition"}), // allow service to be discovered with name: "Addition"
		server.WithRegistry(r), // register service on etcd registry 'r' (as declared earlier)
		server.WithServiceAddr(&net.TCPAddr{Port: 8892}),
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
