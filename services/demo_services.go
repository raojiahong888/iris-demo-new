package services

import (
	"iris-demo-new/commons"
	"iris-demo-new/entity/request"
	"iris-demo-new/entity/response"
	"iris-demo-new/model"
)

type IDemoService interface {
	DemoList(params *request.DemoListRequest) (result response.DemoPageData, errCode commons.ResponseCode)
	CreateDemo(params *request.CreateDemoRequest) (id uint64, errCode commons.ResponseCode)
}

type demoService struct {
	demoModel model.IDemoModel
}

func NewDemoService() IDemoService {
	return &demoService{
		demoModel: model.NewDemoModel(),
	}
}

func (s *demoService) DemoList(params *request.DemoListRequest) (result response.DemoPageData, errCode commons.ResponseCode) {
	return s.demoModel.List(params)
}

func (s *demoService) CreateDemo(params *request.CreateDemoRequest) (id uint64, errCode commons.ResponseCode) {
	return s.demoModel.Create(params)
}
