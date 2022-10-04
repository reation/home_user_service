package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	OrderRpc zrpc.RpcClientConf
	zrpc.RpcServerConf
	mysql.Config
}
