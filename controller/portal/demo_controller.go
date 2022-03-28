package portal

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-demo-new/commons"
	"iris-demo-new/entity/request"
	"iris-demo-new/services"
)

type demoController struct {
	Ctx         iris.Context
	DemoService services.IDemoService
}

func NewDemoController() *demoController {
	return &demoController{
		DemoService: services.NewDemoService(),
	}
}

func (c *demoController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "list", "GetDemoList")
	b.Handle("POST", "create", "CreateDemo")
}

// @tags Demo
// @Summary get demo list
// @Accept json
// @Produce json
// @Param   keywords     query    string     false        "keywords"
// @Param   page     query    int     true        "page"
// @Param   page_size     query    int     true        "page_size"
// @Success 200
// @Router /jh/demo/list [get]
func (c *demoController) GetDemoList() {
	// read body params
	params := request.DemoListRequest{}
	if err := c.Ctx.ReadForm(&params); err != nil {
		c.Ctx.JSON(commons.ResponseError(commons.ParamsFormatError, err.Error()))
		return
	}
	// validate params
	if errMsg := commons.Validate(&params); errMsg != nil {
		c.Ctx.JSON(commons.ResponseError(commons.ParamsFormatError, errMsg.Error()))
		return
	}
	data, errCode := c.DemoService.DemoList(&params)
	if errCode != commons.OK {
		c.Ctx.JSON(commons.ResponseError(errCode))
		return
	}
	c.Ctx.JSON(commons.ResponseSuccess("success", data))
}

// @tags Demo
// @Summary create demo
// @Accept json
// @Produce json
// @Param data body request.CreateDemoRequest true "data"
// @Success 200
// @Router /jh/demo/create [post]
func (c *demoController) CreateDemo() {
	// read body params
	params := request.CreateDemoRequest{}
	if err := c.Ctx.ReadJSON(&params); err != nil {
		c.Ctx.JSON(commons.ResponseError(commons.ParamsFormatError, err.Error()))
		return
	}
	// validate params
	if errMsg := commons.Validate(&params); errMsg != nil {
		c.Ctx.JSON(commons.ResponseError(commons.ParamsFormatError, errMsg.Error()))
		return
	}
	data, errCode := c.DemoService.CreateDemo(&params)
	if errCode != commons.OK {
		c.Ctx.JSON(commons.ResponseError(errCode))
		return
	}
	c.Ctx.JSON(commons.ResponseSuccess("success", data))
}
