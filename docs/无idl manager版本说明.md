# api-gateway

an api-gateway based on Cloudwego

团队名称：ynnek76

小组成员：黄炜、倪匀博、张宇杰、柏帅

#### 如何创建一个新的 division 微服务，并将其应用到 API 网关上。

1. 目录结构:
    - `hertz-http-server`文件夹包含接受HTTP请求的代码和API的主要业务代码，即网关的实现。
    - `utils`文件夹中包含文件`gateway_utils.go`，它提供了许多函数来帮助开发人员使用API网关。
    - `thrift-idl`文件夹中包含thrift接口定义语言文件。这些文件可以自动为 Hertz 和 Kitex 服务器生成脚手架代码。
    - `microservices`文件夹中
        - `addition-service`、`multiplication-service`文件夹包含了 两数之和、两数之积的RPC服务器的代码，也就是这些微服务的实现。现在我们将实现第三个服务`division-service`，来展示使用我们的API网关来构建一个新的微服务， 并添加它以将其与现有的微服务基础设施集成。
    - 在执行下面的步骤之前，请将我们的存储库克隆到您的本地计算机上，并确保您位于根目录下。
2. 创建一个提供求两数之商的功能的新 RPC 服务器(即division微服务)
    - `cd`到`thrift-idl`文件夹下
    - 创建一个名为`division_management.thrift`的新文件
    - 输入以下代码：

```thrift
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
```

- 上面的文件将被Kitex框架用来为你的RPC服务器生成脚手架代码。生成的代码为存储在子目录kitex_gen中。
- 然而，在此之前，您需要在`microservice`文件夹下创建一个名为`division-service`的新目录，其中将包含该服务的Kitex代码。
- 然后执行以下命令，该命令将基于您的thrift IDL文件生成相应代码：

```bash
kitex -module github.com/huangwei021230/api-gateway -service division ../thrift-idl/division-management.thrift
```

- 完成后，请打开你的`handler.go`文件并添加以下代码以实现division服务的逻辑

```go
package main

import (
	"context"
	"fmt"
	api "github.com/huangwei021230/api-gateway/microservices/division-service/kitex_gen/division/management"
	"strconv"
)

type DivisionManagementImpl struct{}

func (s *DivisionManagementImpl) DivideNumbers(ctx context.Context, req *api.DivisionRequest) (resp *api.DivisionResponse, err error) {
	firstNumFloat, err := strconv.ParseFloat(req.FirstNum, 64)
	if err != nil {
		return nil, err
	}

	secondNumFloat, err := strconv.ParseFloat(req.SecondNum, 64)
	if err != nil {
		return nil, err
	}

	finalQuotient := firstNumFloat / secondNumFloat

	return &api.DivisionResponse{
		Quotient: fmt.Sprintf("%.2f", finalQuotient),
	}, nil
}

```

- 打开`division-service`文件夹中的`main.go`文件并添加以下代码。
- 这段代码对etcd注册中心进行了初始化
    - 通过调用`etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})`，用于服务发现和注册。
    - 您每次都需要更改ServiceName字段，并且在您想要决定你希望服务运行在哪个端口上时， 需要在`api.NewServer()`方法中添加以下属性: `server.WithServiceAddr(&net.TCPAddr{Port: XXXX})`。`XXXX`是你希望它运行的端口号
- 整体而言，这段代码的主要功能是创建一个 Kitex 微服务服务器，监听本地的8891端口，并使用 etcd 注册中心进行服务注册和发现。它将实现`DivisionManagement`接口的请求处理委托给`DivisionManagementImpl`对象，从而对外提供微服务功能。

```go
package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	management "github.com/huangwei021230/api-gateway/microservices/division-service/kitex_gen/division/management/divisionmanagement"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	// initate new etcd registry at port 2379
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := management.NewServer(
		new(DivisionManagementImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Division"}),
		server.WithRegistry(r),
        // please edit your port number here
		server.WithServiceAddr(&net.TCPAddr{Port: 8891}),
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}

```

- 恭喜您现在已经完成了 RPC服务器的设置!

下一步是对http服务器(即API网关)进行轻微修改，以确保您的服务集成到它上。

- 进入`thrift-idl`文件夹中

```bash
cd ../../thrift-idl
```

- 打开`gateway.thrift`文件，并在`service Gateway`中添加以下代码：

```thrift
include "division_management.thrift"
......
......
division_management.DivisionResponse divideNumbers(1: division_management.DivisionRequest req) (api.post="/div");
```

- 创建新的`division-thrift`文件

```thirft
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
```

- cd到`hertz-http-request`文件夹。
- 运行命令`hz update -idl ../thrift-idl/gateway_api.thrift`
- 此命令将更新 API 网关，以便当任何请求被发送到 url/div 地址时， DivideNumbers函数将被执行。这只与 hertz http服务器相关，我们将在这里实现调用 RPC服务器上的 calculateLength 方法的业务逻辑。
- cd到`./biz/handler/api`，打开`gateway.go`文件
- 在`gateway_service.go`文件中，在底部您将看到未实现的 DivideNumbers 方法。 这和之前定义在`division-service`中的方法是不一样的。这个方法定义了当请求被发送到 `/length`地址时的处理。请将如下 DivideNumbers 的方法实现放入`gateway_service.go`中：

```go
// DivideNumbers .
// @router /divide [POST]
func DivideNumbers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req divManagement.DivisionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// create new client (with loadbalancing, service discovery capabilities) using utils.GenerateClient feature
	divisionClient, err := utils.GenerateClient("Division")
	if err != nil {
		panic(err)
	}

	// binding req params to RPC reqest struct (following the request format declared in RPC service IDL)
	reqRpc := &divisionService.DivisionRequest{
		FirstNum:  req.FirstNum,
		SecondNum: req.SecondNum,
	}

	// initate new RPC response struct (as declared in RPC service IDL). This response variable will be populated by MakeRpcRequst function
	var respRpc divisionService.DivisionResponse

	// calling MakeRpcRequest method declared in the utils package
	err = utils.MakeRpcRequest(ctx, divisionClient, "divideNumbers", reqRpc, &respRpc)
	if err != nil {
		panic(err)
	}

	// initating and repackaging RPC response into new HTTP DivisionResponse
	resp := &divManagement.DivisionResponse{
		Quotient: respRpc.Quotient,
	}

	// return to client as JSON HTTP response
	c.JSON(consts.StatusOK, resp)
}
```

- 解析请求参数：通过`c.BindAndValidate(&req)`将接收到的 HTTP 请求的 JSON 数据解析到`divManagement.DivisionRequest`结构体中，并进行参数验证。
- 创建远程微服务客户端：使用`utils.GenerateClient`函数创建一个名为"Division"的远程微服务客户端。这个函数提供了负载均衡和服务发现的功能，确保请求会被发送到可用的服务实例。
- 组装 RPC 请求：将接收到的数值参数转换为`divisionService.DivisionRequest`结构体，并准备将其发送给远程服务进行处理。
- 发起 RPC 请求：通过`utils.MakeRpcRequest`函数发起 RPC 请求，将准备好的`reqRpc`请求参数发送给远程微服务，并将计算结果存储在`respRpc`中。
- 封装 RPC 响应：将得到的 RPC 响应`respRpc`中的计算结果（商）提取出来，并封装成`divManagement.DivisionResponse`结构体，以便返回给 HTTP 客户端。
- 返回 HTTP 响应：将封装好的`resp`作为 JSON 格式的 HTTP 响应返回给客户端。



​	现在你完成了API网关的编辑!