package service

import (
	"context"
	"gospacex/order-service/basic/config"
	"gospacex/order-service/model"
	__ "gospacex/proto"
	"net/http"
)

type Server struct {
	__.UnimplementedStreamGreeterServer
}

func (s *Server) OrderCreate(_ context.Context, in *__.OrderCreateReq) (*__.OrderCreateResp, error) {
	var order model.Order
	err := order.FindOrder(config.DB, in.OrderNo)
	if err == nil {
		return &__.OrderCreateResp{
			Code: http.StatusBadRequest,
			Msg:  "订单已存在",
		}, nil
	}
	order = model.Order{
		OrderNo:        in.OrderNo,
		UserID:         in.UserID,
		OrderStatus:    int8(in.OrderStatus),
		PaymentStatus:  int8(in.PaymentStatus),
		TotalAmount:    in.TotalAmount,
		PaidAmount:     in.PaidAmount,
		ConsigneeName:  in.ConsigneeName,
		ConsigneePhone: in.ConsigneePhone,
		Address:        in.Address,
		Remark:         in.Remark,
		ShipmentNo:     in.ShipmentNo,
		ExpressName:    in.ExpressName,
	}
	err = order.OrderCreate(config.DB)
	if err != nil {
		return &__.OrderCreateResp{
			Code: http.StatusBadRequest,
			Msg:  "订单添加失败",
		}, nil
	}
	return &__.OrderCreateResp{
		Code: http.StatusOK,
		Msg:  "订单添加成功",
	}, nil
}
func (s *Server) OrderDel(_ context.Context, in *__.OrderDelReq) (*__.OrderDelResp, error) {
	var order model.Order
	err := order.FindOrderId(config.DB, in.Id)
	if err != nil {
		return &__.OrderDelResp{
			Code: http.StatusBadRequest,
			Msg:  "订单不存在",
		}, nil
	}
	err = order.OrderDel(config.DB, in.Id)
	if err != nil {
		return &__.OrderDelResp{
			Code: http.StatusBadRequest,
			Msg:  "订单删除失败",
		}, nil
	}
	return &__.OrderDelResp{
		Code: http.StatusOK,
		Msg:  "订单删除成功",
	}, nil
}
func (s *Server) OrderUpdate(_ context.Context, in *__.OrderUpdateReq) (*__.OrderUpdateResp, error) {
	var order model.Order
	err := order.FindOrderId(config.DB, in.Id)
	if err != nil {
		return &__.OrderUpdateResp{
			Code: http.StatusBadRequest,
			Msg:  "订单不存在",
		}, nil
	}
	order = model.Order{
		OrderNo:        in.OrderNo,
		UserID:         in.UserID,
		OrderStatus:    int8(in.OrderStatus),
		PaymentStatus:  int8(in.PaymentStatus),
		TotalAmount:    in.TotalAmount,
		PaidAmount:     in.PaidAmount,
		ConsigneeName:  in.ConsigneeName,
		ConsigneePhone: in.ConsigneePhone,
		Address:        in.Address,
		Remark:         in.Remark,
		ShipmentNo:     in.ShipmentNo,
		ExpressName:    in.ExpressName,
	}
	err = order.OrderUpdate(config.DB, in.Id)
	if err != nil {
		return &__.OrderUpdateResp{
			Code: http.StatusBadRequest,
			Msg:  "订单修改失败",
		}, nil
	}
	return &__.OrderUpdateResp{
		Code: http.StatusOK,
		Msg:  "订单修改成功",
	}, nil
}
func (s *Server) OrderList(_ context.Context, in *__.OrderListReq) (*__.OrderListResp, error) {
	var order model.Order
	list, err := order.FindOrderList(config.DB, in)
	if err != nil {
		return &__.OrderListResp{
			Code: http.StatusBadRequest,
			Msg:  "订单列表获取失败",
		}, nil
	}
	var lists []*__.Order
	for _, o := range list {
		lists = append(lists, &__.Order{
			OrderNo:        o.OrderNo,
			UserID:         o.UserID,
			OrderStatus:    int64(o.OrderStatus),
			PaymentStatus:  int64(o.PaymentStatus),
			TotalAmount:    o.TotalAmount,
			PaidAmount:     o.PaidAmount,
			ConsigneeName:  o.ConsigneeName,
			ConsigneePhone: o.ConsigneePhone,
			Address:        o.Address,
			Remark:         o.Remark,
			ShipmentNo:     o.ShipmentNo,
			ExpressName:    o.ExpressName,
		})
	}
	return &__.OrderListResp{
		List: lists,
		Code: http.StatusOK,
		Msg:  "列表获取成功",
	}, nil
}
