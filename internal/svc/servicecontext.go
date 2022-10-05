package svc

import (
	go_redis "github.com/go-redis/redis/v8"
	"github.com/reation/home_order_service/order"
	"github.com/reation/home_user_service/internal/config"
	"github.com/reation/home_user_service/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"runtime"
	"time"
)

type ServiceContext struct {
	Config        config.Config
	OrderRpc      order.Order
	UserInfoModel model.UserInfoModel
	RedisModel    go_redis.ClusterClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.UserData.DataSourceName)

	return &ServiceContext{
		Config:        c,
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		UserInfoModel: model.NewUserInfoModel(conn),
		RedisModel:    GetRedisClusterConn(c),
	}
}

func GetRedisClusterConn(c config.Config) go_redis.ClusterClient {
	redisCluster := go_redis.NewClusterClient(&go_redis.ClusterOptions{
		Addrs:        c.RedisCluster.Host,
		Password:     c.RedisCluster.Pass,
		ReadOnly:     false,
		PoolSize:     20 * runtime.NumCPU(), // 连接池最大socket连接数，默认为5倍CPU数， 5 * runtime.NumCPU
		MinIdleConns: 10,                    //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量。
		//命令执行失败时的重试策略
		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时，-1表示取消读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
		IdleTimeout:  5 * time.Minute, //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:   0 * time.Second, //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
	})

	return *redisCluster
}
