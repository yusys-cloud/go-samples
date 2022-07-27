## Protocol Buffer编译器安装
``` 
//Linux
apt install -y protobuf-compiler
protoc --version

//Mac
brew install protobuf
protoc --version  # Ensure compiler version is 3+
```

``` 
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## .proto文件生成gRPC代码

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/hello.proto
```

## Links
- https://grpc.io/docs/languages/go/quickstart/