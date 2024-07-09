package asset

import "github.com/dnevsky/veryGoodProject/internal/dto"

type DeleteAssetDTO struct {
	dto.ServiceDTO `swaggerignore:"true"`
	Name           string
}
