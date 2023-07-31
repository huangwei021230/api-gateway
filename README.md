# api-gateway-doc

> an api-gateway based on Cloudwego

团队名称：ynnek76

小组成员：黄炜、倪匀博、张宇杰、柏帅

[TOC]

## 项目结构

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

## 项目部署

### etcd

```bash
etcd --log-level debug
```

### hertz client

- in directory `hertz-server`

```bash
go run .
```

### services

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

### interfaces

#### idl managementPlatform

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

#### services

please use post man

```
POST:127.0.0.1/gateway/serviceName/methodName
```

## 项目测试

### 测试目的

本次测试主要为验证网关对于服务基本功能是否成功支持，各模块功能是否符合预期，并且通过Golang性能测试与`Apache Benchmark`压力测试初步验证网关性能，并为后期优化提供参考。

### 测试环境

| 类型 | 说明                                                       |
| ---- | ---------------------------------------------------------- |
| OS   | wsl2                                                       |
| CPU  | AMD Ryzen 7 5800H with Radeon Graphics 3.20 GHz，8核16线程 |

### 测试步骤

#### 服务启动

##### 启动etcd

```
etcd --log-level debug
```

##### 启动HTTP Server

```c
// 在 api-gateway/hertz-http-server 目录下
go run .
```

##### 启动响应微服务

```c
// 在相应微服务文件目录下，如 api-gateway/microserviceaddition-service 下
go run .
```

#### 功能测试

##### 测试过程

1. 发送一个包含 `FirstNum` 与 `SecondNum` 的 JSON 请求到服务端。
2. 读取并关闭响应体。
3. 检查响应体内结果是否与预期的计算结果匹配。

##### 测试指令

```sh
// ADD
go test test/unit_test/addition_test.go
// DIV
go test test/unit_test/division_test.go
// MUL
go test test/unit_test/multiplication_test.go
```

##### 测试用例

| FirstNum | SecondNum | ExpResult | Op   |
| -------- | --------- | --------- | ---- |
| 100      | 20        | 120       | ADD  |
| 10       | -10       | 0         | ADD  |
| 100      | 20        | 5         | DIV  |
| 0        | 20        | 0         | DIV  |
| 100      | 20        | 2000      | MUL  |
| 0        | 20        | 0         | MUL  |

##### 测试结果

<img src="https://img1.imgtp.com/2023/07/23/hbc8milU.png" alt="image-20230723210744866" style="zoom: 70%; float: left;" />

结果表明网关成功提供了对服务的支持，能够正确接受与响应POST请求，并且根据请求路由确认目标服务和方法，并且可以根据相应IDL文件与微服务中的处理逻辑完成相应业务

#### 性能测试

##### 测试过程

对于每项已有服务，分别进行了串行测试与并行测试，来验证代码的横向可扩展性

> 以Addition服务为例，`BenchmarkAddition` 和 `BenchmarkAdditionParallel`两个基准测试函数都使用了 http.Post 方法发送请求，并读取响应内容。测试场景均为发送相同的请求数据，并模拟多次循环调用的场景，用于测试接口的性能。

##### 测试指令

```sh
// ADD
go test -bench=. test/unit_test/addition_test.go
// DIV
go test -bench=. test/unit_test/division_test.go
// MUL
go test -bench=. test/unit_test/multiplication_test.go
```

##### 测试结果

串行测试

<img src="https://img1.imgtp.com/2023/07/23/IAhQ7qUU.png" alt="image-20230723220023490" style="zoom:88%;float:left" />

并行测试

<img src="https://img1.imgtp.com/2023/07/23/DVmbjbls.png" alt="image-20230723220041952" style="zoom:80%;float:left" />

#### 压力测试

##### 测试过程

