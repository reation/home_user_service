package logic

import (
	"context"
	"errors"
	"time"

	"github.com/reation/home_user_service/internal/svc"
	"github.com/reation/home_user_service/proto/types/proto"

	order_service "github.com/reation/home_order_service/order"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOrderListLogic {
	return &GetUserOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserOrderListLogic) GetUserOrderList(in *proto.IdRequest) (*proto.UserOrderList, error) {
	// todo: add your logic here and delete this line
	order, err := l.svcCtx.OrderRpc.GetOrderListByUid(l.ctx, &order_service.UserId{
		Uid: 789,
	})

	if err != nil {
		return nil, err
	}

	if len(order.OrderList) == 0 {
		return nil, errors.New("用户没用订单")
	}

	l.svcCtx.RedisModel.Set(l.ctx, "babalili", "value333", 4*time.Hour).Err()

	userInfo, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if userInfo.Name == "" {
		return nil, errors.New("没有此用户")
	}

	var result proto.UserOrderList
	var orderInfo proto.OrderInfo
	orderList := make(map[int64]*proto.OrderInfo)
	orderInfo.Id = 7
	orderInfo.Oid = 73811
	orderInfo.Uid = userInfo.Id
	orderInfo.Gid = 663098
	orderInfo.Price = 574.39
	orderList[0] = &orderInfo

	orderInfo.Id = 17
	orderInfo.Oid = 733441
	orderInfo.Uid = userInfo.Id
	orderInfo.Gid = 33098
	orderInfo.Price = 99.99
	orderList[1] = &orderInfo

	result.Id = userInfo.Id
	result.Name = userInfo.Name
	result.Gender = userInfo.Gender
	result.OrderList = orderList

	return &result, nil
}
