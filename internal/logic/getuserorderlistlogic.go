package logic

import (
	"context"
	"errors"

	"github.com/reation/home_user_service/internal/svc"
	"github.com/reation/home_user_service/proto/types/proto"

	"github.com/reation/home_order_service/order"
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
	order, err := l.svcCtx.OrderRpc.GetOrderListByUid(l.ctx, &order.UserId{
		Uid: 789,
	})

	if err != nil {
		return nil, err
	}

	if len(order.OrderList) == 0 {
		return nil, errors.New("用户没用订单")
	}

	var result proto.UserOrderList
	for k,_ := range order.OrderList {
		result.OrderList[k].Id = order.OrderList[k].Id
		result.OrderList[k].Oid = order.OrderList[k].Oid
		result.OrderList[k].Uid = order.OrderList[k].Uid
		result.OrderList[k].Gid = order.OrderList[k].Gid
		result.OrderList[k].Price = order.OrderList[k].Price
	}


	return &proto.UserOrderList{
		Id: in.Id,
		Name: "reation",
		Gender: 0,
		OrderList: result.OrderList,
	}, nil
}
