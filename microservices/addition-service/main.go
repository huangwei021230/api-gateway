package main

import (
	management "github.com/huangwei021230/api-gateway/microservices/addition-service/kitex_gen/addition/management/additionmanagement"
	"log"
)

func main() {
	svr := management.NewServer(new(AdditionManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
