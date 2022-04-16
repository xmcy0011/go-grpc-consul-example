package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"test-consul/api/user"
	"test-consul/app/user/rpc"
	"test-consul/pkg/registry"
)

var (
	serverName    = "user"
	consulAddress = "127.0.0.1:8500"
	port          = flag.Int("port", 8000, "-port=8000")
)

func main() {
	flag.Parse()

	endpoint := fmt.Sprintf("%s:%d", getIntranetIP(), *port)
	log.Println("listen on :", endpoint)

	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Println(err)
		return
	}

	s := grpc.NewServer()
	user.RegisterUserServer(s, &rpc.UserService{})

	// 服务注册
	err = registry.Init(consulAddress)
	if err != nil {
		log.Println(err)
		return
	}
	svc := registry.ServerInstance{
		ID:       endpoint,
		Name:     serverName,
		Endpoint: endpoint,
	}
	err = registry.Registry(&svc, registry.Client)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("server register on ", consulAddress)
	// 退出前注销
	defer registry.DeRegister(&svc, registry.Client)

	// grpc
	log.Println("listen on :" + strconv.Itoa(*port))
	if err = s.Serve(listen); err != nil {
		log.Println(err.Error())
	}
}

func getIntranetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
