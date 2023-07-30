# IDL Manager管理平台功能测试

an api-gateway based on Cloudwego

团队名称：ynnek76

小组成员：黄炜、倪匀博、张宇杰、柏帅



## 一、背景说明

- 在项目中后期，我们为API网关添加了IDL Manager管理平台，用于对于IDL版本进行管理



## 二、注意事项

- 注意：请在 `api-gateway/hertz-server` 目录下启动Hertz Server

  ```bash
  go run .
  ```

- 在IDL管理平台进行服务创建时，请注意服务名需要与相应微服务保持匹配关系

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



## 三、功能记录

#### 1.通过 `list-all-service` 查询当前已有服务

<img src="https://img1.imgtp.com/2023/07/30/i2yJeQRv.png" style="zoom: 67%; float: left;" />

> add服务为之前测试时添加

#### 2.通过 `add-service` 添加服务，body中 `Name` 和 `Idl` 分别为服务名与对应IDL路径

<img src="https://img1.imgtp.com/2023/07/30/YSOuYR0Y.png" style="zoom: 67%; float: left;" />

> `Response`中`success`为`true`，表示成功添加

#### 3.再次通过 `list-all-service` 查询当前已有服务，发现 `mul` 已被添加

<img src="https://img1.imgtp.com/2023/07/30/VVoFmUHM.png" style="zoom: 67%; float: left;" />

#### 4.或可通过 `search-for-service` 查询指定服务，body中`serviceNme`为欲查询的服务名

<img src="https://img1.imgtp.com/2023/07/30/vc51x8Ff.png" style="zoom: 67%; float: left;" />

#### 5.通过 `gateway/serviceName/methosName` 发起请求，验证结果正确

<img src="https://img1.imgtp.com/2023/07/30/Rw8unYqE.png" style="zoom: 67%; float: left;" />

#### 6.通过 `delete-service` 删除指定服务，body中`serviceNme`为欲删除的服务名

<img src="https://img1.imgtp.com/2023/07/30/8cUZdFxY.png" style="zoom: 67%; float: left;" />

> `Response`中`success`为`true`，表示成功添加

#### 7.再次查询，发现服务确被删除

<img src="https://img1.imgtp.com/2023/07/30/uMrrI2oW.png" style="zoom: 67%; float: left;" />

#### 8.在没有响应服务的情况下，无response，符合IDL定义内容

<img src="https://img1.imgtp.com/2023/07/30/kkEjuQCn.png" style="zoom: 67%; float: left;" />

- 此时，由于使用panic对可能错误进行预处理，网关内部提示，在查询对应IDL文件时错误

  <img src="https://img1.imgtp.com/2023/07/30/t6pSBjYa.png" style="zoom:70%;float:left" />



