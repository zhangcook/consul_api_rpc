package unilt

import capi "github.com/hashicorp/consul/api"

type Consul struct {
	Address *capi.Client
}

type ConsulData struct {
	ID      string
	Name    string
	Tags    []string
	Port    int
	Address string
	Check   *capi.AgentServiceCheck
}

// NewConsul 实例化Consul
func NewConsul(address string) (Consul, error) {
	config := capi.DefaultConfig()
	config.Address = address
	client, err := capi.NewClient(config)
	return Consul{Address: client}, err
}

// RegisterConsul 注册Consul
func (c *Consul) RegisterConsul(data ConsulData) error {
	registration := &capi.AgentServiceRegistration{
		ID:      data.ID,
		Name:    data.Name,
		Tags:    data.Tags,
		Port:    data.Port,
		Address: data.Address,
		Check:   data.Check,
	}
	return c.Address.Agent().ServiceRegister(registration)
}

// GetConsulService 获取Consul下所有的服务
func (c *Consul) GetConsulService() (map[string]*capi.AgentService, error) {
	return c.Address.Agent().Services()
}
