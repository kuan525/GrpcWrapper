package plugin

import (
	"errors"
	"fmt"

	"github.com/kuan525/grpcwrapper/config"
	"github.com/kuan525/grpcwrapper/discov"
	"github.com/kuan525/grpcwrapper/discov/etcd"
)

func GetDiscovInstance() (discov.Discovery, error) {
	name := config.GetDiscovName()
	switch name {
	case "etcd":
		return etcd.NewETCDRegister(etcd.WithEndpoints(config.GetDiscovEndpoints()))
	}
	return nil, errors.New(fmt.Sprintf("not exist plugin:%s", name))
}
