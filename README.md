#整体框架
```
使用gin框架 ，完成服务
（1）新增注册与登录接口
（2）验证礼品码接口修改
```
#目录结构
```
├── PreTest                         #压力测试
│   ├── giftCode.py
│   └── report_giftCode.html
├── README.md                       #介绍
├── config                          #配置redis与mongodb文件
│   ├── mongodb.go
│   └── redis.go
├── ctrl                            #处理router的请求
│   └── routerCtrl.go
├── gift_test.go                    #单元测试代码
├── go.mod
├── handle                          #处理具体的业务逻辑
│   └── handle.go
├── main.go                         #代码入口
├── model                           #数据模型与操作数据库的方法
│   ├── DBoperation             
│   │   ├── mongodb.go
│   │   └── redisSer.go
│   ├── Protobuf
│   │   ├── result.pb.go
│   │   └── result.proto
│   ├── gift
│   │   └── gift.go
│   └── userInfo
│       └── userInfo.go
├── router                          #路由转发
│   └── router.go
└── util                            #工具方法
    └── getRandCode.go              #获取随机数
```

#代码逻辑分层  gift-MongoDB
| 层     | 文件夹|主要职责 |调用关系|
| :----: | :----|:---- | :-----|
|router  | /router|路由转发 |调用handle|
| ctrl   | /ctrl  |请求参数验证 处理请求后构造回复消息 | 调用handle|
|handle  | /handle|处理路由 |调用module util|
|module  | /module|数据模型 操作数据库 |被handle调用|
|util    | /util | 通用工具 | 被handle调用|

#接口设计

##1.注册与登录接口

```
接口地址 
/giftCode/login 
```
### 请求方式
GET
### 请求示例
```
http://127.0.0.1:8080/giftCode/login?id=nccKwM9O
```
### 参数  说明

``` 
id 类型string 用户的id
```

```
成功示例 
{
    "condition": "success",
    "data": {
        "Diamond": "60",
        "Gold": "300",
        "Uid": "nccKwM9O"
    }
}
```

##2.验证礼品码接口

```
接口地址 
/giftCode/VerGiftCode 
```
### 请求方式
GET
### 请求示例
```
http://127.0.0.1:8080/giftCode/VerGiftCode?giftCode=nAyUzwZh&usr=nccKwM9O
```
### 参数  说明
``` 
giftCode 类型string 此字段为需要查询的礼包码
usr 类型string 用户的id
```

```
成功示例 
{
    "condition": "pass",
    "data": "GgQIARBkGgQIAhAUIgUIARCsAiIECAIQPCoECAIQUCoFCAEQkAM="
}
```

#第三方库
## redis
```
用于  操作redis 
代码 https://github.com/go-redis/redis
```

## gin
```
用于  开发服务器框架
代码  https://github.com/gin-gonic/gin 
```

## mongodb
```
用于 连接mongdb数据库 并进行操作
代码  https://go.mongodb.org/mongo-driver/mongo
```

##
```
用于 通信数据的传输格式protobuf
代码  https://google.golang.org/protobuf      
```
