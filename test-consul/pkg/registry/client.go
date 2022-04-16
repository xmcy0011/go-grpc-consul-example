package registry

import (
	kratos "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
	"sync"
)

var (
	Client *kratos.Registry

	consul *api.Client
	once   sync.Once
)

func Init(consulAddr string) error {
	if Client == nil {
		var err error = nil
		once.Do(func() {
			consul, err = api.NewClient(&api.Config{Address: consulAddr})
			if err != nil {
				return
			}

			opts := []kratos.Option{
				kratos.WithHeartbeat(true),
				kratos.WithHealthCheck(true),
				kratos.WithHealthCheckInterval(5),
			}
			Client = kratos.New(consul, opts...)
		})
		return err
	}
	return nil
}
