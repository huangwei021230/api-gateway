namespace go gateway
struct Service {
    1: string Name(api.body="Name"),
    2: string Idl(api.body="Idl"),
}

struct ServiceReq {
    1: string serviceName(api.body='serviceName'),
}

struct SuccessResp {
    1: bool success(api.body='success'),
    2: string msg(api.body='msg'),
}

service IdlService {
    Service SearchForService(1: ServiceReq serviceReq)(api.post = '/search-for-service')
    SuccessResp AddService(1: Service service)(api.post = '/add-service')
    SuccessResp DelService(1: ServiceReq serviceReq)(api.post = '/delete-service')
    SuccessResp EditService(1: Service service)(api.post = '/edit-service')
    list<Service> ListAllService()(api.post = '/list-all-service')
}