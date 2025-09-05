package api_consul

import (
	"fmt"
	"github.com/zhangcook/consul_api_rpc/unilt"
	"google.golang.org/grpc"
	"log"
	"time"
)

type ConsulServiceApi struct {
	ConsulAddress string
	Name          string
	Service       []Service
	SleepTime     time.Duration
}

type Service struct {
	ServiceName    string
	ServiceAddress string
}

func (c *ConsulServiceApi) ConsulApi() {
	Grpc(c)
}

func (c *ConsulServiceApi) ConsulService() {
	return
}

func ServiceGrpc(address string) *grpc.ClientConn {
	newGrpc := unilt.NewGrpc(address)
	return newGrpc.RegisterGrpcApi()
}

func Grpc(c *ConsulServiceApi) {
	consul, err := unilt.NewConsul(c.ConsulAddress)
	if err != nil {
		return
	}

	log.Println("Consul连接成功")

	go func() {
		for {
			log.Println("==============CONSUL*GRPC======================")
			ConsulService, err := consul.GetConsulService()
			if err != nil {
				log.Println("|获取所有Consul服务失败:" + err.Error())
				return
			}
			if ConsulService == nil {
				log.Println("|Consul服务不存在")
				return
			}
			for serviceID, serviceAdd := range ConsulService {
				if serviceAdd.Service == c.Name {
					var ServiceAddress string
					for _, i2 := range c.Service {
						if serviceID == i2.ServiceName {
							ServiceAddress = fmt.Sprintf("%v:%v", serviceAdd.Address, serviceAdd.Port)
							log.Println("|获取"+i2.ServiceName+"服务成功:", fmt.Sprintf("%v:%v", serviceAdd.Address, serviceAdd.Port))

						} else {
							ServiceAddress = i2.ServiceAddress
							log.Println("|" + i2.ServiceName + "获取服务失败,使用默认地址进行连接")
						}
						log.Println("|服务端口:" + ServiceAddress)
						ServiceGrpc(ServiceAddress)
					}

				}
			}
			log.Println("===============================================")
			time.Sleep(c.SleepTime)
			log.Println("进入Consul心跳检查")
		}
	}()
}
