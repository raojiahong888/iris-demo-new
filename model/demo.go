package model

import (
	"iris-demo-new/commons"
	"iris-demo-new/datasource"
	"iris-demo-new/entity/request"
	"iris-demo-new/entity/response"
	"iris-demo-new/slog"
)

type IDemoModel interface {
	List(params *request.DemoListRequest) (data response.DemoPageData, errCode commons.ResponseCode)
	Create(params *request.CreateDemoRequest) (id uint64, errCode commons.ResponseCode)
}

type demo struct {

}

func NewDemoModel() IDemoModel {
	return &demo{}
}

// get demo list
func (m *demo) List(params *request.DemoListRequest) (data response.DemoPageData, errCode commons.ResponseCode) {
	errCode = commons.OK
	CurrentPage := 1
	PageSize := 10
	if params.Page > 0 {
		CurrentPage = params.Page
	}
	if params.PageSize > 0 {
		PageSize = params.PageSize
	}
	offset := (CurrentPage - 1) *PageSize
	query := datasource.DB().Table("demo")
	if params.Keywords != "" {
		query = query.Where("name like ?", params.Keywords+"%")
	}
	var total int64
	query.Count(&total)
	if total == 0 {
		return
	}
	var demoList []response.Demo
	res := query.Select("id,name").Offset(offset).Limit(PageSize).Order("id desc").Scan(&demoList)
	if res.Error != nil {
		slog.Errorf("get event list failed, %s", res.Error)
		return data, commons.UnKnowError
	}
	data.Total = total
	data.List = demoList
	return
}

// create demo data
func (m *demo) Create(params *request.CreateDemoRequest) (id uint64, errCode commons.ResponseCode) {
	createData := &response.Demo{
		Name: params.Name,
	}
	res := datasource.DB().Save(createData)
	if res.Error != nil {
		slog.Errorf("save demo data failed, %s", res.Error)
		return 0, commons.UnKnowError
	}
	return createData.ID, commons.OK
}