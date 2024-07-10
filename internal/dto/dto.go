package dto

type ServiceDTO struct {
	XUserID uint64
}

type PaginationDTO struct {
	Limit  int `form:"limit,default=50" json:"limit"`
	Page   int `form:"page,default=1" json:"page"`
	Offset int `swaggerignore:"true"`
}
