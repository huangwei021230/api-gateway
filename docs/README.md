# api-gateway

an api-gateway based on Cloudwego

团队名称：ynnek76

小组成员：黄炜、倪匀博、张宇杰、柏帅

### 项目结构

```tree
.
├── docs
│   ├── 测试方案说明与测试数据.md
│   └── 无idl manager版本说明.md
├── go.mod
├── go.sum
├── hertz-http-server (无idl manager的hertz-server)
│   ├── biz
│   │   ├── handler
│   │   │   ├── api
│   │   │   │   └── gateway.go
│   │   │   ├── ping.go
│   │   │   └── utils
│   │   │       └── utils.go
│   │   ├── model
│   │   │   ├── addition
│   │   │   │   └── management
│   │   │   │       └── addition_management.go
│   │   │   ├── api
│   │   │   │   └── gateway_api.go
│   │   │   ├── division
│   │   │   │   └── management
│   │   │   │       └── division_management.go
│   │   │   └── multiplication
│   │   │       └── management
│   │   │           └── multiplication_management.go
│   │   └── router
│   │       ├── api
│   │       │   ├── gateway_api.go
│   │       │   └── middleware.go
│   │       └── register.go
│   ├── build.sh
│   ├── kitex_gen
│   │   └── api
│   │       ├── gateway
│   │       │   ├── client.go
│   │       │   ├── gateway.go
│   │       │   ├── invoker.go
│   │       │   └── server.go
│   │       ├── gateway_api.go
│   │       ├── k-consts.go
│   │       └── k-gateway_api.go
│   ├── main.go
│   ├── router_gen.go
│   ├── router.go
│   └── script
│       └── bootstrap.sh
├── hertz-server (带idl manager的hertz-server)
│   ├── biz
│   │   ├── handler
│   │   │   ├── gateway
│   │   │   │   └── idl_service.go
│   │   │   ├── idl
│   │   │   │   └── managementPlatform.go
│   │   │   ├── ping.go
│   │   │   ├── Router.go (路由跳转)
│   │   │   └── utils
│   │   │       └── utils.go
│   │   ├── model
│   │   │   └── gateway
│   │   │       └── gateway.go
│   │   └── router
│   │       ├── gateway
│   │       │   ├── gateway.go
│   │       │   └── middleware.go
│   │       └── register.go
│   ├── build.sh
│   ├── main.go
│   ├── router_gen.go
│   ├── router.go
│   └── script
│       └── bootstrap.sh
├── LICENSE
├── microservices (微服务集合实现)
│   ├── addition-service
│   │   ├── build.sh
│   │   ├── handler.go
│   │   ├── kitex_gen
│   │   │   └── addition
│   │   │       └── management
│   │   │           ├── additionmanagement
│   │   │           │   ├── additionmanagement.go
│   │   │           │   ├── client.go
│   │   │           │   ├── invoker.go
│   │   │           │   └── server.go
│   │   │           ├── addition_management.go
│   │   │           ├── k-addition_management.go
│   │   │           └── k-consts.go
│   │   ├── kitex_info.yaml
│   │   ├── main.go
│   │   └── script
│   │       └── bootstrap.sh
│   ├── division-service
│   │   ├── build.sh
│   │   ├── handler.go
│   │   ├── kitex_gen
│   │   │   └── division
│   │   │       └── management
│   │   │           ├── divisionmanagement
│   │   │           │   ├── client.go
│   │   │           │   ├── divisionmanagement.go
│   │   │           │   ├── invoker.go
│   │   │           │   └── server.go
│   │   │           ├── division_management.go
│   │   │           ├── k-consts.go
│   │   │           └── k-division_management.go
│   │   ├── kitex_info.yaml
│   │   ├── main.go
│   │   └── script
│   │       └── bootstrap.sh
│   └── multiplication-service
│       ├── build.sh
│       ├── handler.go
│       ├── kitex_gen
│       │   └── multiplication
│       │       └── management
│       │           ├── k-consts.go
│       │           ├── k-multiplication_management.go
│       │           ├── multiplicationmanagement
│       │           │   ├── client.go
│       │           │   ├── invoker.go
│       │           │   ├── multiplicationmanagement.go
│       │           │   └── server.go
│       │           └── multiplication_management.go
│       ├── kitex_info.yaml
│       ├── main.go
│       └── script
│           └── bootstrap.sh
├── README.md
├── test
│   ├── data
│   │   └── addition.data
│   ├── performance_test
│   │   ├── addition_test.go
│   │   ├── division_test.go
│   │   └── multiplication_test.go
│   └── unit_test
│       ├── addition_test.go
│       ├── division_test.go
│       ├── multiplication_test.go
│       └── utils
│           ├── addition_handler.go
│           ├── division_handler.go
│           └── multiplication_handler.go
├── thrift-idl
│   ├── addition_management.thrift
│   ├── division_management.thrift
│   ├── gateway_api.thrift
│   ├── gateway.thrift
│   └── multiplication_management.thrift
└── utils
    └── gateway_utils.go

```

### deployment

#### etcd

```bash
etcd --log-level debug
```

#### hertz client

- in directory `hertz-server`

```bash
go run .
```

#### services

- in directory `addition-service`

```
go run .
```

- in directory `division-service`

```bash
go run .
```

- in directory `multiplication-service`

```bash
go run .
```

#### interfaces

##### idl managementPlatform

```thrift
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
```

##### services

please use post man

```
POST:127.0.0.1/gateway/serviceName/methodName
```