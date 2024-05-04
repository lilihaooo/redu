package request

type PaginationParam struct {
	Page     int `form:"page" json:"page" label:"页码"`
	PageSize int `form:"page_size" json:"page_size" validate:"oneof=0 5 10 20 30" label:"每页条数"`
}
