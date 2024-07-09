package asset

import "github.com/dnevsky/veryGoodProject/internal/dto"

type GetAssetsDTO struct {
	dto.ServiceDTO `swaggerignore:"true"`
	dto.PaginationDTO
}
