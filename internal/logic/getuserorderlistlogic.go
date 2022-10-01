package logic

import (
	"context"
	"errors"

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

	var result proto.UserOrderList
	var orderInfo proto.OrderInfo
	orderList := make(map[int64]*proto.OrderInfo)
	orderInfo.Id = 7
	orderInfo.Oid = 73811
	orderInfo.Uid = 44433
	orderInfo.Gid = 663098
	orderInfo.Price = 574.39
	orderList[0] = &orderInfo

	orderInfo.Id = 17
	orderInfo.Oid = 733441
	orderInfo.Uid = 311222
	orderInfo.Gid = 33098
	orderInfo.Price = 99.99
	orderList[1] = &orderInfo

	result.Id = 488
	result.Name = "reation"
	result.Gender = 0
	result.OrderList = orderList

	return &result, nil
}
