namespace go api

include "addition_management.thrift"
include "multiplication_management.thrift"
include "division_management.thrift"

service Gateway {
   addition_management.AdditionResponse addNumbers(1: addition_management.AdditionRequest req) (api.post="/add");
   multiplication_management.MultiplicationResponse multiplyNumbers(1: multiplication_management.MultiplicationRequest req) (api.post="/mul");
   division_management.DivisionResponse divideNumbers(1: division_management.DivisionRequest req) (api.post="/div");
}