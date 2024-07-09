package asset

import "github.com/dnevsky/veryGoodProject/internal/dto"

type GetAssetDTO struct {
	dto.ServiceDTO `swaggerignore:"true"`
	Name           string
}
