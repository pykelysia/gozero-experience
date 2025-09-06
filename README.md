# GoZero-Experience
## 技术栈
![](https://img.shields.io/badge/Language:-golang-blue)
![](https://img.shields.io/badge/use:-gozero-blue)
![](https://img.shields.io/badge/database:-mysql-blue)

## 项目介绍
一个体验性质的 Go-Zero 项目，包含了常见的功能模块，方便大家学习和参考。

## 项目参考
具体的编写过程可以看[这篇博客](https://dinglz.cn/p/golang%E5%BE%AE%E6%9C%8D%E5%8A%A1%E5%AE%9E%E6%88%98/)

需要注意的是这篇博客中有一些小错误：在这里指出

- 1.在安装工具 `protoc-gen-go` 和 `goctl` 时，文章使用的是 `go get` 命令，而实际应使用使用 `go install` 命令（虽然用这条命令也不一定能直接安装成功）。如果使用 `go install` 命令后，工具未正常安装，建议自行上网搜索解决方案。

- 2.在 "中心网关对接微服务" 中，有一步需要在 `api/internal/svc/servicecontext.go` 中添加代码,此处文章中写的是
```go
type ServiceContext struct {
    Config  config.Config
    Adder   adder.Adder          // 手动代码
    Checker checker.Checker      // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:  c,
        Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),         // 手动代码
        Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),   // 手动代码
    }
}
```
其中的 `adder` 和 `checker` 包名是错误的，正确的是 `addclient` 和 `checkclient`。（该错误可能是由于 `goctl` 版本更新导致的）
```go
type ServiceContext struct {
	Config  config.Config
	Adder   addclient.Add       // 修改后代码
	Checker checkclient.Check   // 修改后代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Adder: addclient.NewAdd(zrpc.MustNewClient(c.Add)),         // 修改后代码
		Checker: checkclient.NewCheck(zrpc.MustNewClient(c.Check)), // 修改后代码
	}
}
```

- 3.在 "rpc中处理数据库" 中，`Config` 结构体中还需要加入 `Cache` 属性，并在对应的 `etc/add(check).yaml` 文件中补充相关如下配置
```yaml
Cache:
    - Host: ip:port
```
否则在对应的 `scc/servicecontext.go` 文件中添加 `MySQL` 配置时会报错，提示 `Config` 结构体中没有 `Cache` 属性。