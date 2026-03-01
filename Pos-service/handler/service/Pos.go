package service

import (
	"context"
	"gospaacex/Pos-service/basic/config"
	"gospaacex/Pos-service/model"
	__ "gospaacex/proto"
	"net/http"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedStreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) PosCreate(_ context.Context, in *__.PosCreateReq) (*__.PosCreateResp, error) {
	var position model.Position
	err := position.FindTitle(config.DB, in.Title)
	if err == nil {
		return &__.PosCreateResp{
			Code: http.StatusBadRequest,
			Msg:  "职位已存在",
		}, nil
	}
	position = model.Position{
		Title:        in.Title,
		Company:      in.Company,
		Salary:       in.Salary,
		Location:     in.Location,
		Description:  in.Description,
		Requirements: in.Requirements,
	}
	err = position.PosCreate(config.DB)
	if err != nil {
		return &__.PosCreateResp{
			Code: http.StatusOK,
			Msg:  "职位添加失败",
		}, nil
	}
	return &__.PosCreateResp{
		Code: http.StatusOK,
		Msg:  "职位添加成功",
	}, nil
}
func (s *Server) PosDel(_ context.Context, in *__.PosDelReq) (*__.PosDelResp, error) {
	var position model.Position
	err := position.FindId(config.DB, in.Id)
	if err != nil {
		return &__.PosDelResp{
			Code: http.StatusBadRequest,
			Msg:  "职位不存在",
		}, nil
	}
	err = position.PosDel(config.DB, in.Id)
	if err != nil {
		return &__.PosDelResp{
			Code: http.StatusOK,
			Msg:  "职位删除失败",
		}, nil
	}
	return &__.PosDelResp{
		Code: http.StatusOK,
		Msg:  "职位删除成功",
	}, nil
}
func (s *Server) PosUpdate(_ context.Context, in *__.PosUpdateReq) (*__.PosUpdateResp, error) {
	var position model.Position
	err := position.FindId(config.DB, in.Id)
	if err != nil {
		return &__.PosUpdateResp{
			Code: http.StatusBadRequest,
			Msg:  "职位不存在",
		}, nil
	}
	position = model.Position{
		Title:        in.Title,
		Company:      in.Company,
		Salary:       in.Salary,
		Location:     in.Location,
		Description:  in.Description,
		Requirements: in.Requirements,
	}
	err = position.PosUpdate(config.DB, in.Id)
	if err != nil {
		return &__.PosUpdateResp{
			Code: http.StatusOK,
			Msg:  "职位修改失败",
		}, nil
	}
	return &__.PosUpdateResp{
		Code: http.StatusOK,
		Msg:  "职位修改成功",
	}, nil
}
func (s *Server) PosList(_ context.Context, in *__.PosListReq) (*__.PosListResp, error) {
	var position model.Position
	list, err := position.PosList(config.DB, in)
	if err != nil {
		return &__.PosListResp{
			Code: http.StatusBadRequest,
			Msg:  "职位列表获取失败",
		}, nil
	}
	var lists []*__.Position
	for _, p := range list {
		lists = append(lists, &__.Position{
			Id:           int64(p.ID),
			Title:        p.Title,
			Company:      p.Company,
			Salary:       p.Salary,
			Location:     p.Location,
			Description:  p.Description,
			Requirements: p.Requirements,
		})
	}
	return &__.PosListResp{
		List: lists,
		Code: http.StatusOK,
		Msg:  "职位列表获取成功",
	}, nil
}
