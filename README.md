# iot-platform

> 基于 go-zero(微服务) 实现的物联网平台
> 
> B站视频地址：https://www.bilibili.com/video/BV13G4y1R71m

## 技术栈

+ 后端：go-zero、gorm
+ 前端：vue
+ 硬件：arduino、esp8266

## 安装

1. 搭建 [Golang](https://golang.google.cn/) 环境
2. 安装 goctl
```shell
go get -u github.com/zeromicro/go-zero/tools/goctl@latest 
```
3. 下载安装 [arduino](https://www.arduino.cc/en/donate/)
4. 搭建 ETCD 环境
```shell
# 参考如下，使用Docker安装ETCD
docker run -d --name Etcd-server \
    --network app-tier \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
```
5. 搭建 EMQX 环境，参考地址：https://www.emqx.io/downloads
```shell
docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx/emqx:5.0.12 
```
6. 开启 EMQX 认证，访问地址：http://192.168.1.8:18083/ ，默认的端口是 18083，根据自己的地址调整即可；默认用户名和密码为 admin/public
7. 修改 `define.go` 中的EmqxAddr\EmqxKey\EmqxSecret为自己在EMQX后台生成的密钥对

## 命令

+ 创建API服务

```shell
goctl api new 服务名称
# 1. 创建 user 服务
goctl api new user
# 2. 创建 admin 服务
goctl api new admin
# 3. 创建 open 服务
goctl api new open
```

+ 生成服务代码

```shell
goctl api go -api 服务名称.api -dir . -style go_zero
# 1. 生成 user api 服务代码
goctl api go -api user.api -dir . -style go_zero
# 2. 生成 admin api 服务代码
goctl api go -api admin.api -dir . -style go_zero
# 3. 生成 user rpc 服务代码
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
# 4. 生成 device rpc 服务代码
goctl rpc protoc device.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
# 5. 生成 open api 服务代码
goctl api go -api open.api -dir . -style go_zero
```

+ 启动服务

```shell
go run 服务名称.go -f 配置文件地址
# 1. 启动 user api 服务
go run user.go -f etc/user-api.yaml
# 2. 启动 admin api 服务
go run admin.go -f etc/admin-api.yaml
# 3. 启动 user rpc 服务
go run user.go -f etc/user.yaml
# 4. 启动 device rpc 服务
go run device.go -f etc/device.yaml
# 5. 启动 open api 服务
go run open.go -f etc/open-api.yaml
```

## 适用场景

共享单车、共享充电宝、外卖柜

## Topic

1. 心跳
```text
/sys/产品key/设备key/ping
```

2. 设备接受消息的订阅地址

```text
/sys/产品key/设备key/receive
```

## 相关说明

### 设备连接

1. 连接的客户端ID即为设备的Key，密码为md5(key+secret)

### 签名规则

对请求参数的key从小到大排序后，拼接key(不包括sign字段)所对应的参数值后求md5。
```go
// 1. 例如发送的参数如下所示
map[string]interface{}{
    "app_key":     "app_key",
    "product_key": "1",
    "device_key":  "device_key",
    "data":        "hello world",
    "sign":        "4d62a91d0588320d314001828da9e1db",
}
// 2. 则签名为： app_key的值 + data的值 + device_key的值 + product_key的值
md5("app_keyhello worlddevice_key1") = "4d62a91d0588320d314001828da9e1db"
```


## 功能模块

+ [ ] 用户模块
  + [x] 登录
+ [ ] 后台管理模块
  + [x] 设备管理
    + [x] 设备列表
    + [x] 创建、修改、删除设备
  + [x] 产品管理
    + [x] 产品列表
    + [x] 创建、修改、删除产品
+ [x] 开放平台模块
  + [x] 签名
  + [x] 发送消息
+ [x] 设备服务模块
  + [x] 设备状态管理
  + [x] 发送消息
  