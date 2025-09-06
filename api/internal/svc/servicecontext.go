package svc

import (
	"gozero-exp/api/internal/config"
	"gozero-exp/rpc/add/addclient"
	"gozero-exp/rpc/check/checkclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Adder   addclient.Add
	Checker checkclient.Check
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Adder: addclient.NewAdd(zrpc.MustNewClient(c.Add)),
		Checker: checkclient.NewCheck(zrpc.MustNewClient(c.Check)),
	}
}
