namespace go division.management

struct DivisionRequest {
    1: required string FirstNum;
    2: required string SecondNum;
}

struct DivisionResponse {
    1: string Quotient;
}

service DivisionManagement {
    DivisionResponse divideNumbers(1: DivisionRequest req)(api.post = '/div');
}