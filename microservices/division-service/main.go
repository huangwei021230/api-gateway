package main

import (
	api "github.com/huangwei021230/api-gateway/microservices/division-service/kitex_gen/division/api/divisionmanagement"
	"log"
)

func main() {
	svr := api.NewServer(new(DivisionManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
