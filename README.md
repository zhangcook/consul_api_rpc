# Consul-Api-Rpc

>调用包 `github.com/zhangcook/consul_api_rpc`

* 实现了consul的注册和发现 并且结合grpc实现了分布式服务的调用
* NewConsulService 实现实例化了Api的Consul
>示例 API
```

import (
	"github.com/zhangcook/consul_api_rpc"
	"github.com/zhangcook/consul_api_rpc/api_consul"
)

func Grpc() {
	grpc := new_consul.NewConsulService()
	grpc.ConsulAddress = "you-consul-address"
	grpc.Name = "electron-srv" // 要监测的consul的通道名称
	service := []api_consul.Service{ // 要发现的consul的服务
		{
			ServiceName:    "user-srv", // 服务名称
			ServiceAddress: "127.0.0.1:50051", // 服务路由(注意如果consul没有发现该服务名称下的路由将使用该路由，这是一个防护措施,可以不填)
		},
		{
			ServiceName:    "order-ser",
			ServiceAddress: "127.0.0.1:50052",
		},
	}
	grpc.Service = service
	grpc.ConsulApi()

}
```
> 示例RPC
```
func main() {
    userService:="127.0.0.1:50051"
    grpcs := new_consul.NewConsulGrpc()
    service:=service_consul.ConsulGrpcs{
        	ID            "user-srv"
	        Name          ""// consul通道名称
	        Tage          ""
	        Address       "127.0.0.1"
	        Http          ""//你的api健康检查路由
	        Interval      ""//检查数据例如 10s 10秒发送一次
	        Port          50051
	        ConsulAddress "you-consul-address"
    }
    grpcs=service
	grpcs.ConsulService()
	flag.Parse()
	lis, err := net.Listen("tcp",userService)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```