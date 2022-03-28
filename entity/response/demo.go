package response

type Demo struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
}

type DemoPageData struct {
	Total int64 `json:"total"`
	List []Demo `json:"list"`
}

func (*Demo) TableName() string {
	return `demo`  //demo table name
}