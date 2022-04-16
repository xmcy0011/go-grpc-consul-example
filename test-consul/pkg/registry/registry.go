package registry

import (
	"context"
	"fmt"
	kratos "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"strconv"
	"time"
)

type ServerInstance struct {
	ID       string // 全局需唯一
	Name     string // 服务名
	Endpoint string // 监听ip端口
}

func Registry(instance *ServerInstance, client *kratos.Registry) error {
	// 使用kratos封装的consul sdk
	svcInfo := registry.ServiceInstance{
		ID:        instance.ID,
		Name:      instance.Name,
		Version:   strconv.FormatInt(time.Now().Unix(), 10),
		Metadata:  map[string]string{"app": "kratos"},
		Endpoints: []string{fmt.Sprintf("tcp://%s?isSecure=false", instance.Endpoint)},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return client.Register(ctx, &svcInfo)
}

func DeRegister(instance *ServerInstance, client *kratos.Registry) error {
	svcInfo := registry.ServiceInstance{
		ID:        instance.ID,
		Name:      instance.Name,
		Version:   strconv.FormatInt(time.Now().Unix(), 10),
		Metadata:  map[string]string{"app": "kratos"},
		Endpoints: []string{fmt.Sprintf("tcp://%s?isSecure=false", instance.Endpoint)},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return client.Deregister(ctx, &svcInfo)
}
