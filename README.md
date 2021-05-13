# hash4me
rpcx demo: 调用hashpower微服务，计算难度在10以内的hash


## demn环境运行

### 一、 启动hashpower微服务

下载[hashpower](https://github.com/rpcxio/hashpower),执行 `go run .`运行这个微服务。

这个微服务模拟比特币的挖矿算法，给定一个指定的难度，它能够挖出一个哈希值。

默认它监听`8972`端口。

### 二、 启动web服务

在本项目主文件夹中执行`go run backend/server.go`, 它会启动一个web程序,接收用户的挖矿难度，调用hashpower微服务计算哈希值，并显示结果。

默认监听端口`9981`， 你可以访问 http://localhost:9981/ 测试。


## 重新编译web前端

进入frontend文件夹，执行:

```
yarn install
yarn build
```

生成的前端代码位于 frontend/dist, 本项目默认已经生成了前端。