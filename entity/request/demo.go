package request

type DemoListRequest struct {
	Keywords string `json:"keywords" form:"keywords"`
	Page     int    `json:"page" form:"page" validate:"required"`           // current page
	PageSize int    `json:"page_size" form:"page_size" validate:"required"` // page size
}

type CreateDemoRequest struct {
	Name string `json:"name" validate:"required"`
}
