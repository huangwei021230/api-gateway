package addition_service

import (
	"context"
	management "github.com/huangwei021230/api-gateway/microservices/addition-service/kitex_gen/addition/management"
)

// AdditionManagementImpl implements the last service interface defined in the IDL.
type AdditionManagementImpl struct{}

// AddNumbers implements the AdditionManagementImpl interface.
func (s *AdditionManagementImpl) AddNumbers(ctx context.Context, req *management.AdditionRequest) (resp *management.AdditionResponse, err error) {
	// TODO: Your code here...
	return
}
