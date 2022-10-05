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
	orderLists, err := l.svcCtx.OrderRpc.GetOrderListByUid(l.ctx, &order_service.UserId{
		Uid: in.Id,
	})

	if err != nil {
		return nil, err
	}

	if len(orderLists.OrderList) == 0 {
		return nil, errors.New("用户没用订单")
	}
	//
	//l.svcCtx.RedisModel.Set(l.ctx, "babalili", "value333", 4*time.Hour).Err()

	userInfo, err := l.svcCtx.UserInfoModel.FindList(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if len(*userInfo) == 0 {
		return nil, errors.New("没有此用户")
	}
	userI := *userInfo

	var result proto.UserOrderList

	orderList := make(map[int64]*proto.OrderInfo)
	for k, v := range orderLists.OrderList {
		var orderInfo proto.OrderInfo
		orderInfo.Id = v.Id
		orderInfo.Oid = v.Oid
		orderInfo.Uid = v.Uid
		orderInfo.Gid = v.Gid
		orderInfo.Price = v.Price
		orderList[k] = &orderInfo
	}

	//
	//orderInfo.Id = 7
	//orderInfo.Oid = 73811
	//orderInfo.Uid = 333
	//orderInfo.Gid = 663098
	//orderInfo.Price = 574.39
	//orderList[0] = &orderInfo
	//
	//orderInfo.Id = 17
	//orderInfo.Oid = 733441
	//orderInfo.Uid = 444
	//orderInfo.Gid = 33098
	//orderInfo.Price = 99.99
	//orderList[1] = &orderInfo

	result.Id = userI[0].Id
	result.Name = userI[0].Name
	result.Gender = userI[0].Gender
	result.OrderList = orderList

	return &result, nil
}
