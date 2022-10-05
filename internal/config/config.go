package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	OrderRpc zrpc.RpcClientConf
	zrpc.RpcServerConf
	Mysql struct {
		UserData struct {
			DataSourceName string
		}
	}
	RedisCluster struct {
		Host []string
		Pass string
	}
}
