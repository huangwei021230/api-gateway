package main

import (
	management "github.com/huangwei021230/api-gateway/microservices/multiplication-service/kitex_gen/multiplication/management/multiplicationmanagement"
	"log"
)

func main() {
	svr := management.NewServer(new(MultiplicationManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
