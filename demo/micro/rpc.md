# RPC学习
## ProtoBuf
- Message - 描述一个请求或响应的消息格式
    - singular : 表示成员0 或 1个,一般省略不写
    - repeated : 表示该字段可以包含0-N个元素
- 字段表示符 - 每个消息的定义时都有为一个数值标签
- 常用数据类型 double float int32/64 bool string bytes
- Service 服务定义- 在service中定义一个RPC 的服务接口

### go-micro v3使用
- 害,学了半天还是弃坑吧 k8s
- go-micro：微服务开发库
- Micro：基于Go-micro开发的运行时工具集
- 安装 [官网](https://github.com/micro/micro)
```bazaar
go get github.com/micro/micro/v3

#docker
# install
docker pull micro/micro

# run it
docker run -p 8080-8081:8080-8081/tcp micro/micro server
```
- 二进制文件下载方式
```bazaar
# MacOS
curl -fsSL https://raw.githubusercontent.com/micro/micro/master/scripts/install.sh | /bin/bash

# Linux
wget -q  https://raw.githubusercontent.com/micro/micro/master/scripts/install.sh -O - | /bin/bash

# Windows
powershell -Command "iwr -useb https://raw.githubusercontent.com/micro/micro/master/scripts/install.ps1 | iex"
```

- 使用该工具生成 .micro.pd.go & .pb.go 文件 
```bazaar
# 当前目录下可以简写为,各个路径在后面自己的修改(. 即可)
# protoc --proto_path=. --micro_out=. --go_out=. 路径/xx.proto 
#注意,其中的--go_out和--micro_out参数指定了生成文件的输出路径，
#应与上一步我们修改的go_package保持一致

``` 
- 遇到问题
```bazaar
#'protoc-gen-micro' 不是内部或外部命令，也不是可运行的程序或批处理文件
#安装 插件工具 protoc-gen-go protoc-gen-micro(缺啥装啥)
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v3/cmd/protoc-gen-micro
```

- go代码里使用
```bazaar
# 导入这两部分
"github.com/micro/micro/v3"
proto "github.com/micro/services/helloworld/proto"
```

