package service_consul

import (
	capi "github.com/hashicorp/consul/api"
	"github.com/zhangcook/consul_api_rpc/unilt"
	"log"
)

type ConsulGrpcs struct {
	ID            string
	Name          string
	Tage          []string
	Address       string
	Http          string
	Interval      string
	Port          int
	ConsulAddress string
}

func (c *ConsulGrpcs) ConsulApi() {
	return
}
func (c *ConsulGrpcs) ConsulService() {
	Grpc(c)

}
func Grpc(c *ConsulGrpcs) {
	consul, err := unilt.NewConsul(c.ConsulAddress)
	if err != nil {
		log.Println("Consul连接失败:" + err.Error())
		return
	}
	data := unilt.ConsulData{
		ID:      c.ID,
		Name:    c.Name,
		Tags:    c.Tage,
		Port:    c.Port,
		Address: c.Address,
		Check: &capi.AgentServiceCheck{
			HTTP:     c.Http,
			Interval: c.Interval,
		},
	}
	err = consul.RegisterConsul(data)
	if err != nil {
		log.Println("Consul注册服务失败:" + err.Error())
		return
	}
	log.Println("Consul连接成功")
}