使用工具`Apache Benchmark`，指令参考[Apache Benchmark](https://httpd.apache.org/docs/2.2/programs/ab.html)

##### 测试指令

```sh
ab -n 100000 -c 10 -T application/json -p test/data/addition.data http://127.0.0.1:8888/add
```

> - 参数说明
>
>   - `-n 1000` 表示执行 1000 次请求。
>
>   - `-c 10` 表示并发请求数量为 10。
>
>   - `-T 'application/json'` 表示提交的数据类型为 `application/json`。
>
>   - `-p test/data/addition.data` 表示使用 `test/data/addition.data` 文件中的数据作为 POST 请求的 body。

##### 测试结果

<img src="https://img1.imgtp.com/2023/07/23/AEoPGQkh.png" alt="image-20230723222402302" style="zoom:80%;float:left" />

测试结果表明：

- 可以发现99%的请求都在1ms内完成，最长耗时也在5ms以内
- 用户平均请求等待时间0.334ms
- 服务器平均请求等待时间0.033ms

#### 管理平台测试

##### 测试说明

在IDL管理平台进行服务创建时，请注意**服务名需要与相应微服务保持匹配关系**

- 例如，微服务etcd集成后，服务端创建服务名为`addition-sever`，则自定义service名需要与服务器端服务名前缀保持一致

- 原因：这与我们的代码实现与etcd发现逻辑有关

  - 在代码实现中，我们根据请求的serviceName字段创建客户端，并以此去匹配服务器端的服务名

    ```go
    // ...	
    serviceName := c.Param("service")
    // ...
    cli, err := genericclient.NewClient(serviceName, g, client.WithResolver(r),
    	client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()))
    if err != nil {
    	panic(err)
    }
    ```

  - 浏览etcd查询信息时我们发现，对于`mul`服务，etcd其实查询了`mul*`的所有服务

    <img src="https://img1.imgtp.com/2023/07/31/ztgfxZAh.png" style="zoom:67%;float:left" />

##### 测试过程

使用`Postman`工具测试

1. 通过`list-all-service`查询当前已有服务
2.  通过`add-service` 添加服务，body中 `Name` 和 `Idl` 分别为服务名与对应IDL路径
3. 通过 `list-all-service` 查询当前已有服务，查询是否增加新服务
4. 通过 `search-for-service` 查询指定服务，body中`serviceNme`为欲查询的服务名
5. 通过 `gateway/serviceName/methosName` 发起请求，验证结果正确
6. 通过 `delete-service` 删除指定服务，body中`serviceNme`为欲删除的服务名
7. 通过 `list-all-service` 查询当前已有服务，查询服务是否已被删除
8. 在没有响应服务的情况下，检查是否有response，以符合IDL定义内容

##### 测试结果

1. add服务为之前测试时添加

<img src="https://img1.imgtp.com/2023/07/30/i2yJeQRv.png" style="zoom: 67%; float: left;" />

2. `Response`中`success`为`true`，表示成功添加

<img src="https://img1.imgtp.com/2023/07/30/YSOuYR0Y.png" style="zoom: 67%; float: left;" />

3. `Response`中新增`mul`及其定义

<img src="https://img1.imgtp.com/2023/07/30/VVoFmUHM.png" style="zoom: 67%; float: left;" />

4. `Response`中成功返回`mul`

<img src="https://img1.imgtp.com/2023/07/30/vc51x8Ff.png" style="zoom: 67%; float: left;" />

5. `mul`返回正确结果

<img src="https://img1.imgtp.com/2023/07/30/Rw8unYqE.png" style="zoom: 67%; float: left;" />

6. `Response`中`success`为`true`，表示成功删除

<img src="https://img1.imgtp.com/2023/07/30/8cUZdFxY.png" style="zoom: 67%; float: left;" />

7. `Response`中不存在`mul`，删除成功

<img src="https://img1.imgtp.com/2023/07/30/uMrrI2oW.png" style="zoom: 67%; float: left;" />

8. 无响应，说明服务删除

<img src="https://img1.imgtp.com/2023/07/30/kkEjuQCn.png" style="zoom: 67%; float: left;" />

> 此时，由于使用panic对可能错误进行预处理，网关内部提示，在查询对应IDL文件时错误
>
> <img src="https://img1.imgtp.com/2023/07/30/t6pSBjYa.png" style="zoom:70%;float:left" />