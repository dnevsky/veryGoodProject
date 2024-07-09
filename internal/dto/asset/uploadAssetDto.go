package asset

import "github.com/dnevsky/veryGoodProject/internal/dto"

type UploadAssetDTO struct {
	dto.ServiceDTO `swaggerignore:"true"`
	Name           string
	Body           []byte
}
