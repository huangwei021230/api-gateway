namespace go api

include "addition_management.thrift"
include "multiplication_management.thrift"
include "division_management.thrift"

struct UpdateReq {
    1: string idl(api.body = 'idl')
}

struct UpdateResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

service Gateway {
   addition_management.AdditionResponse addNumbers(1: addition_management.AdditionRequest req) (api.post="/add");
   multiplication_management.MultiplicationResponse multiplyNumbers(1: multiplication_management.MultiplicationRequest req) (api.post="/mul");
   division_management.DivisionResponse divideNumbers(1: division_management.DivisionRequest req) (api.post="/div");
   UpdateResp Update(1: UpdateReq req)(api.get = '/update')
}
