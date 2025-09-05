package new_consul

import (
	"github.com/zhangcook/consul_api_rpc/api_consul"
	"github.com/zhangcook/consul_api_rpc/service_consul"
)

type Construct interface {
	ConsulApi()
	ConsulService()
}

func NewConsulService() api_consul.ConsulServiceApi {
	return api_consul.ConsulServiceApi{}
}

func NewConsulGrpc() service_consul.ConsulGrpcs {
	return service_consul.ConsulGrpcs{}
}
