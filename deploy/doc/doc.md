### mysql
```
跨机器访问
$ docker exec -it mysql mysql -uroot -p
$ use mysql;
$ update user set host='%' where user='root';
$ FLUSH PRIVILEGES;
```

```
执行 deploy\sql\zshop.sql 和 genModel.sh
```

```
分别执行 app\xxx\cmd\api\desc\genApi.sh
```

```
# 生成rpc业务代码
# 【注】 需要安装下面3个插件
#       protoc >= 3.13.0 ， 如果没安装请先安装 https://github.com/protocolbuffers/protobuf，下载解压到$GOPATH/bin下即可，前提是$GOPATH/bin已经加入$PATH中
#       protoc-gen-go ，如果没有安装请先安装 go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#       protoc-gen-go-grpc  ，如果没有安装请先安装 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
#
#       如果有要使用grpc-gateway，也请安装如下两个插件 , 没有使用就忽略下面2个插件
#       go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
#       go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
#
# 1）goctl >= 1.3 进入"服务/cmd/rpc/pb"目录下，执行下面命令
#    goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero
#    去除proto中的json的omitempty
#    mac: sed -i "" 's/,omitempty//g' *.pb.go
#    linux: sed -i 's/,omitempty//g' *.pb.go
# 2）goctl < 1.3 进入"服务/cmd"目录下，执行下面命令
#    goctl rpc proto -src rpc/pb/*.proto -dir ./rpc --style=goZero
#    去除proto中的json的omitempty
#    mac: sed -i "" 's/,omitempty//g'  ./rpc/pb/*.pb.go
#    linux: sed -i 's/,omitempty//g'  ./rpc/pb/*.pb.go
```

```
分别执行 app\xxx\cmd\rpc\pb\genRpc.sh
```