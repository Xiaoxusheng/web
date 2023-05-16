# 初始化
```txt
go mod tidy
```

```txt
go mod init
```

```txt
Go语言的包管理工具为go mod，以下是常用的go mod指令：
- 初始化模块：go mod init
- 添加依赖项：go get
- 下载依赖项：go mod download
- 更新依赖项：go get -u
- 清理未使用依赖项：go mod tidy
- 查看依赖项：go list -m all
- 编辑依赖项版本：go mod edit
- 转储模块依赖图：go mod graph
- 验证依赖项：go mod verify
- 生成vendor目录：go mod vendor
```

code参数解释
```txt
code
101   必填参数为空
201   参数检验失败
301
401   数据库出错
500   系统内部错误