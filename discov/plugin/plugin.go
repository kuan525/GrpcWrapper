package plugin

import (
	"errors"
	"fmt"
	"grpcwrapper/config"
	"grpcwrapper/discov"
	"grpcwrapper/discov/etcd"
)

func GetDiscovInstance() (discov.Discovery, error) {
	name := config.GetDiscovName()
	switch name {
	case "etcd":
		return etcd.NewETCDRegister(etcd.WithEndpoints(config.GetDiscovEndpoints()))
	}
	return nil, errors.New(fmt.Sprintf("not exist plugin:%s", name))
}
